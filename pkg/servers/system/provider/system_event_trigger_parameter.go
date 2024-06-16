package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/system/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type SystemEventTriggerParameterProvider struct {
	proto.UnimplementedSystemEventTriggerParameterServer
}

func (u *SystemEventTriggerParameterProvider) Add(ctx context.Context, in *proto.SystemEventTriggerParameterInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	id, err := logic.CreateSystemEventTriggerParameter(model.PBToSystemEventTriggerParameter(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}
