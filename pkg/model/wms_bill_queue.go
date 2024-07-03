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
	// WMSBillItems     *WMSBillItem   `gorm:""`
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
	return m
}
