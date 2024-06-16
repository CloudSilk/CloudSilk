package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/material/logic"
)

type MaterialTrayProvider struct {
	proto.UnimplementedMaterialTrayServer
}

func (u *MaterialTrayProvider) Get(ctx context.Context, in *proto.GetMaterialTrayRequest) (*proto.GetMaterialTrayDetailResponse, error) {
	resp := &proto.GetMaterialTrayDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetMaterialTray(in)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialTrayToPB(f)
	}
	return resp, nil
}

func (u *MaterialTrayProvider) Update(ctx context.Context, in *proto.MaterialTrayInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}

	if err := logic.UpdateMaterialTray(model.PBToMaterialTray(in)); err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}
