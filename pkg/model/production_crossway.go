package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 生产路口
type ProductionCrossway struct {
	ModelID
	Code                                string                                `json:"code" gorm:"index;size:100;comment:代号"`
	Description                         string                                `json:"description" gorm:"size:1000;comment:描述"`
	DefaultTurn                         int32                                 `json:"defaultTurn" gorm:"comment:默认走向"`
	Remark                              string                                `json:"remark" gorm:"size:1000;comment:备注"`
	ProductionLineID                    string                                `json:"productionLineID" gorm:"index;size:36;comment:生产产线ID"`
	ProductionLine                      *ProductionLine                       `json:"productionLine" gorm:"constraint:OnDelete:CASCADE"`
	ProductionCrosswayLeftTurnStations  []*ProductionCrosswayLeftTurnStation  `json:"productionCrosswayLeftTurnStations" gorm:"constraint:OnDelete:CASCADE;"`  //产线路口左转工站
	ProductionCrosswayRightTurnStations []*ProductionCrosswayRightTurnStation `json:"productionCrosswayRightTurnStations" gorm:"constraint:OnDelete:CASCADE;"` //产线路口右转工站
	ProductionCrosswayStraightStations  []*ProductionCrosswayStraightStation  `json:"productionCrosswayStraightStations" gorm:"constraint:OnDelete:CASCADE;"`  //产线路口交叉工站
}

// 产线路口左转工站
type ProductionCrosswayLeftTurnStation struct {
	ModelID
	ProductionCrosswayID string             `json:"productionCrosswayID" gorm:"index;size:36;comment:产线路口ID"`
	ProductionStationID  string             `json:"productionStationID" gorm:"size:36;comment:生产工站ID"`
	ProductionStation    *ProductionStation `gorm:"constraint:OnDelete:CASCADE"`
}

// 产线路口右转工站
type ProductionCrosswayRightTurnStation struct {
	ModelID
	ProductionCrosswayID string             `json:"productionCrosswayID" gorm:"index;size:36;comment:产线路口ID"`
	ProductionStationID  string             `json:"productionStationID" gorm:"size:36;comment:生产工站ID"`
	ProductionStation    *ProductionStation `gorm:"constraint:OnDelete:CASCADE"`
}

