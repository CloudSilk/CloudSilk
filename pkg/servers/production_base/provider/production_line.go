package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production_base/logic"
)

type ProductionLineProvider struct {
	proto.UnimplementedProductionLineServer
}

func (u *ProductionLineProvider) GetDetail(ctx context.Context, in *proto.GetDetailRequest) (*proto.GetProductionLineDetailResponse, error) {
	resp := &proto.GetProductionLineDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductionLineByID(in.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionLineToPB(f)
	}
	return resp, nil
}

func (u *ProductionLineProvider) GetAll(ctx context.Context, in *proto.GetAllRequest) (*proto.GetAllProductionLineResponse, error) {
	resp := &proto.GetAllProductionLineResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionLines()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionLinesToPB(list)
	}

	return resp, nil
}
