package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 标签模板
type LabelTemplate struct {
	ModelID
	Code            string            `json:"code" gorm:"size:200;comment:代号"`
	Description     string            `json:"description" gorm:"size:200;comment:描述"`
	FilePath        string            `json:"filePath" gorm:"size:200;comment:文件路径"`
	Remark          string            `json:"remark" gorm:"size:200;comment:备注"`
	LabelTypeID     *string           `json:"labelTypeID" gorm:"size:36;comment:标签类型ID;"`
	LabelType       *LabelType        `json:"labelType" gorm:"constraint:OnDelete:SET NULL"`        //标签类型
	LabelParameters []*LabelParameter `json:"labelParameters" gorm:"constraint:OnDelete:CASCADE;"` //参数
}

// 标签模板参数
type LabelParameter struct {
	ModelID
	LabelTemplateID string `json:"labelTemplateID" gorm:"index;size:36;comment:标签模板ID"`
	Name            string `json:"name" gorm:"size:200;comment:名称"`
	DefaultValue    string `json:"defaultValue" gorm:"size:200;comment:预设值"`
}

func PBToLabelTemplates(in []*proto.LabelTemplateInfo) []*LabelTemplate {
	var result []*LabelTemplate
	for _, c := range in {
		result = append(result, PBToLabelTemplate(c))
	}
	return result
}

func PBToLabelTemplate(in *proto.LabelTemplateInfo) *LabelTemplate {
	if in == nil {
		return nil
	}

	var labelTypeID *string
	if in.LabelTypeID != "" {
		labelTypeID = &in.LabelTypeID
	}
	return &LabelTemplate{
		ModelID:         ModelID{ID: in.Id},
		Code:            in.Code,
		Description:     in.Description,
		FilePath:        in.FilePath,
		Remark:          in.Remark,
		LabelTypeID:     labelTypeID,
		LabelParameters: PBToLabelParameters(in.LabelParameters),
	}
}

func LabelTemplatesToPB(in []*LabelTemplate) []*proto.LabelTemplateInfo {
	var list []*proto.LabelTemplateInfo
	for _, f := range in {
		list = append(list, LabelTemplateToPB(f))
	}
	return list
}

func LabelTemplateToPB(in *LabelTemplate) *proto.LabelTemplateInfo {
	if in == nil {
		return nil
	}
	var labelTypeID string
	if in.LabelTypeID != nil {
		labelTypeID = *in.LabelTypeID
	}
	m := &proto.LabelTemplateInfo{
		Id:              in.ID,
		Code:            in.Code,
		Description:     in.Description,
		FilePath:        in.FilePath,
		Remark:          in.Remark,
		LabelTypeID:     labelTypeID,
		LabelParameters: LabelParametersToPB(in.LabelParameters),
		LabelType:       LabelTypeToPB(in.LabelType),
	}
	return m
}

func PBToLabelParameters(in []*proto.LabelParameterInfo) []*LabelParameter {
	var result []*LabelParameter
	for _, c := range in {
		result = append(result, PBToLabelParameter(c))
	}
	return result
}

func PBToLabelParameter(in *proto.LabelParameterInfo) *LabelParameter {
	if in == nil {
		return nil
	}
	return &LabelParameter{
		ModelID:         ModelID{ID: in.Id},
		Name:            in.Name,
		DefaultValue:    in.DefaultValue,
		LabelTemplateID: in.LabelTemplateID,
	}
}

func LabelParametersToPB(in []*LabelParameter) []*proto.LabelParameterInfo {
	var list []*proto.LabelParameterInfo
	for _, f := range in {
		list = append(list, LabelParameterToPB(f))
	}
	return list
}

func LabelParameterToPB(in *LabelParameter) *proto.LabelParameterInfo {
	if in == nil {
		return nil
	}
	m := &proto.LabelParameterInfo{
		Id:              in.ID,
		Name:            in.Name,
		DefaultValue:    in.DefaultValue,
		LabelTemplateID: in.LabelTemplateID,
	}
	return m
}
