package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 生产工序
type ProductionProcess struct {
	ModelID
	SortIndex                          int32                                `json:"sortIndex" gorm:"comment:顺序"`
	Code                               string                               `json:"code" gorm:"index;size:100;comment:代号"`
	Description                        string                               `json:"description" gorm:"size:1000;comment:描述"`
	Identifier                         string                               `json:"identifier" gorm:"size:100;comment:识别码"`
	Enable                             bool                                 `json:"enable" gorm:"comment:是否启用"`
	InitialValue                       bool                                 `json:"initialValue" gorm:"comment:默认匹配"`
	EnableReport                       bool                                 `json:"enableReport" gorm:"comment:是否报工"`
	EnableControl                      bool                                 `json:"enableControl" gorm:"comment:是否管控"`
	ProcessType                        int32                                `json:"processType" gorm:"comment:工序类型"`
	VehicleType                        int32                                `json:"vehicleType" gorm:"comment:载具类型"`
	ProductState                       string                               `json:"productState" gorm:"size:100;comment:产品状态"`
	Remark                             string                               `json:"remark" gorm:"size:1000;comment:备注"`
	ProductionLineID                   string                               `json:"productionLineID" gorm:"size:36;comment:生产产线ID"`
	ProductionLine                     *ProductionLine                      `gorm:"constraint:OnDelete:CASCADE"`                         //生产产线
	ProductionProcessAvailableStations []*ProductionProcessAvailableStation `gorm:"constraint:OnDelete:CASCADE"`                         //支持工站
	AttributeExpressions               []*AttributeExpression               `gorm:"polymorphic:Rule;polymorphicValue:ProductionProcess"` //特性表达式
	ProductionProcessSteps             []*AvailableProcess                  `gorm:"constraint:OnDelete:CASCADE"`
}

type ProductionProcessAvailableStation struct {
	ModelID
	ProductionProcessID string             `gorm:"index;size:36;comment:生产工序ID"`
	ProductionProcess   *ProductionProcess `gorm:"constraint:OnDelete:CASCADE"`
	ProductionStationID string             `gorm:"index;size:36;comment:生产工站ID"`
	ProductionStation   *ProductionStation `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToProductionProcesss(in []*proto.ProductionProcessInfo) []*ProductionProcess {
	var result []*ProductionProcess
	for _, c := range in {
		result = append(result, PBToProductionProcess(c))
	}
	return result
}

func PBToProductionProcess(in *proto.ProductionProcessInfo) *ProductionProcess {
	if in == nil {
		return nil
	}
	return &ProductionProcess{
		ModelID:                            ModelID{ID: in.Id},
		SortIndex:                          in.SortIndex,
		Code:                               in.Code,
		Description:                        in.Description,
		Identifier:                         in.Identifier,
		Enable:                             in.Enable,
		InitialValue:                       in.InitialValue,
		EnableReport:                       in.EnableReport,
		EnableControl:                      in.EnableControl,
		ProcessType:                        in.ProcessType,
		VehicleType:                        in.VehicleType,
		ProductState:                       in.ProductState,
		Remark:                             in.Remark,
		ProductionLineID:                   in.ProductionLineID,
		ProductionProcessAvailableStations: PBToProductionProcessAvailableStations(in.ProductionProcessAvailableStations),
		AttributeExpressions:               PBToAttributeExpressions(in.AttributeExpressions),
	}
}

func ProductionProcesssToPB(in []*ProductionProcess) []*proto.ProductionProcessInfo {
	var list []*proto.ProductionProcessInfo
	for _, f := range in {
		list = append(list, ProductionProcessToPB(f))
	}
	return list
}

func ProductionProcessToPB(in *ProductionProcess) *proto.ProductionProcessInfo {
	if in == nil {
		return nil
	}

	var availableStationIDs []string
	for _, ProductionProcessAvailableStation := range in.ProductionProcessAvailableStations {
		availableStationIDs = append(availableStationIDs, ProductionProcessAvailableStation.ProductionStationID)
	}

	m := &proto.ProductionProcessInfo{
		Id:                                 in.ID,
		SortIndex:                          in.SortIndex,
		Code:                               in.Code,
		Description:                        in.Description,
		Identifier:                         in.Identifier,
		Enable:                             in.Enable,
		InitialValue:                       in.InitialValue,
		EnableReport:                       in.EnableReport,
		EnableControl:                      in.EnableControl,
		ProcessType:                        in.ProcessType,
		VehicleType:                        in.VehicleType,
		ProductState:                       in.ProductState,
		Remark:                             in.Remark,
		ProductionLineID:                   in.ProductionLineID,
		ProductionLine:                     ProductionLineToPB(in.ProductionLine),
		AvailableStationIDs:                availableStationIDs,
		ProductionProcessAvailableStations: ProductionProcessAvailableStationsToPB(in.ProductionProcessAvailableStations),
		AttributeExpressions:               AttributeExpressionsToPB(in.AttributeExpressions),
		ProductionProcessSteps:             AvailableProcesssToPB(in.ProductionProcessSteps),
	}
	return m
}

func PBToProductionProcessAvailableStations(in []*proto.ProductionProcessAvailableStationInfo) []*ProductionProcessAvailableStation {
	var result []*ProductionProcessAvailableStation
	for _, c := range in {
		result = append(result, PBToProductionProcessAvailableStation(c))
	}
	return result
}

func PBToProductionProcessAvailableStation(in *proto.ProductionProcessAvailableStationInfo) *ProductionProcessAvailableStation {
	if in == nil {
		return nil
	}
	return &ProductionProcessAvailableStation{
		ModelID:             ModelID{ID: in.Id},
		ProductionProcessID: in.ProductionProcessID,
		ProductionStationID: in.ProductionStationID,
	}
}

func ProductionProcessAvailableStationsToPB(in []*ProductionProcessAvailableStation) []*proto.ProductionProcessAvailableStationInfo {
	var list []*proto.ProductionProcessAvailableStationInfo
	for _, f := range in {
		list = append(list, ProductionProcessAvailableStationToPB(f))
	}
	return list
}

func ProductionProcessAvailableStationToPB(in *ProductionProcessAvailableStation) *proto.ProductionProcessAvailableStationInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductionProcessAvailableStationInfo{
		Id:                  in.ID,
		ProductionProcessID: in.ProductionProcessID,
		ProductionStationID: in.ProductionStationID,
		ProductionStation:   ProductionStationToPB(in.ProductionStation),
	}
	return m
}
