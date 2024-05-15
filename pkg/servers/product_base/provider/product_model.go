package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/product_base/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductModelProvider struct {
	proto.UnimplementedProductModelServer
}

func (u *ProductModelProvider) GetDetail(ctx context.Context, in *proto.GetDetailRequest) (*proto.GetProductModelDetailResponse, error) {
	resp := &proto.GetProductModelDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductModelByID(in.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductModelToPB(f)
	}
	return resp, nil
}
