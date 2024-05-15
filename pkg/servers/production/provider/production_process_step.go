package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductionProcessStepProvider struct {
	proto.UnimplementedProductionProcessStepServer
}

func (u *ProductionProcessStepProvider) Query(ctx context.Context, in *proto.QueryProductionProcessStepRequest) (*proto.QueryProductionProcessStepResponse, error) {
	resp := &proto.QueryProductionProcessStepResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProductionProcessStep(in, resp, false)
	return resp, nil
}

func (u *ProductionProcessStepProvider) Get(ctx context.Context, in *proto.GetProductionProcessStepRequest) (*proto.GetProductionProcessStepDetailResponse, error) {
	resp := &proto.GetProductionProcessStepDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductionProcessStep(in)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionProcessStepToPB(f)
	}
	return resp, nil
}
