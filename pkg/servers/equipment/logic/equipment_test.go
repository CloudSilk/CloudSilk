package logic

import (
	"database/sql"
	"testing"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/testutil"
	"gorm.io/gorm"
)

func setupEMDB(t *testing.T) *gorm.DB {
	t.Helper()
	gdb, err := testutil.SetupTestDB()
	if err != nil {
		t.Fatalf("初始化测试数据库失败: %v", err)
	}
	return gdb
}

// 维保周期变更：按新周期重算下次执行（基准=上次执行时间）
func TestUpdateMaintenancePlan_RecalcNextExecution(t *testing.T) {
	gdb := setupEMDB(t)
	eq := &model.Equipment{Code: "EQ-1", CurrentState: EquipmentStateInUse}
	if err := gdb.Create(eq).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	lastExec := time.Now().AddDate(0, 0, -10)
	plan := &model.EquipmentMaintenancePlan{
		EquipmentID:     &eq.ID,
		MaintenanceType: "保养",
		CycleDays:       30,
		LastExecution:   sql.NullTime{Time: lastExec, Valid: true},
		NextExecution:   sql.NullTime{Time: lastExec.AddDate(0, 0, 30), Valid: true},
		Enable:          true,
	}
	if err := gdb.Create(plan).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}

	plan.CycleDays = 15
	if err := UpdateEquipmentMaintenancePlan(plan); err != nil {
		t.Fatalf("更新失败: %v", err)
	}
	after := &model.EquipmentMaintenancePlan{}
	gdb.First(after, "`id` = ?", plan.ID)
	expect := lastExec.AddDate(0, 0, 15)
	if after.NextExecution.Time.Format("2006-01-02") != expect.Format("2006-01-02") {
		t.Fatalf("周期改15天后下次执行应为%v，实际%v", expect, after.NextExecution.Time)
	}

	// 非法周期拒绝
	after.CycleDays = 0
	if err := UpdateEquipmentMaintenancePlan(after); err == nil {
		t.Fatalf("周期0应拒绝")
	}
}

// OEE：未完成故障按时限截断计入停机（防时间稼动率高估）
func TestCalcOee_OpenBreakdownCounted(t *testing.T) {
	gdb := setupEMDB(t)
	station := &model.ProductionStation{Code: "ST-OEE"}
	if err := gdb.Create(station).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	// 30分钟前开始、未完成的故障
	openBD := &model.ProductionStationBreakdown{
		ProductionStationID: station.ID,
		CreateTime:          time.Now().Add(-30 * time.Minute),
		Duration:            0, // 未完成无时长
	}
	if err := gdb.Create(openBD).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}

	start := time.Now().Add(-time.Hour).Format("2006-01-02 15:04:05")
	end := time.Now().Format("2006-01-02 15:04:05")
	resp, err := CalcOee(&proto.EquipmentOeeRequest{ProductionStationID: station.ID, StartTime: start, EndTime: end})
	if err != nil {
		t.Fatalf("OEE计算失败: %v", err)
	}
	if resp.DowntimeMinutes < 29 { // 截断计入≈30分钟（允许秒级误差）
		t.Fatalf("未完成故障应计入停机（约30分钟），实际%v", resp.DowntimeMinutes)
	}
	if resp.Availability > 0.51 {
		t.Fatalf("30/60分钟停机下时间稼动率应≤0.5，实际%v", resp.Availability)
	}
}
