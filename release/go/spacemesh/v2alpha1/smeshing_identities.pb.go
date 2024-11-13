// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: spacemesh/v2alpha1/smeshing_identities.proto

package spacemeshv2alpha1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IdentityStatus int32

const (
	IdentityStatus_UNSPECIFIED                          IdentityStatus = 0
	IdentityStatus_ERROR                                IdentityStatus = 1
	IdentityStatus_WAIT_FOR_ATX_SYNCING                 IdentityStatus = 2
	IdentityStatus_WAITING_FOR_POET_REGISTRATION_WINDOW IdentityStatus = 6
	IdentityStatus_POET_CHALLENGE_READY                 IdentityStatus = 7
	IdentityStatus_POET_REGISTERED                      IdentityStatus = 8
	IdentityStatus_WAIT_FOR_POET_ROUND_END              IdentityStatus = 9
	IdentityStatus_POET_PROOF_RECEIVED                  IdentityStatus = 10
	IdentityStatus_POET_PROOF_FAILED                    IdentityStatus = 11
	IdentityStatus_GENERATING_POST_PROOF                IdentityStatus = 12
	IdentityStatus_POST_PROOF_READY                     IdentityStatus = 13
	IdentityStatus_POST_PROOF_FAILED                    IdentityStatus = 14
	IdentityStatus_ATX_READY                            IdentityStatus = 15
	IdentityStatus_ATX_BROADCASTED                      IdentityStatus = 16
)

// Enum value maps for IdentityStatus.
var (
	IdentityStatus_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "ERROR",
		2:  "WAIT_FOR_ATX_SYNCING",
		6:  "WAITING_FOR_POET_REGISTRATION_WINDOW",
		7:  "POET_CHALLENGE_READY",
		8:  "POET_REGISTERED",
		9:  "WAIT_FOR_POET_ROUND_END",
		10: "POET_PROOF_RECEIVED",
		11: "POET_PROOF_FAILED",
		12: "GENERATING_POST_PROOF",
		13: "POST_PROOF_READY",
		14: "POST_PROOF_FAILED",
		15: "ATX_READY",
		16: "ATX_BROADCASTED",
	}
	IdentityStatus_value = map[string]int32{
		"UNSPECIFIED":                          0,
		"ERROR":                                1,
		"WAIT_FOR_ATX_SYNCING":                 2,
		"WAITING_FOR_POET_REGISTRATION_WINDOW": 6,
		"POET_CHALLENGE_READY":                 7,
		"POET_REGISTERED":                      8,
		"WAIT_FOR_POET_ROUND_END":              9,
		"POET_PROOF_RECEIVED":                  10,
		"POET_PROOF_FAILED":                    11,
		"GENERATING_POST_PROOF":                12,
		"POST_PROOF_READY":                     13,
		"POST_PROOF_FAILED":                    14,
		"ATX_READY":                            15,
		"ATX_BROADCASTED":                      16,
	}
)

func (x IdentityStatus) Enum() *IdentityStatus {
	p := new(IdentityStatus)
	*p = x
	return p
}

func (x IdentityStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdentityStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_spacemesh_v2alpha1_smeshing_identities_proto_enumTypes[0].Descriptor()
}

func (IdentityStatus) Type() protoreflect.EnumType {
	return &file_spacemesh_v2alpha1_smeshing_identities_proto_enumTypes[0]
}

func (x IdentityStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdentityStatus.Descriptor instead.
func (IdentityStatus) EnumDescriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_smeshing_identities_proto_rawDescGZIP(), []int{0}
}

type IdentityStatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IdentityStatesRequest) Reset() {
	*x = IdentityStatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityStatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityStatesRequest) ProtoMessage() {}

