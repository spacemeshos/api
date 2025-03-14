// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: spacemesh/v2beta1/layer.proto

package spacemeshv2beta1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
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

type Layer_LayerStatus int32

const (
	Layer_LAYER_STATUS_UNSPECIFIED Layer_LayerStatus = 0
	Layer_LAYER_STATUS_APPLIED     Layer_LayerStatus = 1 // applied by hare
	Layer_LAYER_STATUS_VERIFIED    Layer_LayerStatus = 2 // verified by tortoise
)

// Enum value maps for Layer_LayerStatus.
var (
	Layer_LayerStatus_name = map[int32]string{
		0: "LAYER_STATUS_UNSPECIFIED",
		1: "LAYER_STATUS_APPLIED",
		2: "LAYER_STATUS_VERIFIED",
	}
	Layer_LayerStatus_value = map[string]int32{
		"LAYER_STATUS_UNSPECIFIED": 0,
		"LAYER_STATUS_APPLIED":     1,
		"LAYER_STATUS_VERIFIED":    2,
	}
)

func (x Layer_LayerStatus) Enum() *Layer_LayerStatus {
	p := new(Layer_LayerStatus)
	*p = x
	return p
}

func (x Layer_LayerStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Layer_LayerStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_spacemesh_v2beta1_layer_proto_enumTypes[0].Descriptor()
}

func (Layer_LayerStatus) Type() protoreflect.EnumType {
	return &file_spacemesh_v2beta1_layer_proto_enumTypes[0]
}

func (x Layer_LayerStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Layer_LayerStatus.Descriptor instead.
func (Layer_LayerStatus) EnumDescriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_layer_proto_rawDescGZIP(), []int{0, 0}
}

type Layer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number              uint32            `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"` // layer number - not hash - layer content may change
	Status              Layer_LayerStatus `protobuf:"varint,2,opt,name=status,proto3,enum=spacemesh.v2beta1.Layer_LayerStatus" json:"status,omitempty"`
	ConsensusHash       string            `protobuf:"bytes,3,opt,name=consensus_hash,json=consensusHash,proto3" json:"consensus_hash,omitempty"`
	StateHash           []byte            `protobuf:"bytes,4,opt,name=state_hash,json=stateHash,proto3" json:"state_hash,omitempty"`                                 // fingerprint of the computed state at the layer
	CumulativeStateHash []byte            `protobuf:"bytes,5,opt,name=cumulative_state_hash,json=cumulativeStateHash,proto3" json:"cumulative_state_hash,omitempty"` // cumulative fingerprint that uniquely identifies state since genesis
	Block               *Block            `protobuf:"bytes,6,opt,name=block,proto3" json:"block,omitempty"`                                                          // layer's block
}

func (x *Layer) Reset() {
	*x = Layer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2beta1_layer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Layer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Layer) ProtoMessage() {}

func (x *Layer) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_layer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Layer.ProtoReflect.Descriptor instead.
func (*Layer) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_layer_proto_rawDescGZIP(), []int{0}
}

func (x *Layer) GetNumber() uint32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Layer) GetStatus() Layer_LayerStatus {
	if x != nil {
		return x.Status
	}
	return Layer_LAYER_STATUS_UNSPECIFIED
}

func (x *Layer) GetConsensusHash() string {
	if x != nil {
		return x.ConsensusHash
	}
	return ""
}

func (x *Layer) GetStateHash() []byte {
	if x != nil {
		return x.StateHash
	}
	return nil
}

func (x *Layer) GetCumulativeStateHash() []byte {
	if x != nil {
		return x.CumulativeStateHash
	}
	return nil
}

func (x *Layer) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

type LayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartLayer uint32    `protobuf:"varint,1,opt,name=start_layer,json=startLayer,proto3" json:"start_layer,omitempty"`                               // starting layer for the query
	EndLayer   uint32    `protobuf:"varint,2,opt,name=end_layer,json=endLayer,proto3" json:"end_layer,omitempty"`                                     // ending layer for the query
	Offset     uint64    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`                                                         // adjusts the starting point for data
	Limit      uint64    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`                                                           // specifies max number of items to fetch
	SortOrder  SortOrder `protobuf:"varint,5,opt,name=sort_order,json=sortOrder,proto3,enum=spacemesh.v2beta1.SortOrder" json:"sort_order,omitempty"` // specifies the sort order (default is ASC)
}

