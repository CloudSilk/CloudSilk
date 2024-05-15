// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: product_rework_solution.proto

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

type ProductReworkSolutionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 代号
	Code string `protobuf:"bytes,3,opt,name=code,proto3" json:"code"`
	// 描述
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description"`
	// 备注
	Remark string `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark"`
	// 返工原因
	ProductReworkCauseAvailableSolutions []*ProductReworkCauseAvailableSolutionInfo `protobuf:"bytes,6,rep,name=productReworkCauseAvailableSolutions,proto3" json:"productReworkCauseAvailableSolutions"`
}

func (x *ProductReworkSolutionInfo) Reset() {
	*x = ProductReworkSolutionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_solution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductReworkSolutionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductReworkSolutionInfo) ProtoMessage() {}

func (x *ProductReworkSolutionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_solution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductReworkSolutionInfo.ProtoReflect.Descriptor instead.
func (*ProductReworkSolutionInfo) Descriptor() ([]byte, []int) {
	return file_product_rework_solution_proto_rawDescGZIP(), []int{0}
}

func (x *ProductReworkSolutionInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductReworkSolutionInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ProductReworkSolutionInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductReworkSolutionInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProductReworkSolutionInfo) GetProductReworkCauseAvailableSolutions() []*ProductReworkCauseAvailableSolutionInfo {
	if x != nil {
		return x.ProductReworkCauseAvailableSolutions
	}
	return nil
}

type ProductReworkCauseAvailableSolutionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id"`
	// 返工解决方案ID
	ProductReworkSolutionID string `protobuf:"bytes,2,opt,name=productReworkSolutionID,proto3" json:"productReworkSolutionID"`
	// 返工原因ID
	ProductReworkCauseID string `protobuf:"bytes,3,opt,name=productReworkCauseID,proto3" json:"productReworkCauseID"`
}

func (x *ProductReworkCauseAvailableSolutionInfo) Reset() {
	*x = ProductReworkCauseAvailableSolutionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_solution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductReworkCauseAvailableSolutionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductReworkCauseAvailableSolutionInfo) ProtoMessage() {}

func (x *ProductReworkCauseAvailableSolutionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_solution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductReworkCauseAvailableSolutionInfo.ProtoReflect.Descriptor instead.
func (*ProductReworkCauseAvailableSolutionInfo) Descriptor() ([]byte, []int) {
	return file_product_rework_solution_proto_rawDescGZIP(), []int{1}
}

func (x *ProductReworkCauseAvailableSolutionInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductReworkCauseAvailableSolutionInfo) GetProductReworkSolutionID() string {
	if x != nil {
		return x.ProductReworkSolutionID
	}
	return ""
}

func (x *ProductReworkCauseAvailableSolutionInfo) GetProductReworkCauseID() string {
	if x != nil {
		return x.ProductReworkCauseID
	}
	return ""
}

type QueryProductReworkSolutionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 代号或描述
	// @inject_tag: uri:"code" form:"code"
	Code string `protobuf:"bytes,5,opt,name=code,proto3" json:"code" uri:"code" form:"code"`
}

func (x *QueryProductReworkSolutionRequest) Reset() {
	*x = QueryProductReworkSolutionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_solution_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductReworkSolutionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductReworkSolutionRequest) ProtoMessage() {}

