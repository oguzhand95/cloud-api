// Copyright 2021-2024 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: cerbos/cloud/apikey/v1/apikey.proto

package apikeyv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/visibility"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IssueAccessTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
}

func (x *IssueAccessTokenRequest) Reset() {
	*x = IssueAccessTokenRequest{}
	mi := &file_cerbos_cloud_apikey_v1_apikey_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueAccessTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueAccessTokenRequest) ProtoMessage() {}

func (x *IssueAccessTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_apikey_v1_apikey_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueAccessTokenRequest.ProtoReflect.Descriptor instead.
func (*IssueAccessTokenRequest) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_apikey_v1_apikey_proto_rawDescGZIP(), []int{0}
}

func (x *IssueAccessTokenRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *IssueAccessTokenRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

type IssueAccessTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string               `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ExpiresIn   *durationpb.Duration `protobuf:"bytes,2,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
}

func (x *IssueAccessTokenResponse) Reset() {
	*x = IssueAccessTokenResponse{}
	mi := &file_cerbos_cloud_apikey_v1_apikey_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueAccessTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueAccessTokenResponse) ProtoMessage() {}

func (x *IssueAccessTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_apikey_v1_apikey_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueAccessTokenResponse.ProtoReflect.Descriptor instead.
func (*IssueAccessTokenResponse) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_apikey_v1_apikey_proto_rawDescGZIP(), []int{1}
}

func (x *IssueAccessTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *IssueAccessTokenResponse) GetExpiresIn() *durationpb.Duration {
	if x != nil {
		return x.ExpiresIn
	}
	return nil
}

var File_cerbos_cloud_apikey_v1_apikey_proto protoreflect.FileDescriptor

var file_cerbos_cloud_apikey_v1_apikey_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x6b, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x17, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x98, 0x01, 0x0c, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x0d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x88, 0x01, 0x0a, 0x18, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x40, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x49, 0x6e, 0x32, 0x9e, 0x01, 0x0a, 0x0d, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x10, 0x49, 0x73, 0x73, 0x75, 0x65, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2f, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x1a, 0x14, 0xfa,
	0xd2, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e,
	0x54, 0x41, 0x4c, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x62, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x3b,
	0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cerbos_cloud_apikey_v1_apikey_proto_rawDescOnce sync.Once
	file_cerbos_cloud_apikey_v1_apikey_proto_rawDescData = file_cerbos_cloud_apikey_v1_apikey_proto_rawDesc
)

func file_cerbos_cloud_apikey_v1_apikey_proto_rawDescGZIP() []byte {
	file_cerbos_cloud_apikey_v1_apikey_proto_rawDescOnce.Do(func() {
		file_cerbos_cloud_apikey_v1_apikey_proto_rawDescData = protoimpl.X.CompressGZIP(file_cerbos_cloud_apikey_v1_apikey_proto_rawDescData)
	})
	return file_cerbos_cloud_apikey_v1_apikey_proto_rawDescData
}

var file_cerbos_cloud_apikey_v1_apikey_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cerbos_cloud_apikey_v1_apikey_proto_goTypes = []any{
	(*IssueAccessTokenRequest)(nil),  // 0: cerbos.cloud.apikey.v1.IssueAccessTokenRequest
	(*IssueAccessTokenResponse)(nil), // 1: cerbos.cloud.apikey.v1.IssueAccessTokenResponse
	(*durationpb.Duration)(nil),      // 2: google.protobuf.Duration
}
var file_cerbos_cloud_apikey_v1_apikey_proto_depIdxs = []int32{
	2, // 0: cerbos.cloud.apikey.v1.IssueAccessTokenResponse.expires_in:type_name -> google.protobuf.Duration
	0, // 1: cerbos.cloud.apikey.v1.ApiKeyService.IssueAccessToken:input_type -> cerbos.cloud.apikey.v1.IssueAccessTokenRequest
	1, // 2: cerbos.cloud.apikey.v1.ApiKeyService.IssueAccessToken:output_type -> cerbos.cloud.apikey.v1.IssueAccessTokenResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cerbos_cloud_apikey_v1_apikey_proto_init() }
func file_cerbos_cloud_apikey_v1_apikey_proto_init() {
	if File_cerbos_cloud_apikey_v1_apikey_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cerbos_cloud_apikey_v1_apikey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cerbos_cloud_apikey_v1_apikey_proto_goTypes,
		DependencyIndexes: file_cerbos_cloud_apikey_v1_apikey_proto_depIdxs,
		MessageInfos:      file_cerbos_cloud_apikey_v1_apikey_proto_msgTypes,
	}.Build()
	File_cerbos_cloud_apikey_v1_apikey_proto = out.File
	file_cerbos_cloud_apikey_v1_apikey_proto_rawDesc = nil
	file_cerbos_cloud_apikey_v1_apikey_proto_goTypes = nil
	file_cerbos_cloud_apikey_v1_apikey_proto_depIdxs = nil
}
