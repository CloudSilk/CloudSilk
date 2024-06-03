package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 生产节拍
type ProductionRhythm struct {
	ModelID
	Priority             int32                  `json:"priority" gorm:"comment:优先级"`
	StandardTime         int32                  `json:"standardTime" gorm:"comment:标准时长(秒)"`
	Enable               bool                   `json:"enable" gorm:"comment:是否启用"`
	InitialValue         bool                   `json:"initialValue" gorm:"comment:默认匹配"`
	Remark               string                 `json:"remark" gorm:"size:500;comment:备注"`
	ProductionLineID     string                 `json:"productionLineID" gorm:"index;size:36;comment:生产产线ID"`
	ProductionLine       *ProductionLine        `json:"productionLine" gorm:"constraint:OnDelete:CASCADE"`
	AttributeExpressions []*AttributeExpression `json:"attributeExpressions" gorm:"polymorphic:Rule;polymorphicValue:ProductionRhythm"` //特性表达式
}

func PBToProductionRhythms(in []*proto.ProductionRhythmInfo) []*ProductionRhythm {
	var result []*ProductionRhythm
	for _, c := range in {
		result = append(result, PBToProductionRhythm(c))
	}
	return result
}

func PBToProductionRhythm(in *proto.ProductionRhythmInfo) *ProductionRhythm {
	if in == nil {
		return nil
	}
	return &ProductionRhythm{
		ModelID:              ModelID{ID: in.Id},
		Priority:             in.Priority,
		StandardTime:         in.StandardTime,
		Enable:               in.Enable,
		InitialValue:         in.InitialValue,
		Remark:               in.Remark,
		ProductionLineID:     in.ProductionLineID,
		AttributeExpressions: PBToAttributeExpressions(in.AttributeExpressions),
	}
}

func ProductionRhythmsToPB(in []*ProductionRhythm) []*proto.ProductionRhythmInfo {
	var list []*proto.ProductionRhythmInfo
	for _, f := range in {
		list = append(list, ProductionRhythmToPB(f))
	}
	return list
}

func ProductionRhythmToPB(in *ProductionRhythm) *proto.ProductionRhythmInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductionRhythmInfo{
		AttributeExpressions: AttributeExpressionsToPB(in.AttributeExpressions),
		Id:                   in.ID,
		Priority:             in.Priority,
		StandardTime:         in.StandardTime,
		Enable:               in.Enable,
		InitialValue:         in.InitialValue,
		Remark:               in.Remark,
		ProductionLineID:     in.ProductionLineID,
	}
	return m
}
