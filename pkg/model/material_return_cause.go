package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 物料退料原因
type MaterialReturnCause struct {
	ModelID
	Code                string                                  `gorm:"size:50;comment:代号"`          //代号
	Description         string                                  `gorm:"size:500;comment:描述"`         //描述
	Remark              string                                  `gorm:"size:500;comment:备注"`         //备注
	MaterialCategories  []*MaterialReturnCauseAvailableCategory `gorm:"constraint:OnDelete:CASCADE"` //物料类别
	MaterialReturnTypes []*MaterialReturnCauseAvailableType     `gorm:"constraint:OnDelete:CASCADE"` //归属类型
}

// 物料退料原因-物料类别
type MaterialReturnCauseAvailableCategory struct {
	ModelID
	MaterialReturnCauseID string            `gorm:"index;size:36;comment:物料退料原因ID"`
	MaterialCategoryID    string            `gorm:"size:36;comment:物料类别ID"`
	MaterialCategory      *MaterialCategory `gorm:"constraint:OnDelete:CASCADE"`
}

// 物料退料原因-归属类型
type MaterialReturnCauseAvailableType struct {
	ModelID
	MaterialReturnCauseID string              `gorm:"index;size:36;comment:物料退料原因ID"`
	MaterialReturnTypeID  string              `gorm:"size:36;comment:物料退料类型ID"`
	MaterialReturnType    *MaterialReturnType `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToMaterialReturnCauses(in []*proto.MaterialReturnCauseInfo) []*MaterialReturnCause {
	var result []*MaterialReturnCause
	for _, c := range in {
		result = append(result, PBToMaterialReturnCause(c))
	}
	return result
}

func PBToMaterialReturnCause(in *proto.MaterialReturnCauseInfo) *MaterialReturnCause {
	if in == nil {
		return nil
	}

	return &MaterialReturnCause{
		ModelID:             ModelID{ID: in.Id},
		Code:                in.Code,
		Description:         in.Description,
		Remark:              in.Remark,
		MaterialCategories:  PBToMaterialReturnCauseAvailableCategorys(in.MaterialCategories),
		MaterialReturnTypes: PBToMaterialReturnCauseAvailableTypes(in.MaterialReturnTypes),
	}
}

func MaterialReturnCausesToPB(in []*MaterialReturnCause) []*proto.MaterialReturnCauseInfo {
	var list []*proto.MaterialReturnCauseInfo
	for _, f := range in {
		list = append(list, MaterialReturnCauseToPB(f))
	}
	return list
}

func MaterialReturnCauseToPB(in *MaterialReturnCause) *proto.MaterialReturnCauseInfo {
	if in == nil {
		return nil
	}

	m := &proto.MaterialReturnCauseInfo{
		Id:                  in.ID,
		Code:                in.Code,
		Description:         in.Description,
		Remark:              in.Remark,
		MaterialCategories:  MaterialReturnCauseAvailableCategorysToPB(in.MaterialCategories),
		MaterialReturnTypes: MaterialReturnCauseAvailableTypesToPB(in.MaterialReturnTypes),
	}
	return m
}

func PBToMaterialReturnCauseAvailableCategorys(in []*proto.MaterialReturnCauseAvailableCategoryInfo) []*MaterialReturnCauseAvailableCategory {
	var result []*MaterialReturnCauseAvailableCategory
	for _, c := range in {
		result = append(result, PBToMaterialReturnCauseAvailableCategory(c))
	}
	return result
}

func PBToMaterialReturnCauseAvailableCategory(in *proto.MaterialReturnCauseAvailableCategoryInfo) *MaterialReturnCauseAvailableCategory {
	if in == nil {
		return nil
	}

	return &MaterialReturnCauseAvailableCategory{
		ModelID:               ModelID{ID: in.Id},
		MaterialReturnCauseID: in.MaterialReturnCauseID,
		MaterialCategoryID:    in.MaterialCategoryID,
	}
}

func MaterialReturnCauseAvailableCategorysToPB(in []*MaterialReturnCauseAvailableCategory) []*proto.MaterialReturnCauseAvailableCategoryInfo {
	var list []*proto.MaterialReturnCauseAvailableCategoryInfo
	for _, f := range in {
		list = append(list, MaterialReturnCauseAvailableCategoryToPB(f))
	}
	return list
}

func MaterialReturnCauseAvailableCategoryToPB(in *MaterialReturnCauseAvailableCategory) *proto.MaterialReturnCauseAvailableCategoryInfo {
	if in == nil {
		return nil
	}

	m := &proto.MaterialReturnCauseAvailableCategoryInfo{
		Id:                    in.ID,
		MaterialReturnCauseID: in.MaterialReturnCauseID,
		MaterialCategoryID:    in.MaterialCategoryID,
		MaterialCategory:      MaterialCategoryToPB(in.MaterialCategory),
	}
	return m
}

func PBToMaterialReturnCauseAvailableTypes(in []*proto.MaterialReturnCauseAvailableTypeInfo) []*MaterialReturnCauseAvailableType {
	var result []*MaterialReturnCauseAvailableType
	for _, c := range in {
		result = append(result, PBToMaterialReturnCauseAvailableType(c))
	}
	return result
}

func PBToMaterialReturnCauseAvailableType(in *proto.MaterialReturnCauseAvailableTypeInfo) *MaterialReturnCauseAvailableType {
	if in == nil {
		return nil
	}

	return &MaterialReturnCauseAvailableType{
		ModelID:               ModelID{ID: in.Id},
		MaterialReturnCauseID: in.MaterialReturnCauseID,
		MaterialReturnTypeID:  in.MaterialReturnTypeID,
	}
}

func MaterialReturnCauseAvailableTypesToPB(in []*MaterialReturnCauseAvailableType) []*proto.MaterialReturnCauseAvailableTypeInfo {
	var list []*proto.MaterialReturnCauseAvailableTypeInfo
	for _, f := range in {
		list = append(list, MaterialReturnCauseAvailableTypeToPB(f))
	}
	return list
}

func MaterialReturnCauseAvailableTypeToPB(in *MaterialReturnCauseAvailableType) *proto.MaterialReturnCauseAvailableTypeInfo {
	if in == nil {
		return nil
	}

	m := &proto.MaterialReturnCauseAvailableTypeInfo{
		Id:                    in.ID,
		MaterialReturnCauseID: in.MaterialReturnCauseID,
		MaterialReturnTypeID:  in.MaterialReturnTypeID,
		MaterialReturnType:    MaterialReturnTypeToPB(in.MaterialReturnType),
	}
	return m
}
