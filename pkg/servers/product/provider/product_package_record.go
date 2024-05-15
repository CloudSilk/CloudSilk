package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductPackageRecordProvider struct {
	proto.UnimplementedProductPackageRecordServer
}

func (u *ProductPackageRecordProvider) Get(ctx context.Context, in *proto.GetProductPackageRecordRequest) (*proto.GetProductPackageRecordDetailResponse, error) {
	resp := &proto.GetProductPackageRecordDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductPackageRecordByPackageNo(in.PackageNo)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductPackageRecordToPB(f)
	}
	return resp, nil
}
