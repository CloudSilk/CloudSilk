package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

type AGVTaskQueue struct {
	ModelID
	TaskNo           string            `gorm:"size:50;comment:任务编号"`
	CreateTime       time.Time         `gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID     string            `gorm:"size:36;comment:配料人员"`
	ContinueTask     bool              `gorm:"comment:后续任务"`
	CurrentState     string            `gorm:"size:50;comment:当前状态"`
	TransactionState string            `gorm:"size:50;comment:事务状态"`
	LastUpdateTime   time.Time         `gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark           string            `gorm:"size:500;comment:备注"`
	AGVTaskTypeID    *string           `gorm:"size:36;comment:任务类型ID"`
	AGVTaskType      *AGVTaskType      `gorm:"constraint:OnDelete:SET NULL"` //任务类型
	DepartureID      string            `gorm:"size:36;comment:起点库位ID"`
	Departure        *MaterialShelfBin `gorm:"constraint:OnDelete:CASCADE"` //起点库位
	DestinationID    string            `gorm:"size:36;comment:终点库位ID"`
	Destination      *MaterialShelfBin `gorm:"constraint:OnDelete:CASCADE"` //终点库位
}

func PBToAGVTaskQueues(in []*proto.AGVTaskQueueInfo) []*AGVTaskQueue {
	var result []*AGVTaskQueue
	for _, c := range in {
		result = append(result, PBToAGVTaskQueue(c))
	}
	return result
}

func PBToAGVTaskQueue(in *proto.AGVTaskQueueInfo) *AGVTaskQueue {
	if in == nil {
		return nil
	}

	var aGVTaskTypeID *string
	if in.AGVTaskTypeID != "" {
		aGVTaskTypeID = &in.AGVTaskTypeID
	}

	return &AGVTaskQueue{
		ModelID:          ModelID{ID: in.Id},
		TaskNo:           in.TaskNo,
		CreateUserID:     in.CreateUserID,
		ContinueTask:     in.ContinueTask,
		CurrentState:     in.CurrentState,
		TransactionState: in.TransactionState,
		Remark:           in.Remark,
		AGVTaskTypeID:    aGVTaskTypeID,
		DepartureID:      in.DepartureID,
		DestinationID:    in.DestinationID,
	}
}

func AGVTaskQueuesToPB(in []*AGVTaskQueue) []*proto.AGVTaskQueueInfo {
	var list []*proto.AGVTaskQueueInfo
	for _, f := range in {
		list = append(list, AGVTaskQueueToPB(f))
	}
	return list
}

func AGVTaskQueueToPB(in *AGVTaskQueue) *proto.AGVTaskQueueInfo {
	if in == nil {
		return nil
	}

	var aGVTaskTypeID string
	if in.AGVTaskTypeID != nil {
		aGVTaskTypeID = *in.AGVTaskTypeID
	}

	m := &proto.AGVTaskQueueInfo{
		Id:               in.ID,
		TaskNo:           in.TaskNo,
		CreateTime:       utils.FormatTime(in.CreateTime),
		CreateUserID:     in.CreateUserID,
		ContinueTask:     in.ContinueTask,
		CurrentState:     in.CurrentState,
		TransactionState: in.TransactionState,
		LastUpdateTime:   utils.FormatTime(in.LastUpdateTime),
		Remark:           in.Remark,
		AGVTaskTypeID:    aGVTaskTypeID,
		AGVTaskType:      AGVTaskTypeToPB(in.AGVTaskType),
		DepartureID:      in.DepartureID,
		Departure:        MaterialShelfBinToPB(in.Departure),
		DestinationID:    in.DestinationID,
		Destination:      MaterialShelfBinToPB(in.Destination),
	}
	return m
}
