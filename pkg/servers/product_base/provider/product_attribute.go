package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product_base/logic"
)

type ProductAttributeProvider struct {
	proto.UnimplementedProductAttributeServer
}

func (u *ProductAttributeProvider) Query(ctx context.Context, in *proto.QueryProductAttributeRequest) (*proto.QueryProductAttributeResponse, error) {
	resp := &proto.QueryProductAttributeResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProductAttribute(in, resp, false)
	return resp, nil
}
