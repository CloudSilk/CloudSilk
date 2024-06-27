package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 标签适配规则
type LabelAdaptationRule struct {
	ModelID
	Priority             int32                  `gorm:"comment:优先级"`
	LabelTemplateID      *string                `gorm:"size:36;comment:适配标签模板ID"`
	LabelTemplate        *LabelTemplate         `gorm:"constraint:OnDelete:SET NULL"` //适配标签模板
	Enable               bool                   `gorm:"comment:是否启用"`
	InitialValue         bool                   `gorm:"comment:默认适配"`
	DoubleCheck          bool                   `gorm:"comment:需要复核"`
	Remark               string                 `gorm:"size:500;comment:备注"`
	AttributeExpressions []*AttributeExpression `json:"attributeExpressions" gorm:"polymorphic:Rule;polymorphicValue:LabelAdaptationRule"` //特征表达式
}

func PBToLabelAdaptationRules(in []*proto.LabelAdaptationRuleInfo) []*LabelAdaptationRule {
	var result []*LabelAdaptationRule
	for _, c := range in {
		result = append(result, PBToLabelAdaptationRule(c))
	}
	return result
}

func PBToLabelAdaptationRule(in *proto.LabelAdaptationRuleInfo) *LabelAdaptationRule {
	if in == nil {
		return nil
	}

	var labelTemplateID *string
	if in.LabelTemplateID != "" {
		labelTemplateID = &in.LabelTemplateID
	}

	return &LabelAdaptationRule{
		ModelID:              ModelID{ID: in.Id},
		Priority:             in.Priority,
		LabelTemplateID:      labelTemplateID,
		Enable:               in.Enable,
		InitialValue:         in.InitialValue,
		DoubleCheck:          in.DoubleCheck,
		Remark:               in.Remark,
		AttributeExpressions: PBToAttributeExpressions(in.AttributeExpressions),
	}
}

func LabelAdaptationRulesToPB(in []*LabelAdaptationRule) []*proto.LabelAdaptationRuleInfo {
	var list []*proto.LabelAdaptationRuleInfo
	for _, f := range in {
		list = append(list, LabelAdaptationRuleToPB(f))
	}
	return list
}

func LabelAdaptationRuleToPB(in *LabelAdaptationRule) *proto.LabelAdaptationRuleInfo {
	if in == nil {
		return nil
	}

	var labelTemplateID string
	if in.LabelTemplateID != nil {
		labelTemplateID = *in.LabelTemplateID
	}

	m := &proto.LabelAdaptationRuleInfo{
		Id:                   in.ID,
		Priority:             in.Priority,
		LabelTemplateID:      labelTemplateID,
		Enable:               in.Enable,
		InitialValue:         in.InitialValue,
		DoubleCheck:          in.DoubleCheck,
		Remark:               in.Remark,
		AttributeExpressions: AttributeExpressionsToPB(in.AttributeExpressions),
	}
	return m
}
