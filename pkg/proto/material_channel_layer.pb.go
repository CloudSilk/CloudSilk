// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: material_channel_layer.proto

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

type MaterialChannelLayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 顺序
	SortIndex int32 `protobuf:"varint,2,opt,name=sortIndex,proto3" json:"sortIndex"`
	// 代号
	Code string `protobuf:"bytes,3,opt,name=code,proto3" json:"code"`
	// 状态寄存器地址
	StatusRegisterAddress int32 `protobuf:"varint,4,opt,name=statusRegisterAddress,proto3" json:"statusRegisterAddress"`
	// 亮灯寄存器地址
	LightRegisterAddress int32 `protobuf:"varint,5,opt,name=lightRegisterAddress,proto3" json:"lightRegisterAddress"`
	// 备注
	Remark string `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark"`
	// 隶属工站ID
	ProductionStationID string `protobuf:"bytes,7,opt,name=productionStationID,proto3" json:"productionStationID"`
	// 隶属工站
	ProductionStation *ProductionStationInfo `protobuf:"bytes,8,opt,name=productionStation,proto3" json:"productionStation"`
}

func (x *MaterialChannelLayerInfo) Reset() {
	*x = MaterialChannelLayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_channel_layer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialChannelLayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialChannelLayerInfo) ProtoMessage() {}

