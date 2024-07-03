package model

import "github.com/CloudSilk/CloudSilk/pkg/proto"

// 物料仓库
type MaterialStore struct {
	ModelID
	Code            string                        `gorm:"size:50;comment:代号"`
	Description     string                        `gorm:"size:500;comment:描述"`
	Remark          string                        `gorm:"size:500;comment:备注"`
	ProductionLines []*MaterialStoreAvailableLine `gorm:"constraint:OnDelete:CASCADE"` //支援产线
}

// 物料仓库支援产线
type MaterialStoreAvailableLine struct {
	ModelID
	MaterialStoreID  string          `gorm:"index;size:36;comment:物料仓库ID"`
	ProductionLineID string          `gorm:"size:36;comment:生产产线ID"`
	ProductionLine   *ProductionLine `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToMaterialStores(in []*proto.MaterialStoreInfo) []*MaterialStore {
	var result []*MaterialStore
	for _, c := range in {
		result = append(result, PBToMaterialStore(c))
	}
	return result
}

func PBToMaterialStore(in *proto.MaterialStoreInfo) *MaterialStore {
	if in == nil {
		return nil
	}
	return &MaterialStore{
		ModelID:         ModelID{ID: in.Id},
		Code:            in.Code,
		Description:     in.Description,
		Remark:          in.Remark,
		ProductionLines: PBToMaterialStoreAvailableLines(in.ProductionLines),
	}
}

func MaterialStoresToPB(in []*MaterialStore) []*proto.MaterialStoreInfo {
	var list []*proto.MaterialStoreInfo
	for _, f := range in {
		list = append(list, MaterialStoreToPB(f))
	}
	return list
}

func MaterialStoreToPB(in *MaterialStore) *proto.MaterialStoreInfo {
	if in == nil {
		return nil
	}
	m := &proto.MaterialStoreInfo{
		Id:              in.ID,
		Code:            in.Code,
		Description:     in.Description,
		Remark:          in.Remark,
		ProductionLines: MaterialStoreAvailableLinesToPB(in.ProductionLines),
	}
	return m
}

func PBToMaterialStoreAvailableLines(in []*proto.MaterialStoreAvailableLineInfo) []*MaterialStoreAvailableLine {
	var result []*MaterialStoreAvailableLine
	for _, c := range in {
		result = append(result, PBToMaterialStoreAvailableLine(c))
	}
	return result
}

func PBToMaterialStoreAvailableLine(in *proto.MaterialStoreAvailableLineInfo) *MaterialStoreAvailableLine {
	if in == nil {
		return nil
	}
	return &MaterialStoreAvailableLine{
		ModelID:          ModelID{ID: in.Id},
		MaterialStoreID:  in.MaterialStoreID,
		ProductionLineID: in.ProductionLineID,
	}
}

func MaterialStoreAvailableLinesToPB(in []*MaterialStoreAvailableLine) []*proto.MaterialStoreAvailableLineInfo {
	var list []*proto.MaterialStoreAvailableLineInfo
	for _, f := range in {
		list = append(list, MaterialStoreAvailableLineToPB(f))
	}
	return list
}

func MaterialStoreAvailableLineToPB(in *MaterialStoreAvailableLine) *proto.MaterialStoreAvailableLineInfo {
	if in == nil {
		return nil
	}
	m := &proto.MaterialStoreAvailableLineInfo{
		Id:               in.ID,
		MaterialStoreID:  in.MaterialStoreID,
		ProductionLineID: in.ProductionLineID,
		ProductionLine:   ProductionLineToPB(in.ProductionLine),
	}
	return m
}
