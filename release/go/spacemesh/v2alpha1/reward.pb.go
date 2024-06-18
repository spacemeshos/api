// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: spacemesh/v2alpha1/reward.proto

package spacemeshv2alpha1

import (
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

type Reward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Layer       uint32 `protobuf:"varint,1,opt,name=layer,proto3" json:"layer,omitempty"`                                // layer award was paid in
	Total       uint64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`                                // total reward paid in smidge (sum of tx fee and layer reward)
	LayerReward uint64 `protobuf:"varint,3,opt,name=layer_reward,json=layerReward,proto3" json:"layer_reward,omitempty"` // tx_fee = total - layer_reward
	Coinbase    string `protobuf:"bytes,5,opt,name=coinbase,proto3" json:"coinbase,omitempty"`                           // account awarded this reward
	Smesher     []byte `protobuf:"bytes,6,opt,name=smesher,proto3" json:"smesher,omitempty"`                             // id of smesher who earned this reward
}

func (x *Reward) Reset() {
	*x = Reward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_reward_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reward) ProtoMessage() {}

func (x *Reward) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_reward_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reward.ProtoReflect.Descriptor instead.
func (*Reward) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_reward_proto_rawDescGZIP(), []int{0}
}

func (x *Reward) GetLayer() uint32 {
	if x != nil {
		return x.Layer
	}
	return 0
}

func (x *Reward) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Reward) GetLayerReward() uint64 {
	if x != nil {
		return x.LayerReward
	}
	return 0
}

func (x *Reward) GetCoinbase() string {
	if x != nil {
		return x.Coinbase
	}
	return ""
}

func (x *Reward) GetSmesher() []byte {
	if x != nil {
		return x.Smesher
	}
	return nil
}

type RewardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartLayer uint32 `protobuf:"varint,1,opt,name=start_layer,json=startLayer,proto3" json:"start_layer,omitempty"`
	EndLayer   uint32 `protobuf:"varint,2,opt,name=end_layer,json=endLayer,proto3" json:"end_layer,omitempty"`
	// Types that are assignable to FilterBy:
	//
	//	*RewardRequest_Coinbase
	//	*RewardRequest_Smesher
	FilterBy isRewardRequest_FilterBy `protobuf_oneof:"filter_by"`
	Offset   uint64                   `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"` // adjusts the starting point for data
	Limit    uint64                   `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`   // specifies max number of items to fetch
}

func (x *RewardRequest) Reset() {
	*x = RewardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_reward_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardRequest) ProtoMessage() {}

func (x *RewardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_reward_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardRequest.ProtoReflect.Descriptor instead.
func (*RewardRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_reward_proto_rawDescGZIP(), []int{1}
}

func (x *RewardRequest) GetStartLayer() uint32 {
	if x != nil {
		return x.StartLayer
	}
	return 0
}

func (x *RewardRequest) GetEndLayer() uint32 {
	if x != nil {
		return x.EndLayer
	}
	return 0
}

func (m *RewardRequest) GetFilterBy() isRewardRequest_FilterBy {
	if m != nil {
		return m.FilterBy
	}
	return nil
}

func (x *RewardRequest) GetCoinbase() string {
	if x, ok := x.GetFilterBy().(*RewardRequest_Coinbase); ok {
		return x.Coinbase
	}
	return ""
}

func (x *RewardRequest) GetSmesher() []byte {
	if x, ok := x.GetFilterBy().(*RewardRequest_Smesher); ok {
		return x.Smesher
	}
	return nil
}

func (x *RewardRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *RewardRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type isRewardRequest_FilterBy interface {
	isRewardRequest_FilterBy()
}

type RewardRequest_Coinbase struct {
	Coinbase string `protobuf:"bytes,3,opt,name=coinbase,proto3,oneof"`
}

type RewardRequest_Smesher struct {
	Smesher []byte `protobuf:"bytes,4,opt,name=smesher,proto3,oneof"`
}

func (*RewardRequest_Coinbase) isRewardRequest_FilterBy() {}

func (*RewardRequest_Smesher) isRewardRequest_FilterBy() {}

type RewardList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rewards []*Reward `protobuf:"bytes,1,rep,name=rewards,proto3" json:"rewards,omitempty"`
	Total   uint32    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"` // Total number of rewards in the requested range.
}

func (x *RewardList) Reset() {
	*x = RewardList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_reward_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardList) ProtoMessage() {}

func (x *RewardList) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_reward_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardList.ProtoReflect.Descriptor instead.
func (*RewardList) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_reward_proto_rawDescGZIP(), []int{2}
}

