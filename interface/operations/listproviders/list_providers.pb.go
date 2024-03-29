//
// Copyright 2019 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: list_providers.proto

package listproviders

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

type ProviderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Vendor      string `protobuf:"bytes,3,opt,name=vendor,proto3" json:"vendor,omitempty"`
	VersionMaj  uint32 `protobuf:"varint,4,opt,name=version_maj,json=versionMaj,proto3" json:"version_maj,omitempty"`
	VersionMin  uint32 `protobuf:"varint,5,opt,name=version_min,json=versionMin,proto3" json:"version_min,omitempty"`
	VersionRev  uint32 `protobuf:"varint,6,opt,name=version_rev,json=versionRev,proto3" json:"version_rev,omitempty"`
	Id          uint32 `protobuf:"varint,7,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProviderInfo) Reset() {
	*x = ProviderInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_providers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderInfo) ProtoMessage() {}

func (x *ProviderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_list_providers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderInfo.ProtoReflect.Descriptor instead.
func (*ProviderInfo) Descriptor() ([]byte, []int) {
	return file_list_providers_proto_rawDescGZIP(), []int{0}
}

func (x *ProviderInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ProviderInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProviderInfo) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *ProviderInfo) GetVersionMaj() uint32 {
	if x != nil {
		return x.VersionMaj
	}
	return 0
}

func (x *ProviderInfo) GetVersionMin() uint32 {
	if x != nil {
		return x.VersionMin
	}
	return 0
}

func (x *ProviderInfo) GetVersionRev() uint32 {
	if x != nil {
		return x.VersionRev
	}
	return 0
}

func (x *ProviderInfo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_providers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_list_providers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_list_providers_proto_rawDescGZIP(), []int{1}
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Providers []*ProviderInfo `protobuf:"bytes,1,rep,name=providers,proto3" json:"providers,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_providers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_list_providers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_list_providers_proto_rawDescGZIP(), []int{2}
}

func (x *Result) GetProviders() []*ProviderInfo {
	if x != nil {
		return x.Providers
	}
	return nil
}

var File_list_providers_proto protoreflect.FileDescriptor

var file_list_providers_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x22, 0xcf, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x61, 0x6a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6a, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x76, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0b, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x44, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x3a, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x42, 0x4f, 0x5a, 0x4d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c,
	0x61, 0x78, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x63, 0x2d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_list_providers_proto_rawDescOnce sync.Once
	file_list_providers_proto_rawDescData = file_list_providers_proto_rawDesc
)

func file_list_providers_proto_rawDescGZIP() []byte {
	file_list_providers_proto_rawDescOnce.Do(func() {
		file_list_providers_proto_rawDescData = protoimpl.X.CompressGZIP(file_list_providers_proto_rawDescData)
	})
	return file_list_providers_proto_rawDescData
}

var file_list_providers_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_list_providers_proto_goTypes = []interface{}{
	(*ProviderInfo)(nil), // 0: list_providers.ProviderInfo
	(*Operation)(nil),    // 1: list_providers.Operation
	(*Result)(nil),       // 2: list_providers.Result
}
var file_list_providers_proto_depIdxs = []int32{
	0, // 0: list_providers.Result.providers:type_name -> list_providers.ProviderInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_list_providers_proto_init() }
func file_list_providers_proto_init() {
	if File_list_providers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_list_providers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderInfo); i {
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
		file_list_providers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
		file_list_providers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_list_providers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_list_providers_proto_goTypes,
		DependencyIndexes: file_list_providers_proto_depIdxs,
		MessageInfos:      file_list_providers_proto_msgTypes,
	}.Build()
	File_list_providers_proto = out.File
	file_list_providers_proto_rawDesc = nil
	file_list_providers_proto_goTypes = nil
	file_list_providers_proto_depIdxs = nil
}
