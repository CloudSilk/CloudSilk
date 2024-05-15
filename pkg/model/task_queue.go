package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 任务队列
type TaskQueue struct {
	ModelID
	Name                string                `json:"name" gorm:"size:100;comment:名称"`
	Description         string                `json:"description" gorm:"size:1000;comment:描述"`
	Identity            string                `json:"identity" gorm:"size:100;comment:身份标识"`
	Enable              bool                  `json:"enable" gorm:"comment:是否启用"`
	Interval            int32                 `json:"interval" gorm:"comment:运行间隔"`
	RunningState        string                `json:"runningState" gorm:"size:-1;comment:运行状态"`
	CauseOfState        string                `json:"causeOfState" gorm:"size:-1;comment:状态分析"`
	JobCount            int32                 `json:"jobCount" gorm:"comment:排队数量"`
	TaskQueueParameters []*TaskQueueParameter `json:"taskQueueParameters" gorm:"constraint:OnDelete:CASCADE"` // 任务队列参数
}

type TaskQueueParameter struct {
	ModelID
	TaskQueueID string `json:"taskQueueID" gorm:"index;size:36;comment:任务队列ID"`
	Key         string `json:"key" gorm:"size:100;comment:键"`
	Value       string `json:"value" gorm:"size:8000;comment:值"`
	Remark      string `json:"remark" gorm:"size:1000;comment:备注"`
}

func PBToTaskQueues(in []*proto.TaskQueueInfo) []*TaskQueue {
	var result []*TaskQueue
	for _, c := range in {
		result = append(result, PBToTaskQueue(c))
	}
	return result
}

func PBToTaskQueue(in *proto.TaskQueueInfo) *TaskQueue {
	if in == nil {
		return nil
	}
	return &TaskQueue{
		ModelID:             ModelID{ID: in.Id},
		Name:                in.Name,
		Description:         in.Description,
		Identity:            in.Identity,
		Enable:              in.Enable,
		Interval:            in.Interval,
		RunningState:        in.RunningState,
		CauseOfState:        in.CauseOfState,
		JobCount:            in.JobCount,
		TaskQueueParameters: PBToTaskQueueParameters(in.TaskQueueParameters),
	}
}

func PBToTaskQueueParameters(in []*proto.TaskQueueParameterInfo) []*TaskQueueParameter {
	var result []*TaskQueueParameter
	for _, c := range in {
		result = append(result, PBToTaskQueueParameter(c))
	}
	return result
}

func PBToTaskQueueParameter(in *proto.TaskQueueParameterInfo) *TaskQueueParameter {
	if in == nil {
		return nil
	}
	return &TaskQueueParameter{
		ModelID:     ModelID{ID: in.Id},
		Key:         in.Key,
		Value:       in.Value,
		Remark:      in.Remark,
		TaskQueueID: in.TaskQueueID,
	}
}

func TaskQueuesToPB(in []*TaskQueue) []*proto.TaskQueueInfo {
	var list []*proto.TaskQueueInfo
	for _, f := range in {
		list = append(list, TaskQueueToPB(f))
	}
	return list
}

func TaskQueueToPB(in *TaskQueue) *proto.TaskQueueInfo {
	if in == nil {
		return nil
	}
	m := &proto.TaskQueueInfo{
		TaskQueueParameters: TaskQueueParametersToPB(in.TaskQueueParameters),
		Id:                  in.ID,
		Name:                in.Name,
		Description:         in.Description,
		Identity:            in.Identity,
		Enable:              in.Enable,
		Interval:            in.Interval,
		RunningState:        in.RunningState,
		CauseOfState:        in.CauseOfState,
		JobCount:            in.JobCount,
	}
	return m
}

func TaskQueueParametersToPB(in []*TaskQueueParameter) []*proto.TaskQueueParameterInfo {
	var list []*proto.TaskQueueParameterInfo
	for _, f := range in {
		list = append(list, TaskQueueParameterToPB(f))
	}
	return list
}

func TaskQueueParameterToPB(in *TaskQueueParameter) *proto.TaskQueueParameterInfo {
	if in == nil {
		return nil
	}
	m := &proto.TaskQueueParameterInfo{
		Id:          in.ID,
		Key:         in.Key,
		Value:       in.Value,
		Remark:      in.Remark,
		TaskQueueID: in.TaskQueueID,
	}
	return m
}
