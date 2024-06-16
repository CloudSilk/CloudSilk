package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
)

type ProductInfoProvider struct {
	proto.UnimplementedProductInfoServer
}

func (u *ProductInfoProvider) Get(ctx context.Context, in *proto.GetProductInfoRequest) (*proto.GetProductInfoDetailResponse, error) {
	resp := &proto.GetProductInfoDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductInfo(in.ProductSerialNo)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductInfoToPB(f)
	}
	return resp, nil
}

func (u *ProductInfoProvider) Query(ctx context.Context, in *proto.QueryProductInfoRequest) (*proto.QueryProductInfoResponse, error) {
	resp := &proto.QueryProductInfoResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProductInfo(in, resp, false)
	return resp, nil
}

func (u *ProductInfoProvider) Update(ctx context.Context, in *proto.ProductInfoInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}

	if err := logic.UpdateProductInfo(model.PBToProductInfo(in)); err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}
