package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
)

type ProductionStationOutputProvider struct {
	proto.UnimplementedProductionStationOutputServer
}

func (u *ProductionStationOutputProvider) Add(ctx context.Context, in *proto.ProductionStationOutputInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	id, err := logic.CreateProductionStationOutput(model.PBToProductionStationOutput(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}

func (u *ProductionStationOutputProvider) Update(ctx context.Context, in *proto.ProductionStationOutputInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := logic.UpdateProductionStationOutput(model.PBToProductionStationOutput(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *ProductionStationOutputProvider) Get(ctx context.Context, in *proto.GetProductionStationOutputRequest) (*proto.GetProductionStationOutputDetailResponse, error) {
	resp := &proto.GetProductionStationOutputDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductionStationOutput(in)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationOutputToPB(f)
	}
	return resp, nil
}
