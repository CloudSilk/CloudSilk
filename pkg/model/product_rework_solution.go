package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 产品返工方案
type ProductReworkSolution struct {
	ModelID
	Code                                 string                                 `json:"code" gorm:"index;size:100;comment:代号"`
	Description                          string                                 `json:"description" gorm:"size:1000;comment:描述"`
	Remark                               string                                 `json:"remark" gorm:"size:1000;comment:备注"`
	ProductReworkCauseAvailableSolutions []*ProductReworkCauseAvailableSolution `json:"productReworkCauseAvailableSolutions" gorm:"constraint:OnDelete:CASCADE"` //返工原因
}

// 支持返工原因
type ProductReworkCauseAvailableSolution struct {
	ModelID
	ProductReworkSolutionID string              `json:"productReworkSolutionID" gorm:"index;size:36;comment:返工解决方案ID"`
	ProductReworkCauseID    string              `json:"productReworkCauseID" gorm:"size:36;comment:返工原因ID"`
	ProductReworkCause      *ProductReworkCause `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToProductReworkSolutions(in []*proto.ProductReworkSolutionInfo) []*ProductReworkSolution {
	var result []*ProductReworkSolution
	for _, c := range in {
		result = append(result, PBToProductReworkSolution(c))
	}
	return result
}

func PBToProductReworkSolution(in *proto.ProductReworkSolutionInfo) *ProductReworkSolution {
	if in == nil {
		return nil
	}
	return &ProductReworkSolution{
		ModelID:                              ModelID{ID: in.Id},
		Code:                                 in.Code,
		Description:                          in.Description,
		Remark:                               in.Remark,
		ProductReworkCauseAvailableSolutions: PBToProductReworkCauseAvailableSolutions(in.ProductReworkCauseAvailableSolutions),
	}
}

func PBToProductReworkCauseAvailableSolutions(in []*proto.ProductReworkCauseAvailableSolutionInfo) []*ProductReworkCauseAvailableSolution {
	var result []*ProductReworkCauseAvailableSolution
	for _, c := range in {
		result = append(result, PBToProductReworkCauseAvailableSolution(c))
	}
	return result
}

func PBToProductReworkCauseAvailableSolution(in *proto.ProductReworkCauseAvailableSolutionInfo) *ProductReworkCauseAvailableSolution {
	if in == nil {
		return nil
	}
	return &ProductReworkCauseAvailableSolution{
		ModelID:                 ModelID{ID: in.Id},
		ProductReworkSolutionID: in.ProductReworkSolutionID,
		ProductReworkCauseID:    in.ProductReworkCauseID,
	}
}

func ProductReworkSolutionsToPB(in []*ProductReworkSolution) []*proto.ProductReworkSolutionInfo {
	var list []*proto.ProductReworkSolutionInfo
	for _, f := range in {
		list = append(list, ProductReworkSolutionToPB(f))
	}
	return list
}

func ProductReworkSolutionToPB(in *ProductReworkSolution) *proto.ProductReworkSolutionInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductReworkSolutionInfo{
		Id:                                   in.ID,
		Code:                                 in.Code,
		Description:                          in.Description,
		Remark:                               in.Remark,
		ProductReworkCauseAvailableSolutions: ProductReworkCauseAvailableSolutionsToPB(in.ProductReworkCauseAvailableSolutions),
	}
	return m
}

func ProductReworkCauseAvailableSolutionsToPB(in []*ProductReworkCauseAvailableSolution) []*proto.ProductReworkCauseAvailableSolutionInfo {
	var list []*proto.ProductReworkCauseAvailableSolutionInfo
	for _, f := range in {
		list = append(list, ProductReworkCauseAvailableSolutionToPB(f))
	}
	return list
}

func ProductReworkCauseAvailableSolutionToPB(in *ProductReworkCauseAvailableSolution) *proto.ProductReworkCauseAvailableSolutionInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductReworkCauseAvailableSolutionInfo{
		Id:                      in.ID,
		ProductReworkSolutionID: in.ProductReworkSolutionID,
		ProductReworkCauseID:    in.ProductReworkCauseID,
	}
	return m
}
