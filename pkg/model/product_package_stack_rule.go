package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductPackageStackRule struct {
	ModelID
	Priority             int32                  `json:"Priority" gorm:"comment:优先级"`
	NumberOfLayers       int32                  `json:"NumberOfLayers" gorm:"comment:层数"`
	NumberOfPackages     int32                  `json:"NumberOfPackages" gorm:"comment:箱数"`
	Enable               bool                   `json:"Enable" gorm:"comment:是否启用"`
	InitialValue         bool                   `json:"InitialValue" gorm:"comment:默认匹配"`
	Remark               string                 `json:"Remark" gorm:"size:1000;comment:备注"`
	AttributeExpressions []*AttributeExpression `json:"AttributeExpressions" gorm:"polymorphic:Rule;polymorphicValue:ProductPackageStackRule"` //特征表达式
}

func PBToProductPackageStackRules(in []*proto.ProductPackageStackRuleInfo) []*ProductPackageStackRule {
	var result []*ProductPackageStackRule
	for _, c := range in {
		result = append(result, PBToProductPackageStackRule(c))
	}
	return result
}

func PBToProductPackageStackRule(in *proto.ProductPackageStackRuleInfo) *ProductPackageStackRule {
	if in == nil {
		return nil
	}
	return &ProductPackageStackRule{
		AttributeExpressions: PBToAttributeExpressions(in.AttributeExpressions),
		ModelID:              ModelID{ID: in.Id},
		Priority:             in.Priority,
		NumberOfLayers:       in.NumberOfLayers,
		NumberOfPackages:     in.NumberOfPackages,
		Enable:               in.Enable,
		InitialValue:         in.InitialValue,
		Remark:               in.Remark,
	}
}

func ProductPackageStackRulesToPB(in []*ProductPackageStackRule) []*proto.ProductPackageStackRuleInfo {
	var list []*proto.ProductPackageStackRuleInfo
	for _, f := range in {
		list = append(list, ProductPackageStackRuleToPB(f))
	}
	return list
}

func ProductPackageStackRuleToPB(in *ProductPackageStackRule) *proto.ProductPackageStackRuleInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductPackageStackRuleInfo{
		AttributeExpressions: AttributeExpressionsToPB(in.AttributeExpressions),
		Id:                   in.ID,
		Priority:             in.Priority,
		NumberOfLayers:       in.NumberOfLayers,
		NumberOfPackages:     in.NumberOfPackages,
		Enable:               in.Enable,
		InitialValue:         in.InitialValue,
		Remark:               in.Remark,
	}
	return m
}
