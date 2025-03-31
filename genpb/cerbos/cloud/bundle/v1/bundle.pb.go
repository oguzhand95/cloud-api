// Copyright 2021-2025 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: cerbos/cloud/bundle/v1/bundle.proto

package bundlev1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v1 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/pdp/v1"
	_ "google.golang.org/genproto/googleapis/api/visibility"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// BundleInfo holds information about a bundle and its download URLs.
type BundleInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Label         string                 `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	BundleHash    []byte                 `protobuf:"bytes,2,opt,name=bundle_hash,json=bundleHash,proto3" json:"bundle_hash,omitempty"`
	Segments      []*BundleInfo_Segment  `protobuf:"bytes,3,rep,name=segments,proto3" json:"segments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BundleInfo) Reset() {
	*x = BundleInfo{}
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BundleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BundleInfo) ProtoMessage() {}

func (x *BundleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BundleInfo.ProtoReflect.Descriptor instead.
func (*BundleInfo) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_bundle_v1_bundle_proto_rawDescGZIP(), []int{0}
}

func (x *BundleInfo) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *BundleInfo) GetBundleHash() []byte {
	if x != nil {
		return x.BundleHash
	}
	return nil
}

func (x *BundleInfo) GetSegments() []*BundleInfo_Segment {
	if x != nil {
		return x.Segments
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Identifier    string                 `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Source        string                 `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Meta) Reset() {
	*x = Meta{}
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_bundle_v1_bundle_proto_rawDescGZIP(), []int{1}
}

func (x *Meta) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Meta) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type Manifest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ApiVersion    string                 `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	PolicyIndex   map[string]string      `protobuf:"bytes,2,rep,name=policy_index,json=policyIndex,proto3" json:"policy_index,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Schemas       []string               `protobuf:"bytes,3,rep,name=schemas,proto3" json:"schemas,omitempty"`
	Meta          *Meta                  `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Manifest) Reset() {
	*x = Manifest{}
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Manifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manifest) ProtoMessage() {}

func (x *Manifest) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manifest.ProtoReflect.Descriptor instead.
func (*Manifest) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_bundle_v1_bundle_proto_rawDescGZIP(), []int{2}
}

func (x *Manifest) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *Manifest) GetPolicyIndex() map[string]string {
	if x != nil {
		return x.PolicyIndex
	}
	return nil
}

func (x *Manifest) GetSchemas() []string {
	if x != nil {
		return x.Schemas
	}
	return nil
}

func (x *Manifest) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type GetBundleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PdpId         *v1.Identifier         `protobuf:"bytes,1,opt,name=pdp_id,json=pdpId,proto3" json:"pdp_id,omitempty"`
	BundleLabel   string                 `protobuf:"bytes,2,opt,name=bundle_label,json=bundleLabel,proto3" json:"bundle_label,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBundleRequest) Reset() {
	*x = GetBundleRequest{}
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBundleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBundleRequest) ProtoMessage() {}

func (x *GetBundleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBundleRequest.ProtoReflect.Descriptor instead.
func (*GetBundleRequest) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_bundle_v1_bundle_proto_rawDescGZIP(), []int{3}
}

func (x *GetBundleRequest) GetPdpId() *v1.Identifier {
	if x != nil {
		return x.PdpId
	}
	return nil
}

func (x *GetBundleRequest) GetBundleLabel() string {
	if x != nil {
		return x.BundleLabel
	}
	return ""
}

type GetBundleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BundleInfo    *BundleInfo            `protobuf:"bytes,1,opt,name=bundle_info,json=bundleInfo,proto3" json:"bundle_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBundleResponse) Reset() {
	*x = GetBundleResponse{}
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBundleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBundleResponse) ProtoMessage() {}

func (x *GetBundleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBundleResponse.ProtoReflect.Descriptor instead.
func (*GetBundleResponse) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_bundle_v1_bundle_proto_rawDescGZIP(), []int{4}
}

func (x *GetBundleResponse) GetBundleInfo() *BundleInfo {
	if x != nil {
		return x.BundleInfo
	}
	return nil
}

type WatchBundleRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	PdpId *v1.Identifier         `protobuf:"bytes,1,opt,name=pdp_id,json=pdpId,proto3" json:"pdp_id,omitempty"`
	// Types that are valid to be assigned to Msg:
	//
	//	*WatchBundleRequest_WatchLabel_
	//	*WatchBundleRequest_Heartbeat_
	Msg           isWatchBundleRequest_Msg `protobuf_oneof:"msg"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WatchBundleRequest) Reset() {
	*x = WatchBundleRequest{}
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchBundleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchBundleRequest) ProtoMessage() {}

