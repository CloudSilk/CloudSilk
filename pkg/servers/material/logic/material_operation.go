package logic

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	system "github.com/CloudSilk/CloudSilk/pkg/servers/system/logic"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
)

// 库存事务类型
const (
	InventoryTxReceive   = "入库"
	InventoryTxIssue     = "出库"
	InventoryTxLock      = "锁定"
	InventoryTxUnlock    = "解锁"
	InventoryTxStocktake = "盘点调整"
)

// 拣货单状态
const (
	WMSBillStateWaitPick  = "待拣货"
	WMSBillStateCompleted = "已完成"
	WMSBillStateCancelled = "已取消"
)

// ReceiveMaterial 收货入库：增加账面库存并写入事务流水
func ReceiveMaterial(req *proto.ReceiveMaterialRequest, userID string) (*proto.ReceiveMaterialResponse, error) {
	if req.MaterialInfoID == "" || req.MaterialStoreID == "" {
		return nil, errors.New("materialInfoID与materialStoreID不能为空")
	}
	if req.Qty <= 0 {
		return nil, errors.New("入库数量必须大于0")
	}

	resp := &proto.ReceiveMaterialResponse{}
	err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
		inventory, err := lockOrCreateInventory(tx, req.MaterialInfoID, req.MaterialStoreID, userID)
		if err != nil {
			return err
		}

		before := inventory.StoredQTY
		inventory.StoredQTY += req.Qty
		if err := tx.Model(&model.MaterialInventory{}).Where("`id` = ?", inventory.ID).
			Update("stored_qty", inventory.StoredQTY).Error; err != nil {
			return err
		}

		if err := createInventoryTx(tx, req.MaterialInfoID, req.MaterialStoreID, InventoryTxReceive, req.Qty, before, inventory.StoredQTY, req.BatchNo, userID, req.Remark); err != nil {
			return err
		}

		resp.MaterialInventoryID = inventory.ID
		resp.StoredQTY = inventory.StoredQTY
		return nil
	})
	if err != nil {
		return nil, err
	}
	resp.Code = proto.Code_Success
	return resp, nil
}