func (x *MaterialChannelLayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_material_channel_layer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialChannelLayerInfo.ProtoReflect.Descriptor instead.
func (*MaterialChannelLayerInfo) Descriptor() ([]byte, []int) {
	return file_material_channel_layer_proto_rawDescGZIP(), []int{0}
}

func (x *MaterialChannelLayerInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MaterialChannelLayerInfo) GetSortIndex() int32 {
	if x != nil {
		return x.SortIndex
	}
	return 0
}

func (x *MaterialChannelLayerInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MaterialChannelLayerInfo) GetStatusRegisterAddress() int32 {
	if x != nil {
		return x.StatusRegisterAddress
	}
	return 0
}

func (x *MaterialChannelLayerInfo) GetLightRegisterAddress() int32 {
	if x != nil {
		return x.LightRegisterAddress
	}
	return 0
}

func (x *MaterialChannelLayerInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *MaterialChannelLayerInfo) GetProductionStationID() string {
	if x != nil {
		return x.ProductionStationID
	}
	return ""
}

func (x *MaterialChannelLayerInfo) GetProductionStation() *ProductionStationInfo {
	if x != nil {
		return x.ProductionStation
	}
	return nil
}

type QueryMaterialChannelLayerRequest struct {
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

func (x *QueryMaterialChannelLayerRequest) Reset() {
	*x = QueryMaterialChannelLayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_channel_layer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMaterialChannelLayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMaterialChannelLayerRequest) ProtoMessage() {}

func (x *QueryMaterialChannelLayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_material_channel_layer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMaterialChannelLayerRequest.ProtoReflect.Descriptor instead.
func (*QueryMaterialChannelLayerRequest) Descriptor() ([]byte, []int) {
	return file_material_channel_layer_proto_rawDescGZIP(), []int{1}
}

func (x *QueryMaterialChannelLayerRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryMaterialChannelLayerRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryMaterialChannelLayerRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

type QueryMaterialChannelLayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                        `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialChannelLayerInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                       `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                       `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                       `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryMaterialChannelLayerResponse) Reset() {
	*x = QueryMaterialChannelLayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_channel_layer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMaterialChannelLayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMaterialChannelLayerResponse) ProtoMessage() {}

func (x *QueryMaterialChannelLayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_channel_layer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMaterialChannelLayerResponse.ProtoReflect.Descriptor instead.
func (*QueryMaterialChannelLayerResponse) Descriptor() ([]byte, []int) {
	return file_material_channel_layer_proto_rawDescGZIP(), []int{2}
}

func (x *QueryMaterialChannelLayerResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryMaterialChannelLayerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryMaterialChannelLayerResponse) GetData() []*MaterialChannelLayerInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryMaterialChannelLayerResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryMaterialChannelLayerResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryMaterialChannelLayerResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllMaterialChannelLayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                        `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialChannelLayerInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllMaterialChannelLayerResponse) Reset() {
	*x = GetAllMaterialChannelLayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_channel_layer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMaterialChannelLayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMaterialChannelLayerResponse) ProtoMessage() {}

func (x *GetAllMaterialChannelLayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_channel_layer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMaterialChannelLayerResponse.ProtoReflect.Descriptor instead.
func (*GetAllMaterialChannelLayerResponse) Descriptor() ([]byte, []int) {
	return file_material_channel_layer_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllMaterialChannelLayerResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllMaterialChannelLayerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllMaterialChannelLayerResponse) GetData() []*MaterialChannelLayerInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetMaterialChannelLayerDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                      `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                    `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *MaterialChannelLayerInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetMaterialChannelLayerDetailResponse) Reset() {
	*x = GetMaterialChannelLayerDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_channel_layer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMaterialChannelLayerDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMaterialChannelLayerDetailResponse) ProtoMessage() {}

func (x *GetMaterialChannelLayerDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_channel_layer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMaterialChannelLayerDetailResponse.ProtoReflect.Descriptor instead.
func (*GetMaterialChannelLayerDetailResponse) Descriptor() ([]byte, []int) {
	return file_material_channel_layer_proto_rawDescGZIP(), []int{4}
}

func (x *GetMaterialChannelLayerDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetMaterialChannelLayerDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetMaterialChannelLayerDetailResponse) GetData() *MaterialChannelLayerInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type MaterialChannelInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 顺序
	SortIndex int32 `protobuf:"varint,2,opt,name=sortIndex,proto3" json:"sortIndex"`
	// 代号
	Code string `protobuf:"bytes,3,opt,name=code,proto3" json:"code"`
	// 描述
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description"`
	// 尺寸
	Size string `protobuf:"bytes,5,opt,name=size,proto3" json:"size"`
	// 规格
	Spec float32 `protobuf:"fixed32,6,opt,name=spec,proto3" json:"spec"`
	// 启用监控
	EnableMonitor bool `protobuf:"varint,7,opt,name=enableMonitor,proto3" json:"enableMonitor"`
	// 当前状态
	CurrentState string `protobuf:"bytes,8,opt,name=currentState,proto3" json:"currentState"`
	// 状态变更时间
	LastUpdateTime string `protobuf:"bytes,9,opt,name=lastUpdateTime,proto3" json:"lastUpdateTime"`
	// 备注
	Remark string `protobuf:"bytes,10,opt,name=remark,proto3" json:"remark"`
	// 物料通道层ID
	MaterialChannelLayerID string `protobuf:"bytes,11,opt,name=materialChannelLayerID,proto3" json:"materialChannelLayerID"`
	// 物料通道层
	MaterialChannelLayer *MaterialChannelLayerInfo `protobuf:"bytes,12,opt,name=materialChannelLayer,proto3" json:"materialChannelLayer"`
	// 物料信息ID
	MaterialInfoID string `protobuf:"bytes,13,opt,name=materialInfoID,proto3" json:"materialInfoID"`
	// 物料通道层
	MaterialInfo *MaterialInfoInfo `protobuf:"bytes,14,opt,name=materialInfo,proto3" json:"materialInfo"`
}

func (x *MaterialChannelInfo) Reset() {
	*x = MaterialChannelInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_channel_layer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialChannelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialChannelInfo) ProtoMessage() {}

func (x *MaterialChannelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_material_channel_layer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialChannelInfo.ProtoReflect.Descriptor instead.
func (*MaterialChannelInfo) Descriptor() ([]byte, []int) {
	return file_material_channel_layer_proto_rawDescGZIP(), []int{5}
}

func (x *MaterialChannelInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MaterialChannelInfo) GetSortIndex() int32 {
	if x != nil {
		return x.SortIndex
	}
	return 0
}

func (x *MaterialChannelInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MaterialChannelInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MaterialChannelInfo) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *MaterialChannelInfo) GetSpec() float32 {
	if x != nil {
		return x.Spec
	}
	return 0
}

func (x *MaterialChannelInfo) GetEnableMonitor() bool {
	if x != nil {
		return x.EnableMonitor
	}
	return false
}

func (x *MaterialChannelInfo) GetCurrentState() string {
	if x != nil {
		return x.CurrentState
	}
	return ""
}

func (x *MaterialChannelInfo) GetLastUpdateTime() string {
	if x != nil {
		return x.LastUpdateTime
	}
	return ""
}

func (x *MaterialChannelInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *MaterialChannelInfo) GetMaterialChannelLayerID() string {
	if x != nil {
		return x.MaterialChannelLayerID
	}
	return ""
}

func (x *MaterialChannelInfo) GetMaterialChannelLayer() *MaterialChannelLayerInfo {
	if x != nil {
		return x.MaterialChannelLayer
	}
	return nil
}

func (x *MaterialChannelInfo) GetMaterialInfoID() string {
	if x != nil {
		return x.MaterialInfoID
	}
	return ""
}

func (x *MaterialChannelInfo) GetMaterialInfo() *MaterialInfoInfo {
	if x != nil {
		return x.MaterialInfo
	}
	return nil
}

type GetMaterialChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductionStationID string `protobuf:"bytes,1,opt,name=productionStationID,proto3" json:"productionStationID"`
}

func (x *GetMaterialChannelRequest) Reset() {
	*x = GetMaterialChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_channel_layer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMaterialChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMaterialChannelRequest) ProtoMessage() {}

func (x *GetMaterialChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_material_channel_layer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMaterialChannelRequest.ProtoReflect.Descriptor instead.
func (*GetMaterialChannelRequest) Descriptor() ([]byte, []int) {
	return file_material_channel_layer_proto_rawDescGZIP(), []int{6}
}

func (x *GetMaterialChannelRequest) GetProductionStationID() string {
	if x != nil {
		return x.ProductionStationID
	}
	return ""
}

type GetAllMaterialChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                   `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MaterialChannelInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllMaterialChannelResponse) Reset() {
	*x = GetAllMaterialChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_material_channel_layer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMaterialChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMaterialChannelResponse) ProtoMessage() {}

func (x *GetAllMaterialChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_material_channel_layer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMaterialChannelResponse.ProtoReflect.Descriptor instead.
func (*GetAllMaterialChannelResponse) Descriptor() ([]byte, []int) {
	return file_material_channel_layer_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllMaterialChannelResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllMaterialChannelResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllMaterialChannelResponse) GetData() []*MaterialChannelInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_material_channel_layer_proto protoreflect.FileDescriptor

var file_material_channel_layer_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x02, 0x0a, 0x18, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x15, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x4a, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x7c, 0x0a, 0x20, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0xd9, 0x01, 0x0a, 0x21, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x94, 0x01, 0x0a,
	0x22, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x97, 0x01, 0x0a, 0x25, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9d, 0x04,
	0x0a, 0x13, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6c,
	0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x36, 0x0a, 0x16, 0x6d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x53, 0x0a, 0x14, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x14, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44,
	0x12, 0x3b, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0c, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4d, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x8a, 0x01, 0x0a,
	0x1d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_material_channel_layer_proto_rawDescOnce sync.Once
	file_material_channel_layer_proto_rawDescData = file_material_channel_layer_proto_rawDesc
)

func file_material_channel_layer_proto_rawDescGZIP() []byte {
	file_material_channel_layer_proto_rawDescOnce.Do(func() {
		file_material_channel_layer_proto_rawDescData = protoimpl.X.CompressGZIP(file_material_channel_layer_proto_rawDescData)
	})
	return file_material_channel_layer_proto_rawDescData
}

var file_material_channel_layer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_material_channel_layer_proto_goTypes = []interface{}{
	(*MaterialChannelLayerInfo)(nil),              // 0: proto.MaterialChannelLayerInfo
	(*QueryMaterialChannelLayerRequest)(nil),      // 1: proto.QueryMaterialChannelLayerRequest
	(*QueryMaterialChannelLayerResponse)(nil),     // 2: proto.QueryMaterialChannelLayerResponse
	(*GetAllMaterialChannelLayerResponse)(nil),    // 3: proto.GetAllMaterialChannelLayerResponse
	(*GetMaterialChannelLayerDetailResponse)(nil), // 4: proto.GetMaterialChannelLayerDetailResponse
	(*MaterialChannelInfo)(nil),                   // 5: proto.MaterialChannelInfo
	(*GetMaterialChannelRequest)(nil),             // 6: proto.GetMaterialChannelRequest
	(*GetAllMaterialChannelResponse)(nil),         // 7: proto.GetAllMaterialChannelResponse
	(*ProductionStationInfo)(nil),                 // 8: proto.ProductionStationInfo
	(Code)(0),                                     // 9: proto.Code
	(*MaterialInfoInfo)(nil),                      // 10: proto.MaterialInfoInfo
}
var file_material_channel_layer_proto_depIdxs = []int32{
	8,  // 0: proto.MaterialChannelLayerInfo.productionStation:type_name -> proto.ProductionStationInfo
	9,  // 1: proto.QueryMaterialChannelLayerResponse.code:type_name -> proto.Code
	0,  // 2: proto.QueryMaterialChannelLayerResponse.data:type_name -> proto.MaterialChannelLayerInfo
	9,  // 3: proto.GetAllMaterialChannelLayerResponse.code:type_name -> proto.Code
	0,  // 4: proto.GetAllMaterialChannelLayerResponse.data:type_name -> proto.MaterialChannelLayerInfo
	9,  // 5: proto.GetMaterialChannelLayerDetailResponse.code:type_name -> proto.Code
	0,  // 6: proto.GetMaterialChannelLayerDetailResponse.data:type_name -> proto.MaterialChannelLayerInfo
	0,  // 7: proto.MaterialChannelInfo.materialChannelLayer:type_name -> proto.MaterialChannelLayerInfo
	10, // 8: proto.MaterialChannelInfo.materialInfo:type_name -> proto.MaterialInfoInfo
	9,  // 9: proto.GetAllMaterialChannelResponse.code:type_name -> proto.Code
	5,  // 10: proto.GetAllMaterialChannelResponse.data:type_name -> proto.MaterialChannelInfo
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_material_channel_layer_proto_init() }
func file_material_channel_layer_proto_init() {
	if File_material_channel_layer_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_production_line_proto_init()
	file_material_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_material_channel_layer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialChannelLayerInfo); i {
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
		file_material_channel_layer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMaterialChannelLayerRequest); i {
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
		file_material_channel_layer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMaterialChannelLayerResponse); i {
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
		file_material_channel_layer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMaterialChannelLayerResponse); i {
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
		file_material_channel_layer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMaterialChannelLayerDetailResponse); i {
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
		file_material_channel_layer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialChannelInfo); i {
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
		file_material_channel_layer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMaterialChannelRequest); i {
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
		file_material_channel_layer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMaterialChannelResponse); i {
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
			RawDescriptor: file_material_channel_layer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_material_channel_layer_proto_goTypes,
		DependencyIndexes: file_material_channel_layer_proto_depIdxs,
		MessageInfos:      file_material_channel_layer_proto_msgTypes,
	}.Build()
	File_material_channel_layer_proto = out.File
	file_material_channel_layer_proto_rawDesc = nil
	file_material_channel_layer_proto_goTypes = nil
	file_material_channel_layer_proto_depIdxs = nil
}
