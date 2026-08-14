package logic

import (
	"sync"
	"time"
)

// Scheduler 排程算法接口：输入带约束的工单集合，输出时段安排。
// 内置实现为 GreedyScheduler（前向贪心：优先级→交期→型号聚簇，
// 含换型/齐套/日窗口约束）。接口存在的意义是把"策略"与"编排"解耦：
// 后续引入约束求解器（如 OR-Tools CP-SAT、遗传算法）时，
// 实现 Scheduler 并 SetScheduler 注册即可，编排层（生成/插单/下发）零改动。
type Scheduler interface {
	Schedule(orders []SchedOrderEx, startTime time.Time, opts SchedOptions) []SchedSlot
}

// GreedyScheduler 内置前向贪心实现（纯函数，可并发复用）
type GreedyScheduler struct{}

func (GreedyScheduler) Schedule(orders []SchedOrderEx, startTime time.Time, opts SchedOptions) []SchedSlot {
	return ConstrainedSchedule(orders, startTime, opts)
}

var (
	schedulerMu  sync.RWMutex
	schedulerImpl Scheduler = GreedyScheduler{}
)

// SetScheduler 注册自定义排程算法（并发安全）
func SetScheduler(s Scheduler) {
	schedulerMu.Lock()
	defer schedulerMu.Unlock()
	if s != nil {
		schedulerImpl = s
	}
}

// CurrentScheduler 返回当前生效的排程算法
func CurrentScheduler() Scheduler {
	schedulerMu.RLock()
	defer schedulerMu.RUnlock()
	return schedulerImpl
}
