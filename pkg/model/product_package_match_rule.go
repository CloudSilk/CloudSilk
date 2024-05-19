package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

//产品包装匹配规则
type ProductPackageMatchRule struct {
	ModelID
	Priority             int32                  `json:"priority" gorm:"comment:优先级"`
	ProductPackageID     *string                `json:"productPackageID" gorm:"size:36;comment:使用包装ID"`
	ProductPackage       *ProductPackage        `json:"productPackage" gorm:"comment:使用包装ID"`
	Enable               bool                   `json:"enable" gorm:"comment:是否启用"`
	InitialValue         bool                   `json:"initialValue" gorm:"comment:默认匹配"`
	Remark               string                 `json:"remark" gorm:"size:1000;comment:备注"`
	AttributeExpressions []*AttributeExpression `json:"attributeExpressions" gorm:"polymorphic:Rule;polymorphicValue:ProductPackageMatchRule"` //特征表达式
}

func PBToProductPackageMatchRules(in []*proto.ProductPackageMatchRuleInfo) []*ProductPackageMatchRule {
	var result []*ProductPackageMatchRule
	for _, c := range in {
		result = append(result, PBToProductPackageMatchRule(c))
	}
	return result
}

func PBToProductPackageMatchRule(in *proto.ProductPackageMatchRuleInfo) *ProductPackageMatchRule {
	if in == nil {
		return nil
	}
	var productPackageID *string
	if in.ProductPackageID != "" {
		productPackageID = &in.ProductPackageID
	}
	return &ProductPackageMatchRule{
		ModelID:              ModelID{ID: in.Id},
		Priority:             in.Priority,
		ProductPackageID:     productPackageID,
		Enable:               in.Enable,
		InitialValue:         in.InitialValue,
		Remark:               in.Remark,
		AttributeExpressions: PBToAttributeExpressions(in.AttributeExpressions),
	}
}

func ProductPackageMatchRulesToPB(in []*ProductPackageMatchRule) []*proto.ProductPackageMatchRuleInfo {
	var list []*proto.ProductPackageMatchRuleInfo
	for _, f := range in {
		list = append(list, ProductPackageMatchRuleToPB(f))
	}
	return list
}

func ProductPackageMatchRuleToPB(in *ProductPackageMatchRule) *proto.ProductPackageMatchRuleInfo {
	if in == nil {
		return nil
	}
	var productPackageID string
	if in.ProductPackageID != nil {
		productPackageID = *in.ProductPackageID
	}
	m := &proto.ProductPackageMatchRuleInfo{
		AttributeExpressions: AttributeExpressionsToPB(in.AttributeExpressions),
		Id:                   in.ID,
		Priority:             in.Priority,
		ProductPackageID:     productPackageID,
		ProductPackage:       ProductPackageToPB(in.ProductPackage),
		Enable:               in.Enable,
		InitialValue:         in.InitialValue,
		Remark:               in.Remark,
	}
	return m
}
