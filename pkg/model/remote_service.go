package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 远程服务
type RemoteService struct {
	ModelID
	Name          string `json:"name" gorm:"size:100;comment:名称"`
	Enable        bool   `json:"enable" gorm:"comment:是否启用"`
	Address       string `json:"address" gorm:"size:1000;comment:服务地址"`
	Headers       string `json:"headers" gorm:"size:-1;comment:请求头"`
	Timeout       int32  `json:"timeout" gorm:"comment:超时时间"`
	UseCredential bool   `json:"useCredential" gorm:"comment:使用凭证"`
	UserName      string `json:"userName" gorm:"size:100;comment:用户名"`
	Password      string `json:"password" gorm:"size:100;comment:密码"`
}

func PBToRemoteServices(in []*proto.RemoteServiceInfo) []*RemoteService {
	var result []*RemoteService
	for _, c := range in {
		result = append(result, PBToRemoteService(c))
	}
	return result
}

func PBToRemoteService(in *proto.RemoteServiceInfo) *RemoteService {
	if in == nil {
		return nil
	}
	return &RemoteService{
		ModelID:       ModelID{ID: in.Id},
		Name:          in.Name,
		Enable:        in.Enable,
		Address:       in.Address,
		Headers:       in.Headers,
		Timeout:       in.Timeout,
		UseCredential: in.UseCredential,
		UserName:      in.UserName,
		Password:      in.Password,
	}
}

func RemoteServicesToPB(in []*RemoteService) []*proto.RemoteServiceInfo {
	var list []*proto.RemoteServiceInfo
	for _, f := range in {
		list = append(list, RemoteServiceToPB(f))
	}
	return list
}

func RemoteServiceToPB(in *RemoteService) *proto.RemoteServiceInfo {
	if in == nil {
		return nil
	}
	m := &proto.RemoteServiceInfo{
		Id:            in.ID,
		Name:          in.Name,
		Enable:        in.Enable,
		Address:       in.Address,
		Headers:       in.Headers,
		Timeout:       in.Timeout,
		UseCredential: in.UseCredential,
		UserName:      in.UserName,
		Password:      in.Password,
	}
	return m
}
