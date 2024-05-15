package model

import (
	"database/sql"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 工站故障记录
type ProductionStationBreakdown struct {
	ModelID
	ProductionStationID          string             `json:"productionStationID" gorm:"size:36;comment:生产工站ID"`
	ProductionStation            *ProductionStation ``
	CreateTime                   time.Time          `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID                 string             `json:"createUserID" gorm:"size:36;comment:创建人员ID"`
	EquipmentBreakdownTypeID     string             `json:"equipmentBreakdownTypeID" gorm:"size:36;comment:故障类型ID"`
	EquipmentBreakdownCauseID    string             `json:"equipmentBreakdownCauseID" gorm:"size:36;comment:故障原因ID"`
	EquipmentBreakdownSolutionID string             `json:"equipmentBreakdownSolutionID" gorm:"size:36;comment:故障方案ID"`
	MaintainBrief                string             `json:"maintainBrief" gorm:"size:2000;comment:维修简述"`
	CompleteTime                 sql.NullTime       `json:"completeTime" gorm:"comment:完成时间"`
	Duration                     int32              `json:"duration" gorm:"comment:停机时长"`
	MaintainUserID               string             `json:"maintainUserID" gorm:"size:36;comment:维修人员ID"`
	Remark                       string             `json:"remark" gorm:"size:1000;comment:备注信息"`
	EquipmentID                  string             `json:"equipmentID" gorm:"size:36;comment:设备ID"`
}

func PBToProductionStationBreakdowns(in []*proto.ProductionStationBreakdownInfo) []*ProductionStationBreakdown {
	var result []*ProductionStationBreakdown
	for _, c := range in {
		result = append(result, PBToProductionStationBreakdown(c))
	}
	return result
}

func PBToProductionStationBreakdown(in *proto.ProductionStationBreakdownInfo) *ProductionStationBreakdown {
	if in == nil {
		return nil
	}
	return &ProductionStationBreakdown{
		ModelID:             ModelID{ID: in.Id},
		ProductionStationID: in.ProductionStationID,
		// CreateTime:                   utils.ParseSqlNullTime(in.CreateTime),
		CreateUserID:                 in.CreateUserID,
		EquipmentBreakdownTypeID:     in.EquipmentBreakdownTypeID,
		EquipmentBreakdownCauseID:    in.EquipmentBreakdownCauseID,
		EquipmentBreakdownSolutionID: in.EquipmentBreakdownSolutionID,
		MaintainBrief:                in.MaintainBrief,
		CompleteTime:                 utils.ParseSqlNullTime(in.CompleteTime),
		Duration:                     in.Duration,
		MaintainUserID:               in.MaintainUserID,
		Remark:                       in.Remark,
		EquipmentID:                  in.EquipmentID,
	}
}

func ProductionStationBreakdownsToPB(in []*ProductionStationBreakdown) []*proto.ProductionStationBreakdownInfo {
	var list []*proto.ProductionStationBreakdownInfo
	for _, f := range in {
		list = append(list, ProductionStationBreakdownToPB(f))
	}
	return list
}

func ProductionStationBreakdownToPB(in *ProductionStationBreakdown) *proto.ProductionStationBreakdownInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductionStationBreakdownInfo{
		Id:                           in.ID,
		ProductionStationID:          in.ProductionStationID,
		ProductionStation:            ProductionStationToPB(in.ProductionStation),
		CreateTime:                   utils.FormatTime(in.CreateTime),
		CreateUserID:                 in.CreateUserID,
		EquipmentBreakdownTypeID:     in.EquipmentBreakdownTypeID,
		EquipmentBreakdownCauseID:    in.EquipmentBreakdownCauseID,
		EquipmentBreakdownSolutionID: in.EquipmentBreakdownSolutionID,
		MaintainBrief:                in.MaintainBrief,
		CompleteTime:                 utils.FormatSqlNullTime(in.CompleteTime),
		Duration:                     in.Duration,
		MaintainUserID:               in.MaintainUserID,
		Remark:                       in.Remark,
		EquipmentID:                  in.EquipmentID,
	}
	return m
}
