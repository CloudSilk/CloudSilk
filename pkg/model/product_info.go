package model

import (
	"database/sql"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 产品
type ProductInfo struct {
	ModelID
	ProductSerialNo     string        `json:"productSerialNo" gorm:"size:100;comment:产品序列号"`
	SecurityCode        string        `json:"securityCode" gorm:"size:100;comment:防伪码"`
	SecurityUrl         string        `json:"securityUrl" gorm:"size:1000;comment:防伪链接"`
	CreateTime          time.Time     `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID        string        `json:"createUserID" gorm:"size:36;comment:创建人员ID"`
	ReleaseTime         sql.NullTime  `json:"releaseTime" gorm:"comment:发放时间"`
	IssuedTime          sql.NullTime  `json:"issuedTime" gorm:"comment:发料时间"`
	StartedTime         sql.NullTime  `json:"startedTime" gorm:"comment:上线时间"`
	RemainingRoutes     int32         `json:"remainingRoutes" gorm:"comment:剩余工序"`
	EstimateTime        sql.NullTime  `json:"estimateTime" gorm:"comment:预计下线时间"`
	FinishedTime        sql.NullTime  `json:"finishedTime" gorm:"comment:下线时间"`
	DepositedTime       sql.NullTime  `json:"depositedTime" gorm:"comment:入库时间"`
	CurrentState        string        `json:"currentState" gorm:"size:100;comment:当前状态"`
	LastUpdateTime      time.Time     `json:"lastUpdateTime" gorm:"autoUpdateTime:nano;comment:更新时间"`
	Remark              string        `json:"remark" gorm:"size:1000;comment:备注"`
	ProductOrderID      string        `json:"productOrderID" gorm:"size:36;comment:产品工单ID"`
	ProductOrder        *ProductOrder `json:"productOrder" gorm:"constraint:OnDelete:CASCADE"` //产品工单
	ProductionProcessID *string       `json:"productionProcessID" gorm:"size:36;comment:当前工序ID"`
}

func PBToProductInfos(in []*proto.ProductInfoInfo) []*ProductInfo {
	var result []*ProductInfo
	for _, c := range in {
		result = append(result, PBToProductInfo(c))
	}
	return result
}

func PBToProductInfo(in *proto.ProductInfoInfo) *ProductInfo {
	if in == nil {
		return nil
	}

	var productionProcessID *string
	if in.ProductionProcessID != "" {
		productionProcessID = &in.ProductionProcessID
	}

	return &ProductInfo{
		ModelID:         ModelID{ID: in.Id},
		ProductSerialNo: in.ProductSerialNo,
		SecurityCode:    in.SecurityCode,
		SecurityUrl:     in.SecurityUrl,
		// CreateTime:          utils.ParseTime(in.CreateTime),
		CreateUserID:    in.CreateUserID,
		ReleaseTime:     utils.ParseSqlNullTime(in.ReleaseTime),
		IssuedTime:      utils.ParseSqlNullTime(in.IssuedTime),
		StartedTime:     utils.ParseSqlNullTime(in.StartedTime),
		RemainingRoutes: in.RemainingRoutes,
		EstimateTime:    utils.ParseSqlNullTime(in.EstimateTime),
		FinishedTime:    utils.ParseSqlNullTime(in.FinishedTime),
		DepositedTime:   utils.ParseSqlNullTime(in.DepositedTime),
		CurrentState:    in.CurrentState,
		// LastUpdateTime:      utils.ParseTime(in.LastUpdateTime),
		Remark:              in.Remark,
		ProductOrderID:      in.ProductOrderID,
		ProductionProcessID: productionProcessID,
	}
}

func ProductInfosToPB(in []*ProductInfo) []*proto.ProductInfoInfo {
	var list []*proto.ProductInfoInfo
	for _, f := range in {
		list = append(list, ProductInfoToPB(f))
	}
	return list
}

func ProductInfoToPB(in *ProductInfo) *proto.ProductInfoInfo {
	if in == nil {
		return nil
	}

	var productionProcessID string
	if in.ProductionProcessID != nil {
		productionProcessID = *in.ProductionProcessID
	}

	m := &proto.ProductInfoInfo{
		Id:                  in.ID,
		ProductSerialNo:     in.ProductSerialNo,
		SecurityCode:        in.SecurityCode,
		SecurityUrl:         in.SecurityUrl,
		CreateTime:          utils.FormatTime(in.CreateTime),
		CreateUserID:        in.CreateUserID,
		ReleaseTime:         utils.FormatSqlNullTime(in.ReleaseTime),
		IssuedTime:          utils.FormatSqlNullTime(in.IssuedTime),
		StartedTime:         utils.FormatSqlNullTime(in.StartedTime),
		RemainingRoutes:     in.RemainingRoutes,
		EstimateTime:        utils.FormatSqlNullTime(in.EstimateTime),
		FinishedTime:        utils.FormatSqlNullTime(in.FinishedTime),
		DepositedTime:       utils.FormatSqlNullTime(in.DepositedTime),
		CurrentState:        in.CurrentState,
		LastUpdateTime:      utils.FormatTime(in.LastUpdateTime),
		Remark:              in.Remark,
		ProductOrderID:      in.ProductOrderID,
		ProductOrder:        ProductOrderToPB(in.ProductOrder),
		ProductionProcessID: productionProcessID,
	}
	return m
}
