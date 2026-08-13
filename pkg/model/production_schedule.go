package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 排程计划（APS）
type ProductionSchedulePlan struct {
	ModelID
	PlanNo            string                      `json:"planNo" gorm:"index;size:50;comment:计划批次号"`
	ProductionLineID  string                      `json:"productionLineID" gorm:"size:36;comment:产线ID"`
	ProductionLine    *ProductionLine             `json:"productionLine" gorm:"constraint:OnDelete:CASCADE"`
	StartTime         time.Time                   `json:"startTime" gorm:"comment:计划开始时间"`
	EndTime           time.Time                   `json:"endTime" gorm:"comment:计划结束时间"`
	OrderCount        int32                       `json:"orderCount" gorm:"comment:纳入排程工单数"`
	CurrentState      string                      `json:"currentState" gorm:"size:50;comment:状态"`
	Remark            string                      `json:"remark" gorm:"size:500;comment:备注"`
	CreateTime        time.Time                   `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	Items             []*ProductionScheduleItem   `json:"items" gorm:"constraint:OnDelete:CASCADE"`
}

// 排程明细
type ProductionScheduleItem struct {
	ModelID
	ProductionSchedulePlanID string                 `json:"productionSchedulePlanID" gorm:"index;size:36;comment:排程计划ID"`
	ProductionSchedulePlan   *ProductionSchedulePlan `json:"productionSchedulePlan" gorm:"constraint:OnDelete:CASCADE"`
	ProductOrderID           string                 `json:"productOrderID" gorm:"size:36;comment:工单ID"`
	ProductOrder             *ProductOrder          `json:"productOrder" gorm:"constraint:OnDelete:CASCADE"`
	Sequence                 int32                  `json:"sequence" gorm:"comment:顺序号"`
	PlannedStartTime         time.Time              `json:"plannedStartTime" gorm:"comment:计划开工时间"`
	PlannedEndTime           time.Time              `json:"plannedEndTime" gorm:"comment:计划完工时间"`
	DurationSeconds          int64                  `json:"durationSeconds" gorm:"comment:工时(秒)"`
	PriorityLevel            int32                  `json:"priorityLevel" gorm:"comment:优先级"`
}

func PBToProductionSchedulePlans(in []*proto.ProductionSchedulePlanInfo) []*ProductionSchedulePlan {
	var result []*ProductionSchedulePlan
	for _, c := range in {
		result = append(result, PBToProductionSchedulePlan(c))
	}
	return result
}

func PBToProductionSchedulePlan(in *proto.ProductionSchedulePlanInfo) *ProductionSchedulePlan {
	if in == nil {
		return nil
	}
	return &ProductionSchedulePlan{
		ModelID:           ModelID{ID: in.Id},
		PlanNo:            in.PlanNo,
		ProductionLineID:  in.ProductionLineID,
		OrderCount:        in.OrderCount,
		CurrentState:      in.CurrentState,
		Remark:            in.Remark,
	}
}

func ProductionSchedulePlansToPB(in []*ProductionSchedulePlan) []*proto.ProductionSchedulePlanInfo {
	var list []*proto.ProductionSchedulePlanInfo
	for _, f := range in {
		list = append(list, ProductionSchedulePlanToPB(f))
	}
	return list
}

func ProductionSchedulePlanToPB(in *ProductionSchedulePlan) *proto.ProductionSchedulePlanInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductionSchedulePlanInfo{
		Id:               in.ID,
		PlanNo:           in.PlanNo,
		ProductionLineID: in.ProductionLineID,
		OrderCount:       in.OrderCount,
		CurrentState:     in.CurrentState,
		Remark:           in.Remark,
		StartTime:        in.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:          in.EndTime.Format("2006-01-02 15:04:05"),
	}
	if in.ProductionLine != nil {
		m.ProductionLineCode = in.ProductionLine.Code
	}
	for _, item := range in.Items {
		m.Items = append(m.Items, ProductionScheduleItemToPB(item))
	}
	return m
}

func ProductionScheduleItemToPB(in *ProductionScheduleItem) *proto.ProductionScheduleItemInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductionScheduleItemInfo{
		Id:                       in.ID,
		ProductionSchedulePlanID: in.ProductionSchedulePlanID,
		ProductOrderID:           in.ProductOrderID,
		Sequence:                 in.Sequence,
		PlannedStartTime:         in.PlannedStartTime.Format("2006-01-02 15:04:05"),
		PlannedEndTime:           in.PlannedEndTime.Format("2006-01-02 15:04:05"),
		DurationSeconds:          in.DurationSeconds,
		PriorityLevel:            in.PriorityLevel,
	}
	if in.ProductOrder != nil {
		m.ProductOrderNo = in.ProductOrder.ProductOrderNo
		if in.ProductOrder.DeliveryDate.Valid {
			m.DeliveryDate = in.ProductOrder.DeliveryDate.Time.Format("2006-01-02")
		}
	}
	return m
}
