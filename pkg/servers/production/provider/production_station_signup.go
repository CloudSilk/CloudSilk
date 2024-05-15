package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductionStationSignupProvider struct {
	proto.UnimplementedProductionStationSignupServer
}

func (u *ProductionStationSignupProvider) Add(ctx context.Context, in *proto.ProductionStationSignupInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	id, err := logic.CreateProductionStationSignup(model.PBToProductionStationSignup(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}

func (u *ProductionStationSignupProvider) Update(ctx context.Context, in *proto.ProductionStationSignupInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := logic.UpdateProductionStationSignup(model.PBToProductionStationSignup(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *ProductionStationSignupProvider) Delete(ctx context.Context, in *proto.DelRequest) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := logic.DeleteProductionStationSignup(in.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *ProductionStationSignupProvider) Query(ctx context.Context, in *proto.QueryProductionStationSignupRequest) (*proto.QueryProductionStationSignupResponse, error) {
	resp := &proto.QueryProductionStationSignupResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProductionStationSignup(in, resp, false)
	return resp, nil
}

func (u *ProductionStationSignupProvider) GetDetail(ctx context.Context, in *proto.GetDetailRequest) (*proto.GetProductionStationSignupDetailResponse, error) {
	resp := &proto.GetProductionStationSignupDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductionStationSignupByID(in.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationSignupToPB(f)
	}
	return resp, nil
}

func (u *ProductionStationSignupProvider) GetAll(ctx context.Context, in *proto.GetAllRequest) (*proto.GetAllProductionStationSignupResponse, error) {
	resp := &proto.GetAllProductionStationSignupResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionStationSignups()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationSignupsToPB(list)
	}

	return resp, nil
}

func (u *ProductionStationSignupProvider) GetByID(ctx context.Context, in *proto.GetProductionStationSignupRequest) (*proto.GetProductionStationSignupDetailResponse, error) {
	resp := &proto.GetProductionStationSignupDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductionStationSignup(in.ProductionStationID, in.LoginUserID, in.HasLogoutTime)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationSignupToPB(f)
	}
	return resp, nil
}
