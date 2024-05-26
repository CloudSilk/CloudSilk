package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 产品工艺路线
type ProductProcessRoute struct {
	ModelID
	WorkIndex           int32              `json:"workIndex" gorm:"comment:作业顺序"`
	RouteIndex          int32              `json:"routeIndex" gorm:"comment:工序顺序"`
	CreateTime          time.Time          `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	CurrentState        string             `json:"currentState" gorm:"size:-1;comment:当前状态"`
	LastUpdateTime      time.Time          `json:"lastUpdateTime" gorm:"autoUpdateTime:nano;comment:更新时间"`
	Remark              string             `json:"remark" gorm:"size:1000;comment:备注"`
	LastProcessID       string             `json:"lastProcessID" gorm:"size:36;comment:上步工序ID"`
	LastProcess         *ProductionProcess `json:"lastProcess" gorm:"foreignkey:last_process_id"` //上步工序
	CurrentProcessID    string             `json:"currentProcessID" gorm:"size:36;comment:当前工序ID"`
	CurrentProcess      *ProductionProcess `json:"currentProcess" gorm:"foreignkey:current_process_id"` //生产工序
	ProductionStationID string             `json:"productionStationID" gorm:"size:36;comment:执行工站ID"`
	ProductionStation   *ProductionStation `json:"productionStation" gorm:"constraint:OnDelete:CASCADE"` //工站
	ProcessUserID       string             `json:"processUserID" gorm:"size:36;comment:执行人员ID"`
	ProductInfoID       string             `json:"productInfoID" gorm:"size:36;comment:产品信息ID"`
	ProductInfo         *ProductInfo       `json:"productInfo" gorm:"constraint:OnDelete:CASCADE"` //产品
}

func PBToProductProcessRoutes(in []*proto.ProductProcessRouteInfo) []*ProductProcessRoute {
	var result []*ProductProcessRoute
	for _, c := range in {
		result = append(result, PBToProductProcessRoute(c))
	}
	return result
}

func PBToProductProcessRoute(in *proto.ProductProcessRouteInfo) *ProductProcessRoute {
	if in == nil {
		return nil
	}
	return &ProductProcessRoute{
		ModelID:    ModelID{ID: in.Id},
		WorkIndex:  in.WorkIndex,
		RouteIndex: in.RouteIndex,
		// CreateTime:          utils.ParseTime(in.CreateTime),
		CurrentState: in.CurrentState,
		// LastUpdateTime:      utils.ParseTime(in.LastUpdateTime),
		Remark:              in.Remark,
		LastProcessID:       in.LastProcessID,
		CurrentProcessID:    in.CurrentProcessID,
		ProductionStationID: in.ProductionStationID,
		ProcessUserID:       in.ProcessUserID,
		ProductInfoID:       in.ProductInfoID,
	}
}

func ProductProcessRoutesToPB(in []*ProductProcessRoute) []*proto.ProductProcessRouteInfo {
	var list []*proto.ProductProcessRouteInfo
	for _, f := range in {
		list = append(list, ProductProcessRouteToPB(f))
	}
	return list
}

func ProductProcessRouteToPB(in *ProductProcessRoute) *proto.ProductProcessRouteInfo {
	if in == nil {
		return nil
	}

	m := &proto.ProductProcessRouteInfo{
		Id:                  in.ID,
		WorkIndex:           in.WorkIndex,
		RouteIndex:          in.RouteIndex,
		CreateTime:          utils.FormatTime(in.CreateTime),
		CurrentState:        in.CurrentState,
		LastUpdateTime:      utils.FormatTime(in.LastUpdateTime),
		Remark:              in.Remark,
		LastProcessID:       in.LastProcessID,
		LastProcess:         ProductionProcessToPB(in.LastProcess),
		CurrentProcessID:    in.CurrentProcessID,
		CurrentProcess:      ProductionProcessToPB(in.CurrentProcess),
		ProductionStationID: in.ProductionStationID,
		ProductionStation:   ProductionStationToPB(in.ProductionStation),
		ProcessUserID:       in.ProcessUserID,
		ProductInfoID:       in.ProductInfoID,
		ProductInfo:         ProductInfoToPB(in.ProductInfo),
	}
	return m
}
