package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 标签打印队列
type LabelPrintQueue struct {
	ModelID
	TaskNo              string                      `gorm:"size:50;comment:任务编号"`
	PrintTimes          int32                       `gorm:"comment:打印次数"`
	FilePath            string                      `gorm:"size:1000;comment:模板文件"`
	PrintCopies         int32                       `gorm:"comment:打印份数"`
	CreateTime          time.Time                   `gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID        string                      `gorm:"size:36;comment:创建人员ID"`
	PrinterID           string                      `gorm:"size:36;comment:打印机ID"`
	Printer             *Printer                    `gorm:"constraint:OnDelete:CASCADE"`
	CurrentState        string                      `gorm:"size:50;comment:当前状态"`
	LastUpdateTime      time.Time                   `gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	RemoteServiceTaskID *string                     `gorm:"size:36;comment:远程任务ID"`
	RemoteServiceTask   *RemoteServiceTask          `gorm:"constraint:OnDelete:SET NULL"` //远程任务
	LabelParameters     []*LabelPrintQueueParameter `gorm:"constraint:OnDelete:CASCADE"`  //标签打印队列参数
	PrintExecutions     []*LabelPrintQueueExecution `gorm:"constraint:OnDelete:CASCADE"`  //标签打印队列执行
}

// 标签打印队列参数
type LabelPrintQueueParameter struct {
	ModelID
	LabelPrintQueueID string `gorm:"size:36;comment:打印队列ID"`
	Name              string `gorm:"size:50;comment:名称"`
	FixedValue        string `gorm:"size:1000;comment:设定值"`
	Remark            string `gorm:"size:500;comment:备注"`
}

// 标签打印队列执行
type LabelPrintQueueExecution struct {
	ModelID
	LabelPrintQueueID string    `gorm:"size:36;comment:打印队列ID"`
	Success           bool      `gorm:"comment:是否成功"`
	FailureReason     string    `gorm:"size:5000;comment:失败原因"`
	CreateTime        time.Time `gorm:"autoCreateTime:nano;comment:创建时间"`
}

func PBToLabelPrintQueues(in []*proto.LabelPrintQueueInfo) []*LabelPrintQueue {
	var result []*LabelPrintQueue
	for _, c := range in {
		result = append(result, PBToLabelPrintQueue(c))
	}
	return result
}

func PBToLabelPrintQueue(in *proto.LabelPrintQueueInfo) *LabelPrintQueue {
	if in == nil {
		return nil
	}

	var remoteServiceTaskID *string
	if in.RemoteServiceTaskID != "" {
		remoteServiceTaskID = &in.RemoteServiceTaskID
	}

	return &LabelPrintQueue{
		ModelID:             ModelID{ID: in.Id},
		TaskNo:              in.TaskNo,
		PrintTimes:          in.PrintTimes,
		FilePath:            in.FilePath,
		PrintCopies:         in.PrintCopies,
		CreateUserID:        in.CreateUserID,
		PrinterID:           in.PrinterID,
		CurrentState:        in.CurrentState,
		RemoteServiceTaskID: remoteServiceTaskID,
		LabelParameters:     PBToLabelPrintQueueParameters(in.LabelParameters),
		PrintExecutions:     PBToLabelPrintQueueExecutions(in.PrintExecutions),
	}
}

func LabelPrintQueuesToPB(in []*LabelPrintQueue) []*proto.LabelPrintQueueInfo {
	var list []*proto.LabelPrintQueueInfo
	for _, f := range in {
		list = append(list, LabelPrintQueueToPB(f))
	}
	return list
}

