package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 生产工步(生产工序步骤)
type ProductionProcessStep struct {
	ModelID
	SortIndex            int32                  `json:"sortIndex" gorm:"comment:排序"`
	Code                 string                 `json:"code" gorm:"index;size:100;comment:代号"`
	Description          string                 `json:"description" gorm:"size:1000;comment:描述"`
	Graphic              string                 `json:"graphic" gorm:"size:1000;comment:图示"`
	GroupCode            string                 `json:"groupCode" gorm:"size:100;comment:采集组"`
	Enable               bool                   `json:"enable" gorm:"comment:是否启用"`
	InitialValue         bool                   `json:"initialValue" gorm:"comment:默认匹配"`
	ProcessControl       bool                   `json:"processControl" gorm:"comment:工序管控"`
	Remark               string                 `json:"remark" gorm:"size:1000;comment:备注"`
	ProductionLineID     string                 `json:"productionLineID" gorm:"comment:生产产线ID"`
	ProcessStepTypeID    string                 `json:"processStepTypeID" gorm:"comment:工步类型ID"`
	ProcessStepType      *ProcessStepType       `gorm:"constraint:OnDelete:CASCADE"`                                                         //工步类型
	AttributeExpressions []*AttributeExpression `json:"attributeExpressions" gorm:"polymorphic:Rule;polymorphicValue:ProductionProcessStep"` //特性表达式
	AvailableProcesses   []*AvailableProcess    `json:"availableProcesses" gorm:"constraint:OnDelete:CASCADE"`                               //可使用的工序
}

// 可使用的工序
type AvailableProcess struct {
	ModelID
	ProductionProcessStepID string             `json:"productionProcessStepID" gorm:"index;comment:生产工步ID"`
	ProductionProcessID     string             `json:"productionProcessID" gorm:"comment:生产工序ID"`
	ProductionProcess       *ProductionProcess `json:"productionProcess" gorm:"constraint:OnDelete:CASCADE"` //生产工序
}

func PBToProductionProcessSteps(in []*proto.ProductionProcessStepInfo) []*ProductionProcessStep {
	var result []*ProductionProcessStep
	for _, c := range in {
		result = append(result, PBToProductionProcessStep(c))
	}
	return result
}

func PBToProductionProcessStep(in *proto.ProductionProcessStepInfo) *ProductionProcessStep {
	if in == nil {
		return nil
	}
	return &ProductionProcessStep{
		ModelID:              ModelID{ID: in.Id},
		SortIndex:            in.SortIndex,
		Code:                 in.Code,
		Description:          in.Description,
		Graphic:              in.Graphic,
		GroupCode:            in.GroupCode,
		Enable:               in.Enable,
		InitialValue:         in.InitialValue,
		ProcessControl:       in.ProcessControl,
		Remark:               in.Remark,
		ProductionLineID:     in.ProductionLineID,
		ProcessStepTypeID:    in.ProcessStepTypeID,
		AttributeExpressions: PBToAttributeExpressions(in.AttributeExpressions),
		AvailableProcesses:   PBToAvailableProcesss(in.AvailableProcesses),
	}
}

func ProductionProcessStepsToPB(in []*ProductionProcessStep) []*proto.ProductionProcessStepInfo {
	var list []*proto.ProductionProcessStepInfo
	for _, f := range in {
		list = append(list, ProductionProcessStepToPB(f))
	}
	return list
}

func ProductionProcessStepToPB(in *ProductionProcessStep) *proto.ProductionProcessStepInfo {
	if in == nil {
		return nil
	}

	var availableProcessIDs []string
	for _, AvailableProcess := range in.AvailableProcesses {
		availableProcessIDs = append(availableProcessIDs, AvailableProcess.ProductionProcessID)
	}

	m := &proto.ProductionProcessStepInfo{
		Id:                   in.ID,
		SortIndex:            in.SortIndex,
		Code:                 in.Code,
		Description:          in.Description,
		Graphic:              in.Graphic,
		GroupCode:            in.GroupCode,
		Enable:               in.Enable,
		InitialValue:         in.InitialValue,
		ProcessControl:       in.ProcessControl,
		Remark:               in.Remark,
		ProductionLineID:     in.ProductionLineID,
		ProcessStepTypeID:    in.ProcessStepTypeID,
		ProcessStepType:      ProcessStepTypeToPB(in.ProcessStepType),
		AvailableProcessIDs:  availableProcessIDs,
		AvailableProcesses:   AvailableProcesssToPB(in.AvailableProcesses),
		AttributeExpressions: AttributeExpressionsToPB(in.AttributeExpressions),
	}
	return m
}

func PBToAvailableProcesss(in []*proto.AvailableProcessInfo) []*AvailableProcess {
	var result []*AvailableProcess
	for _, c := range in {
		result = append(result, PBToAvailableProcess(c))
	}
	return result
}

func PBToAvailableProcess(in *proto.AvailableProcessInfo) *AvailableProcess {
	if in == nil {
		return nil
	}
	return &AvailableProcess{
		ModelID:                 ModelID{ID: in.Id},
		ProductionProcessStepID: in.ProductionProcessStepID,
		ProductionProcessID:     in.ProductionProcessID,
	}
}

func AvailableProcesssToPB(in []*AvailableProcess) []*proto.AvailableProcessInfo {
	var list []*proto.AvailableProcessInfo
	for _, f := range in {
		list = append(list, AvailableProcessToPB(f))
	}
	return list
}

func AvailableProcessToPB(in *AvailableProcess) *proto.AvailableProcessInfo {
	if in == nil {
		return nil
	}

	m := &proto.AvailableProcessInfo{
		Id:                      in.ID,
		ProductionProcessStepID: in.ProductionProcessStepID,
		ProductionProcessID:     in.ProductionProcessID,
		ProductionProcess:       ProductionProcessToPB(in.ProductionProcess),
	}
	return m
}
