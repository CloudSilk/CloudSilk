package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 系统异常日志
type ExceptionTrace struct {
	ModelID
	Host          string    `json:"host" gorm:"size:100;comment:主机"`
	Source        string    `json:"source" gorm:"size:100;comment:来源"`
	Level         string    `json:"level" gorm:"size:100;comment:等级"`
	Message       string    `json:"message" gorm:"size:-1;comment:异常信息"`
	StackTrace    string    `json:"stackTrace" gorm:"size:-1;comment:堆栈跟踪"`
	ScreenCapture string    `json:"screenCapture" gorm:"size:2000;comment:屏幕截图"`
	TimeReported  time.Time `json:"timeReported" gorm:"autoCreateTime:nano;comment:上报时间"`
	ReportUserID  string    `json:"reportUserID" gorm:"size:36;comment:上报人员ID"`
}

func PBToExceptionTraces(in []*proto.ExceptionTraceInfo) []*ExceptionTrace {
	var result []*ExceptionTrace
	for _, c := range in {
		result = append(result, PBToExceptionTrace(c))
	}
	return result
}

func PBToExceptionTrace(in *proto.ExceptionTraceInfo) *ExceptionTrace {
	if in == nil {
		return nil
	}
	return &ExceptionTrace{
		ModelID:       ModelID{ID: in.Id},
		Host:          in.Host,
		Source:        in.Source,
		Level:         in.Level,
		Message:       in.Message,
		StackTrace:    in.StackTrace,
		ScreenCapture: in.ScreenCapture,
		// TimeReported:  utils.ParseTime(in.TimeReported),
		ReportUserID: in.ReportUserID,
	}
}

func ExceptionTracesToPB(in []*ExceptionTrace) []*proto.ExceptionTraceInfo {
	var list []*proto.ExceptionTraceInfo
	for _, f := range in {
		list = append(list, ExceptionTraceToPB(f))
	}
	return list
}

func ExceptionTraceToPB(in *ExceptionTrace) *proto.ExceptionTraceInfo {
	if in == nil {
		return nil
	}
	m := &proto.ExceptionTraceInfo{
		Id:            in.ID,
		Host:          in.Host,
		Source:        in.Source,
		Level:         in.Level,
		Message:       in.Message,
		StackTrace:    in.StackTrace,
		ScreenCapture: in.ScreenCapture,
		TimeReported:  utils.FormatTime(in.TimeReported),
		ReportUserID:  in.ReportUserID,
	}
	return m
}
