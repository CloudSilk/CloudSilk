package model

import "github.com/CloudSilk/CloudSilk/pkg/proto"

// 物料容器类型
type MaterialContainerType struct {
	ModelID
	Code        string  `grom:"size:50;comment:编号"`
	Description string  `grom:"size:100;comment:描述"`
	Length      int64   `grom:"comment:长(mm)"`
	Width       int64   `grom:"comment:宽(mm)"`
	Height      int64   `grom:"comment:高(mm)"`
	WeightLimit float32 `grom:"comment:限重(kg)"`
	HeightLimit int64   `grom:"comment:限高(mm)"`
	Remark      string  `grom:"size:500;comment:备注"`
}

func PBToMaterialContainerTypes(in []*proto.MaterialContainerTypeInfo) []*MaterialContainerType {
	var result []*MaterialContainerType
	for _, c := range in {
		result = append(result, PBToMaterialContainerType(c))
	}
	return result
}

func PBToMaterialContainerType(in *proto.MaterialContainerTypeInfo) *MaterialContainerType {
	if in == nil {
		return nil
	}
	return &MaterialContainerType{
		ModelID:     ModelID{ID: in.Id},
		Code:        in.Code,
		Description: in.Description,
		Length:      in.Length,
		Width:       in.Width,
		Height:      in.Height,
		WeightLimit: in.WeightLimit,
		HeightLimit: in.HeightLimit,
		Remark:      in.Remark,
	}
}

func MaterialContainerTypesToPB(in []*MaterialContainerType) []*proto.MaterialContainerTypeInfo {
	var list []*proto.MaterialContainerTypeInfo
	for _, f := range in {
		list = append(list, MaterialContainerTypeToPB(f))
	}
	return list
}

func MaterialContainerTypeToPB(in *MaterialContainerType) *proto.MaterialContainerTypeInfo {
	if in == nil {
		return nil
	}
	m := &proto.MaterialContainerTypeInfo{
		Id:          in.ID,
		Code:        in.Code,
		Description: in.Description,
		Length:      in.Length,
		Width:       in.Width,
		Height:      in.Height,
		WeightLimit: in.WeightLimit,
		HeightLimit: in.HeightLimit,
		Remark:      in.Remark,
	}
	return m
}
