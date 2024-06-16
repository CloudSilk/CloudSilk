package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
)

type ProductProcessRouteProvider struct {
	proto.UnimplementedProductProcessRouteServer
}

func (u *ProductProcessRouteProvider) Add(ctx context.Context, in *proto.ProductProcessRouteInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	id, err := logic.CreateProductProcessRoute(model.PBToProductProcessRoute(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}

func (u *ProductProcessRouteProvider) Get(ctx context.Context, in *proto.GetProductProcessRouteRequest) (*proto.GetProductProcessRouteDetailResponse, error) {
	resp := &proto.GetProductProcessRouteDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductProcessRoute(in)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductProcessRouteToPB(f)
	}
	return resp, nil
}

func (u *ProductProcessRouteProvider) Query(ctx context.Context, in *proto.QueryProductProcessRouteRequest) (*proto.QueryProductProcessRouteResponse, error) {
	resp := &proto.QueryProductProcessRouteResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProductProcessRoute(in, resp, false)
	return resp, nil
}

func (u *ProductProcessRouteProvider) Update(ctx context.Context, in *proto.ProductProcessRouteInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := logic.UpdateProductProcessRoute(model.PBToProductProcessRoute(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}

	return resp, nil
}
