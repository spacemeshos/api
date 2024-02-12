// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
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

	Id             []byte        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NodeId         []byte        `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Signature      []byte        `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	PublishEpoch   uint32        `protobuf:"varint,4,opt,name=publish_epoch,json=publishEpoch,proto3" json:"publish_epoch,omitempty"`
	Sequence       uint64        `protobuf:"varint,5,opt,name=sequence,proto3" json:"sequence,omitempty"`
	PrevAtx        []byte        `protobuf:"bytes,6,opt,name=prev_atx,json=prevAtx,proto3" json:"prev_atx,omitempty"`
	PositioningAtx []byte        `protobuf:"bytes,7,opt,name=positioning_atx,json=positioningAtx,proto3" json:"positioning_atx,omitempty"`
	CommittmentAtx []byte        `protobuf:"bytes,8,opt,name=committment_atx,json=committmentAtx,proto3" json:"committment_atx,omitempty"`
	InitialPost    *Post         `protobuf:"bytes,9,opt,name=initial_post,json=initialPost,proto3" json:"initial_post,omitempty"`
	Coinbase       string        `protobuf:"bytes,10,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	Units          uint32        `protobuf:"varint,11,opt,name=units,proto3" json:"units,omitempty"`
	BaseHeight     uint32        `protobuf:"varint,12,opt,name=base_height,json=baseHeight,proto3" json:"base_height,omitempty"`
	Ticks          uint32        `protobuf:"varint,13,opt,name=ticks,proto3" json:"ticks,omitempty"`
	PoetProof      *PoetProof    `protobuf:"bytes,14,opt,name=poet_proof,json=poetProof,proto3" json:"poet_proof,omitempty"`
	Post           *Post         `protobuf:"bytes,15,opt,name=post,proto3" json:"post,omitempty"`
	PostMeta       *PostMeta     `protobuf:"bytes,16,opt,name=post_meta,json=postMeta,proto3" json:"post_meta,omitempty"`
	VrfPostIndex   *VRFPostIndex `protobuf:"bytes,17,opt,name=vrf_post_index,json=vrfPostIndex,proto3" json:"vrf_post_index,omitempty"`
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

func (x *ActivationV1) GetNodeId() []byte {
	if x != nil {
		return x.NodeId
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

func (x *ActivationV1) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *ActivationV1) GetPrevAtx() []byte {
	if x != nil {
		return x.PrevAtx
	}
	return nil
}

func (x *ActivationV1) GetPositioningAtx() []byte {
	if x != nil {
		return x.PositioningAtx
	}
	return nil
}

func (x *ActivationV1) GetCommittmentAtx() []byte {
	if x != nil {
		return x.CommittmentAtx
	}
	return nil
}

func (x *ActivationV1) GetInitialPost() *Post {
	if x != nil {
		return x.InitialPost
	}
	return nil
}

func (x *ActivationV1) GetCoinbase() string {
	if x != nil {
		return x.Coinbase
	}
	return ""
}

func (x *ActivationV1) GetUnits() uint32 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *ActivationV1) GetBaseHeight() uint32 {
	if x != nil {
		return x.BaseHeight
	}
	return 0
}

func (x *ActivationV1) GetTicks() uint32 {
	if x != nil {
		return x.Ticks
	}
	return 0
}

func (x *ActivationV1) GetPoetProof() *PoetProof {
	if x != nil {
		return x.PoetProof
	}
	return nil
}

func (x *ActivationV1) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *ActivationV1) GetPostMeta() *PostMeta {
	if x != nil {
		return x.PostMeta
	}
	return nil
}

func (x *ActivationV1) GetVrfPostIndex() *VRFPostIndex {
	if x != nil {
		return x.VrfPostIndex
	}
	return nil
}

type VRFPostIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *VRFPostIndex) Reset() {
	*x = VRFPostIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VRFPostIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VRFPostIndex) ProtoMessage() {}

func (x *VRFPostIndex) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use VRFPostIndex.ProtoReflect.Descriptor instead.
func (*VRFPostIndex) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_activation_proto_rawDescGZIP(), []int{2}
}

func (x *VRFPostIndex) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

type PoetProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProofNodes [][]byte `protobuf:"bytes,1,rep,name=proof_nodes,json=proofNodes,proto3" json:"proof_nodes,omitempty"`
	Leaf       uint64   `protobuf:"varint,2,opt,name=leaf,proto3" json:"leaf,omitempty"`
}

func (x *PoetProof) Reset() {
	*x = PoetProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PoetProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoetProof) ProtoMessage() {}

func (x *PoetProof) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PoetProof.ProtoReflect.Descriptor instead.
func (*PoetProof) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_activation_proto_rawDescGZIP(), []int{3}
}

func (x *PoetProof) GetProofNodes() [][]byte {
	if x != nil {
		return x.ProofNodes
	}
	return nil
}

func (x *PoetProof) GetLeaf() uint64 {
	if x != nil {
		return x.Leaf
	}
	return 0
}

type PostMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Challenge []byte `protobuf:"bytes,1,opt,name=challenge,proto3" json:"challenge,omitempty"`
	Labels    uint64 `protobuf:"varint,2,opt,name=labels,proto3" json:"labels,omitempty"`
}

func (x *PostMeta) Reset() {
	*x = PostMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostMeta) ProtoMessage() {}

func (x *PostMeta) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PostMeta.ProtoReflect.Descriptor instead.
func (*PostMeta) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_activation_proto_rawDescGZIP(), []int{4}
}

func (x *PostMeta) GetChallenge() []byte {
	if x != nil {
		return x.Challenge
	}
	return nil
}

func (x *PostMeta) GetLabels() uint64 {
	if x != nil {
		return x.Labels
	}
	return 0
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce   uint32 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Indices []byte `protobuf:"bytes,2,opt,name=indices,proto3" json:"indices,omitempty"`
	Pow     uint64 `protobuf:"varint,3,opt,name=pow,proto3" json:"pow,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_activation_proto_rawDescGZIP(), []int{5}
}

