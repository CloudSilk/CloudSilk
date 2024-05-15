package types

const (
	// 产品封箱
	SystemEventProductInfoPackaged = "ProductInfoPackaged"
	// 产品配料
	SystemEventProductInfoIssued = "ProductInfoIssued"
	// 产品上线
	SystemEventProductInfoOnlined = "ProductInfoOnlined"
	// 产品测试记录创建
	SystemEventProductInfoTested = "ProductInfoTested"
	// 产品返工记录更新
	SystemEventProductInfoReworked = "ProductInfoReworked"
	// 箱单打印
	SystemEventPackageLabelPrinted = "PackageLabelPrinted"
	// 工站开机
	SystemEventProductionStationStartup = "ProductionStationStartup"
	// 工站停机
	SystemEventProductionStationShutdown = "ProductionStationShutdown"
	// 工站报警
	SystemEventProductionStationAlarmed = "ProductionStationAlarmed"
	// 工站报警恢复
	SystemEventProductionStationRecovered = "ProductionStationRecovered"
	// 工站故障
	SystemEventProductionStationBreakdown = "ProductionStationBreakdown"
	// 工站作业
	SystemEventProductionStationOccupied = "ProductionStationOccupied"
	// 工站待机(清空，释放)
	SystemEventProductionStationReleased = "ProductionStationReleased"
	// 工序开始
	SystemEventProductionProcessStarted = "ProductionProcessStarted"
	// 工序结束
	SystemEventProductionProcessFinished = "ProductionProcessFinished"
	// 工单开工
	SystemEventProductOrderStarted = "ProductOrderStarted"
	// 工单完工
	SystemEventProductOrderFinished = "ProductOrderFinished"
	// 工线班次切换
	SystemEventProductionLineShifted = "ProductionLineShifted"
	// 创建消息发送队列
	SystemEventMessageSendQueueCreated = "MessageSendQueueCreated"
)
