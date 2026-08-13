package types

const (
	AGVTaskQueueStateWaitDispatch = "待签派"
	AGVTaskQueueStateDispatched   = "已签派"
	AGVTaskQueueStateExecuting    = "执行中"
	AGVTaskQueueStateWaitRelease  = "待放行"
	AGVTaskQueueStateTransporting = "运送中"
	AGVTaskQueueStateCompleted    = "已完成"
	AGVTaskQueueStateCancelled    = "已取消"
)
