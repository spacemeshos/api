// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: post/v1/service.proto

package postv1

import (
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

type OperatorStatusResponse_Status int32

const (
	OperatorStatusResponse_UNUSED  OperatorStatusResponse_Status = 0
	OperatorStatusResponse_IDLE    OperatorStatusResponse_Status = 1
	OperatorStatusResponse_PROVING OperatorStatusResponse_Status = 2
)

// Enum value maps for OperatorStatusResponse_Status.
var (
	OperatorStatusResponse_Status_name = map[int32]string{
		0: "UNUSED",
		1: "IDLE",
		2: "PROVING",
	}
	OperatorStatusResponse_Status_value = map[string]int32{
		"UNUSED":  0,
		"IDLE":    1,
		"PROVING": 2,
	}
)

func (x OperatorStatusResponse_Status) Enum() *OperatorStatusResponse_Status {
	p := new(OperatorStatusResponse_Status)
	*p = x
	return p
}

func (x OperatorStatusResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperatorStatusResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_post_v1_service_proto_enumTypes[0].Descriptor()
}

func (OperatorStatusResponse_Status) Type() protoreflect.EnumType {
	return &file_post_v1_service_proto_enumTypes[0]
}

func (x OperatorStatusResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperatorStatusResponse_Status.Descriptor instead.
func (OperatorStatusResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_post_v1_service_proto_rawDescGZIP(), []int{1, 0}
}

type OperatorStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperatorStatusRequest) Reset() {
	*x = OperatorStatusRequest{}
	mi := &file_post_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorStatusRequest) ProtoMessage() {}

func (x *OperatorStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorStatusRequest.ProtoReflect.Descriptor instead.
func (*OperatorStatusRequest) Descriptor() ([]byte, []int) {
	return file_post_v1_service_proto_rawDescGZIP(), []int{0}
}

type OperatorStatusResponse struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Status        OperatorStatusResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=post.v1.OperatorStatusResponse_Status" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperatorStatusResponse) Reset() {
	*x = OperatorStatusResponse{}
	mi := &file_post_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorStatusResponse) ProtoMessage() {}

func (x *OperatorStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorStatusResponse.ProtoReflect.Descriptor instead.
func (*OperatorStatusResponse) Descriptor() ([]byte, []int) {
	return file_post_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *OperatorStatusResponse) GetStatus() OperatorStatusResponse_Status {
	if x != nil {
		return x.Status
	}
	return OperatorStatusResponse_UNUSED
}

var File_post_v1_service_proto protoreflect.FileDescriptor

const file_post_v1_service_proto_rawDesc = "" +
	"\n" +
	"\x15post/v1/service.proto\x12\apost.v1\x1a\x1bgoogle/api/visibility.proto\"\x17\n" +
	"\x15OperatorStatusRequest\"\x85\x01\n" +
	"\x16OperatorStatusResponse\x12>\n" +
	"\x06status\x18\x01 \x01(\x0e2&.post.v1.OperatorStatusResponse.StatusR\x06status\"+\n" +
	"\x06Status\x12\n" +
	"\n" +
	"\x06UNUSED\x10\x00\x12\b\n" +
	"\x04IDLE\x10\x01\x12\v\n" +
	"\aPROVING\x10\x022h\n" +
	"\x0fOperatorService\x12I\n" +
	"\x06Status\x12\x1e.post.v1.OperatorStatusRequest\x1a\x1f.post.v1.OperatorStatusResponse\x1a\n" +
	"\xfa\xd2\xe4\x93\x02\x04\x12\x02V1B\x8e\x01\n" +
	"\vcom.post.v1B\fServiceProtoP\x01Z4github.com/spacemeshos/api/release/go/post/v1;postv1\xa2\x02\x03PXX\xaa\x02\aPost.V1\xca\x02\aPost\\V1\xe2\x02\x13Post\\V1\\GPBMetadata\xea\x02\bPost::V1b\x06proto3"

var (
	file_post_v1_service_proto_rawDescOnce sync.Once
	file_post_v1_service_proto_rawDescData []byte
)

func file_post_v1_service_proto_rawDescGZIP() []byte {
	file_post_v1_service_proto_rawDescOnce.Do(func() {
		file_post_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_post_v1_service_proto_rawDesc), len(file_post_v1_service_proto_rawDesc)))
	})
	return file_post_v1_service_proto_rawDescData
}

var file_post_v1_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_post_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_post_v1_service_proto_goTypes = []any{
	(OperatorStatusResponse_Status)(0), // 0: post.v1.OperatorStatusResponse.Status
	(*OperatorStatusRequest)(nil),      // 1: post.v1.OperatorStatusRequest
	(*OperatorStatusResponse)(nil),     // 2: post.v1.OperatorStatusResponse
}
var file_post_v1_service_proto_depIdxs = []int32{
	0, // 0: post.v1.OperatorStatusResponse.status:type_name -> post.v1.OperatorStatusResponse.Status
	1, // 1: post.v1.OperatorService.Status:input_type -> post.v1.OperatorStatusRequest
	2, // 2: post.v1.OperatorService.Status:output_type -> post.v1.OperatorStatusResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_post_v1_service_proto_init() }
func file_post_v1_service_proto_init() {
	if File_post_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_post_v1_service_proto_rawDesc), len(file_post_v1_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_v1_service_proto_goTypes,
		DependencyIndexes: file_post_v1_service_proto_depIdxs,
		EnumInfos:         file_post_v1_service_proto_enumTypes,
		MessageInfos:      file_post_v1_service_proto_msgTypes,
	}.Build()
	File_post_v1_service_proto = out.File
	file_post_v1_service_proto_goTypes = nil
	file_post_v1_service_proto_depIdxs = nil
}