func (x *QueryProductReworkSolutionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_solution_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductReworkSolutionRequest.ProtoReflect.Descriptor instead.
func (*QueryProductReworkSolutionRequest) Descriptor() ([]byte, []int) {
	return file_product_rework_solution_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProductReworkSolutionRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductReworkSolutionRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductReworkSolutionRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductReworkSolutionRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type QueryProductReworkSolutionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                         `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                       `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductReworkSolutionInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                        `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                        `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                        `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductReworkSolutionResponse) Reset() {
	*x = QueryProductReworkSolutionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_solution_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductReworkSolutionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductReworkSolutionResponse) ProtoMessage() {}

func (x *QueryProductReworkSolutionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_solution_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductReworkSolutionResponse.ProtoReflect.Descriptor instead.
func (*QueryProductReworkSolutionResponse) Descriptor() ([]byte, []int) {
	return file_product_rework_solution_proto_rawDescGZIP(), []int{3}
}

func (x *QueryProductReworkSolutionResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductReworkSolutionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductReworkSolutionResponse) GetData() []*ProductReworkSolutionInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductReworkSolutionResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductReworkSolutionResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductReworkSolutionResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductReworkSolutionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                         `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                       `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductReworkSolutionInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductReworkSolutionResponse) Reset() {
	*x = GetAllProductReworkSolutionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_solution_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductReworkSolutionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductReworkSolutionResponse) ProtoMessage() {}

func (x *GetAllProductReworkSolutionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_solution_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductReworkSolutionResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductReworkSolutionResponse) Descriptor() ([]byte, []int) {
	return file_product_rework_solution_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllProductReworkSolutionResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductReworkSolutionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductReworkSolutionResponse) GetData() []*ProductReworkSolutionInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductReworkSolutionDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                       `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                     `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductReworkSolutionInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductReworkSolutionDetailResponse) Reset() {
	*x = GetProductReworkSolutionDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_solution_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductReworkSolutionDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductReworkSolutionDetailResponse) ProtoMessage() {}

func (x *GetProductReworkSolutionDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_solution_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductReworkSolutionDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductReworkSolutionDetailResponse) Descriptor() ([]byte, []int) {
	return file_product_rework_solution_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductReworkSolutionDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductReworkSolutionDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductReworkSolutionDetailResponse) GetData() *ProductReworkSolutionInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_product_rework_solution_proto protoreflect.FileDescriptor

var file_product_rework_solution_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x01, 0x0a, 0x19, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x82, 0x01, 0x0a, 0x24, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x75, 0x73, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x75, 0x73, 0x65, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x24, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x43, 0x61, 0x75, 0x73, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x27, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x75, 0x73, 0x65,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x17, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x32, 0x0a, 0x14, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x43, 0x61, 0x75, 0x73, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x75, 0x73,
	0x65, 0x49, 0x44, 0x22, 0x91, 0x01, 0x0a, 0x21, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xdb, 0x01, 0x0a, 0x22, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x96, 0x01, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x99,
	0x01, 0x0a, 0x26, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_rework_solution_proto_rawDescOnce sync.Once
	file_product_rework_solution_proto_rawDescData = file_product_rework_solution_proto_rawDesc
)

func file_product_rework_solution_proto_rawDescGZIP() []byte {
	file_product_rework_solution_proto_rawDescOnce.Do(func() {
		file_product_rework_solution_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_rework_solution_proto_rawDescData)
	})
	return file_product_rework_solution_proto_rawDescData
}

var file_product_rework_solution_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_product_rework_solution_proto_goTypes = []interface{}{
	(*ProductReworkSolutionInfo)(nil),               // 0: proto.ProductReworkSolutionInfo
	(*ProductReworkCauseAvailableSolutionInfo)(nil), // 1: proto.ProductReworkCauseAvailableSolutionInfo
	(*QueryProductReworkSolutionRequest)(nil),       // 2: proto.QueryProductReworkSolutionRequest
	(*QueryProductReworkSolutionResponse)(nil),      // 3: proto.QueryProductReworkSolutionResponse
	(*GetAllProductReworkSolutionResponse)(nil),     // 4: proto.GetAllProductReworkSolutionResponse
	(*GetProductReworkSolutionDetailResponse)(nil),  // 5: proto.GetProductReworkSolutionDetailResponse
	(Code)(0), // 6: proto.Code
}
var file_product_rework_solution_proto_depIdxs = []int32{
	1, // 0: proto.ProductReworkSolutionInfo.productReworkCauseAvailableSolutions:type_name -> proto.ProductReworkCauseAvailableSolutionInfo
	6, // 1: proto.QueryProductReworkSolutionResponse.code:type_name -> proto.Code
	0, // 2: proto.QueryProductReworkSolutionResponse.data:type_name -> proto.ProductReworkSolutionInfo
	6, // 3: proto.GetAllProductReworkSolutionResponse.code:type_name -> proto.Code
	0, // 4: proto.GetAllProductReworkSolutionResponse.data:type_name -> proto.ProductReworkSolutionInfo
	6, // 5: proto.GetProductReworkSolutionDetailResponse.code:type_name -> proto.Code
	0, // 6: proto.GetProductReworkSolutionDetailResponse.data:type_name -> proto.ProductReworkSolutionInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_product_rework_solution_proto_init() }
func file_product_rework_solution_proto_init() {
	if File_product_rework_solution_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_rework_solution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductReworkSolutionInfo); i {
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
		file_product_rework_solution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductReworkCauseAvailableSolutionInfo); i {
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
		file_product_rework_solution_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductReworkSolutionRequest); i {
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
		file_product_rework_solution_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductReworkSolutionResponse); i {
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
		file_product_rework_solution_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductReworkSolutionResponse); i {
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
		file_product_rework_solution_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductReworkSolutionDetailResponse); i {
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
			RawDescriptor: file_product_rework_solution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_rework_solution_proto_goTypes,
		DependencyIndexes: file_product_rework_solution_proto_depIdxs,
		MessageInfos:      file_product_rework_solution_proto_msgTypes,
	}.Build()
	File_product_rework_solution_proto = out.File
	file_product_rework_solution_proto_rawDesc = nil
	file_product_rework_solution_proto_goTypes = nil
	file_product_rework_solution_proto_depIdxs = nil
}