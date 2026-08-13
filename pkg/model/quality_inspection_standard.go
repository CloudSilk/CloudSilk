package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 检验标准（检验项主数据：规格限/单位/AQL）
type QualityInspectionStandard struct {
	ModelID
	Code                   string                 `json:"code" gorm:"index;size:100;comment:代号"`
	Description            string                 `json:"description" gorm:"size:1000;comment:描述"`
	Unit                   string                 `json:"unit" gorm:"size:100;comment:单位"`
	StandardValue          string                 `json:"standardValue" gorm:"size:100;comment:标准值"`
	LowerLimit             string                 `json:"lowerLimit" gorm:"size:100;comment:规格下限"`
	UpperLimit             string                 `json:"upperLimit" gorm:"size:100;comment:规格上限"`
	CriterionMethod        string                 `json:"criterionMethod" gorm:"size:100;comment:比较方式"`
	Aql                    float64                `json:"aql" gorm:"comment:AQL接收质量限"`
	SampleLevel            string                 `json:"sampleLevel" gorm:"size:10;comment:样本量字码"`
	QualityInspectionTypeID *string               `json:"qualityInspectionTypeID" gorm:"size:36;comment:检验类型ID"`
	QualityInspectionType  *QualityInspectionType `json:"qualityInspectionType" gorm:"constraint:OnDelete:SET NULL"`
	Enable                 bool                   `json:"enable" gorm:"default:true;comment:是否启用"`
	Remark                 string                 `json:"remark" gorm:"size:1000;comment:备注"`
}

func PBToQualityInspectionStandards(in []*proto.QualityInspectionStandardInfo) []*QualityInspectionStandard {
	var result []*QualityInspectionStandard
	for _, c := range in {
		result = append(result, PBToQualityInspectionStandard(c))
	}
	return result
}

func PBToQualityInspectionStandard(in *proto.QualityInspectionStandardInfo) *QualityInspectionStandard {
	if in == nil {
		return nil
	}
	var qualityInspectionTypeID *string
	if in.QualityInspectionTypeID != "" {
		qualityInspectionTypeID = &in.QualityInspectionTypeID
	}
	return &QualityInspectionStandard{
		ModelID:                ModelID{ID: in.Id},
		Code:                   in.Code,
		Description:            in.Description,
		Unit:                   in.Unit,
		StandardValue:          in.StandardValue,
		LowerLimit:             in.LowerLimit,
		UpperLimit:             in.UpperLimit,
		CriterionMethod:        in.CriterionMethod,
		Aql:                    in.Aql,
		SampleLevel:            in.SampleLevel,
		QualityInspectionTypeID: qualityInspectionTypeID,
		Enable:                 in.Enable,
		Remark:                 in.Remark,
	}
}

func QualityInspectionStandardsToPB(in []*QualityInspectionStandard) []*proto.QualityInspectionStandardInfo {
	var list []*proto.QualityInspectionStandardInfo
	for _, f := range in {
		list = append(list, QualityInspectionStandardToPB(f))
	}
	return list
}

func QualityInspectionStandardToPB(in *QualityInspectionStandard) *proto.QualityInspectionStandardInfo {
	if in == nil {
		return nil
	}
	m := &proto.QualityInspectionStandardInfo{
		Id:          in.ID,
		Code:        in.Code,
		Description: in.Description,
		Unit:        in.Unit,
		StandardValue: in.StandardValue,
		LowerLimit:  in.LowerLimit,
		UpperLimit:  in.UpperLimit,
		CriterionMethod: in.CriterionMethod,
		Aql:         in.Aql,
		SampleLevel: in.SampleLevel,
		Enable:      in.Enable,
		Remark:      in.Remark,
	}
	if in.QualityInspectionTypeID != nil {
		m.QualityInspectionTypeID = *in.QualityInspectionTypeID
	}
	if in.QualityInspectionType != nil {
		m.QualityInspectionTypeCode = in.QualityInspectionType.Code
	}
	return m
}
