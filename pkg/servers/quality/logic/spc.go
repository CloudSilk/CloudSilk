package logic

import (
	"fmt"
	"math"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// SpcResult SPC 统计结果
type SpcResult struct {
	N      int
	Mean   float64
	StdDev float64
	Cpu    float64
	Cpl    float64
	Cpk    float64
	Ucl    float64
	Cl     float64
	Lcl    float64
}

// CalcSpc 计算过程能力与控制图界限（单值-移动极差简化的均值控制图）
// lowerSpec/upperSpec 为 nil 时表示单边规格；两者皆空时仅计算统计量与控制限。
func CalcSpc(values []float64, lowerSpec, upperSpec *float64) (*SpcResult, error) {
	n := len(values)
	if n < 2 {
		return nil, fmt.Errorf("SPC计算至少需要2个样本，当前%d个", n)
	}

	var sum float64
	for _, v := range values {
		sum += v
	}
	mean := sum / float64(n)

	var sumSq float64
	for _, v := range values {
		sumSq += (v - mean) * (v - mean)
	}
	// 样本标准差（n-1）
	stdDev := math.Sqrt(sumSq / float64(n-1))

	result := &SpcResult{
		N:      n,
		Mean:   mean,
		StdDev: stdDev,
		Cl:     mean,
		// 休哈特控制图：CL ± 3σ
		Ucl: mean + 3*stdDev,
		Lcl: mean - 3*stdDev,
	}

	if stdDev == 0 {
		// 无波动：过程能力视为充足
		result.Cpu, result.Cpl, result.Cpk = math.Inf(1), math.Inf(1), math.Inf(1)
		if lowerSpec != nil {
			result.Cpl = math.Inf(1)
		}
		if upperSpec == nil && lowerSpec == nil {
			result.Cpu, result.Cpl, result.Cpk = 0, 0, 0
		}
		return result, nil
	}

	if upperSpec != nil {
		result.Cpu = (*upperSpec - mean) / (3 * stdDev)
	}
	if lowerSpec != nil {
		result.Cpl = (mean - *lowerSpec) / (3 * stdDev)
	}

	switch {
	case upperSpec != nil && lowerSpec != nil:
		result.Cpk = math.Min(result.Cpu, result.Cpl)
	case upperSpec != nil:
		result.Cpk = result.Cpu
	case lowerSpec != nil:
		result.Cpk = result.Cpl
	default:
		result.Cpu, result.Cpl, result.Cpk = 0, 0, 0
	}

	return result, nil
}

// CalcSpcFromStandard 从检验标准加载规格限并计算 SPC
func CalcSpcFromStandard(standardID string, values []string, lowerSpec, upperSpec string) (*SpcResult, error) {
	if len(values) < 2 {
		return nil, fmt.Errorf("SPC计算至少需要2个样本，当前%d个", len(values))
	}

	var ls, us *float64
	if lowerSpec != "" {
		v, err := parseFloat(lowerSpec)
		if err != nil {
			return nil, fmt.Errorf("规格下限格式无效：%w", err)
		}
		ls = &v
	}
	if upperSpec != "" {
		v, err := parseFloat(upperSpec)
		if err != nil {
			return nil, fmt.Errorf("规格上限格式无效：%w", err)
		}
		us = &v
	}

	if standardID != "" {
		standard := &model.QualityInspectionStandard{}
		if err := model.DB.DB().First(standard, "`id` = ?", standardID).Error; err == nil {
			if ls == nil && standard.LowerLimit != "" {
				if v, err := parseFloat(standard.LowerLimit); err == nil {
					ls = &v
				}
			}
			if us == nil && standard.UpperLimit != "" {
				if v, err := parseFloat(standard.UpperLimit); err == nil {
					us = &v
				}
			}
		}
	}

	floats := make([]float64, 0, len(values))
	for _, v := range values {
		f, err := parseFloat(v)
		if err != nil {
			return nil, fmt.Errorf("样本值%q不是有效数值", v)
		}
		floats = append(floats, f)
	}

	return CalcSpc(floats, ls, us)
}

func parseFloat(s string) (float64, error) {
	var f float64
	if _, err := fmt.Sscanf(s, "%g", &f); err != nil {
		return 0, err
	}
	return f, nil
}

// SpcResultToPB 转换为响应结构
func SpcResultToPB(r *SpcResult) *proto.SpcCalcResponse {
	resp := &proto.SpcCalcResponse{
		Code:   1,
		N:      int32(r.N),
		Mean:   round6(r.Mean),
		StdDev: round6(r.StdDev),
		Ucl:    round6(r.Ucl),
		Cl:     round6(r.Cl),
		Lcl:    round6(r.Lcl),
	}
	if !math.IsInf(r.Cpu, 0) {
		resp.Cpu = round6(r.Cpu)
	}
	if !math.IsInf(r.Cpl, 0) {
		resp.Cpl = round6(r.Cpl)
	}
	if !math.IsInf(r.Cpk, 0) {
		resp.Cpk = round6(r.Cpk)
	}
	return resp
}

func round6(v float64) float64 {
	return math.Round(v*1e6) / 1e6
}
