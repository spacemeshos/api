// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: spacemesh/v1/global_state.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_spacemesh_v1_global_state_proto protoreflect.FileDescriptor

var file_spacemesh_v1_global_state_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x1a,
	0x25, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe0, 0x08, 0x0a, 0x12, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x55, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x12, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0f,
	0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x24, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x07,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x10, 0x53, 0x6d, 0x65, 0x73, 0x68,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x25, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6d, 0x65, 0x73, 0x68,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x11, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x12, 0x6c, 0x0a, 0x13, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x72, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x28, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x12, 0x5d, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x23, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12,
	0x66, 0x0a, 0x11, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x6f,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_spacemesh_v1_global_state_proto_goTypes = []interface{}{
	(*GetTemplateRequest)(nil),          // 0: spacemesh.v1.GetTemplateRequest
	(*GetTemplatesRequest)(nil),         // 1: spacemesh.v1.GetTemplatesRequest
	(*GetAppRequest)(nil),               // 2: spacemesh.v1.GetAppRequest
	(*GetStorageRequest)(nil),           // 3: spacemesh.v1.GetStorageRequest
	(*GlobalStateHashRequest)(nil),      // 4: spacemesh.v1.GlobalStateHashRequest
	(*AccountRequest)(nil),              // 5: spacemesh.v1.AccountRequest
	(*AccountDataQueryRequest)(nil),     // 6: spacemesh.v1.AccountDataQueryRequest
	(*SmesherDataQueryRequest)(nil),     // 7: spacemesh.v1.SmesherDataQueryRequest
	(*AccountDataStreamRequest)(nil),    // 8: spacemesh.v1.AccountDataStreamRequest
	(*SmesherRewardStreamRequest)(nil),  // 9: spacemesh.v1.SmesherRewardStreamRequest
	(*AppEventStreamRequest)(nil),       // 10: spacemesh.v1.AppEventStreamRequest
	(*GlobalStateStreamRequest)(nil),    // 11: spacemesh.v1.GlobalStateStreamRequest
	(*GetTemplateResponse)(nil),         // 12: spacemesh.v1.GetTemplateResponse
	(*GetTemplatesResponse)(nil),        // 13: spacemesh.v1.GetTemplatesResponse
	(*GetAppResponse)(nil),              // 14: spacemesh.v1.GetAppResponse
	(*GetStorageResponse)(nil),          // 15: spacemesh.v1.GetStorageResponse
	(*GlobalStateHashResponse)(nil),     // 16: spacemesh.v1.GlobalStateHashResponse
	(*AccountResponse)(nil),             // 17: spacemesh.v1.AccountResponse
	(*AccountDataQueryResponse)(nil),    // 18: spacemesh.v1.AccountDataQueryResponse
	(*SmesherDataQueryResponse)(nil),    // 19: spacemesh.v1.SmesherDataQueryResponse
	(*AccountDataStreamResponse)(nil),   // 20: spacemesh.v1.AccountDataStreamResponse
	(*SmesherRewardStreamResponse)(nil), // 21: spacemesh.v1.SmesherRewardStreamResponse
	(*AppEventStreamResponse)(nil),      // 22: spacemesh.v1.AppEventStreamResponse
	(*GlobalStateStreamResponse)(nil),   // 23: spacemesh.v1.GlobalStateStreamResponse
}
var file_spacemesh_v1_global_state_proto_depIdxs = []int32{
	0,  // 0: spacemesh.v1.GlobalStateService.GetTemplate:input_type -> spacemesh.v1.GetTemplateRequest
	1,  // 1: spacemesh.v1.GlobalStateService.GetTemplates:input_type -> spacemesh.v1.GetTemplatesRequest
	2,  // 2: spacemesh.v1.GlobalStateService.GetApp:input_type -> spacemesh.v1.GetAppRequest
	3,  // 3: spacemesh.v1.GlobalStateService.GetStorage:input_type -> spacemesh.v1.GetStorageRequest
	4,  // 4: spacemesh.v1.GlobalStateService.GlobalStateHash:input_type -> spacemesh.v1.GlobalStateHashRequest
	5,  // 5: spacemesh.v1.GlobalStateService.Account:input_type -> spacemesh.v1.AccountRequest
	6,  // 6: spacemesh.v1.GlobalStateService.AccountDataQuery:input_type -> spacemesh.v1.AccountDataQueryRequest
	7,  // 7: spacemesh.v1.GlobalStateService.SmesherDataQuery:input_type -> spacemesh.v1.SmesherDataQueryRequest
	8,  // 8: spacemesh.v1.GlobalStateService.AccountDataStream:input_type -> spacemesh.v1.AccountDataStreamRequest
	9,  // 9: spacemesh.v1.GlobalStateService.SmesherRewardStream:input_type -> spacemesh.v1.SmesherRewardStreamRequest
	10, // 10: spacemesh.v1.GlobalStateService.AppEventStream:input_type -> spacemesh.v1.AppEventStreamRequest
	11, // 11: spacemesh.v1.GlobalStateService.GlobalStateStream:input_type -> spacemesh.v1.GlobalStateStreamRequest
	12, // 12: spacemesh.v1.GlobalStateService.GetTemplate:output_type -> spacemesh.v1.GetTemplateResponse
	13, // 13: spacemesh.v1.GlobalStateService.GetTemplates:output_type -> spacemesh.v1.GetTemplatesResponse
	14, // 14: spacemesh.v1.GlobalStateService.GetApp:output_type -> spacemesh.v1.GetAppResponse
	15, // 15: spacemesh.v1.GlobalStateService.GetStorage:output_type -> spacemesh.v1.GetStorageResponse
	16, // 16: spacemesh.v1.GlobalStateService.GlobalStateHash:output_type -> spacemesh.v1.GlobalStateHashResponse
	17, // 17: spacemesh.v1.GlobalStateService.Account:output_type -> spacemesh.v1.AccountResponse
	18, // 18: spacemesh.v1.GlobalStateService.AccountDataQuery:output_type -> spacemesh.v1.AccountDataQueryResponse
	19, // 19: spacemesh.v1.GlobalStateService.SmesherDataQuery:output_type -> spacemesh.v1.SmesherDataQueryResponse
	20, // 20: spacemesh.v1.GlobalStateService.AccountDataStream:output_type -> spacemesh.v1.AccountDataStreamResponse
	21, // 21: spacemesh.v1.GlobalStateService.SmesherRewardStream:output_type -> spacemesh.v1.SmesherRewardStreamResponse
	22, // 22: spacemesh.v1.GlobalStateService.AppEventStream:output_type -> spacemesh.v1.AppEventStreamResponse
	23, // 23: spacemesh.v1.GlobalStateService.GlobalStateStream:output_type -> spacemesh.v1.GlobalStateStreamResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_spacemesh_v1_global_state_proto_init() }
func file_spacemesh_v1_global_state_proto_init() {
	if File_spacemesh_v1_global_state_proto != nil {
		return
	}
	file_spacemesh_v1_global_state_types_proto_init()
	file_spacemesh_v1_app_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spacemesh_v1_global_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spacemesh_v1_global_state_proto_goTypes,
		DependencyIndexes: file_spacemesh_v1_global_state_proto_depIdxs,
	}.Build()
	File_spacemesh_v1_global_state_proto = out.File
	file_spacemesh_v1_global_state_proto_rawDesc = nil
	file_spacemesh_v1_global_state_proto_goTypes = nil
	file_spacemesh_v1_global_state_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GlobalStateServiceClient is the client API for GlobalStateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GlobalStateServiceClient interface {
	// Get an app template meta-data and code
	GetTemplate(ctx context.Context, in *GetTemplateRequest, opts ...grpc.CallOption) (*GetTemplateResponse, error)
	// Get all app templates
	GetTemplates(ctx context.Context, in *GetTemplatesRequest, opts ...grpc.CallOption) (*GetTemplatesResponse, error)
	// Get app metadata and optionally current state
	GetApp(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*GetAppResponse, error)
	// Get an app's variable metadata and optionally value
	GetStorage(ctx context.Context, in *GetStorageRequest, opts ...grpc.CallOption) (*GetStorageResponse, error)
	// Latest computed global state - layer and its root hash
	GlobalStateHash(ctx context.Context, in *GlobalStateHashRequest, opts ...grpc.CallOption) (*GlobalStateHashResponse, error)
	// Account info in the current global state.
	Account(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	// Query for account related data such as rewards, tx receipts and account info
	//
	// Note: it might be too expensive to add a param for layer to get these results from
	// as it may require indexing all global state changes per account by layer.
	// If it is possible to index by layer then we should add param start_layer to
	// AccountDataParams. Currently it will return data from genesis.
	AccountDataQuery(ctx context.Context, in *AccountDataQueryRequest, opts ...grpc.CallOption) (*AccountDataQueryResponse, error)
	// Query for smesher data. Currently returns smesher rewards.
	// Note: Not supporting start_layer yet as it may require to index all rewards by
	// smesher and by layer id or allow for queries from a layer and later....
	SmesherDataQuery(ctx context.Context, in *SmesherDataQueryRequest, opts ...grpc.CallOption) (*SmesherDataQueryResponse, error)
	// Get a stream of account related changes such as account balance change,
	// tx receipts and rewards
	AccountDataStream(ctx context.Context, in *AccountDataStreamRequest, opts ...grpc.CallOption) (GlobalStateService_AccountDataStreamClient, error)
	// Rewards awarded to a smesher id
	SmesherRewardStream(ctx context.Context, in *SmesherRewardStreamRequest, opts ...grpc.CallOption) (GlobalStateService_SmesherRewardStreamClient, error)
	// App Events - emitted by app methods impl code trigged by an
	// app transaction
	AppEventStream(ctx context.Context, in *AppEventStreamRequest, opts ...grpc.CallOption) (GlobalStateService_AppEventStreamClient, error)
	// New global state computed for a layer by the STF
	GlobalStateStream(ctx context.Context, in *GlobalStateStreamRequest, opts ...grpc.CallOption) (GlobalStateService_GlobalStateStreamClient, error)
}

type globalStateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGlobalStateServiceClient(cc grpc.ClientConnInterface) GlobalStateServiceClient {
	return &globalStateServiceClient{cc}
}

func (c *globalStateServiceClient) GetTemplate(ctx context.Context, in *GetTemplateRequest, opts ...grpc.CallOption) (*GetTemplateResponse, error) {
	out := new(GetTemplateResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.GlobalStateService/GetTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalStateServiceClient) GetTemplates(ctx context.Context, in *GetTemplatesRequest, opts ...grpc.CallOption) (*GetTemplatesResponse, error) {
	out := new(GetTemplatesResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.GlobalStateService/GetTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalStateServiceClient) GetApp(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*GetAppResponse, error) {
	out := new(GetAppResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.GlobalStateService/GetApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalStateServiceClient) GetStorage(ctx context.Context, in *GetStorageRequest, opts ...grpc.CallOption) (*GetStorageResponse, error) {
	out := new(GetStorageResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.GlobalStateService/GetStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalStateServiceClient) GlobalStateHash(ctx context.Context, in *GlobalStateHashRequest, opts ...grpc.CallOption) (*GlobalStateHashResponse, error) {
	out := new(GlobalStateHashResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.GlobalStateService/GlobalStateHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalStateServiceClient) Account(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.GlobalStateService/Account", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalStateServiceClient) AccountDataQuery(ctx context.Context, in *AccountDataQueryRequest, opts ...grpc.CallOption) (*AccountDataQueryResponse, error) {
	out := new(AccountDataQueryResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.GlobalStateService/AccountDataQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalStateServiceClient) SmesherDataQuery(ctx context.Context, in *SmesherDataQueryRequest, opts ...grpc.CallOption) (*SmesherDataQueryResponse, error) {
	out := new(SmesherDataQueryResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.GlobalStateService/SmesherDataQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalStateServiceClient) AccountDataStream(ctx context.Context, in *AccountDataStreamRequest, opts ...grpc.CallOption) (GlobalStateService_AccountDataStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GlobalStateService_serviceDesc.Streams[0], "/spacemesh.v1.GlobalStateService/AccountDataStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &globalStateServiceAccountDataStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GlobalStateService_AccountDataStreamClient interface {
	Recv() (*AccountDataStreamResponse, error)
	grpc.ClientStream
}

type globalStateServiceAccountDataStreamClient struct {
	grpc.ClientStream
}

func (x *globalStateServiceAccountDataStreamClient) Recv() (*AccountDataStreamResponse, error) {
	m := new(AccountDataStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *globalStateServiceClient) SmesherRewardStream(ctx context.Context, in *SmesherRewardStreamRequest, opts ...grpc.CallOption) (GlobalStateService_SmesherRewardStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GlobalStateService_serviceDesc.Streams[1], "/spacemesh.v1.GlobalStateService/SmesherRewardStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &globalStateServiceSmesherRewardStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GlobalStateService_SmesherRewardStreamClient interface {
	Recv() (*SmesherRewardStreamResponse, error)
	grpc.ClientStream
}

type globalStateServiceSmesherRewardStreamClient struct {
	grpc.ClientStream
}

func (x *globalStateServiceSmesherRewardStreamClient) Recv() (*SmesherRewardStreamResponse, error) {
	m := new(SmesherRewardStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *globalStateServiceClient) AppEventStream(ctx context.Context, in *AppEventStreamRequest, opts ...grpc.CallOption) (GlobalStateService_AppEventStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GlobalStateService_serviceDesc.Streams[2], "/spacemesh.v1.GlobalStateService/AppEventStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &globalStateServiceAppEventStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GlobalStateService_AppEventStreamClient interface {
	Recv() (*AppEventStreamResponse, error)
	grpc.ClientStream
}

type globalStateServiceAppEventStreamClient struct {
	grpc.ClientStream
}

func (x *globalStateServiceAppEventStreamClient) Recv() (*AppEventStreamResponse, error) {
	m := new(AppEventStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *globalStateServiceClient) GlobalStateStream(ctx context.Context, in *GlobalStateStreamRequest, opts ...grpc.CallOption) (GlobalStateService_GlobalStateStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GlobalStateService_serviceDesc.Streams[3], "/spacemesh.v1.GlobalStateService/GlobalStateStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &globalStateServiceGlobalStateStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GlobalStateService_GlobalStateStreamClient interface {
	Recv() (*GlobalStateStreamResponse, error)
	grpc.ClientStream
}

type globalStateServiceGlobalStateStreamClient struct {
	grpc.ClientStream
}

func (x *globalStateServiceGlobalStateStreamClient) Recv() (*GlobalStateStreamResponse, error) {
	m := new(GlobalStateStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GlobalStateServiceServer is the server API for GlobalStateService service.
type GlobalStateServiceServer interface {
	// Get an app template meta-data and code
	GetTemplate(context.Context, *GetTemplateRequest) (*GetTemplateResponse, error)
	// Get all app templates
	GetTemplates(context.Context, *GetTemplatesRequest) (*GetTemplatesResponse, error)
	// Get app metadata and optionally current state
	GetApp(context.Context, *GetAppRequest) (*GetAppResponse, error)
	// Get an app's variable metadata and optionally value
	GetStorage(context.Context, *GetStorageRequest) (*GetStorageResponse, error)
	// Latest computed global state - layer and its root hash
	GlobalStateHash(context.Context, *GlobalStateHashRequest) (*GlobalStateHashResponse, error)
	// Account info in the current global state.
	Account(context.Context, *AccountRequest) (*AccountResponse, error)
	// Query for account related data such as rewards, tx receipts and account info
	//
	// Note: it might be too expensive to add a param for layer to get these results from
	// as it may require indexing all global state changes per account by layer.
	// If it is possible to index by layer then we should add param start_layer to
	// AccountDataParams. Currently it will return data from genesis.
	AccountDataQuery(context.Context, *AccountDataQueryRequest) (*AccountDataQueryResponse, error)
	// Query for smesher data. Currently returns smesher rewards.
	// Note: Not supporting start_layer yet as it may require to index all rewards by
	// smesher and by layer id or allow for queries from a layer and later....
	SmesherDataQuery(context.Context, *SmesherDataQueryRequest) (*SmesherDataQueryResponse, error)
	// Get a stream of account related changes such as account balance change,
	// tx receipts and rewards
	AccountDataStream(*AccountDataStreamRequest, GlobalStateService_AccountDataStreamServer) error
	// Rewards awarded to a smesher id
	SmesherRewardStream(*SmesherRewardStreamRequest, GlobalStateService_SmesherRewardStreamServer) error
	// App Events - emitted by app methods impl code trigged by an
	// app transaction
	AppEventStream(*AppEventStreamRequest, GlobalStateService_AppEventStreamServer) error
	// New global state computed for a layer by the STF
	GlobalStateStream(*GlobalStateStreamRequest, GlobalStateService_GlobalStateStreamServer) error
}

// UnimplementedGlobalStateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGlobalStateServiceServer struct {
}

func (*UnimplementedGlobalStateServiceServer) GetTemplate(context.Context, *GetTemplateRequest) (*GetTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplate not implemented")
}
func (*UnimplementedGlobalStateServiceServer) GetTemplates(context.Context, *GetTemplatesRequest) (*GetTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplates not implemented")
}
func (*UnimplementedGlobalStateServiceServer) GetApp(context.Context, *GetAppRequest) (*GetAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApp not implemented")
}
func (*UnimplementedGlobalStateServiceServer) GetStorage(context.Context, *GetStorageRequest) (*GetStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorage not implemented")
}
func (*UnimplementedGlobalStateServiceServer) GlobalStateHash(context.Context, *GlobalStateHashRequest) (*GlobalStateHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GlobalStateHash not implemented")
}
func (*UnimplementedGlobalStateServiceServer) Account(context.Context, *AccountRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Account not implemented")
}
func (*UnimplementedGlobalStateServiceServer) AccountDataQuery(context.Context, *AccountDataQueryRequest) (*AccountDataQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountDataQuery not implemented")
}
func (*UnimplementedGlobalStateServiceServer) SmesherDataQuery(context.Context, *SmesherDataQueryRequest) (*SmesherDataQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SmesherDataQuery not implemented")
}
func (*UnimplementedGlobalStateServiceServer) AccountDataStream(*AccountDataStreamRequest, GlobalStateService_AccountDataStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AccountDataStream not implemented")
}
func (*UnimplementedGlobalStateServiceServer) SmesherRewardStream(*SmesherRewardStreamRequest, GlobalStateService_SmesherRewardStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SmesherRewardStream not implemented")
}
func (*UnimplementedGlobalStateServiceServer) AppEventStream(*AppEventStreamRequest, GlobalStateService_AppEventStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AppEventStream not implemented")
}
func (*UnimplementedGlobalStateServiceServer) GlobalStateStream(*GlobalStateStreamRequest, GlobalStateService_GlobalStateStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GlobalStateStream not implemented")
}

func RegisterGlobalStateServiceServer(s *grpc.Server, srv GlobalStateServiceServer) {
	s.RegisterService(&_GlobalStateService_serviceDesc, srv)
}

func _GlobalStateService_GetTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalStateServiceServer).GetTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.GlobalStateService/GetTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalStateServiceServer).GetTemplate(ctx, req.(*GetTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalStateService_GetTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalStateServiceServer).GetTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.GlobalStateService/GetTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalStateServiceServer).GetTemplates(ctx, req.(*GetTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalStateService_GetApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalStateServiceServer).GetApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.GlobalStateService/GetApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalStateServiceServer).GetApp(ctx, req.(*GetAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalStateService_GetStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalStateServiceServer).GetStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.GlobalStateService/GetStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalStateServiceServer).GetStorage(ctx, req.(*GetStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalStateService_GlobalStateHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GlobalStateHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalStateServiceServer).GlobalStateHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.GlobalStateService/GlobalStateHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalStateServiceServer).GlobalStateHash(ctx, req.(*GlobalStateHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalStateService_Account_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalStateServiceServer).Account(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.GlobalStateService/Account",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalStateServiceServer).Account(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalStateService_AccountDataQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountDataQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalStateServiceServer).AccountDataQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.GlobalStateService/AccountDataQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalStateServiceServer).AccountDataQuery(ctx, req.(*AccountDataQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalStateService_SmesherDataQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmesherDataQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalStateServiceServer).SmesherDataQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.GlobalStateService/SmesherDataQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalStateServiceServer).SmesherDataQuery(ctx, req.(*SmesherDataQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalStateService_AccountDataStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AccountDataStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GlobalStateServiceServer).AccountDataStream(m, &globalStateServiceAccountDataStreamServer{stream})
}

type GlobalStateService_AccountDataStreamServer interface {
	Send(*AccountDataStreamResponse) error
	grpc.ServerStream
}

type globalStateServiceAccountDataStreamServer struct {
	grpc.ServerStream
}

func (x *globalStateServiceAccountDataStreamServer) Send(m *AccountDataStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GlobalStateService_SmesherRewardStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SmesherRewardStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GlobalStateServiceServer).SmesherRewardStream(m, &globalStateServiceSmesherRewardStreamServer{stream})
}

type GlobalStateService_SmesherRewardStreamServer interface {
	Send(*SmesherRewardStreamResponse) error
	grpc.ServerStream
}

type globalStateServiceSmesherRewardStreamServer struct {
	grpc.ServerStream
}

func (x *globalStateServiceSmesherRewardStreamServer) Send(m *SmesherRewardStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GlobalStateService_AppEventStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AppEventStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GlobalStateServiceServer).AppEventStream(m, &globalStateServiceAppEventStreamServer{stream})
}

type GlobalStateService_AppEventStreamServer interface {
	Send(*AppEventStreamResponse) error
	grpc.ServerStream
}

type globalStateServiceAppEventStreamServer struct {
	grpc.ServerStream
}

func (x *globalStateServiceAppEventStreamServer) Send(m *AppEventStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GlobalStateService_GlobalStateStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GlobalStateStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GlobalStateServiceServer).GlobalStateStream(m, &globalStateServiceGlobalStateStreamServer{stream})
}

type GlobalStateService_GlobalStateStreamServer interface {
	Send(*GlobalStateStreamResponse) error
	grpc.ServerStream
}

type globalStateServiceGlobalStateStreamServer struct {
	grpc.ServerStream
}

func (x *globalStateServiceGlobalStateStreamServer) Send(m *GlobalStateStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GlobalStateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v1.GlobalStateService",
	HandlerType: (*GlobalStateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTemplate",
			Handler:    _GlobalStateService_GetTemplate_Handler,
		},
		{
			MethodName: "GetTemplates",
			Handler:    _GlobalStateService_GetTemplates_Handler,
		},
		{
			MethodName: "GetApp",
			Handler:    _GlobalStateService_GetApp_Handler,
		},
		{
			MethodName: "GetStorage",
			Handler:    _GlobalStateService_GetStorage_Handler,
		},
		{
			MethodName: "GlobalStateHash",
			Handler:    _GlobalStateService_GlobalStateHash_Handler,
		},
		{
			MethodName: "Account",
			Handler:    _GlobalStateService_Account_Handler,
		},
		{
			MethodName: "AccountDataQuery",
			Handler:    _GlobalStateService_AccountDataQuery_Handler,
		},
		{
			MethodName: "SmesherDataQuery",
			Handler:    _GlobalStateService_SmesherDataQuery_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AccountDataStream",
			Handler:       _GlobalStateService_AccountDataStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SmesherRewardStream",
			Handler:       _GlobalStateService_SmesherRewardStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AppEventStream",
			Handler:       _GlobalStateService_AppEventStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GlobalStateStream",
			Handler:       _GlobalStateService_GlobalStateStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "spacemesh/v1/global_state.proto",
}
