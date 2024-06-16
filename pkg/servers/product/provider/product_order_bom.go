package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
)

type ProductOrderBomProvider struct {
	proto.UnimplementedProductOrderBomServer
}

func (u *ProductOrderBomProvider) Query(ctx context.Context, in *proto.QueryProductOrderBomRequest) (*proto.QueryProductOrderBomResponse, error) {
	resp := &proto.QueryProductOrderBomResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProductOrderBom(in, resp, false)
	return resp, nil
}
