// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: product_process_route.proto

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

type GetProductProcessRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductInfoID    string   `protobuf:"bytes,1,opt,name=productInfoID,proto3" json:"productInfoID"`
	CurrentProcessID string   `protobuf:"bytes,2,opt,name=currentProcessID,proto3" json:"currentProcessID"`
	CurrentStates    []string `protobuf:"bytes,3,rep,name=currentStates,proto3" json:"currentStates"`
}

func (x *GetProductProcessRouteRequest) Reset() {
	*x = GetProductProcessRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_process_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductProcessRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductProcessRouteRequest) ProtoMessage() {}

func (x *GetProductProcessRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_process_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductProcessRouteRequest.ProtoReflect.Descriptor instead.
func (*GetProductProcessRouteRequest) Descriptor() ([]byte, []int) {
	return file_product_process_route_proto_rawDescGZIP(), []int{0}
}

func (x *GetProductProcessRouteRequest) GetProductInfoID() string {
	if x != nil {
		return x.ProductInfoID
	}
	return ""
}

func (x *GetProductProcessRouteRequest) GetCurrentProcessID() string {
	if x != nil {
		return x.CurrentProcessID
	}
	return ""
}

func (x *GetProductProcessRouteRequest) GetCurrentStates() []string {
	if x != nil {
		return x.CurrentStates
	}
	return nil
}

type ProductProcessRouteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 作业顺序
	WorkIndex int32 `protobuf:"varint,2,opt,name=workIndex,proto3" json:"workIndex"`
	// 工序顺序
	RouteIndex int32 `protobuf:"varint,3,opt,name=routeIndex,proto3" json:"routeIndex"`
	// 创建时间
	CreateTime string `protobuf:"bytes,4,opt,name=createTime,proto3" json:"createTime"`
	// 当前状态
	CurrentState string `protobuf:"bytes,5,opt,name=currentState,proto3" json:"currentState"`
	// 更新时间
	LastUpdateTime string `protobuf:"bytes,6,opt,name=lastUpdateTime,proto3" json:"lastUpdateTime"`
	// 备注
	Remark string `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark"`
	// 上步工序ID
	LastProcessID string                 `protobuf:"bytes,8,opt,name=lastProcessID,proto3" json:"lastProcessID"`
	LastProcess   *ProductionProcessInfo `protobuf:"bytes,9,opt,name=lastProcess,proto3" json:"lastProcess"`
	// 当前工序ID
	CurrentProcessID string                 `protobuf:"bytes,10,opt,name=currentProcessID,proto3" json:"currentProcessID"`
	CurrentProcess   *ProductionProcessInfo `protobuf:"bytes,11,opt,name=currentProcess,proto3" json:"currentProcess"`
	// 执行工站ID
	ProductionStationID string                 `protobuf:"bytes,12,opt,name=productionStationID,proto3" json:"productionStationID"`
	ProductionStation   *ProductionStationInfo `protobuf:"bytes,13,opt,name=productionStation,proto3" json:"productionStation"`
	// 执行人员ID
	ProcessUserID string `protobuf:"bytes,14,opt,name=processUserID,proto3" json:"processUserID"`
	// 产品信息ID
	ProductInfoID string           `protobuf:"bytes,15,opt,name=productInfoID,proto3" json:"productInfoID"`
	ProductInfo   *ProductInfoInfo `protobuf:"bytes,16,opt,name=productInfo,proto3" json:"productInfo"`
}

func (x *ProductProcessRouteInfo) Reset() {
	*x = ProductProcessRouteInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_process_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductProcessRouteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductProcessRouteInfo) ProtoMessage() {}

