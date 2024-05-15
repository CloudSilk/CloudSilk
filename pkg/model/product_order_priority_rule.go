package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductOrderPriorityRule struct {
	ModelID
	Priority             int32                  `json:"priority" gorm:"comment:优先级"`
	PriorityLevel        int32                  `json:"priorityLevel" gorm:"comment:生产优先级"`
	Enable               bool                   `json:"enable" gorm:"comment:是否启用"`
	InitialValue         bool                   `json:"initialValue" gorm:"comment:默认排序"`
	Remark               string                 `json:"remark" gorm:"size:1000;comment:备注"`
	AttributeExpressions []*AttributeExpression `json:"attributeExpressions" gorm:"polymorphic:Rule;polymorphicValue:ProductOrderPriorityRule"` //特征表达式
}

func PBToProductOrderPriorityRules(in []*proto.ProductOrderPriorityRuleInfo) []*ProductOrderPriorityRule {
	var result []*ProductOrderPriorityRule
	for _, c := range in {
		result = append(result, PBToProductOrderPriorityRule(c))
	}
	return result
}

func PBToProductOrderPriorityRule(in *proto.ProductOrderPriorityRuleInfo) *ProductOrderPriorityRule {
	if in == nil {
		return nil
	}
	return &ProductOrderPriorityRule{
		ModelID:              ModelID{ID: in.Id},
		Priority:             in.Priority,
		PriorityLevel:        in.PriorityLevel,
		Enable:               in.Enable,
		InitialValue:         in.InitialValue,
		Remark:               in.Remark,
		AttributeExpressions: PBToAttributeExpressions(in.AttributeExpressions),
	}
}

func ProductOrderPriorityRulesToPB(in []*ProductOrderPriorityRule) []*proto.ProductOrderPriorityRuleInfo {
	var list []*proto.ProductOrderPriorityRuleInfo
	for _, f := range in {
		list = append(list, ProductOrderPriorityRuleToPB(f))
	}
	return list
}

func ProductOrderPriorityRuleToPB(in *ProductOrderPriorityRule) *proto.ProductOrderPriorityRuleInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductOrderPriorityRuleInfo{
		Id:                   in.ID,
		Priority:             in.Priority,
		PriorityLevel:        in.PriorityLevel,
		Enable:               in.Enable,
		InitialValue:         in.InitialValue,
		Remark:               in.Remark,
		AttributeExpressions: AttributeExpressionsToPB(in.AttributeExpressions),
	}
	return m
}
