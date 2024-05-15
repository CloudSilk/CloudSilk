package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductOrderProvider struct {
	proto.UnimplementedProductOrderServer
}

func (u *ProductOrderProvider) GetDetail(ctx context.Context, in *proto.GetDetailRequest) (*proto.GetProductOrderDetailResponse, error) {
	resp := &proto.GetProductOrderDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductOrderByID(in.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderToPB(f)
	}
	return resp, nil
}
