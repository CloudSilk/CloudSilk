package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

type ProductIssueRecord struct {
	ModelID
	MaterialTraceNo   string             `json:"MaterialTraceNo" gorm:"size:200;comment:物料追溯号"`
	CreateTime        time.Time          `json:"CreateTime" gorm:"autoCreateTime:nano;comment:发料时间"`
	CreateUserID      string             `json:"CreateUserID" gorm:"size:36;comment:发料人员ID"`
	CurrentState      string             `json:"CurrentState" gorm:"size:-1;comment:当前状态"`
	LastUpdateTime    time.Time          `json:"LastUpdateTime" gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	ProductInfoID     string             `json:"ProductInfoID" gorm:"size:36;comment:绑定产品ID"`
	ProductInfo       *ProductInfo       `json:"ProductInfo"`
	ProductOrderBomID string             `json:"ProductOrderBomID" gorm:"size:36;comment:工单物料ID"`
	ProductOrderBom   *ProductOrderBom   `json:"ProductOrderBom"`
	IssuanceProcessID string             `json:"IssuanceProcessID" gorm:"size:36;comment:发料工序ID"`
	IssuanceProcess   *ProductionProcess `json:"IssuanceProcess"`
}

func PBToProductIssueRecords(in []*proto.ProductIssueRecordInfo) []*ProductIssueRecord {
	var result []*ProductIssueRecord
	for _, c := range in {
		result = append(result, PBToProductIssueRecord(c))
	}
	return result
}

func PBToProductIssueRecord(in *proto.ProductIssueRecordInfo) *ProductIssueRecord {
	if in == nil {
		return nil
	}
	return &ProductIssueRecord{
		ModelID:         ModelID{ID: in.Id},
		MaterialTraceNo: in.MaterialTraceNo,
		// CreateTime:        utils.ParseSqlNullTime(in.CreateTime),
		CreateUserID: in.CreateUserID,
		CurrentState: in.CurrentState,
		// LastUpdateTime:    utils.ParseSqlNullTime(in.LastUpdateTime),
		ProductInfoID:     in.ProductInfoID,
		ProductOrderBomID: in.ProductOrderBomID,
		IssuanceProcessID: in.IssuanceProcessID,
	}
}

func ProductIssueRecordsToPB(in []*ProductIssueRecord) []*proto.ProductIssueRecordInfo {
	var list []*proto.ProductIssueRecordInfo
	for _, f := range in {
		list = append(list, ProductIssueRecordToPB(f))
	}
	return list
}

func ProductIssueRecordToPB(in *ProductIssueRecord) *proto.ProductIssueRecordInfo {
	if in == nil {
		return nil
	}
	productSerialNo := ""
	productOrderNo := ""
	materialNo := ""
	materialDescription := ""
	productionProcess := ""
	if in.ProductInfo != nil {
		productSerialNo = in.ProductInfo.ProductSerialNo
		if in.ProductInfo.ProductOrder != nil {
			productOrderNo = in.ProductInfo.ProductOrder.ProductOrderNo
		}
	}
	if in.ProductOrderBom != nil {
		materialNo = in.ProductOrderBom.MaterialNo
		materialDescription = in.ProductOrderBom.MaterialDescription
	}
	if in.IssuanceProcess != nil {
		productionProcess = in.IssuanceProcess.Description
	}
	m := &proto.ProductIssueRecordInfo{
		Id:                  in.ID,
		MaterialTraceNo:     in.MaterialTraceNo,
		CreateTime:          utils.FormatTime(in.CreateTime),
		CreateUserID:        in.CreateUserID,
		CurrentState:        in.CurrentState,
		LastUpdateTime:      utils.FormatTime(in.LastUpdateTime),
		ProductInfoID:       in.ProductInfoID,
		ProductOrderBomID:   in.ProductOrderBomID,
		IssuanceProcessID:   in.IssuanceProcessID,
		ProductSerialNo:     productSerialNo,
		ProductOrderNo:      productOrderNo,
		MaterialNo:          materialNo,
		MaterialDescription: materialDescription,
		ProductionProcess:   productionProcess,
	}
	return m
}