func (x *ProductProcessRouteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_process_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductProcessRouteInfo.ProtoReflect.Descriptor instead.
func (*ProductProcessRouteInfo) Descriptor() ([]byte, []int) {
	return file_product_process_route_proto_rawDescGZIP(), []int{1}
}

func (x *ProductProcessRouteInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductProcessRouteInfo) GetWorkIndex() int32 {
	if x != nil {
		return x.WorkIndex
	}
	return 0
}

func (x *ProductProcessRouteInfo) GetRouteIndex() int32 {
	if x != nil {
		return x.RouteIndex
	}
	return 0
}

func (x *ProductProcessRouteInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *ProductProcessRouteInfo) GetCurrentState() string {
	if x != nil {
		return x.CurrentState
	}
	return ""
}

func (x *ProductProcessRouteInfo) GetLastUpdateTime() string {
	if x != nil {
		return x.LastUpdateTime
	}
	return ""
}

func (x *ProductProcessRouteInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ProductProcessRouteInfo) GetLastProcessID() string {
	if x != nil {
		return x.LastProcessID
	}
	return ""
}

func (x *ProductProcessRouteInfo) GetLastProcess() *ProductionProcessInfo {
	if x != nil {
		return x.LastProcess
	}
	return nil
}

func (x *ProductProcessRouteInfo) GetCurrentProcessID() string {
	if x != nil {
		return x.CurrentProcessID
	}
	return ""
}

func (x *ProductProcessRouteInfo) GetCurrentProcess() *ProductionProcessInfo {
	if x != nil {
		return x.CurrentProcess
	}
	return nil
}

func (x *ProductProcessRouteInfo) GetProductionStationID() string {
	if x != nil {
		return x.ProductionStationID
	}
	return ""
}

func (x *ProductProcessRouteInfo) GetProductionStation() *ProductionStationInfo {
	if x != nil {
		return x.ProductionStation
	}
	return nil
}

func (x *ProductProcessRouteInfo) GetProcessUserID() string {
	if x != nil {
		return x.ProcessUserID
	}
	return ""
}

func (x *ProductProcessRouteInfo) GetProductInfoID() string {
	if x != nil {
		return x.ProductInfoID
	}
	return ""
}

func (x *ProductProcessRouteInfo) GetProductInfo() *ProductInfoInfo {
	if x != nil {
		return x.ProductInfo
	}
	return nil
}

type QueryProductProcessRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"sortConfig" form:"sortConfig"
	SortConfig string `protobuf:"bytes,3,opt,name=sortConfig,proto3" json:"sortConfig" uri:"sortConfig" form:"sortConfig"`
	// 当前工序
	// @inject_tag: uri:"currentProcessID" form:"currentProcessID"
	CurrentProcessID string `protobuf:"bytes,4,opt,name=currentProcessID,proto3" json:"currentProcessID" uri:"currentProcessID" form:"currentProcessID"`
	// 序列号
	// @inject_tag: uri:"productSerialNo" form:"productSerialNo"
	ProductSerialNo string `protobuf:"bytes,5,opt,name=productSerialNo,proto3" json:"productSerialNo" uri:"productSerialNo" form:"productSerialNo"`
	// 生产工单号
	// @inject_tag: uri:"productOrderNo" form:"productOrderNo"
	ProductOrderNo string `protobuf:"bytes,6,opt,name=productOrderNo,proto3" json:"productOrderNo" uri:"productOrderNo" form:"productOrderNo"`
	// 创建时间开始
	// @inject_tag: uri:"createTime0" form:"createTime0"
	CreateTime0 string `protobuf:"bytes,7,opt,name=createTime0,proto3" json:"createTime0" uri:"createTime0" form:"createTime0"`
	// 创建时间结束
	// @inject_tag: uri:"createTime1" form:"createTime1"
	CreateTime1 string `protobuf:"bytes,8,opt,name=createTime1,proto3" json:"createTime1" uri:"createTime1" form:"createTime1"`
	// @inject_tag: uri:"productInfoID" form:"productInfoID"
	ProductInfoID string `protobuf:"bytes,9,opt,name=productInfoID,proto3" json:"productInfoID" uri:"productInfoID" form:"productInfoID"`
	// @inject_tag: uri:"routeIndex" form:"routeIndex"
	RouteIndex int32 `protobuf:"varint,10,opt,name=routeIndex,proto3" json:"routeIndex" uri:"routeIndex" form:"routeIndex"`
	// @inject_tag: uri:"currentState" form:"currentState"
	CurrentState string `protobuf:"bytes,11,opt,name=currentState,proto3" json:"currentState" uri:"currentState" form:"currentState"`
}

func (x *QueryProductProcessRouteRequest) Reset() {
	*x = QueryProductProcessRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_process_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductProcessRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductProcessRouteRequest) ProtoMessage() {}

