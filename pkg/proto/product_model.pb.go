// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: product_model.proto

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

type ProductModelInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 型号
	Code string `protobuf:"bytes,3,opt,name=code,proto3" json:"code"`
	// 物料号
	MaterialNo string `protobuf:"bytes,4,opt,name=materialNo,proto3" json:"materialNo"`
	// 物料描述
	MaterialDescription string `protobuf:"bytes,5,opt,name=materialDescription,proto3" json:"materialDescription"`
	// 识别码
	Identifier string `protobuf:"bytes,6,opt,name=identifier,proto3" json:"identifier"`
	// 是否预制
	IsPrefabricated bool `protobuf:"varint,7,opt,name=isPrefabricated,proto3" json:"isPrefabricated"`
	// 是否授权
	IsAuthorized bool `protobuf:"varint,8,opt,name=isAuthorized,proto3" json:"isAuthorized"`
	// 备注
	Remark string `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark"`
	// 产品类别ID
	ProductCategoryID string               `protobuf:"bytes,10,opt,name=productCategoryID,proto3" json:"productCategoryID"`
	ProductCategory   *ProductCategoryInfo `protobuf:"bytes,13,opt,name=productCategory,proto3" json:"productCategory"`
	// 产品型号特征值
	ProductModelAttributeValues []*ProductModelAttributeValueInfo `protobuf:"bytes,11,rep,name=productModelAttributeValues,proto3" json:"productModelAttributeValues"`
	// 产品型号Bom
	ProductModelBoms []*ProductModelBomInfo `protobuf:"bytes,12,rep,name=productModelBoms,proto3" json:"productModelBoms"`
}

func (x *ProductModelInfo) Reset() {
	*x = ProductModelInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductModelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductModelInfo) ProtoMessage() {}

func (x *ProductModelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductModelInfo.ProtoReflect.Descriptor instead.
func (*ProductModelInfo) Descriptor() ([]byte, []int) {
	return file_product_model_proto_rawDescGZIP(), []int{0}
}

func (x *ProductModelInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductModelInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ProductModelInfo) GetMaterialNo() string {
	if x != nil {
		return x.MaterialNo
	}
	return ""
}

func (x *ProductModelInfo) GetMaterialDescription() string {
	if x != nil {
		return x.MaterialDescription
	}
	return ""
}

func (x *ProductModelInfo) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *ProductModelInfo) GetIsPrefabricated() bool {
	if x != nil {
		return x.IsPrefabricated
	}
	return false
}

func (x *ProductModelInfo) GetIsAuthorized() bool {
	if x != nil {
		return x.IsAuthorized
	}
	return false
}

func (x *ProductModelInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProductModelInfo) GetProductCategoryID() string {
	if x != nil {
		return x.ProductCategoryID
	}
	return ""
}

func (x *ProductModelInfo) GetProductCategory() *ProductCategoryInfo {
	if x != nil {
		return x.ProductCategory
	}
	return nil
}

func (x *ProductModelInfo) GetProductModelAttributeValues() []*ProductModelAttributeValueInfo {
	if x != nil {
		return x.ProductModelAttributeValues
	}
	return nil
}

func (x *ProductModelInfo) GetProductModelBoms() []*ProductModelBomInfo {
	if x != nil {
		return x.ProductModelBoms
	}
	return nil
}

type ProductModelAttributeValueInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 产品型号ID
	ProductModelID string `protobuf:"bytes,2,opt,name=productModelID,proto3" json:"productModelID"`
	// 产品特性ID
	ProductAttributeID string                `protobuf:"bytes,3,opt,name=productAttributeID,proto3" json:"productAttributeID"`
	ProductAttribute   *ProductAttributeInfo `protobuf:"bytes,4,opt,name=productAttribute,proto3" json:"productAttribute"`
	// 设定值
	AssignedValue string `protobuf:"bytes,5,opt,name=assignedValue,proto3" json:"assignedValue"`
	// 允许空缺
	AllowNullOrBlank bool `protobuf:"varint,6,opt,name=allowNullOrBlank,proto3" json:"allowNullOrBlank"`
}

func (x *ProductModelAttributeValueInfo) Reset() {
	*x = ProductModelAttributeValueInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductModelAttributeValueInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductModelAttributeValueInfo) ProtoMessage() {}

