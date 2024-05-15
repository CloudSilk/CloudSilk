package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 工步匹配规则
type ProcessStepMatchRule struct {
	ModelID
	ProductionLineID        string                 `json:"productionLineID" gorm:"size:36;comment:生产产线ID"`
	Priority                int32                  `json:"priority" gorm:"comment:优先级"`
	GroupCode               string                 `json:"groupCode" gorm:"size:100;comment:采集组"`
	Enable                  bool                   `json:"enable" gorm:"comment:是否启用"`
	InitialValue            bool                   `json:"initialValue" gorm:"comment:默认匹配"`
	Description             string                 `json:"description" gorm:"size:1000;comment:描述"`
	Remark                  string                 `json:"remark" gorm:"size:1000;comment:备注"`
	AttributeExpressions    []*AttributeExpression `json:"attributeExpressions" gorm:"polymorphic:Rule;polymorphicValue:ProcessStepMatchRule"` //特性表达式
	ProductionProcessStepID string                 `json:"productionProcessStepID" gorm:"size:36;comment:生产工步ID"`
	ProductionProcessStep   *ProductionProcessStep `json:"productionProcessStep" gorm:""` //生产工步
}

func PBToProcessStepMatchRules(in []*proto.ProcessStepMatchRuleInfo) []*ProcessStepMatchRule {
	var result []*ProcessStepMatchRule
	for _, c := range in {
		result = append(result, PBToProcessStepMatchRule(c))
	}
	return result
}

func PBToProcessStepMatchRule(in *proto.ProcessStepMatchRuleInfo) *ProcessStepMatchRule {
	if in == nil {
		return nil
	}
	return &ProcessStepMatchRule{
		ModelID:                 ModelID{ID: in.Id},
		ProductionLineID:        in.ProductionLineID,
		Priority:                in.Priority,
		GroupCode:               in.GroupCode,
		Enable:                  in.Enable,
		InitialValue:            in.InitialValue,
		Description:             in.Description,
		Remark:                  in.Remark,
		AttributeExpressions:    PBToAttributeExpressions(in.AttributeExpressions),
		ProductionProcessStepID: in.ProductionProcessStepID,
	}
}

func ProcessStepMatchRulesToPB(in []*ProcessStepMatchRule) []*proto.ProcessStepMatchRuleInfo {
	var list []*proto.ProcessStepMatchRuleInfo
	for _, f := range in {
		list = append(list, ProcessStepMatchRuleToPB(f))
	}
	return list
}

func ProcessStepMatchRuleToPB(in *ProcessStepMatchRule) *proto.ProcessStepMatchRuleInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProcessStepMatchRuleInfo{
		Id:                      in.ID,
		ProductionLineID:        in.ProductionLineID,
		Priority:                in.Priority,
		GroupCode:               in.GroupCode,
		Enable:                  in.Enable,
		InitialValue:            in.InitialValue,
		Description:             in.Description,
		Remark:                  in.Remark,
		AttributeExpressions:    AttributeExpressionsToPB(in.AttributeExpressions),
		ProductionProcessStepID: in.ProductionProcessStepID,
		ProductionProcessStep:   ProductionProcessStepToPB(in.ProductionProcessStep),
	}
	return m
}
