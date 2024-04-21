// protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: layout/v1/layout.proto

package v1

import (
	reflect "reflect"
	pbentity "seat/api/pbentity"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationId int32 `protobuf:"varint,1,opt,name=LocationId,proto3" json:"LocationId,omitempty" v:"required"` // v:required
	PolicyCId  int32 `protobuf:"varint,2,opt,name=PolicyCId,proto3" json:"PolicyCId,omitempty" v:"required"`   // v:required
	// 如果已经使用了 PolicyCId，可以不传PolicyInfo
	PolicyInfo string `protobuf:"bytes,3,opt,name=PolicyInfo,proto3" json:"PolicyInfo,omitempty" dc:"如果已经使用了 PolicyCId，可以不传PolicyInfo" v:"json"` // v:json
	LayoutName string `protobuf:"bytes,4,opt,name=LayoutName,proto3" json:"LayoutName,omitempty" v:"required"`                                   // v:required
	Map        string `protobuf:"bytes,5,opt,name=Map,proto3" json:"Map,omitempty" v:"json"`                                                     // v:json
	Status     int32  `protobuf:"varint,6,opt,name=Status,proto3" json:"Status,omitempty" v:"required"`                                          // v:required
	Sort       int32  `protobuf:"varint,7,opt,name=Sort,proto3" json:"Sort,omitempty" v:"required"`                                              // v:required
}

func (x *CreateReq) Reset() {
	*x = CreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layout_v1_layout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReq) ProtoMessage() {}

func (x *CreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_layout_v1_layout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReq.ProtoReflect.Descriptor instead.
func (*CreateReq) Descriptor() ([]byte, []int) {
	return file_layout_v1_layout_proto_rawDescGZIP(), []int{0}
}

func (x *CreateReq) GetLocationId() int32 {
	if x != nil {
		return x.LocationId
	}
	return 0
}

func (x *CreateReq) GetPolicyCId() int32 {
	if x != nil {
		return x.PolicyCId
	}
	return 0
}

func (x *CreateReq) GetPolicyInfo() string {
	if x != nil {
		return x.PolicyInfo
	}
	return ""
}

func (x *CreateReq) GetLayoutName() string {
	if x != nil {
		return x.LayoutName
	}
	return ""
}

func (x *CreateReq) GetMap() string {
	if x != nil {
		return x.Map
	}
	return ""
}

func (x *CreateReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateReq) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

type CreateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" v:"required"` // v:required
}

func (x *CreateRes) Reset() {
	*x = CreateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layout_v1_layout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRes) ProtoMessage() {}

func (x *CreateRes) ProtoReflect() protoreflect.Message {
	mi := &file_layout_v1_layout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRes.ProtoReflect.Descriptor instead.
func (*CreateRes) Descriptor() ([]byte, []int) {
	return file_layout_v1_layout_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRes) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetOneReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" v:"required"` // v:required
}

func (x *GetOneReq) Reset() {
	*x = GetOneReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layout_v1_layout_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneReq) ProtoMessage() {}