func (x *LayerRequest) Reset() {
	*x = LayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2beta1_layer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LayerRequest) ProtoMessage() {}

func (x *LayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_layer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LayerRequest.ProtoReflect.Descriptor instead.
func (*LayerRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_layer_proto_rawDescGZIP(), []int{1}
}

func (x *LayerRequest) GetStartLayer() uint32 {
	if x != nil {
		return x.StartLayer
	}
	return 0
}

func (x *LayerRequest) GetEndLayer() uint32 {
	if x != nil {
		return x.EndLayer
	}
	return 0
}

func (x *LayerRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *LayerRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *LayerRequest) GetSortOrder() SortOrder {
	if x != nil {
		return x.SortOrder
	}
	return SortOrder_ASC
}

type LayerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Layers []*Layer `protobuf:"bytes,1,rep,name=layers,proto3" json:"layers,omitempty"` // list of layers
}

func (x *LayerList) Reset() {
	*x = LayerList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2beta1_layer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LayerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LayerList) ProtoMessage() {}

func (x *LayerList) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_layer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LayerList.ProtoReflect.Descriptor instead.
func (*LayerList) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_layer_proto_rawDescGZIP(), []int{2}
}

func (x *LayerList) GetLayers() []*Layer {
	if x != nil {
		return x.Layers
	}
	return nil
}

type LayerStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartLayer uint32 `protobuf:"varint,1,opt,name=start_layer,json=startLayer,proto3" json:"start_layer,omitempty"`
	EndLayer   uint32 `protobuf:"varint,2,opt,name=end_layer,json=endLayer,proto3" json:"end_layer,omitempty"`
	Watch      bool   `protobuf:"varint,3,opt,name=watch,proto3" json:"watch,omitempty"`
}

func (x *LayerStreamRequest) Reset() {
	*x = LayerStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2beta1_layer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LayerStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LayerStreamRequest) ProtoMessage() {}

func (x *LayerStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_layer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LayerStreamRequest.ProtoReflect.Descriptor instead.
func (*LayerStreamRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_layer_proto_rawDescGZIP(), []int{3}
}

func (x *LayerStreamRequest) GetStartLayer() uint32 {
	if x != nil {
		return x.StartLayer
	}
	return 0
}

func (x *LayerStreamRequest) GetEndLayer() uint32 {
	if x != nil {
		return x.EndLayer
	}
	return 0
}

func (x *LayerStreamRequest) GetWatch() bool {
	if x != nil {
		return x.Watch
	}
	return false
}

var File_spacemesh_v2beta1_layer_proto protoreflect.FileDescriptor

var file_spacemesh_v2beta1_layer_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x73,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x02,
	0x0a, 0x05, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x24, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x13, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2e, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x60, 0x0a, 0x0b, 0x4c, 0x61, 0x79, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x18, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x45, 0x44, 0x10, 0x01, 0x12, 0x19,
	0x0a, 0x15, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x56,
	0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x02, 0x22, 0xb7, 0x01, 0x0a, 0x0c, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x65,
	0x6e, 0x64, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x65, 0x6e, 0x64, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53,
	0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x22, 0x3d, 0x0a, 0x09, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x30, 0x0a, 0x06, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x22, 0x68, 0x0a, 0x12, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64,
	0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e,
	0x64, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x61, 0x74, 0x63, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x77, 0x61, 0x74, 0x63, 0x68, 0x32, 0x94, 0x01, 0x0a,
	0x0c, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x4c, 0x69,
	0x73, 0x74, 0x1a, 0x0f, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x76, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x32, 0x73, 0x0a, 0x12, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c,
	0x61, 0x79, 0x65, 0x72, 0x30, 0x01, 0x1a, 0x10, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x42, 0xd2, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x42, 0x0a, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58,
	0xaa, 0x02, 0x11, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x56, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x11, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1d, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spacemesh_v2beta1_layer_proto_rawDescOnce sync.Once
	file_spacemesh_v2beta1_layer_proto_rawDescData = file_spacemesh_v2beta1_layer_proto_rawDesc
)

