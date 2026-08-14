package logic

import (
	"errors"
	"fmt"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	system "github.com/CloudSilk/CloudSilk/pkg/servers/system/logic"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 排程计划状态
const (
	ScheduleStateGenerated = "已生成"
	ScheduleStateReleased  = "已下发"
	ScheduleStateVoided    = "已作废"
)

// GenerateSchedule 为产线生成排程计划：
// 取产线下已发放/已签派/生产中的工单，按前向贪心算法（优先级→交期）分配时段，落库计划与明细
func GenerateSchedule(req *proto.GenerateScheduleRequest, userID string) (*proto.GenerateScheduleResponse, error) {
	if req.ProductionLineID == "" {
		return nil, errors.New("productionLineID不能为空")
	}

	startTime := time.Now()
	if req.StartTime != "" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", req.StartTime, time.Local)
		if err != nil {
			return nil, errors.New("startTime格式无效，应为yyyy-MM-dd HH:mm:ss")
		}
		startTime = t
	}

	//取待排产工单：已发放/已签派（生产中的按剩余数量顺延排入）
	orders := []*model.ProductOrder{}
	query := model.DB.DB().Preload("ProductModel").Where("`production_line_id` = ? AND `current_state` in ?",
		req.ProductionLineID,
		[]string{types.ProductOrderStateReleased, types.ProductOrderStateDispatched, types.ProductOrderStateProducting})
	if len(req.ProductOrderIDs) > 0 {
		query = query.Where("`id` in (?)", req.ProductOrderIDs)
	}
	if err := query.Find(&orders).Error; err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, errors.New("此产线下没有可排程的工单（需为已发放/已签派状态）")
	}

	schedOrders := make([]SchedOrderEx, 0, len(orders))
	for _, o := range orders {
		std := o.StandardWorkTime
		if std <= 0 {
			return nil, fmt.Errorf("工单%s缺少标准节拍，无法排程（请先完成发放以匹配生产节拍）", o.ProductOrderNo)
		}
		qty := o.OrderQTY
		if o.CurrentState == types.ProductOrderStateProducting && o.StartedQTY > 0 {
			qty = o.OrderQTY - o.StartedQTY
			if qty <= 0 {
				continue
			}
		}
		so := SchedOrderEx{
			ProductModel: orderProductModel(o),
		}
		so.OrderID = o.ID
		so.OrderNo = o.ProductOrderNo
		so.QTY = qty
		so.StdWorkTimeSec = std
		so.PriorityLevel = o.PriorityLevel
		if o.DeliveryDate.Valid {
			so.DeliveryDate = o.DeliveryDate.Time
			so.HasDelivery = true
		}
		//物料齐套时间：取该工单最近一张已完成拣货单的时间（未领料视为未齐套，不限制开工时间）
		if ready, ok := orderMaterialReadyTime(o.ID); ok {
			so.ReadyTime = ready
			so.HasReady = true
		}
		schedOrders = append(schedOrders, so)
	}

	slots := ConstrainedSchedule(schedOrders, startTime, SchedOptions{
		DailyWorkMinutes: req.DailyWorkMinutes,
		ChangeoverSec:    req.ChangeoverSeconds,
	})
	if len(slots) == 0 {
		return nil, errors.New("排程结果为空")
	}

	planNo, err := system.GenerateSerialNumber("排程批次号", "APS排程计划流水号", fmt.Sprintf("SC%s", startTime.Format("20060102")), 4, 1)
	if err != nil {
		return nil, err
	}

	plan := &model.ProductionSchedulePlan{
		PlanNo:           planNo,
		ProductionLineID: req.ProductionLineID,
		StartTime:        slots[0].Start,
		EndTime:          slots[len(slots)-1].End,
		OrderCount:       int32(len(slots)),
		CurrentState:     ScheduleStateGenerated,
	}
	for i, slot := range slots {
		plan.Items = append(plan.Items, &model.ProductionScheduleItem{
			ProductOrderID:   slot.OrderID,
			Sequence:         int32(i + 1),
			PlannedStartTime: slot.Start,
			PlannedEndTime:   slot.End,
			DurationSeconds:  slot.Seconds,
			PriorityLevel:    findPriority(orders, slot.OrderID),
		})
	}

	if err := model.DB.DB().Create(plan).Error; err != nil {
		return nil, err
	}

	resp := &proto.GenerateScheduleResponse{
		Code:      proto.Code_Success,
		PlanID:    plan.ID,
		PlanNo:    plan.PlanNo,
		OrderCount: plan.OrderCount,
		EndTime:   plan.EndTime.Format("2006-01-02 15:04:05"),
	}
	return resp, nil
}

