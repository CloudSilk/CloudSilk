package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 物料类别
type MaterialCategory struct {
	ModelID
	Code        string `json:"code" gorm:"size:50;comment:代号"`
	Description string `json:"description" gorm:"size:100;comment:描述"`
	Remark      string `json:"remark" gorm:"size:500;comment:备注"`
}

func PBToMaterialCategorys(in []*proto.MaterialCategoryInfo) []*MaterialCategory {
	var result []*MaterialCategory
	for _, c := range in {
		result = append(result, PBToMaterialCategory(c))
	}
	return result
}

func PBToMaterialCategory(in *proto.MaterialCategoryInfo) *MaterialCategory {
	if in == nil {
		return nil
	}
	return &MaterialCategory{
		ModelID:     ModelID{ID: in.Id},
		Code:        in.Code,
		Description: in.Description,
		Remark:      in.Remark,
	}
}

func MaterialCategorysToPB(in []*MaterialCategory) []*proto.MaterialCategoryInfo {
	var list []*proto.MaterialCategoryInfo
	for _, f := range in {
		list = append(list, MaterialCategoryToPB(f))
	}
	return list
}

func MaterialCategoryToPB(in *MaterialCategory) *proto.MaterialCategoryInfo {
	if in == nil {
		return nil
	}
	m := &proto.MaterialCategoryInfo{
		Id:          in.ID,
		Code:        in.Code,
		Description: in.Description,
		Remark:      in.Remark,
	}
	return m
}
