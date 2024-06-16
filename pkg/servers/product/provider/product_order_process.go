package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductOrderProcessProvider struct {
	proto.UnimplementedProductOrderProcessServer
}

func (u ProductOrderProcessProvider) Query(ctx context.Context, in *proto.QueryProductOrderProcessRequest) (*proto.QueryProductOrderProcessResponse, error) {
	resp := &proto.QueryProductOrderProcessResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProductOrderProcess(in, resp, false)
	return resp, nil
}
