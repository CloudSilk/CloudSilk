package model

import (
	"database/sql"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 工站登录记录
type ProductionStationSignup struct {
	ModelID
	LoginTime           time.Time          `json:"loginTime" gorm:"autoCreateTime:nano;comment:登入时间"`
	LastHeartbeatTime   sql.NullTime       `json:"lastHeartbeatTime" gorm:"comment:最后心跳时间"`
	LogoutTime          sql.NullTime       `json:"logoutTime" gorm:"comment:注销时间"`
	Duration            int32              `json:"duration" gorm:"comment:上机时长"`
	ProductionStationID string             `json:"productionStationID" gorm:"index;size:36;comment:生产工站ID"`
	ProductionStation   *ProductionStation `json:"productionStation" gorm:"constraint:OnDelete:CASCADE"` //生产工站
	LoginUserID         string             `json:"loginUserID" gorm:"index;size:36;comment:登录人员ID"`
	Remark              string             `json:"remark" gorm:"column:Remark;size:1000;comment:备注信息"`
}

func PBToProductionStationSignups(in []*proto.ProductionStationSignupInfo) []*ProductionStationSignup {
	var result []*ProductionStationSignup
	for _, c := range in {
		result = append(result, PBToProductionStationSignup(c))
	}
	return result
}

func PBToProductionStationSignup(in *proto.ProductionStationSignupInfo) *ProductionStationSignup {
	if in == nil {
		return nil
	}
	return &ProductionStationSignup{
		ModelID: ModelID{ID: in.Id},
		// LoginTime:           utils.ParseSqlNullTime(in.LoginTime),
		LastHeartbeatTime:   utils.ParseSqlNullTime(in.LastHeartbeatTime),
		LogoutTime:          utils.ParseSqlNullTime(in.LogoutTime),
		Duration:            in.Duration,
		ProductionStationID: in.ProductionStationID,
		LoginUserID:         in.LoginUserID,
		Remark:              in.Remark,
	}
}

func ProductionStationSignupsToPB(in []*ProductionStationSignup) []*proto.ProductionStationSignupInfo {
	var list []*proto.ProductionStationSignupInfo
	for _, f := range in {
		list = append(list, ProductionStationSignupToPB(f))
	}
	return list
}

func ProductionStationSignupToPB(in *ProductionStationSignup) *proto.ProductionStationSignupInfo {
	if in == nil {
		return nil
	}

	m := &proto.ProductionStationSignupInfo{
		Id:                  in.ID,
		LoginTime:           utils.FormatTime(in.LoginTime),
		LastHeartbeatTime:   utils.FormatSqlNullTime(in.LastHeartbeatTime),
		LogoutTime:          utils.FormatSqlNullTime(in.LogoutTime),
		Duration:            in.Duration,
		ProductionStationID: in.ProductionStationID,
		ProductionStation:   ProductionStationToPB(in.ProductionStation),
		LoginUserID:         in.LoginUserID,
		Remark:              in.Remark,
	}
	return m
}
