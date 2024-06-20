package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 工站产量记录
type ProductionStationOutput struct {
	ModelID
	OutputTime          time.Time          `json:"OutputTime" gorm:"autoUpdateTime:nano;comment:产出时间"`
	LoginUserID         string             `json:"LoginUserID" gorm:"size:36;comment:登录人员ID"`
	ProductionStationID string             `json:"ProductionStationID" gorm:"size:36;comment:生产工站ID"`
	ProductionStation   *ProductionStation `json:"productionStation" gorm:"constraint:OnDelete:CASCADE"` //生产工站
	ProductionProcessID string             `json:"ProductionProcessID" gorm:"size:36;comment:生产工序ID"`
	ProductionProcess   *ProductionProcess `json:"productionProcess" gorm:"constraint:OnDelete:CASCADE"` //生产工序
	ProductInfoID       string             `json:"ProductInfoID" gorm:"size:36;comment:产品信息ID"`
	ProductInfo         *ProductInfo       `json:"productInfo" gorm:"constraint:OnDelete:CASCADE"` //产品信息
}

func PBToProductionStationOutputs(in []*proto.ProductionStationOutputInfo) []*ProductionStationOutput {
	var result []*ProductionStationOutput
	for _, c := range in {
		result = append(result, PBToProductionStationOutput(c))
	}
	return result
}

func PBToProductionStationOutput(in *proto.ProductionStationOutputInfo) *ProductionStationOutput {
	if in == nil {
		return nil
	}
	return &ProductionStationOutput{
		ModelID: ModelID{ID: in.Id},
		// OutputTime:          utils.ParseSqlNullTime(in.OutputTime),
		LoginUserID:         in.LoginUserID,
		ProductionStationID: in.ProductionStationID,
		ProductionProcessID: in.ProductionProcessID,
		ProductInfoID:       in.ProductInfoID,
	}
}

func ProductionStationOutputsToPB(in []*ProductionStationOutput) []*proto.ProductionStationOutputInfo {
	var list []*proto.ProductionStationOutputInfo
	for _, f := range in {
		list = append(list, ProductionStationOutputToPB(f))
	}
	return list
}

func ProductionStationOutputToPB(in *ProductionStationOutput) *proto.ProductionStationOutputInfo {
	if in == nil {
		return nil
	}

	m := &proto.ProductionStationOutputInfo{
		Id:                  in.ID,
		OutputTime:          utils.FormatTime(in.OutputTime),
		LoginUserID:         in.LoginUserID,
		ProductionStationID: in.ProductionStationID,
		ProductionStation:   ProductionStationToPB(in.ProductionStation),
		ProductionProcessID: in.ProductionProcessID,
		ProductionProcess:   ProductionProcessToPB(in.ProductionProcess),
		ProductInfoID:       in.ProductInfoID,
		ProductInfo:         ProductInfoToPB(in.ProductInfo),
	}
	return m
}
