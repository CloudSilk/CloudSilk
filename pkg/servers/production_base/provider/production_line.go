package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/production_base/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductionLineProvider struct {
	proto.UnimplementedProductionLineServer
}

func (u *ProductionLineProvider) Add(ctx context.Context, in *proto.ProductionLineInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	id, err := logic.CreateProductionLine(model.PBToProductionLine(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}

func (u *ProductionLineProvider) Update(ctx context.Context, in *proto.ProductionLineInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := logic.UpdateProductionLine(model.PBToProductionLine(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *ProductionLineProvider) Delete(ctx context.Context, in *proto.DelRequest) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := logic.DeleteProductionLine(in.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *ProductionLineProvider) Query(ctx context.Context, in *proto.QueryProductionLineRequest) (*proto.QueryProductionLineResponse, error) {
	resp := &proto.QueryProductionLineResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProductionLine(in, resp, false)
	return resp, nil
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
