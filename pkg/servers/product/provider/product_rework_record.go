package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductReworkRecordProvider struct {
	proto.UnimplementedProductReworkRecordServer
}

func (u *ProductReworkRecordProvider) Add(ctx context.Context, in *proto.ProductReworkRecordInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	id, err := logic.CreateProductReworkRecord(model.PBToProductReworkRecord(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}
