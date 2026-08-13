package logic

import (
	"fmt"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	equipment "github.com/CloudSilk/CloudSilk/pkg/servers/equipment/logic"
	"github.com/CloudSilk/CloudSilk/pkg/types"
)

// GetOverview 厂级总览：产线/工站规模、工单进度、告警与工站状态分布
func GetOverview() (*proto.MonitoringOverviewResponse, error) {
	resp := &proto.MonitoringOverviewResponse{}

	var lineCount, stationCount, producingOrders, unhandledAlarms, todayAlarms int64
	var idleStations, producingStations, breakdownStations int64
	type qtyAgg struct {
		OrderQTY    int64
		FinishedQTY int64
		StartedQTY  int64
	}
	orders := qtyAgg{}

	if err := model.DB.DB().Model(&model.ProductionLine{}).Count(&lineCount).Error; err != nil {
		return nil, err
	}
	if err := model.DB.DB().Model(&model.ProductionStation{}).Count(&stationCount).Error; err != nil {
		return nil, err
	}
	if err := model.DB.DB().Model(&model.ProductOrder{}).Where("`current_state` = ?", types.ProductOrderStateProducting).Count(&producingOrders).Error; err != nil {
		return nil, err
	}
	if err := model.DB.DB().Model(&model.ProductOrder{}).Where("`current_state` in ?", []string{
		types.ProductOrderStateReleased, types.ProductOrderStateDispatched, types.ProductOrderStateVerified,
		types.ProductOrderStateReceipted, types.ProductOrderStateProducting,
	}).Select("COALESCE(SUM(order_qty),0) as order_qty, COALESCE(SUM(finished_qty),0) as finished_qty, COALESCE(SUM(started_qty),0) as started_qty").Scan(&orders).Error; err != nil {
		return nil, err
	}
	if err := model.DB.DB().Model(&model.ProductionStationAlarm{}).Where("`current_state` = ?", "未处理").Count(&unhandledAlarms).Error; err != nil {
		return nil, err
	}
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
	if err := model.DB.DB().Model(&model.ProductionStationAlarm{}).Where("`create_time` >= ?", today).Count(&todayAlarms).Error; err != nil {
		return nil, err
	}
	if err := model.DB.DB().Model(&model.ProductionStation{}).Where("`current_state` = ?", types.ProductionStationStateStandby).Count(&idleStations).Error; err != nil {
		return nil, err
	}
	if err := model.DB.DB().Model(&model.ProductionStation{}).Where("`current_state` = ?", types.ProductionStationStateOccupied).Count(&producingStations).Error; err != nil {
		return nil, err
	}
	if err := model.DB.DB().Model(&model.ProductionStation{}).Where("`current_state` = ?", types.ProductionStationStateBreakdown).Count(&breakdownStations).Error; err != nil {
		return nil, err
	}

	resp.Code = proto.Code_Success
	resp.ProductionLineCount = int32(lineCount)
	resp.ProductionStationCount = int32(stationCount)
	resp.ProducingOrderCount = int32(producingOrders)
	resp.TotalOrderQTY = orders.OrderQTY
	resp.TotalFinishedQTY = orders.FinishedQTY
	resp.TotalStartedQTY = orders.StartedQTY
	resp.UnhandledAlarmCount = int32(unhandledAlarms)
	resp.TodayAlarmCount = int32(todayAlarms)
	resp.IdleStationCount = int32(idleStations)
	resp.ProducingStationCount = int32(producingStations)
	resp.BreakdownStationCount = int32(breakdownStations)
	return resp, nil
}

