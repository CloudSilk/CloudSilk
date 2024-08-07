// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: product_rework_cause.proto

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

type ProductReworkCauseInfo struct {
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
	// 返工类型
	ProductReworkTypePossibleCauses []*ProductReworkTypePossibleCauseInfo `protobuf:"bytes,6,rep,name=productReworkTypePossibleCauses,proto3" json:"productReworkTypePossibleCauses"`
}

func (x *ProductReworkCauseInfo) Reset() {
	*x = ProductReworkCauseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_cause_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductReworkCauseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductReworkCauseInfo) ProtoMessage() {}

func (x *ProductReworkCauseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_cause_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductReworkCauseInfo.ProtoReflect.Descriptor instead.
func (*ProductReworkCauseInfo) Descriptor() ([]byte, []int) {
	return file_product_rework_cause_proto_rawDescGZIP(), []int{0}
}

func (x *ProductReworkCauseInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductReworkCauseInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ProductReworkCauseInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductReworkCauseInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProductReworkCauseInfo) GetProductReworkTypePossibleCauses() []*ProductReworkTypePossibleCauseInfo {
	if x != nil {
		return x.ProductReworkTypePossibleCauses
	}
	return nil
}

type ProductReworkTypePossibleCauseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id"`
	// 返工原因ID
	ProductReworkCauseID string `protobuf:"bytes,2,opt,name=productReworkCauseID,proto3" json:"productReworkCauseID"`
	// 返工类型ID
	ProductReworkTypeID string `protobuf:"bytes,3,opt,name=productReworkTypeID,proto3" json:"productReworkTypeID"`
}

func (x *ProductReworkTypePossibleCauseInfo) Reset() {
	*x = ProductReworkTypePossibleCauseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_cause_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductReworkTypePossibleCauseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductReworkTypePossibleCauseInfo) ProtoMessage() {}

func (x *ProductReworkTypePossibleCauseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_cause_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductReworkTypePossibleCauseInfo.ProtoReflect.Descriptor instead.
func (*ProductReworkTypePossibleCauseInfo) Descriptor() ([]byte, []int) {
	return file_product_rework_cause_proto_rawDescGZIP(), []int{1}
}

func (x *ProductReworkTypePossibleCauseInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductReworkTypePossibleCauseInfo) GetProductReworkCauseID() string {
	if x != nil {
		return x.ProductReworkCauseID
	}
	return ""
}

func (x *ProductReworkTypePossibleCauseInfo) GetProductReworkTypeID() string {
	if x != nil {
		return x.ProductReworkTypeID
	}
	return ""
}

type QueryProductReworkCauseRequest struct {
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

func (x *QueryProductReworkCauseRequest) Reset() {
	*x = QueryProductReworkCauseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_cause_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductReworkCauseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductReworkCauseRequest) ProtoMessage() {}

func (x *QueryProductReworkCauseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_cause_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductReworkCauseRequest.ProtoReflect.Descriptor instead.
func (*QueryProductReworkCauseRequest) Descriptor() ([]byte, []int) {
	return file_product_rework_cause_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProductReworkCauseRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductReworkCauseRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductReworkCauseRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductReworkCauseRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type QueryProductReworkCauseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                      `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                    `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductReworkCauseInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                     `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                     `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                     `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductReworkCauseResponse) Reset() {
	*x = QueryProductReworkCauseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_cause_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductReworkCauseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductReworkCauseResponse) ProtoMessage() {}

func (x *QueryProductReworkCauseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_cause_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductReworkCauseResponse.ProtoReflect.Descriptor instead.
func (*QueryProductReworkCauseResponse) Descriptor() ([]byte, []int) {
	return file_product_rework_cause_proto_rawDescGZIP(), []int{3}
}

func (x *QueryProductReworkCauseResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductReworkCauseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductReworkCauseResponse) GetData() []*ProductReworkCauseInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductReworkCauseResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductReworkCauseResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductReworkCauseResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductReworkCauseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                      `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                    `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductReworkCauseInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductReworkCauseResponse) Reset() {
	*x = GetAllProductReworkCauseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_cause_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductReworkCauseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductReworkCauseResponse) ProtoMessage() {}

func (x *GetAllProductReworkCauseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_cause_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductReworkCauseResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductReworkCauseResponse) Descriptor() ([]byte, []int) {
	return file_product_rework_cause_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllProductReworkCauseResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductReworkCauseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductReworkCauseResponse) GetData() []*ProductReworkCauseInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductReworkCauseDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                    `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductReworkCauseInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductReworkCauseDetailResponse) Reset() {
	*x = GetProductReworkCauseDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_rework_cause_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductReworkCauseDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductReworkCauseDetailResponse) ProtoMessage() {}

func (x *GetProductReworkCauseDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_rework_cause_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductReworkCauseDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductReworkCauseDetailResponse) Descriptor() ([]byte, []int) {
	return file_product_rework_cause_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductReworkCauseDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductReworkCauseDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductReworkCauseDetailResponse) GetData() *ProductReworkCauseInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_product_rework_cause_proto protoreflect.FileDescriptor

var file_product_rework_cause_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x63, 0x61, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x75, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x73,
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x50, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x50, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x1f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x50, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x43, 0x61, 0x75,
	0x73, 0x65, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x22, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x50, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c,
	0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x75, 0x73, 0x65,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x75, 0x73, 0x65, 0x49, 0x44, 0x12, 0x30,
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44,
	0x22, 0x8e, 0x01, 0x0a, 0x1e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0xd5, 0x01, 0x0a, 0x1f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x75, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x90, 0x01, 0x0a, 0x20, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x43, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x75,
	0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x93, 0x01, 0x0a,
	0x23, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x43, 0x61, 0x75, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x43, 0x61, 0x75, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_rework_cause_proto_rawDescOnce sync.Once
	file_product_rework_cause_proto_rawDescData = file_product_rework_cause_proto_rawDesc
)

func file_product_rework_cause_proto_rawDescGZIP() []byte {
	file_product_rework_cause_proto_rawDescOnce.Do(func() {
		file_product_rework_cause_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_rework_cause_proto_rawDescData)
	})
	return file_product_rework_cause_proto_rawDescData
}

var file_product_rework_cause_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_product_rework_cause_proto_goTypes = []interface{}{
	(*ProductReworkCauseInfo)(nil),              // 0: proto.ProductReworkCauseInfo
	(*ProductReworkTypePossibleCauseInfo)(nil),  // 1: proto.ProductReworkTypePossibleCauseInfo
	(*QueryProductReworkCauseRequest)(nil),      // 2: proto.QueryProductReworkCauseRequest
	(*QueryProductReworkCauseResponse)(nil),     // 3: proto.QueryProductReworkCauseResponse
	(*GetAllProductReworkCauseResponse)(nil),    // 4: proto.GetAllProductReworkCauseResponse
	(*GetProductReworkCauseDetailResponse)(nil), // 5: proto.GetProductReworkCauseDetailResponse
	(Code)(0), // 6: proto.Code
}
var file_product_rework_cause_proto_depIdxs = []int32{
	1, // 0: proto.ProductReworkCauseInfo.productReworkTypePossibleCauses:type_name -> proto.ProductReworkTypePossibleCauseInfo
	6, // 1: proto.QueryProductReworkCauseResponse.code:type_name -> proto.Code
	0, // 2: proto.QueryProductReworkCauseResponse.data:type_name -> proto.ProductReworkCauseInfo
	6, // 3: proto.GetAllProductReworkCauseResponse.code:type_name -> proto.Code
	0, // 4: proto.GetAllProductReworkCauseResponse.data:type_name -> proto.ProductReworkCauseInfo
	6, // 5: proto.GetProductReworkCauseDetailResponse.code:type_name -> proto.Code
	0, // 6: proto.GetProductReworkCauseDetailResponse.data:type_name -> proto.ProductReworkCauseInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_product_rework_cause_proto_init() }
func file_product_rework_cause_proto_init() {
	if File_product_rework_cause_proto != nil {
		return
	}
	file_mom_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_rework_cause_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductReworkCauseInfo); i {
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
		file_product_rework_cause_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductReworkTypePossibleCauseInfo); i {
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
		file_product_rework_cause_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductReworkCauseRequest); i {
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
		file_product_rework_cause_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductReworkCauseResponse); i {
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
		file_product_rework_cause_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductReworkCauseResponse); i {
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
		file_product_rework_cause_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductReworkCauseDetailResponse); i {
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
			RawDescriptor: file_product_rework_cause_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_rework_cause_proto_goTypes,
		DependencyIndexes: file_product_rework_cause_proto_depIdxs,
		MessageInfos:      file_product_rework_cause_proto_msgTypes,
	}.Build()
	File_product_rework_cause_proto = out.File
	file_product_rework_cause_proto_rawDesc = nil
	file_product_rework_cause_proto_goTypes = nil
	file_product_rework_cause_proto_depIdxs = nil
}
