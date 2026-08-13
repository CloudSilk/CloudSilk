package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 物料容器
type MaterialContainer struct {
	ModelID
	Code                    string                 `gorm:"size:50;comment:编号"`
	Description             string                 `gorm:"size:100;comment:描述"`
	Identifier              string                 `gorm:"size:50;comment:识别码"`
	CurrentState            string                 `gorm:"size:50;comment:当前状态"`
	LastUpdateTime          time.Time              `gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark                  string                 `gorm:"size:500;comment:备注"`
	MaterialContainerTypeID *string                `gorm:"size:36;comment:容器类型ID"`
	MaterialContainerType   *MaterialContainerType `gorm:"constraint:OnDelete:SET NULL"`
	MaterialShelfBinID      *string                `gorm:"size:36;comment:当前库位ID"`
	MaterialShelfBin        *MaterialShelfBin      `gorm:"constraint:OnDelete:SET NULL"`
}

func PBToMaterialContainers(in []*proto.MaterialContainerInfo) []*MaterialContainer {
	var result []*MaterialContainer
	for _, c := range in {
		result = append(result, PBToMaterialContainer(c))
	}
	return result
}

func PBToMaterialContainer(in *proto.MaterialContainerInfo) *MaterialContainer {
	if in == nil {
		return nil
	}

	var materialContainerTypeID, materialShelfBinID *string
	if in.MaterialContainerTypeID != "" {
		materialContainerTypeID = &in.MaterialContainerTypeID
	}
	if in.MaterialShelfBinID != "" {
		materialShelfBinID = &in.MaterialShelfBinID
	}

	return &MaterialContainer{
		ModelID:                 ModelID{ID: in.Id},
		Code:                    in.Code,
		Description:             in.Description,
		Identifier:              in.Identifier,
		CurrentState:            in.CurrentState,
		Remark:                  in.Remark,
		MaterialContainerTypeID: materialContainerTypeID,
		MaterialShelfBinID:      materialShelfBinID,
	}
}

func MaterialContainersToPB(in []*MaterialContainer) []*proto.MaterialContainerInfo {
	var list []*proto.MaterialContainerInfo
	for _, f := range in {
		list = append(list, MaterialContainerToPB(f))
	}
	return list
}

func MaterialContainerToPB(in *MaterialContainer) *proto.MaterialContainerInfo {
	if in == nil {
		return nil
	}

	var materialContainerTypeID, materialShelfBinID string
	if in.MaterialContainerTypeID != nil {
		materialContainerTypeID = *in.MaterialContainerTypeID
	}
	if in.MaterialShelfBinID != nil {
		materialShelfBinID = *in.MaterialShelfBinID
	}

	m := &proto.MaterialContainerInfo{
		Id:                      in.ID,
		Code:                    in.Code,
		Description:             in.Description,
		Identifier:              in.Identifier,
		CurrentState:            in.CurrentState,
		LastUpdateTime:          utils.FormatTime(in.LastUpdateTime),
		Remark:                  in.Remark,
		MaterialContainerTypeID: materialContainerTypeID,
		MaterialContainerType:   MaterialContainerTypeToPB(in.MaterialContainerType),
		MaterialShelfBinID:      materialShelfBinID,
		MaterialShelfBin:        MaterialShelfBinToPB(in.MaterialShelfBin),
	}
	return m
}
