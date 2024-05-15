package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/production_base/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductionStationProvider struct {
	proto.UnimplementedProductionStationServer
}

func (u *ProductionStationProvider) Add(ctx context.Context, in *proto.ProductionStationInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	id, err := logic.CreateProductionStation(model.PBToProductionStation(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}

func (u *ProductionStationProvider) Update(ctx context.Context, in *proto.ProductionStationInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := logic.UpdateProductionStation(model.PBToProductionStation(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *ProductionStationProvider) Delete(ctx context.Context, in *proto.DelRequest) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := logic.DeleteProductionStation(in.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *ProductionStationProvider) Query(ctx context.Context, in *proto.QueryProductionStationRequest) (*proto.QueryProductionStationResponse, error) {
	resp := &proto.QueryProductionStationResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProductionStation(in, resp, false)
	return resp, nil
}

func (u *ProductionStationProvider) GetDetail(ctx context.Context, in *proto.GetDetailRequest) (*proto.GetProductionStationDetailResponse, error) {
	resp := &proto.GetProductionStationDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductionStationByID(in.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationToPB(f)
	}
	return resp, nil
}

func (u *ProductionStationProvider) GetAll(ctx context.Context, in *proto.GetAllRequest) (*proto.GetAllProductionStationResponse, error) {
	resp := &proto.GetAllProductionStationResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionStations()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationsToPB(list)
	}

	return resp, nil
}

func (u *ProductionStationProvider) Get(ctx context.Context, in *proto.GetProductionStationRequest) (*proto.GetProductionStationDetailResponse, error) {
	resp := &proto.GetProductionStationDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductionStationByCode(in.Code)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationToPB(f)
	}
	return resp, nil
}
