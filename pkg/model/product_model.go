package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 产品型号
type ProductModel struct {
	ModelID
	Code                        string                        `json:"code" gorm:"index;size:200;comment:型号"`
	MaterialNo                  string                        `json:"materialNo" gorm:"index;size:100;comment:物料号"`
	MaterialDescription         string                        `json:"materialDescription" gorm:"size:1000;comment:物料描述"`
	Identifier                  string                        `json:"identifier" gorm:"size:100;comment:识别码"`
	IsPrefabricated             bool                          `json:"isPrefabricated" gorm:"comment:是否预制"`
	IsAuthorized                bool                          `json:"isAuthorized" gorm:"comment:是否授权"`
	Remark                      string                        `json:"remark" gorm:"size:1000;comment:备注"`
	ProductCategoryID           string                        `json:"productCategoryID" gorm:"index;size:36;comment:产品类别ID"`
	ProductCategory             *ProductCategory              `json:"productCategory"`                                                 //产品类别
	ProductModelAttributeValues []*ProductModelAttributeValue `json:"productModelAttributeValues" gorm:"constraint:OnDelete:CASCADE;"` //产品型号特征值
	ProductModelBoms            []*ProductModelBom            `json:"productModelBoms" gorm:"constraint:OnDelete:CASCADE;"`            //产品型号Bom
}

// 产品型号特征值
type ProductModelAttributeValue struct {
	ModelID
	ProductModelID     string            `json:"productModelID" gorm:"index;size:36;comment:产品型号ID"`
	ProductAttributeID string            `json:"productAttributeID" gorm:"size:36;comment:产品特性ID"`
	ProductAttribute   *ProductAttribute `gorm:"constraint:OnDelete:CASCADE"`
	AssignedValue      string            `json:"assignedValue" gorm:"size:100;comment:设定值"`
	AllowNullOrBlank   bool              `json:"allowNullOrBlank" gorm:"-"` //允许空缺
}

func PBToProductModels(in []*proto.ProductModelInfo) []*ProductModel {
	var result []*ProductModel
	for _, c := range in {
		result = append(result, PBToProductModel(c))
	}
	return result
}

func PBToProductModel(in *proto.ProductModelInfo) *ProductModel {
	if in == nil {
		return nil
	}
	return &ProductModel{
		ModelID:                     ModelID{ID: in.Id},
		Code:                        in.Code,
		MaterialNo:                  in.MaterialNo,
		MaterialDescription:         in.MaterialDescription,
		Identifier:                  in.Identifier,
		IsPrefabricated:             in.IsPrefabricated,
		IsAuthorized:                in.IsAuthorized,
		Remark:                      in.Remark,
		ProductCategoryID:           in.ProductCategoryID,
		ProductModelAttributeValues: PBToProductModelAttributeValues(in.ProductModelAttributeValues),
		ProductModelBoms:            PBToProductModelBoms(in.ProductModelBoms),
	}
}

func ProductModelsToPB(in []*ProductModel) []*proto.ProductModelInfo {
	var list []*proto.ProductModelInfo
	for _, f := range in {
		list = append(list, ProductModelToPB(f))
	}
	return list
}

func ProductModelToPB(in *ProductModel) *proto.ProductModelInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductModelInfo{
		Id:                          in.ID,
		Code:                        in.Code,
		MaterialNo:                  in.MaterialNo,
		MaterialDescription:         in.MaterialDescription,
		Identifier:                  in.Identifier,
		IsPrefabricated:             in.IsPrefabricated,
		IsAuthorized:                in.IsAuthorized,
		Remark:                      in.Remark,
		ProductCategoryID:           in.ProductCategoryID,
		ProductCategory:             ProductCategoryToPB(in.ProductCategory),
		ProductModelAttributeValues: ProductModelAttributeValuesToPB(in.ProductModelAttributeValues),
		ProductModelBoms:            ProductModelBomsToPB(in.ProductModelBoms),
	}
	return m
}

func PBToProductModelAttributeValues(in []*proto.ProductModelAttributeValueInfo) []*ProductModelAttributeValue {
	var result []*ProductModelAttributeValue
	for _, c := range in {
		result = append(result, PBToProductModelAttributeValue(c))
	}
	return result
}

func PBToProductModelAttributeValue(in *proto.ProductModelAttributeValueInfo) *ProductModelAttributeValue {
	if in == nil {
		return nil
	}
	return &ProductModelAttributeValue{
		ModelID:            ModelID{ID: in.Id},
		ProductModelID:     in.ProductModelID,
		ProductAttributeID: in.ProductAttributeID,
		AssignedValue:      in.AssignedValue,
	}
}

func ProductModelAttributeValuesToPB(in []*ProductModelAttributeValue) []*proto.ProductModelAttributeValueInfo {
	var list []*proto.ProductModelAttributeValueInfo
	for _, f := range in {
		list = append(list, ProductModelAttributeValueToPB(f))
	}
	return list
}

func ProductModelAttributeValueToPB(in *ProductModelAttributeValue) *proto.ProductModelAttributeValueInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductModelAttributeValueInfo{
		Id:                 in.ID,
		ProductModelID:     in.ProductModelID,
		ProductAttributeID: in.ProductAttributeID,
		ProductAttribute:   ProductAttributeToPB(in.ProductAttribute),
		AssignedValue:      in.AssignedValue,
		AllowNullOrBlank:   in.AllowNullOrBlank,
	}
	return m
}