func (x *WatchBundleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchBundleRequest.ProtoReflect.Descriptor instead.
func (*WatchBundleRequest) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_bundle_v1_bundle_proto_rawDescGZIP(), []int{5}
}

func (x *WatchBundleRequest) GetPdpId() *v1.Identifier {
	if x != nil {
		return x.PdpId
	}
	return nil
}

func (x *WatchBundleRequest) GetMsg() isWatchBundleRequest_Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *WatchBundleRequest) GetWatchLabel() *WatchBundleRequest_WatchLabel {
	if x != nil {
		if x, ok := x.Msg.(*WatchBundleRequest_WatchLabel_); ok {
			return x.WatchLabel
		}
	}
	return nil
}

func (x *WatchBundleRequest) GetHeartbeat() *WatchBundleRequest_Heartbeat {
	if x != nil {
		if x, ok := x.Msg.(*WatchBundleRequest_Heartbeat_); ok {
			return x.Heartbeat
		}
	}
	return nil
}

type isWatchBundleRequest_Msg interface {
	isWatchBundleRequest_Msg()
}

type WatchBundleRequest_WatchLabel_ struct {
	WatchLabel *WatchBundleRequest_WatchLabel `protobuf:"bytes,2,opt,name=watch_label,json=watchLabel,proto3,oneof"`
}

type WatchBundleRequest_Heartbeat_ struct {
	Heartbeat *WatchBundleRequest_Heartbeat `protobuf:"bytes,3,opt,name=heartbeat,proto3,oneof"`
}

func (*WatchBundleRequest_WatchLabel_) isWatchBundleRequest_Msg() {}

func (*WatchBundleRequest_Heartbeat_) isWatchBundleRequest_Msg() {}

type WatchBundleResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Msg:
	//
	//	*WatchBundleResponse_BundleUpdate
	//	*WatchBundleResponse_Reconnect_
	//	*WatchBundleResponse_BundleRemoved_
	Msg           isWatchBundleResponse_Msg `protobuf_oneof:"msg"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WatchBundleResponse) Reset() {
	*x = WatchBundleResponse{}
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchBundleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchBundleResponse) ProtoMessage() {}

func (x *WatchBundleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchBundleResponse.ProtoReflect.Descriptor instead.
func (*WatchBundleResponse) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_bundle_v1_bundle_proto_rawDescGZIP(), []int{6}
}

func (x *WatchBundleResponse) GetMsg() isWatchBundleResponse_Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *WatchBundleResponse) GetBundleUpdate() *BundleInfo {
	if x != nil {
		if x, ok := x.Msg.(*WatchBundleResponse_BundleUpdate); ok {
			return x.BundleUpdate
		}
	}
	return nil
}

func (x *WatchBundleResponse) GetReconnect() *WatchBundleResponse_Reconnect {
	if x != nil {
		if x, ok := x.Msg.(*WatchBundleResponse_Reconnect_); ok {
			return x.Reconnect
		}
	}
	return nil
}

func (x *WatchBundleResponse) GetBundleRemoved() *WatchBundleResponse_BundleRemoved {
	if x != nil {
		if x, ok := x.Msg.(*WatchBundleResponse_BundleRemoved_); ok {
			return x.BundleRemoved
		}
	}
	return nil
}

type isWatchBundleResponse_Msg interface {
	isWatchBundleResponse_Msg()
}

type WatchBundleResponse_BundleUpdate struct {
	BundleUpdate *BundleInfo `protobuf:"bytes,1,opt,name=bundle_update,json=bundleUpdate,proto3,oneof"`
}

type WatchBundleResponse_Reconnect_ struct {
	Reconnect *WatchBundleResponse_Reconnect `protobuf:"bytes,2,opt,name=reconnect,proto3,oneof"`
}

type WatchBundleResponse_BundleRemoved_ struct {
	BundleRemoved *WatchBundleResponse_BundleRemoved `protobuf:"bytes,3,opt,name=bundle_removed,json=bundleRemoved,proto3,oneof"`
}

func (*WatchBundleResponse_BundleUpdate) isWatchBundleResponse_Msg() {}

func (*WatchBundleResponse_Reconnect_) isWatchBundleResponse_Msg() {}

func (*WatchBundleResponse_BundleRemoved_) isWatchBundleResponse_Msg() {}

type BundleInfo_Segment struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SegmentId     uint32                 `protobuf:"varint,1,opt,name=segment_id,json=segmentId,proto3" json:"segment_id,omitempty"`
	Checksum      []byte                 `protobuf:"bytes,2,opt,name=checksum,proto3" json:"checksum,omitempty"`
	DownloadUrls  []string               `protobuf:"bytes,3,rep,name=download_urls,json=downloadUrls,proto3" json:"download_urls,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BundleInfo_Segment) Reset() {
	*x = BundleInfo_Segment{}
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BundleInfo_Segment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BundleInfo_Segment) ProtoMessage() {}