func (x *ProductModelAttributeValueInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductModelAttributeValueInfo.ProtoReflect.Descriptor instead.
func (*ProductModelAttributeValueInfo) Descriptor() ([]byte, []int) {
	return file_product_model_proto_rawDescGZIP(), []int{1}
}

func (x *ProductModelAttributeValueInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductModelAttributeValueInfo) GetProductModelID() string {
	if x != nil {
		return x.ProductModelID
	}
	return ""
}

func (x *ProductModelAttributeValueInfo) GetProductAttributeID() string {
	if x != nil {
		return x.ProductAttributeID
	}
	return ""
}

func (x *ProductModelAttributeValueInfo) GetProductAttribute() *ProductAttributeInfo {
	if x != nil {
		return x.ProductAttribute
	}
	return nil
}

func (x *ProductModelAttributeValueInfo) GetAssignedValue() string {
	if x != nil {
		return x.AssignedValue
	}
	return ""
}

func (x *ProductModelAttributeValueInfo) GetAllowNullOrBlank() bool {
	if x != nil {
		return x.AllowNullOrBlank
	}
	return false
}

type QueryProductModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 型号
	// @inject_tag: uri:"code" form:"code"
	Code string `protobuf:"bytes,5,opt,name=code,proto3" json:"code" uri:"code" form:"code"`
	// 产品类别ID
	// @inject_tag: uri:"productCategoryID" form:"productCategoryID"
	ProductCategoryID string `protobuf:"bytes,6,opt,name=productCategoryID,proto3" json:"productCategoryID" uri:"productCategoryID" form:"productCategoryID"`
	// 是否预制
	// @inject_tag: uri:"isPrefabricated" form:"isPrefabricated"
	IsPrefabricated bool `protobuf:"varint,7,opt,name=isPrefabricated,proto3" json:"isPrefabricated" uri:"isPrefabricated" form:"isPrefabricated"`
}

func (x *QueryProductModelRequest) Reset() {
	*x = QueryProductModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductModelRequest) ProtoMessage() {}

func (x *QueryProductModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductModelRequest.ProtoReflect.Descriptor instead.
func (*QueryProductModelRequest) Descriptor() ([]byte, []int) {
	return file_product_model_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProductModelRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductModelRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductModelRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductModelRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *QueryProductModelRequest) GetProductCategoryID() string {
	if x != nil {
		return x.ProductCategoryID
	}
	return ""
}

func (x *QueryProductModelRequest) GetIsPrefabricated() bool {
	if x != nil {
		return x.IsPrefabricated
	}
	return false
}

type QueryProductModelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string              `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductModelInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64               `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64               `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64               `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductModelResponse) Reset() {
	*x = QueryProductModelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductModelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductModelResponse) ProtoMessage() {}

func (x *QueryProductModelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductModelResponse.ProtoReflect.Descriptor instead.
func (*QueryProductModelResponse) Descriptor() ([]byte, []int) {
	return file_product_model_proto_rawDescGZIP(), []int{3}
}

func (x *QueryProductModelResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductModelResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductModelResponse) GetData() []*ProductModelInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductModelResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductModelResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductModelResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductModelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string              `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductModelInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductModelResponse) Reset() {
	*x = GetAllProductModelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductModelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductModelResponse) ProtoMessage() {}

func (x *GetAllProductModelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductModelResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductModelResponse) Descriptor() ([]byte, []int) {
	return file_product_model_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllProductModelResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductModelResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductModelResponse) GetData() []*ProductModelInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductModelDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code              `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductModelInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductModelDetailResponse) Reset() {
	*x = GetProductModelDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductModelDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductModelDetailResponse) ProtoMessage() {}

func (x *GetProductModelDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductModelDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductModelDetailResponse) Descriptor() ([]byte, []int) {
	return file_product_model_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductModelDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductModelDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductModelDetailResponse) GetData() *ProductModelInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_product_model_proto protoreflect.FileDescriptor

