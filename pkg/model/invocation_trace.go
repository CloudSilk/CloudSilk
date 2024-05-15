package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 接口调用日志
type InvocationTrace struct {
	ModelID
	IPAddress      string    `json:"iPAddress" gorm:"size:30;comment:IP地址"`
	ControllerName string    `json:"controllerName" gorm:"size:200;comment:控制器"`
	ActionName     string    `json:"actionName" gorm:"size:200;comment:路由"`
	RequestUrl     string    `json:"requestUrl" gorm:"size:1000;comment:请求地址"`
	RequestTime    time.Time `json:"requestTime" gorm:"autoCreateTime:nano;comment:请求时间"`
	Duration       int32     `json:"duration" gorm:"comment:耗时"`
	RequestText    string    `json:"requestText" gorm:"comment:请求文本"`
	ResponseText   string    `json:"responseText" gorm:"comment:响应文本"`
	ResponseCode   int32     `json:"responseCode" gorm:"comment:响应码"`
	Annotation     string    `json:"annotation" gorm:"size:1000;comment:注释"`
}

func PBToInvocationTraces(in []*proto.InvocationTraceInfo) []*InvocationTrace {
	var result []*InvocationTrace
	for _, c := range in {
		result = append(result, PBToInvocationTrace(c))
	}
	return result
}

func PBToInvocationTrace(in *proto.InvocationTraceInfo) *InvocationTrace {
	if in == nil {
		return nil
	}
	return &InvocationTrace{
		ModelID:        ModelID{ID: in.Id},
		IPAddress:      in.IPAddress,
		ControllerName: in.ControllerName,
		ActionName:     in.ActionName,
		RequestUrl:     in.RequestUrl,
		// RequestTime:    utils.ParseTime(in.RequestTime),
		Duration:     in.Duration,
		RequestText:  in.RequestText,
		ResponseText: in.ResponseText,
		ResponseCode: in.ResponseCode,
		Annotation:   in.Annotation,
	}
}

func InvocationTracesToPB(in []*InvocationTrace) []*proto.InvocationTraceInfo {
	var list []*proto.InvocationTraceInfo
	for _, f := range in {
		list = append(list, InvocationTraceToPB(f))
	}
	return list
}

func InvocationTraceToPB(in *InvocationTrace) *proto.InvocationTraceInfo {
	if in == nil {
		return nil
	}
	m := &proto.InvocationTraceInfo{
		Id:             in.ID,
		IPAddress:      in.IPAddress,
		ControllerName: in.ControllerName,
		ActionName:     in.ActionName,
		RequestUrl:     in.RequestUrl,
		RequestTime:    utils.FormatTime(in.RequestTime),
		Duration:       in.Duration,
		RequestText:    in.RequestText,
		ResponseText:   in.ResponseText,
		ResponseCode:   in.ResponseCode,
		Annotation:     in.Annotation,
	}
	return m
}
