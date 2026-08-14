package logic

import (
	"testing"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/testutil"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"gorm.io/gorm"
)

func setupOrderDB(t *testing.T) *gorm.DB {
	t.Helper()
	gdb, err := testutil.SetupTestDB()
	if err != nil {
		t.Fatalf("初始化测试数据库失败: %v", err)
	}
	return gdb
}

func seedOrder(t *testing.T, gdb *gorm.DB, state string) *model.ProductOrder {
	t.Helper()
	order := &model.ProductOrder{
		ProductOrderNo: "WO-TEST-001",
		OrderQTY:       10,
		CurrentState:   state,
		ProductionTeam: "甲班",
	}
	if err := gdb.Create(order).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	return order
}

// 状态机：已核验 → 签派 → 已签派，并记录 OperationTrace
func TestDispatchProductOrder_HappyPath(t *testing.T) {
	gdb := setupOrderDB(t)
	order := seedOrder(t, gdb, types.ProductOrderStateVerified)

	if err := DispatchProductOrder(gdb, order.ID, "user-1", "甲班"); err != nil {
		t.Fatalf("签派失败: %v", err)
	}

	after := &model.ProductOrder{}
	if err := gdb.First(after, "`id` = ?", order.ID).Error; err != nil {
		t.Fatalf("回读失败: %v", err)
	}
	if after.CurrentState != types.ProductOrderStateDispatched {
		t.Fatalf("签派后状态应为%s，实际%s", types.ProductOrderStateDispatched, after.CurrentState)
	}
	if after.ProductionTeam != "甲班" {
		t.Fatalf("签派应写入生产班组")
	}

	var traces int64
	gdb.Model(&model.OperationTrace{}).Where("`action_name` = ?", "签派").Count(&traces)
	if traces != 1 {
		t.Fatalf("应记录1条签派轨迹，实际%d", traces)
	}
}

// 状态机：非已核验状态不能签派
func TestDispatchProductOrder_WrongState(t *testing.T) {
	gdb := setupOrderDB(t)
	order := seedOrder(t, gdb, types.ProductOrderStateUploaded)

	if err := DispatchProductOrder(gdb, order.ID, "user-1", ""); err == nil {
		t.Fatalf("已上传状态不应允许签派")
	}
}

// 批量签派：混入错误状态的工单整批拒绝
func TestBatchDispatchProductOrder_RejectMixed(t *testing.T) {
	gdb := setupOrderDB(t)
	okOrder := seedOrder(t, gdb, types.ProductOrderStateVerified)
	badOrder := &model.ProductOrder{ProductOrderNo: "WO-BAD", CurrentState: types.ProductOrderStateReleased}
	if err := gdb.Create(badOrder).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}

	if err := BatchDispatchProductOrder([]string{okOrder.ID, badOrder.ID}, "user-1", ""); err == nil {
		t.Fatalf("混合状态批量签派应失败")
	}

	// 整批拒绝：okOrder 也不应被签派（事务回滚）
	after := &model.ProductOrder{}
	gdb.First(after, "`id` = ?", okOrder.ID)
	if after.CurrentState != types.ProductOrderStateVerified {
		t.Fatalf("批量失败时应保持原状态，实际%s", after.CurrentState)
	}
}

// 发放：已签派状态可直接发放（签派环节打通的回归验证）
func TestReleaseProductOrder_AcceptsDispatched(t *testing.T) {
	gdb := setupOrderDB(t)

	line := &model.ProductionLine{Code: "L1"}
	if err := gdb.Create([]*model.ProductionLine{line}).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	rhythm := &model.ProductionRhythm{Priority: 1, Enable: true, StandardTime: 60, ProductionLineID: line.ID, InitialValue: true}
	if err := gdb.Create([]*model.ProductionRhythm{rhythm}).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	process := &model.ProductionProcess{Code: "P1", SortIndex: 1, Enable: true, InitialValue: true}
	if err := gdb.Create([]*model.ProductionProcess{process}).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}

	order := &model.ProductOrder{
		ProductOrderNo:   "WO-REL-001",
		OrderQTY:         2,
		CurrentState:     types.ProductOrderStateDispatched,
		ProductionLineID: &line.ID,
	}
	if err := gdb.Create(order).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}

	if err := ReleaseProductOrder(gdb, order.ID); err != nil {
		t.Fatalf("已签派工单发放失败: %v", err)
	}

	after := &model.ProductOrder{}
	gdb.First(after, "`id` = ?", order.ID)
	if after.CurrentState != types.ProductOrderStateReleased {
		t.Fatalf("发放后状态应为%s，实际%s", types.ProductOrderStateReleased, after.CurrentState)
	}
}
