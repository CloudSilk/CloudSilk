package logic

import (
	"testing"
	"time"
)

func testBase() time.Time {
	return time.Date(2026, 8, 13, 8, 0, 0, 0, time.Local)
}

// 连续排程：按时长顺序排布，无间隔无重叠
func TestForwardSchedule_Continuous(t *testing.T) {
	base := testBase()
	orders := []SchedOrder{
		{OrderID: "a", OrderNo: "WO-A", QTY: 10, StdWorkTimeSec: 60},
		{OrderID: "b", OrderNo: "WO-B", QTY: 30, StdWorkTimeSec: 60},
	}
	slots := ForwardSchedule(orders, base, 0)
	if len(slots) != 2 {
		t.Fatalf("应生成2个时段，实际%d", len(slots))
	}
	if !slots[0].Start.Equal(base) {
		t.Fatalf("首时段应从基准时间开始")
	}
	if !slots[0].End.Equal(base.Add(10 * time.Minute)) {
		t.Fatalf("首时段应持续10分钟")
	}
	if !slots[1].Start.Equal(slots[0].End) {
		t.Fatalf("第二时段应紧接首时段")
	}
	if !slots[1].End.Equal(base.Add(40 * time.Minute)) {
		t.Fatalf("第二时段应结束于40分钟处")
	}
}

// 优先级排序：高优先级先排
func TestForwardSchedule_Priority(t *testing.T) {
	base := testBase()
	orders := []SchedOrder{
		{OrderID: "low", OrderNo: "WO-LOW", QTY: 5, StdWorkTimeSec: 60, PriorityLevel: 1},
		{OrderID: "high", OrderNo: "WO-HIGH", QTY: 5, StdWorkTimeSec: 60, PriorityLevel: 9},
	}
	slots := ForwardSchedule(orders, base, 0)
	if slots[0].OrderID != "high" {
		t.Fatalf("高优先级工单应先排，实际首时段为%s", slots[0].OrderID)
	}
}

// 同优先级按交期升序，无交期排后
func TestForwardSchedule_DeliveryDate(t *testing.T) {
	base := testBase()
	early := base.AddDate(0, 0, 2)
	late := base.AddDate(0, 0, 10)
	orders := []SchedOrder{
		{OrderID: "late", QTY: 1, StdWorkTimeSec: 60, DeliveryDate: late, HasDelivery: true},
		{OrderID: "none", QTY: 1, StdWorkTimeSec: 60},
		{OrderID: "early", QTY: 1, StdWorkTimeSec: 60, DeliveryDate: early, HasDelivery: true},
	}
	slots := ForwardSchedule(orders, base, 0)
	if slots[0].OrderID != "early" || slots[1].OrderID != "late" || slots[2].OrderID != "none" {
		t.Fatalf("顺序应为 early→late→none，实际 %s→%s→%s", slots[0].OrderID, slots[1].OrderID, slots[2].OrderID)
	}
}

// 按日窗口排程：跨天工时顺延到次日窗口
func TestForwardSchedule_DailyWindow(t *testing.T) {
	base := testBase()
	// 每天480分钟（8小时），工单需600分钟 → 当天480 + 次日120
	orders := []SchedOrder{
		{OrderID: "x", QTY: 1, StdWorkTimeSec: 600 * 60},
	}
	slots := ForwardSchedule(orders, base, 480)
	if len(slots) != 1 {
		t.Fatalf("应生成1个时段")
	}
	slot := slots[0]
	// 基准时间8点不在当天0点窗口内，窗口起点为基准时间本身
	if !slot.Start.Equal(base) {
		t.Fatalf("跨天工单应从基准时间开始")
	}
	// 当天窗口剩余 = 16:00(当天0点+480分钟=8:00？) —— 基准8点即当天0点+480分钟窗口末端，应滚动到次日
	nextDayStart := time.Date(2026, 8, 14, 0, 0, 0, 0, time.Local).Add(0)
	_ = nextDayStart
	if !slot.End.After(slot.Start) {
		t.Fatalf("结束时间应晚于开始时间")
	}
	// 总工时600分钟：当天用满480分钟，剩余120分钟顺延至次日窗口起点（08:00）之后
	expectEnd := time.Date(2026, 8, 14, 8, 0, 0, 0, time.Local).Add(120 * time.Minute)
	if !slot.End.Equal(expectEnd) {
		t.Fatalf("跨天工单应于次日窗口内120分钟处结束，期望%v，实际%v", expectEnd, slot.End)
	}
}

// 零工时工单被跳过
func TestForwardSchedule_SkipZero(t *testing.T) {
	base := testBase()
	orders := []SchedOrder{
		{OrderID: "zero", QTY: 0, StdWorkTimeSec: 60},
		{OrderID: "ok", QTY: 2, StdWorkTimeSec: 30},
	}
	slots := ForwardSchedule(orders, base, 0)
	if len(slots) != 1 || slots[0].OrderID != "ok" {
		t.Fatalf("零工时工单应被跳过")
	}
}
