package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/product_base/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductAttributeProvider struct {
	proto.UnimplementedProductAttributeServer
}

func (u *ProductAttributeProvider) Add(ctx context.Context, in *proto.ProductAttributeInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	id, err := logic.CreateProductAttribute(model.PBToProductAttribute(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}

func (u *ProductAttributeProvider) Update(ctx context.Context, in *proto.ProductAttributeInfo) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := logic.UpdateProductAttribute(model.PBToProductAttribute(in))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *ProductAttributeProvider) Delete(ctx context.Context, in *proto.DelRequest) (*proto.CommonResponse, error) {
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := logic.DeleteProductAttribute(in.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *ProductAttributeProvider) Query(ctx context.Context, in *proto.QueryProductAttributeRequest) (*proto.QueryProductAttributeResponse, error) {
	resp := &proto.QueryProductAttributeResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProductAttribute(in, resp, false)
	return resp, nil
}

func (u *ProductAttributeProvider) GetDetail(ctx context.Context, in *proto.GetDetailRequest) (*proto.GetProductAttributeDetailResponse, error) {
	resp := &proto.GetProductAttributeDetailResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetProductAttributeByID(in.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductAttributeToPB(f)
	}
	return resp, nil
}

func (u *ProductAttributeProvider) GetAll(ctx context.Context, in *proto.GetAllRequest) (*proto.GetAllProductAttributeResponse, error) {
	resp := &proto.GetAllProductAttributeResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductAttributes()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductAttributesToPB(list)
	}

	return resp, nil
}
