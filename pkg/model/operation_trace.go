package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 后台操作日志
type OperationTrace struct {
	ModelID
	OperateUserID  string    `json:"operateUserID" gorm:"size:36;comment:操作人员ID"`
	IPAddress      string    `json:"iPAddress" gorm:"size:30;comment:IP地址"`
	OperateTime    time.Time `json:"operateTime" gorm:"autoCreateTime:nano;comment:操作时间"`
	ControllerName string    `json:"controllerName" gorm:"size:200;comment:模块"`
	ActionName     string    `json:"actionName" gorm:"size:200;comment:操作"`
	RequestContent string    `json:"requestContent" gorm:"comment:提交内容"`
	Annotation     string    `json:"annotation" gorm:"size:1000;comment:注释"`
}

func PBToOperationTraces(in []*proto.OperationTraceInfo) []*OperationTrace {
	var result []*OperationTrace
	for _, c := range in {
		result = append(result, PBToOperationTrace(c))
	}
	return result
}

func PBToOperationTrace(in *proto.OperationTraceInfo) *OperationTrace {
	if in == nil {
		return nil
	}
	return &OperationTrace{
		ModelID:       ModelID{ID: in.Id},
		OperateUserID: in.OperateUserID,
		IPAddress:     in.IPAddress,
		// OperateTime:    utils.ParseTime(in.OperateTime),
		ControllerName: in.ControllerName,
		ActionName:     in.ActionName,
		RequestContent: in.RequestContent,
		Annotation:     in.Annotation,
	}
}

func OperationTracesToPB(in []*OperationTrace) []*proto.OperationTraceInfo {
	var list []*proto.OperationTraceInfo
	for _, f := range in {
		list = append(list, OperationTraceToPB(f))
	}
	return list
}

func OperationTraceToPB(in *OperationTrace) *proto.OperationTraceInfo {
	if in == nil {
		return nil
	}
	m := &proto.OperationTraceInfo{
		Id:             in.ID,
		OperateUserID:  in.OperateUserID,
		IPAddress:      in.IPAddress,
		OperateTime:    utils.FormatTime(in.OperateTime),
		ControllerName: in.ControllerName,
		ActionName:     in.ActionName,
		RequestContent: in.RequestContent,
		Annotation:     in.Annotation,
	}
	return m
}
