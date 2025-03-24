// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: spacemesh/v1/node.proto

package spacemeshv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_spacemesh_v1_node_proto protoreflect.FileDescriptor

const file_spacemesh_v1_node_proto_rawDesc = "" +
	"\n" +
	"\x17spacemesh/v1/node.proto\x12\fspacemesh.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1dspacemesh/v1/node_types.proto\x1a\x1bgoogle/api/visibility.proto2\xf8\x04\n" +
	"\vNodeService\x12W\n" +
	"\x04Echo\x12\x19.spacemesh.v1.EchoRequest\x1a\x1a.spacemesh.v1.EchoResponse\"\x18\x82\xd3\xe4\x93\x02\x12:\x01*\"\r/v1/node/echo\x12Z\n" +
	"\aVersion\x12\x16.google.protobuf.Empty\x1a\x1d.spacemesh.v1.VersionResponse\"\x18\x82\xd3\xe4\x93\x02\x12\"\x10/v1/node/version\x12T\n" +
	"\x05Build\x12\x16.google.protobuf.Empty\x1a\x1b.spacemesh.v1.BuildResponse\"\x16\x82\xd3\xe4\x93\x02\x10\"\x0e/v1/node/build\x12_\n" +
	"\x06Status\x12\x1b.spacemesh.v1.StatusRequest\x1a\x1c.spacemesh.v1.StatusResponse\"\x1a\x82\xd3\xe4\x93\x02\x14:\x01*\"\x0f/v1/node/status\x12B\n" +
	"\bNodeInfo\x12\x16.google.protobuf.Empty\x1a\x1e.spacemesh.v1.NodeInfoResponse\x12W\n" +
	"\fStatusStream\x12!.spacemesh.v1.StatusStreamRequest\x1a\".spacemesh.v1.StatusStreamResponse0\x01\x12T\n" +
	"\vErrorStream\x12 .spacemesh.v1.ErrorStreamRequest\x1a!.spacemesh.v1.ErrorStreamResponse0\x01\x1a\n" +
	"\xfa\xd2\xe4\x93\x02\x04\x12\x02V1B\xae\x01\n" +
	"\x10com.spacemesh.v1B\tNodeProtoP\x01Z>github.com/spacemeshos/api/release/go/spacemesh/v1;spacemeshv1\xa2\x02\x03SXX\xaa\x02\fSpacemesh.V1\xca\x02\fSpacemesh\\V1\xe2\x02\x18Spacemesh\\V1\\GPBMetadata\xea\x02\rSpacemesh::V1b\x06proto3"

var file_spacemesh_v1_node_proto_goTypes = []any{
	(*EchoRequest)(nil),          // 0: spacemesh.v1.EchoRequest
	(*emptypb.Empty)(nil),        // 1: google.protobuf.Empty
	(*StatusRequest)(nil),        // 2: spacemesh.v1.StatusRequest
	(*StatusStreamRequest)(nil),  // 3: spacemesh.v1.StatusStreamRequest
	(*ErrorStreamRequest)(nil),   // 4: spacemesh.v1.ErrorStreamRequest
	(*EchoResponse)(nil),         // 5: spacemesh.v1.EchoResponse
	(*VersionResponse)(nil),      // 6: spacemesh.v1.VersionResponse
	(*BuildResponse)(nil),        // 7: spacemesh.v1.BuildResponse
	(*StatusResponse)(nil),       // 8: spacemesh.v1.StatusResponse
	(*NodeInfoResponse)(nil),     // 9: spacemesh.v1.NodeInfoResponse
	(*StatusStreamResponse)(nil), // 10: spacemesh.v1.StatusStreamResponse
	(*ErrorStreamResponse)(nil),  // 11: spacemesh.v1.ErrorStreamResponse
}
var file_spacemesh_v1_node_proto_depIdxs = []int32{
	0,  // 0: spacemesh.v1.NodeService.Echo:input_type -> spacemesh.v1.EchoRequest
	1,  // 1: spacemesh.v1.NodeService.Version:input_type -> google.protobuf.Empty
	1,  // 2: spacemesh.v1.NodeService.Build:input_type -> google.protobuf.Empty
	2,  // 3: spacemesh.v1.NodeService.Status:input_type -> spacemesh.v1.StatusRequest
	1,  // 4: spacemesh.v1.NodeService.NodeInfo:input_type -> google.protobuf.Empty
	3,  // 5: spacemesh.v1.NodeService.StatusStream:input_type -> spacemesh.v1.StatusStreamRequest
	4,  // 6: spacemesh.v1.NodeService.ErrorStream:input_type -> spacemesh.v1.ErrorStreamRequest
	5,  // 7: spacemesh.v1.NodeService.Echo:output_type -> spacemesh.v1.EchoResponse
	6,  // 8: spacemesh.v1.NodeService.Version:output_type -> spacemesh.v1.VersionResponse
	7,  // 9: spacemesh.v1.NodeService.Build:output_type -> spacemesh.v1.BuildResponse
	8,  // 10: spacemesh.v1.NodeService.Status:output_type -> spacemesh.v1.StatusResponse
	9,  // 11: spacemesh.v1.NodeService.NodeInfo:output_type -> spacemesh.v1.NodeInfoResponse
	10, // 12: spacemesh.v1.NodeService.StatusStream:output_type -> spacemesh.v1.StatusStreamResponse
	11, // 13: spacemesh.v1.NodeService.ErrorStream:output_type -> spacemesh.v1.ErrorStreamResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_spacemesh_v1_node_proto_init() }
func file_spacemesh_v1_node_proto_init() {
	if File_spacemesh_v1_node_proto != nil {
		return
	}
	file_spacemesh_v1_node_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spacemesh_v1_node_proto_rawDesc), len(file_spacemesh_v1_node_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spacemesh_v1_node_proto_goTypes,
		DependencyIndexes: file_spacemesh_v1_node_proto_depIdxs,
	}.Build()
	File_spacemesh_v1_node_proto = out.File
	file_spacemesh_v1_node_proto_goTypes = nil
	file_spacemesh_v1_node_proto_depIdxs = nil
}
