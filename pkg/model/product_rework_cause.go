package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 产品返工理由
type ProductReworkCause struct {
	ModelID
	Code                            string                            `json:"code" gorm:"index;size:100;comment:代号"`
	Description                     string                            `json:"description" gorm:"size:1000;comment:描述"`
	Remark                          string                            `json:"remark" gorm:"size:1000;comment:备注"`
	ProductReworkTypePossibleCauses []*ProductReworkTypePossibleCause `json:"productReworkTypePossibleCauses" gorm:"constraint:OnDelete:CASCADE"` //支持返工类型
}

// 支持返工类型
type ProductReworkTypePossibleCause struct {
	ModelID
	ProductReworkCauseID string             `json:"productReworkCauseID" gorm:"index;size:36;comment:返工原因ID"`
	ProductReworkTypeID  string             `json:"productReworkTypeID" gorm:"size:36;comment:返工类型ID;"`
	ProductReworkType    *ProductReworkType `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToProductReworkCauses(in []*proto.ProductReworkCauseInfo) []*ProductReworkCause {
	var result []*ProductReworkCause
	for _, c := range in {
		result = append(result, PBToProductReworkCause(c))
	}
	return result
}

func PBToProductReworkCause(in *proto.ProductReworkCauseInfo) *ProductReworkCause {
	if in == nil {
		return nil
	}
	return &ProductReworkCause{
		ModelID:                         ModelID{ID: in.Id},
		Code:                            in.Code,
		Description:                     in.Description,
		Remark:                          in.Remark,
		ProductReworkTypePossibleCauses: PBToProductReworkTypePossibleCauses(in.ProductReworkTypePossibleCauses),
	}
}

func PBToProductReworkTypePossibleCauses(in []*proto.ProductReworkTypePossibleCauseInfo) []*ProductReworkTypePossibleCause {
	var result []*ProductReworkTypePossibleCause
	for _, c := range in {
		result = append(result, PBToProductReworkTypePossibleCause(c))
	}
	return result
}

func PBToProductReworkTypePossibleCause(in *proto.ProductReworkTypePossibleCauseInfo) *ProductReworkTypePossibleCause {
	if in == nil {
		return nil
	}
	return &ProductReworkTypePossibleCause{
		ModelID:              ModelID{ID: in.Id},
		ProductReworkCauseID: in.ProductReworkCauseID,
		ProductReworkTypeID:  in.ProductReworkTypeID,
	}
}

func ProductReworkCausesToPB(in []*ProductReworkCause) []*proto.ProductReworkCauseInfo {
	var list []*proto.ProductReworkCauseInfo
	for _, f := range in {
		list = append(list, ProductReworkCauseToPB(f))
	}
	return list
}

func ProductReworkCauseToPB(in *ProductReworkCause) *proto.ProductReworkCauseInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductReworkCauseInfo{
		Id:                              in.ID,
		Code:                            in.Code,
		Description:                     in.Description,
		Remark:                          in.Remark,
		ProductReworkTypePossibleCauses: ProductReworkTypePossibleCausesToPB(in.ProductReworkTypePossibleCauses),
	}
	return m
}

func ProductReworkTypePossibleCausesToPB(in []*ProductReworkTypePossibleCause) []*proto.ProductReworkTypePossibleCauseInfo {
	var list []*proto.ProductReworkTypePossibleCauseInfo
	for _, f := range in {
		list = append(list, ProductReworkTypePossibleCauseToPB(f))
	}
	return list
}

func ProductReworkTypePossibleCauseToPB(in *ProductReworkTypePossibleCause) *proto.ProductReworkTypePossibleCauseInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductReworkTypePossibleCauseInfo{
		Id:                   in.ID,
		ProductReworkCauseID: in.ProductReworkCauseID,
		ProductReworkTypeID:  in.ProductReworkTypeID,
	}
	return m
}
