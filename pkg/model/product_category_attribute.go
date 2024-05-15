package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 产品类别特性
type ProductCategoryAttribute struct {
	ModelID
	AllowNullOrBlank              bool                             `json:"allowNullOrBlank" gorm:"comment:允许空缺"`
	DefaultValue                  string                           `json:"defaultValue" gorm:"size:200;comment:预设值"`
	ProductCategoryID             string                           `json:"productCategoryID" gorm:"size:36;comment:产品类别ID"`
	ProductAttributeID            string                           `json:"productAttributeID" gorm:"size:36;comment:产品特性ID"`
	ProductAttribute              *ProductAttribute                `json:"productAttribute"`                                                  ///产品特性
	ProductCategoryAttributeValue []*ProductCategoryAttributeValue `json:"productCategoryAttributeValue" gorm:"constraint:OnDelete:CASCADE;"` //产品类别特征值
}

// 产品类别特征值
type ProductCategoryAttributeValue struct {
	ModelID
	ProductCategoryAttributeID string `json:"productCategoryAttributeID" gorm:"index;size:36;comment:产品类别特性ID"`
	Value                      string `json:"value" gorm:"size:200;comment:值"`
	Description                string `json:"description" gorm:"size:1000;comment:描述"`
}

func PBToProductCategoryAttributes(in []*proto.ProductCategoryAttributeInfo) []*ProductCategoryAttribute {
	var result []*ProductCategoryAttribute
	for _, c := range in {
		result = append(result, PBToProductCategoryAttribute(c))
	}
	return result
}

func PBToProductCategoryAttribute(in *proto.ProductCategoryAttributeInfo) *ProductCategoryAttribute {
	if in == nil {
		return nil
	}
	return &ProductCategoryAttribute{
		ModelID:                       ModelID{ID: in.Id},
		AllowNullOrBlank:              in.AllowNullOrBlank,
		DefaultValue:                  in.DefaultValue,
		ProductCategoryID:             in.ProductCategoryID,
		ProductAttributeID:            in.ProductAttributeID,
		ProductCategoryAttributeValue: PBToProductCategoryAttributeValues(in.ProductCategoryAttributeValue),
	}
}

func ProductCategoryAttributesToPB(in []*ProductCategoryAttribute) []*proto.ProductCategoryAttributeInfo {
	var list []*proto.ProductCategoryAttributeInfo
	for _, f := range in {
		list = append(list, ProductCategoryAttributeToPB(f))
	}
	return list
}

func ProductCategoryAttributeToPB(in *ProductCategoryAttribute) *proto.ProductCategoryAttributeInfo {
	if in == nil {
		return nil
	}
	productAttributeCode := ""
	productAttributeDescription := ""
	if in.ProductAttribute != nil {
		productAttributeCode = in.ProductAttribute.Code
		productAttributeDescription = in.ProductAttribute.Description
	}
	m := &proto.ProductCategoryAttributeInfo{
		Id:                            in.ID,
		AllowNullOrBlank:              in.AllowNullOrBlank,
		DefaultValue:                  in.DefaultValue,
		ProductCategoryID:             in.ProductCategoryID,
		ProductAttributeID:            in.ProductAttributeID,
		ProductAttributeCode:          productAttributeCode,
		ProductAttributeDescription:   productAttributeDescription,
		ProductCategoryAttributeValue: ProductCategoryAttributeValuesToPB(in.ProductCategoryAttributeValue),
	}
	return m
}

func PBToProductCategoryAttributeValues(in []*proto.ProductCategoryAttributeValueInfo) []*ProductCategoryAttributeValue {
	var result []*ProductCategoryAttributeValue
	for _, c := range in {
		result = append(result, PBToProductCategoryAttributeValue(c))
	}
	return result
}

func PBToProductCategoryAttributeValue(in *proto.ProductCategoryAttributeValueInfo) *ProductCategoryAttributeValue {
	if in == nil {
		return nil
	}
	return &ProductCategoryAttributeValue{
		ModelID:                    ModelID{ID: in.Id},
		Value:                      in.Value,
		Description:                in.Description,
		ProductCategoryAttributeID: in.ProductCategoryAttributeID,
	}
}

func ProductCategoryAttributeValuesToPB(in []*ProductCategoryAttributeValue) []*proto.ProductCategoryAttributeValueInfo {
	var list []*proto.ProductCategoryAttributeValueInfo
	for _, f := range in {
		list = append(list, ProductCategoryAttributeValueToPB(f))
	}
	return list
}

func ProductCategoryAttributeValueToPB(in *ProductCategoryAttributeValue) *proto.ProductCategoryAttributeValueInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductCategoryAttributeValueInfo{
		Id:                         in.ID,
		Value:                      in.Value,
		Description:                in.Description,
		ProductCategoryAttributeID: in.ProductCategoryAttributeID,
	}
	return m
}