func file_spacemesh_v2beta1_layer_proto_rawDescGZIP() []byte {
	file_spacemesh_v2beta1_layer_proto_rawDescOnce.Do(func() {
		file_spacemesh_v2beta1_layer_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacemesh_v2beta1_layer_proto_rawDescData)
	})
	return file_spacemesh_v2beta1_layer_proto_rawDescData
}

var file_spacemesh_v2beta1_layer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spacemesh_v2beta1_layer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spacemesh_v2beta1_layer_proto_goTypes = []interface{}{
	(Layer_LayerStatus)(0),     // 0: spacemesh.v2beta1.Layer.LayerStatus
	(*Layer)(nil),              // 1: spacemesh.v2beta1.Layer
	(*LayerRequest)(nil),       // 2: spacemesh.v2beta1.LayerRequest
	(*LayerList)(nil),          // 3: spacemesh.v2beta1.LayerList
	(*LayerStreamRequest)(nil), // 4: spacemesh.v2beta1.LayerStreamRequest
	(*Block)(nil),              // 5: spacemesh.v2beta1.Block
	(SortOrder)(0),             // 6: spacemesh.v2beta1.SortOrder
}
var file_spacemesh_v2beta1_layer_proto_depIdxs = []int32{
	0, // 0: spacemesh.v2beta1.Layer.status:type_name -> spacemesh.v2beta1.Layer.LayerStatus
	5, // 1: spacemesh.v2beta1.Layer.block:type_name -> spacemesh.v2beta1.Block
	6, // 2: spacemesh.v2beta1.LayerRequest.sort_order:type_name -> spacemesh.v2beta1.SortOrder
	1, // 3: spacemesh.v2beta1.LayerList.layers:type_name -> spacemesh.v2beta1.Layer
	2, // 4: spacemesh.v2beta1.LayerService.List:input_type -> spacemesh.v2beta1.LayerRequest
	4, // 5: spacemesh.v2beta1.LayerStreamService.Stream:input_type -> spacemesh.v2beta1.LayerStreamRequest
	3, // 6: spacemesh.v2beta1.LayerService.List:output_type -> spacemesh.v2beta1.LayerList
	1, // 7: spacemesh.v2beta1.LayerStreamService.Stream:output_type -> spacemesh.v2beta1.Layer
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_spacemesh_v2beta1_layer_proto_init() }
func file_spacemesh_v2beta1_layer_proto_init() {
	if File_spacemesh_v2beta1_layer_proto != nil {
		return
	}
	file_spacemesh_v2beta1_block_proto_init()
	file_spacemesh_v2beta1_v2beta1_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spacemesh_v2beta1_layer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Layer); i {
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
		file_spacemesh_v2beta1_layer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LayerRequest); i {
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
		file_spacemesh_v2beta1_layer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LayerList); i {
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
		file_spacemesh_v2beta1_layer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LayerStreamRequest); i {
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
			RawDescriptor: file_spacemesh_v2beta1_layer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_spacemesh_v2beta1_layer_proto_goTypes,
		DependencyIndexes: file_spacemesh_v2beta1_layer_proto_depIdxs,
		EnumInfos:         file_spacemesh_v2beta1_layer_proto_enumTypes,
		MessageInfos:      file_spacemesh_v2beta1_layer_proto_msgTypes,
	}.Build()
	File_spacemesh_v2beta1_layer_proto = out.File
	file_spacemesh_v2beta1_layer_proto_rawDesc = nil
	file_spacemesh_v2beta1_layer_proto_goTypes = nil
	file_spacemesh_v2beta1_layer_proto_depIdxs = nil
}
