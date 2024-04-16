// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: spacemesh/v2alpha1/activation.proto

package spacemeshv2alpha1

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

type Activation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Versioned:
	//
	//	*Activation_V1
	Versioned isActivation_Versioned `protobuf_oneof:"versioned"`
}

func (x *Activation) Reset() {
	*x = Activation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activation) ProtoMessage() {}

func (x *Activation) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activation.ProtoReflect.Descriptor instead.
func (*Activation) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_activation_proto_rawDescGZIP(), []int{0}
}

func (m *Activation) GetVersioned() isActivation_Versioned {
	if m != nil {
		return m.Versioned
	}
	return nil
}

func (x *Activation) GetV1() *ActivationV1 {
	if x, ok := x.GetVersioned().(*Activation_V1); ok {
		return x.V1
	}
	return nil
}

type isActivation_Versioned interface {
	isActivation_Versioned()
}

type Activation_V1 struct {
	V1 *ActivationV1 `protobuf:"bytes,1,opt,name=v1,proto3,oneof"`
}

func (*Activation_V1) isActivation_Versioned() {}

type ActivationV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SmesherId      []byte `protobuf:"bytes,2,opt,name=smesher_id,json=smesherId,proto3" json:"smesher_id,omitempty"`
	Signature      []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	PublishEpoch   uint32 `protobuf:"varint,4,opt,name=publish_epoch,json=publishEpoch,proto3" json:"publish_epoch,omitempty"`
	PreviousAtx    []byte `protobuf:"bytes,5,opt,name=previous_atx,json=previousAtx,proto3" json:"previous_atx,omitempty"`
	PositioningAtx []byte `protobuf:"bytes,6,opt,name=positioning_atx,json=positioningAtx,proto3" json:"positioning_atx,omitempty"`
	Coinbase       string `protobuf:"bytes,7,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	Weight         uint32 `protobuf:"varint,8,opt,name=weight,proto3" json:"weight,omitempty"`
	Height         uint32 `protobuf:"varint,9,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *ActivationV1) Reset() {
	*x = ActivationV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivationV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivationV1) ProtoMessage() {}

func (x *ActivationV1) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivationV1.ProtoReflect.Descriptor instead.
func (*ActivationV1) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_activation_proto_rawDescGZIP(), []int{1}
}

func (x *ActivationV1) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ActivationV1) GetSmesherId() []byte {
	if x != nil {
		return x.SmesherId
	}
	return nil
}

func (x *ActivationV1) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *ActivationV1) GetPublishEpoch() uint32 {
	if x != nil {
		return x.PublishEpoch
	}
	return 0
}

func (x *ActivationV1) GetPreviousAtx() []byte {
	if x != nil {
		return x.PreviousAtx
	}
	return nil
}

func (x *ActivationV1) GetPositioningAtx() []byte {
	if x != nil {
		return x.PositioningAtx
	}
	return nil
}

func (x *ActivationV1) GetCoinbase() string {
	if x != nil {
		return x.Coinbase
	}
	return ""
}

func (x *ActivationV1) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *ActivationV1) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type ActivationStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartEpoch uint32 `protobuf:"varint,1,opt,name=start_epoch,json=startEpoch,proto3" json:"start_epoch,omitempty"` // Apply `start_epoch/end_epoch` filters together with `coinbase` filter for better performance.
	EndEpoch   uint32 `protobuf:"varint,2,opt,name=end_epoch,json=endEpoch,proto3" json:"end_epoch,omitempty"`
	Id         []byte `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	NodeId     []byte `protobuf:"bytes,4,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Coinbase   string `protobuf:"bytes,5,opt,name=coinbase,proto3" json:"coinbase,omitempty"` // `coinbase` filter is not supported by database index and will result in sequential scan.
	Watch      bool   `protobuf:"varint,6,opt,name=watch,proto3" json:"watch,omitempty"`
}

func (x *ActivationStreamRequest) Reset() {
	*x = ActivationStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivationStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivationStreamRequest) ProtoMessage() {}

