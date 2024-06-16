package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
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

func (u *ProductOrderProvider) Update(ctx context.Context, in *proto.ProductOrderInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}

	if err := logic.UpdateProductOrder(model.PBToProductOrder(in)); err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}