// ReleaseSchedule 下发排程：回写各工单的预计开工/完工时间（甘特基准）
func ReleaseSchedule(req *proto.ReleaseScheduleRequest, userID string) error {
	if req.PlanID == "" {
		return errors.New("planID不能为空")
	}

	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		plan := &model.ProductionSchedulePlan{}
		if err := tx.Preload("Items").First(plan, "`id` = ?", req.PlanID).Error; err != nil {
			return errors.New("读取排程计划失败")
		}
		if plan.CurrentState != ScheduleStateGenerated {
			return fmt.Errorf("排程计划状态为%s，只有%s状态可以下发", plan.CurrentState, ScheduleStateGenerated)
		}

		for _, item := range plan.Items {
			if err := tx.Model(&model.ProductOrder{}).Where("`id` = ?", item.ProductOrderID).
				Updates(map[string]interface{}{
					"estimate_start_time": item.PlannedStartTime,
					"estimate_finish_time": item.PlannedEndTime,
				}).Error; err != nil {
				return err
			}
		}

		if err := tx.Model(&model.ProductionSchedulePlan{}).Where("`id` = ?", plan.ID).
			Update("current_state", ScheduleStateReleased).Error; err != nil {
			return err
		}
		return tx.Create(&model.OperationTrace{
			OperateUserID:  userID,
			ControllerName: "APS排程",
			ActionName:     "下发排程",
			RequestContent: fmt.Sprintf("批次:%s,工单数:%d,计划区间:%s→%s", plan.PlanNo, plan.OrderCount, plan.StartTime.Format("01-02 15:04"), plan.EndTime.Format("01-02 15:04")),
		}).Error
	})
}

// VoidSchedule 作废排程计划
func VoidSchedule(planID string) error {
	plan := &model.ProductionSchedulePlan{}
	if err := model.DB.DB().First(plan, "`id` = ?", planID).Error; err != nil {
		return errors.New("读取排程计划失败")
	}
	if plan.CurrentState == ScheduleStateReleased {
		return errors.New("已下发的排程计划不能作废")
	}
	return model.DB.DB().Model(&model.ProductionSchedulePlan{}).Where("`id` = ?", planID).
		Update("current_state", ScheduleStateVoided).Error
}

func QueryProductionSchedulePlan(req *proto.QueryProductionSchedulePlanRequest, resp *proto.QueryProductionSchedulePlanResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductionSchedulePlan{})
	if preload {
		db = db.Preload("ProductionLine").Preload("Items").Preload("Items.ProductOrder")
	}
	if req.ProductionLineID != "" {
		db = db.Where("`production_line_id` = ?", req.ProductionLineID)
	}
	if req.CurrentState != "" {
		db = db.Where("`current_state` = ?", req.CurrentState)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "create_time desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductionSchedulePlan
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionSchedulePlansToPB(list)
	}
	resp.Total = resp.Records
}

func GetProductionSchedulePlanByID(id string) (*model.ProductionSchedulePlan, error) {
	m := &model.ProductionSchedulePlan{}
	err := model.DB.DB().Preload(clause.Associations).Preload("Items.ProductOrder").Where("`id` = ?", id).First(m).Error
	return m, err
}

func DeleteProductionSchedulePlan(id string) error {
	plan := &model.ProductionSchedulePlan{}
	if err := model.DB.DB().First(plan, "`id` = ?", id).Error; err != nil {
		return errors.New("读取排程计划失败")
	}
	if plan.CurrentState == ScheduleStateReleased {
		return errors.New("已下发的排程计划不能删除")
	}
	return model.DB.DB().Delete(&model.ProductionSchedulePlan{}, "`id` = ?", id).Error
}

func findPriority(orders []*model.ProductOrder, orderID string) int32 {
	for _, o := range orders {
		if o.ID == orderID {
			return o.PriorityLevel
		}
	}
	return 0
}

// orderProductModel 读取工单的产品型号代号（用于换型约束）
func orderProductModel(o *model.ProductOrder) string {
	if o.ProductModel != nil {
		return o.ProductModel.Code
	}
	return ""
}

// orderMaterialReadyTime 物料齐套时间：最近一张已完成拣货单的完成时间
func orderMaterialReadyTime(orderID string) (time.Time, bool) {
	bill := &model.WMSBillQueue{}
	err := model.DB.DB().Where("`product_order_id` = ? AND `current_state` = ?", orderID, "已完成").
		Order("last_update_time desc").First(bill).Error
	if err != nil {
		return time.Time{}, false
	}
	return bill.LastUpdateTime, true
}