func (x *ActivationStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivationStreamRequest.ProtoReflect.Descriptor instead.
func (*ActivationStreamRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_activation_proto_rawDescGZIP(), []int{2}
}

func (x *ActivationStreamRequest) GetStartEpoch() uint32 {
	if x != nil {
		return x.StartEpoch
	}
	return 0
}

func (x *ActivationStreamRequest) GetEndEpoch() uint32 {
	if x != nil {
		return x.EndEpoch
	}
	return 0
}

func (x *ActivationStreamRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ActivationStreamRequest) GetNodeId() []byte {
	if x != nil {
		return x.NodeId
	}
	return nil
}

func (x *ActivationStreamRequest) GetCoinbase() string {
	if x != nil {
		return x.Coinbase
	}
	return ""
}

func (x *ActivationStreamRequest) GetWatch() bool {
	if x != nil {
		return x.Watch
	}
	return false
}

type ActivationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartEpoch uint32 `protobuf:"varint,1,opt,name=start_epoch,json=startEpoch,proto3" json:"start_epoch,omitempty"` // Apply `start_epoch/end_epoch` filters together with `coinbase` filter for better performance.
	EndEpoch   uint32 `protobuf:"varint,2,opt,name=end_epoch,json=endEpoch,proto3" json:"end_epoch,omitempty"`
	Id         []byte `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	NodeId     []byte `protobuf:"bytes,4,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Coinbase   string `protobuf:"bytes,5,opt,name=coinbase,proto3" json:"coinbase,omitempty"` // `coinbase` filter is not supported by database index and will result in sequential scan.
	Offset     uint64 `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit      uint64 `protobuf:"varint,7,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ActivationRequest) Reset() {
	*x = ActivationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivationRequest) ProtoMessage() {}

func (x *ActivationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivationRequest.ProtoReflect.Descriptor instead.
func (*ActivationRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_activation_proto_rawDescGZIP(), []int{3}
}

func (x *ActivationRequest) GetStartEpoch() uint32 {
	if x != nil {
		return x.StartEpoch
	}
	return 0
}

func (x *ActivationRequest) GetEndEpoch() uint32 {
	if x != nil {
		return x.EndEpoch
	}
	return 0
}

func (x *ActivationRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ActivationRequest) GetNodeId() []byte {
	if x != nil {
		return x.NodeId
	}
	return nil
}

func (x *ActivationRequest) GetCoinbase() string {
	if x != nil {
		return x.Coinbase
	}
	return ""
}

func (x *ActivationRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ActivationRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ActivationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activations []*Activation `protobuf:"bytes,1,rep,name=activations,proto3" json:"activations,omitempty"`
}

func (x *ActivationList) Reset() {
	*x = ActivationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivationList) ProtoMessage() {}

func (x *ActivationList) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivationList.ProtoReflect.Descriptor instead.
func (*ActivationList) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_activation_proto_rawDescGZIP(), []int{4}
}

func (x *ActivationList) GetActivations() []*Activation {
	if x != nil {
		return x.Activations
	}
	return nil
}

type ActivationsCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Epoch *uint32 `protobuf:"varint,1,opt,name=epoch,proto3,oneof" json:"epoch,omitempty"`
}

func (x *ActivationsCountRequest) Reset() {
	*x = ActivationsCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivationsCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivationsCountRequest) ProtoMessage() {}

func (x *ActivationsCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivationsCountRequest.ProtoReflect.Descriptor instead.
func (*ActivationsCountRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_activation_proto_rawDescGZIP(), []int{5}
}

func (x *ActivationsCountRequest) GetEpoch() uint32 {
	if x != nil && x.Epoch != nil {
		return *x.Epoch
	}
	return 0
}

type ActivationsCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ActivationsCountResponse) Reset() {
	*x = ActivationsCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivationsCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivationsCountResponse) ProtoMessage() {}

func (x *ActivationsCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivationsCountResponse.ProtoReflect.Descriptor instead.
func (*ActivationsCountResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_activation_proto_rawDescGZIP(), []int{6}
}

func (x *ActivationsCountResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_spacemesh_v2alpha1_activation_proto protoreflect.FileDescriptor

var file_spacemesh_v2alpha1_activation_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x4d, 0x0a, 0x0a, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x02, 0x76, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x56, 0x31, 0x48, 0x00, 0x52, 0x02, 0x76, 0x31, 0x42, 0x0b, 0x0a, 0x09, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x22, 0x98, 0x02, 0x0a, 0x0c, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6d, 0x65,
	0x73, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x6d, 0x65, 0x73, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x61, 0x74, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x41, 0x74, 0x78, 0x12, 0x27,
	0x0a, 0x0f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x74,
	0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x67, 0x41, 0x74, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x22, 0xb2, 0x01, 0x0a, 0x17, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68,
	0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x61, 0x74, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x77, 0x61, 0x74, 0x63, 0x68, 0x22, 0xc4, 0x01, 0x0a, 0x11, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12,
	0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x52, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x40, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x3e, 0x0a, 0x17, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52,
	0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x22, 0x30, 0x0a, 0x18, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x72, 0x0a, 0x17, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x57, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x01, 0x32, 0xd5, 0x01, 0x0a, 0x11, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x51, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x6d, 0x0a, 0x10, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0xde, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0f, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58,
	0x58, 0xaa, 0x02, 0x12, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x56, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x12, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x1e, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spacemesh_v2alpha1_activation_proto_rawDescOnce sync.Once
	file_spacemesh_v2alpha1_activation_proto_rawDescData = file_spacemesh_v2alpha1_activation_proto_rawDesc
)

func file_spacemesh_v2alpha1_activation_proto_rawDescGZIP() []byte {
	file_spacemesh_v2alpha1_activation_proto_rawDescOnce.Do(func() {
		file_spacemesh_v2alpha1_activation_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacemesh_v2alpha1_activation_proto_rawDescData)
	})
	return file_spacemesh_v2alpha1_activation_proto_rawDescData
}

