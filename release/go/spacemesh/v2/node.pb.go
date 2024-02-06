// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: spacemesh/v2/node.proto

package spacemeshv2

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

type NodeStatusResponse_SyncStatus int32

const (
	NodeStatusResponse_SYNC_STATUS_UNSPECIFIED NodeStatusResponse_SyncStatus = 0
	NodeStatusResponse_SYNC_STATUS_OFFLINE     NodeStatusResponse_SyncStatus = 1
	NodeStatusResponse_SYNC_STATUS_SYNCING     NodeStatusResponse_SyncStatus = 2
	NodeStatusResponse_SYNC_STATUS_SYNCED      NodeStatusResponse_SyncStatus = 3
)

// Enum value maps for NodeStatusResponse_SyncStatus.
var (
	NodeStatusResponse_SyncStatus_name = map[int32]string{
		0: "SYNC_STATUS_UNSPECIFIED",
		1: "SYNC_STATUS_OFFLINE",
		2: "SYNC_STATUS_SYNCING",
		3: "SYNC_STATUS_SYNCED",
	}
	NodeStatusResponse_SyncStatus_value = map[string]int32{
		"SYNC_STATUS_UNSPECIFIED": 0,
		"SYNC_STATUS_OFFLINE":     1,
		"SYNC_STATUS_SYNCING":     2,
		"SYNC_STATUS_SYNCED":      3,
	}
)

func (x NodeStatusResponse_SyncStatus) Enum() *NodeStatusResponse_SyncStatus {
	p := new(NodeStatusResponse_SyncStatus)
	*p = x
	return p
}

func (x NodeStatusResponse_SyncStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeStatusResponse_SyncStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_spacemesh_v2_node_proto_enumTypes[0].Descriptor()
}

func (NodeStatusResponse_SyncStatus) Type() protoreflect.EnumType {
	return &file_spacemesh_v2_node_proto_enumTypes[0]
}

func (x NodeStatusResponse_SyncStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeStatusResponse_SyncStatus.Descriptor instead.
func (NodeStatusResponse_SyncStatus) EnumDescriptor() ([]byte, []int) {
	return file_spacemesh_v2_node_proto_rawDescGZIP(), []int{1, 0}
}

type NodeStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NodeStatusRequest) Reset() {
	*x = NodeStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatusRequest) ProtoMessage() {}

func (x *NodeStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStatusRequest.ProtoReflect.Descriptor instead.
func (*NodeStatusRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2_node_proto_rawDescGZIP(), []int{0}
}

type NodeStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectedPeers uint64                        `protobuf:"varint,1,opt,name=connected_peers,json=connectedPeers,proto3" json:"connected_peers,omitempty"` // number of connected neighbors
	Status         NodeStatusResponse_SyncStatus `protobuf:"varint,2,opt,name=status,proto3,enum=spacemesh.v2.NodeStatusResponse_SyncStatus" json:"status,omitempty"`
	SyncedLayer    uint32                        `protobuf:"varint,3,opt,name=synced_layer,json=syncedLayer,proto3" json:"synced_layer,omitempty"`       // the last layer node has synced
	VerifiedLayer  uint32                        `protobuf:"varint,4,opt,name=verified_layer,json=verifiedLayer,proto3" json:"verified_layer,omitempty"` // the last layer node has verified
	HeadLayer      uint32                        `protobuf:"varint,5,opt,name=head_layer,json=headLayer,proto3" json:"head_layer,omitempty"`             // head layer is the tip
	HeadBlock      uint32                        `protobuf:"varint,6,opt,name=head_block,json=headBlock,proto3" json:"head_block,omitempty"`
}

func (x *NodeStatusResponse) Reset() {
	*x = NodeStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatusResponse) ProtoMessage() {}

