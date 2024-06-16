package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/system/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type SystemEventTriggerProvider struct {
	proto.UnimplementedSystemEventTriggerServer
}

func (u *SystemEventTriggerProvider) Add(ctx context.Context, in *proto.SystemEventTriggerInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	id, err := logic.CreateSystemEventTrigger(model.PBToSystemEventTrigger(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}