// 产线路口交叉工站
type ProductionCrosswayStraightStation struct {
	ModelID
	ProductionCrosswayID string             `json:"productionCrosswayID" gorm:"index;size:36;comment:产线路口ID"`
	ProductionStationID  string             `json:"productionStationID" gorm:"size:36;comment:生产工站ID"`
	ProductionStation    *ProductionStation `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToProductionCrossways(in []*proto.ProductionCrosswayInfo) []*ProductionCrossway {
	var result []*ProductionCrossway
	for _, c := range in {
		result = append(result, PBToProductionCrossway(c))
	}
	return result
}

func PBToProductionCrossway(in *proto.ProductionCrosswayInfo) *ProductionCrossway {
	if in == nil {
		return nil
	}
	return &ProductionCrossway{
		ModelID:                             ModelID{ID: in.Id},
		Code:                                in.Code,
		Description:                         in.Description,
		DefaultTurn:                         in.DefaultTurn,
		Remark:                              in.Remark,
		ProductionLineID:                    in.ProductionLineID,
		ProductionCrosswayLeftTurnStations:  PBToProductionCrosswayLeftTurnStations(in.ProductionCrosswayLeftTurnStations),
		ProductionCrosswayRightTurnStations: PBToProductionCrosswayRightTurnStations(in.ProductionCrosswayRightTurnStations),
		ProductionCrosswayStraightStations:  PBToProductionCrosswayStraightStations(in.ProductionCrosswayStraightStations),
	}
}

func ProductionCrosswaysToPB(in []*ProductionCrossway) []*proto.ProductionCrosswayInfo {
	var list []*proto.ProductionCrosswayInfo
	for _, f := range in {
		list = append(list, ProductionCrosswayToPB(f))
	}
	return list
}

func ProductionCrosswayToPB(in *ProductionCrossway) *proto.ProductionCrosswayInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductionCrosswayInfo{
		Id:                                  in.ID,
		Code:                                in.Code,
		Description:                         in.Description,
		DefaultTurn:                         in.DefaultTurn,
		Remark:                              in.Remark,
		ProductionLineID:                    in.ProductionLineID,
		ProductionCrosswayLeftTurnStations:  ProductionCrosswayLeftTurnStationsToPB(in.ProductionCrosswayLeftTurnStations),
		ProductionCrosswayRightTurnStations: ProductionCrosswayRightTurnStationsToPB(in.ProductionCrosswayRightTurnStations),
		ProductionCrosswayStraightStations:  ProductionCrosswayStraightStationsToPB(in.ProductionCrosswayStraightStations),
	}
	return m
}

func PBToProductionCrosswayLeftTurnStations(in []*proto.ProductionCrosswayStationInfo) []*ProductionCrosswayLeftTurnStation {
	var result []*ProductionCrosswayLeftTurnStation
	for _, c := range in {
		result = append(result, PBToProductionCrosswayLeftTurnStation(c))
	}
	return result
}

func PBToProductionCrosswayLeftTurnStation(in *proto.ProductionCrosswayStationInfo) *ProductionCrosswayLeftTurnStation {
	if in == nil {
		return nil
	}
	return &ProductionCrosswayLeftTurnStation{
		ModelID:              ModelID{ID: in.Id},
		ProductionCrosswayID: in.ProductionCrosswayID,
		ProductionStationID:  in.ProductionStationID,
	}
}

func ProductionCrosswayLeftTurnStationsToPB(in []*ProductionCrosswayLeftTurnStation) []*proto.ProductionCrosswayStationInfo {
	var list []*proto.ProductionCrosswayStationInfo
	for _, f := range in {
		list = append(list, ProductionCrosswayLeftTurnStationToPB(f))
	}
	return list
}

func ProductionCrosswayLeftTurnStationToPB(in *ProductionCrosswayLeftTurnStation) *proto.ProductionCrosswayStationInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductionCrosswayStationInfo{
		Id:                   in.ID,
		ProductionCrosswayID: in.ProductionCrosswayID,
		ProductionStationID:  in.ProductionStationID,
	}
	return m
}

func PBToProductionCrosswayRightTurnStations(in []*proto.ProductionCrosswayStationInfo) []*ProductionCrosswayRightTurnStation {
	var result []*ProductionCrosswayRightTurnStation
	for _, c := range in {
		result = append(result, PBToProductionCrosswayRightTurnStation(c))
	}
	return result
}

func PBToProductionCrosswayRightTurnStation(in *proto.ProductionCrosswayStationInfo) *ProductionCrosswayRightTurnStation {
	if in == nil {
		return nil
	}
	return &ProductionCrosswayRightTurnStation{
		ModelID:              ModelID{ID: in.Id},
		ProductionCrosswayID: in.ProductionCrosswayID,
		ProductionStationID:  in.ProductionStationID,
	}
}

func ProductionCrosswayRightTurnStationsToPB(in []*ProductionCrosswayRightTurnStation) []*proto.ProductionCrosswayStationInfo {
	var list []*proto.ProductionCrosswayStationInfo
	for _, f := range in {
		list = append(list, ProductionCrosswayRightTurnStationToPB(f))
	}
	return list
}

func ProductionCrosswayRightTurnStationToPB(in *ProductionCrosswayRightTurnStation) *proto.ProductionCrosswayStationInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductionCrosswayStationInfo{
		Id:                   in.ID,
		ProductionCrosswayID: in.ProductionCrosswayID,
		ProductionStationID:  in.ProductionStationID,
	}
	return m
}

func PBToProductionCrosswayStraightStations(in []*proto.ProductionCrosswayStationInfo) []*ProductionCrosswayStraightStation {
	var result []*ProductionCrosswayStraightStation
	for _, c := range in {
		result = append(result, PBToProductionCrosswayStraightStation(c))
	}
	return result
}

func PBToProductionCrosswayStraightStation(in *proto.ProductionCrosswayStationInfo) *ProductionCrosswayStraightStation {
	if in == nil {
		return nil
	}
	return &ProductionCrosswayStraightStation{
		ModelID:              ModelID{ID: in.Id},
		ProductionCrosswayID: in.ProductionCrosswayID,
		ProductionStationID:  in.ProductionStationID,
	}
}

func ProductionCrosswayStraightStationsToPB(in []*ProductionCrosswayStraightStation) []*proto.ProductionCrosswayStationInfo {
	var list []*proto.ProductionCrosswayStationInfo
	for _, f := range in {
		list = append(list, ProductionCrosswayStraightStationToPB(f))
	}
	return list
}

func ProductionCrosswayStraightStationToPB(in *ProductionCrosswayStraightStation) *proto.ProductionCrosswayStationInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductionCrosswayStationInfo{
		Id:                   in.ID,
		ProductionCrosswayID: in.ProductionCrosswayID,
		ProductionStationID:  in.ProductionStationID,
	}
	return m
}
