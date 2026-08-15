package logic

import (
	"database/sql"
	"testing"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
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