func LabelPrintQueueToPB(in *LabelPrintQueue) *proto.LabelPrintQueueInfo {
	if in == nil {
		return nil
	}

	var remoteServiceTaskID string
	if in.RemoteServiceTaskID != nil {
		remoteServiceTaskID = *in.RemoteServiceTaskID
	}

	m := &proto.LabelPrintQueueInfo{
		Id:                  in.ID,
		TaskNo:              in.TaskNo,
		PrintTimes:          in.PrintTimes,
		FilePath:            in.FilePath,
		PrintCopies:         in.PrintCopies,
		CreateTime:          utils.FormatTime(in.CreateTime),
		CreateUserID:        in.CreateUserID,
		PrinterID:           in.PrinterID,
		Printer:             PrinterToPB(in.Printer),
		CurrentState:        in.CurrentState,
		LastUpdateTime:      utils.FormatTime(in.LastUpdateTime),
		RemoteServiceTaskID: remoteServiceTaskID,
		RemoteServiceTask:   RemoteServiceTaskToPB(in.RemoteServiceTask),
		LabelParameters:     LabelPrintQueueParametersToPB(in.LabelParameters),
		PrintExecutions:     LabelPrintQueueExecutionsToPB(in.PrintExecutions),
	}
	return m
}

func PBToLabelPrintQueueParameters(in []*proto.LabelPrintQueueParameterInfo) []*LabelPrintQueueParameter {
	var result []*LabelPrintQueueParameter
	for _, c := range in {
		result = append(result, PBToLabelPrintQueueParameter(c))
	}
	return result
}

func PBToLabelPrintQueueParameter(in *proto.LabelPrintQueueParameterInfo) *LabelPrintQueueParameter {
	if in == nil {
		return nil
	}

	return &LabelPrintQueueParameter{
		ModelID:           ModelID{ID: in.Id},
		LabelPrintQueueID: in.LabelPrintQueueID,
		Name:              in.Name,
		FixedValue:        in.FixedValue,
		Remark:            in.Remark,
	}
}

func LabelPrintQueueParametersToPB(in []*LabelPrintQueueParameter) []*proto.LabelPrintQueueParameterInfo {
	var list []*proto.LabelPrintQueueParameterInfo
	for _, f := range in {
		list = append(list, LabelPrintQueueParameterToPB(f))
	}
	return list
}

func LabelPrintQueueParameterToPB(in *LabelPrintQueueParameter) *proto.LabelPrintQueueParameterInfo {
	if in == nil {
		return nil
	}

	m := &proto.LabelPrintQueueParameterInfo{
		Id:                in.ID,
		LabelPrintQueueID: in.LabelPrintQueueID,
		Name:              in.Name,
		FixedValue:        in.FixedValue,
		Remark:            in.Remark,
	}
	return m
}

func PBToLabelPrintQueueExecutions(in []*proto.LabelPrintQueueExecutionInfo) []*LabelPrintQueueExecution {
	var result []*LabelPrintQueueExecution
	for _, c := range in {
		result = append(result, PBToLabelPrintQueueExecution(c))
	}
	return result
}

func PBToLabelPrintQueueExecution(in *proto.LabelPrintQueueExecutionInfo) *LabelPrintQueueExecution {
	if in == nil {
		return nil
	}

	return &LabelPrintQueueExecution{
		ModelID:           ModelID{ID: in.Id},
		LabelPrintQueueID: in.LabelPrintQueueID,
		Success:           in.Success,
		FailureReason:     in.FailureReason,
	}
}

func LabelPrintQueueExecutionsToPB(in []*LabelPrintQueueExecution) []*proto.LabelPrintQueueExecutionInfo {
	var list []*proto.LabelPrintQueueExecutionInfo
	for _, f := range in {
		list = append(list, LabelPrintQueueExecutionToPB(f))
	}
	return list
}

func LabelPrintQueueExecutionToPB(in *LabelPrintQueueExecution) *proto.LabelPrintQueueExecutionInfo {
	if in == nil {
		return nil
	}

	m := &proto.LabelPrintQueueExecutionInfo{
		Id:                in.ID,
		LabelPrintQueueID: in.LabelPrintQueueID,
		Success:           in.Success,
		FailureReason:     in.FailureReason,
		CreateTime:        utils.FormatTime(in.CreateTime),
	}
	return m
}