// CreatePickBillFromOrder 按工单BOM生成拣货单并锁定库存
func CreatePickBillFromOrder(req *proto.CreatePickBillRequest, userID string) (*proto.CreatePickBillResponse, error) {
	if req.ProductOrderID == "" {
		return nil, errors.New("productOrderID不能为空")
	}

	resp := &proto.CreatePickBillResponse{}
	err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
		order := &model.ProductOrder{}
		if err := tx.Preload("ProductOrderBoms").First(order, "`id` = ?", req.ProductOrderID).Error; err != nil {
			return errors.New("读取生产工单失败")
		}
		if len(order.ProductOrderBoms) == 0 {
			return errors.New("此工单没有BOM明细，无法生成拣货单")
		}

		billNo, err := system.GenerateSerialNumber("拣货单号", "WMS拣货单流水号", fmt.Sprintf("PK%s", time.Now().Format("20060102")), 4, 1)
		if err != nil {
			return err
		}

		bill := &model.WMSBillQueue{
			BillNo:         billNo,
			MaterialStoreID: req.MaterialStoreID,
			ProductOrderID: order.ID,
			CreateUserID:   userID,
			CurrentState:   WMSBillStateWaitPick,
			Remark:         fmt.Sprintf("工单%s领料", order.ProductOrderNo),
		}
		if err := tx.Create(bill).Error; err != nil {
			return err
		}

		//逐行锁库存：需求 = 单件需求量 × 工单数量（按物料号匹配库存）
		materialInfoIDs := map[string]string{}
		for _, bom := range order.ProductOrderBoms {
			if bom.MaterialNo != "" {
				materialInfoIDs[bom.MaterialNo] = ""
			}
		}
		if len(materialInfoIDs) > 0 {
			materialNos := make([]string, 0, len(materialInfoIDs))
			for no := range materialInfoIDs {
				materialNos = append(materialNos, no)
			}
			var materialInfos []*model.MaterialInfo
			if err := tx.Select("id", "material_no").Where("`material_no` in (?)", materialNos).Find(&materialInfos).Error; err != nil {
				return err
			}
			for _, mi := range materialInfos {
				materialInfoIDs[mi.MaterialNo] = mi.ID
			}
		}

		var lockedLines int32
		var lockedQTY int64
		var shortage []string
		for _, bom := range order.ProductOrderBoms {
			if !bom.EnableControl {
				continue
			}
			required := int64(bom.RequireQTY*float32(order.OrderQTY) + 0.999)
			if required <= 0 {
				continue
			}

			materialInfoID := materialInfoIDs[bom.MaterialNo]
			if materialInfoID == "" {
				shortage = append(shortage, fmt.Sprintf("%s(物料主数据缺失)", bom.MaterialNo))
				continue
			}

			var inventory *model.MaterialInventory
			query := tx.Where("`material_info_id` = ?", materialInfoID)
			if req.MaterialStoreID != "" {
				query = query.Where("`material_store_id` = ?", req.MaterialStoreID)
			}
			if err := query.Order("stored_qty desc").First(&inventory).Error; err == gorm.ErrRecordNotFound {
				inventory = nil
			} else if err != nil {
				return err
			}

			if inventory == nil || inventory.StoredQTY-inventory.IssuedQTY < required {
				shortage = append(shortage, fmt.Sprintf("%s(需%d)", bom.MaterialNo, required))
				continue
			}

			before := inventory.StoredQTY
			if err := tx.Model(&model.MaterialInventory{}).Where("`id` = ?", inventory.ID).
				Update("issued_qty", gorm.Expr("issued_qty + ?", required)).Error; err != nil {
				return err
			}
			//锁定不改变账面库存，流水记录数量为0、锁定明细体现在备注
			if err := createInventoryTx(tx, inventory.MaterialInfoID, inventory.MaterialStoreID, InventoryTxLock, required, before, before, billNo, userID,
				fmt.Sprintf("工单%s锁定，物料：%s", order.ProductOrderNo, bom.MaterialNo)); err != nil {
				return err
			}
			lockedLines++
			lockedQTY += required
		}

		if lockedLines == 0 {
			return fmt.Errorf("拣货单生成失败，无足够库存可锁定。缺料：%s", strings.Join(shortage, "、"))
		}

		resp.BillID = bill.ID
		resp.BillNo = billNo
		resp.LockedLines = lockedLines
		resp.LockedQTY = lockedQTY
		if len(shortage) > 0 {
			resp.Message = fmt.Sprintf("部分物料缺料未锁定：%s", strings.Join(shortage, "、"))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	resp.Code = proto.Code_Success
	return resp, nil
}

// CompletePickBill 完成拣货：扣减账面库存、解锁、联动工单发料数量，可选生成AGV任务
func CompletePickBill(req *proto.CompletePickBillRequest, userID string) error {
	if req.BillID == "" {
		return errors.New("billID不能为空")
	}

	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		bill := &model.WMSBillQueue{}
		if err := tx.First(bill, "`id` = ?", req.BillID).Error; err != nil {
			return errors.New("读取拣货单失败")
		}
		if bill.CurrentState != WMSBillStateWaitPick {
			return fmt.Errorf("拣货单状态为%s，只有%s状态可以完成", bill.CurrentState, WMSBillStateWaitPick)
		}

		//按流水中该单的锁定记录逐行出库
		locks := []*model.MaterialInventoryTransaction{}
		if err := tx.Where("`ref_bill_no` = ? AND `transaction_type` = ?", bill.BillNo, InventoryTxLock).Find(&locks).Error; err != nil {
			return err
		}

		var totalIssued int64
		for _, lock := range locks {
			inventory := &model.MaterialInventory{}
			if err := tx.First(inventory, "`id` = ?", lock.MaterialInfoID).Error; err != nil {
				return fmt.Errorf("读取库存记录失败：%w", err)
			}
			before := inventory.StoredQTY
			if err := tx.Model(&model.MaterialInventory{}).Where("`id` = ?", inventory.ID).
				Updates(map[string]interface{}{
					"stored_qty": gorm.Expr("stored_qty - ?", lock.Qty),
					"issued_qty": gorm.Expr("issued_qty - ?", lock.Qty),
				}).Error; err != nil {
				return err
			}
			if err := createInventoryTx(tx, lock.MaterialInfoID, lock.MaterialStoreID, InventoryTxIssue, -lock.Qty, before, before-lock.Qty, bill.BillNo, userID, "拣货出库"); err != nil {
				return err
			}
			totalIssued += lock.Qty
		}

		//联动工单发料信息
		if bill.ProductOrderID != "" {
			now := time.Now()
			if err := tx.Model(&model.ProductOrder{}).Where("`id` = ?", bill.ProductOrderID).
				Updates(map[string]interface{}{
					"issued_qty":      gorm.Expr("issued_qty + ?", totalIssued),
					"last_issue_time": now,
				}).Error; err != nil {
				return err
			}
		}

		//可选：生成AGV搬运任务（待签派状态），复用AGV任务队列
		if req.CreateAGVTask {
			taskNo, err := system.GenerateSerialNumber("AGV任务号", "拣货AGV搬运任务流水号", fmt.Sprintf("AG%s", time.Now().Format("20060102")), 4, 1)
			if err != nil {
				return err
			}
			if err := tx.Create(&model.AGVTaskQueue{
				TaskNo:       taskNo,
				CreateUserID: userID,
				CurrentState: types.AGVTaskQueueStateWaitDispatch,
				Remark:       fmt.Sprintf("拣货单%s出库搬运", bill.BillNo),
			}).Error; err != nil {
				return err
			}
		}

		if err := tx.Model(&model.WMSBillQueue{}).Where("`id` = ?", bill.ID).
			Update("current_state", WMSBillStateCompleted).Error; err != nil {
			return err
		}
		return nil
	})
}

