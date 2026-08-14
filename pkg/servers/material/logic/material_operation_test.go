package logic

import (
	"testing"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/testutil"
	"gorm.io/gorm"
)

func setupWMSDB(t *testing.T) *gorm.DB {
	t.Helper()
	gdb, err := testutil.SetupTestDB()
	if err != nil {
		t.Fatalf("初始化测试数据库失败: %v", err)
	}
	return gdb
}

func seedPickFixture(t *testing.T, gdb *gorm.DB, storedQty int64) (*model.MaterialStore, *model.MaterialInfo, *model.ProductOrder) {
	t.Helper()
	store := &model.MaterialStore{Code: "WH1"}
	material := &model.MaterialInfo{MaterialNo: "MAT-001", MaterialDescription: "电阻"}
	order := &model.ProductOrder{ProductOrderNo: "WO-PICK-001", OrderQTY: 10, CurrentState: "已发放"}
	if err := gdb.Create([]*model.MaterialStore{store}).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	if err := gdb.Create([]*model.MaterialInfo{material}).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	if err := gdb.Create([]*model.ProductOrder{order}).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	//BOM：单件需求1个，管控
	bom := &model.ProductOrderBom{
		ProductOrderID: order.ID,
		MaterialNo:     "MAT-001",
		RequireQTY:     1,
		EnableControl:  true,
		Warehouse:      "WH1",
	}
	if err := gdb.Create([]*model.ProductOrderBom{bom}).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	//库存
	inventory := &model.MaterialInventory{
		MaterialInfoID:  material.ID,
		MaterialStoreID: store.ID,
		StoredQTY:       storedQty,
	}
	if err := gdb.Create([]*model.MaterialInventory{inventory}).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	return store, material, order
}

func getInventory(t *testing.T, gdb *gorm.DB, materialID, storeID string) *model.MaterialInventory {
	t.Helper()
	inv := &model.MaterialInventory{}
	if err := gdb.Where("`material_info_id` = ? AND `material_store_id` = ?", materialID, storeID).First(inv).Error; err != nil {
		t.Fatalf("读取库存失败: %v", err)
	}
	return inv
}

// 全流程：入库建账 → 生成拣货单（锁定） → 完成拣货（扣减+工单联动）
func TestPickFlow_LockThenComplete(t *testing.T) {
	gdb := setupWMSDB(t)
	store, material, order := seedPickFixture(t, gdb, 0)

	//1. 收货入库 100
	recvResp, err := ReceiveMaterial(&proto.ReceiveMaterialRequest{
		MaterialInfoID: material.ID, MaterialStoreID: store.ID, Qty: 100,
	}, "user-1")
	if err != nil {
		t.Fatalf("入库失败: %v", err)
	}
	if recvResp.StoredQTY != 100 {
		t.Fatalf("入库后库存应为100，实际%d", recvResp.StoredQTY)
	}

	//2. 按工单生成拣货单（需求 = 1×10）
	pickResp, err := CreatePickBillFromOrder(&proto.CreatePickBillRequest{ProductOrderID: order.ID}, "user-1")
	if err != nil {
		t.Fatalf("生成拣货单失败: %v", err)
	}
	if pickResp.LockedQTY != 10 || pickResp.LockedLines != 1 {
		t.Fatalf("应锁定1行共10个，实际%d行%d个", pickResp.LockedLines, pickResp.LockedQTY)
	}

	//锁定后：账面不变、锁定量增加
	inv := getInventory(t, gdb, material.ID, store.ID)
	if inv.StoredQTY != 100 || inv.IssuedQTY != 10 {
		t.Fatalf("锁定后应 stored=100 issued=10，实际 stored=%d issued=%d", inv.StoredQTY, inv.IssuedQTY)
	}

	//3. 完成拣货（含AGV任务生成）
	if err := CompletePickBill(&proto.CompletePickBillRequest{BillID: pickResp.BillID, CreateAGVTask: true}, "user-1"); err != nil {
		t.Fatalf("完成拣货失败: %v", err)
	}

	inv = getInventory(t, gdb, material.ID, store.ID)
	if inv.StoredQTY != 90 || inv.IssuedQTY != 0 {
		t.Fatalf("完成拣货后应 stored=90 issued=0，实际 stored=%d issued=%d", inv.StoredQTY, inv.IssuedQTY)
	}

	//工单发料联动
	after := &model.ProductOrder{}
	gdb.First(after, "`id` = ?", order.ID)
	if after.IssuedQTY != 10 || !after.LastIssueTime.Valid {
		t.Fatalf("工单发料数量应为10且记录发料时间，实际 issued=%d valid=%v", after.IssuedQTY, after.LastIssueTime.Valid)
	}

	//AGV任务已生成且为待签派
	var agvCount int64
	gdb.Model(&model.AGVTaskQueue{}).Where("`current_state` = ?", "待签派").Count(&agvCount)
	if agvCount != 1 {
		t.Fatalf("应生成1条待签派AGV任务，实际%d", agvCount)
	}

	//流水：入库1 + 锁定1 + 出库1
	var txCount int64
	gdb.Model(&model.MaterialInventoryTransaction{}).Count(&txCount)
	if txCount != 3 {
		t.Fatalf("应产生3条事务流水，实际%d", txCount)
	}

	//重复完成应拒绝
	if err := CompletePickBill(&proto.CompletePickBillRequest{BillID: pickResp.BillID}, "user-1"); err == nil {
		t.Fatalf("已完成的拣货单不能重复完成")
	}
}

