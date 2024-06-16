package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductOrderAttributeProvider struct {
	proto.UnimplementedProductOrderAttributeServer
}

func (u *ProductOrderAttributeProvider) Query(ctx context.Context, in *proto.QueryProductOrderAttributeRequest) (*proto.QueryProductOrderAttributeResponse, error) {
	resp := &proto.QueryProductOrderAttributeResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProductOrderAttribute(in, resp, false)
	return resp, nil
}
