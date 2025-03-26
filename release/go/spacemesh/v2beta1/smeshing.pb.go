// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
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
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SmeshingVersionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SmeshingVersionRequest) Reset() {
	*x = SmeshingVersionRequest{}
	mi := &file_spacemesh_v2beta1_smeshing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SmeshingVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmeshingVersionRequest) ProtoMessage() {}

func (x *SmeshingVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_smeshing_proto_msgTypes[0]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"` // semantic version of the node software
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SmeshingVersionResponse) Reset() {
	*x = SmeshingVersionResponse{}
	mi := &file_spacemesh_v2beta1_smeshing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SmeshingVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmeshingVersionResponse) ProtoMessage() {}

func (x *SmeshingVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_smeshing_proto_msgTypes[1]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SmeshingBuildRequest) Reset() {
	*x = SmeshingBuildRequest{}
	mi := &file_spacemesh_v2beta1_smeshing_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SmeshingBuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmeshingBuildRequest) ProtoMessage() {}

func (x *SmeshingBuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_smeshing_proto_msgTypes[2]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Build         string                 `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty"` // git commit hash of the build
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SmeshingBuildResponse) Reset() {
	*x = SmeshingBuildResponse{}
	mi := &file_spacemesh_v2beta1_smeshing_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SmeshingBuildResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmeshingBuildResponse) ProtoMessage() {}

func (x *SmeshingBuildResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_smeshing_proto_msgTypes[3]
	if x != nil {
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

const file_spacemesh_v2beta1_smeshing_proto_rawDesc = "" +
	"\n" +
	" spacemesh/v2beta1/smeshing.proto\x12\x11spacemesh.v2beta1\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/api/visibility.proto\"\x18\n" +
	"\x16SmeshingVersionRequest\"3\n" +
	"\x17SmeshingVersionResponse\x12\x18\n" +
	"\aversion\x18\x01 \x01(\tR\aversion\"\x16\n" +
	"\x14SmeshingBuildRequest\"-\n" +
	"\x15SmeshingBuildResponse\x12\x14\n" +
	"\x05build\x18\x01 \x01(\tR\x05build2\xc8\x02\n" +
	"\x0fSmeshingService\x12\x94\x01\n" +
	"\aVersion\x12).spacemesh.v2beta1.SmeshingVersionRequest\x1a*.spacemesh.v2beta1.SmeshingVersionResponse\"2\x82\xd3\xe4\x93\x02,\x12*/spacemesh.v2beta1.SmeshingService/Version\x12\x8c\x01\n" +
	"\x05Build\x12'.spacemesh.v2beta1.SmeshingBuildRequest\x1a(.spacemesh.v2beta1.SmeshingBuildResponse\"0\x82\xd3\xe4\x93\x02*\x12(/spacemesh.v2beta1.SmeshingService/Build\x1a\x0f\xfa\xd2\xe4\x93\x02\t\x12\av2beta1B\xd5\x01\n" +
	"\x15com.spacemesh.v2beta1B\rSmeshingProtoP\x01ZHgithub.com/spacemeshos/api/release/go/spacemesh/v2beta1;spacemeshv2beta1\xa2\x02\x03SXX\xaa\x02\x11Spacemesh.V2beta1\xca\x02\x11Spacemesh\\V2beta1\xe2\x02\x1dSpacemesh\\V2beta1\\GPBMetadata\xea\x02\x12Spacemesh::V2beta1b\x06proto3"

var (
	file_spacemesh_v2beta1_smeshing_proto_rawDescOnce sync.Once
	file_spacemesh_v2beta1_smeshing_proto_rawDescData []byte
)

func file_spacemesh_v2beta1_smeshing_proto_rawDescGZIP() []byte {
	file_spacemesh_v2beta1_smeshing_proto_rawDescOnce.Do(func() {
		file_spacemesh_v2beta1_smeshing_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spacemesh_v2beta1_smeshing_proto_rawDesc), len(file_spacemesh_v2beta1_smeshing_proto_rawDesc)))
	})
	return file_spacemesh_v2beta1_smeshing_proto_rawDescData
}

var file_spacemesh_v2beta1_smeshing_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spacemesh_v2beta1_smeshing_proto_goTypes = []any{
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spacemesh_v2beta1_smeshing_proto_rawDesc), len(file_spacemesh_v2beta1_smeshing_proto_rawDesc)),
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
	file_spacemesh_v2beta1_smeshing_proto_goTypes = nil
	file_spacemesh_v2beta1_smeshing_proto_depIdxs = nil
}