func (x *Post) GetNonce() uint32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Post) GetIndices() []byte {
	if x != nil {
		return x.Indices
	}
	return nil
}

func (x *Post) GetPow() uint64 {
	if x != nil {
		return x.Pow
	}
	return 0
}

type ActivationStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartEpoch uint32 `protobuf:"varint,1,opt,name=start_epoch,json=startEpoch,proto3" json:"start_epoch,omitempty"`
	EndEpoch   uint32 `protobuf:"varint,2,opt,name=end_epoch,json=endEpoch,proto3" json:"end_epoch,omitempty"`
	Id         []byte `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	NodeId     []byte `protobuf:"bytes,4,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// CoinbaseFilter is not supported by database index and will result in sequential scan.
	// Apply EpochFilter together with CoinbaseFilter for better performance.
	Coinbase string `protobuf:"bytes,5,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	Watch    bool   `protobuf:"varint,6,opt,name=watch,proto3" json:"watch,omitempty"`
}

func (x *ActivationStreamRequest) Reset() {
	*x = ActivationStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivationStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivationStreamRequest) ProtoMessage() {}

func (x *ActivationStreamRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ActivationStreamRequest.ProtoReflect.Descriptor instead.
func (*ActivationStreamRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_activation_proto_rawDescGZIP(), []int{6}
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

	StartEpoch uint32 `protobuf:"varint,1,opt,name=start_epoch,json=startEpoch,proto3" json:"start_epoch,omitempty"`
	EndEpoch   uint32 `protobuf:"varint,2,opt,name=end_epoch,json=endEpoch,proto3" json:"end_epoch,omitempty"`
	Id         []byte `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	NodeId     []byte `protobuf:"bytes,4,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// CoinbaseFilter is not supported by database index and will result in sequential scan.
	// Apply EpochFilter together with CoinbaseFilter for better performance.
	Coinbase string `protobuf:"bytes,5,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	Offset   uint64 `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit    uint64 `protobuf:"varint,7,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ActivationRequest) Reset() {
	*x = ActivationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivationRequest) ProtoMessage() {}

func (x *ActivationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[7]
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
	return file_spacemesh_v2alpha1_activation_proto_rawDescGZIP(), []int{7}
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
		mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivationList) ProtoMessage() {}

func (x *ActivationList) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_activation_proto_msgTypes[8]
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
	return file_spacemesh_v2alpha1_activation_proto_rawDescGZIP(), []int{8}
}

