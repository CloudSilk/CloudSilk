package model

import "github.com/CloudSilk/CloudSilk/pkg/proto"

// 工步参数
type ProcessStepParameter struct {
	ModelID
	Code                       string                       `gorm:"index;size:50;comment:代号"`
	Priority                   int32                        `gorm:"comment:优先级"`
	Description                string                       `gorm:"size:500;comment:描述"`
	Enable                     bool                         `gorm:"comment:是否启用"`
	Remark                     string                       `gorm:"size:500;comment:备注"`
	ProductionLineID           string                       `gorm:"size:36;comment:产线ID"`
	ProductionLine             *ProductionLine              `gorm:""`                                                       //生产产线
	ProcessStepParameterValues []*ProcessStepParameterValue `gorm:"constraint:OnDelete:CASCADE"`                            //生产工步参数值
	AttributeExpressions       []*AttributeExpression       `gorm:"polymorphic:Rule;polymorphicValue:ProcessStepParameter"` //特性表达式
}

// 工步参数值
type ProcessStepParameterValue struct {
	ModelID
	ProcessStepParameterID     string                 `gorm:"index;size:36;comment:生产工步参数ID"`
	ProductionProcessID        string                 `gorm:"size:36;comment:生产工序ID"`
	ProductionProcessStepID    string                 `gorm:"size:36;comment:生产工步ID"`
	ProductionProcessStep      *ProductionProcessStep `gorm:""` //生产工步
	ProcessStepTypeParameterID string                 `gorm:"size:36;comment:生产工步类型参数ID"`
	StandardValue              string                 `gorm:"size:100;comment:标准值"`
	MaximumValue               string                 `gorm:"size:100;comment:最大值"`
	MinimumValue               string                 `gorm:"size:100;comment:最小值"`
}

func PBToProcessStepParameters(in []*proto.ProcessStepParameterInfo) []*ProcessStepParameter {
	var result []*ProcessStepParameter
	for _, c := range in {
		result = append(result, PBToProcessStepParameter(c))
	}
	return result
}

func PBToProcessStepParameter(in *proto.ProcessStepParameterInfo) *ProcessStepParameter {
	if in == nil {
		return nil
	}
	return &ProcessStepParameter{
		ModelID:                    ModelID{ID: in.Id},
		Code:                       in.Code,
		Priority:                   in.Priority,
		Description:                in.Description,
		Enable:                     in.Enable,
		Remark:                     in.Remark,
		ProductionLineID:           in.ProductionLineID,
		ProcessStepParameterValues: PBToProcessStepParameterValues(in.ProcessStepParameterValues),
		AttributeExpressions:       PBToAttributeExpressions(in.AttributeExpressions),
	}
}

func ProcessStepParametersToPB(in []*ProcessStepParameter) []*proto.ProcessStepParameterInfo {
	var list []*proto.ProcessStepParameterInfo
	for _, f := range in {
		list = append(list, ProcessStepParameterToPB(f))
	}
	return list
}

func ProcessStepParameterToPB(in *ProcessStepParameter) *proto.ProcessStepParameterInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProcessStepParameterInfo{
		Id:                         in.ID,
		Code:                       in.Code,
		Priority:                   in.Priority,
		Description:                in.Description,
		Enable:                     in.Enable,
		Remark:                     in.Remark,
		ProductionLineID:           in.ProductionLineID,
		ProductionLine:             ProductionLineToPB(in.ProductionLine),
		ProcessStepParameterValues: ProcessStepParameterValuesToPB(in.ProcessStepParameterValues),
		AttributeExpressions:       AttributeExpressionsToPB(in.AttributeExpressions),
	}
	return m
}

func PBToProcessStepParameterValues(in []*proto.ProcessStepParameterValueInfo) []*ProcessStepParameterValue {
	var result []*ProcessStepParameterValue
	for _, c := range in {
		result = append(result, PBToProcessStepParameterValue(c))
	}
	return result
}

func PBToProcessStepParameterValue(in *proto.ProcessStepParameterValueInfo) *ProcessStepParameterValue {
	if in == nil {
		return nil
	}
	return &ProcessStepParameterValue{
		ModelID:                    ModelID{ID: in.Id},
		ProductionProcessID:        in.ProductionProcessID,
		ProductionProcessStepID:    in.ProductionProcessStepID,
		ProcessStepTypeParameterID: in.ProcessStepTypeParameterID,
		StandardValue:              in.StandardValue,
		MaximumValue:               in.MaximumValue,
		MinimumValue:               in.MinimumValue,
	}
}

func ProcessStepParameterValuesToPB(in []*ProcessStepParameterValue) []*proto.ProcessStepParameterValueInfo {
	var list []*proto.ProcessStepParameterValueInfo
	for _, f := range in {
		list = append(list, ProcessStepParameterValueToPB(f))
	}
	return list
}

func ProcessStepParameterValueToPB(in *ProcessStepParameterValue) *proto.ProcessStepParameterValueInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProcessStepParameterValueInfo{
		Id:                         in.ID,
		ProcessStepParameterID:     in.ProcessStepParameterID,
		ProductionProcessID:        in.ProductionProcessID,
		ProductionProcessStepID:    in.ProductionProcessStepID,
		ProductionProcessStep:      ProductionProcessStepToPB(in.ProductionProcessStep),
		ProcessStepTypeParameterID: in.ProcessStepTypeParameterID,
		StandardValue:              in.StandardValue,
		MaximumValue:               in.MaximumValue,
		MinimumValue:               in.MinimumValue,
	}
	return m
}
