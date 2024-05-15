package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/production_base/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductionCrosswayProvider struct {
	proto.UnimplementedProductionCrosswayServer
}

func (u *ProductionCrosswayProvider) Add(ctx context.Context, in *proto.ProductionCrosswayInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	id, err := logic.CreateProductionCrossway(model.PBToProductionCrossway(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}

func (u *ProductionCrosswayProvider) Update(ctx context.Context, in *proto.ProductionCrosswayInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := logic.UpdateProductionCrossway(model.PBToProductionCrossway(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *ProductionCrosswayProvider) Delete(ctx context.Context, in *proto.DelRequest) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := logic.DeleteProductionCrossway(in.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *ProductionCrosswayProvider) Query(ctx context.Context, in *proto.QueryProductionCrosswayRequest) (*proto.QueryProductionCrosswayResponse, error) {
	resp := &proto.QueryProductionCrosswayResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProductionCrossway(in, resp, false)
	return resp, nil
}

func (u *ProductionCrosswayProvider) GetDetail(ctx context.Context, in *proto.GetDetailRequest) (*proto.GetProductionCrosswayDetailResponse, error) {
	resp := &proto.GetProductionCrosswayDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductionCrosswayByID(in.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionCrosswayToPB(f)
	}
	return resp, nil
}

func (u *ProductionCrosswayProvider) GetAll(ctx context.Context, in *proto.GetAllRequest) (*proto.GetAllProductionCrosswayResponse, error) {
	resp := &proto.GetAllProductionCrosswayResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionCrossways()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionCrosswaysToPB(list)
	}

	return resp, nil
}
