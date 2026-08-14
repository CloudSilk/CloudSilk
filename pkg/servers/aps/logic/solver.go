package logic

import (
	"time"
)

// SchedOptions 排程约束选项
type SchedOptions struct {
	// 每日工作分钟数（<=0 表示连续排程不跨天中断）
	DailyWorkMinutes int32
	// 换型时间（秒）：相邻工单产品型号不同时插入的换型间隔
	ChangeoverSec int32
}

// SchedOrderEx 带约束信息的排程工单
type SchedOrderEx struct {
	SchedOrder
	// 产品型号（用于换型约束判断）
	ProductModel string
	// 物料齐套时间（可选）：工单不得早于该时间开工
	ReadyTime time.Time
	HasReady  bool
}

// ConstrainedSchedule 带约束的前向排程：
//  1. 排序：优先级降序 → 交期升序（无交期排后）；同优先级内按产品型号聚簇以减少换型（稳定，不跨优先级合并）
//  2. 约束：
//     - 物料齐套：开工时间不早于 ReadyTime
//     - 换型时间：与前一个工单产品型号不同时，插入 ChangeoverSec 的换型间隔
//     - 每日工作窗口：跨天顺延（见 fitToDailyWindow）
func ConstrainedSchedule(orders []SchedOrderEx, startTime time.Time, opts SchedOptions) []SchedSlot {
	sorted := make([]SchedOrderEx, len(orders))
	copy(sorted, orders)
	sortSchedOrders(sorted)

	slots := make([]SchedSlot, 0, len(sorted))
	var cursor time.Time = startTime
	lastModel := ""
	for _, o := range sorted {
		seconds := int64(o.QTY) * int64(o.StdWorkTimeSec)
		if seconds <= 0 {
			continue
		}

		//物料齐套约束：不早于齐套时间开工
		begin := cursor
		if o.HasReady && o.ReadyTime.After(begin) {
			begin = o.ReadyTime
		}

		//换型约束：型号切换插入换型间隔
		if opts.ChangeoverSec > 0 && lastModel != "" && o.ProductModel != lastModel {
			begin = begin.Add(time.Duration(opts.ChangeoverSec) * time.Second)
		}

		end := begin.Add(time.Duration(seconds) * time.Second)
		if opts.DailyWorkMinutes > 0 {
			begin, end = fitToDailyWindow(begin, seconds, startTime, opts.DailyWorkMinutes)
		}

		slots = append(slots, SchedSlot{
			OrderID: o.OrderID,
			OrderNo: o.OrderNo,
			Start:   begin,
			End:     end,
			Seconds: seconds,
		})
		cursor = end
		lastModel = o.ProductModel
	}
	return slots
}

// InsertOrder 插单重排：将新工单插入既有排程，
// 插入位置为最后一个"优先级严格更高"的工单之后（保持既有工单相对顺序），
// 从插入点起重新应用齐套/换型/日窗口约束，后续工单顺延。
// 返回新时段列表与新工单所在下标。
func InsertOrder(existing []SchedSlot, orders []SchedOrderEx, newOrder SchedOrderEx, startTime time.Time, opts SchedOptions) ([]SchedSlot, int) {
	//按 ID 找到既有工单的约束信息
	byID := map[string]SchedOrderEx{}
	for _, o := range orders {
		byID[o.OrderID] = o
	}

	//确定插入下标：最后一个优先级严格更高的工单之后
	insertAt := len(existing)
	for i, slot := range existing {
		if o, ok := byID[slot.OrderID]; ok && o.PriorityLevel > newOrder.PriorityLevel {
			insertAt = i + 1
		}
	}

	//重组工单序列（保持既有顺序），插入新工单
	sequence := make([]SchedOrderEx, 0, len(existing)+1)
	for i, slot := range existing {
		if i == insertAt {
			sequence = append(sequence, newOrder)
		}
		if o, ok := byID[slot.OrderID]; ok {
			sequence = append(sequence, o)
		}
	}
	if insertAt >= len(existing) {
		sequence = append(sequence, newOrder)
	}

	slots := ConstrainedSchedule(sequence, startTime, opts)
	for i, slot := range slots {
		if slot.OrderID == newOrder.OrderID {
			return slots, i
		}
	}
	return slots, -1
}

// sortSchedOrders 排序规则：优先级降序 → 交期升序（无交期排后）→ 同优先级内产品型号聚簇
func sortSchedOrders(sorted []SchedOrderEx) {
	//第一遍：按优先级/交期稳定排序（复用既有语义）
	stableSortByPriorityDelivery(sorted)

	//第二遍：同优先级分区内按产品型号聚簇（稳定，减少换型次数）
	partitionSortByModel(sorted)
}

func stableSortByPriorityDelivery(sorted []SchedOrderEx) {
	//插入排序保证稳定性且避免引入新依赖
	for i := 1; i < len(sorted); i++ {
		for j := i; j > 0; j-- {
			a, b := sorted[j-1], sorted[j]
			swap := false
			if a.PriorityLevel != b.PriorityLevel {
				swap = a.PriorityLevel < b.PriorityLevel
			} else if a.HasDelivery != b.HasDelivery {
				swap = b.HasDelivery // 有交期者在前
			} else if a.HasDelivery && !a.DeliveryDate.Equal(b.DeliveryDate) {
				swap = a.DeliveryDate.After(b.DeliveryDate)
			}
			if !swap {
				break
			}
			sorted[j-1], sorted[j] = sorted[j], sorted[j-1]
		}
	}
}

// partitionSortByModel 在同一优先级分区内，按产品型号对相邻工单做稳定聚簇
// （聚簇仅移动同分区工单，不影响跨分区顺序）
func partitionSortByModel(sorted []SchedOrderEx) {
	start := 0
	for start < len(sorted) {
		end := start + 1
		for end < len(sorted) && sorted[end].PriorityLevel == sorted[start].PriorityLevel {
			end++
		}
		// 分区 [start, end)：把相同型号聚到一起，保持各型号组内原有顺序
		grouped := make([]SchedOrderEx, 0, end-start)
		emitted := make([]bool, end-start)
		for i := start; i < end; i++ {
			if emitted[i-start] {
				continue
			}
			model := sorted[i].ProductModel
			for j := i; j < end; j++ {
				if !emitted[j-start] && sorted[j].ProductModel == model {
					grouped = append(grouped, sorted[j])
					emitted[j-start] = true
				}
			}
		}
		copy(sorted[start:end], grouped)
		start = end
	}
}
