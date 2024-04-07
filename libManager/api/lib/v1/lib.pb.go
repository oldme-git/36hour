// protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.0--rc3
// source: lib/v1/lib.proto

package v1

import (
	pbentity "libManager/api/pbentity"
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

	LibName string `protobuf:"bytes,1,opt,name=libName,proto3" json:"libName,omitempty" v:"required"` // v:required
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" v:"required"` // v:required
	Active  bool   `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty" v:"required"`  // v:required
}

func (x *CreateReq) Reset() {
	*x = CreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_v1_lib_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReq) ProtoMessage() {}

func (x *CreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_lib_v1_lib_proto_msgTypes[0]
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
	return file_lib_v1_lib_proto_rawDescGZIP(), []int{0}
}

func (x *CreateReq) GetLibName() string {
	if x != nil {
		return x.LibName
	}
	return ""
}

func (x *CreateReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateReq) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type CreateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateRes) Reset() {
	*x = CreateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_v1_lib_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRes) ProtoMessage() {}

func (x *CreateRes) ProtoReflect() protoreflect.Message {
	mi := &file_lib_v1_lib_proto_msgTypes[1]
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
	return file_lib_v1_lib_proto_rawDescGZIP(), []int{1}
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
		mi := &file_lib_v1_lib_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneReq) ProtoMessage() {}

func (x *GetOneReq) ProtoReflect() protoreflect.Message {
	mi := &file_lib_v1_lib_proto_msgTypes[2]
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
	return file_lib_v1_lib_proto_rawDescGZIP(), []int{2}
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

	Lib *pbentity.Lib `protobuf:"bytes,1,opt,name=lib,proto3" json:"lib,omitempty"`
}

func (x *GetOneRes) Reset() {
	*x = GetOneRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_v1_lib_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneRes) ProtoMessage() {}

func (x *GetOneRes) ProtoReflect() protoreflect.Message {
	mi := &file_lib_v1_lib_proto_msgTypes[3]
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
	return file_lib_v1_lib_proto_rawDescGZIP(), []int{3}
}

func (x *GetOneRes) GetLib() *pbentity.Lib {
	if x != nil {
		return x.Lib
	}
	return nil
}

type GetListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty" v:"required"`                 // v:required
	PageSize int32  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty" v:"required|max:100"` // v:required|max:100
	LibName  string `protobuf:"bytes,3,opt,name=libName,proto3" json:"libName,omitempty"`
	Address  string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Active   bool   `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *GetListReq) Reset() {
	*x = GetListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_v1_lib_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListReq) ProtoMessage() {}

func (x *GetListReq) ProtoReflect() protoreflect.Message {
	mi := &file_lib_v1_lib_proto_msgTypes[4]
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
	return file_lib_v1_lib_proto_rawDescGZIP(), []int{4}
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

func (x *GetListReq) GetLibName() string {
	if x != nil {
		return x.LibName
	}
	return ""
}

func (x *GetListReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetListReq) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type GetListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Libs []*pbentity.Lib `protobuf:"bytes,1,rep,name=libs,proto3" json:"libs,omitempty"`
}

func (x *GetListRes) Reset() {
	*x = GetListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_v1_lib_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRes) ProtoMessage() {}

func (x *GetListRes) ProtoReflect() protoreflect.Message {
	mi := &file_lib_v1_lib_proto_msgTypes[5]
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
	return file_lib_v1_lib_proto_rawDescGZIP(), []int{5}
}

func (x *GetListRes) GetLibs() []*pbentity.Lib {
	if x != nil {
		return x.Libs
	}
	return nil
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" v:"required"`          // v:required
	LibName string `protobuf:"bytes,2,opt,name=libName,proto3" json:"libName,omitempty" v:"required"` // v:required
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty" v:"required"` // v:required
	Active  bool   `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty" v:"required"`  // v:required
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_v1_lib_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_lib_v1_lib_proto_msgTypes[6]
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
	return file_lib_v1_lib_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateReq) GetLibName() string {
	if x != nil {
		return x.LibName
	}
	return ""
}

func (x *UpdateReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateReq) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type UpdateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateRes) Reset() {
	*x = UpdateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_v1_lib_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRes) ProtoMessage() {}

func (x *UpdateRes) ProtoReflect() protoreflect.Message {
	mi := &file_lib_v1_lib_proto_msgTypes[7]
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
	return file_lib_v1_lib_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateRes) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
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
		mi := &file_lib_v1_lib_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_lib_v1_lib_proto_msgTypes[8]
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
	return file_lib_v1_lib_proto_rawDescGZIP(), []int{8}
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

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRes) Reset() {
	*x = DeleteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_v1_lib_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRes) ProtoMessage() {}