func (x *RewardList) GetRewards() []*Reward {
	if x != nil {
		return x.Rewards
	}
	return nil
}

func (x *RewardList) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type RewardStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartLayer uint32 `protobuf:"varint,1,opt,name=start_layer,json=startLayer,proto3" json:"start_layer,omitempty"`
	EndLayer   uint32 `protobuf:"varint,2,opt,name=end_layer,json=endLayer,proto3" json:"end_layer,omitempty"`
	// Types that are assignable to FilterBy:
	//
	//	*RewardStreamRequest_Coinbase
	//	*RewardStreamRequest_Smesher
	FilterBy isRewardStreamRequest_FilterBy `protobuf_oneof:"filter_by"`
	Watch    bool                           `protobuf:"varint,5,opt,name=watch,proto3" json:"watch,omitempty"`
}

func (x *RewardStreamRequest) Reset() {
	*x = RewardStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_reward_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardStreamRequest) ProtoMessage() {}

func (x *RewardStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_reward_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardStreamRequest.ProtoReflect.Descriptor instead.
func (*RewardStreamRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_reward_proto_rawDescGZIP(), []int{3}
}

func (x *RewardStreamRequest) GetStartLayer() uint32 {
	if x != nil {
		return x.StartLayer
	}
	return 0
}

func (x *RewardStreamRequest) GetEndLayer() uint32 {
	if x != nil {
		return x.EndLayer
	}
	return 0
}

func (m *RewardStreamRequest) GetFilterBy() isRewardStreamRequest_FilterBy {
	if m != nil {
		return m.FilterBy
	}
	return nil
}

func (x *RewardStreamRequest) GetCoinbase() string {
	if x, ok := x.GetFilterBy().(*RewardStreamRequest_Coinbase); ok {
		return x.Coinbase
	}
	return ""
}

func (x *RewardStreamRequest) GetSmesher() []byte {
	if x, ok := x.GetFilterBy().(*RewardStreamRequest_Smesher); ok {
		return x.Smesher
	}
	return nil
}

func (x *RewardStreamRequest) GetWatch() bool {
	if x != nil {
		return x.Watch
	}
	return false
}

type isRewardStreamRequest_FilterBy interface {
	isRewardStreamRequest_FilterBy()
}

type RewardStreamRequest_Coinbase struct {
	Coinbase string `protobuf:"bytes,3,opt,name=coinbase,proto3,oneof"`
}

type RewardStreamRequest_Smesher struct {
	Smesher []byte `protobuf:"bytes,4,opt,name=smesher,proto3,oneof"`
}

func (*RewardStreamRequest_Coinbase) isRewardStreamRequest_FilterBy() {}

func (*RewardStreamRequest_Smesher) isRewardStreamRequest_FilterBy() {}

var File_spacemesh_v2alpha1_reward_proto protoreflect.FileDescriptor

var file_spacemesh_v2alpha1_reward_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6d, 0x65, 0x73,
	0x68, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x6d, 0x65, 0x73, 0x68,
	0x65, 0x72, 0x22, 0xc2, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x07, 0x73, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x07, 0x73, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x22, 0x58, 0x0a, 0x0a, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0xb0, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e,
	0x64, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65,
	0x6e, 0x64, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x73, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x73, 0x6d, 0x65, 0x73, 0x68, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x61, 0x74, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x77, 0x61, 0x74, 0x63, 0x68, 0x42, 0x0b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x5f, 0x62, 0x79, 0x32, 0x66, 0x0a, 0x0d, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x0a, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x04, 0x12, 0x02, 0x56, 0x32, 0x32, 0x78, 0x0a, 0x13,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x27, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x30, 0x01, 0x1a, 0x10, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x42, 0xda, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x42, 0x0b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x53,
	0x58, 0x58, 0xaa, 0x02, 0x12, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x56,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x12, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x1e, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x56, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spacemesh_v2alpha1_reward_proto_rawDescOnce sync.Once
	file_spacemesh_v2alpha1_reward_proto_rawDescData = file_spacemesh_v2alpha1_reward_proto_rawDesc
)