func (x *QueryProductProcessRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_process_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductProcessRouteRequest.ProtoReflect.Descriptor instead.
func (*QueryProductProcessRouteRequest) Descriptor() ([]byte, []int) {
	return file_product_process_route_proto_rawDescGZIP(), []int{2}
}

func (x *QueryProductProcessRouteRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryProductProcessRouteRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryProductProcessRouteRequest) GetSortConfig() string {
	if x != nil {
		return x.SortConfig
	}
	return ""
}

func (x *QueryProductProcessRouteRequest) GetCurrentProcessID() string {
	if x != nil {
		return x.CurrentProcessID
	}
	return ""
}

func (x *QueryProductProcessRouteRequest) GetProductSerialNo() string {
	if x != nil {
		return x.ProductSerialNo
	}
	return ""
}

func (x *QueryProductProcessRouteRequest) GetProductOrderNo() string {
	if x != nil {
		return x.ProductOrderNo
	}
	return ""
}

func (x *QueryProductProcessRouteRequest) GetCreateTime0() string {
	if x != nil {
		return x.CreateTime0
	}
	return ""
}

func (x *QueryProductProcessRouteRequest) GetCreateTime1() string {
	if x != nil {
		return x.CreateTime1
	}
	return ""
}

func (x *QueryProductProcessRouteRequest) GetProductInfoID() string {
	if x != nil {
		return x.ProductInfoID
	}
	return ""
}

func (x *QueryProductProcessRouteRequest) GetRouteIndex() int32 {
	if x != nil {
		return x.RouteIndex
	}
	return 0
}

func (x *QueryProductProcessRouteRequest) GetCurrentState() string {
	if x != nil {
		return x.CurrentState
	}
	return ""
}

type QueryProductProcessRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                       `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                     `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductProcessRouteInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64                      `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64                      `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64                      `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryProductProcessRouteResponse) Reset() {
	*x = QueryProductProcessRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_process_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductProcessRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductProcessRouteResponse) ProtoMessage() {}

func (x *QueryProductProcessRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_process_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductProcessRouteResponse.ProtoReflect.Descriptor instead.
func (*QueryProductProcessRouteResponse) Descriptor() ([]byte, []int) {
	return file_product_process_route_proto_rawDescGZIP(), []int{3}
}

func (x *QueryProductProcessRouteResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryProductProcessRouteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryProductProcessRouteResponse) GetData() []*ProductProcessRouteInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductProcessRouteResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryProductProcessRouteResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryProductProcessRouteResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllProductProcessRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                       `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                     `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*ProductProcessRouteInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllProductProcessRouteResponse) Reset() {
	*x = GetAllProductProcessRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_process_route_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductProcessRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductProcessRouteResponse) ProtoMessage() {}

func (x *GetAllProductProcessRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_process_route_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductProcessRouteResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductProcessRouteResponse) Descriptor() ([]byte, []int) {
	return file_product_process_route_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllProductProcessRouteResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllProductProcessRouteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProductProcessRouteResponse) GetData() []*ProductProcessRouteInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductProcessRouteDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code                     `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Code" json:"code"`
	Message string                   `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *ProductProcessRouteInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetProductProcessRouteDetailResponse) Reset() {
	*x = GetProductProcessRouteDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_process_route_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductProcessRouteDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductProcessRouteDetailResponse) ProtoMessage() {}

func (x *GetProductProcessRouteDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_process_route_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductProcessRouteDetailResponse.ProtoReflect.Descriptor instead.
func (*GetProductProcessRouteDetailResponse) Descriptor() ([]byte, []int) {
	return file_product_process_route_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductProcessRouteDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetProductProcessRouteDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductProcessRouteDetailResponse) GetData() *ProductProcessRouteInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_product_process_route_proto protoreflect.FileDescriptor

var file_product_process_route_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0xc7, 0x05, 0x0a,
	0x17, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x77, 0x6f, 0x72,
	0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x61,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x61,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44,
	0x12, 0x3e, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x2a, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x12, 0x44, 0x0a, 0x0e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x12, 0x4a, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xa7, 0x03, 0x0a, 0x1f, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44,
	0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x4e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x4e, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x30, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x30, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x31, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x31, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x0c,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x22, 0xd7, 0x01, 0x0a, 0x20, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x92, 0x01, 0x0a, 0x21, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x95, 0x01, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x8d, 0x02, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12,
	0x3e, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x05, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_process_route_proto_rawDescOnce sync.Once
	file_product_process_route_proto_rawDescData = file_product_process_route_proto_rawDesc
)

func file_product_process_route_proto_rawDescGZIP() []byte {
	file_product_process_route_proto_rawDescOnce.Do(func() {
		file_product_process_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_process_route_proto_rawDescData)
	})
	return file_product_process_route_proto_rawDescData
}

var file_product_process_route_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_product_process_route_proto_goTypes = []interface{}{
	(*GetProductProcessRouteRequest)(nil),        // 0: proto.GetProductProcessRouteRequest
	(*ProductProcessRouteInfo)(nil),              // 1: proto.ProductProcessRouteInfo
	(*QueryProductProcessRouteRequest)(nil),      // 2: proto.QueryProductProcessRouteRequest
	(*QueryProductProcessRouteResponse)(nil),     // 3: proto.QueryProductProcessRouteResponse
	(*GetAllProductProcessRouteResponse)(nil),    // 4: proto.GetAllProductProcessRouteResponse
	(*GetProductProcessRouteDetailResponse)(nil), // 5: proto.GetProductProcessRouteDetailResponse
	(*ProductionProcessInfo)(nil),                // 6: proto.ProductionProcessInfo
	(*ProductionStationInfo)(nil),                // 7: proto.ProductionStationInfo
	(*ProductInfoInfo)(nil),                      // 8: proto.ProductInfoInfo
	(Code)(0),                                    // 9: proto.Code
	(*CommonResponse)(nil),                       // 10: proto.CommonResponse
}
var file_product_process_route_proto_depIdxs = []int32{
	6,  // 0: proto.ProductProcessRouteInfo.lastProcess:type_name -> proto.ProductionProcessInfo
	6,  // 1: proto.ProductProcessRouteInfo.currentProcess:type_name -> proto.ProductionProcessInfo
	7,  // 2: proto.ProductProcessRouteInfo.productionStation:type_name -> proto.ProductionStationInfo
	8,  // 3: proto.ProductProcessRouteInfo.productInfo:type_name -> proto.ProductInfoInfo
	9,  // 4: proto.QueryProductProcessRouteResponse.code:type_name -> proto.Code
	1,  // 5: proto.QueryProductProcessRouteResponse.data:type_name -> proto.ProductProcessRouteInfo
	9,  // 6: proto.GetAllProductProcessRouteResponse.code:type_name -> proto.Code
	1,  // 7: proto.GetAllProductProcessRouteResponse.data:type_name -> proto.ProductProcessRouteInfo
	9,  // 8: proto.GetProductProcessRouteDetailResponse.code:type_name -> proto.Code
	1,  // 9: proto.GetProductProcessRouteDetailResponse.data:type_name -> proto.ProductProcessRouteInfo
	1,  // 10: proto.ProductProcessRoute.Add:input_type -> proto.ProductProcessRouteInfo
	0,  // 11: proto.ProductProcessRoute.Get:input_type -> proto.GetProductProcessRouteRequest
	2,  // 12: proto.ProductProcessRoute.Query:input_type -> proto.QueryProductProcessRouteRequest
	10, // 13: proto.ProductProcessRoute.Add:output_type -> proto.CommonResponse
	5,  // 14: proto.ProductProcessRoute.Get:output_type -> proto.GetProductProcessRouteDetailResponse
	3,  // 15: proto.ProductProcessRoute.Query:output_type -> proto.QueryProductProcessRouteResponse
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_product_process_route_proto_init() }
func file_product_process_route_proto_init() {
	if File_product_process_route_proto != nil {
		return
	}
	file_mom_common_proto_init()
	file_production_line_proto_init()
	file_product_order_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_process_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductProcessRouteRequest); i {
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
		file_product_process_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductProcessRouteInfo); i {
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
		file_product_process_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductProcessRouteRequest); i {
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
		file_product_process_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductProcessRouteResponse); i {
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
		file_product_process_route_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductProcessRouteResponse); i {
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
		file_product_process_route_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductProcessRouteDetailResponse); i {
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
			RawDescriptor: file_product_process_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_process_route_proto_goTypes,
		DependencyIndexes: file_product_process_route_proto_depIdxs,
		MessageInfos:      file_product_process_route_proto_msgTypes,
	}.Build()
	File_product_process_route_proto = out.File
	file_product_process_route_proto_rawDesc = nil
	file_product_process_route_proto_goTypes = nil
	file_product_process_route_proto_depIdxs = nil
}