func (x *DeleteRes) ProtoReflect() protoreflect.Message {
	mi := &file_lib_v1_lib_proto_msgTypes[9]
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
	return file_lib_v1_lib_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteRes) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_lib_v1_lib_proto protoreflect.FileDescriptor

var file_lib_v1_lib_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6c, 0x69, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x6c, 0x69, 0x62, 0x1a, 0x12, 0x70, 0x62, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x6c, 0x69, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x62, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x22, 0x1b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x1b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x03, 0x6c,
	0x69, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x62, 0x52, 0x03, 0x6c, 0x69, 0x62, 0x22, 0x88, 0x01, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c,
	0x69, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69,
	0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x2f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x6c, 0x69, 0x62, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x4c,
	0x69, 0x62, 0x52, 0x04, 0x6c, 0x69, 0x62, 0x73, 0x22, 0x67, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x62, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x22, 0x1b, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b,
	0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x09, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0xe4, 0x01, 0x0a, 0x03, 0x4c, 0x69, 0x62,
	0x12, 0x2a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x6c, 0x69, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x6c, 0x69, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x0e, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x0e, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x0e, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e,
	0x6c, 0x69, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e,
	0x6c, 0x69, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42,
	0x17, 0x5a, 0x15, 0x6c, 0x69, 0x62, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lib_v1_lib_proto_rawDescOnce sync.Once
	file_lib_v1_lib_proto_rawDescData = file_lib_v1_lib_proto_rawDesc
)

func file_lib_v1_lib_proto_rawDescGZIP() []byte {
	file_lib_v1_lib_proto_rawDescOnce.Do(func() {
		file_lib_v1_lib_proto_rawDescData = protoimpl.X.CompressGZIP(file_lib_v1_lib_proto_rawDescData)
	})
	return file_lib_v1_lib_proto_rawDescData
}

var file_lib_v1_lib_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_lib_v1_lib_proto_goTypes = []interface{}{
	(*CreateReq)(nil),    // 0: lib.CreateReq
	(*CreateRes)(nil),    // 1: lib.CreateRes
	(*GetOneReq)(nil),    // 2: lib.GetOneReq
	(*GetOneRes)(nil),    // 3: lib.GetOneRes
	(*GetListReq)(nil),   // 4: lib.GetListReq
	(*GetListRes)(nil),   // 5: lib.GetListRes
	(*UpdateReq)(nil),    // 6: lib.UpdateReq
	(*UpdateRes)(nil),    // 7: lib.UpdateRes
	(*DeleteReq)(nil),    // 8: lib.DeleteReq
	(*DeleteRes)(nil),    // 9: lib.DeleteRes
	(*pbentity.Lib)(nil), // 10: pbentity.Lib
}
var file_lib_v1_lib_proto_depIdxs = []int32{
	10, // 0: lib.GetOneRes.lib:type_name -> pbentity.Lib
	10, // 1: lib.GetListRes.libs:type_name -> pbentity.Lib
	0,  // 2: lib.Lib.Create:input_type -> lib.CreateReq
	2,  // 3: lib.Lib.GetOne:input_type -> lib.GetOneReq
	4,  // 4: lib.Lib.GetList:input_type -> lib.GetListReq
	6,  // 5: lib.Lib.Update:input_type -> lib.UpdateReq
	8,  // 6: lib.Lib.Delete:input_type -> lib.DeleteReq
	1,  // 7: lib.Lib.Create:output_type -> lib.CreateRes
	3,  // 8: lib.Lib.GetOne:output_type -> lib.GetOneRes
	5,  // 9: lib.Lib.GetList:output_type -> lib.GetListRes
	7,  // 10: lib.Lib.Update:output_type -> lib.UpdateRes
	9,  // 11: lib.Lib.Delete:output_type -> lib.DeleteRes
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_lib_v1_lib_proto_init() }
func file_lib_v1_lib_proto_init() {
	if File_lib_v1_lib_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lib_v1_lib_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_lib_v1_lib_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_lib_v1_lib_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_lib_v1_lib_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_lib_v1_lib_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_lib_v1_lib_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_lib_v1_lib_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_lib_v1_lib_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_lib_v1_lib_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_lib_v1_lib_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_lib_v1_lib_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lib_v1_lib_proto_goTypes,
		DependencyIndexes: file_lib_v1_lib_proto_depIdxs,
		MessageInfos:      file_lib_v1_lib_proto_msgTypes,
	}.Build()
	File_lib_v1_lib_proto = out.File
	file_lib_v1_lib_proto_rawDesc = nil
	file_lib_v1_lib_proto_goTypes = nil
	file_lib_v1_lib_proto_depIdxs = nil
}
