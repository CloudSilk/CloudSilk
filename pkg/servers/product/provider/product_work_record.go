package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
)

type ProductWorkRecordProvider struct {
	proto.UnimplementedProductWorkRecordServer
}

func (u *ProductWorkRecordProvider) Query(ctx context.Context, in *proto.QueryProductWorkRecordRequest) (*proto.QueryProductWorkRecordResponse, error) {
	resp := &proto.QueryProductWorkRecordResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProductWorkRecord(in, resp, false)
	return resp, nil
}
