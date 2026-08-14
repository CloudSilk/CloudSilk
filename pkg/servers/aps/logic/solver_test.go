package logic

import (
	"testing"
	"time"
)

// 换型约束：相邻工单型号不同时插入换型间隔，相同时不插入
func TestConstrainedSchedule_Changeover(t *testing.T) {
	base := testBase()
	orders := []SchedOrderEx{
		{SchedOrder: SchedOrder{OrderID: "a", QTY: 10, StdWorkTimeSec: 60, PriorityLevel: 1}, ProductModel: "M1"},
		{SchedOrder: SchedOrder{OrderID: "b", QTY: 10, StdWorkTimeSec: 60, PriorityLevel: 1}, ProductModel: "M2"},
		{SchedOrder: SchedOrder{OrderID: "c", QTY: 10, StdWorkTimeSec: 60, PriorityLevel: 1}, ProductModel: "M2"},
	}
	slots := ConstrainedSchedule(orders, base, SchedOptions{ChangeoverSec: 300})

	if len(slots) != 3 {
		t.Fatalf("应生成3个时段，实际%d", len(slots))
	}
	// a(M1)→b(M2) 换型5分钟；b→c 同型号无换型
	if !slots[1].Start.Equal(base.Add(10 * time.Minute).Add(5 * time.Minute)) {
		t.Fatalf("型号切换应插入5分钟换型，实际第二时段起点%v", slots[1].Start)
	}
	if !slots[2].Start.Equal(slots[1].End) {
		t.Fatalf("同型号相邻工单不应插入换型")
	}
}

// 型号聚簇：同优先级内相同型号被排到一起（减少换型次数）
func TestConstrainedSchedule_ModelClustering(t *testing.T) {
	base := testBase()
	orders := []SchedOrderEx{
		{SchedOrder: SchedOrder{OrderID: "1", QTY: 1, StdWorkTimeSec: 60}, ProductModel: "A"},
		{SchedOrder: SchedOrder{OrderID: "2", QTY: 1, StdWorkTimeSec: 60}, ProductModel: "B"},
		{SchedOrder: SchedOrder{OrderID: "3", QTY: 1, StdWorkTimeSec: 60}, ProductModel: "A"},
	}
	slots := ConstrainedSchedule(orders, base, SchedOptions{ChangeoverSec: 600})
	models := []string{}
	byID := map[string]string{"1": "A", "2": "B", "3": "A"}
	for _, s := range slots {
		models = append(models, byID[s.OrderID])
	}
	// 期望 A,A,B（仅一次换型）
	if models[0] != "A" || models[1] != "A" || models[2] != "B" {
		t.Fatalf("同优先级应按型号聚簇为 A,A,B，实际 %v", models)
	}
}

// 物料齐套约束：开工时间不早于齐套时间
func TestConstrainedSchedule_MaterialReady(t *testing.T) {
	base := testBase()
	ready := base.Add(30 * time.Minute)
	orders := []SchedOrderEx{
		{SchedOrder: SchedOrder{OrderID: "a", QTY: 10, StdWorkTimeSec: 60}, ReadyTime: ready, HasReady: true},
		{SchedOrder: SchedOrder{OrderID: "b", QTY: 10, StdWorkTimeSec: 60}},
	}
	slots := ConstrainedSchedule(orders, base, SchedOptions{})
	if !slots[0].Start.Equal(ready) {
		t.Fatalf("工单a应等待物料齐套后开工，期望%v，实际%v", ready, slots[0].Start)
	}
	if !slots[1].Start.Equal(slots[0].End) {
		t.Fatalf("后续工单应顺延")
	}
}

// 高优先级不受低优先级齐套时间影响顺序
func TestConstrainedSchedule_PriorityBeatsReady(t *testing.T) {
	base := testBase()
	ready := base.Add(60 * time.Minute)
	orders := []SchedOrderEx{
		{SchedOrder: SchedOrder{OrderID: "low", QTY: 1, StdWorkTimeSec: 60, PriorityLevel: 1}, ReadyTime: ready, HasReady: true},
		{SchedOrder: SchedOrder{OrderID: "high", QTY: 1, StdWorkTimeSec: 60, PriorityLevel: 9}},
	}
	slots := ConstrainedSchedule(orders, base, SchedOptions{})
	if slots[0].OrderID != "high" || slots[0].Start.Equal(ready) {
		t.Fatalf("高优先级应先排且不等低优先级的齐套时间")
	}
}

// 插单重排：保持既有顺序，新工单插入到严格更高优先级之后，后续顺延
func TestInsertOrder(t *testing.T) {
	base := testBase()
	orders := []SchedOrderEx{
		{SchedOrder: SchedOrder{OrderID: "high", QTY: 10, StdWorkTimeSec: 60, PriorityLevel: 9}, ProductModel: "M1"},
		{SchedOrder: SchedOrder{OrderID: "mid", QTY: 10, StdWorkTimeSec: 60, PriorityLevel: 5}, ProductModel: "M1"},
		{SchedOrder: SchedOrder{OrderID: "low", QTY: 10, StdWorkTimeSec: 60, PriorityLevel: 1}, ProductModel: "M1"},
	}
	existing := ConstrainedSchedule(orders, base, SchedOptions{})

	// 插入中优先级新单（优先级7）：应位于 high 之后、mid 之前
	newOrder := SchedOrderEx{SchedOrder: SchedOrder{OrderID: "new", QTY: 5, StdWorkTimeSec: 60, PriorityLevel: 7}, ProductModel: "M1"}
	newSlots, idx := InsertOrder(existing, orders, newOrder, base, SchedOptions{})

	if len(newSlots) != 4 {
		t.Fatalf("插单后应有4个时段，实际%d", len(newSlots))
	}
	if idx != 1 || newSlots[idx].OrderID != "new" {
		t.Fatalf("新工单应插在下标1（高优先级之后），实际下标%d（%s）", idx, newSlots[idx].OrderID)
	}
	if newSlots[2].OrderID != "mid" || newSlots[3].OrderID != "low" {
		t.Fatalf("既有工单相对顺序应保持")
	}
	// 后续工单顺延5分钟
	if !newSlots[2].Start.Equal(existing[1].Start.Add(5 * time.Minute)) {
		t.Fatalf("后续工单应顺延新工单工时")
	}
}

// 插入最高优先级：新工单排最前
func TestInsertOrder_TopPriority(t *testing.T) {
	base := testBase()
	orders := []SchedOrderEx{
		{SchedOrder: SchedOrder{OrderID: "a", QTY: 10, StdWorkTimeSec: 60, PriorityLevel: 5}, ProductModel: "M1"},
	}
	existing := ConstrainedSchedule(orders, base, SchedOptions{})
	newOrder := SchedOrderEx{SchedOrder: SchedOrder{OrderID: "urgent", QTY: 2, StdWorkTimeSec: 60, PriorityLevel: 10}, ProductModel: "M1"}

	newSlots, idx := InsertOrder(existing, orders, newOrder, base, SchedOptions{})
	if idx != 0 || newSlots[0].OrderID != "urgent" {
		t.Fatalf("最高优先级插单应排最前")
	}
	if !newSlots[1].Start.Equal(base.Add(2 * time.Minute)) {
		t.Fatalf("原工单应顺延2分钟")
	}
}
