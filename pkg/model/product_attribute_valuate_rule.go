package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 特性赋值规则
type ProductAttributeValuateRule struct {
	ModelID
	Priority            int32                 `json:"priority" gorm:"comment:优先级"`
	ProductCategoryID   *string               `json:"productCategoryID" gorm:"size:36;comment:产品类别ID"`
	ProductCategory     *ProductCategory      `json:"productCategory"`
	ProductAttributeID  string                `json:"productAttributeID" gorm:"size:36;comment:目标特性ID"`
	ProductAttribute    *ProductAttribute     `json:"productAttribute"`
	Value               string                `json:"value" gorm:"size:200;comment:设定值"`
	Description         string                `json:"description" gorm:"size:1000;comment:值描述"`
	Enable              bool                  `json:"enable" gorm:"comment:是否启用"`
	InitialValue        bool                  `json:"initialValue" gorm:"comment:默认赋值"`
	Remark              string                `json:"remark" gorm:"size:1000;comment:备注"`
	PropertyExpressions []*PropertyExpression `json:"propertyExpressions" gorm:"polymorphic:Rule;polymorphicValue:ProductAttributeValuateRule"` //属性表达式
}

func PBToProductAttributeValuateRules(in []*proto.ProductAttributeValuateRuleInfo) []*ProductAttributeValuateRule {
	var result []*ProductAttributeValuateRule
	for _, c := range in {
		result = append(result, PBToProductAttributeValuateRule(c))
	}
	return result
}

func PBToProductAttributeValuateRule(in *proto.ProductAttributeValuateRuleInfo) *ProductAttributeValuateRule {
	if in == nil {
		return nil
	}

	var productCategoryID *string
	if in.ProductCategoryID != "" {
		productCategoryID = &in.ProductCategoryID
	}
	return &ProductAttributeValuateRule{
		ModelID:             ModelID{ID: in.Id},
		Priority:            in.Priority,
		ProductCategoryID:   productCategoryID,
		ProductAttributeID:  in.ProductAttributeID,
		Value:               in.Value,
		Description:         in.Description,
		Enable:              in.Enable,
		InitialValue:        in.InitialValue,
		Remark:              in.Remark,
		PropertyExpressions: PBToPropertyExpressions(in.PropertyExpressions),
	}
}

func ProductAttributeValuateRulesToPB(in []*ProductAttributeValuateRule) []*proto.ProductAttributeValuateRuleInfo {
	var list []*proto.ProductAttributeValuateRuleInfo
	for _, f := range in {
		list = append(list, ProductAttributeValuateRuleToPB(f))
	}
	return list
}

func ProductAttributeValuateRuleToPB(in *ProductAttributeValuateRule) *proto.ProductAttributeValuateRuleInfo {
	if in == nil {
		return nil
	}

	var productCategoryID string
	if in.ProductCategoryID != nil {
		productCategoryID = *in.ProductCategoryID
	}
	m := &proto.ProductAttributeValuateRuleInfo{
		PropertyExpressions: PropertyExpressionsToPB(in.PropertyExpressions),
		Id:                  in.ID,
		Priority:            in.Priority,
		ProductCategoryID:   productCategoryID,
		ProductAttributeID:  in.ProductAttributeID,
		Value:               in.Value,
		Description:         in.Description,
		Enable:              in.Enable,
		InitialValue:        in.InitialValue,
		Remark:              in.Remark,
		ProductCategory:     ProductCategoryToPB(in.ProductCategory),
		ProductAttribute:    ProductAttributeToPB(in.ProductAttribute),
	}
	return m
}