func (x *IdentityStatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityStatesRequest.ProtoReflect.Descriptor instead.
func (*IdentityStatesRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_smeshing_identities_proto_rawDescGZIP(), []int{0}
}

//	type IdentityStateInfo struct {
//		Date string
//	}
//
//	type IdentityInfo struct {
//		States map[IdentityState]IdentityStateInfo
//	}
//
//	type IdentityStateStorage struct {
//		mu     sync.RWMutex
//		states map[types.NodeID]*IdentityInfo
//	}
type IdentityStateEpoch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Epoch  uint32           `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	States []*IdentityState `protobuf:"bytes,2,rep,name=states,proto3" json:"states,omitempty"`
}

func (x *IdentityStateEpoch) Reset() {
	*x = IdentityStateEpoch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityStateEpoch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityStateEpoch) ProtoMessage() {}

func (x *IdentityStateEpoch) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityStateEpoch.ProtoReflect.Descriptor instead.
func (*IdentityStateEpoch) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_smeshing_identities_proto_rawDescGZIP(), []int{1}
}

func (x *IdentityStateEpoch) GetEpoch() uint32 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *IdentityStateEpoch) GetStates() []*IdentityState {
	if x != nil {
		return x.States
	}
	return nil
}

type IdentityState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State IdentityStatus         `protobuf:"varint,1,opt,name=state,proto3,enum=spacemesh.v2alpha1.IdentityStatus" json:"state,omitempty"`
	Time  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *IdentityState) Reset() {
	*x = IdentityState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityState) ProtoMessage() {}

func (x *IdentityState) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityState.ProtoReflect.Descriptor instead.
func (*IdentityState) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_smeshing_identities_proto_rawDescGZIP(), []int{2}
}

func (x *IdentityState) GetState() IdentityStatus {
	if x != nil {
		return x.State
	}
	return IdentityStatus_UNSPECIFIED
}

func (x *IdentityState) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmesherId []byte                `protobuf:"bytes,1,opt,name=smesher_id,json=smesherId,proto3" json:"smesher_id,omitempty"`
	Epochs    []*IdentityStateEpoch `protobuf:"bytes,2,rep,name=epochs,proto3" json:"epochs,omitempty"`
	States    []*IdentityState      `protobuf:"bytes,3,rep,name=states,proto3" json:"states,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_smeshing_identities_proto_rawDescGZIP(), []int{3}
}

func (x *Identity) GetSmesherId() []byte {
	if x != nil {
		return x.SmesherId
	}
	return nil
}

func (x *Identity) GetEpochs() []*IdentityStateEpoch {
	if x != nil {
		return x.Epochs
	}
	return nil
}

func (x *Identity) GetStates() []*IdentityState {
	if x != nil {
		return x.States
	}
	return nil
}

type IdentityStatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identities []*Identity `protobuf:"bytes,3,rep,name=identities,proto3" json:"identities,omitempty"`
}

func (x *IdentityStatesResponse) Reset() {
	*x = IdentityStatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityStatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityStatesResponse) ProtoMessage() {}

func (x *IdentityStatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityStatesResponse.ProtoReflect.Descriptor instead.
func (*IdentityStatesResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_smeshing_identities_proto_rawDescGZIP(), []int{4}
}

func (x *IdentityStatesResponse) GetIdentities() []*Identity {
	if x != nil {
		return x.Identities
	}
	return nil
}

var File_spacemesh_v2alpha1_smeshing_identities_proto protoreflect.FileDescriptor

var file_spacemesh_v2alpha1_smeshing_identities_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17,
	0x0a, 0x15, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x65, 0x0a, 0x12, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0x79,
	0x0a, 0x0d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x38, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x08, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6d, 0x65, 0x73, 0x68, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x6d, 0x65, 0x73,
	0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x06, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x52, 0x06, 0x65,
	0x70, 0x6f, 0x63, 0x68, 0x73, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x22, 0x56, 0x0a, 0x16, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2a, 0xd8, 0x02, 0x0a, 0x0e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x57, 0x41, 0x49, 0x54, 0x5f,
	0x46, 0x4f, 0x52, 0x5f, 0x41, 0x54, 0x58, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x28, 0x0a, 0x24, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x4f, 0x52,
	0x5f, 0x50, 0x4f, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x10, 0x06, 0x12, 0x18, 0x0a, 0x14, 0x50,
	0x4f, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x5f, 0x52, 0x45,
	0x41, 0x44, 0x59, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4f, 0x45, 0x54, 0x5f, 0x52, 0x45,
	0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45, 0x44, 0x10, 0x08, 0x12, 0x1b, 0x0a, 0x17, 0x57, 0x41,
	0x49, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x50, 0x4f, 0x45, 0x54, 0x5f, 0x52, 0x4f, 0x55, 0x4e,
	0x44, 0x5f, 0x45, 0x4e, 0x44, 0x10, 0x09, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x4f, 0x45, 0x54, 0x5f,
	0x50, 0x52, 0x4f, 0x4f, 0x46, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x44, 0x10, 0x0a,
	0x12, 0x15, 0x0a, 0x11, 0x50, 0x4f, 0x45, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x4f, 0x46, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x19, 0x0a, 0x15, 0x47, 0x45, 0x4e, 0x45, 0x52,
	0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x4f, 0x46,
	0x10, 0x0c, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x4f, 0x46,
	0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x0d, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x4f, 0x53, 0x54,
	0x5f, 0x50, 0x52, 0x4f, 0x4f, 0x46, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x0e, 0x12,
	0x0d, 0x0a, 0x09, 0x41, 0x54, 0x58, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x0f, 0x12, 0x13,
	0x0a, 0x0f, 0x41, 0x54, 0x58, 0x5f, 0x42, 0x52, 0x4f, 0x41, 0x44, 0x43, 0x41, 0x53, 0x54, 0x45,
	0x44, 0x10, 0x10, 0x32, 0x8a, 0x01, 0x0a, 0x19, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x61, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x29, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x1a, 0x0a, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x04, 0x12, 0x02, 0x56, 0x32,
	0x42, 0xe6, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x17, 0x53, 0x6d, 0x65,
	0x73, 0x68, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x12, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x12,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xe2, 0x02, 0x1e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x3a,
	0x3a, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_spacemesh_v2alpha1_smeshing_identities_proto_rawDescOnce sync.Once
	file_spacemesh_v2alpha1_smeshing_identities_proto_rawDescData = file_spacemesh_v2alpha1_smeshing_identities_proto_rawDesc
)

