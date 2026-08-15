package logic

import (
	"errors"
	"fmt"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 设备状态
const (
	EquipmentStateInUse      = "在用"
	EquipmentStateStopped    = "停用"
	EquipmentStateMaintaining = "维修中"
	EquipmentStateScrapped   = "报废"
)

func CreateEquipment(m *model.Equipment) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " `code`  = ? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同设备代号")
	}
	if m.CurrentState == "" {
		m.CurrentState = EquipmentStateInUse
	}
	return m.ID, nil
}

// UpdateEquipment 白名单更新：仅写显式提供的字段，避免全量 Save 把漏传字段清零
//（设备状态走独立的 changestate 接口，此处不接受状态修改）
func UpdateEquipment(m *model.Equipment) error {
	if m.ID == "" {
		return errors.New("id不能为空")
	}
	updates := map[string]interface{}{}
	if m.Code != "" {
		updates["code"] = m.Code
	}
	if m.Name != "" {
		updates["name"] = m.Name
	}
	if m.Model != "" {
		updates["model"] = m.Model
	}
	if m.Manufacturer != "" {
		updates["manufacturer"] = m.Manufacturer
	}
	if m.SerialNo != "" {
		updates["serial_no"] = m.SerialNo
	}
	if m.Location != "" {
		updates["location"] = m.Location
	}
	if m.Team != "" {
		updates["team"] = m.Team
	}
	if m.CommissionDate.Valid {
		updates["commission_date"] = m.CommissionDate
	}
	if m.ProductionStationID != nil && *m.ProductionStationID != "" {
		updates["production_station_id"] = *m.ProductionStationID
	}
	updates["enable"] = m.Enable
	updates["remark"] = m.Remark

	if m.Code != "" {
		var count int64
		if err := model.DB.DB().Model(&model.Equipment{}).
			Where("`id` != ? AND `code` = ?", m.ID, m.Code).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return errors.New("存在相同设备代号")
		}
	}
	if len(updates) <= 2 { // 仅 enable/remark 默认项
		updates["enable"] = m.Enable
	}
	return model.DB.DB().Model(&model.Equipment{}).Where("`id` = ?", m.ID).Updates(updates).Error
}

