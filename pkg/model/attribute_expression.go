package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 特性表达式
type AttributeExpression struct {
	ModelID
	SortIndex          int32             `json:"sortIndex" gorm:"comment:优先级"`
	ProductAttributeID string            `json:"productAttributeID" gorm:"index;size:36;comment:产品特性ID;"`
	ProductAttribute   *ProductAttribute `json:"productAttribute" gorm:"constraint:OnDelete:CASCADE"`
	MathOperator       string            `json:"mathOperator" gorm:"size:-1;comment:运算符"`
	AttributeValue     string            `json:"attributeValue" gorm:"size:-1;comment:特性值"`
	Remark             string            `json:"remark" gorm:"size:500;comment:备注"`
	RuleID             string            `json:"ruleID" gorm:"index;size:36;comment:归属规则ID"`
	RuleType           string            `json:"ruleType" gorm:"index;size:500;comment:规则名称"`
}

// 特性表达式组
type AttributeExpressionGroup struct {
	ModelID
	Operator   string
	RuleID     string `json:"ruleID" gorm:"index;size:36;comment:归属规则ID"`
	RuleType   string `json:"ruleType" gorm:"index;size:500;comment:规则名称"`
	DateSource []*AttributeExpressionAttribute
	Children   []*AttributeExpression
}

// 特性表达式组参数
type AttributeExpressionAttribute struct {
	ModelID
	AttributeExpressionGroupID string
	SortIndex                  int32  `json:"sortIndex" gorm:"comment:优先级"`
	ProductAttributeID         string `json:"productAttributeID" gorm:"index;size:36;comment:产品特性ID;"`
	MathOperator               string `json:"mathOperator" gorm:"size:-1;comment:运算符"`
	AttributeValue             string `json:"attributeValue" gorm:"size:-1;comment:特性值"`
	Remark                     string `json:"remark" gorm:"size:500;comment:备注"`
}

func PBToAttributeExpressions(in []*proto.AttributeExpressionInfo) []*AttributeExpression {
	var result []*AttributeExpression
	for _, c := range in {
		result = append(result, PBToAttributeExpression(c))
	}
	return result
}

func PBToAttributeExpression(in *proto.AttributeExpressionInfo) *AttributeExpression {
	if in == nil {
		return nil
	}
	return &AttributeExpression{
		ModelID:            ModelID{ID: in.Id},
		SortIndex:          in.SortIndex,
		ProductAttributeID: in.ProductAttributeID,
		MathOperator:       in.MathOperator,
		AttributeValue:     in.AttributeValue,
		Remark:             in.Remark,
	}
}

func AttributeExpressionsToPB(in []*AttributeExpression) []*proto.AttributeExpressionInfo {
	var list []*proto.AttributeExpressionInfo
	for _, f := range in {
		list = append(list, AttributeExpressionToPB(f))
	}
	return list
}

func AttributeExpressionToPB(in *AttributeExpression) *proto.AttributeExpressionInfo {
	if in == nil {
		return nil
	}
	m := &proto.AttributeExpressionInfo{
		Id:                 in.ID,
		SortIndex:          in.SortIndex,
		ProductAttributeID: in.ProductAttributeID,
		MathOperator:       in.MathOperator,
		AttributeValue:     in.AttributeValue,
		Remark:             in.Remark,
	}
	return m
}
