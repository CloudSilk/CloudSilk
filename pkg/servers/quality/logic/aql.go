package logic

// GB/T 2828.1（ISO 2859-1）抽样方案简化实现：
// 一般检验水平 II 的批量→样本量区间表 + 按样本字码与 AQL 的接收数（Ac）查询表。

// sampleSizes 一般检验水平II：批量区间 → 样本量
var sampleSizes = []struct {
	min, max, n int
}{
	{2, 8, 2},
	{9, 15, 3},
	{16, 25, 5},
	{26, 50, 8},
	{51, 90, 13},
	{91, 150, 20},
	{151, 280, 32},
	{281, 500, 50},
	{501, 1200, 80},
	{1201, 3200, 125},
	{3201, 10000, 200},
	{10001, 35000, 315},
}

// SampleSizeByLot 根据批量确定样本量（正常检验一般水平II，简化表）
func SampleSizeByLot(lotQTY int32, sampleLevel string) int32 {
	if lotQTY <= 1 {
		return lotQTY
	}
	if sampleLevel == "I" {
		// 水平I：样本量约为水平II的一半，且不低于2
		n := SampleSizeByLot(lotQTY, "II")
		if n/2 > 2 {
			return n / 2
		}
		return 2
	}
	if sampleLevel == "III" {
		// 水平III：样本量约为水平II的两倍，且不超过批量
		n := SampleSizeByLot(lotQTY, "II") * 2
		if n < lotQTY {
			return n
		}
		return lotQTY
	}
	// 默认水平II
	n := lotQTY
	for _, s := range sampleSizes {
		if int(lotQTY) >= s.min && int(lotQTY) <= s.max {
			n = int32(s.n)
			break
		}
	}
	if n > lotQTY {
		n = lotQTY
	}
	return n
}

// acceptTable 样本量 → 常见AQL的接收数Ac（正常检验一次抽样，简化）
var acceptTable = map[int]map[float64]int{
	2:   {1.0: 0, 1.5: 0, 2.5: 0, 4.0: 0},
	3:   {1.0: 0, 1.5: 0, 2.5: 0, 4.0: 0},
	5:   {1.0: 0, 1.5: 0, 2.5: 0, 4.0: 0},
	8:   {1.0: 0, 1.5: 0, 2.5: 0, 4.0: 1},
	13:  {1.0: 0, 1.5: 0, 2.5: 1, 4.0: 1},
	20:  {1.0: 0, 1.5: 1, 2.5: 1, 4.0: 2},
	32:  {1.0: 1, 1.5: 1, 2.5: 2, 4.0: 3},
	50:  {1.0: 1, 1.5: 2, 2.5: 3, 4.0: 5},
	80:  {1.0: 2, 1.5: 3, 2.5: 5, 4.0: 7},
	125: {1.0: 3, 1.5: 5, 2.5: 7, 4.0: 10},
	200: {1.0: 5, 1.5: 7, 2.5: 10, 4.0: 14},
	315: {1.0: 7, 1.5: 10, 2.5: 14, 4.0: 21},
}

// AcceptNumber 根据样本量与AQL确定接收数Ac
func AcceptNumber(sampleQTY int32, aql float64) int {
	if aql <= 0 {
		// 未启用AQL：零缺陷判定
		return 0
	}
	// 就近匹配样本量档位
	best := 2
	for _, s := range sampleSizes {
		if int32(s.n) <= sampleQTY {
			best = s.n
		}
	}
	if m, ok := acceptTable[best]; ok {
		// 就近向下匹配AQL
		aqls := []float64{1.0, 1.5, 2.5, 4.0}
		chosen := aqls[0]
		for _, a := range aqls {
			if a <= aql {
				chosen = a
			}
		}
		return m[chosen]
	}
	// 兜底：按二项近似 n*p 上取整
	ac := int(float64(sampleQTY)*aql/100.0 + 0.999)
	return ac
}
