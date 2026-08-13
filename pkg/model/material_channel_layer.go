package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 物料通道层
type MaterialChannelLayer struct {
	ModelID
	SortIndex             int32              `gorm:"comment:顺序"`
	Code                  string             `gorm:"size:50;comment:代号"`
	StatusRegisterAddress int32              `gorm:"comment:状态寄存器地址"`
	LightRegisterAddress  int32              `gorm:"comment:亮灯寄存器地址"`
	Remark                string             `gorm:"size:500;comment:备注"`
	ProductionStationID   string             `gorm:"size:36;comment:隶属工站ID"`
	ProductionStation     *ProductionStation `gorm:"constraint:OnDelete:CASCADE"` //隶属工站
	MaterialChannels      []*MaterialChannel `gorm:"constraint:OnDelete:CASCADE"` //物料通道
}

// 物料通道
type MaterialChannel struct {
	ModelID
	SortIndex              int32                 `gorm:"comment:顺序"`
	Code                   string                `gorm:"size:50;comment:代号"`
	Description            string                `gorm:"size:100;comment:描述"`
	Size                   string                `gorm:"size:50;comment:尺寸"`
	Spec                   float32               `gorm:"comment:规格"`
	EnableMonitor          bool                  `gorm:"comment:启用监控"`
	CurrentState           string                `gorm:"size:50;comment:当前状态"`
	LastUpdateTime         time.Time             `gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark                 string                `gorm:"size:500;comment:备注"`
	MaterialChannelLayerID string                `gorm:"size:36;comment:物料通道层ID"`
	MaterialChannelLayer   *MaterialChannelLayer `gorm:"constraint:OnDelete:CASCADE"` //物料通道层
	MaterialInfoID         string                `gorm:"size:36;comment:物料信息ID"`
	MaterialInfo           *MaterialInfo         `gorm:"constraint:OnDelete:CASCADE"` //物料通道层
}

func PBToMaterialChannelLayers(in []*proto.MaterialChannelLayerInfo) []*MaterialChannelLayer {
	var result []*MaterialChannelLayer
	for _, c := range in {
		result = append(result, PBToMaterialChannelLayer(c))
	}
	return result
}

func PBToMaterialChannelLayer(in *proto.MaterialChannelLayerInfo) *MaterialChannelLayer {
	if in == nil {
		return nil
	}
	return &MaterialChannelLayer{
		ModelID:               ModelID{ID: in.Id},
		SortIndex:             in.SortIndex,
		Code:                  in.Code,
		StatusRegisterAddress: in.StatusRegisterAddress,
		LightRegisterAddress:  in.LightRegisterAddress,
		Remark:                in.Remark,
		ProductionStationID:   in.ProductionStationID,
		MaterialChannels:      PBToMaterialChannels(in.MaterialChannels),
	}
}

func MaterialChannelLayersToPB(in []*MaterialChannelLayer) []*proto.MaterialChannelLayerInfo {
	var list []*proto.MaterialChannelLayerInfo
	for _, f := range in {
		list = append(list, MaterialChannelLayerToPB(f))
	}
	return list
}

func MaterialChannelLayerToPB(in *MaterialChannelLayer) *proto.MaterialChannelLayerInfo {
	if in == nil {
		return nil
	}
	m := &proto.MaterialChannelLayerInfo{
		Id:                    in.ID,
		SortIndex:             in.SortIndex,
		Code:                  in.Code,
		StatusRegisterAddress: in.StatusRegisterAddress,
		LightRegisterAddress:  in.LightRegisterAddress,
		Remark:                in.Remark,
		ProductionStationID:   in.ProductionStationID,
		ProductionStation:     ProductionStationToPB(in.ProductionStation),
		MaterialChannels:      MaterialChannelsToPB(in.MaterialChannels),
	}
	return m
}

func PBToMaterialChannels(in []*proto.MaterialChannelInfo) []*MaterialChannel {
	var result []*MaterialChannel
	for _, c := range in {
		result = append(result, PBToMaterialChannel(c))
	}
	return result
}

func PBToMaterialChannel(in *proto.MaterialChannelInfo) *MaterialChannel {
	if in == nil {
		return nil
	}
	return &MaterialChannel{
		ModelID:                ModelID{ID: in.Id},
		SortIndex:              in.SortIndex,
		Code:                   in.Code,
		Description:            in.Description,
		Size:                   in.Size,
		Spec:                   in.Spec,
		EnableMonitor:          in.EnableMonitor,
		CurrentState:           in.CurrentState,
		Remark:                 in.Remark,
		MaterialChannelLayerID: in.MaterialChannelLayerID,
		MaterialInfoID:         in.MaterialInfoID,
	}
}

func MaterialChannelsToPB(in []*MaterialChannel) []*proto.MaterialChannelInfo {
	var list []*proto.MaterialChannelInfo
	for _, f := range in {
		list = append(list, MaterialChannelToPB(f))
	}
	return list
}

func MaterialChannelToPB(in *MaterialChannel) *proto.MaterialChannelInfo {
	if in == nil {
		return nil
	}
	m := &proto.MaterialChannelInfo{
		Id:                     in.ID,
		SortIndex:              in.SortIndex,
		Code:                   in.Code,
		Description:            in.Description,
		Size:                   in.Size,
		Spec:                   in.Spec,
		EnableMonitor:          in.EnableMonitor,
		CurrentState:           in.CurrentState,
		LastUpdateTime:         utils.FormatTime(in.LastUpdateTime),
		Remark:                 in.Remark,
		MaterialChannelLayerID: in.MaterialChannelLayerID,
		MaterialChannelLayer:   MaterialChannelLayerToPB(in.MaterialChannelLayer),
		MaterialInfoID:         in.MaterialInfoID,
		MaterialInfo:           MaterialInfoToPB(in.MaterialInfo),
	}
	return m
}
