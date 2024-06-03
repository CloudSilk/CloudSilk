package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
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

func (u *ProductionStationSignupProvider) Get(ctx context.Context, in *proto.GetProductionStationSignupRequest) (*proto.GetProductionStationSignupDetailResponse, error) {
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
