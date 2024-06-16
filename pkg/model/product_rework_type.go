package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 返工类型
type ProductReworkType struct {
	ModelID
	Code        string `json:"code" gorm:"index;size:100;comment:代号"`
	Description string `json:"description" gorm:"size:1000;comment:描述"`
	Remark      string `json:"remark" gorm:"size:1000;comment:备注"`
}

func PBToProductReworkTypes(in []*proto.ProductReworkTypeInfo) []*ProductReworkType {
	var result []*ProductReworkType
	for _, c := range in {
		result = append(result, PBToProductReworkType(c))
	}
	return result
}

func PBToProductReworkType(in *proto.ProductReworkTypeInfo) *ProductReworkType {
	if in == nil {
		return nil
	}
	return &ProductReworkType{
		ModelID:     ModelID{ID: in.Id},
		Code:        in.Code,
		Description: in.Description,
		Remark:      in.Remark,
	}
}

func ProductReworkTypesToPB(in []*ProductReworkType) []*proto.ProductReworkTypeInfo {
	var list []*proto.ProductReworkTypeInfo
	for _, f := range in {
		list = append(list, ProductReworkTypeToPB(f))
	}
	return list
}

func ProductReworkTypeToPB(in *ProductReworkType) *proto.ProductReworkTypeInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductReworkTypeInfo{
		Id:          in.ID,
		Code:        in.Code,
		Description: in.Description,
		Remark:      in.Remark,
	}
	return m
}
