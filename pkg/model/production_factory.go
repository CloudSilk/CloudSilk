package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 生产工厂
type ProductionFactory struct {
	ModelID
	Code            string            `json:"code" gorm:"index;size:100;comment:代号"`
	Description     string            `json:"description" gorm:"size:200;comment:描述"`
	Identifier      string            `json:"identifier" gorm:"size:100;comment:识别码"`
	Remark          string            `json:"remark" gorm:"size:1000;comment:备注"`
	ProductionLines []*ProductionLine `json:"productionLines" gorm:"constraint:OnDelete:CASCADE;"` //生产厂线
}

func PBToProductionFactorys(in []*proto.ProductionFactoryInfo) []*ProductionFactory {
	var result []*ProductionFactory
	for _, c := range in {
		result = append(result, PBToProductionFactory(c))
	}
	return result
}

func PBToProductionFactory(in *proto.ProductionFactoryInfo) *ProductionFactory {
	if in == nil {
		return nil
	}
	return &ProductionFactory{
		ModelID:         ModelID{ID: in.Id},
		Code:            in.Code,
		Description:     in.Description,
		Identifier:      in.Identifier,
		Remark:          in.Remark,
		ProductionLines: PBToProductionLines(in.ProductionLines),
	}
}

func ProductionFactorysToPB(in []*ProductionFactory) []*proto.ProductionFactoryInfo {
	var list []*proto.ProductionFactoryInfo
	for _, f := range in {
		list = append(list, ProductionFactoryToPB(f))
	}
	return list
}

func ProductionFactoryToPB(in *ProductionFactory) *proto.ProductionFactoryInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductionFactoryInfo{
		Id:              in.ID,
		Code:            in.Code,
		Description:     in.Description,
		Identifier:      in.Identifier,
		Remark:          in.Remark,
		ProductionLines: ProductionLinesToPB(in.ProductionLines),
	}
	return m
}