var file_spacemesh_v2alpha1_activation_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spacemesh_v2alpha1_activation_proto_goTypes = []interface{}{
	(*Activation)(nil),               // 0: spacemesh.v2alpha1.Activation
	(*ActivationV1)(nil),             // 1: spacemesh.v2alpha1.ActivationV1
	(*ActivationStreamRequest)(nil),  // 2: spacemesh.v2alpha1.ActivationStreamRequest
	(*ActivationRequest)(nil),        // 3: spacemesh.v2alpha1.ActivationRequest
	(*ActivationList)(nil),           // 4: spacemesh.v2alpha1.ActivationList
	(*ActivationsCountRequest)(nil),  // 5: spacemesh.v2alpha1.ActivationsCountRequest
	(*ActivationsCountResponse)(nil), // 6: spacemesh.v2alpha1.ActivationsCountResponse
}
var file_spacemesh_v2alpha1_activation_proto_depIdxs = []int32{
	1, // 0: spacemesh.v2alpha1.Activation.v1:type_name -> spacemesh.v2alpha1.ActivationV1
	0, // 1: spacemesh.v2alpha1.ActivationList.activations:type_name -> spacemesh.v2alpha1.Activation
	2, // 2: spacemesh.v2alpha1.ActivationStreamService.Stream:input_type -> spacemesh.v2alpha1.ActivationStreamRequest
	3, // 3: spacemesh.v2alpha1.ActivationService.List:input_type -> spacemesh.v2alpha1.ActivationRequest
	5, // 4: spacemesh.v2alpha1.ActivationService.ActivationsCount:input_type -> spacemesh.v2alpha1.ActivationsCountRequest
	0, // 5: spacemesh.v2alpha1.ActivationStreamService.Stream:output_type -> spacemesh.v2alpha1.Activation
	4, // 6: spacemesh.v2alpha1.ActivationService.List:output_type -> spacemesh.v2alpha1.ActivationList
	6, // 7: spacemesh.v2alpha1.ActivationService.ActivationsCount:output_type -> spacemesh.v2alpha1.ActivationsCountResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_spacemesh_v2alpha1_activation_proto_init() }
func file_spacemesh_v2alpha1_activation_proto_init() {
	if File_spacemesh_v2alpha1_activation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spacemesh_v2alpha1_activation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activation); i {
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
		file_spacemesh_v2alpha1_activation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivationV1); i {
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
		file_spacemesh_v2alpha1_activation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivationStreamRequest); i {
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
		file_spacemesh_v2alpha1_activation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivationRequest); i {
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
		file_spacemesh_v2alpha1_activation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivationList); i {
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
		file_spacemesh_v2alpha1_activation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivationsCountRequest); i {
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
		file_spacemesh_v2alpha1_activation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivationsCountResponse); i {
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
	file_spacemesh_v2alpha1_activation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Activation_V1)(nil),
	}
	file_spacemesh_v2alpha1_activation_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spacemesh_v2alpha1_activation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_spacemesh_v2alpha1_activation_proto_goTypes,
		DependencyIndexes: file_spacemesh_v2alpha1_activation_proto_depIdxs,
		MessageInfos:      file_spacemesh_v2alpha1_activation_proto_msgTypes,
	}.Build()
	File_spacemesh_v2alpha1_activation_proto = out.File
	file_spacemesh_v2alpha1_activation_proto_rawDesc = nil
	file_spacemesh_v2alpha1_activation_proto_goTypes = nil
	file_spacemesh_v2alpha1_activation_proto_depIdxs = nil
}
