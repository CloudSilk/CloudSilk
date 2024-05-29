// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: mom_common.proto

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

type Code int32

const (
	Code_None Code = 0
	// 成功
	Code_Success Code = 20000
	// 服务端错误
	Code_InternalServerError Code = 50000
	// 错误请求参数
	Code_BadRequest Code = 40000
	// 未授权
	Code_Unauthorized Code = 40001
	// 资源不存在
	Code_ErrRecordNotFound Code = 40002
	// 用户名或者密码错误
	Code_UserNameOrPasswordIsWrong Code = 41001
	// 用户不存在
	Code_UserIsNotExist Code = 41002
	// 没有权限
	Code_NoPermission Code = 41003
	// 无效Token
	Code_TokenInvalid Code = 41004
	// Token过期
	Code_TokenExpired Code = 41005
	// 已禁用用户
	Code_UserDisabled Code = 41006
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:     "None",
		20000: "Success",
		50000: "InternalServerError",
		40000: "BadRequest",
		40001: "Unauthorized",
		40002: "ErrRecordNotFound",
		41001: "UserNameOrPasswordIsWrong",
		41002: "UserIsNotExist",
		41003: "NoPermission",
		41004: "TokenInvalid",
		41005: "TokenExpired",
		41006: "UserDisabled",
	}
	Code_value = map[string]int32{
		"None":                      0,
		"Success":                   20000,
		"InternalServerError":       50000,
		"BadRequest":                40000,
		"Unauthorized":              40001,
		"ErrRecordNotFound":         40002,
		"UserNameOrPasswordIsWrong": 41001,
		"UserIsNotExist":            41002,
		"NoPermission":              41003,
		"TokenInvalid":              41004,
		"TokenExpired":              41005,
		"UserDisabled":              41006,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_mom_common_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_mom_common_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_mom_common_proto_rawDescGZIP(), []int{0}
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code   `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mom_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mom_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_mom_common_proto_rawDescGZIP(), []int{0}
}

func (x *CommonResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *CommonResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	TenantID string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID"`
}

func (x *DelRequest) Reset() {
	*x = DelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mom_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelRequest) ProtoMessage() {}

func (x *DelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mom_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelRequest.ProtoReflect.Descriptor instead.
func (*DelRequest) Descriptor() ([]byte, []int) {
	return file_mom_common_proto_rawDescGZIP(), []int{1}
}

func (x *DelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DelRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

type EnableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	TenantID string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID"`
	Enable   bool   `protobuf:"varint,3,opt,name=enable,proto3" json:"enable"`
}

func (x *EnableRequest) Reset() {
	*x = EnableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mom_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableRequest) ProtoMessage() {}

func (x *EnableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mom_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableRequest.ProtoReflect.Descriptor instead.
func (*EnableRequest) Descriptor() ([]byte, []int) {
	return file_mom_common_proto_rawDescGZIP(), []int{2}
}

func (x *EnableRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EnableRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *EnableRequest) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mom_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mom_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_mom_common_proto_rawDescGZIP(), []int{3}
}

type GetDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	TenantID string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID"`
}

func (x *GetDetailRequest) Reset() {
	*x = GetDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mom_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDetailRequest) ProtoMessage() {}

func (x *GetDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mom_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDetailRequest.ProtoReflect.Descriptor instead.
func (*GetDetailRequest) Descriptor() ([]byte, []int) {
	return file_mom_common_proto_rawDescGZIP(), []int{4}
}

func (x *GetDetailRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetDetailRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

type GetByIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids"`
}

func (x *GetByIDsRequest) Reset() {
	*x = GetByIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mom_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDsRequest) ProtoMessage() {}

func (x *GetByIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mom_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDsRequest.ProtoReflect.Descriptor instead.
func (*GetByIDsRequest) Descriptor() ([]byte, []int) {
	return file_mom_common_proto_rawDescGZIP(), []int{5}
}

func (x *GetByIDsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_mom_common_proto protoreflect.FileDescriptor

var file_mom_common_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x0e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x38, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44,
	0x22, 0x53, 0x0a, 0x0d, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x2a, 0x80, 0x02, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0xa0, 0x9c, 0x01, 0x12, 0x19, 0x0a,
	0x13, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0xd0, 0x86, 0x03, 0x12, 0x10, 0x0a, 0x0a, 0x42, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0xc0, 0xb8, 0x02, 0x12, 0x12, 0x0a, 0x0c, 0x55, 0x6e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x10, 0xc1, 0xb8, 0x02, 0x12, 0x17,
	0x0a, 0x11, 0x45, 0x72, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x6f, 0x74, 0x46, 0x6f,
	0x75, 0x6e, 0x64, 0x10, 0xc2, 0xb8, 0x02, 0x12, 0x1f, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x4f, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x73, 0x57,
	0x72, 0x6f, 0x6e, 0x67, 0x10, 0xa9, 0xc0, 0x02, 0x12, 0x14, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x73, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0xaa, 0xc0, 0x02, 0x12, 0x12,
	0x0a, 0x0c, 0x4e, 0x6f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0xab,
	0xc0, 0x02, 0x12, 0x12, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x10, 0xac, 0xc0, 0x02, 0x12, 0x12, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x10, 0xad, 0xc0, 0x02, 0x12, 0x12, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x10, 0xae, 0xc0, 0x02, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_mom_common_proto_rawDescOnce sync.Once
	file_mom_common_proto_rawDescData = file_mom_common_proto_rawDesc
)

func file_mom_common_proto_rawDescGZIP() []byte {
	file_mom_common_proto_rawDescOnce.Do(func() {
		file_mom_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_mom_common_proto_rawDescData)
	})
	return file_mom_common_proto_rawDescData
}

var file_mom_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mom_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mom_common_proto_goTypes = []interface{}{
	(Code)(0),                // 0: proto.Code
	(*CommonResponse)(nil),   // 1: proto.CommonResponse
	(*DelRequest)(nil),       // 2: proto.DelRequest
	(*EnableRequest)(nil),    // 3: proto.EnableRequest
	(*GetAllRequest)(nil),    // 4: proto.GetAllRequest
	(*GetDetailRequest)(nil), // 5: proto.GetDetailRequest
	(*GetByIDsRequest)(nil),  // 6: proto.GetByIDsRequest
}
var file_mom_common_proto_depIdxs = []int32{
	0, // 0: proto.CommonResponse.code:type_name -> proto.Code
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mom_common_proto_init() }
func file_mom_common_proto_init() {
	if File_mom_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mom_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
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
		file_mom_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelRequest); i {
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
		file_mom_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableRequest); i {
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
		file_mom_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_mom_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDetailRequest); i {
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
		file_mom_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDsRequest); i {
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
			RawDescriptor: file_mom_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mom_common_proto_goTypes,
		DependencyIndexes: file_mom_common_proto_depIdxs,
		EnumInfos:         file_mom_common_proto_enumTypes,
		MessageInfos:      file_mom_common_proto_msgTypes,
	}.Build()
	File_mom_common_proto = out.File
	file_mom_common_proto_rawDesc = nil
	file_mom_common_proto_goTypes = nil
	file_mom_common_proto_depIdxs = nil
}
