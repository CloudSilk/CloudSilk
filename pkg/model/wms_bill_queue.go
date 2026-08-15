package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

type WMSBillQueue struct {
	ModelID
	BillNo           string         `gorm:"size:50;comment:拣货单号"`
	TaskNo           string         `gorm:"size:50;comment:任务单号"`
	MaterialStoreID  string         `gorm:"size:36;comment:物料仓库ID"`
	MaterialStore    *MaterialStore `gorm:"constraint:OnDelete:CASCADE"`
	ProductOrderID   string         `gorm:"size:36;comment:生产工单号ID"`
	ProductOrder     *ProductOrder  `gorm:"constraint:OnDelete:CASCADE"`
	CreateTime       time.Time      `gorm:"autoCreateTime:nano;comment:申请时间"`
	CreateUserID     string         `gorm:"size:36;comment:申请人员ID"`
	CurrentState     string         `gorm:"size:50;comment:当前状态"`
	TransactionState string         `gorm:"size:50;comment:事务状态"`
	LastUpdateTime   time.Time      `gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark           string         `gorm:"size:500;comment:备注"`
	// 行项目（拣货明细）
	WMSBillItems []*WMSBillItem `gorm:"constraint:OnDelete:CASCADE"`
}

// WMSBillItem 拣货单行项目（一行=一个物料的锁定/出库承诺）
type WMSBillItem struct {
	ModelID
	WMSBillQueueID      string         `json:"wmsBillQueueID" gorm:"index;size:36;comment:拣货单ID"`
	WMSBillQueue        *WMSBillQueue  `json:"wmsBillQueue" gorm:"constraint:OnDelete:CASCADE"`
	ProductOrderBomID   string         `json:"productOrderBomID" gorm:"size:36;comment:工单BOM行ID"`
	MaterialInfoID      string         `json:"materialInfoID" gorm:"size:36;comment:物料ID"`
	MaterialInfo        *MaterialInfo  `json:"materialInfo" gorm:"constraint:OnDelete:CASCADE"`
	MaterialNo          string         `json:"materialNo" gorm:"size:100;comment:物料号"`
	MaterialDescription string         `json:"materialDescription" gorm:"size:500;comment:物料描述"`
	MaterialStoreID     string         `json:"materialStoreID" gorm:"size:36;comment:仓库ID"`
	MaterialStore       *MaterialStore `json:"materialStore" gorm:"constraint:OnDelete:CASCADE"`
	RequiredQTY         int64          `json:"requiredQTY" gorm:"comment:需求数量"`
	LockedQTY           int64          `json:"lockedQTY" gorm:"comment:锁定数量"`
	CurrentState        string         `json:"currentState" gorm:"size:50;comment:当前状态(待拣货/已完成/已取消)"`
	Remark              string         `json:"remark" gorm:"size:500;comment:备注"`
}

func PBToWMSBillQueues(in []*proto.WMSBillQueueInfo) []*WMSBillQueue {
	var result []*WMSBillQueue
	for _, c := range in {
		result = append(result, PBToWMSBillQueue(c))
	}
	return result
}

func PBToWMSBillQueue(in *proto.WMSBillQueueInfo) *WMSBillQueue {
	if in == nil {
		return nil
	}

	return &WMSBillQueue{
		ModelID:          ModelID{ID: in.Id},
		BillNo:           in.BillNo,
		TaskNo:           in.TaskNo,
		MaterialStoreID:  in.MaterialStoreID,
		ProductOrderID:   in.ProductOrderID,
		CreateUserID:     in.CreateUserID,
		CurrentState:     in.CurrentState,
		TransactionState: in.TransactionState,
		Remark:           in.Remark,
	}
}

func WMSBillQueuesToPB(in []*WMSBillQueue) []*proto.WMSBillQueueInfo {
	var list []*proto.WMSBillQueueInfo
	for _, f := range in {
		list = append(list, WMSBillQueueToPB(f))
	}
	return list
}

func WMSBillQueueToPB(in *WMSBillQueue) *proto.WMSBillQueueInfo {
	if in == nil {
		return nil
	}

	m := &proto.WMSBillQueueInfo{
		Id:               in.ID,
		BillNo:           in.BillNo,
		TaskNo:           in.TaskNo,
		MaterialStoreID:  in.MaterialStoreID,
		MaterialStore:    MaterialStoreToPB(in.MaterialStore),
		ProductOrderID:   in.ProductOrderID,
		ProductOrder:     ProductOrderToPB(in.ProductOrder),
		CreateTime:       utils.FormatTime(in.CreateTime),
		CreateUserID:     in.CreateUserID,
		CurrentState:     in.CurrentState,
		TransactionState: in.TransactionState,
		LastUpdateTime:   utils.FormatTime(in.LastUpdateTime),
		Remark:           in.Remark,
	}
	for _, item := range in.WMSBillItems {
		m.WMSBillItems = append(m.WMSBillItems, WMSBillItemToPB(item))
	}
	return m
}

func WMSBillItemsToPB(in []*WMSBillItem) []*proto.WMSBillItemInfo {
	var list []*proto.WMSBillItemInfo
	for _, f := range in {
		list = append(list, WMSBillItemToPB(f))
	}
	return list
}

func WMSBillItemToPB(in *WMSBillItem) *proto.WMSBillItemInfo {
	if in == nil {
		return nil
	}
	m := &proto.WMSBillItemInfo{
		Id:                  in.ID,
		WMSBillQueueID:      in.WMSBillQueueID,
		ProductOrderBomID:   in.ProductOrderBomID,
		MaterialInfoID:      in.MaterialInfoID,
		MaterialNo:          in.MaterialNo,
		MaterialDescription: in.MaterialDescription,
		MaterialStoreID:     in.MaterialStoreID,
		RequiredQTY:         in.RequiredQTY,
		LockedQTY:           in.LockedQTY,
		CurrentState:        in.CurrentState,
		Remark:              in.Remark,
	}
	if in.MaterialStore != nil {
		m.MaterialStoreCode = in.MaterialStore.Code
	}
	return m
}