// CancelPickBill 取消拣货单：解除库存锁定
func CancelPickBill(req *proto.CancelPickBillRequest, userID string) error {
	if req.BillID == "" {
		return errors.New("billID不能为空")
	}

	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		bill := &model.WMSBillQueue{}
		if err := tx.First(bill, "`id` = ?", req.BillID).Error; err != nil {
			return errors.New("读取拣货单失败")
		}
		if bill.CurrentState != WMSBillStateWaitPick {
			return fmt.Errorf("拣货单状态为%s，只有%s状态可以取消", bill.CurrentState, WMSBillStateWaitPick)
		}

		locks := []*model.MaterialInventoryTransaction{}
		if err := tx.Where("`ref_bill_no` = ? AND `transaction_type` = ?", bill.BillNo, InventoryTxLock).Find(&locks).Error; err != nil {
			return err
		}

		for _, lock := range locks {
			inventory := &model.MaterialInventory{}
			if err := tx.First(inventory, "`id` = ?", lock.MaterialInfoID).Error; err != nil {
				return fmt.Errorf("读取库存记录失败：%w", err)
			}
			before := inventory.StoredQTY
			if err := tx.Model(&model.MaterialInventory{}).Where("`id` = ?", lock.MaterialInfoID).
				Update("issued_qty", gorm.Expr("issued_qty - ?", lock.Qty)).Error; err != nil {
				return err
			}
			if err := createInventoryTx(tx, lock.MaterialInfoID, lock.MaterialStoreID, InventoryTxUnlock, lock.Qty, before, before, bill.BillNo, userID, "取消拣货解锁"); err != nil {
				return err
			}
		}

		return tx.Model(&model.WMSBillQueue{}).Where("`id` = ?", bill.ID).
			Update("current_state", WMSBillStateCancelled).Error
	})
}

// StocktakeInventory 库存盘点：按实盘数量调整账面并记录差异
func StocktakeInventory(req *proto.StocktakeInventoryRequest, userID string) (*proto.StocktakeInventoryResponse, error) {
	if req.MaterialInventoryID == "" {
		return nil, errors.New("materialInventoryID不能为空")
	}
	if req.CountedQTY < 0 {
		return nil, errors.New("实盘数量不能为负数")
	}

	resp := &proto.StocktakeInventoryResponse{}
	err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
		inventory := &model.MaterialInventory{}
		if err := tx.First(inventory, "`id` = ?", req.MaterialInventoryID).Error; err != nil {
			return errors.New("读取库存记录失败")
		}

		before := inventory.StoredQTY
		difference := req.CountedQTY - before
		if err := tx.Model(&model.MaterialInventory{}).Where("`id` = ?", inventory.ID).
			Update("stored_qty", req.CountedQTY).Error; err != nil {
			return err
		}

		if err := createInventoryTx(tx, inventory.MaterialInfoID, inventory.MaterialStoreID, InventoryTxStocktake, difference, before, req.CountedQTY, "", userID,
			fmt.Sprintf("盘点差异%d（账面%d→实盘%d）：%s", difference, before, req.CountedQTY, req.Remark)); err != nil {
			return err
		}

		resp.Difference = difference
		return nil
	})
	if err != nil {
		return nil, err
	}
	resp.Code = proto.Code_Success
	return resp, nil
}