func (x *GetOneReq) ProtoReflect() protoreflect.Message {
	mi := &file_layout_v1_layout_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneReq.ProtoReflect.Descriptor instead.
func (*GetOneReq) Descriptor() ([]byte, []int) {
	return file_layout_v1_layout_proto_rawDescGZIP(), []int{2}
}

func (x *GetOneReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetOneRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Layout     *pbentity.Layout `protobuf:"bytes,1,opt,name=layout,proto3" json:"layout,omitempty"`
	PolicyInfo string           `protobuf:"bytes,2,opt,name=PolicyInfo,proto3" json:"PolicyInfo,omitempty"`
}

func (x *GetOneRes) Reset() {
	*x = GetOneRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layout_v1_layout_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneRes) ProtoMessage() {}

func (x *GetOneRes) ProtoReflect() protoreflect.Message {
	mi := &file_layout_v1_layout_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneRes.ProtoReflect.Descriptor instead.
func (*GetOneRes) Descriptor() ([]byte, []int) {
	return file_layout_v1_layout_proto_rawDescGZIP(), []int{3}
}

func (x *GetOneRes) GetLayout() *pbentity.Layout {
	if x != nil {
		return x.Layout
	}
	return nil
}

func (x *GetOneRes) GetPolicyInfo() string {
	if x != nil {
		return x.PolicyInfo
	}
	return ""
}

type GetListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int32  `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty" v:"required"`                 // v:required
	PageSize   int32  `protobuf:"varint,2,opt,name=PageSize,proto3" json:"PageSize,omitempty" v:"required|max:100"` // v:required|max:100
	LayoutName string `protobuf:"bytes,3,opt,name=LayoutName,proto3" json:"LayoutName,omitempty" v:"length:1,20"`   // v:length:1,20
	Status     int32  `protobuf:"varint,4,opt,name=Status,proto3" json:"Status,omitempty"`
	SeatsMin   int32  `protobuf:"varint,5,opt,name=SeatsMin,proto3" json:"SeatsMin,omitempty" v:"length:0,99999999"` // v:length:0,99999999
	SeatsMax   int32  `protobuf:"varint,6,opt,name=SeatsMax,proto3" json:"SeatsMax,omitempty" v:"length:1,99999999"` // v:length:1,99999999
}

func (x *GetListReq) Reset() {
	*x = GetListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layout_v1_layout_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListReq) ProtoMessage() {}

func (x *GetListReq) ProtoReflect() protoreflect.Message {
	mi := &file_layout_v1_layout_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListReq.ProtoReflect.Descriptor instead.
func (*GetListReq) Descriptor() ([]byte, []int) {
	return file_layout_v1_layout_proto_rawDescGZIP(), []int{4}
}

func (x *GetListReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetListReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetListReq) GetLayoutName() string {
	if x != nil {
		return x.LayoutName
	}
	return ""
}

func (x *GetListReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetListReq) GetSeatsMin() int32 {
	if x != nil {
		return x.SeatsMin
	}
	return 0
}

func (x *GetListReq) GetSeatsMax() int32 {
	if x != nil {
		return x.SeatsMax
	}
	return 0
}

type LayoutList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LocationId   int32  `protobuf:"varint,2,opt,name=LocationId,proto3" json:"LocationId,omitempty"`
	LocationName string `protobuf:"bytes,3,opt,name=LocationName,proto3" json:"LocationName,omitempty"`
	LayoutName   string `protobuf:"bytes,4,opt,name=LayoutName,proto3" json:"LayoutName,omitempty"`
	Status       int32  `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
	Sort         int32  `protobuf:"varint,6,opt,name=Sort,proto3" json:"Sort,omitempty"`
	Seats        int32  `protobuf:"varint,7,opt,name=Seats,proto3" json:"Seats,omitempty"`
}

func (x *LayoutList) Reset() {
	*x = LayoutList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layout_v1_layout_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LayoutList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LayoutList) ProtoMessage() {}

func (x *LayoutList) ProtoReflect() protoreflect.Message {
	mi := &file_layout_v1_layout_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LayoutList.ProtoReflect.Descriptor instead.
func (*LayoutList) Descriptor() ([]byte, []int) {
	return file_layout_v1_layout_proto_rawDescGZIP(), []int{5}
}

func (x *LayoutList) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LayoutList) GetLocationId() int32 {
	if x != nil {
		return x.LocationId
	}
	return 0
}

func (x *LayoutList) GetLocationName() string {
	if x != nil {
		return x.LocationName
	}
	return ""
}

func (x *LayoutList) GetLayoutName() string {
	if x != nil {
		return x.LayoutName
	}
	return ""
}

func (x *LayoutList) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *LayoutList) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *LayoutList) GetSeats() int32 {
	if x != nil {
		return x.Seats
	}
	return 0
}

type GetListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Layouts []*LayoutList `protobuf:"bytes,1,rep,name=layouts,proto3" json:"layouts,omitempty"`
	Total   int32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetListRes) Reset() {
	*x = GetListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layout_v1_layout_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRes) ProtoMessage() {}

func (x *GetListRes) ProtoReflect() protoreflect.Message {
	mi := &file_layout_v1_layout_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListRes.ProtoReflect.Descriptor instead.
func (*GetListRes) Descriptor() ([]byte, []int) {
	return file_layout_v1_layout_proto_rawDescGZIP(), []int{6}
}

func (x *GetListRes) GetLayouts() []*LayoutList {
	if x != nil {
		return x.Layouts
	}
	return nil
}

func (x *GetListRes) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" v:"required"`                 // v:required
	LocationId int32  `protobuf:"varint,2,opt,name=LocationId,proto3" json:"LocationId,omitempty" v:"required"` // v:required
	PolicyCId  int32  `protobuf:"varint,3,opt,name=PolicyCId,proto3" json:"PolicyCId,omitempty" v:"required"`   // v:required
	PolicyInfo string `protobuf:"bytes,4,opt,name=PolicyInfo,proto3" json:"PolicyInfo,omitempty" v:"json"`      // v:json
	LayoutName string `protobuf:"bytes,5,opt,name=LayoutName,proto3" json:"LayoutName,omitempty" v:"required"`  // v:required
	Map        string `protobuf:"bytes,6,opt,name=Map,proto3" json:"Map,omitempty" v:"json"`                    // v:json
	Status     int32  `protobuf:"varint,7,opt,name=Status,proto3" json:"Status,omitempty" v:"required"`         // v:required
	Sort       int32  `protobuf:"varint,8,opt,name=Sort,proto3" json:"Sort,omitempty" v:"required"`             // v:required
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layout_v1_layout_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_layout_v1_layout_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_layout_v1_layout_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateReq) GetLocationId() int32 {
	if x != nil {
		return x.LocationId
	}
	return 0
}

func (x *UpdateReq) GetPolicyCId() int32 {
	if x != nil {
		return x.PolicyCId
	}
	return 0
}

func (x *UpdateReq) GetPolicyInfo() string {
	if x != nil {
		return x.PolicyInfo
	}
	return ""
}

func (x *UpdateReq) GetLayoutName() string {
	if x != nil {
		return x.LayoutName
	}
	return ""
}

func (x *UpdateReq) GetMap() string {
	if x != nil {
		return x.Map
	}
	return ""
}

func (x *UpdateReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateReq) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

type UpdateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Empty *emptypb.Empty `protobuf:"bytes,1,opt,name=empty,proto3" json:"empty,omitempty"`
}

