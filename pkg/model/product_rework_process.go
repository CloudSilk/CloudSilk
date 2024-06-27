package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 产品返工工序
type ProductReworkProcess struct {
	ModelID
	Code                string                                  `gorm:"size:50;comment:代号"`
	Description         string                                  `gorm:"size:500;comment:描述"`
	Enable              bool                                    `gorm:"comment:是否启用"`
	EnableReport        bool                                    `gorm:"comment:是否报工"`
	Remark              string                                  `gorm:"size:500;comment:备注"`
	ProductionLineID    string                                  `gorm:"size:36;comment:生产产线ID"`
	ProductionLine      *ProductionLine                         `gorm:"constraint:OnDelete:CASCADE"`
	ProductionStations  []*ProductReworkProcessAvailableStation `gorm:"constraint:OnDelete:CASCADE"` //可用生产工站
	ProductionProcesses []*ProductReworkProcessAvailableProcess `gorm:"constraint:OnDelete:CASCADE"` //支持生产工序
}

// 产品返工工序可用生产工站
type ProductReworkProcessAvailableStation struct {
	ModelID
	ProductReworkProcessID string             `gorm:"index;size:36;comment:产品返工工序ID"`
	ProductionStationID    string             `gorm:"size:36;comment:生产工站ID"`
	ProductionStation      *ProductionStation `gorm:"constraint:OnDelete:CASCADE"`
}

// 产品返工工序支持生产工序
type ProductReworkProcessAvailableProcess struct {
	ModelID
	ProductReworkProcessID string             `gorm:"index;size:36;comment:产品返工工序ID"`
	ProductionProcessID    string             `gorm:"size:36;comment:生产工序ID"`
	ProductionProcess      *ProductionProcess `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToProductReworkProcesss(in []*proto.ProductReworkProcessInfo) []*ProductReworkProcess {
	var result []*ProductReworkProcess
	for _, c := range in {
		result = append(result, PBToProductReworkProcess(c))
	}
	return result
}

func PBToProductReworkProcess(in *proto.ProductReworkProcessInfo) *ProductReworkProcess {
	if in == nil {
		return nil
	}
	return &ProductReworkProcess{
		ModelID:             ModelID{ID: in.Id},
		Code:                in.Code,
		Description:         in.Description,
		Enable:              in.Enable,
		EnableReport:        in.EnableReport,
		Remark:              in.Remark,
		ProductionLineID:    in.ProductionLineID,
		ProductionStations:  PBToProductReworkProcessAvailableStations(in.ProductionStations),
		ProductionProcesses: PBToProductReworkProcessAvailableProcesss(in.ProductionProcesses),
	}
}

func ProductReworkProcesssToPB(in []*ProductReworkProcess) []*proto.ProductReworkProcessInfo {
	var list []*proto.ProductReworkProcessInfo
	for _, f := range in {
		list = append(list, ProductReworkProcessToPB(f))
	}
	return list
}

func ProductReworkProcessToPB(in *ProductReworkProcess) *proto.ProductReworkProcessInfo {
	if in == nil {
		return nil
	}

	var availableStationIDs, availableProcessIDs []string
	for _, availableStation := range in.ProductionStations {
		availableStationIDs = append(availableStationIDs, availableStation.ProductionStationID)
	}
	for _, availableProcess := range in.ProductionProcesses {
		availableProcessIDs = append(availableProcessIDs, availableProcess.ProductionProcessID)
	}

	m := &proto.ProductReworkProcessInfo{
		Id:                  in.ID,
		Code:                in.Code,
		Description:         in.Description,
		Enable:              in.Enable,
		EnableReport:        in.EnableReport,
		Remark:              in.Remark,
		ProductionLineID:    in.ProductionLineID,
		ProductionLine:      ProductionLineToPB(in.ProductionLine),
		AvailableStationIDs: availableStationIDs,
		ProductionStations:  ProductReworkProcessAvailableStationsToPB(in.ProductionStations),
		AvailableProcessIDs: availableProcessIDs,
		ProductionProcesses: ProductReworkProcessAvailableProcesssToPB(in.ProductionProcesses),
	}
	return m
}

func PBToProductReworkProcessAvailableStations(in []*proto.ProductReworkProcessAvailableStationInfo) []*ProductReworkProcessAvailableStation {
	var result []*ProductReworkProcessAvailableStation
	for _, c := range in {
		result = append(result, PBToProductReworkProcessAvailableStation(c))
	}
	return result
}

func PBToProductReworkProcessAvailableStation(in *proto.ProductReworkProcessAvailableStationInfo) *ProductReworkProcessAvailableStation {
	if in == nil {
		return nil
	}
	return &ProductReworkProcessAvailableStation{
		ModelID:                ModelID{ID: in.Id},
		ProductReworkProcessID: in.ProductReworkProcessID,
		ProductionStationID:    in.ProductionStationID,
	}
}

func ProductReworkProcessAvailableStationsToPB(in []*ProductReworkProcessAvailableStation) []*proto.ProductReworkProcessAvailableStationInfo {
	var list []*proto.ProductReworkProcessAvailableStationInfo
	for _, f := range in {
		list = append(list, ProductReworkProcessAvailableStationToPB(f))
	}
	return list
}

func ProductReworkProcessAvailableStationToPB(in *ProductReworkProcessAvailableStation) *proto.ProductReworkProcessAvailableStationInfo {
	if in == nil {
		return nil
	}

	m := &proto.ProductReworkProcessAvailableStationInfo{
		Id:                     in.ID,
		ProductReworkProcessID: in.ProductReworkProcessID,
		ProductionStationID:    in.ProductionStationID,
		ProductionStation:      ProductionStationToPB(in.ProductionStation),
	}
	return m
}

func PBToProductReworkProcessAvailableProcesss(in []*proto.ProductReworkProcessAvailableProcessInfo) []*ProductReworkProcessAvailableProcess {
	var result []*ProductReworkProcessAvailableProcess
	for _, c := range in {
		result = append(result, PBToProductReworkProcessAvailableProcess(c))
	}
	return result
}

func PBToProductReworkProcessAvailableProcess(in *proto.ProductReworkProcessAvailableProcessInfo) *ProductReworkProcessAvailableProcess {
	if in == nil {
		return nil
	}
	return &ProductReworkProcessAvailableProcess{
		ModelID:                ModelID{ID: in.Id},
		ProductReworkProcessID: in.ProductReworkProcessID,
		ProductionProcessID:    in.ProductionProcessID,
	}
}

func ProductReworkProcessAvailableProcesssToPB(in []*ProductReworkProcessAvailableProcess) []*proto.ProductReworkProcessAvailableProcessInfo {
	var list []*proto.ProductReworkProcessAvailableProcessInfo
	for _, f := range in {
		list = append(list, ProductReworkProcessAvailableProcessToPB(f))
	}
	return list
}

func ProductReworkProcessAvailableProcessToPB(in *ProductReworkProcessAvailableProcess) *proto.ProductReworkProcessAvailableProcessInfo {
	if in == nil {
		return nil
	}

	m := &proto.ProductReworkProcessAvailableProcessInfo{
		Id:                     in.ID,
		ProductReworkProcessID: in.ProductReworkProcessID,
		ProductionProcessID:    in.ProductionProcessID,
		ProductionProcess:      ProductionProcessToPB(in.ProductionProcess),
	}
	return m
}
