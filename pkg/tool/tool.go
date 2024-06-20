package tool

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func MathOperator(value, method, attributeExpressionValue string) (bool, error) {
	switch method {
	case "等于":
		return value == attributeExpressionValue, nil
	case "不等于":
		return value != attributeExpressionValue, nil
	case "大于":
		a, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return false, err
		}
		b, err := strconv.ParseFloat(attributeExpressionValue, 64)
		if err != nil {
			return false, err
		}
		return a > b, nil
	case "大于等于":
		a, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return false, err
		}
		b, err := strconv.ParseFloat(attributeExpressionValue, 64)
		if err != nil {
			return false, err
		}
		return a >= b, nil
	case "小于":
		a, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return false, err
		}
		b, err := strconv.ParseFloat(attributeExpressionValue, 64)
		if err != nil {
			return false, err
		}
		return a < b, nil
	case "小于等于":
		a, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return false, err
		}
		b, err := strconv.ParseFloat(attributeExpressionValue, 64)
		if err != nil {
			return false, err
		}
		return a <= b, nil
	case "包含":
		return strings.Contains(value, attributeExpressionValue), nil
	case "不包含":
		return !strings.Contains(value, attributeExpressionValue), nil
	case "起始于":
		return strings.HasPrefix(value, attributeExpressionValue), nil
	case "结束于":
		return strings.HasSuffix(value, attributeExpressionValue), nil
	case "包括":
		return strings.Contains(value, attributeExpressionValue), nil
	case "排除":
		return !strings.Contains(value, attributeExpressionValue), nil
	default:
		return false, fmt.Errorf("未知的比较方法:%s", method)
	}
}

func Contains(v string, list []string) bool {
	for _, item := range list {
		if item == v {
			return true
		}
	}
	return false
}

func Time2NullTime(time time.Time) sql.NullTime {
	return sql.NullTime{Time: time, Valid: true}
}

func NullTime2Time(time sql.NullTime) time.Time {
	return time.Time
}
