package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 生产作业手册
type ProductionProcessSop struct {
	ModelID
	ProductionProcessID string             `json:"productionProcessID" gorm:"size:36;comment:生产工序ID"`
	ProductionProcess   *ProductionProcess `gorm:"constraint:OnDelete:CASCADE"` //生产工序
	ProductModelID      string             `json:"productModelID" gorm:"size:36;comment:产品型号ID"`
	ProductModel        *ProductModel      `gorm:"constraint:OnDelete:CASCADE"` //产品型号
	FileLink            string             `json:"fileLink" gorm:"size:-1;comment:文件链接"`
}

func PBToProductionProcessSops(in []*proto.ProductionProcessSopInfo) []*ProductionProcessSop {
	var result []*ProductionProcessSop
	for _, c := range in {
		result = append(result, PBToProductionProcessSop(c))
	}
	return result
}

func PBToProductionProcessSop(in *proto.ProductionProcessSopInfo) *ProductionProcessSop {
	if in == nil {
		return nil
	}
	return &ProductionProcessSop{
		ModelID:             ModelID{ID: in.Id},
		ProductionProcessID: in.ProductionProcessID,
		ProductModelID:      in.ProductModelID,
		FileLink:            in.FileLink,
	}
}

func ProductionProcessSopsToPB(in []*ProductionProcessSop) []*proto.ProductionProcessSopInfo {
	var list []*proto.ProductionProcessSopInfo
	for _, f := range in {
		list = append(list, ProductionProcessSopToPB(f))
	}
	return list
}

func ProductionProcessSopToPB(in *ProductionProcessSop) *proto.ProductionProcessSopInfo {
	if in == nil {
		return nil
	}
	code := ""
	materialNo := ""
	materialDescription := ""
	if in.ProductModel != nil {
		code = in.ProductModel.Code
		materialNo = in.ProductModel.MaterialNo
		materialDescription = in.ProductModel.MaterialDescription
	}
	var productionLineID string
	if in.ProductionProcess != nil {
		productionLineID = in.ProductionProcess.ProductionLineID
	}
	m := &proto.ProductionProcessSopInfo{
		Id:                  in.ID,
		ProductionProcessID: in.ProductionProcessID,
		ProductModelID:      in.ProductModelID,
		FileLink:            in.FileLink,
		Code:                code,
		MaterialNo:          materialNo,
		MaterialDescription: materialDescription,
		ProductionLineID:    productionLineID,
	}
	return m
}
