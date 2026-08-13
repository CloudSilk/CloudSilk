package model

import "github.com/CloudSilk/CloudSilk/pkg/proto"

type MaterialStoreFeedRule struct {
	ModelID
	SortIndex            int32                  `gorm:"comment:排序"`
	MaterialStoreID      string                 `gorm:"size:36;comment:物料仓库ID"`
	MaterialStore        *MaterialStore         `gorm:"constraint:OnDelete:CASCADE"` //物料仓库
	MaterialInfoID       string                 `gorm:"size:36;comment:物料信息ID"`
	MaterialInfo         *MaterialInfo          `gorm:"constraint:OnDelete:CASCADE"` //物料信息
	MinimumStoreQTY      int64                  `gorm:"comment:最低库存量"`
	DefaultRequestQTY    int64                  `gorm:"comment:默认申请量"`
	Enable               bool                   `gorm:"comment:是否启用"`
	Remark               string                 `gorm:"size:500;comment:备注"`
	Priority             int32                  `gorm:"comment:优先级"`
	AttributeExpressions []*AttributeExpression `gorm:"polymorphic:Rule;polymorphicValue:MaterialStoreFeedRule"` //特征表达式
}

func PBToMaterialStoreFeedRules(in []*proto.MaterialStoreFeedRuleInfo) []*MaterialStoreFeedRule {
	var result []*MaterialStoreFeedRule
	for _, c := range in {
		result = append(result, PBToMaterialStoreFeedRule(c))
	}
	return result
}

func PBToMaterialStoreFeedRule(in *proto.MaterialStoreFeedRuleInfo) *MaterialStoreFeedRule {
	if in == nil {
		return nil
	}

	return &MaterialStoreFeedRule{
		ModelID:              ModelID{ID: in.Id},
		SortIndex:            in.SortIndex,
		MaterialStoreID:      in.MaterialStoreID,
		MaterialInfoID:       in.MaterialInfoID,
		MinimumStoreQTY:      in.MinimumStoreQTY,
		DefaultRequestQTY:    in.DefaultRequestQTY,
		Enable:               in.Enable,
		Remark:               in.Remark,
		Priority:             in.Priority,
		AttributeExpressions: PBToAttributeExpressions(in.AttributeExpressions),
	}
}

func MaterialStoreFeedRulesToPB(in []*MaterialStoreFeedRule) []*proto.MaterialStoreFeedRuleInfo {
	var list []*proto.MaterialStoreFeedRuleInfo
	for _, f := range in {
		list = append(list, MaterialStoreFeedRuleToPB(f))
	}
	return list
}

func MaterialStoreFeedRuleToPB(in *MaterialStoreFeedRule) *proto.MaterialStoreFeedRuleInfo {
	if in == nil {
		return nil
	}

	m := &proto.MaterialStoreFeedRuleInfo{
		Id:                   in.ID,
		SortIndex:            in.SortIndex,
		MaterialStoreID:      in.MaterialStoreID,
		MaterialStore:        MaterialStoreToPB(in.MaterialStore),
		MaterialInfoID:       in.MaterialInfoID,
		MaterialInfo:         MaterialInfoToPB(in.MaterialInfo),
		MinimumStoreQTY:      in.MinimumStoreQTY,
		DefaultRequestQTY:    in.DefaultRequestQTY,
		Enable:               in.Enable,
		Remark:               in.Remark,
		Priority:             in.Priority,
		AttributeExpressions: AttributeExpressionsToPB(in.AttributeExpressions),
	}
	return m
}
