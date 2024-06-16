package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production_base/logic"
)

type ProductionStationProvider struct {
	proto.UnimplementedProductionStationServer
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

func (u *ProductionStationProvider) Get(ctx context.Context, in *proto.GetProductionStationRequest) (*proto.GetProductionStationDetailResponse, error) {
	resp := &proto.GetProductionStationDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductionStation(in)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationToPB(f)
	}
	return resp, nil
}
