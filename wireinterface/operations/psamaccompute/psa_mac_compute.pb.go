//
// Copyright 2020 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: psa_mac_compute.proto

package psamaccompute

import (
	psaalgorithm "github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psaalgorithm"
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

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyName string                      `protobuf:"bytes,1,opt,name=key_name,json=keyName,proto3" json:"key_name,omitempty"`
	Alg     *psaalgorithm.Algorithm_Mac `protobuf:"bytes,2,opt,name=alg,proto3" json:"alg,omitempty"`
	Input   []byte                      `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_psa_mac_compute_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_psa_mac_compute_proto_msgTypes[0]
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
	return file_psa_mac_compute_proto_rawDescGZIP(), []int{0}
}

func (x *Operation) GetKeyName() string {
	if x != nil {
		return x.KeyName
	}
	return ""
}

func (x *Operation) GetAlg() *psaalgorithm.Algorithm_Mac {
	if x != nil {
		return x.Alg
	}
	return nil
}

func (x *Operation) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mac []byte `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_psa_mac_compute_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_psa_mac_compute_proto_msgTypes[1]
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
	return file_psa_mac_compute_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetMac() []byte {
	if x != nil {
		return x.Mac
	}
	return nil
}

var File_psa_mac_compute_proto protoreflect.FileDescriptor

var file_psa_mac_compute_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x73, 0x61, 0x5f, 0x6d, 0x61, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x73, 0x61, 0x5f, 0x6d, 0x61, 0x63,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x1a, 0x13, 0x70, 0x73, 0x61, 0x5f, 0x61, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a,
	0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x03, 0x61, 0x6c, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x73, 0x61, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x2e, 0x4d, 0x61, 0x63,
	0x52, 0x03, 0x61, 0x6c, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x1a, 0x0a, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x42, 0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x61, 0x78, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x77, 0x69, 0x72, 0x65, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70,
	0x73, 0x61, 0x6d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_psa_mac_compute_proto_rawDescOnce sync.Once
	file_psa_mac_compute_proto_rawDescData = file_psa_mac_compute_proto_rawDesc
)

func file_psa_mac_compute_proto_rawDescGZIP() []byte {
	file_psa_mac_compute_proto_rawDescOnce.Do(func() {
		file_psa_mac_compute_proto_rawDescData = protoimpl.X.CompressGZIP(file_psa_mac_compute_proto_rawDescData)
	})
	return file_psa_mac_compute_proto_rawDescData
}

var file_psa_mac_compute_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_psa_mac_compute_proto_goTypes = []interface{}{
	(*Operation)(nil),                  // 0: psa_mac_compute.Operation
	(*Result)(nil),                     // 1: psa_mac_compute.Result
	(*psaalgorithm.Algorithm_Mac)(nil), // 2: psa_algorithm.Algorithm.Mac
}
var file_psa_mac_compute_proto_depIdxs = []int32{
	2, // 0: psa_mac_compute.Operation.alg:type_name -> psa_algorithm.Algorithm.Mac
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_psa_mac_compute_proto_init() }
func file_psa_mac_compute_proto_init() {
	if File_psa_mac_compute_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_psa_mac_compute_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_psa_mac_compute_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_psa_mac_compute_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_psa_mac_compute_proto_goTypes,
		DependencyIndexes: file_psa_mac_compute_proto_depIdxs,
		MessageInfos:      file_psa_mac_compute_proto_msgTypes,
	}.Build()
	File_psa_mac_compute_proto = out.File
	file_psa_mac_compute_proto_rawDesc = nil
	file_psa_mac_compute_proto_goTypes = nil
	file_psa_mac_compute_proto_depIdxs = nil
}