func (x *BundleInfo_Segment) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BundleInfo_Segment.ProtoReflect.Descriptor instead.
func (*BundleInfo_Segment) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_bundle_v1_bundle_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BundleInfo_Segment) GetSegmentId() uint32 {
	if x != nil {
		return x.SegmentId
	}
	return 0
}

func (x *BundleInfo_Segment) GetChecksum() []byte {
	if x != nil {
		return x.Checksum
	}
	return nil
}

func (x *BundleInfo_Segment) GetDownloadUrls() []string {
	if x != nil {
		return x.DownloadUrls
	}
	return nil
}

type WatchBundleRequest_WatchLabel struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BundleLabel   string                 `protobuf:"bytes,1,opt,name=bundle_label,json=bundleLabel,proto3" json:"bundle_label,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WatchBundleRequest_WatchLabel) Reset() {
	*x = WatchBundleRequest_WatchLabel{}
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchBundleRequest_WatchLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchBundleRequest_WatchLabel) ProtoMessage() {}

func (x *WatchBundleRequest_WatchLabel) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchBundleRequest_WatchLabel.ProtoReflect.Descriptor instead.
func (*WatchBundleRequest_WatchLabel) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_bundle_v1_bundle_proto_rawDescGZIP(), []int{5, 0}
}

func (x *WatchBundleRequest_WatchLabel) GetBundleLabel() string {
	if x != nil {
		return x.BundleLabel
	}
	return ""
}

type WatchBundleRequest_Heartbeat struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Timestamp      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ActiveBundleId string                 `protobuf:"bytes,2,opt,name=active_bundle_id,json=activeBundleId,proto3" json:"active_bundle_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *WatchBundleRequest_Heartbeat) Reset() {
	*x = WatchBundleRequest_Heartbeat{}
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchBundleRequest_Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchBundleRequest_Heartbeat) ProtoMessage() {}

func (x *WatchBundleRequest_Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchBundleRequest_Heartbeat.ProtoReflect.Descriptor instead.
func (*WatchBundleRequest_Heartbeat) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_bundle_v1_bundle_proto_rawDescGZIP(), []int{5, 1}
}

func (x *WatchBundleRequest_Heartbeat) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *WatchBundleRequest_Heartbeat) GetActiveBundleId() string {
	if x != nil {
		return x.ActiveBundleId
	}
	return ""
}

