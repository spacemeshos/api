// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: spacemesh/v2beta1/account.proto

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

type Account struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`     // account public address
	Current       *AccountState          `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`     // current state
	Projected     *AccountState          `protobuf:"bytes,3,opt,name=projected,proto3" json:"projected,omitempty"` // projected state (includes pending txs)
	Template      string                 `protobuf:"bytes,4,opt,name=template,proto3" json:"template,omitempty"`   // account template address
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Account) Reset() {
	*x = Account{}
	mi := &file_spacemesh_v2beta1_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_account_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Account) GetCurrent() *AccountState {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *Account) GetProjected() *AccountState {
	if x != nil {
		return x.Projected
	}
	return nil
}

func (x *Account) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

type AccountState struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Counter       uint64                 `protobuf:"varint,1,opt,name=counter,proto3" json:"counter,omitempty"` // aka nonce
	Balance       uint64                 `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"` // account balance in smidge
	Layer         uint32                 `protobuf:"varint,3,opt,name=layer,proto3" json:"layer,omitempty"`     // account balance as of layer X
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccountState) Reset() {
	*x = AccountState{}
	mi := &file_spacemesh_v2beta1_account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountState) ProtoMessage() {}

func (x *AccountState) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_account_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountState.ProtoReflect.Descriptor instead.
func (*AccountState) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_account_proto_rawDescGZIP(), []int{1}
}

func (x *AccountState) GetCounter() uint64 {
	if x != nil {
		return x.Counter
	}
	return 0
}

func (x *AccountState) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *AccountState) GetLayer() uint32 {
	if x != nil {
		return x.Layer
	}
	return 0
}

type AccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Addresses     []string               `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"` // list of account addresses
	Offset        uint64                 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`      // adjusts the starting point for data
	Limit         uint64                 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`        // specifies max number of items to fetch// bech32 format including HRP
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccountRequest) Reset() {
	*x = AccountRequest{}
	mi := &file_spacemesh_v2beta1_account_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountRequest) ProtoMessage() {}

func (x *AccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_account_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountRequest.ProtoReflect.Descriptor instead.
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_account_proto_rawDescGZIP(), []int{2}
}

func (x *AccountRequest) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *AccountRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *AccountRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type AccountList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Accounts      []*Account             `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"` // list of accounts
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccountList) Reset() {
	*x = AccountList{}
	mi := &file_spacemesh_v2beta1_account_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountList) ProtoMessage() {}

func (x *AccountList) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v2beta1_account_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountList.ProtoReflect.Descriptor instead.
func (*AccountList) Descriptor() ([]byte, []int) {
	return file_spacemesh_v2beta1_account_proto_rawDescGZIP(), []int{3}
}

func (x *AccountList) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

var File_spacemesh_v2beta1_account_proto protoreflect.FileDescriptor

const file_spacemesh_v2beta1_account_proto_rawDesc = "" +
	"\n" +
	"\x1fspacemesh/v2beta1/account.proto\x12\x11spacemesh.v2beta1\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/api/visibility.proto\"\xb9\x01\n" +
	"\aAccount\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\tR\aaddress\x129\n" +
	"\acurrent\x18\x02 \x01(\v2\x1f.spacemesh.v2beta1.AccountStateR\acurrent\x12=\n" +
	"\tprojected\x18\x03 \x01(\v2\x1f.spacemesh.v2beta1.AccountStateR\tprojected\x12\x1a\n" +
	"\btemplate\x18\x04 \x01(\tR\btemplate\"X\n" +
	"\fAccountState\x12\x18\n" +
	"\acounter\x18\x01 \x01(\x04R\acounter\x12\x18\n" +
	"\abalance\x18\x02 \x01(\x04R\abalance\x12\x14\n" +
	"\x05layer\x18\x03 \x01(\rR\x05layer\"\\\n" +
	"\x0eAccountRequest\x12\x1c\n" +
	"\taddresses\x18\x01 \x03(\tR\taddresses\x12\x16\n" +
	"\x06offset\x18\x02 \x01(\x04R\x06offset\x12\x14\n" +
	"\x05limit\x18\x03 \x01(\x04R\x05limit\"E\n" +
	"\vAccountList\x126\n" +
	"\baccounts\x18\x01 \x03(\v2\x1a.spacemesh.v2beta1.AccountR\baccounts2\x9c\x01\n" +
	"\x0eAccountService\x12y\n" +
	"\x04List\x12!.spacemesh.v2beta1.AccountRequest\x1a\x1e.spacemesh.v2beta1.AccountList\".\x82\xd3\xe4\x93\x02(\x12&/spacemesh.v2beta1.AccountService/List\x1a\x0f\xfa\xd2\xe4\x93\x02\t\x12\av2beta1B\xd4\x01\n" +
	"\x15com.spacemesh.v2beta1B\fAccountProtoP\x01ZHgithub.com/spacemeshos/api/release/go/spacemesh/v2beta1;spacemeshv2beta1\xa2\x02\x03SXX\xaa\x02\x11Spacemesh.V2beta1\xca\x02\x11Spacemesh\\V2beta1\xe2\x02\x1dSpacemesh\\V2beta1\\GPBMetadata\xea\x02\x12Spacemesh::V2beta1b\x06proto3"

var (
	file_spacemesh_v2beta1_account_proto_rawDescOnce sync.Once
	file_spacemesh_v2beta1_account_proto_rawDescData []byte
)

func file_spacemesh_v2beta1_account_proto_rawDescGZIP() []byte {
	file_spacemesh_v2beta1_account_proto_rawDescOnce.Do(func() {
		file_spacemesh_v2beta1_account_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spacemesh_v2beta1_account_proto_rawDesc), len(file_spacemesh_v2beta1_account_proto_rawDesc)))
	})
	return file_spacemesh_v2beta1_account_proto_rawDescData
}

var file_spacemesh_v2beta1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spacemesh_v2beta1_account_proto_goTypes = []any{
	(*Account)(nil),        // 0: spacemesh.v2beta1.Account
	(*AccountState)(nil),   // 1: spacemesh.v2beta1.AccountState
	(*AccountRequest)(nil), // 2: spacemesh.v2beta1.AccountRequest
	(*AccountList)(nil),    // 3: spacemesh.v2beta1.AccountList
}
var file_spacemesh_v2beta1_account_proto_depIdxs = []int32{
	1, // 0: spacemesh.v2beta1.Account.current:type_name -> spacemesh.v2beta1.AccountState
	1, // 1: spacemesh.v2beta1.Account.projected:type_name -> spacemesh.v2beta1.AccountState
	0, // 2: spacemesh.v2beta1.AccountList.accounts:type_name -> spacemesh.v2beta1.Account
	2, // 3: spacemesh.v2beta1.AccountService.List:input_type -> spacemesh.v2beta1.AccountRequest
	3, // 4: spacemesh.v2beta1.AccountService.List:output_type -> spacemesh.v2beta1.AccountList
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_spacemesh_v2beta1_account_proto_init() }
func file_spacemesh_v2beta1_account_proto_init() {
	if File_spacemesh_v2beta1_account_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spacemesh_v2beta1_account_proto_rawDesc), len(file_spacemesh_v2beta1_account_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spacemesh_v2beta1_account_proto_goTypes,
		DependencyIndexes: file_spacemesh_v2beta1_account_proto_depIdxs,
		MessageInfos:      file_spacemesh_v2beta1_account_proto_msgTypes,
	}.Build()
	File_spacemesh_v2beta1_account_proto = out.File
	file_spacemesh_v2beta1_account_proto_goTypes = nil
	file_spacemesh_v2beta1_account_proto_depIdxs = nil
}
