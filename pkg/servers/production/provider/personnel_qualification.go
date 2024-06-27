package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
)

type PersonnelQualificationProvider struct {
	proto.UnimplementedPersonnelQualificationServer
}

func (u *PersonnelQualificationProvider) Query(ctx context.Context, in *proto.QueryPersonnelQualificationRequest) (*proto.QueryPersonnelQualificationResponse, error) {
	resp := &proto.QueryPersonnelQualificationResponse{
		Code: proto.Code_Success,
	}
	logic.QueryPersonnelQualification(in, resp, false)
	return resp, nil
}

func (u *PersonnelQualificationProvider) Get(ctx context.Context, in *proto.GetPersonnelQualificationRequest) (*proto.GetPersonnelQualificationDetailResponse, error) {
	resp := &proto.GetPersonnelQualificationDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetPersonnelQualification(in)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.PersonnelQualificationToPB(f)
	}
	return resp, nil
}
