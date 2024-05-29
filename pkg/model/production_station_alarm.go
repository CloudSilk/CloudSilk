package model

import (
	"database/sql"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 工站报警记录
type ProductionStationAlarm struct {
	ModelID
	CreateTime          time.Time          `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID        string             `json:"createUserID" gorm:"size:36;column:创建人员ID"`
	ProductionStationID string             `json:"productionStationID" gorm:"size:36;comment:生产工站ID"`
	ProductionStation   *ProductionStation `gorm:"constraint:OnDelete:CASCADE"`
	ProductionProcessID string             `json:"productionProcessID" gorm:"size:36;comment:生产工序ID"`
	ProductionProcess   *ProductionProcess `gorm:"constraint:OnDelete:CASCADE"`
	ProductInfoID       string             `json:"productInfoID" gorm:"size:36;comment:产品信息ID"`
	ProductInfo         *ProductInfo       `gorm:"constraint:OnDelete:CASCADE"`
	AlarmNo             string             `json:"alarmNo" gorm:"size:100;comment:报警编号"`
	AlarmMessage        string             `json:"alarmMessage" gorm:"size:1000;comment:报警信息"`
	CurrentState        string             `json:"currentState" gorm:"size:100;comment:当前状态"`
	HandleMethod        string             `json:"handleMethod" gorm:"size:100;comment:处理方式"`
	HandleTime          sql.NullTime       `json:"handleTime" gorm:"comment:处理时间"`
	Remark              string             `json:"remark" gorm:"size:1000;comment:备注信息"`
}

func PBToProductionStationAlarms(in []*proto.ProductionStationAlarmInfo) []*ProductionStationAlarm {
	var result []*ProductionStationAlarm
	for _, c := range in {
		result = append(result, PBToProductionStationAlarm(c))
	}
	return result
}

func PBToProductionStationAlarm(in *proto.ProductionStationAlarmInfo) *ProductionStationAlarm {
	if in == nil {
		return nil
	}
	return &ProductionStationAlarm{
		ModelID: ModelID{ID: in.Id},
		// CreateTime:          utils.ParseSqlNullTime(in.CreateTime),
		CreateUserID:        in.CreateUserID,
		ProductionStationID: in.ProductionStationID,
		ProductionProcessID: in.ProductionProcessID,
		ProductInfoID:       in.ProductInfoID,
		AlarmNo:             in.AlarmNo,
		AlarmMessage:        in.AlarmMessage,
		CurrentState:        in.CurrentState,
		HandleMethod:        in.HandleMethod,
		HandleTime:          utils.ParseSqlNullTime(in.HandleTime),
		Remark:              in.Remark,
	}
}

func ProductionStationAlarmsToPB(in []*ProductionStationAlarm) []*proto.ProductionStationAlarmInfo {
	var list []*proto.ProductionStationAlarmInfo
	for _, f := range in {
		list = append(list, ProductionStationAlarmToPB(f))
	}
	return list
}

func ProductionStationAlarmToPB(in *ProductionStationAlarm) *proto.ProductionStationAlarmInfo {
	if in == nil {
		return nil
	}

	m := &proto.ProductionStationAlarmInfo{
		Id:                  in.ID,
		CreateTime:          utils.FormatTime(in.CreateTime),
		CreateUserID:        in.CreateUserID,
		ProductionStationID: in.ProductionStationID,
		ProductionStation:   ProductionStationToPB(in.ProductionStation),
		ProductionProcessID: in.ProductionProcessID,
		ProductionProcess:   ProductionProcessToPB(in.ProductionProcess),
		ProductInfoID:       in.ProductInfoID,
		ProductInfo:         ProductInfoToPB(in.ProductInfo),
		AlarmNo:             in.AlarmNo,
		AlarmMessage:        in.AlarmMessage,
		CurrentState:        in.CurrentState,
		HandleMethod:        in.HandleMethod,
		HandleTime:          utils.FormatSqlNullTime(in.HandleTime),
		Remark:              in.Remark,
	}
	return m
}
