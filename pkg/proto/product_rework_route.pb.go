// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: product_rework_route.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductReworkRouteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id                 string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	ProductionLineID   string                 `protobuf:"bytes,2,opt,name=productionLineID,proto3" json:"productionLineID"`
	ProductionLine     *ProductionLineInfo    `protobuf:"bytes,3,opt,name=productionLine,proto3" json:"productionLine"`
	MaterialCategoryID string                 `protobuf:"bytes,4,opt,name=materialCategoryID,proto3" json:"materialCategoryID"`
	MaterialCategory   *MaterialCategoryInfo  `protobuf:"bytes,5,opt,name=materialCategory,proto3" json:"materialCategory"`
	FollowProcessID    string                 `protobuf:"bytes,6,opt,name=followProcessID,proto3" json:"followProcessID"`
	FollowProcess      *ProductionProcessInfo `protobuf:"bytes,7,opt,name=followProcess,proto3" json:"followProcess"`
}

func (x *ProductReworkRouteInfo) Reset() {
	*x = ProductReworkRouteInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductReworkRouteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductReworkRouteInfo) ProtoMessage() {}

func (x *ProductReworkRouteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductReworkRouteInfo.ProtoReflect.Descriptor instead.
func (*ProductReworkRouteInfo) Descriptor() ([]byte, []int) {
	return file_product_rework_route_proto_rawDescGZIP(), []int{0}
}

func (x *ProductReworkRouteInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductReworkRouteInfo) GetProductionLineID() string {
	if x != nil {
		return x.ProductionLineID
	}
	return ""
}

func (x *ProductReworkRouteInfo) GetProductionLine() *ProductionLineInfo {
	if x != nil {
		return x.ProductionLine
	}
	return nil
}

func (x *ProductReworkRouteInfo) GetMaterialCategoryID() string {
	if x != nil {
		return x.MaterialCategoryID
	}
	return ""
}

func (x *ProductReworkRouteInfo) GetMaterialCategory() *MaterialCategoryInfo {
	if x != nil {
		return x.MaterialCategory
	}
	return nil
}

func (x *ProductReworkRouteInfo) GetFollowProcessID() string {
	if x != nil {
		return x.FollowProcessID
	}
	return ""
}

func (x *ProductReworkRouteInfo) GetFollowProcess() *ProductionProcessInfo {
	if x != nil {
		return x.FollowProcess
	}
	return nil
}

type QueryProductReworkRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 生产产线ID
	// @inject_tag: uri:"productionLineID" form:"productionLineID"
	ProductionLineID string `protobuf:"bytes,4,opt,name=productionLineID,proto3" json:"productionLineID" uri:"productionLineID" form:"productionLineID"`
	// 代号或描述
	// @inject_tag: uri:"code" form:"code"
	Code string `protobuf:"bytes,5,opt,name=code,proto3" json:"code" uri:"code" form:"code"`
}

func (x *QueryProductReworkRouteRequest) Reset() {
	*x = QueryProductReworkRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductReworkRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductReworkRouteRequest) ProtoMessage() {}

func (x *QueryProductReworkRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductReworkRouteRequest.ProtoReflect.Descriptor instead.
func (*QueryProductReworkRouteRequest) Descriptor() ([]byte, []int) {
	return file_product_rework_route_proto_rawDescGZIP(), []int{1}
}

func (x *QueryProductReworkRouteRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductReworkRouteRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductReworkRouteRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductReworkRouteRequest) GetProductionLineID() string {
	if x != nil {
		return x.ProductionLineID
	}
	return ""
}

func (x *QueryProductReworkRouteRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type QueryProductReworkRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                      `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                    `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductReworkRouteInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                     `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                     `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                     `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductReworkRouteResponse) Reset() {
	*x = QueryProductReworkRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductReworkRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductReworkRouteResponse) ProtoMessage() {}

func (x *QueryProductReworkRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductReworkRouteResponse.ProtoReflect.Descriptor instead.
func (*QueryProductReworkRouteResponse) Descriptor() ([]byte, []int) {
	return file_product_rework_route_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProductReworkRouteResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductReworkRouteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductReworkRouteResponse) GetData() []*ProductReworkRouteInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductReworkRouteResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductReworkRouteResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductReworkRouteResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductReworkRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                      `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                    `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductReworkRouteInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductReworkRouteResponse) Reset() {
	*x = GetAllProductReworkRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductReworkRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductReworkRouteResponse) ProtoMessage() {}

func (x *GetAllProductReworkRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductReworkRouteResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductReworkRouteResponse) Descriptor() ([]byte, []int) {
	return file_product_rework_route_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllProductReworkRouteResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductReworkRouteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductReworkRouteResponse) GetData() []*ProductReworkRouteInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductReworkRouteDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                    `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductReworkRouteInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductReworkRouteDetailResponse) Reset() {
	*x = GetProductReworkRouteDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_route_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductReworkRouteDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductReworkRouteDetailResponse) ProtoMessage() {}

func (x *GetProductReworkRouteDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_route_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductReworkRouteDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductReworkRouteDetailResponse) Descriptor() ([]byte, []int) {
	return file_product_rework_route_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductReworkRouteDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductReworkRouteDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductReworkRouteDetailResponse) GetData() *ProductReworkRouteInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_product_rework_route_proto protoreflect.FileDescriptor

var file_product_rework_route_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x6e, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x41, 0x0a, 0x0e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x12,
	0x2e, 0x0a, 0x12, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12,
	0x47, 0x0a, 0x10, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x49, 0x44, 0x12, 0x42, 0x0a, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x22, 0xba, 0x01, 0x0a, 0x1e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0xd5, 0x01, 0x0a, 0x1f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x90, 0x01, 0x0a, 0x20,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x93,
	0x01, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_rework_route_proto_rawDescOnce sync.Once
	file_product_rework_route_proto_rawDescData = file_product_rework_route_proto_rawDesc
)

