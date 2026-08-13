package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 物料退料类型
type MaterialReturnType struct {
	ModelID
	Code        string `gorm:"size:50;comment:代号"`  //代号
	Description string `gorm:"size:500;comment:描述"` //描述
	Remark      string `gorm:"size:500;comment:备注"` //备注
}

func PBToMaterialReturnTypes(in []*proto.MaterialReturnTypeInfo) []*MaterialReturnType {
	var result []*MaterialReturnType
	for _, c := range in {
		result = append(result, PBToMaterialReturnType(c))
	}
	return result
}

func PBToMaterialReturnType(in *proto.MaterialReturnTypeInfo) *MaterialReturnType {
	if in == nil {
		return nil
	}

	return &MaterialReturnType{
		ModelID:     ModelID{ID: in.Id},
		Code:        in.Code,
		Description: in.Description,
		Remark:      in.Remark,
	}
}

func MaterialReturnTypesToPB(in []*MaterialReturnType) []*proto.MaterialReturnTypeInfo {
	var list []*proto.MaterialReturnTypeInfo
	for _, f := range in {
		list = append(list, MaterialReturnTypeToPB(f))
	}
	return list
}

func MaterialReturnTypeToPB(in *MaterialReturnType) *proto.MaterialReturnTypeInfo {
	if in == nil {
		return nil
	}

	m := &proto.MaterialReturnTypeInfo{
		Id:          in.ID,
		Code:        in.Code,
		Description: in.Description,
		Remark:      in.Remark,
	}
	return m
}
