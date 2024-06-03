package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 产品工单工序
type ProductOrderProcess struct {
	ModelID
	SortIndex                int32                      `json:"sortIndex" gorm:"comment:工序顺序"`
	CreateTime               time.Time                  `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID             string                     `json:"createUserID" gorm:"comment:创建人员ID"`
	Enable                   bool                       `json:"enable" gorm:"comment:是否启用"`
	LastUpdateTime           time.Time                  `json:"lastUpdateTime" gorm:"autoUpdateTime:nano;comment:更新时间"`
	Remark                   string                     `json:"remark" gorm:"size:1000;comment:备注"`
	ProductionProcessID      string                     `json:"productionProcessID" gorm:"size:36;comment:生产工序ID"`
	ProductionProcess        *ProductionProcess         `json:"productionProcess" gorm:"constraint:OnDelete:CASCADE"` //生产工序
	ProductOrderID           string                     `json:"productOrderID" gorm:"size:36;comment:隶属工单ID"`
	ProductOrder             *ProductOrder              `json:"productOrder" gorm:"constraint:OnDelete:CASCADE"` //隶属工单
	ProductOrderProcessSteps []*ProductOrderProcessStep `json:"productOrderProcessSteps" gorm:"constraint:OnDelete:CASCADE;"`
}

func PBToProductOrderProcesss(in []*proto.ProductOrderProcessInfo) []*ProductOrderProcess {
	var result []*ProductOrderProcess
	for _, c := range in {
		result = append(result, PBToProductOrderProcess(c))
	}
	return result
}

func PBToProductOrderProcess(in *proto.ProductOrderProcessInfo) *ProductOrderProcess {
	if in == nil {
		return nil
	}
	return &ProductOrderProcess{
		ModelID:   ModelID{ID: in.Id},
		SortIndex: in.SortIndex,
		// CreateTime:          utils.ParseTime(in.CreateTime),
		CreateUserID: in.CreateUserID,
		Enable:       in.Enable,
		// LastUpdateTime:      utils.ParseTime(in.LastUpdateTime),
		Remark:              in.Remark,
		ProductionProcessID: in.ProductionProcessID,
		ProductOrderID:      in.ProductOrderID,
	}
}

func ProductOrderProcesssToPB(in []*ProductOrderProcess) []*proto.ProductOrderProcessInfo {
	var list []*proto.ProductOrderProcessInfo
	for _, f := range in {
		list = append(list, ProductOrderProcessToPB(f))
	}
	return list
}

func ProductOrderProcessToPB(in *ProductOrderProcess) *proto.ProductOrderProcessInfo {
	if in == nil {
		return nil
	}
	productOrderNo := ""
	if in.ProductOrder != nil {
		productOrderNo = in.ProductOrder.ProductOrderNo
	}
	var productionProcess *proto.ProductionProcessInfo
	if in.ProductionProcess != nil {
		productionProcess = ProductionProcessToPB(in.ProductionProcess)
	}
	m := &proto.ProductOrderProcessInfo{
		Id:                  in.ID,
		SortIndex:           in.SortIndex,
		CreateTime:          utils.FormatTime(in.CreateTime),
		CreateUserID:        in.CreateUserID,
		Enable:              in.Enable,
		LastUpdateTime:      utils.FormatTime(in.LastUpdateTime),
		Remark:              in.Remark,
		ProductionProcessID: in.ProductionProcessID,
		ProductOrderID:      in.ProductOrderID,
		ProductOrderNo:      productOrderNo,
		ProductionProcess:   productionProcess,
	}
	return m
}
