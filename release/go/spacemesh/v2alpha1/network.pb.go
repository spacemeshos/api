// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: spacemesh/v2alpha1/network.proto

package spacemeshv2alpha1

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

type NetworkInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NetworkInfoRequest) Reset() {
	*x = NetworkInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInfoRequest) ProtoMessage() {}

func (x *NetworkInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInfoRequest.ProtoReflect.Descriptor instead.
func (*NetworkInfoRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_network_proto_rawDescGZIP(), []int{0}
}

type NetworkInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GenesisTime           uint64 `protobuf:"varint,1,opt,name=genesis_time,json=genesisTime,proto3" json:"genesis_time,omitempty"`            // network genesis time as unix epoch time
	EpochNumLayers        uint32 `protobuf:"varint,2,opt,name=epoch_num_layers,json=epochNumLayers,proto3" json:"epoch_num_layers,omitempty"` // number of layers per epoch
	LayerDuration         uint32 `protobuf:"varint,3,opt,name=layer_duration,json=layerDuration,proto3" json:"layer_duration,omitempty"`      // layer duration, in seconds
	GenesisId             []byte `protobuf:"bytes,4,opt,name=genesis_id,json=genesisId,proto3" json:"genesis_id,omitempty"`
	Hrp                   string `protobuf:"bytes,5,opt,name=hrp,proto3" json:"hrp,omitempty"`
	FirstGenesis          uint32 `protobuf:"varint,7,opt,name=first_genesis,json=firstGenesis,proto3" json:"first_genesis,omitempty"`                              // first effective genesis layer
	EffectiveGenesisLayer uint32 `protobuf:"varint,8,opt,name=effective_genesis_layer,json=effectiveGenesisLayer,proto3" json:"effective_genesis_layer,omitempty"` // effective genesis layer, i.e., first layer after genesis initialization period
	LayersPerEpoch        uint32 `protobuf:"varint,9,opt,name=layers_per_epoch,json=layersPerEpoch,proto3" json:"layers_per_epoch,omitempty"`                      // number of layers per epoch
}

func (x *NetworkInfoResponse) Reset() {
	*x = NetworkInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v2alpha1_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInfoResponse) ProtoMessage() {}

func (x *NetworkInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2alpha1_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInfoResponse.ProtoReflect.Descriptor instead.
func (*NetworkInfoResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2alpha1_network_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkInfoResponse) GetGenesisTime() uint64 {
	if x != nil {
		return x.GenesisTime
	}
	return 0
}

func (x *NetworkInfoResponse) GetEpochNumLayers() uint32 {
	if x != nil {
		return x.EpochNumLayers
	}
	return 0
}

func (x *NetworkInfoResponse) GetLayerDuration() uint32 {
	if x != nil {
		return x.LayerDuration
	}
	return 0
}

func (x *NetworkInfoResponse) GetGenesisId() []byte {
	if x != nil {
		return x.GenesisId
	}
	return nil
}

func (x *NetworkInfoResponse) GetHrp() string {
	if x != nil {
		return x.Hrp
	}
	return ""
}

func (x *NetworkInfoResponse) GetFirstGenesis() uint32 {
	if x != nil {
		return x.FirstGenesis
	}
	return 0
}

func (x *NetworkInfoResponse) GetEffectiveGenesisLayer() uint32 {
	if x != nil {
		return x.EffectiveGenesisLayer
	}
	return 0
}

func (x *NetworkInfoResponse) GetLayersPerEpoch() uint32 {
	if x != nil {
		return x.LayersPerEpoch
	}
	return 0
}

var File_spacemesh_v2alpha1_network_proto protoreflect.FileDescriptor

var file_spacemesh_v2alpha1_network_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x14, 0x0a, 0x12, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xc1, 0x02, 0x0a,
	0x13, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x67, 0x65, 0x6e, 0x65,
	0x73, 0x69, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x70, 0x6f, 0x63, 0x68,
	0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x4e, 0x75, 0x6d, 0x4c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65,
	0x73, 0x69, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x67, 0x65,
	0x6e, 0x65, 0x73, 0x69, 0x73, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x72, 0x70, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x68, 0x72, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x12, 0x36,
	0x0a, 0x17, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x65,
	0x73, 0x69, 0x73, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x15, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69,
	0x73, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x5f, 0x70, 0x65, 0x72, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x50, 0x65, 0x72, 0x45, 0x70, 0x6f, 0x63, 0x68,
	0x32, 0x69, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x57, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xdb, 0x01, 0x0a, 0x16,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x12, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x12,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xe2, 0x02, 0x1e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x3a,
	0x3a, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_spacemesh_v2alpha1_network_proto_rawDescOnce sync.Once
	file_spacemesh_v2alpha1_network_proto_rawDescData = file_spacemesh_v2alpha1_network_proto_rawDesc
)

func file_spacemesh_v2alpha1_network_proto_rawDescGZIP() []byte {
	file_spacemesh_v2alpha1_network_proto_rawDescOnce.Do(func() {
		file_spacemesh_v2alpha1_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacemesh_v2alpha1_network_proto_rawDescData)
	})
	return file_spacemesh_v2alpha1_network_proto_rawDescData
}

var file_spacemesh_v2alpha1_network_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_spacemesh_v2alpha1_network_proto_goTypes = []interface{}{
	(*NetworkInfoRequest)(nil),  // 0: spacemesh.v2alpha1.NetworkInfoRequest
	(*NetworkInfoResponse)(nil), // 1: spacemesh.v2alpha1.NetworkInfoResponse
}
var file_spacemesh_v2alpha1_network_proto_depIdxs = []int32{
	0, // 0: spacemesh.v2alpha1.NetworkService.Info:input_type -> spacemesh.v2alpha1.NetworkInfoRequest
	1, // 1: spacemesh.v2alpha1.NetworkService.Info:output_type -> spacemesh.v2alpha1.NetworkInfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spacemesh_v2alpha1_network_proto_init() }
func file_spacemesh_v2alpha1_network_proto_init() {
	if File_spacemesh_v2alpha1_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spacemesh_v2alpha1_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInfoRequest); i {
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
		file_spacemesh_v2alpha1_network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInfoResponse); i {
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
			RawDescriptor: file_spacemesh_v2alpha1_network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spacemesh_v2alpha1_network_proto_goTypes,
		DependencyIndexes: file_spacemesh_v2alpha1_network_proto_depIdxs,
		MessageInfos:      file_spacemesh_v2alpha1_network_proto_msgTypes,
	}.Build()
	File_spacemesh_v2alpha1_network_proto = out.File
	file_spacemesh_v2alpha1_network_proto_rawDesc = nil
	file_spacemesh_v2alpha1_network_proto_goTypes = nil
	file_spacemesh_v2alpha1_network_proto_depIdxs = nil
}
