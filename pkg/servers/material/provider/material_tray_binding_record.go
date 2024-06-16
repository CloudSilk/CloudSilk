package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/material/logic"
)

type MaterialTrayBindingRecordProvider struct {
	proto.UnimplementedMaterialTrayBindingRecordServer
}

func (u *MaterialTrayBindingRecordProvider) Add(ctx context.Context, in *proto.MaterialTrayBindingRecordInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	id, err := logic.CreateMaterialTrayBindingRecord(model.PBToMaterialTrayBindingRecord(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}

func (u *MaterialTrayBindingRecordProvider) Get(ctx context.Context, in *proto.GetMaterialTrayBindingRecordRequest) (*proto.GetMaterialTrayBindingRecordDetailResponse, error) {
	resp := &proto.GetMaterialTrayBindingRecordDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetMaterialTrayBindingRecord(in)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialTrayBindingRecordToPB(f)
	}
	return resp, nil
}
