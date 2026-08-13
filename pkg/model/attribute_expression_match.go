package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/tool"
)

// MatchAnyAttributeExpressions 判断是否命中特性表达式集合（表达式间为 OR 关系）。
// 单个表达式命中条件：存在产品工单特性，其特性ID与表达式一致，且特性值满足比较运算符。
// 表达式集合为空时直接返回 initialValue。
func MatchAnyAttributeExpressions(initialValue bool, attributeExpressions []*AttributeExpression, productOrderAttributes []*ProductOrderAttribute) (bool, error) {
	match := initialValue
	for _, attributeExpression := range attributeExpressions {
		match = false
		for _, productOrderAttribute := range productOrderAttributes {
			if productOrderAttribute.ProductAttributeID == attributeExpression.ProductAttributeID {
				b, err := tool.MathOperator(productOrderAttribute.Value, attributeExpression.MathOperator, attributeExpression.AttributeValue)
				if err != nil {
					return false, err
				}
				if b {
					match = true
					break
				}
			}
		}
		if match {
			break
		}
	}
	return match, nil
}
