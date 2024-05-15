package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductRhythmRecordProvider struct {
	proto.UnimplementedProductRhythmRecordServer
}

func (u *ProductRhythmRecordProvider) Add(ctx context.Context, in *proto.ProductRhythmRecordInfo) (*proto.CommonResponse, error) {
	return nil, nil
}

func (u *ProductRhythmRecordProvider) Get(ctx context.Context, in *proto.GetProductRhythmRecordRequest) (*proto.GetProductRhythmRecordDetailResponse, error) {
	resp := &proto.GetProductRhythmRecordDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductRhythmRecord(in.ProductionProcessID, in.ProductInfoID, in.ProductionStationID, in.HasWorkEndTime)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductRhythmRecordToPB(f)
	}
	return resp, nil
}
