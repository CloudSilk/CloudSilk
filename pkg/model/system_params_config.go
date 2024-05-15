package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 系统参数
type SystemParamsConfig struct {
	ModelID
	Code   string `json:"code" gorm:"size:100;comment:代号"`
	Key    string `json:"key" gorm:"size:100;comment:项"`
	Value  string `json:"value" gorm:"size:8000;comment:值"`
	Remark string `json:"remark" gorm:"size:1000;comment:备注"`
}

func PBToSystemParamsConfigs(in []*proto.SystemParamsConfigInfo) []*SystemParamsConfig {
	var result []*SystemParamsConfig
	for _, c := range in {
		result = append(result, PBToSystemParamsConfig(c))
	}
	return result
}

func PBToSystemParamsConfig(in *proto.SystemParamsConfigInfo) *SystemParamsConfig {
	if in == nil {
		return nil
	}
	return &SystemParamsConfig{
		ModelID: ModelID{ID: in.Id},
		Code:    in.Code,
		Key:     in.Key,
		Value:   in.Value,
		Remark:  in.Remark,
	}
}

func SystemParamsConfigsToPB(in []*SystemParamsConfig) []*proto.SystemParamsConfigInfo {
	var list []*proto.SystemParamsConfigInfo
	for _, f := range in {
		list = append(list, SystemParamsConfigToPB(f))
	}
	return list
}

func SystemParamsConfigToPB(in *SystemParamsConfig) *proto.SystemParamsConfigInfo {
	if in == nil {
		return nil
	}
	m := &proto.SystemParamsConfigInfo{
		Id:     in.ID,
		Code:   in.Code,
		Key:    in.Key,
		Value:  in.Value,
		Remark: in.Remark,
	}
	return m
}