func (x *UpdateRes) Reset() {
	*x = UpdateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layout_v1_layout_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRes) ProtoMessage() {}

func (x *UpdateRes) ProtoReflect() protoreflect.Message {
	mi := &file_layout_v1_layout_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRes.ProtoReflect.Descriptor instead.
func (*UpdateRes) Descriptor() ([]byte, []int) {
	return file_layout_v1_layout_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateRes) GetEmpty() *emptypb.Empty {
	if x != nil {
		return x.Empty
	}
	return nil
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" v:"required"` // v:required
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layout_v1_layout_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_layout_v1_layout_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_layout_v1_layout_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Empty *emptypb.Empty `protobuf:"bytes,1,opt,name=empty,proto3" json:"empty,omitempty"`
}

func (x *DeleteRes) Reset() {
	*x = DeleteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layout_v1_layout_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRes) ProtoMessage() {}

func (x *DeleteRes) ProtoReflect() protoreflect.Message {
	mi := &file_layout_v1_layout_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRes.ProtoReflect.Descriptor instead.
func (*DeleteRes) Descriptor() ([]byte, []int) {
	return file_layout_v1_layout_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteRes) GetEmpty() *emptypb.Empty {
	if x != nil {
		return x.Empty
	}
	return nil
}

var File_layout_v1_layout_proto protoreflect.FileDescriptor

var file_layout_v1_layout_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x79, 0x6f,
	0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x65, 0x61, 0x74, 0x1a, 0x15,
	0x70, 0x62, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e,
	0x0a, 0x0a, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x4d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x61, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x22, 0x1b, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x55, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65,
	0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x4c,
	0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x06, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xac, 0x01,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x61, 0x74, 0x73, 0x4d, 0x69, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x53, 0x65, 0x61, 0x74, 0x73, 0x4d, 0x69, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x61, 0x74, 0x73, 0x4d, 0x61, 0x78, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x53, 0x65, 0x61, 0x74, 0x73, 0x4d, 0x61, 0x78, 0x22, 0xc2, 0x01, 0x0a,
	0x0a, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x53,
	0x65, 0x61, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x65, 0x61, 0x74,
	0x73, 0x22, 0x4e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12,
	0x2a, 0x0a, 0x07, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x2e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x07, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0xd7, 0x01, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a,
	0x0a, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x4d, 0x61, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x61, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x22, 0x39, 0x0a, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x12, 0x2c, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xf1,
	0x01, 0x0a, 0x06, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x6e,
	0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x10, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x0f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x0f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f,
	0x2e, 0x73, 0x65, 0x61, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x0f, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x73, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c,
	0x69, 0x62, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_layout_v1_layout_proto_rawDescOnce sync.Once
	file_layout_v1_layout_proto_rawDescData = file_layout_v1_layout_proto_rawDesc
)

func file_layout_v1_layout_proto_rawDescGZIP() []byte {
	file_layout_v1_layout_proto_rawDescOnce.Do(func() {
		file_layout_v1_layout_proto_rawDescData = protoimpl.X.CompressGZIP(file_layout_v1_layout_proto_rawDescData)
	})
	return file_layout_v1_layout_proto_rawDescData
}

var file_layout_v1_layout_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_layout_v1_layout_proto_goTypes = []interface{}{
	(*CreateReq)(nil),       // 0: seat.CreateReq
	(*CreateRes)(nil),       // 1: seat.CreateRes
	(*GetOneReq)(nil),       // 2: seat.GetOneReq
	(*GetOneRes)(nil),       // 3: seat.GetOneRes
	(*GetListReq)(nil),      // 4: seat.GetListReq
	(*LayoutList)(nil),      // 5: seat.LayoutList
	(*GetListRes)(nil),      // 6: seat.GetListRes
	(*UpdateReq)(nil),       // 7: seat.UpdateReq
	(*UpdateRes)(nil),       // 8: seat.UpdateRes
	(*DeleteReq)(nil),       // 9: seat.DeleteReq
	(*DeleteRes)(nil),       // 10: seat.DeleteRes
	(*pbentity.Layout)(nil), // 11: pbentity.Layout
	(*emptypb.Empty)(nil),   // 12: google.protobuf.Empty
}
var file_layout_v1_layout_proto_depIdxs = []int32{
	11, // 0: seat.GetOneRes.layout:type_name -> pbentity.Layout
	5,  // 1: seat.GetListRes.layouts:type_name -> seat.LayoutList
	12, // 2: seat.UpdateRes.empty:type_name -> google.protobuf.Empty
	12, // 3: seat.DeleteRes.empty:type_name -> google.protobuf.Empty
	0,  // 4: seat.Layout.Create:input_type -> seat.CreateReq
	2,  // 5: seat.Layout.GetOne:input_type -> seat.GetOneReq
	4,  // 6: seat.Layout.GetList:input_type -> seat.GetListReq
	7,  // 7: seat.Layout.Update:input_type -> seat.UpdateReq
	9,  // 8: seat.Layout.Delete:input_type -> seat.DeleteReq
	1,  // 9: seat.Layout.Create:output_type -> seat.CreateRes
	3,  // 10: seat.Layout.GetOne:output_type -> seat.GetOneRes
	6,  // 11: seat.Layout.GetList:output_type -> seat.GetListRes
	8,  // 12: seat.Layout.Update:output_type -> seat.UpdateRes
	10, // 13: seat.Layout.Delete:output_type -> seat.DeleteRes
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_layout_v1_layout_proto_init() }
func file_layout_v1_layout_proto_init() {
	if File_layout_v1_layout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_layout_v1_layout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReq); i {
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
		file_layout_v1_layout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRes); i {
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
		file_layout_v1_layout_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneReq); i {
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
		file_layout_v1_layout_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneRes); i {
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
		file_layout_v1_layout_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListReq); i {
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
		file_layout_v1_layout_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LayoutList); i {
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
		file_layout_v1_layout_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListRes); i {
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
		file_layout_v1_layout_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
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
		file_layout_v1_layout_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRes); i {
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
		file_layout_v1_layout_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReq); i {
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
		file_layout_v1_layout_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRes); i {
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
			RawDescriptor: file_layout_v1_layout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_layout_v1_layout_proto_goTypes,
		DependencyIndexes: file_layout_v1_layout_proto_depIdxs,
		MessageInfos:      file_layout_v1_layout_proto_msgTypes,
	}.Build()
	File_layout_v1_layout_proto = out.File
	file_layout_v1_layout_proto_rawDesc = nil
	file_layout_v1_layout_proto_goTypes = nil
	file_layout_v1_layout_proto_depIdxs = nil
}