func (x *NodeStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStatusResponse.ProtoReflect.Descriptor instead.
func (*NodeStatusResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2_node_proto_rawDescGZIP(), []int{1}
}

func (x *NodeStatusResponse) GetConnectedPeers() uint64 {
	if x != nil {
		return x.ConnectedPeers
	}
	return 0
}

func (x *NodeStatusResponse) GetStatus() NodeStatusResponse_SyncStatus {
	if x != nil {
		return x.Status
	}
	return NodeStatusResponse_SYNC_STATUS_UNSPECIFIED
}

func (x *NodeStatusResponse) GetSyncedLayer() uint32 {
	if x != nil {
		return x.SyncedLayer
	}
	return 0
}

func (x *NodeStatusResponse) GetVerifiedLayer() uint32 {
	if x != nil {
		return x.VerifiedLayer
	}
	return 0
}

func (x *NodeStatusResponse) GetHeadLayer() uint32 {
	if x != nil {
		return x.HeadLayer
	}
	return 0
}

func (x *NodeStatusResponse) GetHeadBlock() uint32 {
	if x != nil {
		return x.HeadBlock
	}
	return 0
}

var File_spacemesh_v2_node_proto protoreflect.FileDescriptor

var file_spacemesh_v2_node_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x22, 0x13, 0x0a, 0x11, 0x4e, 0x6f, 0x64, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xff, 0x02, 0x0a,
	0x12, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x5f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x43, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x4c,
	0x61, 0x79, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x68,
	0x65, 0x61, 0x64, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x68, 0x65, 0x61, 0x64, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x65,
	0x61, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x68, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x73, 0x0a, 0x0a, 0x53, 0x79, 0x6e,
	0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x59, 0x4e, 0x43, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x17, 0x0a,
	0x13, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x59, 0x4e,
	0x43, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x45, 0x44, 0x10, 0x03, 0x32, 0x5a,
	0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xae, 0x01, 0x0a, 0x10, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x42,
	0x09, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32,
	0x3b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x53,
	0x58, 0x58, 0xaa, 0x02, 0x0c, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x56,
	0x32, 0xca, 0x02, 0x0c, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32,
	0xe2, 0x02, 0x18, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_spacemesh_v2_node_proto_rawDescOnce sync.Once
	file_spacemesh_v2_node_proto_rawDescData = file_spacemesh_v2_node_proto_rawDesc
)

func file_spacemesh_v2_node_proto_rawDescGZIP() []byte {
	file_spacemesh_v2_node_proto_rawDescOnce.Do(func() {
		file_spacemesh_v2_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacemesh_v2_node_proto_rawDescData)
	})
	return file_spacemesh_v2_node_proto_rawDescData
}

var file_spacemesh_v2_node_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spacemesh_v2_node_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_spacemesh_v2_node_proto_goTypes = []interface{}{
	(NodeStatusResponse_SyncStatus)(0), // 0: spacemesh.v2.NodeStatusResponse.SyncStatus
	(*NodeStatusRequest)(nil),          // 1: spacemesh.v2.NodeStatusRequest
	(*NodeStatusResponse)(nil),         // 2: spacemesh.v2.NodeStatusResponse
}
var file_spacemesh_v2_node_proto_depIdxs = []int32{
	0, // 0: spacemesh.v2.NodeStatusResponse.status:type_name -> spacemesh.v2.NodeStatusResponse.SyncStatus
	1, // 1: spacemesh.v2.NodeService.Status:input_type -> spacemesh.v2.NodeStatusRequest
	2, // 2: spacemesh.v2.NodeService.Status:output_type -> spacemesh.v2.NodeStatusResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spacemesh_v2_node_proto_init() }
func file_spacemesh_v2_node_proto_init() {
	if File_spacemesh_v2_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spacemesh_v2_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStatusRequest); i {
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
		file_spacemesh_v2_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStatusResponse); i {
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
			RawDescriptor: file_spacemesh_v2_node_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spacemesh_v2_node_proto_goTypes,
		DependencyIndexes: file_spacemesh_v2_node_proto_depIdxs,
		EnumInfos:         file_spacemesh_v2_node_proto_enumTypes,
		MessageInfos:      file_spacemesh_v2_node_proto_msgTypes,
	}.Build()
	File_spacemesh_v2_node_proto = out.File
	file_spacemesh_v2_node_proto_rawDesc = nil
	file_spacemesh_v2_node_proto_goTypes = nil
	file_spacemesh_v2_node_proto_depIdxs = nil
}
