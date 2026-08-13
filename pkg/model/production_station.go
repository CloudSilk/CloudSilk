package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 生产工站
type ProductionStation struct {
	ModelID
	Code                         string                               `json:"code" gorm:"size:50;comment:代号"`
	Description                  string                               `json:"description" gorm:"size:500;comment:描述"`
	StationType                  string                               `json:"stationType" gorm:"size:-1;comment:工位类型"`
	AccountControl               bool                                 `json:"accountControl" gorm:"comment:账号管控"`
	MaterialControl              bool                                 `json:"materialControl" gorm:"comment:物料管控"`
	AllowReport                  bool                                 `json:"allowReport" gorm:"comment:允许报工"`
	AllowRework                  bool                                 `json:"allowRework" gorm:"comment:允许返工"`
	CurrentState                 string                               `json:"currentState" gorm:"size:-1;comment:当前状态"`
	LastUpdateTime               time.Time                            `json:"lastUpdateTime" gorm:"autoUpdateTime:nano;comment:最后更新时间"`
	Remark                       string                               `json:"remark" gorm:"size:500;comment:备注"`
	ProductionLineID             string                               `json:"productionLineID" gorm:"index;size:36;comment:生产产线ID"`
	ProductionLine               *ProductionLine                      `json:"productionLine" gorm:"constraint:OnDelete:CASCADE"` //生产产线
	CurrentUserID                *string                              `json:"currentUserID" gorm:"index;size:36;comment:当前登录用户ID"`
	ProductInfoID                *string                              `json:"productInfoID" gorm:"index;size:36;comment:当前产品ID"`
	ProductInfo                  *ProductInfo                         `json:"productInfo" gorm:"constraint:OnDelete:CASCADE"` //当前产品
	AvailableProductionProcesses []*ProductionProcessAvailableStation `json:"availableProductionProcesses" gorm:"constraint:OnDelete:CASCADE"`
}

func PBToProductionStations(in []*proto.ProductionStationInfo) []*ProductionStation {
	var result []*ProductionStation
	for _, c := range in {
		result = append(result, PBToProductionStation(c))
	}
	return result
}

func PBToProductionStation(in *proto.ProductionStationInfo) *ProductionStation {
	if in == nil {
		return nil
	}
	var currentUserID, productInfoID *string
	if in.CurrentUserID != "" {
		currentUserID = &in.CurrentUserID
	}
	if in.ProductInfoID != "" {
		productInfoID = &in.ProductInfoID
	}
	return &ProductionStation{
		ModelID:         ModelID{ID: in.Id},
		Code:            in.Code,
		Description:     in.Description,
		StationType:     in.StationType,
		AccountControl:  in.AccountControl,
		MaterialControl: in.MaterialControl,
		AllowReport:     in.AllowReport,
		AllowRework:     in.AllowRework,
		CurrentState:    in.CurrentState,
		// LastUpdateTime:   utils.ParseTime(in.LastUpdateTime),
		Remark:           in.Remark,
		ProductionLineID: in.ProductionLineID,
		CurrentUserID:    currentUserID,
		ProductInfoID:    productInfoID,
	}
}

func ProductionStationsToPB(in []*ProductionStation) []*proto.ProductionStationInfo {
	var list []*proto.ProductionStationInfo
	for _, f := range in {
		list = append(list, ProductionStationToPB(f))
	}
	return list
}

func ProductionStationToPB(in *ProductionStation) *proto.ProductionStationInfo {
	if in == nil {
		return nil
	}

	var currentUserID, productInfoID string
	if in.CurrentUserID != nil {
		currentUserID = *in.CurrentUserID
	}
	if in.ProductInfoID != nil {
		productInfoID = *in.ProductInfoID
	}
	m := &proto.ProductionStationInfo{
		Id:               in.ID,
		Code:             in.Code,
		Description:      in.Description,
		StationType:      in.StationType,
		AccountControl:   in.AccountControl,
		MaterialControl:  in.MaterialControl,
		AllowReport:      in.AllowReport,
		AllowRework:      in.AllowRework,
		CurrentState:     in.CurrentState,
		LastUpdateTime:   utils.FormatTime(in.LastUpdateTime),
		Remark:           in.Remark,
		ProductionLineID: in.ProductionLineID,
		ProductionLine:   ProductionLineToPB(in.ProductionLine),
		CurrentUserID:    currentUserID,
		ProductInfoID:    productInfoID,
	}
	return m
}
