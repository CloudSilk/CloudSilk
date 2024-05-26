package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 产品品牌
type ProductBrand struct {
	ModelID
	Code             string             `json:"code" gorm:"index;size:100;comment:代号"`
	Description      string             `json:"description" gorm:"size:200;comment:描述"`
	Identifier       string             `json:"identifier" gorm:"size:100;comment:识别码"`
	Remark           string             `json:"remark" gorm:"size:1000;comment:备注"`
}

func PBToProductBrands(in []*proto.ProductBrandInfo) []*ProductBrand {
	var result []*ProductBrand
	for _, c := range in {
		result = append(result, PBToProductBrand(c))
	}
	return result
}

func PBToProductBrand(in *proto.ProductBrandInfo) *ProductBrand {
	if in == nil {
		return nil
	}
	return &ProductBrand{
		ModelID:     ModelID{ID: in.Id},
		Code:        in.Code,
		Description: in.Description,
		Identifier:  in.Identifier,
		Remark:      in.Remark,
	}
}

func ProductBrandsToPB(in []*ProductBrand) []*proto.ProductBrandInfo {
	var list []*proto.ProductBrandInfo
	for _, f := range in {
		list = append(list, ProductBrandToPB(f))
	}
	return list
}

func ProductBrandToPB(in *ProductBrand) *proto.ProductBrandInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductBrandInfo{
		Id:          in.ID,
		Code:        in.Code,
		Description: in.Description,
		Identifier:  in.Identifier,
		Remark:      in.Remark,
	}
	return m
}
