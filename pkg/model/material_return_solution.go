package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 物料退料方案
type MaterialReturnSolution struct {
	ModelID
	Code                 string                                  `gorm:"size:50;comment:代号"`          //代号
	Description          string                                  `gorm:"size:500;comment:描述"`         //描述
	Remark               string                                  `gorm:"size:500;comment:备注"`         //备注
	MaterialReturnCauses []*MaterialReturnSolutionAvailableCause `gorm:"constraint:OnDelete:CASCADE"` //退料原因
}

// 物料退料方案-退料原因
type MaterialReturnSolutionAvailableCause struct {
	ModelID
	MaterialReturnSolutionID string               `gorm:"index;size:36;comment:物料退料方案ID"`
	MaterialReturnCauseID    string               `gorm:"size:36;comment:物料退料原因ID"`
	MaterialReturnCause      *MaterialReturnCause `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToMaterialReturnSolutions(in []*proto.MaterialReturnSolutionInfo) []*MaterialReturnSolution {
	var result []*MaterialReturnSolution
	for _, c := range in {
		result = append(result, PBToMaterialReturnSolution(c))
	}
	return result
}

func PBToMaterialReturnSolution(in *proto.MaterialReturnSolutionInfo) *MaterialReturnSolution {
	if in == nil {
		return nil
	}

	return &MaterialReturnSolution{
		ModelID:              ModelID{ID: in.Id},
		Code:                 in.Code,
		Description:          in.Description,
		Remark:               in.Remark,
		MaterialReturnCauses: PBToMaterialReturnSolutionAvailableCauses(in.MaterialReturnCauses),
	}
}

func MaterialReturnSolutionsToPB(in []*MaterialReturnSolution) []*proto.MaterialReturnSolutionInfo {
	var list []*proto.MaterialReturnSolutionInfo
	for _, f := range in {
		list = append(list, MaterialReturnSolutionToPB(f))
	}
	return list
}

func MaterialReturnSolutionToPB(in *MaterialReturnSolution) *proto.MaterialReturnSolutionInfo {
	if in == nil {
		return nil
	}

	m := &proto.MaterialReturnSolutionInfo{
		Id:                   in.ID,
		Code:                 in.Code,
		Description:          in.Description,
		Remark:               in.Remark,
		MaterialReturnCauses: MaterialReturnSolutionAvailableCausesToPB(in.MaterialReturnCauses),
	}
	return m
}

func PBToMaterialReturnSolutionAvailableCauses(in []*proto.MaterialReturnSolutionAvailableCauseInfo) []*MaterialReturnSolutionAvailableCause {
	var result []*MaterialReturnSolutionAvailableCause
	for _, c := range in {
		result = append(result, PBToMaterialReturnSolutionAvailableCause(c))
	}
	return result
}

func PBToMaterialReturnSolutionAvailableCause(in *proto.MaterialReturnSolutionAvailableCauseInfo) *MaterialReturnSolutionAvailableCause {
	if in == nil {
		return nil
	}

	return &MaterialReturnSolutionAvailableCause{
		ModelID:                  ModelID{ID: in.Id},
		MaterialReturnSolutionID: in.MaterialReturnSolutionID,
		MaterialReturnCauseID:    in.MaterialReturnCauseID,
	}
}

func MaterialReturnSolutionAvailableCausesToPB(in []*MaterialReturnSolutionAvailableCause) []*proto.MaterialReturnSolutionAvailableCauseInfo {
	var list []*proto.MaterialReturnSolutionAvailableCauseInfo
	for _, f := range in {
		list = append(list, MaterialReturnSolutionAvailableCauseToPB(f))
	}
	return list
}

func MaterialReturnSolutionAvailableCauseToPB(in *MaterialReturnSolutionAvailableCause) *proto.MaterialReturnSolutionAvailableCauseInfo {
	if in == nil {
		return nil
	}

	m := &proto.MaterialReturnSolutionAvailableCauseInfo{
		Id:                       in.ID,
		MaterialReturnSolutionID: in.MaterialReturnSolutionID,
		MaterialReturnCauseID:    in.MaterialReturnCauseID,
		MaterialReturnCause:      MaterialReturnCauseToPB(in.MaterialReturnCause),
	}
	return m
}
