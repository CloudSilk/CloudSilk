package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 物料信息
type MaterialInfo struct {
	ModelID
	MaterialCategoryID  string            `json:"materialCategoryID" gorm:"size:36;comment:物料类别ID;"`
	MaterialCategory    *MaterialCategory `json:"materialCategory"` //物料类别
	MaterialNo          string            `json:"materialNo" gorm:"size:50;comment:物料号"`
	MaterialDescription string            `json:"materialDescription" gorm:"size:500;comment:物料描述"`
	Unit                string            `json:"unit" gorm:"size:10;comment:单位"`
	Identifier          string            `json:"identifier" gorm:"size:50;comment:识别码"`
	StartIndex          int32             `json:"startIndex" gorm:"comment:索引"`
	EnableControl       bool              `json:"enableControl" gorm:"comment:是否管控"`
	ControlType         int32             `json:"controlType" gorm:"comment:管控类型"`
	Warehouse           string            `json:"warehouse" gorm:"size:50;comment:物料仓库"`
	Remark              string            `json:"remark" gorm:"size:500;comment:备注"`
}

func PBToMaterialInfos(in []*proto.MaterialInfoInfo) []*MaterialInfo {
	var result []*MaterialInfo
	for _, c := range in {
		result = append(result, PBToMaterialInfo(c))
	}
	return result
}

func PBToMaterialInfo(in *proto.MaterialInfoInfo) *MaterialInfo {
	if in == nil {
		return nil
	}
	return &MaterialInfo{
		ModelID:             ModelID{ID: in.Id},
		MaterialCategoryID:  in.MaterialCategoryID,
		MaterialNo:          in.MaterialNo,
		MaterialDescription: in.MaterialDescription,
		Unit:                in.Unit,
		Identifier:          in.Identifier,
		StartIndex:          in.StartIndex,
		EnableControl:       in.EnableControl,
		ControlType:         in.ControlType,
		Warehouse:           in.Warehouse,
		Remark:              in.Remark,
	}
}

func MaterialInfosToPB(in []*MaterialInfo) []*proto.MaterialInfoInfo {
	var list []*proto.MaterialInfoInfo
	for _, f := range in {
		list = append(list, MaterialInfoToPB(f))
	}
	return list
}

func MaterialInfoToPB(in *MaterialInfo) *proto.MaterialInfoInfo {
	if in == nil {
		return nil
	}
	m := &proto.MaterialInfoInfo{
		Id:                  in.ID,
		MaterialCategoryID:  in.MaterialCategoryID,
		MaterialCategory:    MaterialCategoryToPB(in.MaterialCategory),
		MaterialNo:          in.MaterialNo,
		MaterialDescription: in.MaterialDescription,
		Unit:                in.Unit,
		Identifier:          in.Identifier,
		StartIndex:          in.StartIndex,
		EnableControl:       in.EnableControl,
		ControlType:         in.ControlType,
		Warehouse:           in.Warehouse,
		Remark:              in.Remark,
	}
	return m
}
