package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 物料供应商
type MaterialSupplier struct {
	ModelID
	Code               string               `json:"code" gorm:"size:50;comment:代号"`
	Description        string               `json:"description" gorm:"size:500;comment:描述"`
	Remark             string               `json:"remark" gorm:"size:500;comment:备注"`
	AvailableMaterials []*AvailableMaterial `json:"materialInfoes" gorm:"constraint:OnDelete:CASCADE;"` //持有物料信息
}

// 持有物料信息
type AvailableMaterial struct {
	ModelID
	MaterialSupplierID string        `json:"materialSupplierID" gorm:"index;size:36;comment:物料供应商ID"`
	MaterialInfoID     string        `json:"materialInfoID" gorm:"size:36;comment:物料信息ID"`
	MaterialInfo       *MaterialInfo `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToMaterialSuppliers(in []*proto.MaterialSupplierInfo) []*MaterialSupplier {
	var result []*MaterialSupplier
	for _, c := range in {
		result = append(result, PBToMaterialSupplier(c))
	}
	return result
}

func PBToMaterialSupplier(in *proto.MaterialSupplierInfo) *MaterialSupplier {
	if in == nil {
		return nil
	}
	return &MaterialSupplier{
		ModelID:            ModelID{ID: in.Id},
		Code:               in.Code,
		Description:        in.Description,
		Remark:             in.Remark,
		AvailableMaterials: PBToAvailableMaterials(in.AvailableMaterials),
	}
}

func MaterialSuppliersToPB(in []*MaterialSupplier) []*proto.MaterialSupplierInfo {
	var list []*proto.MaterialSupplierInfo
	for _, f := range in {
		list = append(list, MaterialSupplierToPB(f))
	}
	return list
}

func MaterialSupplierToPB(in *MaterialSupplier) *proto.MaterialSupplierInfo {
	if in == nil {
		return nil
	}
	var AvailableMaterials []*proto.AvailableMaterialInfo
	if in.AvailableMaterials != nil {
		AvailableMaterials = AvailableMaterialsToPB(in.AvailableMaterials)
	}
	m := &proto.MaterialSupplierInfo{
		Id:                 in.ID,
		Code:               in.Code,
		Description:        in.Description,
		Remark:             in.Remark,
		AvailableMaterials: AvailableMaterials,
	}
	return m
}

func PBToAvailableMaterials(in []*proto.AvailableMaterialInfo) []*AvailableMaterial {
	var result []*AvailableMaterial
	for _, c := range in {
		result = append(result, PBToAvailableMaterial(c))
	}
	return result
}

func PBToAvailableMaterial(in *proto.AvailableMaterialInfo) *AvailableMaterial {
	if in == nil {
		return nil
	}
	return &AvailableMaterial{
		ModelID:            ModelID{ID: in.Id},
		MaterialSupplierID: in.MaterialSupplierID,
		MaterialInfoID:     in.MaterialInfoID,
	}
}

func AvailableMaterialsToPB(in []*AvailableMaterial) []*proto.AvailableMaterialInfo {
	var list []*proto.AvailableMaterialInfo
	for _, f := range in {
		list = append(list, AvailableMaterialToPB(f))
	}
	return list
}

func AvailableMaterialToPB(in *AvailableMaterial) *proto.AvailableMaterialInfo {
	if in == nil {
		return nil
	}
	m := &proto.AvailableMaterialInfo{
		Id:                 in.ID,
		MaterialSupplierID: in.MaterialSupplierID,
		MaterialInfoID:     in.MaterialInfoID,
	}
	return m
}
