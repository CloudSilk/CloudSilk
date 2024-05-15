// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: product_order_priority_rule.proto

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

type ProductOrderPriorityRuleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id"`
	// 优先级
	Priority int32 `protobuf:"varint,4,opt,name=priority,proto3" json:"priority"`
	// 生产优先级
	PriorityLevel int32 `protobuf:"varint,5,opt,name=priorityLevel,proto3" json:"priorityLevel"`
	// 是否启用
	Enable bool `protobuf:"varint,6,opt,name=enable,proto3" json:"enable"`
	// 默认排序
	InitialValue bool `protobuf:"varint,7,opt,name=initialValue,proto3" json:"initialValue"`
	// 备注
	Remark string `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark"`
	// 特征表达式
	AttributeExpressions []*AttributeExpressionInfo `protobuf:"bytes,2,rep,name=attributeExpressions,proto3" json:"attributeExpressions"`
}

func (x *ProductOrderPriorityRuleInfo) Reset() {
	*x = ProductOrderPriorityRuleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_order_priority_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOrderPriorityRuleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOrderPriorityRuleInfo) ProtoMessage() {}

func (x *ProductOrderPriorityRuleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_order_priority_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOrderPriorityRuleInfo.ProtoReflect.Descriptor instead.
func (*ProductOrderPriorityRuleInfo) Descriptor() ([]byte, []int) {
	return file_product_order_priority_rule_proto_rawDescGZIP(), []int{0}
}

func (x *ProductOrderPriorityRuleInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductOrderPriorityRuleInfo) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *ProductOrderPriorityRuleInfo) GetPriorityLevel() int32 {
	if x != nil {
		return x.PriorityLevel
	}
	return 0
}

func (x *ProductOrderPriorityRuleInfo) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *ProductOrderPriorityRuleInfo) GetInitialValue() bool {
	if x != nil {
		return x.InitialValue
	}
	return false
}

func (x *ProductOrderPriorityRuleInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProductOrderPriorityRuleInfo) GetAttributeExpressions() []*AttributeExpressionInfo {
	if x != nil {
		return x.AttributeExpressions
	}
	return nil
}

type QueryProductOrderPriorityRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 生产优先级
	// @inject_tag: uri:"priorityLevel" form:"priorityLevel"
	PriorityLevel int32 `protobuf:"varint,6,opt,name=priorityLevel,proto3" json:"priorityLevel" uri:"priorityLevel" form:"priorityLevel"`
}

func (x *QueryProductOrderPriorityRuleRequest) Reset() {
	*x = QueryProductOrderPriorityRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_order_priority_rule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductOrderPriorityRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductOrderPriorityRuleRequest) ProtoMessage() {}

func (x *QueryProductOrderPriorityRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_order_priority_rule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductOrderPriorityRuleRequest.ProtoReflect.Descriptor instead.
func (*QueryProductOrderPriorityRuleRequest) Descriptor() ([]byte, []int) {
	return file_product_order_priority_rule_proto_rawDescGZIP(), []int{1}
}

func (x *QueryProductOrderPriorityRuleRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductOrderPriorityRuleRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductOrderPriorityRuleRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductOrderPriorityRuleRequest) GetPriorityLevel() int32 {
	if x != nil {
		return x.PriorityLevel
	}
	return 0
}

type QueryProductOrderPriorityRuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                            `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                          `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductOrderPriorityRuleInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                           `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                           `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                           `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductOrderPriorityRuleResponse) Reset() {
	*x = QueryProductOrderPriorityRuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_order_priority_rule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductOrderPriorityRuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductOrderPriorityRuleResponse) ProtoMessage() {}

func (x *QueryProductOrderPriorityRuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_order_priority_rule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductOrderPriorityRuleResponse.ProtoReflect.Descriptor instead.
func (*QueryProductOrderPriorityRuleResponse) Descriptor() ([]byte, []int) {
	return file_product_order_priority_rule_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProductOrderPriorityRuleResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductOrderPriorityRuleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductOrderPriorityRuleResponse) GetData() []*ProductOrderPriorityRuleInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductOrderPriorityRuleResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductOrderPriorityRuleResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductOrderPriorityRuleResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductOrderPriorityRuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                            `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                          `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductOrderPriorityRuleInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductOrderPriorityRuleResponse) Reset() {
	*x = GetAllProductOrderPriorityRuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_order_priority_rule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductOrderPriorityRuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductOrderPriorityRuleResponse) ProtoMessage() {}

func (x *GetAllProductOrderPriorityRuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_order_priority_rule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductOrderPriorityRuleResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductOrderPriorityRuleResponse) Descriptor() ([]byte, []int) {
	return file_product_order_priority_rule_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllProductOrderPriorityRuleResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductOrderPriorityRuleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductOrderPriorityRuleResponse) GetData() []*ProductOrderPriorityRuleInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductOrderPriorityRuleDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                          `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                        `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductOrderPriorityRuleInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductOrderPriorityRuleDetailResponse) Reset() {
	*x = GetProductOrderPriorityRuleDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_order_priority_rule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductOrderPriorityRuleDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductOrderPriorityRuleDetailResponse) ProtoMessage() {}

func (x *GetProductOrderPriorityRuleDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_order_priority_rule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductOrderPriorityRuleDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductOrderPriorityRuleDetailResponse) Descriptor() ([]byte, []int) {
	return file_product_order_priority_rule_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductOrderPriorityRuleDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductOrderPriorityRuleDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductOrderPriorityRuleDetailResponse) GetData() *ProductOrderPriorityRuleInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_product_order_priority_rule_proto protoreflect.FileDescriptor

var file_product_order_priority_rule_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x02, 0x0a, 0x1c, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x52, 0x0a, 0x14, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x45,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0xa6, 0x01, 0x0a, 0x24, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0xe1, 0x01, 0x0a,
	0x25, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x9c, 0x01, 0x0a, 0x26, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x9f, 0x01, 0x0a, 0x29, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_order_priority_rule_proto_rawDescOnce sync.Once
	file_product_order_priority_rule_proto_rawDescData = file_product_order_priority_rule_proto_rawDesc
)

func file_product_order_priority_rule_proto_rawDescGZIP() []byte {
	file_product_order_priority_rule_proto_rawDescOnce.Do(func() {
		file_product_order_priority_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_order_priority_rule_proto_rawDescData)
	})
	return file_product_order_priority_rule_proto_rawDescData
}

var file_product_order_priority_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_product_order_priority_rule_proto_goTypes = []interface{}{
	(*ProductOrderPriorityRuleInfo)(nil),              // 0: proto.ProductOrderPriorityRuleInfo
	(*QueryProductOrderPriorityRuleRequest)(nil),      // 1: proto.QueryProductOrderPriorityRuleRequest
	(*QueryProductOrderPriorityRuleResponse)(nil),     // 2: proto.QueryProductOrderPriorityRuleResponse
	(*GetAllProductOrderPriorityRuleResponse)(nil),    // 3: proto.GetAllProductOrderPriorityRuleResponse
	(*GetProductOrderPriorityRuleDetailResponse)(nil), // 4: proto.GetProductOrderPriorityRuleDetailResponse
	(*AttributeExpressionInfo)(nil),                   // 5: proto.AttributeExpressionInfo
	(Code)(0),                                         // 6: proto.Code
}
var file_product_order_priority_rule_proto_depIdxs = []int32{
	5, // 0: proto.ProductOrderPriorityRuleInfo.attributeExpressions:type_name -> proto.AttributeExpressionInfo
	6, // 1: proto.QueryProductOrderPriorityRuleResponse.code:type_name -> proto.Code
	0, // 2: proto.QueryProductOrderPriorityRuleResponse.data:type_name -> proto.ProductOrderPriorityRuleInfo
	6, // 3: proto.GetAllProductOrderPriorityRuleResponse.code:type_name -> proto.Code
	0, // 4: proto.GetAllProductOrderPriorityRuleResponse.data:type_name -> proto.ProductOrderPriorityRuleInfo
	6, // 5: proto.GetProductOrderPriorityRuleDetailResponse.code:type_name -> proto.Code
	0, // 6: proto.GetProductOrderPriorityRuleDetailResponse.data:type_name -> proto.ProductOrderPriorityRuleInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_product_order_priority_rule_proto_init() }
func file_product_order_priority_rule_proto_init() {
	if File_product_order_priority_rule_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_attribute_expression_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_order_priority_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOrderPriorityRuleInfo); i {
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
		file_product_order_priority_rule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductOrderPriorityRuleRequest); i {
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
		file_product_order_priority_rule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductOrderPriorityRuleResponse); i {
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
		file_product_order_priority_rule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductOrderPriorityRuleResponse); i {
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
		file_product_order_priority_rule_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductOrderPriorityRuleDetailResponse); i {
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
			RawDescriptor: file_product_order_priority_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_order_priority_rule_proto_goTypes,
		DependencyIndexes: file_product_order_priority_rule_proto_depIdxs,
		MessageInfos:      file_product_order_priority_rule_proto_msgTypes,
	}.Build()
	File_product_order_priority_rule_proto = out.File
	file_product_order_priority_rule_proto_rawDesc = nil
	file_product_order_priority_rule_proto_goTypes = nil
	file_product_order_priority_rule_proto_depIdxs = nil
}