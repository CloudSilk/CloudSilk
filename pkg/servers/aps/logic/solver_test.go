package logic

import (
	"testing"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/testutil"
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

// 调度器注册：自定义算法可替换内置贪心，编排层无感
func TestSchedulerRegistry(t *testing.T) {
	original := CurrentScheduler()
	defer SetScheduler(original)

	calls := 0
	stub := &schedulerStub{calls: &calls}
	SetScheduler(stub)

	base := testBase()
	CurrentScheduler().Schedule([]SchedOrderEx{{SchedOrder: SchedOrder{OrderID: "a", QTY: 1, StdWorkTimeSec: 60}}}, base, SchedOptions{})
	if calls != 1 {
		t.Fatalf("自定义调度器应被调用，实际%d次", calls)
	}
	if CurrentScheduler() != Scheduler(stub) {
		t.Fatalf("注册后应返回新调度器")
	}
}

type schedulerStub struct{ calls *int }

func (s *schedulerStub) Schedule(orders []SchedOrderEx, startTime time.Time, opts SchedOptions) []SchedSlot {
	*s.calls++
	return ConstrainedSchedule(orders, startTime, opts)
}

// 人工调整：级联顺延保持间隔，前置明细不受影响，计划区间刷新
func TestAdjustScheduleItem_Cascade(t *testing.T) {
	gdb, err := testutil.SetupTestDB()
	if err != nil {
		t.Fatalf("初始化测试数据库失败: %v", err)
	}
	base := testBase()

	plan := &model.ProductionSchedulePlan{
		PlanNo: "SC-001", ProductionLineID: "line-1",
		StartTime: base, EndTime: base.Add(30 * time.Minute),
		OrderCount: 3, CurrentState: ScheduleStateGenerated,
	}
	if err := gdb.Create(plan).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	items := []*model.ProductionScheduleItem{
		{ProductionSchedulePlanID: plan.ID, ProductOrderID: "o1", Sequence: 1, PlannedStartTime: base, PlannedEndTime: base.Add(10 * time.Minute), DurationSeconds: 600},
		{ProductionSchedulePlanID: plan.ID, ProductOrderID: "o2", Sequence: 2, PlannedStartTime: base.Add(10 * time.Minute), PlannedEndTime: base.Add(20 * time.Minute), DurationSeconds: 600},
		{ProductionSchedulePlanID: plan.ID, ProductOrderID: "o3", Sequence: 3, PlannedStartTime: base.Add(20 * time.Minute), PlannedEndTime: base.Add(30 * time.Minute), DurationSeconds: 600},
	}
	for _, it := range items {
		if err := gdb.Create(it).Error; err != nil {
			t.Fatalf("造数失败: %v", err)
		}
	}

	// 把第2条推迟5分钟并顺延
	newStart := base.Add(15 * time.Minute)
	if err := AdjustScheduleItem(&proto.AdjustScheduleItemRequest{
		PlanID: plan.ID, ItemID: items[1].ID,
		NewStartTime: newStart.Format("2006-01-02 15:04:05"), Cascade: true,
	}, "user-9"); err != nil {
		t.Fatalf("调整失败: %v", err)
	}

	var after []*model.ProductionScheduleItem
	gdb.Where("`production_schedule_plan_id` = ?", plan.ID).Order("sequence").Find(&after)
	if !after[0].PlannedStartTime.Equal(base) {
		t.Fatalf("前置明细不应移动")
	}
	if !after[1].PlannedStartTime.Equal(newStart) {
		t.Fatalf("目标明细应移动到新开工时间")
	}
	if !after[2].PlannedStartTime.Equal(base.Add(25 * time.Minute)) {
		t.Fatalf("后续明细应顺延5分钟，实际%v", after[2].PlannedStartTime)
	}

	p := &model.ProductionSchedulePlan{}
	gdb.First(p, "`id` = ?", plan.ID)
	if !p.EndTime.Equal(base.Add(35 * time.Minute)) {
		t.Fatalf("计划结束时间应刷新为+35分钟，实际%v", p.EndTime)
	}

	var traces int64
	gdb.Model(&model.OperationTrace{}).Where("`action_name` = ?", "人工调整").Count(&traces)
	if traces != 1 {
		t.Fatalf("应记录1条人工调整轨迹")
	}
}

// 非级联模式仅移动目标明细；已下发计划拒绝调整
func TestAdjustScheduleItem_NoCascadeAndStateGuard(t *testing.T) {
	gdb, _ := testutil.SetupTestDB()
	base := testBase()

	plan := &model.ProductionSchedulePlan{PlanNo: "SC-002", StartTime: base, EndTime: base.Add(20 * time.Minute), OrderCount: 2, CurrentState: ScheduleStateGenerated}
	gdb.Create(plan)
	itemA := &model.ProductionScheduleItem{ProductionSchedulePlanID: plan.ID, Sequence: 1, PlannedStartTime: base, PlannedEndTime: base.Add(10 * time.Minute)}
	itemB := &model.ProductionScheduleItem{ProductionSchedulePlanID: plan.ID, Sequence: 2, PlannedStartTime: base.Add(10 * time.Minute), PlannedEndTime: base.Add(20 * time.Minute)}
	gdb.Create([]*model.ProductionScheduleItem{itemA, itemB})

	// 非级联：仅A移动
	if err := AdjustScheduleItem(&proto.AdjustScheduleItemRequest{
		PlanID: plan.ID, ItemID: itemA.ID,
		NewStartTime: base.Add(5 * time.Minute).Format("2006-01-02 15:04:05"), Cascade: false,
	}, "u"); err != nil {
		t.Fatalf("调整失败: %v", err)
	}
	var after []*model.ProductionScheduleItem
	gdb.Where("`production_schedule_plan_id` = ?", plan.ID).Order("sequence").Find(&after)
	if !after[0].PlannedStartTime.Equal(base.Add(5 * time.Minute)) {
		t.Fatalf("目标明细应移动")
	}
	if !after[1].PlannedStartTime.Equal(base.Add(10 * time.Minute)) {
		t.Fatalf("非级联模式后续明细不应移动")
	}

	// 已下发拒绝
	gdb.Model(&model.ProductionSchedulePlan{}).Where("`id` = ?", plan.ID).Update("current_state", ScheduleStateReleased)
	if err := AdjustScheduleItem(&proto.AdjustScheduleItemRequest{
		PlanID: plan.ID, ItemID: itemA.ID,
		NewStartTime: base.Add(6 * time.Minute).Format("2006-01-02 15:04:05"),
	}, "u"); err == nil {
		t.Fatalf("已下发计划应拒绝人工调整")
	}
}

// 下发幂等：重复下发第二次被原子门拒绝
func TestReleaseSchedule_TwiceRejected(t *testing.T) {
	gdb, _ := testutil.SetupTestDB()
	base := testBase()
	line := &model.ProductionLine{Code: "L-REL"}
	gdb.Create([]*model.ProductionLine{line})
	rhythm := &model.ProductionRhythm{Priority: 1, Enable: true, StandardTime: 60, ProductionLineID: line.ID, InitialValue: true}
	gdb.Create([]*model.ProductionRhythm{rhythm})
	process := &model.ProductionProcess{Code: "P-REL", SortIndex: 1, Enable: true, InitialValue: true}
	gdb.Create([]*model.ProductionProcess{process})
	order := &model.ProductOrder{ProductOrderNo: "WO-REL2", OrderQTY: 2, CurrentState: "已签派", ProductionLineID: &line.ID, StandardWorkTime: 60}
	gdb.Create(order)

	resp, err := GenerateSchedule(&proto.GenerateScheduleRequest{ProductionLineID: line.ID, StartTime: base.Add(time.Hour).Format("2006-01-02 15:04:05")}, "u")
	if err != nil {
		t.Fatalf("生成失败: %v", err)
	}
	if err := ReleaseSchedule(&proto.ReleaseScheduleRequest{PlanID: resp.PlanID}, "u"); err != nil {
		t.Fatalf("首次下发失败: %v", err)
	}
	err = ReleaseSchedule(&proto.ReleaseScheduleRequest{PlanID: resp.PlanID}, "u")
	if err == nil {
		t.Fatalf("重复下发应被拒绝")
	}
}
