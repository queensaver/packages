//
//Queensaver API
//
//Queensaver API to send in sensor data and retrieve it. It's also used for user management.
//
//The version of the OpenAPI document: 0.0.1
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: models/hive_frame.proto

package models

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

// The format of the frame
type HiveFrame_FunctionEnum int32

const (
	HiveFrame_BREED       HiveFrame_FunctionEnum = 0
	HiveFrame_HONEY       HiveFrame_FunctionEnum = 1
	HiveFrame_UNSPECIFIED HiveFrame_FunctionEnum = 2
)

// Enum value maps for HiveFrame_FunctionEnum.
var (
	HiveFrame_FunctionEnum_name = map[int32]string{
		0: "BREED",
		1: "HONEY",
		2: "UNSPECIFIED",
	}
	HiveFrame_FunctionEnum_value = map[string]int32{
		"BREED":       0,
		"HONEY":       1,
		"UNSPECIFIED": 2,
	}
)

func (x HiveFrame_FunctionEnum) Enum() *HiveFrame_FunctionEnum {
	p := new(HiveFrame_FunctionEnum)
	*p = x
	return p
}

func (x HiveFrame_FunctionEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HiveFrame_FunctionEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_models_hive_frame_proto_enumTypes[0].Descriptor()
}

func (HiveFrame_FunctionEnum) Type() protoreflect.EnumType {
	return &file_models_hive_frame_proto_enumTypes[0]
}

func (x HiveFrame_FunctionEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HiveFrame_FunctionEnum.Descriptor instead.
func (HiveFrame_FunctionEnum) EnumDescriptor() ([]byte, []int) {
	return file_models_hive_frame_proto_rawDescGZIP(), []int{0, 0}
}

type HiveFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Function HiveFrame_FunctionEnum `protobuf:"varint,1,opt,name=function,proto3,enum=openapi.HiveFrame_FunctionEnum" json:"function,omitempty"`
}

func (x *HiveFrame) Reset() {
	*x = HiveFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_hive_frame_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HiveFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HiveFrame) ProtoMessage() {}

func (x *HiveFrame) ProtoReflect() protoreflect.Message {
	mi := &file_models_hive_frame_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HiveFrame.ProtoReflect.Descriptor instead.
func (*HiveFrame) Descriptor() ([]byte, []int) {
	return file_models_hive_frame_proto_rawDescGZIP(), []int{0}
}

func (x *HiveFrame) GetFunction() HiveFrame_FunctionEnum {
	if x != nil {
		return x.Function
	}
	return HiveFrame_BREED
}

var File_models_hive_frame_proto protoreflect.FileDescriptor

var file_models_hive_frame_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x22, 0x7f, 0x0a, 0x09, 0x48, 0x69, 0x76, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12,
	0x3b, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x69, 0x76, 0x65,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e,
	0x75, 0x6d, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x0c,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x09, 0x0a, 0x05,
	0x42, 0x52, 0x45, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x4f, 0x4e, 0x45, 0x59,
	0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_hive_frame_proto_rawDescOnce sync.Once
	file_models_hive_frame_proto_rawDescData = file_models_hive_frame_proto_rawDesc
)

func file_models_hive_frame_proto_rawDescGZIP() []byte {
	file_models_hive_frame_proto_rawDescOnce.Do(func() {
		file_models_hive_frame_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_hive_frame_proto_rawDescData)
	})
	return file_models_hive_frame_proto_rawDescData
}

var file_models_hive_frame_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_models_hive_frame_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_hive_frame_proto_goTypes = []interface{}{
	(HiveFrame_FunctionEnum)(0), // 0: openapi.HiveFrame.FunctionEnum
	(*HiveFrame)(nil),           // 1: openapi.HiveFrame
}
var file_models_hive_frame_proto_depIdxs = []int32{
	0, // 0: openapi.HiveFrame.function:type_name -> openapi.HiveFrame.FunctionEnum
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_hive_frame_proto_init() }
func file_models_hive_frame_proto_init() {
	if File_models_hive_frame_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_hive_frame_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HiveFrame); i {
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
			RawDescriptor: file_models_hive_frame_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_hive_frame_proto_goTypes,
		DependencyIndexes: file_models_hive_frame_proto_depIdxs,
		EnumInfos:         file_models_hive_frame_proto_enumTypes,
		MessageInfos:      file_models_hive_frame_proto_msgTypes,
	}.Build()
	File_models_hive_frame_proto = out.File
	file_models_hive_frame_proto_rawDesc = nil
	file_models_hive_frame_proto_goTypes = nil
	file_models_hive_frame_proto_depIdxs = nil
}