func file_product_rework_route_proto_rawDescGZIP() []byte {
	file_product_rework_route_proto_rawDescOnce.Do(func() {
		file_product_rework_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_rework_route_proto_rawDescData)
	})
	return file_product_rework_route_proto_rawDescData
}

var file_product_rework_route_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_product_rework_route_proto_goTypes = []interface{}{
	(*ProductReworkRouteInfo)(nil),              // 0: proto.ProductReworkRouteInfo
	(*QueryProductReworkRouteRequest)(nil),      // 1: proto.QueryProductReworkRouteRequest
	(*QueryProductReworkRouteResponse)(nil),     // 2: proto.QueryProductReworkRouteResponse
	(*GetAllProductReworkRouteResponse)(nil),    // 3: proto.GetAllProductReworkRouteResponse
	(*GetProductReworkRouteDetailResponse)(nil), // 4: proto.GetProductReworkRouteDetailResponse
	(*ProductionLineInfo)(nil),                  // 5: proto.ProductionLineInfo
	(*MaterialCategoryInfo)(nil),                // 6: proto.MaterialCategoryInfo
	(*ProductionProcessInfo)(nil),               // 7: proto.ProductionProcessInfo
	(Code)(0),                                   // 8: proto.Code
}
var file_product_rework_route_proto_depIdxs = []int32{
	5, // 0: proto.ProductReworkRouteInfo.productionLine:type_name -> proto.ProductionLineInfo
	6, // 1: proto.ProductReworkRouteInfo.materialCategory:type_name -> proto.MaterialCategoryInfo
	7, // 2: proto.ProductReworkRouteInfo.followProcess:type_name -> proto.ProductionProcessInfo
	8, // 3: proto.QueryProductReworkRouteResponse.code:type_name -> proto.Code
	0, // 4: proto.QueryProductReworkRouteResponse.data:type_name -> proto.ProductReworkRouteInfo
	8, // 5: proto.GetAllProductReworkRouteResponse.code:type_name -> proto.Code
	0, // 6: proto.GetAllProductReworkRouteResponse.data:type_name -> proto.ProductReworkRouteInfo
	8, // 7: proto.GetProductReworkRouteDetailResponse.code:type_name -> proto.Code
	0, // 8: proto.GetProductReworkRouteDetailResponse.data:type_name -> proto.ProductReworkRouteInfo
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_product_rework_route_proto_init() }
func file_product_rework_route_proto_init() {
	if File_product_rework_route_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_material_category_proto_init()
	file_production_line_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_rework_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductReworkRouteInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_rework_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductReworkRouteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_rework_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductReworkRouteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_rework_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductReworkRouteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_rework_route_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductReworkRouteDetailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_product_rework_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_rework_route_proto_goTypes,
		DependencyIndexes: file_product_rework_route_proto_depIdxs,
		MessageInfos:      file_product_rework_route_proto_msgTypes,
	}.Build()
	File_product_rework_route_proto = out.File
	file_product_rework_route_proto_rawDesc = nil
	file_product_rework_route_proto_goTypes = nil
	file_product_rework_route_proto_depIdxs = nil
}
