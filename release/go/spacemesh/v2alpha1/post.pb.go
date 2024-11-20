// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: spacemesh/v2alpha1/post.proto

package spacemeshv2alpha1

import (
	_ "google.golang.org/genproto/googleapis/api/visibility"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PostConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BitsPerLabel  uint32 `protobuf:"varint,1,opt,name=bits_per_label,json=bitsPerLabel,proto3" json:"bits_per_label,omitempty"`
	LabelsPerUnit uint64 `protobuf:"varint,2,opt,name=labels_per_unit,json=labelsPerUnit,proto3" json:"labels_per_unit,omitempty"`
	MinNumUnits   uint32 `protobuf:"varint,3,opt,name=min_num_units,json=minNumUnits,proto3" json:"min_num_units,omitempty"`
	MaxNumUnits   uint32 `protobuf:"varint,4,opt,name=max_num_units,json=maxNumUnits,proto3" json:"max_num_units,omitempty"`
	K1            uint32 `protobuf:"varint,5,opt,name=k1,proto3" json:"k1,omitempty"`
	K2            uint32 `protobuf:"varint,6,opt,name=k2,proto3" json:"k2,omitempty"`
}

func (x *PostConfigResponse) Reset() {
	*x = PostConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostConfigResponse) ProtoMessage() {}

func (x *PostConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostConfigResponse.ProtoReflect.Descriptor instead.
func (*PostConfigResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_post_proto_rawDescGZIP(), []int{0}
}

func (x *PostConfigResponse) GetBitsPerLabel() uint32 {
	if x != nil {
		return x.BitsPerLabel
	}
	return 0
}

func (x *PostConfigResponse) GetLabelsPerUnit() uint64 {
	if x != nil {
		return x.LabelsPerUnit
	}
	return 0
}

func (x *PostConfigResponse) GetMinNumUnits() uint32 {
	if x != nil {
		return x.MinNumUnits
	}
	return 0
}

func (x *PostConfigResponse) GetMaxNumUnits() uint32 {
	if x != nil {
		return x.MaxNumUnits
	}
	return 0
}

func (x *PostConfigResponse) GetK1() uint32 {
	if x != nil {
		return x.K1
	}
	return 0
}

func (x *PostConfigResponse) GetK2() uint32 {
	if x != nil {
		return x.K2
	}
	return 0
}

type PostConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostConfigRequest) Reset() {
	*x = PostConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostConfigRequest) ProtoMessage() {}

func (x *PostConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostConfigRequest.ProtoReflect.Descriptor instead.
func (*PostConfigRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_post_proto_rawDescGZIP(), []int{1}
}

var File_spacemesh_v2alpha1_post_proto protoreflect.FileDescriptor

var file_spacemesh_v2alpha1_post_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xca, 0x01, 0x0a, 0x12, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x69, 0x74, 0x73,
	0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x62, 0x69, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x26,
	0x0a, 0x0f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x75, 0x6e, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x50,
	0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x5f, 0x6e, 0x75,
	0x6d, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d,
	0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61,
	0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x6b, 0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x6b, 0x31, 0x12, 0x0e,
	0x0a, 0x02, 0x6b, 0x32, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x6b, 0x32, 0x22, 0x13,
	0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x32, 0x72, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x57, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x0a, 0xfa, 0xd2, 0xe4,
	0x93, 0x02, 0x04, 0x12, 0x02, 0x56, 0x32, 0x42, 0xd8, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x42, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58,
	0x58, 0xaa, 0x02, 0x12, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x56, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x12, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x1e, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spacemesh_v2alpha1_post_proto_rawDescOnce sync.Once
	file_spacemesh_v2alpha1_post_proto_rawDescData = file_spacemesh_v2alpha1_post_proto_rawDesc
)

func file_spacemesh_v2alpha1_post_proto_rawDescGZIP() []byte {
	file_spacemesh_v2alpha1_post_proto_rawDescOnce.Do(func() {
		file_spacemesh_v2alpha1_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacemesh_v2alpha1_post_proto_rawDescData)
	})
	return file_spacemesh_v2alpha1_post_proto_rawDescData
}

var file_spacemesh_v2alpha1_post_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_spacemesh_v2alpha1_post_proto_goTypes = []interface{}{
	(*PostConfigResponse)(nil), // 0: spacemesh.v2alpha1.PostConfigResponse
	(*PostConfigRequest)(nil),  // 1: spacemesh.v2alpha1.PostConfigRequest
}
var file_spacemesh_v2alpha1_post_proto_depIdxs = []int32{
	1, // 0: spacemesh.v2alpha1.PostService.Config:input_type -> spacemesh.v2alpha1.PostConfigRequest
	0, // 1: spacemesh.v2alpha1.PostService.Config:output_type -> spacemesh.v2alpha1.PostConfigResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spacemesh_v2alpha1_post_proto_init() }
func file_spacemesh_v2alpha1_post_proto_init() {
	if File_spacemesh_v2alpha1_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spacemesh_v2alpha1_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostConfigResponse); i {
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
		file_spacemesh_v2alpha1_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostConfigRequest); i {
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
			RawDescriptor: file_spacemesh_v2alpha1_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spacemesh_v2alpha1_post_proto_goTypes,
		DependencyIndexes: file_spacemesh_v2alpha1_post_proto_depIdxs,
		MessageInfos:      file_spacemesh_v2alpha1_post_proto_msgTypes,
	}.Build()
	File_spacemesh_v2alpha1_post_proto = out.File
	file_spacemesh_v2alpha1_post_proto_rawDesc = nil
	file_spacemesh_v2alpha1_post_proto_goTypes = nil
	file_spacemesh_v2alpha1_post_proto_depIdxs = nil
}