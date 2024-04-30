// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: pbentity/policy_prepare.proto

package pbentity

import (
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

type PolicyPrepare struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`              //
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`           //
	Info string `protobuf:"bytes,3,opt,name=Info,proto3" json:"Info,omitempty" dc:"策略信息"` // 策略信息
}

func (x *PolicyPrepare) Reset() {
	*x = PolicyPrepare{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbentity_policy_prepare_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyPrepare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyPrepare) ProtoMessage() {}

func (x *PolicyPrepare) ProtoReflect() protoreflect.Message {
	mi := &file_pbentity_policy_prepare_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyPrepare.ProtoReflect.Descriptor instead.
func (*PolicyPrepare) Descriptor() ([]byte, []int) {
	return file_pbentity_policy_prepare_proto_rawDescGZIP(), []int{0}
}

func (x *PolicyPrepare) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PolicyPrepare) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PolicyPrepare) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_pbentity_policy_prepare_proto protoreflect.FileDescriptor

var file_pbentity_policy_prepare_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x62, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x5f, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x70, 0x62, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x47, 0x0a, 0x0d, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x6c, 0x64, 0x6d, 0x65, 0x2d, 0x67, 0x69, 0x74, 0x2f, 0x33, 0x36, 0x68, 0x6f, 0x75,
	0x72, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x62, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbentity_policy_prepare_proto_rawDescOnce sync.Once
	file_pbentity_policy_prepare_proto_rawDescData = file_pbentity_policy_prepare_proto_rawDesc
)

func file_pbentity_policy_prepare_proto_rawDescGZIP() []byte {
	file_pbentity_policy_prepare_proto_rawDescOnce.Do(func() {
		file_pbentity_policy_prepare_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbentity_policy_prepare_proto_rawDescData)
	})
	return file_pbentity_policy_prepare_proto_rawDescData
}

var file_pbentity_policy_prepare_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pbentity_policy_prepare_proto_goTypes = []interface{}{
	(*PolicyPrepare)(nil), // 0: pbentity.PolicyPrepare
}
var file_pbentity_policy_prepare_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pbentity_policy_prepare_proto_init() }
func file_pbentity_policy_prepare_proto_init() {
	if File_pbentity_policy_prepare_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbentity_policy_prepare_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyPrepare); i {
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
			RawDescriptor: file_pbentity_policy_prepare_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbentity_policy_prepare_proto_goTypes,
		DependencyIndexes: file_pbentity_policy_prepare_proto_depIdxs,
		MessageInfos:      file_pbentity_policy_prepare_proto_msgTypes,
	}.Build()
	File_pbentity_policy_prepare_proto = out.File
	file_pbentity_policy_prepare_proto_rawDesc = nil
	file_pbentity_policy_prepare_proto_goTypes = nil
	file_pbentity_policy_prepare_proto_depIdxs = nil
}