package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 库存事务流水（入库/出库/锁定/解锁/盘点调整）
type MaterialInventoryTransaction struct {
	ModelID
	MaterialInfoID  string         `json:"materialInfoID" gorm:"index;size:36;comment:物料信息ID"`
	MaterialInfo    *MaterialInfo  `json:"materialInfo" gorm:"constraint:OnDelete:CASCADE"`
	MaterialStoreID string         `json:"materialStoreID" gorm:"size:36;comment:物料仓库ID"`
	MaterialStore   *MaterialStore `json:"materialStore" gorm:"constraint:OnDelete:CASCADE"`
	TransactionType string         `json:"transactionType" gorm:"size:50;comment:事务类型"`
	Qty             int64          `json:"qty" gorm:"comment:数量(带符号)"`
	BeforeQTY       int64          `json:"beforeQTY" gorm:"comment:变更前库存"`
	AfterQTY        int64          `json:"afterQTY" gorm:"comment:变更后库存"`
	RefBillNo       string         `json:"refBillNo" gorm:"index;size:50;comment:关联单号"`
	CreateTime      time.Time      `json:"createTime" gorm:"autoCreateTime:nano;comment:操作时间"`
	CreateUserID    string         `json:"createUserID" gorm:"size:36;comment:操作人"`
	Remark          string         `json:"remark" gorm:"size:500;comment:备注"`
}

func PBToMaterialInventoryTransactions(in []*proto.MaterialInventoryTransactionInfo) []*MaterialInventoryTransaction {
	var result []*MaterialInventoryTransaction
	for _, c := range in {
		result = append(result, PBToMaterialInventoryTransaction(c))
	}
	return result
}

func PBToMaterialInventoryTransaction(in *proto.MaterialInventoryTransactionInfo) *MaterialInventoryTransaction {
	if in == nil {
		return nil
	}
	return &MaterialInventoryTransaction{
		ModelID:         ModelID{ID: in.Id},
		MaterialInfoID:  in.MaterialInfoID,
		MaterialStoreID: in.MaterialStoreID,
		TransactionType: in.TransactionType,
		Qty:             in.Qty,
		BeforeQTY:       in.BeforeQTY,
		AfterQTY:        in.AfterQTY,
		RefBillNo:       in.RefBillNo,
		CreateUserID:    in.CreateUserID,
		Remark:          in.Remark,
	}
}

func MaterialInventoryTransactionsToPB(in []*MaterialInventoryTransaction) []*proto.MaterialInventoryTransactionInfo {
	var list []*proto.MaterialInventoryTransactionInfo
	for _, f := range in {
		list = append(list, MaterialInventoryTransactionToPB(f))
	}
	return list
}

func MaterialInventoryTransactionToPB(in *MaterialInventoryTransaction) *proto.MaterialInventoryTransactionInfo {
	if in == nil {
		return nil
	}
	m := &proto.MaterialInventoryTransactionInfo{
		Id:              in.ID,
		MaterialInfoID:  in.MaterialInfoID,
		MaterialStoreID: in.MaterialStoreID,
		TransactionType: in.TransactionType,
		Qty:             in.Qty,
		BeforeQTY:       in.BeforeQTY,
		AfterQTY:        in.AfterQTY,
		RefBillNo:       in.RefBillNo,
		CreateUserID:    in.CreateUserID,
		CreateTime:      in.CreateTime.Format("2006-01-02 15:04:05"),
		Remark:          in.Remark,
	}
	if in.MaterialInfo != nil {
		m.MaterialNo = in.MaterialInfo.MaterialNo
	}
	if in.MaterialStore != nil {
		m.MaterialStoreCode = in.MaterialStore.Code
	}
	return m
}
