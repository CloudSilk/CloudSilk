package model

import (
	"database/sql"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 工站开机记录
type ProductionStationStartup struct {
	ModelID
	StartupTime         time.Time          `json:"startupTime" gorm:"autoCreateTime:nano;comment:开机时间"`
	LastHeartbeatTime   sql.NullTime       `json:"lastHeartbeatTime" gorm:"comment:最后心跳时间"`
	ShutdownTime        sql.NullTime       `json:"shutdownTime" gorm:"comment:停机时间"`
	Duration            int32              `json:"duration" gorm:"comment:开机时长"`
	ProductionStationID string             `json:"productionStationID" gorm:"size:36;comment:生产工站ID"`
	ProductionStation   *ProductionStation `gorm:"constraint:OnDelete:CASCADE"`
	Remark              string             `json:"remark" gorm:"size:1000;comment:备注"`
}

func PBToProductionStationStartups(in []*proto.ProductionStationStartupInfo) []*ProductionStationStartup {
	var result []*ProductionStationStartup
	for _, c := range in {
		result = append(result, PBToProductionStationStartup(c))
	}
	return result
}

func PBToProductionStationStartup(in *proto.ProductionStationStartupInfo) *ProductionStationStartup {
	if in == nil {
		return nil
	}
	return &ProductionStationStartup{
		ModelID: ModelID{ID: in.Id},
		// StartupTime:         utils.ParseSqlNullTime(in.StartupTime),
		LastHeartbeatTime:   utils.ParseSqlNullTime(in.LastHeartbeatTime),
		ShutdownTime:        utils.ParseSqlNullTime(in.ShutdownTime),
		Duration:            in.Duration,
		ProductionStationID: in.ProductionStationID,
		Remark:              in.Remark,
	}
}

func ProductionStationStartupsToPB(in []*ProductionStationStartup) []*proto.ProductionStationStartupInfo {
	var list []*proto.ProductionStationStartupInfo
	for _, f := range in {
		list = append(list, ProductionStationStartupToPB(f))
	}
	return list
}

func ProductionStationStartupToPB(in *ProductionStationStartup) *proto.ProductionStationStartupInfo {
	if in == nil {
		return nil
	}

	m := &proto.ProductionStationStartupInfo{
		Id:                  in.ID,
		StartupTime:         utils.FormatTime(in.StartupTime),
		LastHeartbeatTime:   utils.FormatSqlNullTime(in.LastHeartbeatTime),
		ShutdownTime:        utils.FormatSqlNullTime(in.ShutdownTime),
		Duration:            in.Duration,
		ProductionStationID: in.ProductionStationID,
		ProductionStation:   ProductionStationToPB(in.ProductionStation),
		Remark:              in.Remark,
	}
	return m
}