type WatchBundleResponse_Reconnect struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Backoff       *durationpb.Duration   `protobuf:"bytes,1,opt,name=backoff,proto3" json:"backoff,omitempty"`
	Reason        string                 `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WatchBundleResponse_Reconnect) Reset() {
	*x = WatchBundleResponse_Reconnect{}
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchBundleResponse_Reconnect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchBundleResponse_Reconnect) ProtoMessage() {}

func (x *WatchBundleResponse_Reconnect) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchBundleResponse_Reconnect.ProtoReflect.Descriptor instead.
func (*WatchBundleResponse_Reconnect) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_bundle_v1_bundle_proto_rawDescGZIP(), []int{6, 0}
}

func (x *WatchBundleResponse_Reconnect) GetBackoff() *durationpb.Duration {
	if x != nil {
		return x.Backoff
	}
	return nil
}

func (x *WatchBundleResponse_Reconnect) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type WatchBundleResponse_BundleRemoved struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WatchBundleResponse_BundleRemoved) Reset() {
	*x = WatchBundleResponse_BundleRemoved{}
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchBundleResponse_BundleRemoved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchBundleResponse_BundleRemoved) ProtoMessage() {}

func (x *WatchBundleResponse_BundleRemoved) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchBundleResponse_BundleRemoved.ProtoReflect.Descriptor instead.
func (*WatchBundleResponse_BundleRemoved) Descriptor() ([]byte, []int) {
	return file_cerbos_cloud_bundle_v1_bundle_proto_rawDescGZIP(), []int{6, 1}
}

var File_cerbos_cloud_bundle_v1_bundle_proto protoreflect.FileDescriptor

const file_cerbos_cloud_bundle_v1_bundle_proto_rawDesc = "" +
	"\n" +
	"#cerbos/cloud/bundle/v1/bundle.proto\x12\x16cerbos.cloud.bundle.v1\x1a\x1bbuf/validate/validate.proto\x1a\x1dcerbos/cloud/pdp/v1/pdp.proto\x1a\x1bgoogle/api/visibility.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xb7\x02\n" +
	"\n" +
	"BundleInfo\x12\x1d\n" +
	"\x05label\x18\x01 \x01(\tB\a\xbaH\x04r\x02\x10\x01R\x05label\x12(\n" +
	"\vbundle_hash\x18\x02 \x01(\fB\a\xbaH\x04z\x02h R\n" +
	"bundleHash\x12P\n" +
	"\bsegments\x18\x03 \x03(\v2*.cerbos.cloud.bundle.v1.BundleInfo.SegmentB\b\xbaH\x05\x92\x01\x02\b\x01R\bsegments\x1a\x8d\x01\n" +
	"\aSegment\x12&\n" +
	"\n" +
	"segment_id\x18\x01 \x01(\rB\a\xbaH\x04*\x02 \x00R\tsegmentId\x12#\n" +
	"\bchecksum\x18\x02 \x01(\fB\a\xbaH\x04z\x02h R\bchecksum\x125\n" +
	"\rdownload_urls\x18\x03 \x03(\tB\x10\xbaH\r\x92\x01\n" +
	"\b\x01\x18\x01\"\x04r\x02\x10\x01R\fdownloadUrls\">\n" +
	"\x04Meta\x12\x1e\n" +
	"\n" +
	"identifier\x18\x01 \x01(\tR\n" +
	"identifier\x12\x16\n" +
	"\x06source\x18\x02 \x01(\tR\x06source\"\x8d\x02\n" +
	"\bManifest\x12\x1f\n" +
	"\vapi_version\x18\x01 \x01(\tR\n" +
	"apiVersion\x12T\n" +
	"\fpolicy_index\x18\x02 \x03(\v21.cerbos.cloud.bundle.v1.Manifest.PolicyIndexEntryR\vpolicyIndex\x12\x18\n" +
	"\aschemas\x18\x03 \x03(\tR\aschemas\x120\n" +
	"\x04meta\x18\x04 \x01(\v2\x1c.cerbos.cloud.bundle.v1.MetaR\x04meta\x1a>\n" +
	"\x10PolicyIndexEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"~\n" +
	"\x10GetBundleRequest\x12>\n" +
	"\x06pdp_id\x18\x01 \x01(\v2\x1f.cerbos.cloud.pdp.v1.IdentifierB\x06\xbaH\x03\xc8\x01\x01R\x05pdpId\x12*\n" +
	"\fbundle_label\x18\x02 \x01(\tB\a\xbaH\x04r\x02\x10\x01R\vbundleLabel\"`\n" +
	"\x11GetBundleResponse\x12K\n" +
	"\vbundle_info\x18\x01 \x01(\v2\".cerbos.cloud.bundle.v1.BundleInfoB\x06\xbaH\x03\xc8\x01\x01R\n" +
	"bundleInfo\"\xcf\x03\n" +
	"\x12WatchBundleRequest\x12>\n" +
	"\x06pdp_id\x18\x01 \x01(\v2\x1f.cerbos.cloud.pdp.v1.IdentifierB\x06\xbaH\x03\xc8\x01\x01R\x05pdpId\x12X\n" +
	"\vwatch_label\x18\x02 \x01(\v25.cerbos.cloud.bundle.v1.WatchBundleRequest.WatchLabelH\x00R\n" +
	"watchLabel\x12T\n" +
	"\theartbeat\x18\x03 \x01(\v24.cerbos.cloud.bundle.v1.WatchBundleRequest.HeartbeatH\x00R\theartbeat\x1a8\n" +
	"\n" +
	"WatchLabel\x12*\n" +
	"\fbundle_label\x18\x01 \x01(\tB\a\xbaH\x04r\x02\x10\x01R\vbundleLabel\x1a\x87\x01\n" +
	"\tHeartbeat\x12G\n" +
	"\ttimestamp\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampB\r\xbaH\n" +
	"\xc8\x01\x01\xb2\x01\x04J\x02\b<R\ttimestamp\x121\n" +
	"\x10active_bundle_id\x18\x02 \x01(\tB\a\xbaH\x04r\x02\x10\x01R\x0eactiveBundleIdB\x05\n" +
	"\x03msg\"\x94\x03\n" +
	"\x13WatchBundleResponse\x12I\n" +
	"\rbundle_update\x18\x01 \x01(\v2\".cerbos.cloud.bundle.v1.BundleInfoH\x00R\fbundleUpdate\x12U\n" +
	"\treconnect\x18\x02 \x01(\v25.cerbos.cloud.bundle.v1.WatchBundleResponse.ReconnectH\x00R\treconnect\x12b\n" +
	"\x0ebundle_removed\x18\x03 \x01(\v29.cerbos.cloud.bundle.v1.WatchBundleResponse.BundleRemovedH\x00R\rbundleRemoved\x1aX\n" +
	"\tReconnect\x123\n" +
	"\abackoff\x18\x01 \x01(\v2\x19.google.protobuf.DurationR\abackoff\x12\x16\n" +
	"\x06reason\x18\x02 \x01(\tR\x06reason\x1a\x0f\n" +
	"\rBundleRemovedB\f\n" +
	"\x03msg\x12\x05\xbaH\x02\b\x012\xfd\x01\n" +
	"\x13CerbosBundleService\x12b\n" +
	"\tGetBundle\x12(.cerbos.cloud.bundle.v1.GetBundleRequest\x1a).cerbos.cloud.bundle.v1.GetBundleResponse\"\x00\x12l\n" +
	"\vWatchBundle\x12*.cerbos.cloud.bundle.v1.WatchBundleRequest\x1a+.cerbos.cloud.bundle.v1.WatchBundleResponse\"\x00(\x010\x01\x1a\x14\xfa\xd2\xe4\x93\x02\x0e\x12\fEXPERIMENTALBCZAgithub.com/cerbos/cloud-api/genpb/cerbos/cloud/bundle/v1;bundlev1b\x06proto3"

