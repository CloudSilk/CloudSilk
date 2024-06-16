package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/material/logic"
)

type MaterialChannelLayerProvider struct {
	proto.UnimplementedMaterialChannelLayerServer
}

func (u *MaterialChannelLayerProvider) GetMaterialChannels(ctx context.Context, in *proto.GetMaterialChannelRequest) (*proto.GetAllMaterialChannelResponse, error) {
	resp := &proto.GetAllMaterialChannelResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetMaterialChannels(in)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialChannelsToPB(f)
	}
	return resp, nil
}
