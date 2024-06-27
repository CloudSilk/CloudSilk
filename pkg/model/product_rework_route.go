package model

import "github.com/CloudSilk/CloudSilk/pkg/proto"

// 产品返工路线
type ProductReworkRoute struct {
	ModelID
	ProductionLineID   string             `gorm:"size:36;comment:生产产线ID"`
	ProductionLine     *ProductionLine    `gorm:"constraint:OnDelete:CASCADE"`
	MaterialCategoryID string             `gorm:"size:36;comment:物料类别ID"`
	MaterialCategory   *MaterialCategory  `gorm:"constraint:OnDelete:CASCADE"`
	FollowProcessID    *string            `gorm:"size:36;comment:后续工序ID"`
	FollowProcess      *ProductionProcess `gorm:"constraint:OnDelete:SET NULL"`
}

func PBToProductReworkRoutes(in []*proto.ProductReworkRouteInfo) []*ProductReworkRoute {
	var result []*ProductReworkRoute
	for _, c := range in {
		result = append(result, PBToProductReworkRoute(c))
	}
	return result
}

func PBToProductReworkRoute(in *proto.ProductReworkRouteInfo) *ProductReworkRoute {
	if in == nil {
		return nil
	}

	var followProcessID *string
	if in.FollowProcessID != "" {
		followProcessID = &in.FollowProcessID
	}

	return &ProductReworkRoute{
		ModelID:            ModelID{ID: in.Id},
		ProductionLineID:   in.ProductionLineID,
		MaterialCategoryID: in.MaterialCategoryID,
		FollowProcessID:    followProcessID,
	}
}

func ProductReworkRoutesToPB(in []*ProductReworkRoute) []*proto.ProductReworkRouteInfo {
	var list []*proto.ProductReworkRouteInfo
	for _, f := range in {
		list = append(list, ProductReworkRouteToPB(f))
	}
	return list
}

func ProductReworkRouteToPB(in *ProductReworkRoute) *proto.ProductReworkRouteInfo {
	if in == nil {
		return nil
	}

	var followProcessID string
	if in.FollowProcessID != nil {
		followProcessID = *in.FollowProcessID
	}

	m := &proto.ProductReworkRouteInfo{
		Id:                 in.ID,
		ProductionLineID:   in.ProductionLineID,
		ProductionLine:     ProductionLineToPB(in.ProductionLine),
		MaterialCategoryID: in.MaterialCategoryID,
		MaterialCategory:   MaterialCategoryToPB(in.MaterialCategory),
		FollowProcessID:    followProcessID,
		FollowProcess:      ProductionProcessToPB(in.FollowProcess),
	}
	return m
}
