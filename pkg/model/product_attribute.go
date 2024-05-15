package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 产品特性管理
type ProductAttribute struct {
	ModelID
	Code                        string                        `json:"code" gorm:"index;size:50;comment:代号"`
	Description                 string                        `json:"description" gorm:"size:100;comment:描述"`
	DefaultValue                string                        `json:"defaultValue" gorm:"size:100;comment:预设值"`
	Remark                      string                        `json:"remark" gorm:"size:500;comment:备注"`
	ProductAttributeIdentifiers []*ProductAttributeIdentifier `json:"productAttributeIdentifiers" gorm:"constraint:OnDelete:CASCADE;"` //产品特征标识符
}

// 产品特征标识符
type ProductAttributeIdentifier struct {
	ModelID
	ProductAttributeID                           string                                         `json:"productAttributeID" gorm:"index;size:36;comment:产品特性ID"`
	Identifier                                   string                                         `json:"identifier" gorm:"size:100;comment:识别码"`
	Description                                  string                                         `json:"description" gorm:"size:500;comment:描述"`
	ProductAttributeIdentifierAvailableCategorys []*ProductAttributeIdentifierAvailableCategory `gorm:"constraint:OnDelete:CASCADE"` //可适配的产品类别
}

// 可适配的产品类别
type ProductAttributeIdentifierAvailableCategory struct {
	ModelID
	ProductAttributeIdentifierID string           `gorm:"index;size:36;comment:产品特征标识符ID"`
	ProductCategoryID            string           `gorm:"size:36;comment:产品类别ID"`
	ProductCategory              *ProductCategory `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToProductAttributes(in []*proto.ProductAttributeInfo) []*ProductAttribute {
	var result []*ProductAttribute
	for _, c := range in {
		result = append(result, PBToProductAttribute(c))
	}
	return result
}

func PBToProductAttribute(in *proto.ProductAttributeInfo) *ProductAttribute {
	if in == nil {
		return nil
	}
	return &ProductAttribute{
		ModelID:                     ModelID{ID: in.Id},
		Code:                        in.Code,
		Description:                 in.Description,
		DefaultValue:                in.DefaultValue,
		Remark:                      in.Remark,
		ProductAttributeIdentifiers: PBToProductAttributeIdentifiers(in.ProductAttributeIdentifiers),
	}
}

func ProductAttributesToPB(in []*ProductAttribute) []*proto.ProductAttributeInfo {
	var list []*proto.ProductAttributeInfo
	for _, f := range in {
		list = append(list, ProductAttributeToPB(f))
	}
	return list
}

func ProductAttributeToPB(in *ProductAttribute) *proto.ProductAttributeInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductAttributeInfo{
		Id:                          in.ID,
		Code:                        in.Code,
		Description:                 in.Description,
		DefaultValue:                in.DefaultValue,
		Remark:                      in.Remark,
		ProductAttributeIdentifiers: ProductAttributeIdentifiersToPB(in.ProductAttributeIdentifiers),
	}
	return m
}

func PBToProductAttributeIdentifiers(in []*proto.ProductAttributeIdentifierInfo) []*ProductAttributeIdentifier {
	var result []*ProductAttributeIdentifier
	for _, c := range in {
		result = append(result, PBToProductAttributeIdentifier(c))
	}
	return result
}

func PBToProductAttributeIdentifier(in *proto.ProductAttributeIdentifierInfo) *ProductAttributeIdentifier {
	if in == nil {
		return nil
	}
	return &ProductAttributeIdentifier{
		ModelID:            ModelID{ID: in.Id},
		Identifier:         in.Identifier,
		Description:        in.Description,
		ProductAttributeID: in.ProductAttributeID,
		ProductAttributeIdentifierAvailableCategorys: PBToProductAttributeIdentifierAvailableCategorys(in.ProductAttributeIdentifierAvailableCategorys),
	}
}

func ProductAttributeIdentifiersToPB(in []*ProductAttributeIdentifier) []*proto.ProductAttributeIdentifierInfo {
	var list []*proto.ProductAttributeIdentifierInfo
	for _, f := range in {
		list = append(list, ProductAttributeIdentifierToPB(f))
	}
	return list
}

func ProductAttributeIdentifierToPB(in *ProductAttributeIdentifier) *proto.ProductAttributeIdentifierInfo {
	if in == nil {
		return nil
	}
	var AvailableCategoryIDs []string
	for _, AvailableCategory := range in.ProductAttributeIdentifierAvailableCategorys {
		AvailableCategoryIDs = append(AvailableCategoryIDs, AvailableCategory.ProductCategoryID)
	}

	m := &proto.ProductAttributeIdentifierInfo{
		Id:                   in.ID,
		ProductAttributeID:   in.ProductAttributeID,
		Identifier:           in.Identifier,
		Description:          in.Description,
		AvailableCategoryIDs: AvailableCategoryIDs,
		ProductAttributeIdentifierAvailableCategorys: ProductAttributeIdentifierAvailableCategorysToPB(in.ProductAttributeIdentifierAvailableCategorys),
	}
	return m
}

func PBToProductAttributeIdentifierAvailableCategorys(in []*proto.ProductAttributeIdentifierAvailableCategoryInfo) []*ProductAttributeIdentifierAvailableCategory {
	var result []*ProductAttributeIdentifierAvailableCategory
	for _, c := range in {
		result = append(result, PBToProductAttributeIdentifierAvailableCategory(c))
	}
	return result
}

func PBToProductAttributeIdentifierAvailableCategory(in *proto.ProductAttributeIdentifierAvailableCategoryInfo) *ProductAttributeIdentifierAvailableCategory {
	if in == nil {
		return nil
	}
	return &ProductAttributeIdentifierAvailableCategory{
		ModelID:                      ModelID{ID: in.Id},
		ProductAttributeIdentifierID: in.ProductAttributeIdentifierID,
		ProductCategoryID:            in.ProductCategoryID,
	}
}

func ProductAttributeIdentifierAvailableCategorysToPB(in []*ProductAttributeIdentifierAvailableCategory) []*proto.ProductAttributeIdentifierAvailableCategoryInfo {
	var list []*proto.ProductAttributeIdentifierAvailableCategoryInfo
	for _, f := range in {
		list = append(list, ProductAttributeIdentifierAvailableCategoryToPB(f))
	}
	return list
}

func ProductAttributeIdentifierAvailableCategoryToPB(in *ProductAttributeIdentifierAvailableCategory) *proto.ProductAttributeIdentifierAvailableCategoryInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductAttributeIdentifierAvailableCategoryInfo{
		Id:                           in.ID,
		ProductAttributeIdentifierID: in.ProductAttributeIdentifierID,
		ProductCategoryID:            in.ProductCategoryID,
	}
	return m
}
