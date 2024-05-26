package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 产品型号BOM
type ProductModelBom struct {
	ModelID
	ItemNo              string        `json:"itemNo" gorm:"size:100;comment:项目号"`
	MaterialNo          string        `json:"materialNo" gorm:"size:100;comment:物料号"`
	MaterialDescription string        `json:"materialDescription" gorm:"size:1000;comment:物料描述"`
	RequireQTY          float32       `json:"requireQTY" gorm:"comment:需求数量"`
	Unit                string        `json:"unit" gorm:"size:20;comment:单位"`
	ProductionProcess   string        `json:"productionProcess" gorm:"size:20;comment:需求工序"`
	Remark              string        `json:"remark" gorm:"size:1000;comment:备注"`
	ProductModelID      string        `json:"productModelID" gorm:"index;size:36;comment:产品型号ID"`
	ProductModel        *ProductModel `json:"productModel" gorm:"constraint:OnDelete:CASCADE"` //产品型号
}

func PBToProductModelBoms(in []*proto.ProductModelBomInfo) []*ProductModelBom {
	var result []*ProductModelBom
	for _, c := range in {
		result = append(result, PBToProductModelBom(c))
	}
	return result
}

func PBToProductModelBom(in *proto.ProductModelBomInfo) *ProductModelBom {
	if in == nil {
		return nil
	}
	return &ProductModelBom{
		ModelID:             ModelID{ID: in.Id},
		ItemNo:              in.ItemNo,
		MaterialNo:          in.MaterialNo,
		MaterialDescription: in.MaterialDescription,
		RequireQTY:          in.RequireQTY,
		Unit:                in.Unit,
		ProductionProcess:   in.ProductionProcess,
		Remark:              in.Remark,
		ProductModelID:      in.ProductModelID,
	}
}

func ProductModelBomsToPB(in []*ProductModelBom) []*proto.ProductModelBomInfo {
	var list []*proto.ProductModelBomInfo
	for _, f := range in {
		list = append(list, ProductModelBomToPB(f))
	}
	return list
}

func ProductModelBomToPB(in *ProductModelBom) *proto.ProductModelBomInfo {
	if in == nil {
		return nil
	}
	var productCategoryID string
	if in.ProductModel != nil {
		productCategoryID = in.ProductModel.ProductCategoryID
	}
	m := &proto.ProductModelBomInfo{
		Id:                  in.ID,
		ItemNo:              in.ItemNo,
		MaterialNo:          in.MaterialNo,
		MaterialDescription: in.MaterialDescription,
		RequireQTY:          in.RequireQTY,
		Unit:                in.Unit,
		ProductionProcess:   in.ProductionProcess,
		Remark:              in.Remark,
		ProductModelID:      in.ProductModelID,
		ProductCategoryID:   productCategoryID,
	}
	return m
}
