package logic

import (
	"math"
	"testing"
)

func TestSampleSizeByLot(t *testing.T) {
	cases := []struct {
		lot    int32
		level  string
		expect int32
	}{
		{5, "II", 2},
		{12, "II", 3},
		{20, "II", 5},
		{40, "II", 8},
		{100, "II", 20},
		{300, "II", 50},
		{1000, "II", 80},
		{1, "II", 1},
		{4, "I", 2},
		{4, "III", 4}, // 不超过批量
	}
	for _, c := range cases {
		if got := SampleSizeByLot(c.lot, c.level); got != c.expect {
			t.Fatalf("批量%d水平%s：期望样本量%d，实际%d", c.lot, c.level, c.expect, got)
		}
	}
}

func TestAcceptNumber(t *testing.T) {
	if got := AcceptNumber(20, 0); got != 0 {
		t.Fatalf("未启用AQL应零缺陷接收，实际Ac=%d", got)
	}
	if got := AcceptNumber(5, 2.5); got != 0 {
		t.Fatalf("样本5/AQL2.5期望Ac=0，实际%d", got)
	}
	if got := AcceptNumber(80, 1.5); got != 3 {
		t.Fatalf("样本80/AQL1.5期望Ac=3，实际%d", got)
	}
	if got := AcceptNumber(200, 4.0); got != 14 {
		t.Fatalf("样本200/AQL4.0期望Ac=14，实际%d", got)
	}
}

func TestCalcSpc(t *testing.T) {
	// 均值10、无波动的数据 → 标准差0
	r, err := CalcSpc([]float64{10, 10, 10}, nil, nil)
	if err != nil {
		t.Fatalf("意外错误: %v", err)
	}
	if r.StdDev != 0 {
		t.Fatalf("无波动数据标准差应为0")
	}

	// 双边规格：均值10、σ≈0.5、规格[8,12]
	ls, us := 8.0, 12.0
	values := []float64{10, 10.5, 9.5, 10.2, 9.8, 10.1}
	r, err = CalcSpc(values, &ls, &us)
	if err != nil {
		t.Fatalf("意外错误: %v", err)
	}
	if r.Cpk <= 1.5 || r.Cpk > 2.5 {
		t.Fatalf("双边规格Cpk应在1.5~2.5之间，实际%v", r.Cpk)
	}
	if r.Cpk != math.Min(r.Cpu, r.Cpl) {
		t.Fatalf("双边规格Cpk应为Cpu/Cpl较小者")
	}
	// 控制限中心线=均值
	if r.Cl != r.Mean {
		t.Fatalf("控制中心线应等于均值")
	}
	if !(r.Lcl < r.Cl && r.Cl < r.Ucl) {
		t.Fatalf("控制限应满足 Lcl < Cl < Ucl")
	}

	// 单边规格
	r, err = CalcSpc(values, &ls, nil)
	if err != nil {
		t.Fatalf("意外错误: %v", err)
	}
	if r.Cpk != r.Cpl {
		t.Fatalf("仅下规格时Cpk应等于Cpl")
	}

	// 样本不足
	if _, err := CalcSpc([]float64{1}, nil, nil); err == nil {
		t.Fatalf("单样本应报错")
	}
}
