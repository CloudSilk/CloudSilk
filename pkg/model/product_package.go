package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 产品包装
type ProductPackage struct {
	ModelID
	Code                 string              `json:"code" gorm:"index;size:100;comment:代号"`
	Description          string              `json:"description" gorm:"size:200;comment:描述"`
	Identifier           string              `json:"identifier" gorm:"size:100;comment:识别码"`
	PackageQuantity      int32               `json:"packageQuantity" gorm:"comment:包装数量"`
	QuantityUnit         string              `json:"quantityUnit" gorm:"size:200;comment:数量单位"`
	GrossWeight          float32             `json:"grossWeight" gorm:"comment:毛重"`
	NetWeight            float32             `json:"netWeight" gorm:"comment:净重"`
	WeightUnit           string              `json:"weightUnit" gorm:"size:200;comment:重量单位"`
	Measure              string              `json:"measure" gorm:"size:200;comment:尺寸"`
	MeasureUnit          string              `json:"measureUnit" gorm:"size:200;comment:尺寸单位"`
	Remark               string              `json:"remark" gorm:"size:1000;comment:备注"`
	ProductPackageTypeID *string             `json:"productPackageTypeID" gorm:"comment:包装类型ID"`
	ProductPackageType   *ProductPackageType `json:"productPackageType"` //包装类型
}

func PBToProductPackages(in []*proto.ProductPackageInfo) []*ProductPackage {
	var result []*ProductPackage
	for _, c := range in {
		result = append(result, PBToProductPackage(c))
	}
	return result
}

func PBToProductPackage(in *proto.ProductPackageInfo) *ProductPackage {
	if in == nil {
		return nil
	}
	var productPackageTypeID *string
	if in.ProductPackageTypeID != "" {
		productPackageTypeID = &in.ProductPackageTypeID
	}
	return &ProductPackage{
		ModelID:              ModelID{ID: in.Id},
		Code:                 in.Code,
		Description:          in.Description,
		Identifier:           in.Identifier,
		PackageQuantity:      in.PackageQuantity,
		QuantityUnit:         in.QuantityUnit,
		GrossWeight:          in.GrossWeight,
		NetWeight:            in.NetWeight,
		WeightUnit:           in.WeightUnit,
		Measure:              in.Measure,
		MeasureUnit:          in.MeasureUnit,
		Remark:               in.Remark,
		ProductPackageTypeID: productPackageTypeID,
	}
}

func ProductPackagesToPB(in []*ProductPackage) []*proto.ProductPackageInfo {
	var list []*proto.ProductPackageInfo
	for _, f := range in {
		list = append(list, ProductPackageToPB(f))
	}
	return list
}

func ProductPackageToPB(in *ProductPackage) *proto.ProductPackageInfo {
	if in == nil {
		return nil
	}
	var productPackageTypeID string
	if in.ProductPackageTypeID != nil {
		productPackageTypeID = *in.ProductPackageTypeID
	}
	m := &proto.ProductPackageInfo{
		Id:                   in.ID,
		Code:                 in.Code,
		Description:          in.Description,
		Identifier:           in.Identifier,
		PackageQuantity:      in.PackageQuantity,
		QuantityUnit:         in.QuantityUnit,
		GrossWeight:          in.GrossWeight,
		NetWeight:            in.NetWeight,
		WeightUnit:           in.WeightUnit,
		Measure:              in.Measure,
		MeasureUnit:          in.MeasureUnit,
		Remark:               in.Remark,
		ProductPackageTypeID: productPackageTypeID,
		ProductPackageType:   ProductPackageTypeToPB(in.ProductPackageType),
	}
	return m
}