// QueryMaterialInventoryTransaction 查询库存事务流水
func QueryMaterialInventoryTransaction(req *proto.QueryMaterialInventoryTransactionRequest, resp *proto.QueryMaterialInventoryTransactionResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialInventoryTransaction{})
	if preload {
		db = db.Preload("MaterialInfo").Preload("MaterialStore")
	}
	if req.MaterialInfoID != "" {
		db = db.Where("`material_info_id` = ?", req.MaterialInfoID)
	}
	if req.MaterialStoreID != "" {
		db = db.Where("`material_store_id` = ?", req.MaterialStoreID)
	}
	if req.TransactionType != "" {
		db = db.Where("`transaction_type` = ?", req.TransactionType)
	}
	if req.RefBillNo != "" {
		db = db.Where("`ref_bill_no` = ?", req.RefBillNo)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "create_time desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialInventoryTransaction
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialInventoryTransactionsToPB(list)
	}
	resp.Total = resp.Records
}

// GetInventoryAlerts 库存预警：可用量（账面-锁定）低于补料规则最低库存量或已为负的记录
func GetInventoryAlerts() (*proto.MaterialInventoryAlertResponse, error) {
	resp := &proto.MaterialInventoryAlertResponse{Code: proto.Code_Success}

	inventories := []*model.MaterialInventory{}
	if err := model.DB.DB().Preload("MaterialInfo").Preload("MaterialStore").Find(&inventories).Error; err != nil {
		return nil, err
	}

	rules := []*model.MaterialStoreFeedRule{}
	if err := model.DB.DB().Find(&rules, "`enable` = ?", true).Error; err != nil {
		return nil, err
	}
	minimumByMaterial := map[string]int64{}
	for _, rule := range rules {
		key := rule.MaterialStoreID + "|" + rule.MaterialInfoID
		minimumByMaterial[key] = rule.MinimumStoreQTY
	}

	for _, inv := range inventories {
		available := inv.StoredQTY - inv.IssuedQTY
		minimum := minimumByMaterial[inv.MaterialStoreID+"|"+inv.MaterialInfoID]
		if available > minimum && available > 0 {
			continue
		}
		item := &proto.MaterialInventoryAlertInfo{
			MaterialInventoryID: inv.ID,
			MaterialInfoID:      inv.MaterialInfoID,
			MaterialStoreID:     inv.MaterialStoreID,
			StoredQTY:           inv.StoredQTY,
			IssuedQTY:           inv.IssuedQTY,
			AvailableQTY:        available,
			MinimumQTY:          minimum,
		}
		if inv.MaterialInfo != nil {
			item.MaterialNo = inv.MaterialInfo.MaterialNo
			item.MaterialDescription = inv.MaterialInfo.MaterialDescription
		}
		if inv.MaterialStore != nil {
			item.MaterialStoreCode = inv.MaterialStore.Code
		}
		resp.Data = append(resp.Data, item)
	}
	return resp, nil
}

// lockOrCreateInventory 查找或创建库存记录（SELECT ... 行锁由事务保证）
func lockOrCreateInventory(tx *gorm.DB, materialInfoID, materialStoreID, userID string) (*model.MaterialInventory, error) {
	inventory := &model.MaterialInventory{}
	err := tx.Where("`material_info_id` = ? AND `material_store_id` = ?", materialInfoID, materialStoreID).First(inventory).Error
	if err == gorm.ErrRecordNotFound {
		inventory = &model.MaterialInventory{
			MaterialInfoID:  materialInfoID,
			MaterialStoreID: materialStoreID,
			CreateUserID:    userID,
			CurrentState:    "正常",
		}
		if err := tx.Create(inventory).Error; err != nil {
			return nil, err
		}
		return inventory, nil
	}
	if err != nil {
		return nil, err
	}
	return inventory, nil
}

func createInventoryTx(tx *gorm.DB, materialInfoID, materialStoreID, txType string, qty, before, after int64, refBillNo, userID, remark string) error {
	return tx.Create(&model.MaterialInventoryTransaction{
		MaterialInfoID:  materialInfoID,
		MaterialStoreID: materialStoreID,
		TransactionType: txType,
		Qty:             qty,
		BeforeQTY:       before,
		AfterQTY:        after,
		RefBillNo:       refBillNo,
		CreateUserID:    userID,
		Remark:          remark,
	}).Error
}