// GetLineMonitoring 产线级监控：OEE、产量、达成率、在制、告警
func GetLineMonitoring(productionLineID, startTime, endTime string) (*proto.MonitoringLineResponse, error) {
	if productionLineID == "" {
		return nil, fmt.Errorf("productionLineID不能为空")
	}
	const layout = "2006-01-02 15:04:05"
	var st, et time.Time
	var err error
	if startTime != "" {
		if st, err = time.ParseInLocation(layout, startTime, time.Local); err != nil {
			return nil, fmt.Errorf("startTime格式无效")
		}
	} else {
		st = time.Now().AddDate(0, 0, -1)
	}
	if endTime != "" {
		if et, err = time.ParseInLocation(layout, endTime, time.Local); err != nil {
			return nil, fmt.Errorf("endTime格式无效")
		}
	} else {
		et = time.Now()
	}

	line := &model.ProductionLine{}
	if err := model.DB.DB().First(line, "`id` = ?", productionLineID).Error; err != nil {
		return nil, fmt.Errorf("读取产线失败")
	}

	resp := &proto.MonitoringLineResponse{
		Code:                     proto.Code_Success,
		ProductionLineID:         line.ID,
		ProductionLineCode:       line.Code,
		ProductionLineDescription: line.Description,
	}

	//工站数与告警
	var stationCount int64
	if err := model.DB.DB().Model(&model.ProductionStation{}).Where("`production_line_id` = ?", line.ID).Count(&stationCount).Error; err != nil {
		return nil, err
	}
	resp.StationCount = int32(stationCount)
	var alarmCount int64
	model.DB.DB().Model(&model.ProductionStationAlarm{}).
		Joins("JOIN production_stations ON production_station_alarms.production_station_id=production_stations.id").
		Where("production_stations.production_line_id = ? AND production_station_alarms.current_state = ?", line.ID, "未处理").Count(&alarmCount)
	resp.UnhandledAlarmCount = int32(alarmCount)

	//周期产量：该产线工站在时间段内的节拍记录数
	var outputCount int64
	model.DB.DB().Model(&model.ProductRhythmRecord{}).
		Joins("JOIN production_stations ON product_rhythm_records.production_station_id=production_stations.id").
		Where("production_stations.production_line_id = ? AND product_rhythm_records.work_start_time >= ? AND product_rhythm_records.work_start_time <= ?", line.ID, st, et).
		Count(&outputCount)
	resp.OutputCount = outputCount

	//工单达成率与在制
	type orderAgg struct {
		OrderQTY    int64
		FinishedQTY int64
	}
	orders := orderAgg{}
	model.DB.DB().Model(&model.ProductOrder{}).Where("`production_line_id` = ? AND `current_state` in ?", line.ID,
		[]string{types.ProductOrderStateReleased, types.ProductOrderStateDispatched, types.ProductOrderStateProducting}).
		Select("COALESCE(SUM(order_qty),0) as order_qty, COALESCE(SUM(finished_qty),0) as finished_qty").Scan(&orders)
	if orders.OrderQTY > 0 {
		resp.AchievementRate = float64(orders.FinishedQTY) / float64(orders.OrderQTY)
	}

	var wipCount int64
	model.DB.DB().Model(&model.ProductProcessRoute{}).
		Joins("JOIN production_processes ON product_process_routes.current_process_id=production_processes.id").
		Where("production_processes.production_line_id = ? AND product_process_routes.current_state in ?", line.ID,
			[]string{types.ProductProcessRouteStateWaitProcess, types.ProductProcessRouteStateProcessing}).
		Count(&wipCount)
	resp.WipCount = wipCount

	//OEE：取产线首个工站计算（产线级 OEE 汇总为其工站加权平均的增强项）
	var firstStation model.ProductionStation
	if err := model.DB.DB().Where("`production_line_id` = ?", line.ID).Order("sort_index").First(&firstStation).Error; err == nil {
		oee, err := equipment.CalcOee(&proto.EquipmentOeeRequest{
			ProductionStationID: firstStation.ID,
			StartTime:           st.Format(layout),
			EndTime:             et.Format(layout),
		})
		if err == nil {
			resp.Oee = oee.Oee
			resp.Availability = oee.Availability
			resp.Performance = oee.Performance
			resp.Quality = oee.Quality
		}
	}

	return resp, nil
}

// QueryAlarms 告警查询（大屏右侧滚动列表）
func QueryAlarms(req *proto.QueryMonitoringAlarmRequest) (*proto.QueryMonitoringAlarmResponse, error) {
	resp := &proto.QueryMonitoringAlarmResponse{Code: proto.Code_Success}
	limit := req.Limit
	if limit <= 0 || limit > 200 {
		limit = 50
	}

	db := model.DB.DB().Model(&model.ProductionStationAlarm{}).Preload("ProductionStation")
	if req.CurrentState != "" {
		db = db.Where("production_station_alarms.`current_state` = ?", req.CurrentState)
	}
	db = db.Order("create_time desc")

	var list []*model.ProductionStationAlarm
	if err := db.Limit(int(limit)).Find(&list).Error; err != nil {
		return nil, err
	}

	for _, a := range list {
		item := &proto.MonitoringAlarmInfo{
			Id:           a.ID,
			AlarmNo:      a.AlarmNo,
			AlarmMessage: a.AlarmMessage,
			CurrentState: a.CurrentState,
			HandleMethod: a.HandleMethod,
			CreateTime:   a.CreateTime.Format("2006-01-02 15:04:05"),
		}
		if a.ProductionStation != nil {
			item.ProductionStationCode = a.ProductionStation.Code
		}
		resp.Data = append(resp.Data, item)
	}
	resp.Total = int64(len(list))
	return resp, nil
}
