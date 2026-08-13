package logic

import (
	"fmt"
	"sort"
	"time"
)

// SchedOrder 排程输入工单
type SchedOrder struct {
	OrderID        string
	OrderNo        string
	QTY            int32
	StdWorkTimeSec int32
	PriorityLevel  int32
	DeliveryDate   time.Time
	HasDelivery    bool
}

// SchedSlot 排程输出时段
type SchedSlot struct {
	OrderID string
	OrderNo string
	Start   time.Time
	End     time.Time
	Seconds int64
}

// ForwardSchedule 前向贪心排程算法：
//  1. 排序：优先级降序 → 交期升序（无交期排后）
//  2. 从 startTime 起依次为每个工单分配连续时段，工时 = 数量 × 标准节拍
//  3. dailyWorkMinutes > 0 时按天窗口排程：每天只排该分钟数，跨天自动跳到次日窗口起点
//
// 算法为纯函数，便于单测与后续替换为约束求解器。
func ForwardSchedule(orders []SchedOrder, startTime time.Time, dailyWorkMinutes int32) []SchedSlot {
	sorted := make([]SchedOrder, len(orders))
	copy(sorted, orders)
	sort.SliceStable(sorted, func(i, j int) bool {
		if sorted[i].PriorityLevel != sorted[j].PriorityLevel {
			return sorted[i].PriorityLevel > sorted[j].PriorityLevel
		}
		// 有交期者在前，同按交期升序
		if sorted[i].HasDelivery != sorted[j].HasDelivery {
			return sorted[i].HasDelivery
		}
		if sorted[i].HasDelivery && sorted[i].DeliveryDate.Equal(sorted[j].DeliveryDate) == false {
			return sorted[i].DeliveryDate.Before(sorted[j].DeliveryDate)
		}
		return false // 稳定排序保留原始顺序
	})

	slots := make([]SchedSlot, 0, len(sorted))
	cursor := startTime
	for _, o := range sorted {
		seconds := int64(o.QTY) * int64(o.StdWorkTimeSec)
		if seconds <= 0 {
			continue
		}
		start, end := cursor, cursor.Add(time.Duration(seconds)*time.Second)
		if dailyWorkMinutes > 0 {
			start, end = fitToDailyWindow(cursor, seconds, startTime, dailyWorkMinutes)
		}
		slots = append(slots, SchedSlot{
			OrderID: o.OrderID,
			OrderNo: o.OrderNo,
			Start:   start,
			End:     end,
			Seconds: seconds,
		})
		if dailyWorkMinutes > 0 {
			cursor = end
		} else {
			cursor = end
		}
	}
	return slots
}

// fitToDailyWindow 将 [cursor, cursor+seconds] 的工时切分到按日窗口：
// 每天（自排程基准日0点起算）允许工作 dailyWorkMinutes 分钟，超出部分顺延到次日窗口。
func fitToDailyWindow(cursor time.Time, seconds int64, base time.Time, dailyWorkMinutes int32) (time.Time, time.Time) {
	window := time.Duration(dailyWorkMinutes) * time.Minute
	// 每日窗口起点：当天 0 点（如需班次日历可在此扩展）
	dayStart := time.Date(cursor.Year(), cursor.Month(), cursor.Day(), 0, 0, 0, 0, cursor.Location())
	if dayStart.Before(base) {
		dayStart = base
	}
	dayEnd := dayStart.Add(window)

	remaining := time.Duration(seconds) * time.Second
	start := cursor
	// 若当前时刻已越过当日窗口末端，滚动到次日
	for start.After(dayEnd) || start.Equal(dayEnd) {
		dayStart = dayStart.AddDate(0, 0, 1)
		dayEnd = dayStart.Add(window)
		start = dayStart
	}
	if start.Add(remaining).Before(dayEnd) || start.Add(remaining).Equal(dayEnd) {
		return start, start.Add(remaining)
	}
	// 需要跨天：先耗尽当天剩余
	used := dayEnd.Sub(start)
	remaining -= used
	end := dayEnd
	for remaining > 0 {
		dayStart = dayStart.AddDate(0, 0, 1)
		dayEnd = dayStart.Add(window)
		if remaining <= window {
			end = dayStart.Add(remaining)
			remaining = 0
		} else {
			remaining -= window
		}
	}
	return start, end
}

// formatSlotDesc 输出可读的时段描述（调试/日志用）
func formatSlotDesc(slots []SchedSlot) string {
	s := ""
	for i, slot := range slots {
		s += fmt.Sprintf("%d.%s[%s→%s] ", i+1, slot.OrderNo, slot.Start.Format("01-02 15:04"), slot.End.Format("01-02 15:04"))
	}
	return s
}
