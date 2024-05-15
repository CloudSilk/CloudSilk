package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type PropertyExpression struct {
	ModelID
	SortIndex     int32  `json:"sortIndex" gorm:"comment:顺序"`
	PropertyName  string `json:"propertyName" gorm:"size:-1;comment:属性名称"`
	MathOperator  string `json:"mathOperator" gorm:"size:-1;comment:运算符"`
	PropertyValue string `json:"propertyValue" gorm:"size:-1;comment:属性值"`
	Remark        string `json:"remark" gorm:"size:1000;comment:备注"`
	RuleID        string `json:"ruleID" gorm:"size:36;comment:归属规则ID"`
	RuleType      string `json:"ruleType" gorm:"size:1000;comment:规则名称"`
}

func PBToPropertyExpressions(in []*proto.PropertyExpressionInfo) []*PropertyExpression {
	var result []*PropertyExpression
	for _, c := range in {
		result = append(result, PBToPropertyExpression(c))
	}
	return result
}

func PBToPropertyExpression(in *proto.PropertyExpressionInfo) *PropertyExpression {
	if in == nil {
		return nil
	}
	return &PropertyExpression{
		ModelID:       ModelID{ID: in.Id},
		SortIndex:     in.SortIndex,
		PropertyName:  in.PropertyName,
		MathOperator:  in.MathOperator,
		PropertyValue: in.PropertyValue,
		Remark:        in.Remark,
	}
}

func PropertyExpressionsToPB(in []*PropertyExpression) []*proto.PropertyExpressionInfo {
	var list []*proto.PropertyExpressionInfo
	for _, f := range in {
		list = append(list, PropertyExpressionToPB(f))
	}
	return list
}

func PropertyExpressionToPB(in *PropertyExpression) *proto.PropertyExpressionInfo {
	if in == nil {
		return nil
	}
	m := &proto.PropertyExpressionInfo{
		Id:            in.ID,
		SortIndex:     in.SortIndex,
		PropertyName:  in.PropertyName,
		MathOperator:  in.MathOperator,
		PropertyValue: in.PropertyValue,
		Remark:        in.Remark,
	}
	return m
}
