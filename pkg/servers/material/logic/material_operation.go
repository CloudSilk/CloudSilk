package logic

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	product "github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	system "github.com/CloudSilk/CloudSilk/pkg/servers/system/logic"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	//收货库位（可选）：存在性校验
	var shelfBinID *string
	if req.ShelfBinID != "" {
		bin := &model.MaterialShelfBin{}
		if err := model.DB.DB().Preload("MaterialShelf").First(bin, "`id` = ?", req.ShelfBinID).Error; err != nil {
			return nil, errors.New("无效的库位ID")
		}
		//库位经料架归属仓库，跨仓库上架直接拒绝
		if bin.MaterialShelf != nil && bin.MaterialShelf.MaterialStoreID != "" && bin.MaterialShelf.MaterialStoreID != req.MaterialStoreID {
			return nil, errors.New("库位不属于该仓库")
		}
		shelfBinID = &req.ShelfBinID
	}

	resp := &proto.ReceiveMaterialResponse{}
	err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
		inventory, err := lockOrCreateInventoryWithBin(tx, req.MaterialInfoID, req.MaterialStoreID, shelfBinID, userID)
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
		return tx.Create(&model.OperationTrace{
			OperateUserID:  userID,
			ControllerName: "WMS作业",
			ActionName:     "收货入库",
			RequestContent: fmt.Sprintf("物料:%s,仓库:%s,数量:%d,入库后:%d", req.MaterialInfoID, req.MaterialStoreID, req.Qty, inventory.StoredQTY),
		}).Error
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
	//流水号在事务外生成（避免事务内跨连接写引发 SQLite 表锁互等）
	billNo, err := system.GenerateSerialNumber("拣货单号", "WMS拣货单流水号", fmt.Sprintf("PK%s", time.Now().Format("20060102")), 4, 1)
	if err != nil {
		return nil, err
	}
	err = model.DB.DB().Transaction(func(tx *gorm.DB) error {
		order, err := product.GetProductOrderForPick(tx, req.ProductOrderID)
		if err != nil {
			return err
		}
		if len(order.ProductOrderBoms) == 0 {
			return errors.New("此工单没有BOM明细，无法生成拣货单")
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
			//行项目落库（拣货明细的一等公民，完成/取消按行驱动）
			if err := tx.Create(&model.WMSBillItem{
				WMSBillQueueID:      bill.ID,
				ProductOrderBomID:   bom.ID,
				MaterialInfoID:      materialInfoID,
				MaterialNo:          bom.MaterialNo,
				MaterialDescription: bom.MaterialDescription,
				MaterialStoreID:     inventory.MaterialStoreID,
				RequiredQTY:         required,
				LockedQTY:           required,
				CurrentState:        WMSBillStateWaitPick,
			}).Error; err != nil {
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

	//AGV任务号在事务外生成（同上，避免事务内跨连接写）
	agvTaskNo := ""
	if req.CreateAGVTask {
		taskNo, err := system.GenerateSerialNumber("AGV任务号", "拣货AGV搬运任务流水号", fmt.Sprintf("AG%s", time.Now().Format("20060102")), 4, 1)
		if err != nil {
			return err
		}
		agvTaskNo = taskNo
	}

	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		bill := &model.WMSBillQueue{}
		if err := tx.First(bill, "`id` = ?", req.BillID).Error; err != nil {
			return errors.New("读取拣货单失败")
		}
		//原子状态门：并发完成/取消只有一个能通过（先占状态再干活）
		res := tx.Model(&model.WMSBillQueue{}).
			Where("`id` = ? AND `current_state` = ?", bill.ID, WMSBillStateWaitPick).
			Update("current_state", WMSBillStateCompleted)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("拣货单已被处理（当前状态%s），请刷新后重试", bill.CurrentState)
		}

		//优先按行项目出库（新数据）；无行项目的旧单据回退到流水反查
		items := []*model.WMSBillItem{}
		if err := tx.Where("`wms_bill_queue_id` = ? AND `current_state` = ?", bill.ID, WMSBillStateWaitPick).Find(&items).Error; err != nil {
			return err
		}

		var totalIssued int64
		if len(items) > 0 {
			for _, item := range items {
				inventory := &model.MaterialInventory{}
				if err := tx.Where("`material_info_id` = ? AND `material_store_id` = ?", item.MaterialInfoID, item.MaterialStoreID).First(inventory).Error; err != nil {
					return fmt.Errorf("读取库存记录失败（物料%s）：%w", item.MaterialNo, err)
				}
				before := inventory.StoredQTY
				if err := tx.Model(&model.MaterialInventory{}).Where("`id` = ?", inventory.ID).
					Updates(map[string]interface{}{
						"stored_qty": gorm.Expr("stored_qty - ?", item.LockedQTY),
						"issued_qty": gorm.Expr("issued_qty - ?", item.LockedQTY),
					}).Error; err != nil {
					return err
				}
				if err := createInventoryTx(tx, item.MaterialInfoID, item.MaterialStoreID, InventoryTxIssue, -item.LockedQTY, before, before-item.LockedQTY, bill.BillNo, userID, fmt.Sprintf("拣货出库：%s", item.MaterialNo)); err != nil {
					return err
				}
				if err := tx.Model(&model.WMSBillItem{}).Where("`id` = ?", item.ID).
					Update("current_state", WMSBillStateCompleted).Error; err != nil {
					return err
				}
				totalIssued += item.LockedQTY
			}
		} else {
			//旧单据回退：按锁定流水反查
			locks := []*model.MaterialInventoryTransaction{}
			if err := tx.Where("`ref_bill_no` = ? AND `transaction_type` = ?", bill.BillNo, InventoryTxLock).Find(&locks).Error; err != nil {
				return err
			}
			for _, lock := range locks {
				inventory := &model.MaterialInventory{}
				if err := tx.Where("`material_info_id` = ? AND `material_store_id` = ?", lock.MaterialInfoID, lock.MaterialStoreID).First(inventory).Error; err != nil {
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
		}

		//联动工单发料信息（product 域出口）
		if bill.ProductOrderID != "" {
			if err := product.AddProductOrderIssued(tx, bill.ProductOrderID, totalIssued, time.Now()); err != nil {
				return err
			}
		}

		//可选：生成AGV搬运任务（待签派状态），复用AGV任务队列
		if req.CreateAGVTask {
			if err := tx.Create(&model.AGVTaskQueue{
				TaskNo:       agvTaskNo,
				CreateUserID: userID,
				CurrentState: types.AGVTaskQueueStateWaitDispatch,
				Remark:       fmt.Sprintf("拣货单%s出库搬运", bill.BillNo),
			}).Error; err != nil {
				return err
			}
		}

		return tx.Create(&model.OperationTrace{
			OperateUserID:  userID,
			ControllerName: "WMS作业",
			ActionName:     "完成拣货",
			RequestContent: fmt.Sprintf("拣货单:%s,出库总量:%d,工单:%s,AGV:%v", bill.BillNo, totalIssued, bill.ProductOrderID, req.CreateAGVTask),
		}).Error
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
		//原子状态门（与完成互斥）
		res := tx.Model(&model.WMSBillQueue{}).
			Where("`id` = ? AND `current_state` = ?", bill.ID, WMSBillStateWaitPick).
			Update("current_state", WMSBillStateCancelled)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("拣货单已被处理（当前状态%s），请刷新后重试", bill.CurrentState)
		}

		//优先按行项目解锁（新数据）；旧单据回退流水反查
		items := []*model.WMSBillItem{}
		if err := tx.Where("`wms_bill_queue_id` = ? AND `current_state` = ?", bill.ID, WMSBillStateWaitPick).Find(&items).Error; err != nil {
			return err
		}
		if len(items) > 0 {
			for _, item := range items {
				inventory := &model.MaterialInventory{}
				if err := tx.Where("`material_info_id` = ? AND `material_store_id` = ?", item.MaterialInfoID, item.MaterialStoreID).First(inventory).Error; err != nil {
					return fmt.Errorf("读取库存记录失败（物料%s）：%w", item.MaterialNo, err)
				}
				before := inventory.StoredQTY
				if err := tx.Model(&model.MaterialInventory{}).Where("`id` = ?", inventory.ID).
					Update("issued_qty", gorm.Expr("issued_qty - ?", item.LockedQTY)).Error; err != nil {
					return err
				}
				if err := createInventoryTx(tx, item.MaterialInfoID, item.MaterialStoreID, InventoryTxUnlock, item.LockedQTY, before, before, bill.BillNo, userID, fmt.Sprintf("取消拣货解锁：%s", item.MaterialNo)); err != nil {
					return err
				}
				if err := tx.Model(&model.WMSBillItem{}).Where("`id` = ?", item.ID).
					Update("current_state", WMSBillStateCancelled).Error; err != nil {
					return err
				}
			}
		} else {
			locks := []*model.MaterialInventoryTransaction{}
			if err := tx.Where("`ref_bill_no` = ? AND `transaction_type` = ?", bill.BillNo, InventoryTxLock).Find(&locks).Error; err != nil {
				return err
			}
			for _, lock := range locks {
				inventory := &model.MaterialInventory{}
				if err := tx.Where("`material_info_id` = ? AND `material_store_id` = ?", lock.MaterialInfoID, lock.MaterialStoreID).First(inventory).Error; err != nil {
					return fmt.Errorf("读取库存记录失败：%w", err)
				}
				before := inventory.StoredQTY
				if err := tx.Model(&model.MaterialInventory{}).Where("`id` = ?", inventory.ID).
					Update("issued_qty", gorm.Expr("issued_qty - ?", lock.Qty)).Error; err != nil {
					return err
				}
				if err := createInventoryTx(tx, lock.MaterialInfoID, lock.MaterialStoreID, InventoryTxUnlock, lock.Qty, before, before, bill.BillNo, userID, "取消拣货解锁"); err != nil {
					return err
				}
			}
		}

		return tx.Create(&model.OperationTrace{
			OperateUserID:  userID,
			ControllerName: "WMS作业",
			ActionName:     "取消拣货",
			RequestContent: fmt.Sprintf("拣货单:%s,解除锁定", bill.BillNo),
		}).Error
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

		if req.CountedQTY < inventory.IssuedQTY {
			return fmt.Errorf("实盘数量%d小于当前锁定量%d，请先处理在途拣货单再盘点", req.CountedQTY, inventory.IssuedQTY)
		}
		before := inventory.StoredQTY
		difference := req.CountedQTY - before
		//乐观锁：账面自读取以来未变才允许写（并发盘点/收发互斥）
		res := tx.Model(&model.MaterialInventory{}).
			Where("`id` = ? AND `stored_qty` = ?", inventory.ID, before).
			Update("stored_qty", req.CountedQTY)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return errors.New("库存账面已发生变化（并发收发/盘点），请刷新后重新盘点")
		}

		if err := createInventoryTx(tx, inventory.MaterialInfoID, inventory.MaterialStoreID, InventoryTxStocktake, difference, before, req.CountedQTY, "", userID,
			fmt.Sprintf("盘点差异%d（账面%d→实盘%d）：%s", difference, before, req.CountedQTY, req.Remark)); err != nil {
			return err
		}

		resp.Difference = difference
		return tx.Create(&model.OperationTrace{
			OperateUserID:  userID,
			ControllerName: "WMS作业",
			ActionName:     "库存盘点",
			RequestContent: fmt.Sprintf("库存:%s,账面:%d,实盘:%d,差异:%d", req.MaterialInventoryID, before, req.CountedQTY, difference),
		}).Error
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

// lockOrCreateInventory 查找或创建库存记录；并发下唯一索引兜底（冲突即回查）
func lockOrCreateInventoryWithBin(tx *gorm.DB, materialInfoID, materialStoreID string, shelfBinID *string, userID string) (*model.MaterialInventory, error) {
	inv, err := lockOrCreateInventory(tx, materialInfoID, materialStoreID, userID)
	if err != nil {
		return nil, err
	}
	//首次建账时记录默认库位（已有库位仅在为空时补记，不覆盖人工调整）
	if shelfBinID != nil && inv.ShelfBinID == nil {
		if err := tx.Model(&model.MaterialInventory{}).Where("`id` = ?", inv.ID).
			Update("shelf_bin_id", *shelfBinID).Error; err != nil {
			return nil, err
		}
		inv.ShelfBinID = shelfBinID
	}
	return inv, nil
}

func lockOrCreateInventory(tx *gorm.DB, materialInfoID, materialStoreID, userID string) (*model.MaterialInventory, error) {
	inventory := &model.MaterialInventory{}
	err := tx.Where("`material_info_id` = ? AND `material_store_id` = ?", materialInfoID, materialStoreID).First(inventory).Error
	if err == gorm.ErrRecordNotFound {
		created := &model.MaterialInventory{
			MaterialInfoID:  materialInfoID,
			MaterialStoreID: materialStoreID,
			CreateUserID:    userID,
			CurrentState:    "正常",
		}
		res := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(created)
		if res.Error != nil {
			return nil, res.Error
		}
		if res.RowsAffected == 0 {
			// 并发对手已建账，回查
			if err := tx.Where("`material_info_id` = ? AND `material_store_id` = ?", materialInfoID, materialStoreID).First(inventory).Error; err != nil {
				return nil, err
			}
			return inventory, nil
		}
		return created, nil
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
