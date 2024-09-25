// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: spacemesh/v2alpha1/malfeasance.proto

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

type MalfeasanceProof_MalfeasanceDomain int32

const (
	MalfeasanceProof_DOMAIN_UNSPECIFIED MalfeasanceProof_MalfeasanceDomain = 0 // for legacy proofs
)

// Enum value maps for MalfeasanceProof_MalfeasanceDomain.
var (
	MalfeasanceProof_MalfeasanceDomain_name = map[int32]string{
		0: "DOMAIN_UNSPECIFIED",
	}
	MalfeasanceProof_MalfeasanceDomain_value = map[string]int32{
		"DOMAIN_UNSPECIFIED": 0,
	}
)

func (x MalfeasanceProof_MalfeasanceDomain) Enum() *MalfeasanceProof_MalfeasanceDomain {
	p := new(MalfeasanceProof_MalfeasanceDomain)
	*p = x
	return p
}

func (x MalfeasanceProof_MalfeasanceDomain) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MalfeasanceProof_MalfeasanceDomain) Descriptor() protoreflect.EnumDescriptor {
	return file_spacemesh_v2alpha1_malfeasance_proto_enumTypes[0].Descriptor()
}

func (MalfeasanceProof_MalfeasanceDomain) Type() protoreflect.EnumType {
	return &file_spacemesh_v2alpha1_malfeasance_proto_enumTypes[0]
}

func (x MalfeasanceProof_MalfeasanceDomain) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MalfeasanceProof_MalfeasanceDomain.Descriptor instead.
func (MalfeasanceProof_MalfeasanceDomain) EnumDescriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_malfeasance_proto_rawDescGZIP(), []int{0, 0}
}

type MalfeasanceProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Smesher []byte                             `protobuf:"bytes,1,opt,name=smesher,proto3" json:"smesher,omitempty"`
	Domain  MalfeasanceProof_MalfeasanceDomain `protobuf:"varint,2,opt,name=domain,proto3,enum=spacemesh.v2alpha1.MalfeasanceProof_MalfeasanceDomain" json:"domain,omitempty"`
	// type of the malfeasance proof, depends on domain
	//
	// for legacy proofs the types are
	//
	//	1 - Double publish of ATX
	//	2 - Multiple ballots for a layer by same smesher
	//	3 - Hare Equivocation (currently unused)
	//	4 - ATX with invalid PoST proof publised
	//	5 - ATX referencing an invalid previous ATX published
	Type uint32 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	// Properties of the Malfeasance proof, different for every type of proof
	Properties map[string]string `protobuf:"bytes,4,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MalfeasanceProof) Reset() {
	*x = MalfeasanceProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_malfeasance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MalfeasanceProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MalfeasanceProof) ProtoMessage() {}

func (x *MalfeasanceProof) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_malfeasance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MalfeasanceProof.ProtoReflect.Descriptor instead.
func (*MalfeasanceProof) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_malfeasance_proto_rawDescGZIP(), []int{0}
}

func (x *MalfeasanceProof) GetSmesher() []byte {
	if x != nil {
		return x.Smesher
	}
	return nil
}

func (x *MalfeasanceProof) GetDomain() MalfeasanceProof_MalfeasanceDomain {
	if x != nil {
		return x.Domain
	}
	return MalfeasanceProof_DOMAIN_UNSPECIFIED
}

func (x *MalfeasanceProof) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *MalfeasanceProof) GetProperties() map[string]string {
	if x != nil {
		return x.Properties
	}
	return nil
}

type MalfeasanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmesherId [][]byte `protobuf:"bytes,1,rep,name=smesher_id,json=smesherId,proto3" json:"smesher_id,omitempty"` // list of smesher ids to fetch
	Offset    uint64   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`                       // adjusts the starting point for data
	Limit     uint64   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`                         // specifies max number of items to fetch
}

func (x *MalfeasanceRequest) Reset() {
	*x = MalfeasanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_malfeasance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MalfeasanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MalfeasanceRequest) ProtoMessage() {}

func (x *MalfeasanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_malfeasance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MalfeasanceRequest.ProtoReflect.Descriptor instead.
func (*MalfeasanceRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_malfeasance_proto_rawDescGZIP(), []int{1}
}

func (x *MalfeasanceRequest) GetSmesherId() [][]byte {
	if x != nil {
		return x.SmesherId
	}
	return nil
}

func (x *MalfeasanceRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *MalfeasanceRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type MalfeasanceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Malfeasances []*MalfeasanceProof `protobuf:"bytes,1,rep,name=malfeasances,proto3" json:"malfeasances,omitempty"`
}

func (x *MalfeasanceList) Reset() {
	*x = MalfeasanceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_malfeasance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MalfeasanceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MalfeasanceList) ProtoMessage() {}

func (x *MalfeasanceList) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_malfeasance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MalfeasanceList.ProtoReflect.Descriptor instead.
func (*MalfeasanceList) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_malfeasance_proto_rawDescGZIP(), []int{2}
}

func (x *MalfeasanceList) GetMalfeasances() []*MalfeasanceProof {
	if x != nil {
		return x.Malfeasances
	}
	return nil
}

type MalfeasanceStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmesherId [][]byte `protobuf:"bytes,1,rep,name=smesher_id,json=smesherId,proto3" json:"smesher_id,omitempty"` // list of smesher ids to watch
	Watch     bool     `protobuf:"varint,2,opt,name=watch,proto3" json:"watch,omitempty"`
}

func (x *MalfeasanceStreamRequest) Reset() {
	*x = MalfeasanceStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_malfeasance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MalfeasanceStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MalfeasanceStreamRequest) ProtoMessage() {}

func (x *MalfeasanceStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_malfeasance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MalfeasanceStreamRequest.ProtoReflect.Descriptor instead.
func (*MalfeasanceStreamRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_malfeasance_proto_rawDescGZIP(), []int{3}
}

func (x *MalfeasanceStreamRequest) GetSmesherId() [][]byte {
	if x != nil {
		return x.SmesherId
	}
	return nil
}

func (x *MalfeasanceStreamRequest) GetWatch() bool {
	if x != nil {
		return x.Watch
	}
	return false
}

var File_spacemesh_v2alpha1_malfeasance_proto protoreflect.FileDescriptor

var file_spacemesh_v2alpha1_malfeasance_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x02, 0x0a, 0x10, 0x4d, 0x61, 0x6c, 0x66,
	0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73,
	0x6d, 0x65, 0x73, 0x68, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x61, 0x6c, 0x66,
	0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x2e, 0x4d, 0x61, 0x6c,
	0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x54, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x2b, 0x0a, 0x11, 0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x22, 0x61, 0x0a, 0x12,
	0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x5b, 0x0a, 0x0f, 0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x48, 0x0a, 0x0c, 0x6d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x61,
	0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x0c,
	0x6d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x4f, 0x0a, 0x18,
	0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6d, 0x65, 0x73,
	0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x6d,
	0x65, 0x73, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x77, 0x61, 0x74, 0x63, 0x68, 0x32, 0x75, 0x0a,
	0x12, 0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73,
	0x61, 0x6e, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x0a, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x04,
	0x12, 0x02, 0x56, 0x32, 0x32, 0x8c, 0x01, 0x0a, 0x18, 0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73,
	0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5e, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2c, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d,
	0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x30,
	0x01, 0x1a, 0x10, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x4e, 0x41, 0x4c, 0x42, 0xdf, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x10,
	0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02,
	0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x12, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x12, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02,
	0x1e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x13, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x56, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spacemesh_v2alpha1_malfeasance_proto_rawDescOnce sync.Once
	file_spacemesh_v2alpha1_malfeasance_proto_rawDescData = file_spacemesh_v2alpha1_malfeasance_proto_rawDesc
)

func file_spacemesh_v2alpha1_malfeasance_proto_rawDescGZIP() []byte {
	file_spacemesh_v2alpha1_malfeasance_proto_rawDescOnce.Do(func() {
		file_spacemesh_v2alpha1_malfeasance_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacemesh_v2alpha1_malfeasance_proto_rawDescData)
	})
	return file_spacemesh_v2alpha1_malfeasance_proto_rawDescData
}

var file_spacemesh_v2alpha1_malfeasance_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spacemesh_v2alpha1_malfeasance_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spacemesh_v2alpha1_malfeasance_proto_goTypes = []interface{}{
	(MalfeasanceProof_MalfeasanceDomain)(0), // 0: spacemesh.v2alpha1.MalfeasanceProof.MalfeasanceDomain
	(*MalfeasanceProof)(nil),                // 1: spacemesh.v2alpha1.MalfeasanceProof
	(*MalfeasanceRequest)(nil),              // 2: spacemesh.v2alpha1.MalfeasanceRequest
	(*MalfeasanceList)(nil),                 // 3: spacemesh.v2alpha1.MalfeasanceList
	(*MalfeasanceStreamRequest)(nil),        // 4: spacemesh.v2alpha1.MalfeasanceStreamRequest
	nil,                                     // 5: spacemesh.v2alpha1.MalfeasanceProof.PropertiesEntry
}
var file_spacemesh_v2alpha1_malfeasance_proto_depIdxs = []int32{
	0, // 0: spacemesh.v2alpha1.MalfeasanceProof.domain:type_name -> spacemesh.v2alpha1.MalfeasanceProof.MalfeasanceDomain
	5, // 1: spacemesh.v2alpha1.MalfeasanceProof.properties:type_name -> spacemesh.v2alpha1.MalfeasanceProof.PropertiesEntry
	1, // 2: spacemesh.v2alpha1.MalfeasanceList.malfeasances:type_name -> spacemesh.v2alpha1.MalfeasanceProof
	2, // 3: spacemesh.v2alpha1.MalfeasanceService.List:input_type -> spacemesh.v2alpha1.MalfeasanceRequest
	4, // 4: spacemesh.v2alpha1.MalfeasanceStreamService.Stream:input_type -> spacemesh.v2alpha1.MalfeasanceStreamRequest
	3, // 5: spacemesh.v2alpha1.MalfeasanceService.List:output_type -> spacemesh.v2alpha1.MalfeasanceList
	1, // 6: spacemesh.v2alpha1.MalfeasanceStreamService.Stream:output_type -> spacemesh.v2alpha1.MalfeasanceProof
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_spacemesh_v2alpha1_malfeasance_proto_init() }
func file_spacemesh_v2alpha1_malfeasance_proto_init() {
	if File_spacemesh_v2alpha1_malfeasance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spacemesh_v2alpha1_malfeasance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MalfeasanceProof); i {
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
		file_spacemesh_v2alpha1_malfeasance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MalfeasanceRequest); i {
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
		file_spacemesh_v2alpha1_malfeasance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MalfeasanceList); i {
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
		file_spacemesh_v2alpha1_malfeasance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MalfeasanceStreamRequest); i {
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
			RawDescriptor: file_spacemesh_v2alpha1_malfeasance_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_spacemesh_v2alpha1_malfeasance_proto_goTypes,
		DependencyIndexes: file_spacemesh_v2alpha1_malfeasance_proto_depIdxs,
		EnumInfos:         file_spacemesh_v2alpha1_malfeasance_proto_enumTypes,
		MessageInfos:      file_spacemesh_v2alpha1_malfeasance_proto_msgTypes,
	}.Build()
	File_spacemesh_v2alpha1_malfeasance_proto = out.File
	file_spacemesh_v2alpha1_malfeasance_proto_rawDesc = nil
	file_spacemesh_v2alpha1_malfeasance_proto_goTypes = nil
	file_spacemesh_v2alpha1_malfeasance_proto_depIdxs = nil
}