// InsertSchedule 插单重排：在既有计划中插入新工单（保持既有顺序），重排后落库
func InsertSchedule(req *proto.InsertScheduleRequest) (*proto.InsertScheduleResponse, error) {
	if req.PlanID == "" || req.ProductOrderID == "" {
		return nil, errors.New("planID与productOrderID不能为空")
	}

	plan := &model.ProductionSchedulePlan{}
	if err := model.DB.DB().Preload("Items").Preload("Items.ProductOrder").First(plan, "`id` = ?", req.PlanID).Error; err != nil {
		return nil, errors.New("读取排程计划失败")
	}
	if plan.CurrentState != ScheduleStateGenerated {
		return nil, fmt.Errorf("排程计划状态为%s，只有%s状态可以插单重排", plan.CurrentState, ScheduleStateGenerated)
	}

	//读取新工单
	newOrder := &model.ProductOrder{}
	if err := model.DB.DB().Preload("ProductModel").First(newOrder, "`id` = ?", req.ProductOrderID).Error; err != nil {
		return nil, errors.New("读取新工单失败")
	}
	if newOrder.StandardWorkTime <= 0 {
		return nil, fmt.Errorf("工单%s缺少标准节拍，无法排程", newOrder.ProductOrderNo)
	}

	//重建既有工单的约束信息
	orders := make([]SchedOrderEx, 0, len(plan.Items))
	existing := make([]SchedSlot, 0, len(plan.Items))
	for _, item := range plan.Items {
		if item.ProductOrder == nil {
			continue
		}
		o := item.ProductOrder
		so := SchedOrderEx{ProductModel: orderProductModel(o)}
		so.OrderID = o.ID
		so.OrderNo = o.ProductOrderNo
		so.QTY = o.OrderQTY
		so.StdWorkTimeSec = o.StandardWorkTime
		so.PriorityLevel = item.PriorityLevel
		if o.DeliveryDate.Valid {
			so.DeliveryDate = o.DeliveryDate.Time
			so.HasDelivery = true
		}
		orders = append(orders, so)
		existing = append(existing, SchedSlot{
			OrderID: o.ID,
			OrderNo: o.ProductOrderNo,
			Start:   item.PlannedStartTime,
			End:     item.PlannedEndTime,
			Seconds: item.DurationSeconds,
		})
	}

	newSo := SchedOrderEx{ProductModel: orderProductModel(newOrder)}
	newSo.OrderID = newOrder.ID
	newSo.OrderNo = newOrder.ProductOrderNo
	newSo.QTY = newOrder.OrderQTY
	newSo.StdWorkTimeSec = newOrder.StandardWorkTime
	newSo.PriorityLevel = newOrder.PriorityLevel
	if newOrder.DeliveryDate.Valid {
		newSo.DeliveryDate = newOrder.DeliveryDate.Time
		newSo.HasDelivery = true
	}
	if ready, ok := orderMaterialReadyTime(newOrder.ID); ok {
		newSo.ReadyTime = ready
		newSo.HasReady = true
	}

	slots, idx := InsertOrder(existing, orders, newSo, plan.StartTime, SchedOptions{
		ChangeoverSec: req.ChangeoverSeconds,
	})
	if idx < 0 {
		return nil, errors.New("插单失败，未能确定插入位置")
	}

	resp := &proto.InsertScheduleResponse{
		Code:     proto.Code_Success,
		Sequence: int32(idx + 1),
		EndTime:  slots[len(slots)-1].End.Format("2006-01-02 15:04:05"),
	}

	//重建明细
	return resp, model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("`production_schedule_plan_id` = ?", plan.ID).
			Delete(&model.ProductionScheduleItem{}).Error; err != nil {
			return err
		}
		priorityOf := func(orderID string) int32 {
			for _, o := range orders {
				if o.OrderID == orderID {
					return o.PriorityLevel
				}
			}
			if newSo.OrderID == orderID {
				return newSo.PriorityLevel
			}
			return 0
		}
		items := make([]*model.ProductionScheduleItem, 0, len(slots))
		for i, slot := range slots {
			items = append(items, &model.ProductionScheduleItem{
				ProductionSchedulePlanID: plan.ID,
				ProductOrderID:           slot.OrderID,
				Sequence:                 int32(i + 1),
				PlannedStartTime:         slot.Start,
				PlannedEndTime:           slot.End,
				DurationSeconds:          slot.Seconds,
				PriorityLevel:            priorityOf(slot.OrderID),
			})
		}
		if err := tx.Create(&items).Error; err != nil {
			return err
		}
		return tx.Model(&model.ProductionSchedulePlan{}).Where("`id` = ?", plan.ID).
			Updates(map[string]interface{}{
				"end_time":    slots[len(slots)-1].End,
				"order_count": int32(len(slots)),
			}).Error
	})
}
