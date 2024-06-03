package model

import (
	"database/sql"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 远程服务任务
type RemoteServiceTask struct {
	ModelID
	Code                        string                        `json:"code" gorm:"size:100;comment:代号"`
	Description                 string                        `json:"description" gorm:"size:1000;comment:描述"`
	InvokeMethod                string                        `json:"invokeMethod" gorm:"size:100;comment:调用方法"`
	CallbackMethod              string                        `json:"callbackMethod" gorm:"size:100;comment:回调方法"`
	FailureMeasure              int32                         `json:"failureMeasure" gorm:"comment:失败措施"`
	MaximumInvokeCount          int32                         `json:"maximumInvokeCount" gorm:"comment:调用计数"`
	RoutineInvoke               bool                          `json:"routineInvoke" gorm:"comment:常规调用"`
	RegularInvoke               bool                          `json:"regularInvoke" gorm:"comment:定期调用"`
	StartTime                   sql.NullTime                  `json:"startTime" gorm:"comment:起始时间"`
	Interval                    int32                         `json:"interval" gorm:"comment:运行间隔"`
	FinishTime                  sql.NullTime                  `json:"finishTime" gorm:"comment:结束时间"`
	LastInvokeTime              sql.NullTime                  `json:"lastInvokeTime" gorm:"comment:最后调用时间"`
	RemoteServiceID             string                        `json:"remoteServiceID" gorm:"index;comment:远程服务ID"`
	RemoteService               *RemoteService                `json:"remoteService" gorm:"constraint:OnDelete:CASCADE"`                //远程服务
	RemoteServiceTaskParameters []*RemoteServiceTaskParameter `json:"remoteServiceTaskParameters" gorm:"constraint:OnDelete:CASCADE;"` // 远程服务任务参数
}

// 远程任务参数
type RemoteServiceTaskParameter struct {
	ModelID
	RemoteServiceTaskID string `json:"remoteServiceTaskID" gorm:"index;comment:远程服务任务ID"`
	DataType            string `json:"dataType" gorm:"size:100;comment:数据类型"`
	Name                string `json:"name" gorm:"size:100;comment:名称"`
	Description         string `json:"description" gorm:"size:100;comment:描述"`
	Value               string `json:"value" gorm:"size:1000;comment:参数值"`
}

func PBToRemoteServiceTasks(in []*proto.RemoteServiceTaskInfo) []*RemoteServiceTask {
	var result []*RemoteServiceTask
	for _, c := range in {
		result = append(result, PBToRemoteServiceTask(c))
	}
	return result
}

func PBToRemoteServiceTask(in *proto.RemoteServiceTaskInfo) *RemoteServiceTask {
	if in == nil {
		return nil
	}
	return &RemoteServiceTask{
		RemoteServiceTaskParameters: PBToRemoteServiceTaskParameters(in.RemoteServiceTaskParameters),
		ModelID:                     ModelID{ID: in.Id},
		Code:                        in.Code,
		Description:                 in.Description,
		RemoteServiceID:             in.RemoteServiceID,
		InvokeMethod:                in.InvokeMethod,
		CallbackMethod:              in.CallbackMethod,
		FailureMeasure:              in.FailureMeasure,
		MaximumInvokeCount:          in.MaximumInvokeCount,
		RoutineInvoke:               in.RoutineInvoke,
		RegularInvoke:               in.RegularInvoke,
		StartTime:                   utils.ParseSqlNullDate(in.StartTime),
		Interval:                    in.Interval,
		FinishTime:                  utils.ParseSqlNullDate(in.FinishTime),
		LastInvokeTime:              utils.ParseSqlNullTime(in.LastInvokeTime),
	}
}

func PBToRemoteServiceTaskParameters(in []*proto.RemoteServiceTaskParameterInfo) []*RemoteServiceTaskParameter {
	var result []*RemoteServiceTaskParameter
	for _, c := range in {
		result = append(result, PBToRemoteServiceTaskParameter(c))
	}
	return result
}

func PBToRemoteServiceTaskParameter(in *proto.RemoteServiceTaskParameterInfo) *RemoteServiceTaskParameter {
	if in == nil {
		return nil
	}
	return &RemoteServiceTaskParameter{
		ModelID:             ModelID{ID: in.Id},
		DataType:            in.DataType,
		Name:                in.Name,
		Description:         in.Description,
		Value:               in.Value,
		RemoteServiceTaskID: in.RemoteServiceTaskID,
	}
}

func RemoteServiceTasksToPB(in []*RemoteServiceTask) []*proto.RemoteServiceTaskInfo {
	var list []*proto.RemoteServiceTaskInfo
	for _, f := range in {
		list = append(list, RemoteServiceTaskToPB(f))
	}
	return list
}

func RemoteServiceTaskToPB(in *RemoteServiceTask) *proto.RemoteServiceTaskInfo {
	if in == nil {
		return nil
	}
	remoteServiceName := ""
	if in.RemoteService != nil {
		remoteServiceName = in.RemoteService.Name
	}
	m := &proto.RemoteServiceTaskInfo{
		RemoteServiceTaskParameters: RemoteServiceTaskParametersToPB(in.RemoteServiceTaskParameters),
		Id:                          in.ID,
		Code:                        in.Code,
		Description:                 in.Description,
		RemoteServiceID:             in.RemoteServiceID,
		RemoteServiceName:           remoteServiceName,
		InvokeMethod:                in.InvokeMethod,
		CallbackMethod:              in.CallbackMethod,
		FailureMeasure:              in.FailureMeasure,
		MaximumInvokeCount:          in.MaximumInvokeCount,
		RoutineInvoke:               in.RoutineInvoke,
		RegularInvoke:               in.RegularInvoke,
		StartTime:                   utils.FormatSqlNullDate(in.StartTime),
		Interval:                    in.Interval,
		FinishTime:                  utils.FormatSqlNullDate(in.FinishTime),
		LastInvokeTime:              utils.FormatSqlNullTime(in.LastInvokeTime),
	}
	return m
}

func RemoteServiceTaskParametersToPB(in []*RemoteServiceTaskParameter) []*proto.RemoteServiceTaskParameterInfo {
	var list []*proto.RemoteServiceTaskParameterInfo
	for _, f := range in {
		list = append(list, RemoteServiceTaskParameterToPB(f))
	}
	return list
}

func RemoteServiceTaskParameterToPB(in *RemoteServiceTaskParameter) *proto.RemoteServiceTaskParameterInfo {
	if in == nil {
		return nil
	}
	m := &proto.RemoteServiceTaskParameterInfo{
		Id:                  in.ID,
		DataType:            in.DataType,
		Name:                in.Name,
		Description:         in.Description,
		Value:               in.Value,
		RemoteServiceTaskID: in.RemoteServiceTaskID,
	}
	return m
}
