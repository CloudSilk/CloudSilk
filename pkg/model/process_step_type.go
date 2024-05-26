package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 工步类型
type ProcessStepType struct {
	ModelID
	Code                      string                      `json:"code" gorm:"index;size:100;comment:代号"`
	Description               string                      `json:"description" gorm:"size:1000;comment:描述"`
	Enable                    bool                        `json:"enable" gorm:"comment:是否启用"`
	Remark                    string                      `json:"remark" gorm:"size:1000;comment:备注"`
	ProcessStepTypeParameters []*ProcessStepTypeParameter `json:"processStepTypeParameters" gorm:"constraint:OnDelete:CASCADE;"` // 工步类型参数
}

// 工步类型参数
type ProcessStepTypeParameter struct {
	ModelID
	ProcessStepTypeID string `json:"processStepTypeID" gorm:"index;size:36;comment:工步类型ID"`
	Code              string `json:"code" gorm:"size:100;comment:代号"`
	Description       string `json:"description" gorm:"size:1000;comment:描述"`
	DefaultValue      string `json:"defaultValue" gorm:"size:100;comment:默认值"`
	StandardValue     string `json:"standardValue" gorm:"-"` //标准值
	MaximumValue      string `json:"maximumValue" gorm:"size:100;comment:最大值"`
	MinimumValue      string `json:"minimumValue" gorm:"size:100;comment:最小值"`
	Unit              string `json:"unit" gorm:"size:100;comment:单位"`
	Required          bool   `json:"required" gorm:"comment:是否必填"`
	BoundsRequired    bool   `json:"boundsRequired" gorm:"comment:是否上下限"`
	ParameterType     bool   `json:"parameterType" gorm:"comment:参数类型 true-输入 false-输出"`
	Remark            string `json:"remark" gorm:"size:1000;comment:备注"`
}

func PBToProcessStepTypes(in []*proto.ProcessStepTypeInfo) []*ProcessStepType {
	var result []*ProcessStepType
	for _, c := range in {
		result = append(result, PBToProcessStepType(c))
	}
	return result
}

func PBToProcessStepType(in *proto.ProcessStepTypeInfo) *ProcessStepType {
	if in == nil {
		return nil
	}
	return &ProcessStepType{
		ModelID:                   ModelID{ID: in.Id},
		Code:                      in.Code,
		Description:               in.Description,
		Enable:                    in.Enable,
		Remark:                    in.Remark,
		ProcessStepTypeParameters: PBToProcessStepTypeParameters(in.ProcessStepTypeParameters),
	}
}

func ProcessStepTypesToPB(in []*ProcessStepType) []*proto.ProcessStepTypeInfo {
	var list []*proto.ProcessStepTypeInfo
	for _, f := range in {
		list = append(list, ProcessStepTypeToPB(f))
	}
	return list
}

func ProcessStepTypeToPB(in *ProcessStepType) *proto.ProcessStepTypeInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProcessStepTypeInfo{
		Id:                        in.ID,
		Code:                      in.Code,
		Description:               in.Description,
		Enable:                    in.Enable,
		Remark:                    in.Remark,
		ProcessStepTypeParameters: ProcessStepTypeParametersToPB(in.ProcessStepTypeParameters),
	}
	return m
}

func PBToProcessStepTypeParameters(in []*proto.ProcessStepTypeParameterInfo) []*ProcessStepTypeParameter {
	var result []*ProcessStepTypeParameter
	for _, c := range in {
		result = append(result, PBToProcessStepTypeParameter(c))
	}
	return result
}

func PBToProcessStepTypeParameter(in *proto.ProcessStepTypeParameterInfo) *ProcessStepTypeParameter {
	if in == nil {
		return nil
	}
	return &ProcessStepTypeParameter{
		ModelID:        ModelID{ID: in.Id},
		Code:           in.Code,
		Description:    in.Description,
		DefaultValue:   in.DefaultValue,
		StandardValue:  in.StandardValue,
		MaximumValue:   in.MaximumValue,
		MinimumValue:   in.MinimumValue,
		Unit:           in.Unit,
		Required:       in.Required,
		BoundsRequired: in.BoundsRequired,
		ParameterType:  in.ParameterType,
		Remark:         in.Remark,
	}
}

func ProcessStepTypeParametersToPB(in []*ProcessStepTypeParameter) []*proto.ProcessStepTypeParameterInfo {
	var list []*proto.ProcessStepTypeParameterInfo
	for _, f := range in {
		list = append(list, ProcessStepTypeParameterToPB(f))
	}
	return list
}

func ProcessStepTypeParameterToPB(in *ProcessStepTypeParameter) *proto.ProcessStepTypeParameterInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProcessStepTypeParameterInfo{
		Id:                in.ID,
		ProcessStepTypeID: in.ProcessStepTypeID,
		Code:              in.Code,
		Description:       in.Description,
		DefaultValue:      in.DefaultValue,
		StandardValue:     in.StandardValue,
		MaximumValue:      in.MaximumValue,
		MinimumValue:      in.MinimumValue,
		Unit:              in.Unit,
		Required:          in.Required,
		BoundsRequired:    in.BoundsRequired,
		ParameterType:     in.ParameterType,
		Remark:            in.Remark,
	}
	return m
}
