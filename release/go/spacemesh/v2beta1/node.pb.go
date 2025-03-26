// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: spacemesh/v2beta1/node.proto

package spacemeshv2beta1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
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
	return file_spacemesh_v2beta1_node_proto_enumTypes[0].Descriptor()
}

func (NodeStatusResponse_SyncStatus) Type() protoreflect.EnumType {
	return &file_spacemesh_v2beta1_node_proto_enumTypes[0]
}

func (x NodeStatusResponse_SyncStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeStatusResponse_SyncStatus.Descriptor instead.
func (NodeStatusResponse_SyncStatus) EnumDescriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_node_proto_rawDescGZIP(), []int{0, 0}
}

type NodeStatusResponse struct {
	state          protoimpl.MessageState        `protogen:"open.v1"`
	ConnectedPeers uint64                        `protobuf:"varint,1,opt,name=connected_peers,json=connectedPeers,proto3" json:"connected_peers,omitempty"`                // number of connected neighbors
	Status         NodeStatusResponse_SyncStatus `protobuf:"varint,2,opt,name=status,proto3,enum=spacemesh.v2beta1.NodeStatusResponse_SyncStatus" json:"status,omitempty"` // node sync status
	LatestLayer    uint32                        `protobuf:"varint,3,opt,name=latest_layer,json=latestLayer,proto3" json:"latest_layer,omitempty"`                         // latest layer node has seen from blocks
	AppliedLayer   uint32                        `protobuf:"varint,4,opt,name=applied_layer,json=appliedLayer,proto3" json:"applied_layer,omitempty"`                      // last layer node has applied to the state
	ProcessedLayer uint32                        `protobuf:"varint,5,opt,name=processed_layer,json=processedLayer,proto3" json:"processed_layer,omitempty"`                // last layer whose votes have been processed
	CurrentLayer   uint32                        `protobuf:"varint,6,opt,name=current_layer,json=currentLayer,proto3" json:"current_layer,omitempty"`                      // current layer, based on clock time
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *NodeStatusResponse) Reset() {
	*x = NodeStatusResponse{}
	mi := &file_spacemesh_v2beta1_node_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatusResponse) ProtoMessage() {}

