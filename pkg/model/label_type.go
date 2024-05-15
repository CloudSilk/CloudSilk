package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 标签类型
type LabelType struct {
	ModelID
	Code        string `json:"code" gorm:"size:100;comment:类型"`
	Description string `json:"description" gorm:"size:1000;comment:描述"`
	Remark      string `json:"remark" gorm:"size:1000;comment:备注"`
}

func PBToLabelTypes(in []*proto.LabelTypeInfo) []*LabelType {
	var result []*LabelType
	for _, c := range in {
		result = append(result, PBToLabelType(c))
	}
	return result
}

func PBToLabelType(in *proto.LabelTypeInfo) *LabelType {
	if in == nil {
		return nil
	}
	return &LabelType{
		ModelID:     ModelID{ID: in.Id},
		Code:        in.Code,
		Description: in.Description,
		Remark:      in.Remark,
	}
}

func LabelTypesToPB(in []*LabelType) []*proto.LabelTypeInfo {
	var list []*proto.LabelTypeInfo
	for _, f := range in {
		list = append(list, LabelTypeToPB(f))
	}
	return list
}

func LabelTypeToPB(in *LabelType) *proto.LabelTypeInfo {
	if in == nil {
		return nil
	}
	m := &proto.LabelTypeInfo{
		Id:          in.ID,
		Code:        in.Code,
		Description: in.Description,
		Remark:      in.Remark,
	}
	return m
}
