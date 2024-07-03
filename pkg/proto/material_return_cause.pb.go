// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: material_return_cause.proto

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

type MaterialReturnCauseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 代号
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`
	// 描述
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	// 备注
	Remark string `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark"`
	// 物料类别
	MaterialCategories []*MaterialReturnCauseAvailableCategoryInfo `protobuf:"bytes,5,rep,name=materialCategories,proto3" json:"materialCategories"`
	// 归属类型
	MaterialReturnTypes []*MaterialReturnCauseAvailableTypeInfo `protobuf:"bytes,6,rep,name=materialReturnTypes,proto3" json:"materialReturnTypes"`
}

func (x *MaterialReturnCauseInfo) Reset() {
	*x = MaterialReturnCauseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_return_cause_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialReturnCauseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialReturnCauseInfo) ProtoMessage() {}

func (x *MaterialReturnCauseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_material_return_cause_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialReturnCauseInfo.ProtoReflect.Descriptor instead.
func (*MaterialReturnCauseInfo) Descriptor() ([]byte, []int) {
	return file_material_return_cause_proto_rawDescGZIP(), []int{0}
}

func (x *MaterialReturnCauseInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MaterialReturnCauseInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MaterialReturnCauseInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MaterialReturnCauseInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *MaterialReturnCauseInfo) GetMaterialCategories() []*MaterialReturnCauseAvailableCategoryInfo {
	if x != nil {
		return x.MaterialCategories
	}
	return nil
}

func (x *MaterialReturnCauseInfo) GetMaterialReturnTypes() []*MaterialReturnCauseAvailableTypeInfo {
	if x != nil {
		return x.MaterialReturnTypes
	}
	return nil
}

type MaterialReturnCauseAvailableCategoryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	MaterialReturnCauseID string                `protobuf:"bytes,2,opt,name=materialReturnCauseID,proto3" json:"materialReturnCauseID"`
	MaterialCategoryID    string                `protobuf:"bytes,3,opt,name=materialCategoryID,proto3" json:"materialCategoryID"`
	MaterialCategory      *MaterialCategoryInfo `protobuf:"bytes,4,opt,name=materialCategory,proto3" json:"materialCategory"`
}

func (x *MaterialReturnCauseAvailableCategoryInfo) Reset() {
	*x = MaterialReturnCauseAvailableCategoryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_return_cause_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialReturnCauseAvailableCategoryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialReturnCauseAvailableCategoryInfo) ProtoMessage() {}

func (x *MaterialReturnCauseAvailableCategoryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_material_return_cause_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialReturnCauseAvailableCategoryInfo.ProtoReflect.Descriptor instead.
func (*MaterialReturnCauseAvailableCategoryInfo) Descriptor() ([]byte, []int) {
	return file_material_return_cause_proto_rawDescGZIP(), []int{1}
}

func (x *MaterialReturnCauseAvailableCategoryInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MaterialReturnCauseAvailableCategoryInfo) GetMaterialReturnCauseID() string {
	if x != nil {
		return x.MaterialReturnCauseID
	}
	return ""
}

func (x *MaterialReturnCauseAvailableCategoryInfo) GetMaterialCategoryID() string {
	if x != nil {
		return x.MaterialCategoryID
	}
	return ""
}

func (x *MaterialReturnCauseAvailableCategoryInfo) GetMaterialCategory() *MaterialCategoryInfo {
	if x != nil {
		return x.MaterialCategory
	}
	return nil
}

type MaterialReturnCauseAvailableTypeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	MaterialReturnCauseID string                  `protobuf:"bytes,2,opt,name=materialReturnCauseID,proto3" json:"materialReturnCauseID"`
	MaterialReturnTypeID  string                  `protobuf:"bytes,3,opt,name=materialReturnTypeID,proto3" json:"materialReturnTypeID"`
	MaterialReturnType    *MaterialReturnTypeInfo `protobuf:"bytes,4,opt,name=materialReturnType,proto3" json:"materialReturnType"`
}

func (x *MaterialReturnCauseAvailableTypeInfo) Reset() {
	*x = MaterialReturnCauseAvailableTypeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_return_cause_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialReturnCauseAvailableTypeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialReturnCauseAvailableTypeInfo) ProtoMessage() {}

func (x *MaterialReturnCauseAvailableTypeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_material_return_cause_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialReturnCauseAvailableTypeInfo.ProtoReflect.Descriptor instead.
func (*MaterialReturnCauseAvailableTypeInfo) Descriptor() ([]byte, []int) {
	return file_material_return_cause_proto_rawDescGZIP(), []int{2}
}

func (x *MaterialReturnCauseAvailableTypeInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MaterialReturnCauseAvailableTypeInfo) GetMaterialReturnCauseID() string {
	if x != nil {
		return x.MaterialReturnCauseID
	}
	return ""
}

func (x *MaterialReturnCauseAvailableTypeInfo) GetMaterialReturnTypeID() string {
	if x != nil {
		return x.MaterialReturnTypeID
	}
	return ""
}

func (x *MaterialReturnCauseAvailableTypeInfo) GetMaterialReturnType() *MaterialReturnTypeInfo {
	if x != nil {
		return x.MaterialReturnType
	}
	return nil
}

type QueryMaterialReturnCauseRequest struct {
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
	Code string `protobuf:"bytes,4,opt,name=code,proto3" json:"code" uri:"code" form:"code"`
}

func (x *QueryMaterialReturnCauseRequest) Reset() {
	*x = QueryMaterialReturnCauseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_return_cause_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMaterialReturnCauseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMaterialReturnCauseRequest) ProtoMessage() {}

func (x *QueryMaterialReturnCauseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_material_return_cause_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMaterialReturnCauseRequest.ProtoReflect.Descriptor instead.
func (*QueryMaterialReturnCauseRequest) Descriptor() ([]byte, []int) {
	return file_material_return_cause_proto_rawDescGZIP(), []int{3}
}

func (x *QueryMaterialReturnCauseRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryMaterialReturnCauseRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryMaterialReturnCauseRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryMaterialReturnCauseRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type QueryMaterialReturnCauseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                       `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                     `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialReturnCauseInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                      `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                      `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                      `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryMaterialReturnCauseResponse) Reset() {
	*x = QueryMaterialReturnCauseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_return_cause_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMaterialReturnCauseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMaterialReturnCauseResponse) ProtoMessage() {}

func (x *QueryMaterialReturnCauseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_return_cause_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMaterialReturnCauseResponse.ProtoReflect.Descriptor instead.
func (*QueryMaterialReturnCauseResponse) Descriptor() ([]byte, []int) {
	return file_material_return_cause_proto_rawDescGZIP(), []int{4}
}

func (x *QueryMaterialReturnCauseResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryMaterialReturnCauseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryMaterialReturnCauseResponse) GetData() []*MaterialReturnCauseInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryMaterialReturnCauseResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryMaterialReturnCauseResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryMaterialReturnCauseResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllMaterialReturnCauseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                       `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                     `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialReturnCauseInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllMaterialReturnCauseResponse) Reset() {
	*x = GetAllMaterialReturnCauseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_return_cause_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMaterialReturnCauseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMaterialReturnCauseResponse) ProtoMessage() {}

func (x *GetAllMaterialReturnCauseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_return_cause_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMaterialReturnCauseResponse.ProtoReflect.Descriptor instead.
func (*GetAllMaterialReturnCauseResponse) Descriptor() ([]byte, []int) {
	return file_material_return_cause_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllMaterialReturnCauseResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllMaterialReturnCauseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllMaterialReturnCauseResponse) GetData() []*MaterialReturnCauseInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetMaterialReturnCauseDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                     `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                   `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *MaterialReturnCauseInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetMaterialReturnCauseDetailResponse) Reset() {
	*x = GetMaterialReturnCauseDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_return_cause_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMaterialReturnCauseDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMaterialReturnCauseDetailResponse) ProtoMessage() {}

func (x *GetMaterialReturnCauseDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_return_cause_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMaterialReturnCauseDetailResponse.ProtoReflect.Descriptor instead.
func (*GetMaterialReturnCauseDetailResponse) Descriptor() ([]byte, []int) {
	return file_material_return_cause_proto_rawDescGZIP(), []int{6}
}

func (x *GetMaterialReturnCauseDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetMaterialReturnCauseDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetMaterialReturnCauseDetailResponse) GetData() *MaterialReturnCauseInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_material_return_cause_proto protoreflect.FileDescriptor

var file_material_return_cause_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x5f, 0x63, 0x61, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x02, 0x0a, 0x17,
	0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61,
	0x75, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x5f, 0x0a, 0x12, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x65, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x12, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x5d, 0x0a, 0x13, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x65, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x13, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0xe9, 0x01, 0x0a, 0x28, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x65, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x34, 0x0a, 0x15, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x15, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x43, 0x61, 0x75, 0x73, 0x65, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x12, 0x6d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x47, 0x0a, 0x10, 0x6d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x10, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x22, 0xef, 0x01, 0x0a, 0x24, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x15, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x75, 0x73,
	0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x65, 0x49, 0x44,
	0x12, 0x32, 0x0a, 0x14, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x44, 0x12, 0x4d, 0x0a, 0x12, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x12, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x1f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xd7, 0x01, 0x0a, 0x20, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x75,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x92, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x95, 0x01, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x75,
	0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_material_return_cause_proto_rawDescOnce sync.Once
	file_material_return_cause_proto_rawDescData = file_material_return_cause_proto_rawDesc
)

func file_material_return_cause_proto_rawDescGZIP() []byte {
	file_material_return_cause_proto_rawDescOnce.Do(func() {
		file_material_return_cause_proto_rawDescData = protoimpl.X.CompressGZIP(file_material_return_cause_proto_rawDescData)
	})
	return file_material_return_cause_proto_rawDescData
}

var file_material_return_cause_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_material_return_cause_proto_goTypes = []interface{}{
	(*MaterialReturnCauseInfo)(nil),                  // 0: proto.MaterialReturnCauseInfo
	(*MaterialReturnCauseAvailableCategoryInfo)(nil), // 1: proto.MaterialReturnCauseAvailableCategoryInfo
	(*MaterialReturnCauseAvailableTypeInfo)(nil),     // 2: proto.MaterialReturnCauseAvailableTypeInfo
	(*QueryMaterialReturnCauseRequest)(nil),          // 3: proto.QueryMaterialReturnCauseRequest
	(*QueryMaterialReturnCauseResponse)(nil),         // 4: proto.QueryMaterialReturnCauseResponse
	(*GetAllMaterialReturnCauseResponse)(nil),        // 5: proto.GetAllMaterialReturnCauseResponse
	(*GetMaterialReturnCauseDetailResponse)(nil),     // 6: proto.GetMaterialReturnCauseDetailResponse
	(*MaterialCategoryInfo)(nil),                     // 7: proto.MaterialCategoryInfo
	(*MaterialReturnTypeInfo)(nil),                   // 8: proto.MaterialReturnTypeInfo
	(Code)(0),                                        // 9: proto.Code
}
var file_material_return_cause_proto_depIdxs = []int32{
	1,  // 0: proto.MaterialReturnCauseInfo.materialCategories:type_name -> proto.MaterialReturnCauseAvailableCategoryInfo
	2,  // 1: proto.MaterialReturnCauseInfo.materialReturnTypes:type_name -> proto.MaterialReturnCauseAvailableTypeInfo
	7,  // 2: proto.MaterialReturnCauseAvailableCategoryInfo.materialCategory:type_name -> proto.MaterialCategoryInfo
	8,  // 3: proto.MaterialReturnCauseAvailableTypeInfo.materialReturnType:type_name -> proto.MaterialReturnTypeInfo
	9,  // 4: proto.QueryMaterialReturnCauseResponse.code:type_name -> proto.Code
	0,  // 5: proto.QueryMaterialReturnCauseResponse.data:type_name -> proto.MaterialReturnCauseInfo
	9,  // 6: proto.GetAllMaterialReturnCauseResponse.code:type_name -> proto.Code
	0,  // 7: proto.GetAllMaterialReturnCauseResponse.data:type_name -> proto.MaterialReturnCauseInfo
	9,  // 8: proto.GetMaterialReturnCauseDetailResponse.code:type_name -> proto.Code
	0,  // 9: proto.GetMaterialReturnCauseDetailResponse.data:type_name -> proto.MaterialReturnCauseInfo
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_material_return_cause_proto_init() }
func file_material_return_cause_proto_init() {
	if File_material_return_cause_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_material_category_proto_init()
	file_material_return_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_material_return_cause_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialReturnCauseInfo); i {
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
		file_material_return_cause_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialReturnCauseAvailableCategoryInfo); i {
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
		file_material_return_cause_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialReturnCauseAvailableTypeInfo); i {
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
		file_material_return_cause_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMaterialReturnCauseRequest); i {
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
		file_material_return_cause_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMaterialReturnCauseResponse); i {
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
		file_material_return_cause_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMaterialReturnCauseResponse); i {
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
		file_material_return_cause_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMaterialReturnCauseDetailResponse); i {
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
			RawDescriptor: file_material_return_cause_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_material_return_cause_proto_goTypes,
		DependencyIndexes: file_material_return_cause_proto_depIdxs,
		MessageInfos:      file_material_return_cause_proto_msgTypes,
	}.Build()
	File_material_return_cause_proto = out.File
	file_material_return_cause_proto_rawDesc = nil
	file_material_return_cause_proto_goTypes = nil
	file_material_return_cause_proto_depIdxs = nil
}