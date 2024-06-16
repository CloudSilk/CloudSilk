package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
)

type ProductRhythmRecordProvider struct {
	proto.UnimplementedProductRhythmRecordServer
}

func (u *ProductRhythmRecordProvider) Add(ctx context.Context, in *proto.ProductRhythmRecordInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	id, err := logic.CreateProductRhythmRecord(model.PBToProductRhythmRecord(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}

func (u *ProductRhythmRecordProvider) Get(ctx context.Context, in *proto.GetProductRhythmRecordRequest) (*proto.GetProductRhythmRecordDetailResponse, error) {
	resp := &proto.GetProductRhythmRecordDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductRhythmRecord(in)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductRhythmRecordToPB(f)
	}
	return resp, nil
}

func (u *ProductRhythmRecordProvider) Update(ctx context.Context, in *proto.ProductRhythmRecordInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := logic.UpdateProductRhythmRecord(model.PBToProductRhythmRecord(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}
