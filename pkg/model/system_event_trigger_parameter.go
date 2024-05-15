package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 系统事件触发参数
type SystemEventTriggerParameter struct {
	ModelID
	DataType             string `json:"dataType" gorm:"comment:数据类型"`
	Name                 string `json:"name" gorm:"comment:名称"`
	Description          string `json:"description" gorm:"comment:描述"`
	Value                string `json:"value" gorm:"comment:参数值"`
	SystemEventTriggerID string `json:"SystemEventTriggerID" gorm:"size:36;comment:系统事件触发ID"`
}

func PBToSystemEventTriggerParameters(in []*proto.SystemEventTriggerParameterInfo) []*SystemEventTriggerParameter {
	var result []*SystemEventTriggerParameter
	for _, c := range in {
		result = append(result, PBToSystemEventTriggerParameter(c))
	}
	return result
}

func PBToSystemEventTriggerParameter(in *proto.SystemEventTriggerParameterInfo) *SystemEventTriggerParameter {
	if in == nil {
		return nil
	}
	return &SystemEventTriggerParameter{
		ModelID:              ModelID{ID: in.Id},
		DataType:             in.DataType,
		Name:                 in.Name,
		Description:          in.Description,
		Value:                in.Value,
		SystemEventTriggerID: in.SystemEventTriggerID,
	}
}

func SystemEventTriggerParametersToPB(in []*SystemEventTriggerParameter) []*proto.SystemEventTriggerParameterInfo {
	var list []*proto.SystemEventTriggerParameterInfo
	for _, f := range in {
		list = append(list, SystemEventTriggerParameterToPB(f))
	}
	return list
}

func SystemEventTriggerParameterToPB(in *SystemEventTriggerParameter) *proto.SystemEventTriggerParameterInfo {
	if in == nil {
		return nil
	}
	m := &proto.SystemEventTriggerParameterInfo{
		Id:                   in.ID,
		DataType:             in.DataType,
		Name:                 in.Name,
		Description:          in.Description,
		Value:                in.Value,
		SystemEventTriggerID: in.SystemEventTriggerID,
	}
	return m
}
