package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 产品类别
type ProductCategory struct {
	ModelID
	Code                      string                      `json:"code" gorm:"index;size:100;comment:代号"`
	Description               string                      `json:"description" gorm:"size:200;comment:描述"`
	Identifier                string                      `json:"identifier" gorm:"size:100;comment:识别码"`
	IsAuthorized              bool                        `json:"isAuthorized" gorm:"comment:是否授权"`
	AttributeExpression       string                      `json:"attributeExpression" gorm:"size:4000;comment:特性表达式"`
	Remark                    string                      `json:"remark" gorm:"size:1000;comment:备注"`
	ProductBrandID            string                      `json:"productBrandID" gorm:"size:36;comment:产品品牌"`
	ProductBrand              *ProductBrand               `json:"productBrand" gorm:"constraint:OnDelete:CASCADE"` //产品品牌
	ProductCategoryAttributes []*ProductCategoryAttribute `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToProductCategorys(in []*proto.ProductCategoryInfo) []*ProductCategory {
	var result []*ProductCategory
	for _, c := range in {
		result = append(result, PBToProductCategory(c))
	}
	return result
}

func PBToProductCategory(in *proto.ProductCategoryInfo) *ProductCategory {
	if in == nil {
		return nil
	}
	return &ProductCategory{
		ModelID:             ModelID{ID: in.Id},
		Code:                in.Code,
		Description:         in.Description,
		Identifier:          in.Identifier,
		IsAuthorized:        in.IsAuthorized,
		AttributeExpression: in.AttributeExpression,
		Remark:              in.Remark,
		ProductBrandID:      in.ProductBrandID,
	}
}

func ProductCategorysToPB(in []*ProductCategory) []*proto.ProductCategoryInfo {
	var list []*proto.ProductCategoryInfo
	for _, f := range in {
		list = append(list, ProductCategoryToPB(f))
	}
	return list
}

func ProductCategoryToPB(in *ProductCategory) *proto.ProductCategoryInfo {
	if in == nil {
		return nil
	}
	productBrandName := ""
	if in.ProductBrand != nil {
		productBrandName = in.ProductBrand.Description
	}
	m := &proto.ProductCategoryInfo{
		Id:                  in.ID,
		Code:                in.Code,
		Description:         in.Description,
		Identifier:          in.Identifier,
		IsAuthorized:        in.IsAuthorized,
		AttributeExpression: in.AttributeExpression,
		Remark:              in.Remark,
		ProductBrandID:      in.ProductBrandID,
		ProductBrandName:    productBrandName,
	}
	return m
}
