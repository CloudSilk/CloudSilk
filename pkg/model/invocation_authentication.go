package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 接口认证
type InvocationAuthentication struct {
	ModelID
	Name      string `json:"name" gorm:"size:100;comment:名称"`
	IPAddress string `json:"ipAddress" gorm:"size:30;comment:IP地址"`
	APIKey    string `json:"aPIKey" gorm:"size:100;comment:API密钥"`
	Remark    string `json:"remark" gorm:"size:1000;comment:备注"`
}

func PBToInvocationAuthentications(in []*proto.InvocationAuthenticationInfo) []*InvocationAuthentication {
	var result []*InvocationAuthentication
	for _, c := range in {
		result = append(result, PBToInvocationAuthentication(c))
	}
	return result
}

func PBToInvocationAuthentication(in *proto.InvocationAuthenticationInfo) *InvocationAuthentication {
	if in == nil {
		return nil
	}
	return &InvocationAuthentication{
		ModelID:   ModelID{ID: in.Id},
		Name:      in.Name,
		IPAddress: in.IPAddress,
		APIKey:    in.APIKey,
		Remark:    in.Remark,
	}
}

func InvocationAuthenticationsToPB(in []*InvocationAuthentication) []*proto.InvocationAuthenticationInfo {
	var list []*proto.InvocationAuthenticationInfo
	for _, f := range in {
		list = append(list, InvocationAuthenticationToPB(f))
	}
	return list
}

func InvocationAuthenticationToPB(in *InvocationAuthentication) *proto.InvocationAuthenticationInfo {
	if in == nil {
		return nil
	}
	m := &proto.InvocationAuthenticationInfo{
		Id:        in.ID,
		Name:      in.Name,
		IPAddress: in.IPAddress,
		APIKey:    in.APIKey,
		Remark:    in.Remark,
	}
	return m
}
