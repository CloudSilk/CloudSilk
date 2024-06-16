package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
)

type ProductOrderProcessStepProvider struct {
	proto.UnimplementedProductOrderProcessStepServer
}

func (u *ProductOrderProcessStepProvider) Query(ctx context.Context, in *proto.QueryProductOrderProcessStepRequest) (*proto.QueryProductOrderProcessStepResponse, error) {
	resp := &proto.QueryProductOrderProcessStepResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProductOrderProcessStep(in, resp, false)
	return resp, nil
}
