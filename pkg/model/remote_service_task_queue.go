package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 远程任务队列
type RemoteServiceTaskQueue struct {
	ModelID
	TaskNo                           string                             `json:"taskNo" gorm:";size:100;comment:任务编号"`
	CreateTime                       time.Time                          `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	RequestURL                       string                             `json:"requestURL" gorm:"size:-1;comment:请求路径"`
	RequestText                      string                             `json:"requestText" gorm:"size:-1;comment:请求内容"`
	ResponseText                     string                             `json:"responseText" gorm:"size:-1;comment:响应内容"`
	FinishTime                       time.Time                          `json:"finishTime" gorm:"comment:完成时间"`
	InvokeCount                      int32                              `json:"invokeCount" gorm:"comment:调用计数"`
	CurrentState                     string                             `json:"currentState" gorm:"size:-1;comment:当前状态"`
	TransactionState                 string                             `json:"transactionState" gorm:"size:-1;comment:事务状态"`
	RemoteServiceTaskID              string                             `json:"remoteServiceTaskID" gorm:"index;size:36;comment:远程任务ID"`
	RemoteServiceTask                *RemoteServiceTask                 `json:"remoteServiceTask" gorm:"constraint:OnDelete:CASCADE"`                 //远程任务
	RemoteServiceTaskQueueParameters []*RemoteServiceTaskQueueParameter `json:"remoteServiceTaskQueueParameters" gorm:"constraint:OnDelete:CASCADE;"` //远程任务队列参数
}

// 远程任务队列参数
type RemoteServiceTaskQueueParameter struct {
	ModelID
	RemoteServiceTaskQueueID string `json:"remoteServiceTaskQueueID" gorm:"index;size:36;comment:远程服务任务队列ID"`
	DataType                 string `json:"dataType" gorm:"size:200;comment:数据类型"`
	Name                     string `json:"name" gorm:"size:200;comment:名称"`
	Description              string `json:"description" gorm:"size:2000;comment:备注"`
	Value                    string `json:"value" gorm:"size:2000;comment:值"`
}

func PBToRemoteServiceTaskQueues(in []*proto.RemoteServiceTaskQueueInfo) []*RemoteServiceTaskQueue {
	var result []*RemoteServiceTaskQueue
	for _, c := range in {
		result = append(result, PBToRemoteServiceTaskQueue(c))
	}
	return result
}

func PBToRemoteServiceTaskQueue(in *proto.RemoteServiceTaskQueueInfo) *RemoteServiceTaskQueue {
	if in == nil {
		return nil
	}
	return &RemoteServiceTaskQueue{
		ModelID: ModelID{ID: in.Id},
		TaskNo:  in.TaskNo,
		// CreateTime:                       utils.ParseTime(in.CreateTime),
		RequestURL:   in.RequestURL,
		RequestText:  in.RequestText,
		ResponseText: in.ResponseText,
		// FinishTime:                       utils.ParseTime(in.FinishTime),
		InvokeCount:                      in.InvokeCount,
		CurrentState:                     in.CurrentState,
		TransactionState:                 in.TransactionState,
		RemoteServiceTaskID:              in.RemoteServiceTaskID,
		RemoteServiceTaskQueueParameters: PBToRemoteServiceTaskQueueParameters(in.RemoteServiceTaskQueueParameters),
	}
}

func RemoteServiceTaskQueuesToPB(in []*RemoteServiceTaskQueue) []*proto.RemoteServiceTaskQueueInfo {
	var list []*proto.RemoteServiceTaskQueueInfo
	for _, f := range in {
		list = append(list, RemoteServiceTaskQueueToPB(f))
	}
	return list
}

func RemoteServiceTaskQueueToPB(in *RemoteServiceTaskQueue) *proto.RemoteServiceTaskQueueInfo {
	if in == nil {
		return nil
	}
	remoteServiceTaskName := ""
	remoteServiceName := ""
	if in.RemoteServiceTask != nil {
		remoteServiceTaskName = in.RemoteServiceTask.Description
		if in.RemoteServiceTask.RemoteService != nil {
			remoteServiceName = in.RemoteServiceTask.RemoteService.Name
		}
	}
	m := &proto.RemoteServiceTaskQueueInfo{
		Id:                               in.ID,
		TaskNo:                           in.TaskNo,
		CreateTime:                       utils.FormatTime(in.CreateTime),
		RequestURL:                       in.RequestURL,
		RequestText:                      in.RequestText,
		ResponseText:                     in.ResponseText,
		FinishTime:                       utils.FormatTime(in.FinishTime),
		InvokeCount:                      in.InvokeCount,
		CurrentState:                     in.CurrentState,
		TransactionState:                 in.TransactionState,
		RemoteServiceTaskID:              in.RemoteServiceTaskID,
		RemoteServiceTaskName:            remoteServiceTaskName,
		RemoteServiceName:                remoteServiceName,
		RemoteServiceTaskQueueParameters: RemoteServiceTaskQueueParametersToPB(in.RemoteServiceTaskQueueParameters),
	}
	return m
}

func PBToRemoteServiceTaskQueueParameters(in []*proto.RemoteServiceTaskQueueParameterInfo) []*RemoteServiceTaskQueueParameter {
	var result []*RemoteServiceTaskQueueParameter
	for _, c := range in {
		result = append(result, PBToRemoteServiceTaskQueueParameter(c))
	}
	return result
}

func PBToRemoteServiceTaskQueueParameter(in *proto.RemoteServiceTaskQueueParameterInfo) *RemoteServiceTaskQueueParameter {
	if in == nil {
		return nil
	}
	return &RemoteServiceTaskQueueParameter{
		ModelID:                  ModelID{ID: in.Id},
		DataType:                 in.DataType,
		Name:                     in.Name,
		Description:              in.Description,
		Value:                    in.Value,
		RemoteServiceTaskQueueID: in.RemoteServiceTaskQueueID,
	}
}

func RemoteServiceTaskQueueParametersToPB(in []*RemoteServiceTaskQueueParameter) []*proto.RemoteServiceTaskQueueParameterInfo {
	var list []*proto.RemoteServiceTaskQueueParameterInfo
	for _, f := range in {
		list = append(list, RemoteServiceTaskQueueParameterToPB(f))
	}
	return list
}

func RemoteServiceTaskQueueParameterToPB(in *RemoteServiceTaskQueueParameter) *proto.RemoteServiceTaskQueueParameterInfo {
	if in == nil {
		return nil
	}

	return &proto.RemoteServiceTaskQueueParameterInfo{
		Id:                       in.ID,
		DataType:                 in.DataType,
		Name:                     in.Name,
		Description:              in.Description,
		Value:                    in.Value,
		RemoteServiceTaskQueueID: in.RemoteServiceTaskQueueID,
	}
}
