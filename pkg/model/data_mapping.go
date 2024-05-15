package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 数据词条
type DataMapping struct {
	ModelID
	Group       string `json:"group" gorm:"size:100;comment:分组"`
	Code        string `json:"code" gorm:"size:100;comment:代号"`
	Description string `json:"description" gorm:"size:1000;comment:描述"`
	Remark      string `json:"remark" gorm:"size:1000;comment:备注"`
}

func PBToDataMappings(in []*proto.DataMappingInfo) []*DataMapping {
	var result []*DataMapping
	for _, c := range in {
		result = append(result, PBToDataMapping(c))
	}
	return result
}

func PBToDataMapping(in *proto.DataMappingInfo) *DataMapping {
	if in == nil {
		return nil
	}
	return &DataMapping{
		ModelID:     ModelID{ID: in.Id},
		Group:       in.Group,
		Code:        in.Code,
		Description: in.Description,
		Remark:      in.Remark,
	}
}

func DataMappingsToPB(in []*DataMapping) []*proto.DataMappingInfo {
	var list []*proto.DataMappingInfo
	for _, f := range in {
		list = append(list, DataMappingToPB(f))
	}
	return list
}

func DataMappingToPB(in *DataMapping) *proto.DataMappingInfo {
	if in == nil {
		return nil
	}
	m := &proto.DataMappingInfo{
		Id:          in.ID,
		Group:       in.Group,
		Code:        in.Code,
		Description: in.Description,
		Remark:      in.Remark,
	}
	return m
}