func (x *ActivationList) GetActivations() []*Activation {
	if x != nil {
		return x.Activations
	}
	return nil
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
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x22, 0x98, 0x05, 0x0a, 0x0c, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x61, 0x74, 0x78, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x72, 0x65, 0x76, 0x41, 0x74, 0x78, 0x12, 0x27, 0x0a, 0x0f,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x74, 0x78, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69,
	0x6e, 0x67, 0x41, 0x74, 0x78, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x78, 0x12, 0x3b,
	0x0a, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x0b,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74,
	0x69, 0x63, 0x6b, 0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x70, 0x6f, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x09, 0x70, 0x6f, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x12, 0x2c, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74,
	0x12, 0x39, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x46, 0x0a, 0x0e, 0x76,
	0x72, 0x66, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x52, 0x46, 0x50, 0x6f, 0x73, 0x74,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x0c, 0x76, 0x72, 0x66, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x22, 0x24, 0x0a, 0x0c, 0x56, 0x52, 0x46, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x40, 0x0a, 0x09, 0x50, 0x6f, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x65, 0x61, 0x66, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x6c, 0x65, 0x61, 0x66, 0x22, 0x40, 0x0a, 0x08, 0x50,
	0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x48, 0x0a,
	0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69,
	0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x69, 0x6e,
	0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x70, 0x6f, 0x77, 0x22, 0xb2, 0x01, 0x0a, 0x17, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x70, 0x6f,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x45,
	0x70, 0x6f, 0x63, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x45, 0x70, 0x6f, 0x63,
	0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f,
	0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f,
	0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x61, 0x74, 0x63, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x77, 0x61, 0x74, 0x63, 0x68, 0x22, 0xc4, 0x01, 0x0a,
	0x11, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x45, 0x70, 0x6f, 0x63, 0x68,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x52, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x72, 0x0a, 0x17, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x57, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2b, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x01, 0x32, 0x66, 0x0a, 0x11, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x51, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0xde, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0f,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03,
	0x53, 0x58, 0x58, 0xaa, 0x02, 0x12, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x12, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x1e,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x13, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x56, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_spacemesh_v2alpha1_activation_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_spacemesh_v2alpha1_activation_proto_goTypes = []interface{}{
	(*Activation)(nil),              // 0: spacemesh.v2alpha1.Activation
	(*ActivationV1)(nil),            // 1: spacemesh.v2alpha1.ActivationV1
	(*VRFPostIndex)(nil),            // 2: spacemesh.v2alpha1.VRFPostIndex
	(*PoetProof)(nil),               // 3: spacemesh.v2alpha1.PoetProof
	(*PostMeta)(nil),                // 4: spacemesh.v2alpha1.PostMeta
	(*Post)(nil),                    // 5: spacemesh.v2alpha1.Post
	(*ActivationStreamRequest)(nil), // 6: spacemesh.v2alpha1.ActivationStreamRequest
	(*ActivationRequest)(nil),       // 7: spacemesh.v2alpha1.ActivationRequest
	(*ActivationList)(nil),          // 8: spacemesh.v2alpha1.ActivationList
}
var file_spacemesh_v2alpha1_activation_proto_depIdxs = []int32{
	1, // 0: spacemesh.v2alpha1.Activation.v1:type_name -> spacemesh.v2alpha1.ActivationV1
	5, // 1: spacemesh.v2alpha1.ActivationV1.initial_post:type_name -> spacemesh.v2alpha1.Post
	3, // 2: spacemesh.v2alpha1.ActivationV1.poet_proof:type_name -> spacemesh.v2alpha1.PoetProof
	5, // 3: spacemesh.v2alpha1.ActivationV1.post:type_name -> spacemesh.v2alpha1.Post
	4, // 4: spacemesh.v2alpha1.ActivationV1.post_meta:type_name -> spacemesh.v2alpha1.PostMeta
	2, // 5: spacemesh.v2alpha1.ActivationV1.vrf_post_index:type_name -> spacemesh.v2alpha1.VRFPostIndex
	0, // 6: spacemesh.v2alpha1.ActivationList.activations:type_name -> spacemesh.v2alpha1.Activation
	6, // 7: spacemesh.v2alpha1.ActivationStreamService.Stream:input_type -> spacemesh.v2alpha1.ActivationStreamRequest
	7, // 8: spacemesh.v2alpha1.ActivationService.List:input_type -> spacemesh.v2alpha1.ActivationRequest
	0, // 9: spacemesh.v2alpha1.ActivationStreamService.Stream:output_type -> spacemesh.v2alpha1.Activation
	8, // 10: spacemesh.v2alpha1.ActivationService.List:output_type -> spacemesh.v2alpha1.ActivationList
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
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
			switch v := v.(*VRFPostIndex); i {
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
			switch v := v.(*PoetProof); i {
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
			switch v := v.(*PostMeta); i {
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
			switch v := v.(*Post); i {
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
		file_spacemesh_v2alpha1_activation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_spacemesh_v2alpha1_activation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_spacemesh_v2alpha1_activation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Activation_V1)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spacemesh_v2alpha1_activation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
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
