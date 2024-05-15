package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductTestRecordProvider struct {
	proto.UnimplementedProductTestRecordServer
}

func (u *ProductTestRecordProvider) Add(ctx context.Context, in *proto.ProductTestRecordInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	id, err := logic.CreateProductTestRecord(model.PBToProductTestRecord(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}
