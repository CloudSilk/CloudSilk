package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 物料货架库位
type MaterialShelfBin struct {
	ModelID
	Code                    string                 `gorm:"size:50;comment:编号"`
	RowIndex                int32                  `gorm:"comment:行"`
	ColumnIndex             int32                  `gorm:"comment:列"`
	Identifier              string                 `gorm:"size:50;comment:识别码"`
	MaterialContainerTypeID *string                `gorm:"size:36;comment:兼容容器类型ID"`
	MaterialContainerType   *MaterialContainerType `gorm:"constraint:OnDelete:SET NULL"` //兼容容器类型
	EnableDispatch          bool                   `gorm:"comment:缺料调度"`
	EnableMonitor           bool                   `gorm:"comment:状态监控"`
	CurrentState            string                 `gorm:"size:50;comment:当前状态"`
	LastUpdateTime          time.Time              `gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark                  string                 `gorm:"size:500;comment:备注"`
	MaterialShelfID         string                 `gorm:"size:36;comment:隶属料架ID"`
	MaterialShelf           *MaterialShelf         `gorm:"constraint:OnDelete:CASCADE"` //隶属料架
	MaterialInfoID          *string                `gorm:"size:36;comment:当前物料ID"`
	MaterialInfo            *MaterialInfo          `gorm:"constraint:OnDelete:SET NULL"` //当前物料
}

func PBToMaterialShelfBins(in []*proto.MaterialShelfBinInfo) []*MaterialShelfBin {
	var result []*MaterialShelfBin
	for _, c := range in {
		result = append(result, PBToMaterialShelfBin(c))
	}
	return result
}

func PBToMaterialShelfBin(in *proto.MaterialShelfBinInfo) *MaterialShelfBin {
	if in == nil {
		return nil
	}

	var materialContainerTypeID, materialInfoID *string
	if in.MaterialContainerTypeID != "" {
		materialContainerTypeID = &in.MaterialContainerTypeID
	}
	if in.MaterialInfoID != "" {
		materialInfoID = &in.MaterialInfoID
	}

	return &MaterialShelfBin{
		ModelID:                 ModelID{ID: in.Id},
		Code:                    in.Code,
		RowIndex:                in.RowIndex,
		ColumnIndex:             in.ColumnIndex,
		Identifier:              in.Identifier,
		MaterialContainerTypeID: materialContainerTypeID,
		EnableDispatch:          in.EnableDispatch,
		EnableMonitor:           in.EnableMonitor,
		CurrentState:            in.CurrentState,
		Remark:                  in.Remark,
		MaterialShelfID:         in.MaterialShelfID,
		MaterialInfoID:          materialInfoID,
	}
}

func MaterialShelfBinsToPB(in []*MaterialShelfBin) []*proto.MaterialShelfBinInfo {
	var list []*proto.MaterialShelfBinInfo
	for _, f := range in {
		list = append(list, MaterialShelfBinToPB(f))
	}
	return list
}

func MaterialShelfBinToPB(in *MaterialShelfBin) *proto.MaterialShelfBinInfo {
	if in == nil {
		return nil
	}

	var materialContainerTypeID, materialInfoID string
	if in.MaterialContainerTypeID != nil {
		materialContainerTypeID = *in.MaterialContainerTypeID
	}
	if in.MaterialInfoID != nil {
		materialInfoID = *in.MaterialInfoID
	}

	m := &proto.MaterialShelfBinInfo{
		Id:                      in.ID,
		Code:                    in.Code,
		RowIndex:                in.RowIndex,
		ColumnIndex:             in.ColumnIndex,
		Identifier:              in.Identifier,
		MaterialContainerTypeID: materialContainerTypeID,
		MaterialContainerType:   MaterialContainerTypeToPB(in.MaterialContainerType),
		EnableDispatch:          in.EnableDispatch,
		EnableMonitor:           in.EnableMonitor,
		CurrentState:            in.CurrentState,
		LastUpdateTime:          utils.FormatTime(in.LastUpdateTime),
		Remark:                  in.Remark,
		MaterialShelfID:         in.MaterialShelfID,
		MaterialShelf:           MaterialShelfToPB(in.MaterialShelf),
		MaterialInfoID:          materialInfoID,
		MaterialInfo:            MaterialInfoToPB(in.MaterialInfo),
	}
	return m
}
