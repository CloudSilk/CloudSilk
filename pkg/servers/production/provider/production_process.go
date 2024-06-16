package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductionProcessProvider struct {
	proto.UnimplementedProductionProcessServer
}

func (u *ProductionProcessProvider) Query(ctx context.Context, in *proto.QueryProductionProcessRequest) (*proto.QueryProductionProcessResponse, error) {
	resp := &proto.QueryProductionProcessResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProductionProcess(in, resp, false)
	return resp, nil
}

func (u *ProductionProcessProvider) GetDetail(ctx context.Context, in *proto.GetDetailRequest) (*proto.GetProductionProcessDetailResponse, error) {
	resp := &proto.GetProductionProcessDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductionProcessByID(in.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionProcessToPB(f)
	}
	return resp, nil
}
