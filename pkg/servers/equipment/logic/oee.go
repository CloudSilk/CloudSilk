package logic

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// CalcOee 计算设备/工站在时间范围内的 OEE
//
// OEE = 时间稼动率 × 性能稼动率 × 良品率
//   - 时间稼动率 = 运行时间 / 计划时间（停机时间来自工站故障记录 ProductionStationBreakdown.Duration，单位分钟）
//   - 性能稼动率 = 理论加工时间 / 运行时间（理论加工时间 = 加工数量 × 平均标准节拍，节拍数据来自 ProductRhythmRecord）
//   - 良品率 = 合格数 / 检测总数（检测数据来自 ProductTestRecord；无检测数据时按返工记录折算）
func CalcOee(req *proto.EquipmentOeeRequest) (*proto.EquipmentOeeResponse, error) {
	if req.ProductionStationID == "" && req.EquipmentID == "" {
		return nil, fmt.Errorf("productionStationID与equipmentID至少提供一个")
	}
	if req.EquipmentID != "" && req.ProductionStationID == "" {
		equipment := &model.Equipment{}
		if err := model.DB.DB().First(equipment, "`id` = ?", req.EquipmentID).Error; err != nil {
			return nil, fmt.Errorf("读取设备失败")
		}
		if equipment.ProductionStationID == nil {
			return nil, fmt.Errorf("此设备未关联生产工站，无法按节拍数据计算OEE")
		}
		req.ProductionStationID = *equipment.ProductionStationID
	}

	const layout = "2006-01-02 15:04:05"
	startTime, err := time.ParseInLocation(layout, req.StartTime, time.Local)
	if err != nil {
		return nil, fmt.Errorf("startTime格式无效，应为%s", layout)
	}
	endTime, err := time.ParseInLocation(layout, req.EndTime, time.Local)
	if err != nil {
		return nil, fmt.Errorf("endTime格式无效，应为%s", layout)
	}
	if !endTime.After(startTime) {
		return nil, fmt.Errorf("endTime必须晚于startTime")
	}

	resp := &proto.EquipmentOeeResponse{Code: proto.Code_Success}

	//计划时间（分钟）
	plannedMinutes := req.PlannedMinutes
	if plannedMinutes <= 0 {
		plannedMinutes = int32(endTime.Sub(startTime).Minutes())
	}
	resp.PlannedMinutes = float64(plannedMinutes)

	//停机时间：工站故障记录（Duration 单位分钟），未完成的不计入
	var downtimeSum sql.NullFloat64
	if err := model.DB.DB().Model(&model.ProductionStationBreakdown{}).
		Select("COALESCE(SUM(duration),0)").
		Where("`production_station_id` = ? AND `create_time` >= ? AND `create_time` <= ? AND `complete_time` IS NOT NULL",
			req.ProductionStationID, startTime, endTime).
		Scan(&downtimeSum).Error; err != nil {
		return nil, err
	}
	downtime := downtimeSum.Float64
	if downtime > float64(plannedMinutes) {
		downtime = float64(plannedMinutes)
	}
	resp.DowntimeMinutes = downtime

	runMinutes := float64(plannedMinutes) - downtime
	resp.RunMinutes = runMinutes

	//加工数量与平均标准节拍（秒）
	type rhythmAgg struct {
		Count     int64
		AvgStdSec float64
	}
	agg := rhythmAgg{}
	if err := model.DB.DB().Model(&model.ProductRhythmRecord{}).
		Select("COUNT(*) as count, COALESCE(AVG(standard_work_time),0) as avg_std_sec").
		Where("`production_station_id` = ? AND `work_start_time` >= ? AND `work_start_time` <= ?",
			req.ProductionStationID, startTime, endTime).
		Scan(&agg).Error; err != nil {
		return nil, err
	}
	resp.TotalCount = agg.Count

	//合格数量：优先测试记录，无测试记录时用总数-返工数折算
	var qualified, tested int64
	if err := model.DB.DB().Model(&model.ProductTestRecord{}).
		Where("`production_station_id` = ? AND `test_start_time` >= ? AND `test_end_time` <= ? AND `is_qualified` = ?",
			req.ProductionStationID, startTime, endTime, true).
		Count(&tested).Error; err != nil {
		return nil, err
	}
	var testTotal int64
	if err := model.DB.DB().Model(&model.ProductTestRecord{}).
		Where("`production_station_id` = ? AND `test_start_time` >= ? AND `test_end_time` <= ?",
			req.ProductionStationID, startTime, endTime).
		Count(&testTotal).Error; err != nil {
		return nil, err
	}

	if testTotal > 0 {
		qualified = tested
		resp.QualifiedCount = qualified
		resp.Quality = round4(float64(qualified) / float64(testTotal))
	} else if agg.Count > 0 {
		var reworks int64
		if err := model.DB.DB().Model(&model.ProductReworkRecord{}).
			Where("`production_station_id` = ? AND `rework_time` >= ? AND `rework_time` <= ?",
				req.ProductionStationID, startTime, endTime).
			Count(&reworks).Error; err != nil {
			return nil, err
		}
		qualified = agg.Count - reworks
		if qualified < 0 {
			qualified = 0
		}
		resp.QualifiedCount = qualified
		resp.Quality = round4(float64(qualified) / float64(agg.Count))
	} else {
		resp.Quality = 0
	}

	//时间稼动率
	if plannedMinutes > 0 {
		resp.Availability = round4(runMinutes / float64(plannedMinutes))
	} else {
		resp.Availability = 0
	}

	//性能稼动率：理论加工时间(分钟) = 数量 × 平均节拍(秒) / 60
	idealMinutes := float64(agg.Count) * agg.AvgStdSec / 60.0
	if runMinutes > 0 && idealMinutes > 0 {
		performance := idealMinutes / runMinutes
		if performance > 1 {
			performance = 1
		}
		resp.Performance = round4(performance)
	} else {
		resp.Performance = 0
	}

	resp.Oee = round4(resp.Availability * resp.Performance * resp.Quality)
	return resp, nil
}

func round4(v float64) float64 {
	return float64(int(v*10000+0.5)) / 10000
}