func file_spacemesh_v2alpha1_smeshing_identities_proto_rawDescGZIP() []byte {
	file_spacemesh_v2alpha1_smeshing_identities_proto_rawDescOnce.Do(func() {
		file_spacemesh_v2alpha1_smeshing_identities_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacemesh_v2alpha1_smeshing_identities_proto_rawDescData)
	})
	return file_spacemesh_v2alpha1_smeshing_identities_proto_rawDescData
}

var file_spacemesh_v2alpha1_smeshing_identities_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spacemesh_v2alpha1_smeshing_identities_proto_goTypes = []interface{}{
	(IdentityStatus)(0),            // 0: spacemesh.v2alpha1.IdentityStatus
	(*IdentityStatesRequest)(nil),  // 1: spacemesh.v2alpha1.IdentityStatesRequest
	(*IdentityStateEpoch)(nil),     // 2: spacemesh.v2alpha1.IdentityStateEpoch
	(*IdentityState)(nil),          // 3: spacemesh.v2alpha1.IdentityState
	(*Identity)(nil),               // 4: spacemesh.v2alpha1.Identity
	(*IdentityStatesResponse)(nil), // 5: spacemesh.v2alpha1.IdentityStatesResponse
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
}
var file_spacemesh_v2alpha1_smeshing_identities_proto_depIdxs = []int32{
	3, // 0: spacemesh.v2alpha1.IdentityStateEpoch.states:type_name -> spacemesh.v2alpha1.IdentityState
	0, // 1: spacemesh.v2alpha1.IdentityState.state:type_name -> spacemesh.v2alpha1.IdentityStatus
	6, // 2: spacemesh.v2alpha1.IdentityState.time:type_name -> google.protobuf.Timestamp
	2, // 3: spacemesh.v2alpha1.Identity.epochs:type_name -> spacemesh.v2alpha1.IdentityStateEpoch
	3, // 4: spacemesh.v2alpha1.Identity.states:type_name -> spacemesh.v2alpha1.IdentityState
	4, // 5: spacemesh.v2alpha1.IdentityStatesResponse.identities:type_name -> spacemesh.v2alpha1.Identity
	1, // 6: spacemesh.v2alpha1.SmeshingIdentitiesService.States:input_type -> spacemesh.v2alpha1.IdentityStatesRequest
	5, // 7: spacemesh.v2alpha1.SmeshingIdentitiesService.States:output_type -> spacemesh.v2alpha1.IdentityStatesResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_spacemesh_v2alpha1_smeshing_identities_proto_init() }
func file_spacemesh_v2alpha1_smeshing_identities_proto_init() {
	if File_spacemesh_v2alpha1_smeshing_identities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityStatesRequest); i {
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
		file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityStateEpoch); i {
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
		file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityState); i {
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
		file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
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
		file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityStatesResponse); i {
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
			RawDescriptor: file_spacemesh_v2alpha1_smeshing_identities_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spacemesh_v2alpha1_smeshing_identities_proto_goTypes,
		DependencyIndexes: file_spacemesh_v2alpha1_smeshing_identities_proto_depIdxs,
		EnumInfos:         file_spacemesh_v2alpha1_smeshing_identities_proto_enumTypes,
		MessageInfos:      file_spacemesh_v2alpha1_smeshing_identities_proto_msgTypes,
	}.Build()
	File_spacemesh_v2alpha1_smeshing_identities_proto = out.File
	file_spacemesh_v2alpha1_smeshing_identities_proto_rawDesc = nil
	file_spacemesh_v2alpha1_smeshing_identities_proto_goTypes = nil
	file_spacemesh_v2alpha1_smeshing_identities_proto_depIdxs = nil
}
