package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/system/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type SystemEventProvider struct {
	proto.UnimplementedSystemEventServer
}

func (u *SystemEventProvider) Get(ctx context.Context, in *proto.GetSystemEventRequest) (*proto.GetSystemEventDetailResponse, error) {
	resp := &proto.GetSystemEventDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetSystemEvent(in)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.SystemEventToPB(f)
	}
	return resp, nil
}
