// Copyright 2021-2025 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: cerbos/cloud/epdp/v1/epdp.proto

package epdpv1

import (
	v1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Metadata struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Version string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Deprecated: Marked as deprecated in cerbos/cloud/epdp/v1/epdp.proto.
	Policies         []string                        `protobuf:"bytes,2,rep,name=policies,proto3" json:"policies,omitempty"`
	BuildTimestamp   int64                           `protobuf:"varint,3,opt,name=build_timestamp,json=buildTimestamp,proto3" json:"build_timestamp,omitempty"`
	CommitHash       string                          `protobuf:"bytes,4,opt,name=commit_hash,json=commitHash,proto3" json:"commit_hash,omitempty"`
	SourceAttributes map[string]*v1.SourceAttributes `protobuf:"bytes,5,rep,name=source_attributes,json=sourceAttributes,proto3" json:"source_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	mi := &file_cerbos_cloud_epdp_v1_epdp_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_epdp_v1_epdp_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_epdp_v1_epdp_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// Deprecated: Marked as deprecated in cerbos/cloud/epdp/v1/epdp.proto.
func (x *Metadata) GetPolicies() []string {
	if x != nil {
		return x.Policies
	}
	return nil
}

func (x *Metadata) GetBuildTimestamp() int64 {
	if x != nil {
		return x.BuildTimestamp
	}
	return 0
}

func (x *Metadata) GetCommitHash() string {
	if x != nil {
		return x.CommitHash
	}
	return ""
}

func (x *Metadata) GetSourceAttributes() map[string]*v1.SourceAttributes {
	if x != nil {
		return x.SourceAttributes
	}
	return nil
}

var File_cerbos_cloud_epdp_v1_epdp_proto protoreflect.FileDescriptor

const file_cerbos_cloud_epdp_v1_epdp_proto_rawDesc = "" +
	"\n" +
	"\x1fcerbos/cloud/epdp/v1/epdp.proto\x12\x14cerbos.cloud.epdp.v1\x1a\x1dcerbos/policy/v1/policy.proto\"\xda\x02\n" +
	"\bMetadata\x12\x18\n" +
	"\aversion\x18\x01 \x01(\tR\aversion\x12\x1e\n" +
	"\bpolicies\x18\x02 \x03(\tB\x02\x18\x01R\bpolicies\x12'\n" +
	"\x0fbuild_timestamp\x18\x03 \x01(\x03R\x0ebuildTimestamp\x12\x1f\n" +
	"\vcommit_hash\x18\x04 \x01(\tR\n" +
	"commitHash\x12a\n" +
	"\x11source_attributes\x18\x05 \x03(\v24.cerbos.cloud.epdp.v1.Metadata.SourceAttributesEntryR\x10sourceAttributes\x1ag\n" +
	"\x15SourceAttributesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x128\n" +
	"\x05value\x18\x02 \x01(\v2\".cerbos.policy.v1.SourceAttributesR\x05value:\x028\x01B?Z=github.com/cerbos/cloud-api/genpb/cerbos/cloud/epdp/v1;epdpv1b\x06proto3"

var (
	file_cerbos_cloud_epdp_v1_epdp_proto_rawDescOnce sync.Once
	file_cerbos_cloud_epdp_v1_epdp_proto_rawDescData []byte
)

func file_cerbos_cloud_epdp_v1_epdp_proto_rawDescGZIP() []byte {
	file_cerbos_cloud_epdp_v1_epdp_proto_rawDescOnce.Do(func() {
		file_cerbos_cloud_epdp_v1_epdp_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cerbos_cloud_epdp_v1_epdp_proto_rawDesc), len(file_cerbos_cloud_epdp_v1_epdp_proto_rawDesc)))
	})
	return file_cerbos_cloud_epdp_v1_epdp_proto_rawDescData
}

var file_cerbos_cloud_epdp_v1_epdp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cerbos_cloud_epdp_v1_epdp_proto_goTypes = []any{
	(*Metadata)(nil),            // 0: cerbos.cloud.epdp.v1.Metadata
	nil,                         // 1: cerbos.cloud.epdp.v1.Metadata.SourceAttributesEntry
	(*v1.SourceAttributes)(nil), // 2: cerbos.policy.v1.SourceAttributes
}
var file_cerbos_cloud_epdp_v1_epdp_proto_depIdxs = []int32{
	1, // 0: cerbos.cloud.epdp.v1.Metadata.source_attributes:type_name -> cerbos.cloud.epdp.v1.Metadata.SourceAttributesEntry
	2, // 1: cerbos.cloud.epdp.v1.Metadata.SourceAttributesEntry.value:type_name -> cerbos.policy.v1.SourceAttributes
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cerbos_cloud_epdp_v1_epdp_proto_init() }
func file_cerbos_cloud_epdp_v1_epdp_proto_init() {
	if File_cerbos_cloud_epdp_v1_epdp_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cerbos_cloud_epdp_v1_epdp_proto_rawDesc), len(file_cerbos_cloud_epdp_v1_epdp_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cerbos_cloud_epdp_v1_epdp_proto_goTypes,
		DependencyIndexes: file_cerbos_cloud_epdp_v1_epdp_proto_depIdxs,
		MessageInfos:      file_cerbos_cloud_epdp_v1_epdp_proto_msgTypes,
	}.Build()
	File_cerbos_cloud_epdp_v1_epdp_proto = out.File
	file_cerbos_cloud_epdp_v1_epdp_proto_goTypes = nil
	file_cerbos_cloud_epdp_v1_epdp_proto_depIdxs = nil
}
