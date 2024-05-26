package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductOrderReleaseRule struct {
	ModelID
	Priority             int32                  `json:"priority" gorm:"comment:优先级"`
	ProductionLineID     string                 `json:"productionLineID" gorm:"size:36;comment:发放产线ID"`
	ProductionLine       *ProductionLine        `json:"productionLine" gorm:"constraint:OnDelete:CASCADE"`
	Enable               bool                   `json:"enable" gorm:"comment:是否启用"`
	InitialValue         bool                   `json:"initialValue" gorm:"comment:默认发放"`
	Remark               string                 `json:"remark" gorm:"size:1000;comment:备注"`
	AttributeExpressions []*AttributeExpression `json:"attributeExpressions" gorm:"polymorphic:Rule;polymorphicValue:ProductOrderReleaseRule"` //特征表达式
}

func PBToProductOrderReleaseRules(in []*proto.ProductOrderReleaseRuleInfo) []*ProductOrderReleaseRule {
	var result []*ProductOrderReleaseRule
	for _, c := range in {
		result = append(result, PBToProductOrderReleaseRule(c))
	}
	return result
}

func PBToProductOrderReleaseRule(in *proto.ProductOrderReleaseRuleInfo) *ProductOrderReleaseRule {
	if in == nil {
		return nil
	}
	return &ProductOrderReleaseRule{
		ModelID:              ModelID{ID: in.Id},
		Priority:             in.Priority,
		ProductionLineID:     in.ProductionLineID,
		Enable:               in.Enable,
		InitialValue:         in.InitialValue,
		Remark:               in.Remark,
		AttributeExpressions: PBToAttributeExpressions(in.AttributeExpressions),
	}
}

func ProductOrderReleaseRulesToPB(in []*ProductOrderReleaseRule) []*proto.ProductOrderReleaseRuleInfo {
	var list []*proto.ProductOrderReleaseRuleInfo
	for _, f := range in {
		list = append(list, ProductOrderReleaseRuleToPB(f))
	}
	return list
}

func ProductOrderReleaseRuleToPB(in *ProductOrderReleaseRule) *proto.ProductOrderReleaseRuleInfo {
	if in == nil {
		return nil
	}

	m := &proto.ProductOrderReleaseRuleInfo{
		Id:                   in.ID,
		Priority:             in.Priority,
		ProductionLineID:     in.ProductionLineID,
		ProductionLine:       ProductionLineToPB(in.ProductionLine),
		Enable:               in.Enable,
		InitialValue:         in.InitialValue,
		Remark:               in.Remark,
		AttributeExpressions: AttributeExpressionsToPB(in.AttributeExpressions),
	}
	return m
}
