package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 系统事件
type SystemEvent struct {
	ModelID
	Code                     string                     `json:"code" gorm:"size:100;comment:代号"`
	Description              string                     `json:"description" gorm:"size:1000;comment:描述"`
	Enable                   bool                       `json:"enable" gorm:"comment:是否启用"`
	SystemEventParameters    []*SystemEventParameter    `json:"systemEventParameters" gorm:"constraint:OnDelete:CASCADE;"`    //系统事件参数
	SystemEventSubscriptions []*SystemEventSubscription `json:"systemEventSubscriptions" gorm:"constraint:OnDelete:CASCADE;"` //系统事件订阅
}

// 系统事件参数
type SystemEventParameter struct {
	ModelID
	SystemEventID string `json:"systemEventID" gorm:"index;size:36;comment:系统事件 ID"`
	DataType      string `json:"dataType" gorm:"size:100;comment:数据类型"`
	Name          string `json:"name" gorm:"size:100;comment:名称"`
	Description   string `json:"description" gorm:"size:100;comment:描述"`
	Value         string `json:"value" gorm:"size:1000;comment:参数值"`
}

// 系统事件订阅
type SystemEventSubscription struct {
	ModelID
	SystemEventID       string             `json:"systemEventID" gorm:"index;size:36;comment:系统事件ID"`
	RemoteServiceTaskID string             `json:"remoteServiceTaskID" gorm:"size:36;comment:远程服务任务ID"`
	RemoteServiceTask   *RemoteServiceTask `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToSystemEvents(in []*proto.SystemEventInfo) []*SystemEvent {
	var result []*SystemEvent
	for _, c := range in {
		result = append(result, PBToSystemEvent(c))
	}
	return result
}

func PBToSystemEvent(in *proto.SystemEventInfo) *SystemEvent {
	if in == nil {
		return nil
	}
	return &SystemEvent{
		ModelID:                  ModelID{ID: in.Id},
		Code:                     in.Code,
		Description:              in.Description,
		Enable:                   in.Enable,
		SystemEventParameters:    PBToSystemEventParameters(in.SystemEventParameters),
		SystemEventSubscriptions: PBToSystemEventSubscriptions(in.SystemEventSubscriptions),
	}
}

func PBToSystemEventParameters(in []*proto.SystemEventParameterInfo) []*SystemEventParameter {
	var result []*SystemEventParameter
	for _, c := range in {
		result = append(result, PBToSystemEventParameter(c))
	}
	return result
}

func PBToSystemEventParameter(in *proto.SystemEventParameterInfo) *SystemEventParameter {
	if in == nil {
		return nil
	}
	return &SystemEventParameter{
		ModelID:     ModelID{ID: in.Id},
		DataType:    in.DataType,
		Name:        in.Name,
		Description: in.Description,
		Value:       in.Value,
		// SystemEventID: in.SystemEventID,
	}
}

func SystemEventsToPB(in []*SystemEvent) []*proto.SystemEventInfo {
	var list []*proto.SystemEventInfo
	for _, f := range in {
		list = append(list, SystemEventToPB(f))
	}
	return list
}

func SystemEventToPB(in *SystemEvent) *proto.SystemEventInfo {
	if in == nil {
		return nil
	}
	m := &proto.SystemEventInfo{
		SystemEventParameters:    SystemEventParametersToPB(in.SystemEventParameters),
		Id:                       in.ID,
		Code:                     in.Code,
		Description:              in.Description,
		Enable:                   in.Enable,
		SystemEventSubscriptions: SystemEventSubscriptionsToPB(in.SystemEventSubscriptions),
	}
	return m
}

func SystemEventParametersToPB(in []*SystemEventParameter) []*proto.SystemEventParameterInfo {
	var list []*proto.SystemEventParameterInfo
	for _, f := range in {
		list = append(list, SystemEventParameterToPB(f))
	}
	return list
}

func SystemEventParameterToPB(in *SystemEventParameter) *proto.SystemEventParameterInfo {
	if in == nil {
		return nil
	}
	m := &proto.SystemEventParameterInfo{
		Id:            in.ID,
		DataType:      in.DataType,
		Name:          in.Name,
		Description:   in.Description,
		Value:         in.Value,
		SystemEventID: in.SystemEventID,
	}
	return m
}

func PBToSystemEventSubscriptions(in []*proto.SystemEventSubscriptionInfo) []*SystemEventSubscription {
	var result []*SystemEventSubscription
	for _, c := range in {
		result = append(result, PBToSystemEventSubscription(c))
	}
	return result
}

func PBToSystemEventSubscription(in *proto.SystemEventSubscriptionInfo) *SystemEventSubscription {
	if in == nil {
		return nil
	}
	return &SystemEventSubscription{
		// SystemEventID:       in.SystemEventID,
		RemoteServiceTaskID: in.RemoteServiceTaskID,
	}
}

func SystemEventSubscriptionsToPB(in []*SystemEventSubscription) []*proto.SystemEventSubscriptionInfo {
	var list []*proto.SystemEventSubscriptionInfo
	for _, f := range in {
		list = append(list, SystemEventSubscriptionToPB(f))
	}
	return list
}

func SystemEventSubscriptionToPB(in *SystemEventSubscription) *proto.SystemEventSubscriptionInfo {
	if in == nil {
		return nil
	}
	m := &proto.SystemEventSubscriptionInfo{
		SystemEventID:       in.SystemEventID,
		RemoteServiceTaskID: in.RemoteServiceTaskID,
	}
	return m
}