// 缺料：库存不足时生成拣货单应报错且回滚
func TestPickFlow_Shortage(t *testing.T) {
	gdb := setupWMSDB(t)
	store, material, order := seedPickFixture(t, gdb, 5)

	if _, err := CreatePickBillFromOrder(&proto.CreatePickBillRequest{ProductOrderID: order.ID}, "user-1"); err == nil {
		t.Fatalf("库存不足（5<10）应生成失败")
	}

	//不应留下锁定
	inv := getInventory(t, gdb, material.ID, store.ID)
	if inv.IssuedQTY != 0 {
		t.Fatalf("失败后不应有锁定量，实际%d", inv.IssuedQTY)
	}
}

// 取消拣货单：解除锁定且账面不变
func TestPickFlow_CancelUnlocks(t *testing.T) {
	gdb := setupWMSDB(t)
	store, material, order := seedPickFixture(t, gdb, 100)

	pickResp, err := CreatePickBillFromOrder(&proto.CreatePickBillRequest{ProductOrderID: order.ID}, "user-1")
	if err != nil {
		t.Fatalf("生成拣货单失败: %v", err)
	}
	if err := CancelPickBill(&proto.CancelPickBillRequest{BillID: pickResp.BillID}, "user-1"); err != nil {
		t.Fatalf("取消失败: %v", err)
	}

	inv := getInventory(t, gdb, material.ID, store.ID)
	if inv.StoredQTY != 100 || inv.IssuedQTY != 0 {
		t.Fatalf("取消后应 stored=100 issued=0，实际 stored=%d issued=%d", inv.StoredQTY, inv.IssuedQTY)
	}

	bill := &model.WMSBillQueue{}
	gdb.First(bill, "`id` = ?", pickResp.BillID)
	if bill.CurrentState != "已取消" {
		t.Fatalf("拣货单状态应为已取消，实际%s", bill.CurrentState)
	}
}

// 盘点：差异调整与流水
func TestStocktake(t *testing.T) {
	gdb := setupWMSDB(t)
	store, material, _ := seedPickFixture(t, gdb, 80)

	inv := getInventory(t, gdb, material.ID, store.ID)
	resp, err := StocktakeInventory(&proto.StocktakeInventoryRequest{MaterialInventoryID: inv.ID, CountedQTY: 77}, "user-1")
	if err != nil {
		t.Fatalf("盘点失败: %v", err)
	}
	if resp.Difference != -3 {
		t.Fatalf("盘点差异应为-3，实际%d", resp.Difference)
	}

	after := getInventory(t, gdb, material.ID, store.ID)
	if after.StoredQTY != 77 {
		t.Fatalf("盘点后账面应为77，实际%d", after.StoredQTY)
	}
}