func (x *NodeStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_node_proto_msgTypes[0]
	if x != nil {
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
	return file_spacemesh_v2beta1_node_proto_rawDescGZIP(), []int{0}
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

func (x *NodeStatusResponse) GetLatestLayer() uint32 {
	if x != nil {
		return x.LatestLayer
	}
	return 0
}

func (x *NodeStatusResponse) GetAppliedLayer() uint32 {
	if x != nil {
		return x.AppliedLayer
	}
	return 0
}

func (x *NodeStatusResponse) GetProcessedLayer() uint32 {
	if x != nil {
		return x.ProcessedLayer
	}
	return 0
}

func (x *NodeStatusResponse) GetCurrentLayer() uint32 {
	if x != nil {
		return x.CurrentLayer
	}
	return 0
}

type NodeStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeStatusRequest) Reset() {
	*x = NodeStatusRequest{}
	mi := &file_spacemesh_v2beta1_node_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatusRequest) ProtoMessage() {}

func (x *NodeStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_node_proto_msgTypes[1]
	if x != nil {
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
	return file_spacemesh_v2beta1_node_proto_rawDescGZIP(), []int{1}
}

type NodeVersionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeVersionRequest) Reset() {
	*x = NodeVersionRequest{}
	mi := &file_spacemesh_v2beta1_node_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeVersionRequest) ProtoMessage() {}

func (x *NodeVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_node_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeVersionRequest.ProtoReflect.Descriptor instead.
func (*NodeVersionRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_node_proto_rawDescGZIP(), []int{2}
}

type NodeVersionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"` // semantic version of the node software
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeVersionResponse) Reset() {
	*x = NodeVersionResponse{}
	mi := &file_spacemesh_v2beta1_node_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeVersionResponse) ProtoMessage() {}

func (x *NodeVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_node_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeVersionResponse.ProtoReflect.Descriptor instead.
func (*NodeVersionResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_node_proto_rawDescGZIP(), []int{3}
}

func (x *NodeVersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type NodeBuildRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeBuildRequest) Reset() {
	*x = NodeBuildRequest{}
	mi := &file_spacemesh_v2beta1_node_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeBuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeBuildRequest) ProtoMessage() {}

func (x *NodeBuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_node_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeBuildRequest.ProtoReflect.Descriptor instead.
func (*NodeBuildRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_node_proto_rawDescGZIP(), []int{4}
}

type NodeBuildResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Build         string                 `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty"` // git commit hash of the build
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeBuildResponse) Reset() {
	*x = NodeBuildResponse{}
	mi := &file_spacemesh_v2beta1_node_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeBuildResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeBuildResponse) ProtoMessage() {}

func (x *NodeBuildResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_node_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeBuildResponse.ProtoReflect.Descriptor instead.
func (*NodeBuildResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_node_proto_rawDescGZIP(), []int{5}
}

func (x *NodeBuildResponse) GetBuild() string {
	if x != nil {
		return x.Build
	}
	return ""
}

var File_spacemesh_v2beta1_node_proto protoreflect.FileDescriptor

const file_spacemesh_v2beta1_node_proto_rawDesc = "" +
	"\n" +
	"\x1cspacemesh/v2beta1/node.proto\x12\x11spacemesh.v2beta1\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/api/visibility.proto\"\x92\x03\n" +
	"\x12NodeStatusResponse\x12'\n" +
	"\x0fconnected_peers\x18\x01 \x01(\x04R\x0econnectedPeers\x12H\n" +
	"\x06status\x18\x02 \x01(\x0e20.spacemesh.v2beta1.NodeStatusResponse.SyncStatusR\x06status\x12!\n" +
	"\flatest_layer\x18\x03 \x01(\rR\vlatestLayer\x12#\n" +
	"\rapplied_layer\x18\x04 \x01(\rR\fappliedLayer\x12'\n" +
	"\x0fprocessed_layer\x18\x05 \x01(\rR\x0eprocessedLayer\x12#\n" +
	"\rcurrent_layer\x18\x06 \x01(\rR\fcurrentLayer\"s\n" +
	"\n" +
	"SyncStatus\x12\x1b\n" +
	"\x17SYNC_STATUS_UNSPECIFIED\x10\x00\x12\x17\n" +
	"\x13SYNC_STATUS_OFFLINE\x10\x01\x12\x17\n" +
	"\x13SYNC_STATUS_SYNCING\x10\x02\x12\x16\n" +
	"\x12SYNC_STATUS_SYNCED\x10\x03\"\x13\n" +
	"\x11NodeStatusRequest\"\x14\n" +
	"\x12NodeVersionRequest\"/\n" +
	"\x13NodeVersionResponse\x12\x18\n" +
	"\aversion\x18\x01 \x01(\tR\aversion\"\x12\n" +
	"\x10NodeBuildRequest\")\n" +
	"\x11NodeBuildResponse\x12\x14\n" +
	"\x05build\x18\x01 \x01(\tR\x05build2\xb3\x03\n" +
	"\vNodeService\x12\x84\x01\n" +
	"\x06Status\x12$.spacemesh.v2beta1.NodeStatusRequest\x1a%.spacemesh.v2beta1.NodeStatusResponse\"-\x82\xd3\xe4\x93\x02'\x12%/spacemesh.v2beta1.NodeService/Status\x12\x88\x01\n" +
	"\aVersion\x12%.spacemesh.v2beta1.NodeVersionRequest\x1a&.spacemesh.v2beta1.NodeVersionResponse\".\x82\xd3\xe4\x93\x02(\x12&/spacemesh.v2beta1.NodeService/Version\x12\x80\x01\n" +
	"\x05Build\x12#.spacemesh.v2beta1.NodeBuildRequest\x1a$.spacemesh.v2beta1.NodeBuildResponse\",\x82\xd3\xe4\x93\x02&\x12$/spacemesh.v2beta1.NodeService/Build\x1a\x0f\xfa\xd2\xe4\x93\x02\t\x12\av2beta1B\xd1\x01\n" +
	"\x15com.spacemesh.v2beta1B\tNodeProtoP\x01ZHgithub.com/spacemeshos/api/release/go/spacemesh/v2beta1;spacemeshv2beta1\xa2\x02\x03SXX\xaa\x02\x11Spacemesh.V2beta1\xca\x02\x11Spacemesh\\V2beta1\xe2\x02\x1dSpacemesh\\V2beta1\\GPBMetadata\xea\x02\x12Spacemesh::V2beta1b\x06proto3"

var (
	file_spacemesh_v2beta1_node_proto_rawDescOnce sync.Once
	file_spacemesh_v2beta1_node_proto_rawDescData []byte
)

func file_spacemesh_v2beta1_node_proto_rawDescGZIP() []byte {
	file_spacemesh_v2beta1_node_proto_rawDescOnce.Do(func() {
		file_spacemesh_v2beta1_node_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spacemesh_v2beta1_node_proto_rawDesc), len(file_spacemesh_v2beta1_node_proto_rawDesc)))
	})
	return file_spacemesh_v2beta1_node_proto_rawDescData
}

