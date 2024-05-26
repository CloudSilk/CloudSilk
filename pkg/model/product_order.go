package model

import (
	"database/sql"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 产品工单
type ProductOrder struct {
	ModelID
	ProductOrderNo              string                    `json:"productOrderNo" gorm:"size:100;comment:生产工单号"`
	SalesOrderNo                string                    `json:"salesOrderNo" gorm:"size:100;comment:销售单号"`
	ItemNo                      string                    `json:"itemNo" gorm:"size:100;comment:销售项号"`
	CustomerName                string                    `json:"customerName" gorm:"size:400;comment:客户名称"`
	OrderType                   int32                     `json:"orderType" gorm:"comment:工单类型"`
	OrderTime                   sql.NullTime              `json:"orderTime" gorm:"comment:下单时间"`
	OrderQTY                    int32                     `json:"orderQTY" gorm:"comment:下单数量"`
	Warehouse                   string                    `json:"warehouse" gorm:"size:100;comment:成品仓库"`
	ReceiptNoteNo               string                    `json:"receiptNoteNo" gorm:"size:100;comment:成品入库单"`
	ProductionTeam              string                    `json:"productionTeam" gorm:"size:100;comment:生产班组"`
	PriorityLevel               int32                     `json:"priorityLevel" gorm:"comment:生产优先级"`
	DeliveryDate                sql.NullTime              `json:"deliveryDate" gorm:"comment:交货期限"`
	CreateTime                  time.Time                 `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID                string                    `json:"createUserID" gorm:"size:36;comment:创建人员ID"`
	ReleaseTime                 sql.NullTime              `json:"releaseTime" gorm:"comment:发放时间"`
	SelectedOptions             string                    `json:"selectedOptions" gorm:"size:2000;comment:选配信息"`
	PropertyBrief               string                    `json:"propertyBrief" gorm:"size:4000;comment:特性简述"`
	StandardWorkTime            int32                     `json:"standardWorkTime" gorm:"comment:标准节拍"`
	IssuedQTY                   int32                     `json:"issuedQTY" gorm:"comment:已发料数量"`
	LastIssueTime               sql.NullTime              `json:"lastIssueTime" gorm:"comment:最新发料时间"`
	EstimateStartTime           sql.NullTime              `json:"estimateStartTime" gorm:"comment:预计开工时间"`
	ActualStartTime             sql.NullTime              `json:"actualStartTime" gorm:"comment:实际开工时间"`
	StartedQTY                  int32                     `json:"startedQTY" gorm:"comment:已开工数量"`
	ObsoletedQTY                int32                     `json:"obsoletedQTY" gorm:"comment:已淘汰数量"`
	FinishedQTY                 int32                     `json:"finishedQTY" gorm:"comment:已完工数量"`
	EstimateConsumeTimePerPiece int32                     `json:"estimateConsumeTimePerPiece" gorm:"comment:预计单件耗时"`
	EstimateFinishTime          sql.NullTime              `json:"estimateFinishTime" gorm:"comment:预计完工时间"`
	ActualFinishTime            sql.NullTime              `json:"actualFinishTime" gorm:"comment:实际完工时间"`
	AverageConsumeTimePerPiece  int32                     `json:"averageConsumeTimePerPiece" gorm:"comment:平均单件耗时"`
	DepositedQTY                int32                     `json:"depositedQTY" gorm:"comment:已入库数量"`
	LastDepositTime             sql.NullTime              `json:"lastDepositTime" gorm:"comment:最新入库时间"`
	CurrentState                string                    `json:"currentState" gorm:"size:100;comment:当前状态"`
	LastUpdateTime              time.Time                 `json:"lastUpdateTime" gorm:"autoUpdateTime:nano;comment:最新更新时间"`
	Remark                      string                    `json:"remark" gorm:"size:1000;comment:备注"`
	ProductModelID              *string                   `json:"productModelID" gorm:"size:36;comment:产品型号ID"`
	ProductModel                *ProductModel             `json:"productModel" gorm:"constraint:OnDelete:SET NULL"` //产品型号
	ProductionLineID            *string                   `json:"productionLineID" gorm:"size:36;comment:发放产线ID"`
	ProductionLine              *ProductionLine           `json:"productionLine" gorm:"constraint:OnDelete:SET NULL"` //发放产线
	ProductOrderAttachments     []*ProductOrderAttachment `gorm:"constraint:OnDelete:CASCADE;"`                       //产品工单附件
	ProductOrderBoms            []*ProductOrderBom        `gorm:"constraint:OnDelete:CASCADE;"`                       //产品工单BOM
	ProductOrderAttributes      []*ProductOrderAttribute  `gorm:"constraint:OnDelete:CASCADE;"`                       //产品工单特性
	ProductInfos                []*ProductInfo            `gorm:"constraint:OnDelete:CASCADE;"`                       //产品清单
	ProductOrderProcesses       []*ProductOrderProcess    `gorm:"constraint:OnDelete:CASCADE;"`                       //产品工单工序
	ProductOrderLabels          []*ProductOrderLabel      `gorm:"constraint:OnDelete:CASCADE;"`                       //产品工单标签
	ProductOrderPackages        []*ProductOrderPackage    `gorm:"constraint:OnDelete:CASCADE;"`                       //产品工单包装
	ProductOrderPallets         []*ProductOrderPallet     `gorm:"constraint:OnDelete:CASCADE;"`                       //产品工单栈板
}

func PBToProductOrders(in []*proto.ProductOrderInfo) []*ProductOrder {
	var result []*ProductOrder
	for _, c := range in {
		result = append(result, PBToProductOrder(c))
	}
	return result
}

func PBToProductOrder(in *proto.ProductOrderInfo) *ProductOrder {
	if in == nil {
		return nil
	}

	var productionLineID, productModelID *string
	if in.ProductionLineID != "" {
		productionLineID = &in.ProductionLineID
	}
	if in.ProductModelID != "" {
		productModelID = &in.ProductModelID
	}

	return &ProductOrder{
		ModelID:        ModelID{ID: in.Id},
		ProductOrderNo: in.ProductOrderNo,
		SalesOrderNo:   in.SalesOrderNo,
		ItemNo:         in.ItemNo,
		CustomerName:   in.CustomerName,
		OrderType:      in.OrderType,
		OrderTime:      utils.ParseSqlNullTime(in.OrderTime),
		OrderQTY:       in.OrderQTY,
		Warehouse:      in.Warehouse,
		ReceiptNoteNo:  in.ReceiptNoteNo,
		ProductionTeam: in.ProductionTeam,
		PriorityLevel:  in.PriorityLevel,
		DeliveryDate:   utils.ParseSqlNullDate(in.DeliveryDate),
		// CreateTime:                  utils.ParseTime(in.CreateTime),
		CreateUserID:                in.CreateUserID,
		ReleaseTime:                 utils.ParseSqlNullTime(in.ReleaseTime),
		SelectedOptions:             in.SelectedOptions,
		PropertyBrief:               in.PropertyBrief,
		StandardWorkTime:            in.StandardWorkTime,
		IssuedQTY:                   in.IssuedQTY,
		LastIssueTime:               utils.ParseSqlNullTime(in.LastIssueTime),
		EstimateStartTime:           utils.ParseSqlNullTime(in.EstimateStartTime),
		ActualStartTime:             utils.ParseSqlNullTime(in.ActualStartTime),
		StartedQTY:                  in.StartedQTY,
		ObsoletedQTY:                in.ObsoletedQTY,
		FinishedQTY:                 in.FinishedQTY,
		EstimateConsumeTimePerPiece: in.EstimateConsumeTimePerPiece,
		EstimateFinishTime:          utils.ParseSqlNullTime(in.EstimateFinishTime),
		ActualFinishTime:            utils.ParseSqlNullTime(in.ActualFinishTime),
		AverageConsumeTimePerPiece:  in.AverageConsumeTimePerPiece,
		DepositedQTY:                in.DepositedQTY,
		LastDepositTime:             utils.ParseSqlNullTime(in.LastDepositTime),
		CurrentState:                in.CurrentState,
		// LastUpdateTime:              utils.ParseTime(in.LastUpdateTime),
		Remark:                  in.Remark,
		ProductModelID:          productModelID,
		ProductionLineID:        productionLineID,
		ProductOrderBoms:        PBToProductOrderBoms(in.ProductOrderBoms),
		ProductOrderLabels:      PBToProductOrderLabels(in.ProductOrderLabels),
		ProductOrderAttachments: PBToProductOrderAttachments(in.ProductOrderAttachments),
	}
}

func ProductOrdersToPB(in []*ProductOrder) []*proto.ProductOrderInfo {
	var list []*proto.ProductOrderInfo
	for _, f := range in {
		list = append(list, ProductOrderToPB(f))
	}
	return list
}

func ProductOrderToPB(in *ProductOrder) *proto.ProductOrderInfo {
	if in == nil {
		return nil
	}

	var productionLineID, productModelID string
	if in.ProductionLineID != nil {
		productionLineID = *in.ProductionLineID
	}
	if in.ProductModelID != nil {
		productModelID = *in.ProductModelID
	}

	m := &proto.ProductOrderInfo{
		Id:                          in.ID,
		ProductOrderNo:              in.ProductOrderNo,
		SalesOrderNo:                in.SalesOrderNo,
		ItemNo:                      in.ItemNo,
		CustomerName:                in.CustomerName,
		OrderType:                   in.OrderType,
		OrderTime:                   utils.FormatSqlNullTime(in.OrderTime),
		OrderQTY:                    in.OrderQTY,
		Warehouse:                   in.Warehouse,
		ReceiptNoteNo:               in.ReceiptNoteNo,
		ProductionTeam:              in.ProductionTeam,
		PriorityLevel:               in.PriorityLevel,
		DeliveryDate:                utils.FormatSqlNullDate(in.DeliveryDate),
		CreateTime:                  utils.FormatTime(in.CreateTime),
		CreateUserID:                in.CreateUserID,
		ReleaseTime:                 utils.FormatSqlNullTime(in.ReleaseTime),
		SelectedOptions:             in.SelectedOptions,
		PropertyBrief:               in.PropertyBrief,
		StandardWorkTime:            in.StandardWorkTime,
		IssuedQTY:                   in.IssuedQTY,
		LastIssueTime:               utils.FormatSqlNullTime(in.LastIssueTime),
		EstimateStartTime:           utils.FormatSqlNullTime(in.EstimateStartTime),
		ActualStartTime:             utils.FormatSqlNullTime(in.ActualStartTime),
		StartedQTY:                  in.StartedQTY,
		ObsoletedQTY:                in.ObsoletedQTY,
		FinishedQTY:                 in.FinishedQTY,
		EstimateConsumeTimePerPiece: in.EstimateConsumeTimePerPiece,
		EstimateFinishTime:          utils.FormatSqlNullTime(in.EstimateFinishTime),
		ActualFinishTime:            utils.FormatSqlNullTime(in.ActualFinishTime),
		AverageConsumeTimePerPiece:  in.AverageConsumeTimePerPiece,
		DepositedQTY:                in.DepositedQTY,
		LastDepositTime:             utils.FormatSqlNullTime(in.LastDepositTime),
		CurrentState:                in.CurrentState,
		LastUpdateTime:              utils.FormatTime(in.LastUpdateTime),
		Remark:                      in.Remark,
		ProductModelID:              productModelID,
		ProductModel:                ProductModelToPB(in.ProductModel),
		ProductionLineID:            productionLineID,
		ProductOrderAttachments:     ProductOrderAttachmentsToPB(in.ProductOrderAttachments),
		ProductOrderBoms:            ProductOrderBomsToPB(in.ProductOrderBoms),
		ProductInfos:                ProductInfosToPB(in.ProductInfos),
		ProductOrderAttributes:      ProductOrderAttributesToPB(in.ProductOrderAttributes),
		ProductOrderProcesses:       ProductOrderProcesssToPB(in.ProductOrderProcesses),
		ProductOrderLabels:          ProductOrderLabelsToPB(in.ProductOrderLabels),
		ProductOrderPackages:        ProductOrderPackagesToPB(in.ProductOrderPackages),
		ProductOrderPallets:         ProductOrderPalletsToPB(in.ProductOrderPallets),
	}
	return m
}
