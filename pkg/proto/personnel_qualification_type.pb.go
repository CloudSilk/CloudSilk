// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: personnel_qualification_type.proto

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

type PersonnelQualificationTypeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 代号
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`
	// 描述
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	// 有效时长(天)
	EffectiveDuration int32 `protobuf:"varint,4,opt,name=effectiveDuration,proto3" json:"effectiveDuration"`
	// 失效预警
	ExpirationEarlyWarning bool `protobuf:"varint,5,opt,name=expirationEarlyWarning,proto3" json:"expirationEarlyWarning"`
	// 产品型号
	ProductModels []*ProductModelInfo `protobuf:"bytes,6,rep,name=productModels,proto3" json:"productModels"`
	// 生产工序ID
	ProductionProcessID string `protobuf:"bytes,7,opt,name=productionProcessID,proto3" json:"productionProcessID"`
	// 生产工序
	ProductionProcess *ProductionProcessInfo `protobuf:"bytes,8,opt,name=productionProcess,proto3" json:"productionProcess"`
	// 备注
	Remark string `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark"`
}

func (x *PersonnelQualificationTypeInfo) Reset() {
	*x = PersonnelQualificationTypeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personnel_qualification_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonnelQualificationTypeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonnelQualificationTypeInfo) ProtoMessage() {}

func (x *PersonnelQualificationTypeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_personnel_qualification_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonnelQualificationTypeInfo.ProtoReflect.Descriptor instead.
func (*PersonnelQualificationTypeInfo) Descriptor() ([]byte, []int) {
	return file_personnel_qualification_type_proto_rawDescGZIP(), []int{0}
}

func (x *PersonnelQualificationTypeInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PersonnelQualificationTypeInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PersonnelQualificationTypeInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PersonnelQualificationTypeInfo) GetEffectiveDuration() int32 {
	if x != nil {
		return x.EffectiveDuration
	}
	return 0
}

func (x *PersonnelQualificationTypeInfo) GetExpirationEarlyWarning() bool {
	if x != nil {
		return x.ExpirationEarlyWarning
	}
	return false
}

func (x *PersonnelQualificationTypeInfo) GetProductModels() []*ProductModelInfo {
	if x != nil {
		return x.ProductModels
	}
	return nil
}

func (x *PersonnelQualificationTypeInfo) GetProductionProcessID() string {
	if x != nil {
		return x.ProductionProcessID
	}
	return ""
}

func (x *PersonnelQualificationTypeInfo) GetProductionProcess() *ProductionProcessInfo {
	if x != nil {
		return x.ProductionProcess
	}
	return nil
}

func (x *PersonnelQualificationTypeInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type QueryPersonnelQualificationTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
}

func (x *QueryPersonnelQualificationTypeRequest) Reset() {
	*x = QueryPersonnelQualificationTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personnel_qualification_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPersonnelQualificationTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPersonnelQualificationTypeRequest) ProtoMessage() {}

func (x *QueryPersonnelQualificationTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personnel_qualification_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPersonnelQualificationTypeRequest.ProtoReflect.Descriptor instead.
func (*QueryPersonnelQualificationTypeRequest) Descriptor() ([]byte, []int) {
	return file_personnel_qualification_type_proto_rawDescGZIP(), []int{1}
}

func (x *QueryPersonnelQualificationTypeRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryPersonnelQualificationTypeRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryPersonnelQualificationTypeRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

type QueryPersonnelQualificationTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                              `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                            `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*PersonnelQualificationTypeInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                             `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                             `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                             `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryPersonnelQualificationTypeResponse) Reset() {
	*x = QueryPersonnelQualificationTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personnel_qualification_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPersonnelQualificationTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPersonnelQualificationTypeResponse) ProtoMessage() {}

func (x *QueryPersonnelQualificationTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_personnel_qualification_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPersonnelQualificationTypeResponse.ProtoReflect.Descriptor instead.
func (*QueryPersonnelQualificationTypeResponse) Descriptor() ([]byte, []int) {
	return file_personnel_qualification_type_proto_rawDescGZIP(), []int{2}
}

func (x *QueryPersonnelQualificationTypeResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryPersonnelQualificationTypeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryPersonnelQualificationTypeResponse) GetData() []*PersonnelQualificationTypeInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryPersonnelQualificationTypeResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryPersonnelQualificationTypeResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryPersonnelQualificationTypeResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllPersonnelQualificationTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                              `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                            `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*PersonnelQualificationTypeInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllPersonnelQualificationTypeResponse) Reset() {
	*x = GetAllPersonnelQualificationTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personnel_qualification_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPersonnelQualificationTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPersonnelQualificationTypeResponse) ProtoMessage() {}

func (x *GetAllPersonnelQualificationTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_personnel_qualification_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPersonnelQualificationTypeResponse.ProtoReflect.Descriptor instead.
func (*GetAllPersonnelQualificationTypeResponse) Descriptor() ([]byte, []int) {
	return file_personnel_qualification_type_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllPersonnelQualificationTypeResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllPersonnelQualificationTypeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllPersonnelQualificationTypeResponse) GetData() []*PersonnelQualificationTypeInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetPersonnelQualificationTypeDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                            `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                          `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *PersonnelQualificationTypeInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetPersonnelQualificationTypeDetailResponse) Reset() {
	*x = GetPersonnelQualificationTypeDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personnel_qualification_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonnelQualificationTypeDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonnelQualificationTypeDetailResponse) ProtoMessage() {}

func (x *GetPersonnelQualificationTypeDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_personnel_qualification_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonnelQualificationTypeDetailResponse.ProtoReflect.Descriptor instead.
func (*GetPersonnelQualificationTypeDetailResponse) Descriptor() ([]byte, []int) {
	return file_personnel_qualification_type_proto_rawDescGZIP(), []int{4}
}

func (x *GetPersonnelQualificationTypeDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetPersonnelQualificationTypeDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetPersonnelQualificationTypeDetailResponse) GetData() *PersonnelQualificationTypeInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_personnel_qualification_type_proto protoreflect.FileDescriptor

var file_personnel_qualification_type_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x71, 0x75, 0x61, 0x6c,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x03, 0x0a,
	0x1e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x6e, 0x65, 0x6c, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x11, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x16, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x61, 0x72, 0x6c, 0x79, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x16, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x61, 0x72, 0x6c, 0x79, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x0d, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x12, 0x4a, 0x0a, 0x11,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x22, 0x82, 0x01, 0x0a, 0x26, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x6e, 0x65, 0x6c, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xe5, 0x01, 0x0a, 0x27, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x6e, 0x65, 0x6c, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x6e, 0x65, 0x6c, 0x51, 0x75, 0x61, 0x6c,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xa0, 0x01,
	0x0a, 0x28, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x6e, 0x65,
	0x6c, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x6e, 0x65, 0x6c, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xa3, 0x01, 0x0a, 0x2b, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x6e, 0x65,
	0x6c, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x6e, 0x65, 0x6c, 0x51, 0x75, 0x61, 0x6c, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_personnel_qualification_type_proto_rawDescOnce sync.Once
	file_personnel_qualification_type_proto_rawDescData = file_personnel_qualification_type_proto_rawDesc
)

func file_personnel_qualification_type_proto_rawDescGZIP() []byte {
	file_personnel_qualification_type_proto_rawDescOnce.Do(func() {
		file_personnel_qualification_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_personnel_qualification_type_proto_rawDescData)
	})
	return file_personnel_qualification_type_proto_rawDescData
}

var file_personnel_qualification_type_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_personnel_qualification_type_proto_goTypes = []interface{}{
	(*PersonnelQualificationTypeInfo)(nil),              // 0: proto.PersonnelQualificationTypeInfo
	(*QueryPersonnelQualificationTypeRequest)(nil),      // 1: proto.QueryPersonnelQualificationTypeRequest
	(*QueryPersonnelQualificationTypeResponse)(nil),     // 2: proto.QueryPersonnelQualificationTypeResponse
	(*GetAllPersonnelQualificationTypeResponse)(nil),    // 3: proto.GetAllPersonnelQualificationTypeResponse
	(*GetPersonnelQualificationTypeDetailResponse)(nil), // 4: proto.GetPersonnelQualificationTypeDetailResponse
	(*ProductModelInfo)(nil),                            // 5: proto.ProductModelInfo
	(*ProductionProcessInfo)(nil),                       // 6: proto.ProductionProcessInfo
	(Code)(0),                                           // 7: proto.Code
}
var file_personnel_qualification_type_proto_depIdxs = []int32{
	5, // 0: proto.PersonnelQualificationTypeInfo.productModels:type_name -> proto.ProductModelInfo
	6, // 1: proto.PersonnelQualificationTypeInfo.productionProcess:type_name -> proto.ProductionProcessInfo
	7, // 2: proto.QueryPersonnelQualificationTypeResponse.code:type_name -> proto.Code
	0, // 3: proto.QueryPersonnelQualificationTypeResponse.data:type_name -> proto.PersonnelQualificationTypeInfo
	7, // 4: proto.GetAllPersonnelQualificationTypeResponse.code:type_name -> proto.Code
	0, // 5: proto.GetAllPersonnelQualificationTypeResponse.data:type_name -> proto.PersonnelQualificationTypeInfo
	7, // 6: proto.GetPersonnelQualificationTypeDetailResponse.code:type_name -> proto.Code
	0, // 7: proto.GetPersonnelQualificationTypeDetailResponse.data:type_name -> proto.PersonnelQualificationTypeInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_personnel_qualification_type_proto_init() }
func file_personnel_qualification_type_proto_init() {
	if File_personnel_qualification_type_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_product_model_proto_init()
	file_production_process_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_personnel_qualification_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonnelQualificationTypeInfo); i {
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
		file_personnel_qualification_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPersonnelQualificationTypeRequest); i {
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
		file_personnel_qualification_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPersonnelQualificationTypeResponse); i {
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
		file_personnel_qualification_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPersonnelQualificationTypeResponse); i {
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
		file_personnel_qualification_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonnelQualificationTypeDetailResponse); i {
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
			RawDescriptor: file_personnel_qualification_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_personnel_qualification_type_proto_goTypes,
		DependencyIndexes: file_personnel_qualification_type_proto_depIdxs,
		MessageInfos:      file_personnel_qualification_type_proto_msgTypes,
	}.Build()
	File_personnel_qualification_type_proto = out.File
	file_personnel_qualification_type_proto_rawDesc = nil
	file_personnel_qualification_type_proto_goTypes = nil
	file_personnel_qualification_type_proto_depIdxs = nil
}