var file_spacemesh_v2beta1_node_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spacemesh_v2beta1_node_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_spacemesh_v2beta1_node_proto_goTypes = []any{
	(NodeStatusResponse_SyncStatus)(0), // 0: spacemesh.v2beta1.NodeStatusResponse.SyncStatus
	(*NodeStatusResponse)(nil),         // 1: spacemesh.v2beta1.NodeStatusResponse
	(*NodeStatusRequest)(nil),          // 2: spacemesh.v2beta1.NodeStatusRequest
	(*NodeVersionRequest)(nil),         // 3: spacemesh.v2beta1.NodeVersionRequest
	(*NodeVersionResponse)(nil),        // 4: spacemesh.v2beta1.NodeVersionResponse
	(*NodeBuildRequest)(nil),           // 5: spacemesh.v2beta1.NodeBuildRequest
	(*NodeBuildResponse)(nil),          // 6: spacemesh.v2beta1.NodeBuildResponse
}
var file_spacemesh_v2beta1_node_proto_depIdxs = []int32{
	0, // 0: spacemesh.v2beta1.NodeStatusResponse.status:type_name -> spacemesh.v2beta1.NodeStatusResponse.SyncStatus
	2, // 1: spacemesh.v2beta1.NodeService.Status:input_type -> spacemesh.v2beta1.NodeStatusRequest
	3, // 2: spacemesh.v2beta1.NodeService.Version:input_type -> spacemesh.v2beta1.NodeVersionRequest
	5, // 3: spacemesh.v2beta1.NodeService.Build:input_type -> spacemesh.v2beta1.NodeBuildRequest
	1, // 4: spacemesh.v2beta1.NodeService.Status:output_type -> spacemesh.v2beta1.NodeStatusResponse
	4, // 5: spacemesh.v2beta1.NodeService.Version:output_type -> spacemesh.v2beta1.NodeVersionResponse
	6, // 6: spacemesh.v2beta1.NodeService.Build:output_type -> spacemesh.v2beta1.NodeBuildResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spacemesh_v2beta1_node_proto_init() }
func file_spacemesh_v2beta1_node_proto_init() {
	if File_spacemesh_v2beta1_node_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spacemesh_v2beta1_node_proto_rawDesc), len(file_spacemesh_v2beta1_node_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spacemesh_v2beta1_node_proto_goTypes,
		DependencyIndexes: file_spacemesh_v2beta1_node_proto_depIdxs,
		EnumInfos:         file_spacemesh_v2beta1_node_proto_enumTypes,
		MessageInfos:      file_spacemesh_v2beta1_node_proto_msgTypes,
	}.Build()
	File_spacemesh_v2beta1_node_proto = out.File
	file_spacemesh_v2beta1_node_proto_goTypes = nil
	file_spacemesh_v2beta1_node_proto_depIdxs = nil
}
