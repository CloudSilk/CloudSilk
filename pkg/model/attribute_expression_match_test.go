package model

import "testing"

func TestMatchAnyAttributeExpressions(t *testing.T) {
	attributes := []*ProductOrderAttribute{
		{ProductAttributeID: "attr-color", Value: "红色"},
		{ProductAttributeID: "attr-capacity", Value: "100"},
	}

	cases := []struct {
		name         string
		initialValue bool
		expressions  []*AttributeExpression
		expect       bool
		expectErr    bool
	}{
		{
			name:         "表达式为空返回initialValue-true",
			initialValue: true,
			expect:       true,
		},
		{
			name:         "表达式为空返回initialValue-false",
			initialValue: false,
			expect:       false,
		},
		{
			name:         "单表达式命中-字符串相等",
			initialValue: false,
			expressions:  []*AttributeExpression{{ProductAttributeID: "attr-color", MathOperator: "等于", AttributeValue: "红色"}},
			expect:       true,
		},
		{
			name:         "单表达式未命中-字符串不等",
			initialValue: false,
			expressions:  []*AttributeExpression{{ProductAttributeID: "attr-color", MathOperator: "等于", AttributeValue: "蓝色"}},
			expect:       false,
		},
		{
			name:         "数值比较命中-大于",
			initialValue: false,
			expressions:  []*AttributeExpression{{ProductAttributeID: "attr-capacity", MathOperator: "大于", AttributeValue: "50"}},
			expect:       true,
		},
		{
			name:         "任一表达式命中即可-OR语义",
			initialValue: false,
			expressions: []*AttributeExpression{
				{ProductAttributeID: "attr-color", MathOperator: "等于", AttributeValue: "蓝色"},
				{ProductAttributeID: "attr-capacity", MathOperator: "大于等于", AttributeValue: "100"},
			},
			expect: true,
		},
		{
			name:         "全部未命中-false",
			initialValue: false,
			expressions: []*AttributeExpression{
				{ProductAttributeID: "attr-color", MathOperator: "等于", AttributeValue: "蓝色"},
				{ProductAttributeID: "attr-capacity", MathOperator: "大于", AttributeValue: "999"},
			},
			expect: false,
		},
		{
			name:         "表达式特性ID不存在于工单特性-false",
			initialValue: false,
			expressions:  []*AttributeExpression{{ProductAttributeID: "attr-unknown", MathOperator: "等于", AttributeValue: "x"}},
			expect:       false,
		},
		{
			name:         "数值比较遇到非数字值-报错",
			initialValue: false,
			expressions:  []*AttributeExpression{{ProductAttributeID: "attr-color", MathOperator: "大于", AttributeValue: "1"}},
			expectErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := MatchAnyAttributeExpressions(c.initialValue, c.expressions, attributes)
			if c.expectErr {
				if err == nil {
					t.Fatalf("期望报错，实际通过（结果 %v）", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("意外错误: %v", err)
			}
			if got != c.expect {
				t.Fatalf("期望 %v，实际 %v", c.expect, got)
			}
		})
	}
}
