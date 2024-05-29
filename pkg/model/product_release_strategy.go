package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductReleaseStrategy struct {
	ModelID
	ProductionLineID                           string                                       `json:"productionLineID" gorm:"index;size:36;comment:工厂产线ID"`
	ProductCategoryID                          string                                       `json:"productCategoryID" gorm:"index;size:36;comment:产品类别ID"`
	ProductCategory                            *ProductCategory                             `json:"productCategory" gorm:"constraint:OnDelete:CASCADE"` //产品类别
	ReleaseMethod                              int32                                        `json:"releaseMethod" gorm:"comment:投料方式"`
	Remark                                     string                                       `json:"remark" gorm:"column:Remark;comment:备注"`
	ProductReleaseStrategyComparableAttributes []*ProductReleaseStrategyComparableAttribute `json:"productReleaseStrategyComparableAttributes" gorm:"constraint:OnDelete:CASCADE"` //产品类别特性
}

type ProductReleaseStrategyComparableAttribute struct {
	ModelID
	ProductReleaseStrategyID   string                    `json:"productReleaseStrategyID" gorm:"index;size:36;comment:产品发布策略ID"`
	ProductCategoryAttributeID string                    `json:"productCategoryAttributeID" gorm:"size:36;comment:产品类别特性ID"`
	ProductCategoryAttribute   *ProductCategoryAttribute `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToProductReleaseStrategys(in []*proto.ProductReleaseStrategyInfo) []*ProductReleaseStrategy {
	var result []*ProductReleaseStrategy
	for _, c := range in {
		result = append(result, PBToProductReleaseStrategy(c))
	}
	return result
}

func PBToProductReleaseStrategy(in *proto.ProductReleaseStrategyInfo) *ProductReleaseStrategy {
	if in == nil {
		return nil
	}
	return &ProductReleaseStrategy{
		ModelID:           ModelID{ID: in.Id},
		ProductionLineID:  in.ProductionLineID,
		ProductCategoryID: in.ProductCategoryID,
		ReleaseMethod:     in.ReleaseMethod,
		Remark:            in.Remark,
		ProductReleaseStrategyComparableAttributes: PBToProductReleaseStrategyComparableAttributes(in.ProductReleaseStrategyComparableAttributes),
	}
}

func PBToProductReleaseStrategyComparableAttributes(in []*proto.ProductReleaseStrategyComparableAttributeInfo) []*ProductReleaseStrategyComparableAttribute {
	var result []*ProductReleaseStrategyComparableAttribute
	for _, c := range in {
		result = append(result, PBToProductReleaseStrategyComparableAttribute(c))
	}
	return result
}

func PBToProductReleaseStrategyComparableAttribute(in *proto.ProductReleaseStrategyComparableAttributeInfo) *ProductReleaseStrategyComparableAttribute {
	if in == nil {
		return nil
	}
	return &ProductReleaseStrategyComparableAttribute{
		ModelID:                    ModelID{ID: in.Id},
		ProductReleaseStrategyID:   in.ProductReleaseStrategyID,
		ProductCategoryAttributeID: in.ProductCategoryAttributeID,
	}
}

func ProductReleaseStrategysToPB(in []*ProductReleaseStrategy) []*proto.ProductReleaseStrategyInfo {
	var list []*proto.ProductReleaseStrategyInfo
	for _, f := range in {
		list = append(list, ProductReleaseStrategyToPB(f))
	}
	return list
}

func ProductReleaseStrategyToPB(in *ProductReleaseStrategy) *proto.ProductReleaseStrategyInfo {
	if in == nil {
		return nil
	}
	productCategoryName := ""
	if in.ProductCategory != nil {
		productCategoryName = in.ProductCategory.Description
	}
	m := &proto.ProductReleaseStrategyInfo{
		Id:                  in.ID,
		ProductionLineID:    in.ProductionLineID,
		ProductCategoryID:   in.ProductCategoryID,
		ProductCategoryName: productCategoryName,
		ReleaseMethod:       in.ReleaseMethod,
		Remark:              in.Remark,
		ProductReleaseStrategyComparableAttributes: PBToProductReleaseStrategyComparableAttributesToPB(in.ProductReleaseStrategyComparableAttributes),
	}
	return m
}

func PBToProductReleaseStrategyComparableAttributesToPB(in []*ProductReleaseStrategyComparableAttribute) []*proto.ProductReleaseStrategyComparableAttributeInfo {
	var list []*proto.ProductReleaseStrategyComparableAttributeInfo
	for _, f := range in {
		list = append(list, PBToProductReleaseStrategyComparableAttributeToPB(f))
	}
	return list
}

func PBToProductReleaseStrategyComparableAttributeToPB(in *ProductReleaseStrategyComparableAttribute) *proto.ProductReleaseStrategyComparableAttributeInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductReleaseStrategyComparableAttributeInfo{
		Id:                         in.ID,
		ProductReleaseStrategyID:   in.ProductReleaseStrategyID,
		ProductCategoryAttributeID: in.ProductCategoryAttributeID,
	}
	return m
}