var (
	file_cerbos_cloud_bundle_v1_bundle_proto_rawDescOnce sync.Once
	file_cerbos_cloud_bundle_v1_bundle_proto_rawDescData []byte
)

func file_cerbos_cloud_bundle_v1_bundle_proto_rawDescGZIP() []byte {
	file_cerbos_cloud_bundle_v1_bundle_proto_rawDescOnce.Do(func() {
		file_cerbos_cloud_bundle_v1_bundle_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cerbos_cloud_bundle_v1_bundle_proto_rawDesc), len(file_cerbos_cloud_bundle_v1_bundle_proto_rawDesc)))
	})
	return file_cerbos_cloud_bundle_v1_bundle_proto_rawDescData
}

var file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_cerbos_cloud_bundle_v1_bundle_proto_goTypes = []any{
	(*BundleInfo)(nil),                        // 0: cerbos.cloud.bundle.v1.BundleInfo
	(*Meta)(nil),                              // 1: cerbos.cloud.bundle.v1.Meta
	(*Manifest)(nil),                          // 2: cerbos.cloud.bundle.v1.Manifest
	(*GetBundleRequest)(nil),                  // 3: cerbos.cloud.bundle.v1.GetBundleRequest
	(*GetBundleResponse)(nil),                 // 4: cerbos.cloud.bundle.v1.GetBundleResponse
	(*WatchBundleRequest)(nil),                // 5: cerbos.cloud.bundle.v1.WatchBundleRequest
	(*WatchBundleResponse)(nil),               // 6: cerbos.cloud.bundle.v1.WatchBundleResponse
	(*BundleInfo_Segment)(nil),                // 7: cerbos.cloud.bundle.v1.BundleInfo.Segment
	nil,                                       // 8: cerbos.cloud.bundle.v1.Manifest.PolicyIndexEntry
	(*WatchBundleRequest_WatchLabel)(nil),     // 9: cerbos.cloud.bundle.v1.WatchBundleRequest.WatchLabel
	(*WatchBundleRequest_Heartbeat)(nil),      // 10: cerbos.cloud.bundle.v1.WatchBundleRequest.Heartbeat
	(*WatchBundleResponse_Reconnect)(nil),     // 11: cerbos.cloud.bundle.v1.WatchBundleResponse.Reconnect
	(*WatchBundleResponse_BundleRemoved)(nil), // 12: cerbos.cloud.bundle.v1.WatchBundleResponse.BundleRemoved
	(*v1.Identifier)(nil),                     // 13: cerbos.cloud.pdp.v1.Identifier
	(*timestamppb.Timestamp)(nil),             // 14: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),               // 15: google.protobuf.Duration
}
var file_cerbos_cloud_bundle_v1_bundle_proto_depIdxs = []int32{
	7,  // 0: cerbos.cloud.bundle.v1.BundleInfo.segments:type_name -> cerbos.cloud.bundle.v1.BundleInfo.Segment
	8,  // 1: cerbos.cloud.bundle.v1.Manifest.policy_index:type_name -> cerbos.cloud.bundle.v1.Manifest.PolicyIndexEntry
	1,  // 2: cerbos.cloud.bundle.v1.Manifest.meta:type_name -> cerbos.cloud.bundle.v1.Meta
	13, // 3: cerbos.cloud.bundle.v1.GetBundleRequest.pdp_id:type_name -> cerbos.cloud.pdp.v1.Identifier
	0,  // 4: cerbos.cloud.bundle.v1.GetBundleResponse.bundle_info:type_name -> cerbos.cloud.bundle.v1.BundleInfo
	13, // 5: cerbos.cloud.bundle.v1.WatchBundleRequest.pdp_id:type_name -> cerbos.cloud.pdp.v1.Identifier
	9,  // 6: cerbos.cloud.bundle.v1.WatchBundleRequest.watch_label:type_name -> cerbos.cloud.bundle.v1.WatchBundleRequest.WatchLabel
	10, // 7: cerbos.cloud.bundle.v1.WatchBundleRequest.heartbeat:type_name -> cerbos.cloud.bundle.v1.WatchBundleRequest.Heartbeat
	0,  // 8: cerbos.cloud.bundle.v1.WatchBundleResponse.bundle_update:type_name -> cerbos.cloud.bundle.v1.BundleInfo
	11, // 9: cerbos.cloud.bundle.v1.WatchBundleResponse.reconnect:type_name -> cerbos.cloud.bundle.v1.WatchBundleResponse.Reconnect
	12, // 10: cerbos.cloud.bundle.v1.WatchBundleResponse.bundle_removed:type_name -> cerbos.cloud.bundle.v1.WatchBundleResponse.BundleRemoved
	14, // 11: cerbos.cloud.bundle.v1.WatchBundleRequest.Heartbeat.timestamp:type_name -> google.protobuf.Timestamp
	15, // 12: cerbos.cloud.bundle.v1.WatchBundleResponse.Reconnect.backoff:type_name -> google.protobuf.Duration
	3,  // 13: cerbos.cloud.bundle.v1.CerbosBundleService.GetBundle:input_type -> cerbos.cloud.bundle.v1.GetBundleRequest
	5,  // 14: cerbos.cloud.bundle.v1.CerbosBundleService.WatchBundle:input_type -> cerbos.cloud.bundle.v1.WatchBundleRequest
	4,  // 15: cerbos.cloud.bundle.v1.CerbosBundleService.GetBundle:output_type -> cerbos.cloud.bundle.v1.GetBundleResponse
	6,  // 16: cerbos.cloud.bundle.v1.CerbosBundleService.WatchBundle:output_type -> cerbos.cloud.bundle.v1.WatchBundleResponse
	15, // [15:17] is the sub-list for method output_type
	13, // [13:15] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_cerbos_cloud_bundle_v1_bundle_proto_init() }
func file_cerbos_cloud_bundle_v1_bundle_proto_init() {
	if File_cerbos_cloud_bundle_v1_bundle_proto != nil {
		return
	}
	file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[5].OneofWrappers = []any{
		(*WatchBundleRequest_WatchLabel_)(nil),
		(*WatchBundleRequest_Heartbeat_)(nil),
	}
	file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes[6].OneofWrappers = []any{
		(*WatchBundleResponse_BundleUpdate)(nil),
		(*WatchBundleResponse_Reconnect_)(nil),
		(*WatchBundleResponse_BundleRemoved_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cerbos_cloud_bundle_v1_bundle_proto_rawDesc), len(file_cerbos_cloud_bundle_v1_bundle_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cerbos_cloud_bundle_v1_bundle_proto_goTypes,
		DependencyIndexes: file_cerbos_cloud_bundle_v1_bundle_proto_depIdxs,
		MessageInfos:      file_cerbos_cloud_bundle_v1_bundle_proto_msgTypes,
	}.Build()
	File_cerbos_cloud_bundle_v1_bundle_proto = out.File
	file_cerbos_cloud_bundle_v1_bundle_proto_goTypes = nil
	file_cerbos_cloud_bundle_v1_bundle_proto_depIdxs = nil
}