func file_spacemesh_v2alpha1_reward_proto_rawDescGZIP() []byte {
	file_spacemesh_v2alpha1_reward_proto_rawDescOnce.Do(func() {
		file_spacemesh_v2alpha1_reward_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacemesh_v2alpha1_reward_proto_rawDescData)
	})
	return file_spacemesh_v2alpha1_reward_proto_rawDescData
}

var file_spacemesh_v2alpha1_reward_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spacemesh_v2alpha1_reward_proto_goTypes = []interface{}{
	(*Reward)(nil),              // 0: spacemesh.v2alpha1.Reward
	(*RewardRequest)(nil),       // 1: spacemesh.v2alpha1.RewardRequest
	(*RewardList)(nil),          // 2: spacemesh.v2alpha1.RewardList
	(*RewardStreamRequest)(nil), // 3: spacemesh.v2alpha1.RewardStreamRequest
}
var file_spacemesh_v2alpha1_reward_proto_depIdxs = []int32{
	0, // 0: spacemesh.v2alpha1.RewardList.rewards:type_name -> spacemesh.v2alpha1.Reward
	1, // 1: spacemesh.v2alpha1.RewardService.List:input_type -> spacemesh.v2alpha1.RewardRequest
	3, // 2: spacemesh.v2alpha1.RewardStreamService.Stream:input_type -> spacemesh.v2alpha1.RewardStreamRequest
	2, // 3: spacemesh.v2alpha1.RewardService.List:output_type -> spacemesh.v2alpha1.RewardList
	0, // 4: spacemesh.v2alpha1.RewardStreamService.Stream:output_type -> spacemesh.v2alpha1.Reward
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spacemesh_v2alpha1_reward_proto_init() }
func file_spacemesh_v2alpha1_reward_proto_init() {
	if File_spacemesh_v2alpha1_reward_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spacemesh_v2alpha1_reward_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reward); i {
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
		file_spacemesh_v2alpha1_reward_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardRequest); i {
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
		file_spacemesh_v2alpha1_reward_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardList); i {
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
		file_spacemesh_v2alpha1_reward_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardStreamRequest); i {
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
	file_spacemesh_v2alpha1_reward_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*RewardRequest_Coinbase)(nil),
		(*RewardRequest_Smesher)(nil),
	}
	file_spacemesh_v2alpha1_reward_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*RewardStreamRequest_Coinbase)(nil),
		(*RewardStreamRequest_Smesher)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spacemesh_v2alpha1_reward_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_spacemesh_v2alpha1_reward_proto_goTypes,
		DependencyIndexes: file_spacemesh_v2alpha1_reward_proto_depIdxs,
		MessageInfos:      file_spacemesh_v2alpha1_reward_proto_msgTypes,
	}.Build()
	File_spacemesh_v2alpha1_reward_proto = out.File
	file_spacemesh_v2alpha1_reward_proto_rawDesc = nil
	file_spacemesh_v2alpha1_reward_proto_goTypes = nil
	file_spacemesh_v2alpha1_reward_proto_depIdxs = nil
}
