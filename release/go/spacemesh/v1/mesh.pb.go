// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: spacemesh/v1/mesh.proto

package spacemeshv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_spacemesh_v1_mesh_proto protoreflect.FileDescriptor

var file_spacemesh_v1_mesh_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x85, 0x0d, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x73, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a,
	0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73,
	0x69, 0x73, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x77, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x77, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12,
	0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01,
	0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x6b, 0x0a, 0x09, 0x47, 0x65, 0x6e, 0x65,
	0x73, 0x69, 0x73, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01,
	0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x73, 0x69, 0x73, 0x69, 0x64, 0x12, 0x7f, 0x0a, 0x0e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x4e, 0x75,
	0x6d, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x23, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x4e, 0x75, 0x6d, 0x4c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x70, 0x6f, 0x63,
	0x68, 0x4e, 0x75, 0x6d, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x6e, 0x75, 0x6d,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x7b, 0x0a, 0x0d, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0xa7, 0x01, 0x0a, 0x18, 0x4d, 0x61, 0x78, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x12, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x61, 0x78, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50,
	0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x61, 0x78, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x65,
	0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x22, 0x21, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x68, 0x2f, 0x6d, 0x61, 0x78, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x97, 0x01,
	0x0a, 0x14, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x68, 0x44, 0x61, 0x74,
	0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x6d, 0x65, 0x73, 0x68, 0x64, 0x61,
	0x74, 0x61, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x73, 0x0a, 0x0b, 0x4c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x72, 0x0a, 0x15,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x68,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x12, 0x54, 0x0a, 0x0b, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x20, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x54, 0x0a, 0x0b, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x57, 0x0a, 0x10,
	0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x11, 0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73,
	0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x26, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61,
	0x73, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x61, 0x6c, 0x66, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x1a, 0x0a, 0xfa,
	0xd2, 0xe4, 0x93, 0x02, 0x04, 0x12, 0x02, 0x56, 0x31, 0x42, 0xae, 0x01, 0x0a, 0x10, 0x63, 0x6f,
	0x6d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x09,
	0x4d, 0x65, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f,
	0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x3b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58,
	0x58, 0xaa, 0x02, 0x0c, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0c, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x18, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_spacemesh_v1_mesh_proto_goTypes = []interface{}{
	(*GenesisTimeRequest)(nil),               // 0: spacemesh.v1.GenesisTimeRequest
	(*CurrentLayerRequest)(nil),              // 1: spacemesh.v1.CurrentLayerRequest
	(*CurrentEpochRequest)(nil),              // 2: spacemesh.v1.CurrentEpochRequest
	(*GenesisIDRequest)(nil),                 // 3: spacemesh.v1.GenesisIDRequest
	(*EpochNumLayersRequest)(nil),            // 4: spacemesh.v1.EpochNumLayersRequest
	(*LayerDurationRequest)(nil),             // 5: spacemesh.v1.LayerDurationRequest
	(*MaxTransactionsPerSecondRequest)(nil),  // 6: spacemesh.v1.MaxTransactionsPerSecondRequest
	(*AccountMeshDataQueryRequest)(nil),      // 7: spacemesh.v1.AccountMeshDataQueryRequest
	(*LayersQueryRequest)(nil),               // 8: spacemesh.v1.LayersQueryRequest
	(*AccountMeshDataStreamRequest)(nil),     // 9: spacemesh.v1.AccountMeshDataStreamRequest
	(*LayerStreamRequest)(nil),               // 10: spacemesh.v1.LayerStreamRequest
	(*EpochStreamRequest)(nil),               // 11: spacemesh.v1.EpochStreamRequest
	(*MalfeasanceRequest)(nil),               // 12: spacemesh.v1.MalfeasanceRequest
	(*MalfeasanceStreamRequest)(nil),         // 13: spacemesh.v1.MalfeasanceStreamRequest
	(*GenesisTimeResponse)(nil),              // 14: spacemesh.v1.GenesisTimeResponse
	(*CurrentLayerResponse)(nil),             // 15: spacemesh.v1.CurrentLayerResponse
	(*CurrentEpochResponse)(nil),             // 16: spacemesh.v1.CurrentEpochResponse
	(*GenesisIDResponse)(nil),                // 17: spacemesh.v1.GenesisIDResponse
	(*EpochNumLayersResponse)(nil),           // 18: spacemesh.v1.EpochNumLayersResponse
	(*LayerDurationResponse)(nil),            // 19: spacemesh.v1.LayerDurationResponse
	(*MaxTransactionsPerSecondResponse)(nil), // 20: spacemesh.v1.MaxTransactionsPerSecondResponse
	(*AccountMeshDataQueryResponse)(nil),     // 21: spacemesh.v1.AccountMeshDataQueryResponse
	(*LayersQueryResponse)(nil),              // 22: spacemesh.v1.LayersQueryResponse
	(*AccountMeshDataStreamResponse)(nil),    // 23: spacemesh.v1.AccountMeshDataStreamResponse
	(*LayerStreamResponse)(nil),              // 24: spacemesh.v1.LayerStreamResponse
	(*EpochStreamResponse)(nil),              // 25: spacemesh.v1.EpochStreamResponse
	(*MalfeasanceResponse)(nil),              // 26: spacemesh.v1.MalfeasanceResponse
	(*MalfeasanceStreamResponse)(nil),        // 27: spacemesh.v1.MalfeasanceStreamResponse
}
var file_spacemesh_v1_mesh_proto_depIdxs = []int32{
	0,  // 0: spacemesh.v1.MeshService.GenesisTime:input_type -> spacemesh.v1.GenesisTimeRequest
	1,  // 1: spacemesh.v1.MeshService.CurrentLayer:input_type -> spacemesh.v1.CurrentLayerRequest
	2,  // 2: spacemesh.v1.MeshService.CurrentEpoch:input_type -> spacemesh.v1.CurrentEpochRequest
	3,  // 3: spacemesh.v1.MeshService.GenesisID:input_type -> spacemesh.v1.GenesisIDRequest
	4,  // 4: spacemesh.v1.MeshService.EpochNumLayers:input_type -> spacemesh.v1.EpochNumLayersRequest
	5,  // 5: spacemesh.v1.MeshService.LayerDuration:input_type -> spacemesh.v1.LayerDurationRequest
	6,  // 6: spacemesh.v1.MeshService.MaxTransactionsPerSecond:input_type -> spacemesh.v1.MaxTransactionsPerSecondRequest
	7,  // 7: spacemesh.v1.MeshService.AccountMeshDataQuery:input_type -> spacemesh.v1.AccountMeshDataQueryRequest
	8,  // 8: spacemesh.v1.MeshService.LayersQuery:input_type -> spacemesh.v1.LayersQueryRequest
	9,  // 9: spacemesh.v1.MeshService.AccountMeshDataStream:input_type -> spacemesh.v1.AccountMeshDataStreamRequest
	10, // 10: spacemesh.v1.MeshService.LayerStream:input_type -> spacemesh.v1.LayerStreamRequest
	11, // 11: spacemesh.v1.MeshService.EpochStream:input_type -> spacemesh.v1.EpochStreamRequest
	12, // 12: spacemesh.v1.MeshService.MalfeasanceQuery:input_type -> spacemesh.v1.MalfeasanceRequest
	13, // 13: spacemesh.v1.MeshService.MalfeasanceStream:input_type -> spacemesh.v1.MalfeasanceStreamRequest
	14, // 14: spacemesh.v1.MeshService.GenesisTime:output_type -> spacemesh.v1.GenesisTimeResponse
	15, // 15: spacemesh.v1.MeshService.CurrentLayer:output_type -> spacemesh.v1.CurrentLayerResponse
	16, // 16: spacemesh.v1.MeshService.CurrentEpoch:output_type -> spacemesh.v1.CurrentEpochResponse
	17, // 17: spacemesh.v1.MeshService.GenesisID:output_type -> spacemesh.v1.GenesisIDResponse
	18, // 18: spacemesh.v1.MeshService.EpochNumLayers:output_type -> spacemesh.v1.EpochNumLayersResponse
	19, // 19: spacemesh.v1.MeshService.LayerDuration:output_type -> spacemesh.v1.LayerDurationResponse
	20, // 20: spacemesh.v1.MeshService.MaxTransactionsPerSecond:output_type -> spacemesh.v1.MaxTransactionsPerSecondResponse
	21, // 21: spacemesh.v1.MeshService.AccountMeshDataQuery:output_type -> spacemesh.v1.AccountMeshDataQueryResponse
	22, // 22: spacemesh.v1.MeshService.LayersQuery:output_type -> spacemesh.v1.LayersQueryResponse
	23, // 23: spacemesh.v1.MeshService.AccountMeshDataStream:output_type -> spacemesh.v1.AccountMeshDataStreamResponse
	24, // 24: spacemesh.v1.MeshService.LayerStream:output_type -> spacemesh.v1.LayerStreamResponse
	25, // 25: spacemesh.v1.MeshService.EpochStream:output_type -> spacemesh.v1.EpochStreamResponse
	26, // 26: spacemesh.v1.MeshService.MalfeasanceQuery:output_type -> spacemesh.v1.MalfeasanceResponse
	27, // 27: spacemesh.v1.MeshService.MalfeasanceStream:output_type -> spacemesh.v1.MalfeasanceStreamResponse
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_spacemesh_v1_mesh_proto_init() }
func file_spacemesh_v1_mesh_proto_init() {
	if File_spacemesh_v1_mesh_proto != nil {
		return
	}
	file_spacemesh_v1_mesh_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spacemesh_v1_mesh_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spacemesh_v1_mesh_proto_goTypes,
		DependencyIndexes: file_spacemesh_v1_mesh_proto_depIdxs,
	}.Build()
	File_spacemesh_v1_mesh_proto = out.File
	file_spacemesh_v1_mesh_proto_rawDesc = nil
	file_spacemesh_v1_mesh_proto_goTypes = nil
	file_spacemesh_v1_mesh_proto_depIdxs = nil
}