var file_product_model_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f,
	0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x62, 0x6f,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x04, 0x0a, 0x10, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x6f, 0x12, 0x30, 0x0a, 0x13, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x73, 0x50, 0x72, 0x65, 0x66, 0x61, 0x62, 0x72,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73,
	0x50, 0x72, 0x65, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x44, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x67, 0x0a,
	0x1b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x1b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x6f, 0x6d, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x6f, 0x6d, 0x73, 0x22, 0xa3,
	0x02, 0x0a, 0x1e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x12, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x44, 0x12, 0x47, 0x0a, 0x10, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x4e, 0x75, 0x6c, 0x6c, 0x4f, 0x72, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4e, 0x75, 0x6c, 0x6c, 0x4f, 0x72, 0x42,
	0x6c, 0x61, 0x6e, 0x6b, 0x22, 0xe0, 0x01, 0x0a, 0x18, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x28, 0x0a,
	0x0f, 0x69, 0x73, 0x50, 0x72, 0x65, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x50, 0x72, 0x65, 0x66, 0x61, 0x62,
	0x72, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x22, 0xc9, 0x01, 0x0a, 0x19, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x84, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x87, 0x01, 0x0a, 0x1d, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0x5c, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x4c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_model_proto_rawDescOnce sync.Once
	file_product_model_proto_rawDescData = file_product_model_proto_rawDesc
)

func file_product_model_proto_rawDescGZIP() []byte {
	file_product_model_proto_rawDescOnce.Do(func() {
		file_product_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_model_proto_rawDescData)
	})
	return file_product_model_proto_rawDescData
}

var file_product_model_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_product_model_proto_goTypes = []interface{}{
	(*ProductModelInfo)(nil),               // 0: proto.ProductModelInfo
	(*ProductModelAttributeValueInfo)(nil), // 1: proto.ProductModelAttributeValueInfo
	(*QueryProductModelRequest)(nil),       // 2: proto.QueryProductModelRequest
	(*QueryProductModelResponse)(nil),      // 3: proto.QueryProductModelResponse
	(*GetAllProductModelResponse)(nil),     // 4: proto.GetAllProductModelResponse
	(*GetProductModelDetailResponse)(nil),  // 5: proto.GetProductModelDetailResponse
	(*ProductCategoryInfo)(nil),            // 6: proto.ProductCategoryInfo
	(*ProductModelBomInfo)(nil),            // 7: proto.ProductModelBomInfo
	(*ProductAttributeInfo)(nil),           // 8: proto.ProductAttributeInfo
	(Code)(0),                              // 9: proto.Code
	(*GetDetailRequest)(nil),               // 10: proto.GetDetailRequest
}
var file_product_model_proto_depIdxs = []int32{
	6,  // 0: proto.ProductModelInfo.productCategory:type_name -> proto.ProductCategoryInfo
	1,  // 1: proto.ProductModelInfo.productModelAttributeValues:type_name -> proto.ProductModelAttributeValueInfo
	7,  // 2: proto.ProductModelInfo.productModelBoms:type_name -> proto.ProductModelBomInfo
	8,  // 3: proto.ProductModelAttributeValueInfo.productAttribute:type_name -> proto.ProductAttributeInfo
	9,  // 4: proto.QueryProductModelResponse.code:type_name -> proto.Code
	0,  // 5: proto.QueryProductModelResponse.data:type_name -> proto.ProductModelInfo
	9,  // 6: proto.GetAllProductModelResponse.code:type_name -> proto.Code
	0,  // 7: proto.GetAllProductModelResponse.data:type_name -> proto.ProductModelInfo
	9,  // 8: proto.GetProductModelDetailResponse.code:type_name -> proto.Code
	0,  // 9: proto.GetProductModelDetailResponse.data:type_name -> proto.ProductModelInfo
	10, // 10: proto.ProductModel.GetDetail:input_type -> proto.GetDetailRequest
	5,  // 11: proto.ProductModel.GetDetail:output_type -> proto.GetProductModelDetailResponse
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_product_model_proto_init() }
func file_product_model_proto_init() {
	if File_product_model_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_product_model_bom_proto_init()
	file_product_category_proto_init()
	file_product_attribute_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductModelInfo); i {
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
		file_product_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductModelAttributeValueInfo); i {
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
		file_product_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductModelRequest); i {
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
		file_product_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductModelResponse); i {
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
		file_product_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductModelResponse); i {
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
		file_product_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductModelDetailResponse); i {
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
			RawDescriptor: file_product_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_model_proto_goTypes,
		DependencyIndexes: file_product_model_proto_depIdxs,
		MessageInfos:      file_product_model_proto_msgTypes,
	}.Build()
	File_product_model_proto = out.File
	file_product_model_proto_rawDesc = nil
	file_product_model_proto_goTypes = nil
	file_product_model_proto_depIdxs = nil
}