func QueryEquipment(req *proto.QueryEquipmentRequest, resp *proto.QueryEquipmentResponse, preload bool) {
	db := model.DB.DB().Model(&model.Equipment{})
	if preload {
		db = db.Preload(clause.Associations)
	}
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `name` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}
	if req.CurrentState != "" {
		db = db.Where("`current_state` = ?", req.CurrentState)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.Equipment
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.EquipmentsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllEquipments() (list []*model.Equipment, err error) {
	err = model.DB.DB().Preload(clause.Associations).Find(&list, "`enable` = ?", true).Error
	return
}

func GetEquipmentByID(id string) (*model.Equipment, error) {
	m := &model.Equipment{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func DeleteEquipment(id string) (err error) {
	var count int64
	if err := model.DB.DB().Model(&model.EquipmentMaintenancePlan{}).Where("`equipment_id` = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("此设备存在维保计划，无法删除")
	}
	if err := model.DB.DB().Model(&model.EquipmentMaintenanceRecord{}).Where("`equipment_id` = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("此设备存在维保历史记录，无法删除（可停用）")
	}
	return model.DB.DB().Delete(&model.Equipment{}, "`id` = ?", id).Error
}

// ChangeEquipmentState 设备状态流转（在用/停用/维修中/报废）
func ChangeEquipmentState(id, state string) error {
	valid := map[string]bool{EquipmentStateInUse: true, EquipmentStateStopped: true, EquipmentStateMaintaining: true, EquipmentStateScrapped: true}
	if !valid[state] {
		return fmt.Errorf("无效的设备状态：%s", state)
	}
	return model.DB.DB().Model(&model.Equipment{}).Where("`id` = ?", id).Update("current_state", state).Error
}

// ---------- 维保计划 ----------

func CreateEquipmentMaintenancePlan(m *model.EquipmentMaintenancePlan) (string, error) {
	if m.EquipmentID == nil || *m.EquipmentID == "" {
		return "", errors.New("设备不能为空")
	}
	if m.CycleDays <= 0 {
		return "", errors.New("维保周期必须大于0天")
	}
	// 首次创建：未指定下次执行日期时按当前日期+周期滚动
	if !m.NextExecution.Valid {
		m.NextExecution.Valid = true
		m.NextExecution.Time = time.Now().AddDate(0, 0, int(m.CycleDays))
	}
	return m.ID, model.DB.DB().Create(m).Error
}

func UpdateEquipmentMaintenancePlan(m *model.EquipmentMaintenancePlan) error {
	omits := []string{"created_at"}
	if m.EquipmentID == nil || *m.EquipmentID == "" {
		omits = append(omits, "EquipmentID")
	}
	if m.CycleDays <= 0 {
		return errors.New("维保周期必须大于0天")
	}
	//周期变更时按新周期重算下次执行（基准：上次执行时间，未执行过则当前时间）
	old := &model.EquipmentMaintenancePlan{}
	if err := model.DB.DB().First(old, "`id` = ?", m.ID).Error; err != nil {
		return errors.New("读取维保计划失败")
	}
	if old.CycleDays != m.CycleDays {
		base := time.Now()
		if old.LastExecution.Valid {
			base = old.LastExecution.Time
		}
		m.NextExecution.Valid = true
		m.NextExecution.Time = base.AddDate(0, 0, int(m.CycleDays))
	}
	return model.DB.DB().Omit(omits...).Save(m).Error
}

func QueryEquipmentMaintenancePlan(req *proto.QueryEquipmentMaintenancePlanRequest, resp *proto.QueryEquipmentMaintenancePlanResponse, preload bool) {
	db := model.DB.DB().Model(&model.EquipmentMaintenancePlan{})
	if preload {
		db = db.Preload(clause.Associations)
	}
	if req.EquipmentID != "" {
		db = db.Where("`equipment_id` = ?", req.EquipmentID)
	}
	if req.DueOnly {
		db = db.Where("`enable` = ? AND `next_execution` IS NOT NULL AND `next_execution` <= ?", true, time.Now())
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.EquipmentMaintenancePlan
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.EquipmentMaintenancePlansToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllEquipmentMaintenancePlans() (list []*model.EquipmentMaintenancePlan, err error) {
	err = model.DB.DB().Preload(clause.Associations).Find(&list).Error
	return
}

func GetEquipmentMaintenancePlanByID(id string) (*model.EquipmentMaintenancePlan, error) {
	m := &model.EquipmentMaintenancePlan{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func DeleteEquipmentMaintenancePlan(id string) (err error) {
	return model.DB.DB().Delete(&model.EquipmentMaintenancePlan{}, "`id` = ?", id).Error
}

// ---------- 维保记录 ----------

func QueryEquipmentMaintenanceRecord(req *proto.QueryEquipmentMaintenanceRecordRequest, resp *proto.QueryEquipmentMaintenanceRecordResponse, preload bool) {
	db := model.DB.DB().Model(&model.EquipmentMaintenanceRecord{})
	if preload {
		db = db.Preload(clause.Associations)
	}
	if req.EquipmentID != "" {
		db = db.Where("`equipment_id` = ?", req.EquipmentID)
	}
	if req.StartDate != "" {
		start, err := time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			resp.Code = proto.Code_BadRequest
			resp.Message = "startDate格式无效，应为yyyy-MM-dd"
			return
		}
		db = db.Where("`execution_date` >= ?", start)
	}
	if req.EndDate != "" {
		end, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			resp.Code = proto.Code_BadRequest
			resp.Message = "endDate格式无效，应为yyyy-MM-dd"
			return
		}
		db = db.Where("`execution_date` < ?", end.AddDate(0, 0, 1))
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "execution_date desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.EquipmentMaintenanceRecord
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.EquipmentMaintenanceRecordsToPB(list)
	}
	resp.Total = resp.Records
}

func GetEquipmentMaintenanceRecordByID(id string) (*model.EquipmentMaintenanceRecord, error) {
	m := &model.EquipmentMaintenanceRecord{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

// ExecuteEquipmentMaintenance 执行维保：生成维保记录并滚动计划的下一次执行日期
func ExecuteEquipmentMaintenance(req *proto.ExecuteEquipmentMaintenanceRequest) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		plan := &model.EquipmentMaintenancePlan{}
		if err := tx.Preload(clause.Associations).First(plan, "`id` = ?", req.EquipmentMaintenancePlanID).Error; err != nil {
			return errors.New("读取维保计划失败")
		}

		now := time.Now()
		record := &model.EquipmentMaintenanceRecord{
			EquipmentMaintenancePlanID: &plan.ID,
			EquipmentID:                plan.EquipmentID,
			MaintenanceType:            plan.MaintenanceType,
			ExecutorID:                 req.ExecutorID,
			Content:                    req.Content,
			DurationMinutes:            req.DurationMinutes,
			Result:                     req.Result,
			AbnormalDescription:        req.AbnormalDescription,
		}
		if record.Result == "" {
			record.Result = "正常"
		}
		if err := tx.Create(record).Error; err != nil {
			return err
		}

		updates := map[string]interface{}{
			"last_execution": now,
			"next_execution": now.AddDate(0, 0, int(plan.CycleDays)),
		}
		if err := tx.Model(&model.EquipmentMaintenancePlan{}).Where("`id` = ?", plan.ID).Updates(updates).Error; err != nil {
			return err
		}

		if err := tx.Create(&model.OperationTrace{
			OperateUserID:  req.ExecutorID,
			ControllerName: "设备管理",
			ActionName:     "执行维保",
			RequestContent: fmt.Sprintf("计划:%s,设备:%s,类型:%s,结果:%s,耗时:%d分钟", plan.ID, derefEquipmentID(plan.EquipmentID), plan.MaintenanceType, record.Result, req.DurationMinutes),
		}).Error; err != nil {
			return err
		}

		//异常时联动设备状态为维修中
		if req.Result == "异常" && plan.EquipmentID != nil {
			if err := tx.Model(&model.Equipment{}).Where("`id` = ?", *plan.EquipmentID).
				Update("current_state", EquipmentStateMaintaining).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func derefEquipmentID(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
