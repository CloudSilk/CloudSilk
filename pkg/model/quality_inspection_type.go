package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 检验类型（来料检验/过程检验/成品检验等）
type QualityInspectionType struct {
	ModelID
	Code        string `json:"code" gorm:"index;size:100;comment:代号"`
	Description string `json:"description" gorm:"size:1000;comment:描述"`
	Remark      string `json:"remark" gorm:"size:1000;comment:备注"`
	Enable      bool   `json:"enable" gorm:"default:true;comment:是否启用"`
}

func PBToQualityInspectionTypes(in []*proto.QualityInspectionTypeInfo) []*QualityInspectionType {
	var result []*QualityInspectionType
	for _, c := range in {
		result = append(result, PBToQualityInspectionType(c))
	}
	return result
}

func PBToQualityInspectionType(in *proto.QualityInspectionTypeInfo) *QualityInspectionType {
	if in == nil {
		return nil
	}
	return &QualityInspectionType{
		ModelID:     ModelID{ID: in.Id},
		Code:        in.Code,
		Description: in.Description,
		Remark:      in.Remark,
		Enable:      in.Enable,
	}
}

func QualityInspectionTypesToPB(in []*QualityInspectionType) []*proto.QualityInspectionTypeInfo {
	var list []*proto.QualityInspectionTypeInfo
	for _, f := range in {
		list = append(list, QualityInspectionTypeToPB(f))
	}
	return list
}

func QualityInspectionTypeToPB(in *QualityInspectionType) *proto.QualityInspectionTypeInfo {
	if in == nil {
		return nil
	}
	return &proto.QualityInspectionTypeInfo{
		Id:          in.ID,
		Code:        in.Code,
		Description: in.Description,
		Remark:      in.Remark,
		Enable:      in.Enable,
	}
}
