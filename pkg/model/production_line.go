package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 生产产线
type ProductionLine struct {
	ModelID
	ProductionFactoryID                string                               `json:"productionFactoryID" gorm:"index;size:36;comment:生产工厂ID"`
	ProductionFactory                  *ProductionFactory                   `json:"productionFactory"`
	Code                               string                               `json:"code" gorm:"index;size:100;comment:代号"`
	Description                        string                               `json:"description" gorm:"size:200;comment:描述"`
	Identifier                         string                               `json:"identifier" gorm:"size:100;comment:识别码"`
	AccountControl                     bool                                 `json:"accountControl" gorm:"comment:账号管控"`
	MaterialControl                    bool                                 `json:"materialControl" gorm:"comment:物料管控"`
	EfficiencyStatistic                bool                                 `json:"efficiencyStatistic" gorm:"comment:效率统计"`
	Remark                             string                               `json:"remark" gorm:"size:1000;comment:备注"`
	ProductModelID                     string                               `json:"productModelID" gorm:"index;size:36;comment:当前型号ID"`
	ProductionStations                 []*ProductionStation                 `json:"productionStations" gorm:"constraint:OnDelete:CASCADE;"`                 //工站
	ProductionCrossways                []*ProductionCrossway                `json:"productionCrossways" gorm:"constraint:OnDelete:CASCADE;"`                //产线路口
	ProductionLineSupportableCategorys []*ProductionLineSupportableCategory `json:"productionLineSupportableCategorys" gorm:"constraint:OnDelete:CASCADE;"` //产线支持产品类别
	ProductionProcesses                []*ProductionProcess                 `gorm:"constraint:OnDelete:CASCADE"`
	ProcessStepParameters              []*ProcessStepParameter              `gorm:"constraint:OnDelete:CASCADE"`
}

type ProductionLineSupportableCategory struct {
	ModelID
	ProductionLineID  string           `json:"productionLineID" gorm:"index;size:36;comment:产线ID"`
	ProductCategoryID string           `json:"productCategoryID" gorm:"size:36;comment:产品类别ID"`
	ProductCategory   *ProductCategory `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToProductionLines(in []*proto.ProductionLineInfo) []*ProductionLine {
	var result []*ProductionLine
	for _, c := range in {
		result = append(result, PBToProductionLine(c))
	}
	return result
}

func PBToProductionLine(in *proto.ProductionLineInfo) *ProductionLine {
	if in == nil {
		return nil
	}
	return &ProductionLine{
		ModelID:                            ModelID{ID: in.Id},
		Code:                               in.Code,
		Description:                        in.Description,
		Identifier:                         in.Identifier,
		AccountControl:                     in.AccountControl,
		MaterialControl:                    in.MaterialControl,
		EfficiencyStatistic:                in.EfficiencyStatistic,
		Remark:                             in.Remark,
		ProductionFactoryID:                in.ProductionFactoryID,
		ProductModelID:                     in.ProductModelID,
		ProductionStations:                 PBToProductionStations(in.ProductionStations),
		ProductionCrossways:                PBToProductionCrossways(in.ProductionCrossways),
		ProductionLineSupportableCategorys: PBToProductionLineSupportableCategorys(in.ProductionLineSupportableCategorys),
	}
}

func PBToProductionLineSupportableCategorys(in []*proto.ProductionLineSupportableCategoryInfo) []*ProductionLineSupportableCategory {
	var result []*ProductionLineSupportableCategory
	for _, c := range in {
		result = append(result, PBToProductionLineSupportableCategory(c))
	}
	return result
}

func PBToProductionLineSupportableCategory(in *proto.ProductionLineSupportableCategoryInfo) *ProductionLineSupportableCategory {
	if in == nil {
		return nil
	}
	return &ProductionLineSupportableCategory{
		ModelID:           ModelID{ID: in.Id},
		ProductionLineID:  in.ProductionLineID,
		ProductCategoryID: in.ProductCategoryID,
	}
}

func ProductionLinesToPB(in []*ProductionLine) []*proto.ProductionLineInfo {
	var list []*proto.ProductionLineInfo
	for _, f := range in {
		list = append(list, ProductionLineToPB(f))
	}
	return list
}

func ProductionLineToPB(in *ProductionLine) *proto.ProductionLineInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductionLineInfo{
		Id:                                 in.ID,
		Code:                               in.Code,
		Description:                        in.Description,
		Identifier:                         in.Identifier,
		AccountControl:                     in.AccountControl,
		MaterialControl:                    in.MaterialControl,
		EfficiencyStatistic:                in.EfficiencyStatistic,
		Remark:                             in.Remark,
		ProductionFactoryID:                in.ProductionFactoryID,
		ProductModelID:                     in.ProductModelID,
		ProductionStations:                 ProductionStationsToPB(in.ProductionStations),
		ProductionCrossways:                ProductionCrosswaysToPB(in.ProductionCrossways),
		ProductionLineSupportableCategorys: ProductionLineSupportableCategorysToPB(in.ProductionLineSupportableCategorys),
		ProductionProcesses:                ProductionProcesssToPB(in.ProductionProcesses),
		ProcessStepParameters:              ProcessStepParametersToPB(in.ProcessStepParameters),
	}
	return m
}

func ProductionLineSupportableCategorysToPB(in []*ProductionLineSupportableCategory) []*proto.ProductionLineSupportableCategoryInfo {
	var list []*proto.ProductionLineSupportableCategoryInfo
	for _, f := range in {
		list = append(list, ProductionLineSupportableCategoryToPB(f))
	}
	return list
}

func ProductionLineSupportableCategoryToPB(in *ProductionLineSupportableCategory) *proto.ProductionLineSupportableCategoryInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductionLineSupportableCategoryInfo{
		Id:                in.ID,
		ProductionLineID:  in.ProductionLineID,
		ProductCategoryID: in.ProductCategoryID,
	}
	return m
}
