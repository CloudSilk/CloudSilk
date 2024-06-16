package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductionProcessSopProvider struct {
	proto.UnimplementedProductionProcessSopServer
}

func (u *ProductionProcessSopProvider) Get(c context.Context, in *proto.GetProductionProcessSopRequest) (*proto.GetProductionProcessSopDetailResponse, error) {
	resp := &proto.GetProductionProcessSopDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductionProcessSop(in)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionProcessSopToPB(f)
	}
	return resp, nil
}
