package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 任务队列执行
type TaskQueueExecution struct {
	ModelID
	DataTrace     string     `json:"dataTrace" gorm:"size:-1;comment:数据跟踪"`
	Success       bool       `json:"success" gorm:"comment:是否成功"`
	FailureReason string     `json:"failureReason" gorm:"size:-1;comment:失败原因"`
	CreateTime    time.Time  `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	TaskQueueID   string     `json:"taskQueueID" gorm:"size:36;comment:任务队列ID"`
	TaskQueue     *TaskQueue `json:"taskQueue"` //任务队列
}

func PBToTaskQueueExecutions(in []*proto.TaskQueueExecutionInfo) []*TaskQueueExecution {
	var result []*TaskQueueExecution
	for _, c := range in {
		result = append(result, PBToTaskQueueExecution(c))
	}
	return result
}

func PBToTaskQueueExecution(in *proto.TaskQueueExecutionInfo) *TaskQueueExecution {
	if in == nil {
		return nil
	}
	return &TaskQueueExecution{
		ModelID:       ModelID{ID: in.Id},
		DataTrace:     in.DataTrace,
		Success:       in.Success,
		FailureReason: in.FailureReason,
		// CreateTime:    utils.ParseTime(in.CreateTime),
		TaskQueueID: in.TaskQueueID,
	}
}

func TaskQueueExecutionsToPB(in []*TaskQueueExecution) []*proto.TaskQueueExecutionInfo {
	var list []*proto.TaskQueueExecutionInfo
	for _, f := range in {
		list = append(list, TaskQueueExecutionToPB(f))
	}
	return list
}

func TaskQueueExecutionToPB(in *TaskQueueExecution) *proto.TaskQueueExecutionInfo {
	if in == nil {
		return nil
	}
	taskQueueDescription := ""
	if in.TaskQueue != nil {
		taskQueueDescription = in.TaskQueue.Name
	}
	m := &proto.TaskQueueExecutionInfo{
		Id:                   in.ID,
		DataTrace:            in.DataTrace,
		Success:              in.Success,
		FailureReason:        in.FailureReason,
		CreateTime:           utils.FormatTime(in.CreateTime),
		TaskQueueID:          in.TaskQueueID,
		TaskQueueDescription: taskQueueDescription,
	}
	return m
}
