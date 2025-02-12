// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: spacemesh/v2beta1/smeshing.proto

package spacemeshv2beta1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SmeshingVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SmeshingVersionRequest) Reset() {
	*x = SmeshingVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2beta1_smeshing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmeshingVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmeshingVersionRequest) ProtoMessage() {}

func (x *SmeshingVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_smeshing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmeshingVersionRequest.ProtoReflect.Descriptor instead.
func (*SmeshingVersionRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_smeshing_proto_rawDescGZIP(), []int{0}
}

type SmeshingVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"` // semantic version of the node software
}

func (x *SmeshingVersionResponse) Reset() {
	*x = SmeshingVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2beta1_smeshing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmeshingVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmeshingVersionResponse) ProtoMessage() {}

func (x *SmeshingVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_smeshing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmeshingVersionResponse.ProtoReflect.Descriptor instead.
func (*SmeshingVersionResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_smeshing_proto_rawDescGZIP(), []int{1}
}

func (x *SmeshingVersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type SmeshingBuildRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SmeshingBuildRequest) Reset() {
	*x = SmeshingBuildRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2beta1_smeshing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmeshingBuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmeshingBuildRequest) ProtoMessage() {}

func (x *SmeshingBuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_smeshing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmeshingBuildRequest.ProtoReflect.Descriptor instead.
func (*SmeshingBuildRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_smeshing_proto_rawDescGZIP(), []int{2}
}

type SmeshingBuildResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Build string `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty"` // git commit hash of the build
}

func (x *SmeshingBuildResponse) Reset() {
	*x = SmeshingBuildResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2beta1_smeshing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmeshingBuildResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmeshingBuildResponse) ProtoMessage() {}

func (x *SmeshingBuildResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_smeshing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmeshingBuildResponse.ProtoReflect.Descriptor instead.
func (*SmeshingBuildResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_smeshing_proto_rawDescGZIP(), []int{3}
}

func (x *SmeshingBuildResponse) GetBuild() string {
	if x != nil {
		return x.Build
	}
	return ""
}

var File_spacemesh_v2beta1_smeshing_proto protoreflect.FileDescriptor

var file_spacemesh_v2beta1_smeshing_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x73, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x18, 0x0a, 0x16, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x17, 0x53, 0x6d,
	0x65, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x16, 0x0a, 0x14, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x15, 0x53, 0x6d, 0x65, 0x73, 0x68,
	0x69, 0x6e, 0x67, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x32, 0xc8, 0x02, 0x0a, 0x0f, 0x53, 0x6d, 0x65, 0x73, 0x68,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x94, 0x01, 0x0a, 0x07, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x6d, 0x65, 0x73, 0x68,
	0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x12, 0x2a, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x8c, 0x01, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x27, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x53, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x12, 0x28, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x6d, 0x65, 0x73, 0x68,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x1a, 0x0f, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x42, 0xd5, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0d, 0x53, 0x6d, 0x65,
	0x73, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x11, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xca, 0x02, 0x11, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1d, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_spacemesh_v2beta1_smeshing_proto_rawDescOnce sync.Once
	file_spacemesh_v2beta1_smeshing_proto_rawDescData = file_spacemesh_v2beta1_smeshing_proto_rawDesc
)

func file_spacemesh_v2beta1_smeshing_proto_rawDescGZIP() []byte {
	file_spacemesh_v2beta1_smeshing_proto_rawDescOnce.Do(func() {
		file_spacemesh_v2beta1_smeshing_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacemesh_v2beta1_smeshing_proto_rawDescData)
	})
	return file_spacemesh_v2beta1_smeshing_proto_rawDescData
}

var file_spacemesh_v2beta1_smeshing_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spacemesh_v2beta1_smeshing_proto_goTypes = []interface{}{
	(*SmeshingVersionRequest)(nil),  // 0: spacemesh.v2beta1.SmeshingVersionRequest
	(*SmeshingVersionResponse)(nil), // 1: spacemesh.v2beta1.SmeshingVersionResponse
	(*SmeshingBuildRequest)(nil),    // 2: spacemesh.v2beta1.SmeshingBuildRequest
	(*SmeshingBuildResponse)(nil),   // 3: spacemesh.v2beta1.SmeshingBuildResponse
}
var file_spacemesh_v2beta1_smeshing_proto_depIdxs = []int32{
	0, // 0: spacemesh.v2beta1.SmeshingService.Version:input_type -> spacemesh.v2beta1.SmeshingVersionRequest
	2, // 1: spacemesh.v2beta1.SmeshingService.Build:input_type -> spacemesh.v2beta1.SmeshingBuildRequest
	1, // 2: spacemesh.v2beta1.SmeshingService.Version:output_type -> spacemesh.v2beta1.SmeshingVersionResponse
	3, // 3: spacemesh.v2beta1.SmeshingService.Build:output_type -> spacemesh.v2beta1.SmeshingBuildResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spacemesh_v2beta1_smeshing_proto_init() }
func file_spacemesh_v2beta1_smeshing_proto_init() {
	if File_spacemesh_v2beta1_smeshing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spacemesh_v2beta1_smeshing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmeshingVersionRequest); i {
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
		file_spacemesh_v2beta1_smeshing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmeshingVersionResponse); i {
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
		file_spacemesh_v2beta1_smeshing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmeshingBuildRequest); i {
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
		file_spacemesh_v2beta1_smeshing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmeshingBuildResponse); i {
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
			RawDescriptor: file_spacemesh_v2beta1_smeshing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spacemesh_v2beta1_smeshing_proto_goTypes,
		DependencyIndexes: file_spacemesh_v2beta1_smeshing_proto_depIdxs,
		MessageInfos:      file_spacemesh_v2beta1_smeshing_proto_msgTypes,
	}.Build()
	File_spacemesh_v2beta1_smeshing_proto = out.File
	file_spacemesh_v2beta1_smeshing_proto_rawDesc = nil
	file_spacemesh_v2beta1_smeshing_proto_goTypes = nil
	file_spacemesh_v2beta1_smeshing_proto_depIdxs = nil
}
