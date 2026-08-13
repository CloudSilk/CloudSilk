package model

import "github.com/CloudSilk/CloudSilk/pkg/proto"

// 物料货架
type MaterialShelf struct {
	ModelID
	Code            string         `gorm:"size:50;comment:代号"`
	Description     string         `gorm:"size:500;comment:描述"`
	Identifier      string         `gorm:"size:50;comment:识别码"`
	ShelfType       int32          `gorm:"comment:货架类型"`
	Enable          bool           `gorm:"comment:是否启用"`
	Remark          string         `gorm:"size:500;comment:备注"`
	MaterialStoreID string         `gorm:"size:36;comment:隶属仓库ID"`
	MaterialStore   *MaterialStore `gorm:"constraint:OnDelete:CASCADE"` //隶属仓库
}

func PBToMaterialShelfs(in []*proto.MaterialShelfInfo) []*MaterialShelf {
	var result []*MaterialShelf
	for _, c := range in {
		result = append(result, PBToMaterialShelf(c))
	}
	return result
}

func PBToMaterialShelf(in *proto.MaterialShelfInfo) *MaterialShelf {
	if in == nil {
		return nil
	}
	return &MaterialShelf{
		ModelID:         ModelID{ID: in.Id},
		Code:            in.Code,
		Description:     in.Description,
		Identifier:      in.Identifier,
		ShelfType:       in.ShelfType,
		Enable:          in.Enable,
		Remark:          in.Remark,
		MaterialStoreID: in.MaterialStoreID,
	}
}

func MaterialShelfsToPB(in []*MaterialShelf) []*proto.MaterialShelfInfo {
	var list []*proto.MaterialShelfInfo
	for _, f := range in {
		list = append(list, MaterialShelfToPB(f))
	}
	return list
}

func MaterialShelfToPB(in *MaterialShelf) *proto.MaterialShelfInfo {
	if in == nil {
		return nil
	}
	m := &proto.MaterialShelfInfo{
		Id:              in.ID,
		Code:            in.Code,
		Description:     in.Description,
		Identifier:      in.Identifier,
		ShelfType:       in.ShelfType,
		Enable:          in.Enable,
		Remark:          in.Remark,
		MaterialStoreID: in.MaterialStoreID,
		MaterialStore:   MaterialStoreToPB(in.MaterialStore),
	}
	return m
}
