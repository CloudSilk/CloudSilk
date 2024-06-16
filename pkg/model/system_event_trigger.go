package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 系统事件触发
type SystemEventTrigger struct {
	ModelID
	EventNo                      string                         `json:"eventNo" gorm:"size:100;comment:事件编号"`
	CurrentState                 string                         `json:"currentState" gorm:"size:-1;comment:当前状态"`
	CreateTime                   time.Time                      `json:"createTime" gorm:"autoCreateTime;comment:触发时间"`
	LastUpdateTime               time.Time                      `json:"lastUpdateTime" gorm:"autoCreateTime;comment:状态变更时间"`
	SystemEventID                string                         `json:"systemEventID" gorm:"size:36;comment:系统事件ID"`
	SystemEvent                  *SystemEvent                   `json:"systemEvent" gorm:"constraint:OnDelete:CASCADE"` //系统事件
	SystemEventTriggerParameters []*SystemEventTriggerParameter `gorm:"constraint:OnDelete:CASCADE;"`
	SystemEventTriggerExecutions []*SystemEventTriggerExecution `gorm:"constraint:OnDelete:CASCADE;"`
}
type SystemEventTriggerExecution struct {
	ModelID
	SystemEventTriggerID     string                  `gorm:"index;size:36;comment:系统事件触发ID"`
	SystemEventTrigger       *SystemEventTrigger     `gorm:"constraint:OnDelete:CASCADE;"`
	RemoteServiceTaskQueueID string                  `gorm:"size:36;comment:远程任务队列ID"`
	RemoteServiceTaskQueue   *RemoteServiceTaskQueue `gorm:"constraint:OnDelete:CASCADE;"`
}

func PBToSystemEventTriggers(in []*proto.SystemEventTriggerInfo) []*SystemEventTrigger {
	var result []*SystemEventTrigger
	for _, c := range in {
		result = append(result, PBToSystemEventTrigger(c))
	}
	return result
}

func PBToSystemEventTrigger(in *proto.SystemEventTriggerInfo) *SystemEventTrigger {
	if in == nil {
		return nil
	}
	return &SystemEventTrigger{
		ModelID:       ModelID{ID: in.Id},
		EventNo:       in.EventNo,
		SystemEventID: in.SystemEventID,
		// CreateTime:     utils.ParseTime(in.CreateTime),
		CurrentState: in.CurrentState,
		// LastUpdateTime: utils.ParseTime(in.LastUpdateTime),
		SystemEventTriggerParameters: PBToSystemEventTriggerParameters(in.SystemEventTriggerParameters),
	}
}

func SystemEventTriggersToPB(in []*SystemEventTrigger) []*proto.SystemEventTriggerInfo {
	var list []*proto.SystemEventTriggerInfo
	for _, f := range in {
		list = append(list, SystemEventTriggerToPB(f))
	}
	return list
}

func SystemEventTriggerToPB(in *SystemEventTrigger) *proto.SystemEventTriggerInfo {
	if in == nil {
		return nil
	}
	systemEventName := ""
	if in.SystemEvent != nil {
		systemEventName = in.SystemEvent.Description
	}
	m := &proto.SystemEventTriggerInfo{
		Id:                           in.ID,
		EventNo:                      in.EventNo,
		SystemEventID:                in.SystemEventID,
		SystemEventName:              systemEventName,
		CreateTime:                   utils.FormatTime(in.CreateTime),
		CurrentState:                 in.CurrentState,
		LastUpdateTime:               utils.FormatTime(in.LastUpdateTime),
		SystemEventTriggerParameters: SystemEventTriggerParametersToPB(in.SystemEventTriggerParameters),
	}
	return m
}
