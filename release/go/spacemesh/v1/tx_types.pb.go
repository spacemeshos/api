// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: spacemesh/v1/tx_types.proto

package spacemeshv1

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
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

type TransactionState_TransactionState int32

const (
	TransactionState_TRANSACTION_STATE_UNSPECIFIED        TransactionState_TransactionState = 0 // default state
	TransactionState_TRANSACTION_STATE_REJECTED           TransactionState_TransactionState = 1 // rejected from mempool due to, e.g., invalid syntax
	TransactionState_TRANSACTION_STATE_INSUFFICIENT_FUNDS TransactionState_TransactionState = 2 // rejected from mempool by funds check
	TransactionState_TRANSACTION_STATE_CONFLICTING        TransactionState_TransactionState = 3 // rejected from mempool due to conflicting counter
	TransactionState_TRANSACTION_STATE_MEMPOOL            TransactionState_TransactionState = 4 // in mempool but not on the mesh yet
	TransactionState_TRANSACTION_STATE_MESH               TransactionState_TransactionState = 5 // submitted to the mesh
	TransactionState_TRANSACTION_STATE_PROCESSED          TransactionState_TransactionState = 6 // processed by STF; check Receipt for success or failure
	TransactionState_TRANSACTION_STATE_INEFFECTUAL        TransactionState_TransactionState = 7 // removed from mempool and will be forgotten and never executed
)

// Enum value maps for TransactionState_TransactionState.
var (
	TransactionState_TransactionState_name = map[int32]string{
		0: "TRANSACTION_STATE_UNSPECIFIED",
		1: "TRANSACTION_STATE_REJECTED",
		2: "TRANSACTION_STATE_INSUFFICIENT_FUNDS",
		3: "TRANSACTION_STATE_CONFLICTING",
		4: "TRANSACTION_STATE_MEMPOOL",
		5: "TRANSACTION_STATE_MESH",
		6: "TRANSACTION_STATE_PROCESSED",
		7: "TRANSACTION_STATE_INEFFECTUAL",
	}
	TransactionState_TransactionState_value = map[string]int32{
		"TRANSACTION_STATE_UNSPECIFIED":        0,
		"TRANSACTION_STATE_REJECTED":           1,
		"TRANSACTION_STATE_INSUFFICIENT_FUNDS": 2,
		"TRANSACTION_STATE_CONFLICTING":        3,
		"TRANSACTION_STATE_MEMPOOL":            4,
		"TRANSACTION_STATE_MESH":               5,
		"TRANSACTION_STATE_PROCESSED":          6,
		"TRANSACTION_STATE_INEFFECTUAL":        7,
	}
)

func (x TransactionState_TransactionState) Enum() *TransactionState_TransactionState {
	p := new(TransactionState_TransactionState)
	*p = x
	return p
}

func (x TransactionState_TransactionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionState_TransactionState) Descriptor() protoreflect.EnumDescriptor {
	return file_spacemesh_v1_tx_types_proto_enumTypes[0].Descriptor()
}

func (TransactionState_TransactionState) Type() protoreflect.EnumType {
	return &file_spacemesh_v1_tx_types_proto_enumTypes[0]
}

func (x TransactionState_TransactionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionState_TransactionState.Descriptor instead.
func (TransactionState_TransactionState) EnumDescriptor() ([]byte, []int) {
	return file_spacemesh_v1_tx_types_proto_rawDescGZIP(), []int{7, 0}
}

type TransactionResult_Status int32

const (
	TransactionResult_SUCCESS TransactionResult_Status = 0
	TransactionResult_FAILURE TransactionResult_Status = 1
	TransactionResult_INVALID TransactionResult_Status = 2
)

// Enum value maps for TransactionResult_Status.
var (
	TransactionResult_Status_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILURE",
		2: "INVALID",
	}
	TransactionResult_Status_value = map[string]int32{
		"SUCCESS": 0,
		"FAILURE": 1,
		"INVALID": 2,
	}
)

func (x TransactionResult_Status) Enum() *TransactionResult_Status {
	p := new(TransactionResult_Status)
	*p = x
	return p
}

func (x TransactionResult_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionResult_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_spacemesh_v1_tx_types_proto_enumTypes[1].Descriptor()
}

func (TransactionResult_Status) Type() protoreflect.EnumType {
	return &file_spacemesh_v1_tx_types_proto_enumTypes[1]
}

func (x TransactionResult_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionResult_Status.Descriptor instead.
func (TransactionResult_Status) EnumDescriptor() ([]byte, []int) {
	return file_spacemesh_v1_tx_types_proto_rawDescGZIP(), []int{9, 0}
}

type TransactionsIds struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TransactionId []*TransactionId       `protobuf:"bytes,1,rep,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransactionsIds) Reset() {
	*x = TransactionsIds{}
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionsIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsIds) ProtoMessage() {}

func (x *TransactionsIds) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsIds.ProtoReflect.Descriptor instead.
func (*TransactionsIds) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_tx_types_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionsIds) GetTransactionId() []*TransactionId {
	if x != nil {
		return x.TransactionId
	}
	return nil
}

type SubmitTransactionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Transaction   []byte                 `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"` // signed binary transaction
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitTransactionRequest) Reset() {
	*x = SubmitTransactionRequest{}
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitTransactionRequest) ProtoMessage() {}

func (x *SubmitTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitTransactionRequest.ProtoReflect.Descriptor instead.
func (*SubmitTransactionRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_tx_types_proto_rawDescGZIP(), []int{1}
}

func (x *SubmitTransactionRequest) GetTransaction() []byte {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type SubmitTransactionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        *status.Status         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Txstate       *TransactionState      `protobuf:"bytes,2,opt,name=txstate,proto3" json:"txstate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitTransactionResponse) Reset() {
	*x = SubmitTransactionResponse{}
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitTransactionResponse) ProtoMessage() {}

func (x *SubmitTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitTransactionResponse.ProtoReflect.Descriptor instead.
func (*SubmitTransactionResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_tx_types_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitTransactionResponse) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *SubmitTransactionResponse) GetTxstate() *TransactionState {
	if x != nil {
		return x.Txstate
	}
	return nil
}

type TransactionsStateRequest struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	TransactionId       []*TransactionId       `protobuf:"bytes,1,rep,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	IncludeTransactions bool                   `protobuf:"varint,2,opt,name=include_transactions,json=includeTransactions,proto3" json:"include_transactions,omitempty"` // when true response will include matching transactions in addition to state
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *TransactionsStateRequest) Reset() {
	*x = TransactionsStateRequest{}
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionsStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsStateRequest) ProtoMessage() {}

func (x *TransactionsStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsStateRequest.ProtoReflect.Descriptor instead.
func (*TransactionsStateRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_tx_types_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionsStateRequest) GetTransactionId() []*TransactionId {
	if x != nil {
		return x.TransactionId
	}
	return nil
}

func (x *TransactionsStateRequest) GetIncludeTransactions() bool {
	if x != nil {
		return x.IncludeTransactions
	}
	return false
}

type TransactionsStateResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	TransactionsState []*TransactionState    `protobuf:"bytes,1,rep,name=transactions_state,json=transactionsState,proto3" json:"transactions_state,omitempty"`
	Transactions      []*Transaction         `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *TransactionsStateResponse) Reset() {
	*x = TransactionsStateResponse{}
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionsStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsStateResponse) ProtoMessage() {}

func (x *TransactionsStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsStateResponse.ProtoReflect.Descriptor instead.
func (*TransactionsStateResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_tx_types_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionsStateResponse) GetTransactionsState() []*TransactionState {
	if x != nil {
		return x.TransactionsState
	}
	return nil
}

func (x *TransactionsStateResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type TransactionsStateStreamRequest struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	TransactionId       []*TransactionId       `protobuf:"bytes,1,rep,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	IncludeTransactions bool                   `protobuf:"varint,2,opt,name=include_transactions,json=includeTransactions,proto3" json:"include_transactions,omitempty"` // when true response will include matching transactions in addition to state
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *TransactionsStateStreamRequest) Reset() {
	*x = TransactionsStateStreamRequest{}
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionsStateStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsStateStreamRequest) ProtoMessage() {}

func (x *TransactionsStateStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsStateStreamRequest.ProtoReflect.Descriptor instead.
func (*TransactionsStateStreamRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_tx_types_proto_rawDescGZIP(), []int{5}
}

func (x *TransactionsStateStreamRequest) GetTransactionId() []*TransactionId {
	if x != nil {
		return x.TransactionId
	}
	return nil
}

func (x *TransactionsStateStreamRequest) GetIncludeTransactions() bool {
	if x != nil {
		return x.IncludeTransactions
	}
	return false
}

type TransactionsStateStreamResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	TransactionState *TransactionState      `protobuf:"bytes,1,opt,name=transaction_state,json=transactionState,proto3" json:"transaction_state,omitempty"`
	Transaction      *Transaction           `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *TransactionsStateStreamResponse) Reset() {
	*x = TransactionsStateStreamResponse{}
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionsStateStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsStateStreamResponse) ProtoMessage() {}

func (x *TransactionsStateStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsStateStreamResponse.ProtoReflect.Descriptor instead.
func (*TransactionsStateStreamResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_tx_types_proto_rawDescGZIP(), []int{6}
}

func (x *TransactionsStateStreamResponse) GetTransactionState() *TransactionState {
	if x != nil {
		return x.TransactionState
	}
	return nil
}

func (x *TransactionsStateStreamResponse) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

// TransactionState is the "journey" of a tx from mempool to block inclusion to
// mesh to STF processing. To know whether or not the tx actually succeeded,
// and its side effects, check the Receipt in the GlobalStateService.
type TransactionState struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	Id            *TransactionId                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State         TransactionState_TransactionState `protobuf:"varint,2,opt,name=state,proto3,enum=spacemesh.v1.TransactionState_TransactionState" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransactionState) Reset() {
	*x = TransactionState{}
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionState) ProtoMessage() {}

func (x *TransactionState) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionState.ProtoReflect.Descriptor instead.
func (*TransactionState) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_tx_types_proto_rawDescGZIP(), []int{7}
}

func (x *TransactionState) GetId() *TransactionId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *TransactionState) GetState() TransactionState_TransactionState {
	if x != nil {
		return x.State
	}
	return TransactionState_TRANSACTION_STATE_UNSPECIFIED
}

// TransactionResultsRequest request object for results stream.
type TransactionResultsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// id is filter by transaction id.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// address is a filter by account address, it could be principal or any affected address.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// start streaming from this layer. if 0 - stream will start from genesis.
	Start uint32 `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	// end streaming at this layer. if 0 - stream till the latest available layer.
	End uint32 `protobuf:"varint,4,opt,name=end,proto3" json:"end,omitempty"`
	// watch live data.
	Watch         bool `protobuf:"varint,5,opt,name=watch,proto3" json:"watch,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransactionResultsRequest) Reset() {
	*x = TransactionResultsRequest{}
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionResultsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResultsRequest) ProtoMessage() {}

func (x *TransactionResultsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResultsRequest.ProtoReflect.Descriptor instead.
func (*TransactionResultsRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_tx_types_proto_rawDescGZIP(), []int{8}
}

func (x *TransactionResultsRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *TransactionResultsRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *TransactionResultsRequest) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *TransactionResultsRequest) GetEnd() uint32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *TransactionResultsRequest) GetWatch() bool {
	if x != nil {
		return x.Watch
	}
	return false
}

type TransactionResult struct {
	state            protoimpl.MessageState   `protogen:"open.v1"`
	Tx               *Transaction             `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	Status           TransactionResult_Status `protobuf:"varint,2,opt,name=status,proto3,enum=spacemesh.v1.TransactionResult_Status" json:"status,omitempty"`
	Message          string                   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	GasConsumed      uint64                   `protobuf:"varint,4,opt,name=gas_consumed,json=gasConsumed,proto3" json:"gas_consumed,omitempty"`
	Fee              uint64                   `protobuf:"varint,5,opt,name=fee,proto3" json:"fee,omitempty"`
	Block            []byte                   `protobuf:"bytes,6,opt,name=block,proto3" json:"block,omitempty"`
	Layer            uint32                   `protobuf:"varint,7,opt,name=layer,proto3" json:"layer,omitempty"`
	TouchedAddresses []string                 `protobuf:"bytes,8,rep,name=touched_addresses,json=touchedAddresses,proto3" json:"touched_addresses,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *TransactionResult) Reset() {
	*x = TransactionResult{}
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResult) ProtoMessage() {}

func (x *TransactionResult) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResult.ProtoReflect.Descriptor instead.
func (*TransactionResult) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_tx_types_proto_rawDescGZIP(), []int{9}
}

func (x *TransactionResult) GetTx() *Transaction {
	if x != nil {
		return x.Tx
	}
	return nil
}

func (x *TransactionResult) GetStatus() TransactionResult_Status {
	if x != nil {
		return x.Status
	}
	return TransactionResult_SUCCESS
}

func (x *TransactionResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TransactionResult) GetGasConsumed() uint64 {
	if x != nil {
		return x.GasConsumed
	}
	return 0
}

func (x *TransactionResult) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *TransactionResult) GetBlock() []byte {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *TransactionResult) GetLayer() uint32 {
	if x != nil {
		return x.Layer
	}
	return 0
}

func (x *TransactionResult) GetTouchedAddresses() []string {
	if x != nil {
		return x.TouchedAddresses
	}
	return nil
}

type ParseTransactionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Transaction   []byte                 `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"` // signed binary transaction
	Verify        bool                   `protobuf:"varint,2,opt,name=verify,proto3" json:"verify,omitempty"`          // if true signature verification will be executed
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ParseTransactionRequest) Reset() {
	*x = ParseTransactionRequest{}
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParseTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseTransactionRequest) ProtoMessage() {}

func (x *ParseTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseTransactionRequest.ProtoReflect.Descriptor instead.
func (*ParseTransactionRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_tx_types_proto_rawDescGZIP(), []int{10}
}

func (x *ParseTransactionRequest) GetTransaction() []byte {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *ParseTransactionRequest) GetVerify() bool {
	if x != nil {
		return x.Verify
	}
	return false
}

type ParseTransactionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tx            *Transaction           `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ParseTransactionResponse) Reset() {
	*x = ParseTransactionResponse{}
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParseTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseTransactionResponse) ProtoMessage() {}

func (x *ParseTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_tx_types_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseTransactionResponse.ProtoReflect.Descriptor instead.
func (*ParseTransactionResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_tx_types_proto_rawDescGZIP(), []int{11}
}

func (x *ParseTransactionResponse) GetTx() *Transaction {
	if x != nil {
		return x.Tx
	}
	return nil
}

var File_spacemesh_v1_tx_types_proto protoreflect.FileDescriptor

const file_spacemesh_v1_tx_types_proto_rawDesc = "" +
	"\n" +
	"\x1bspacemesh/v1/tx_types.proto\x12\fspacemesh.v1\x1a\x17google/rpc/status.proto\x1a\x18spacemesh/v1/types.proto\"U\n" +
	"\x0fTransactionsIds\x12B\n" +
	"\x0etransaction_id\x18\x01 \x03(\v2\x1b.spacemesh.v1.TransactionIdR\rtransactionId\"<\n" +
	"\x18SubmitTransactionRequest\x12 \n" +
	"\vtransaction\x18\x01 \x01(\fR\vtransaction\"\x81\x01\n" +
	"\x19SubmitTransactionResponse\x12*\n" +
	"\x06status\x18\x01 \x01(\v2\x12.google.rpc.StatusR\x06status\x128\n" +
	"\atxstate\x18\x02 \x01(\v2\x1e.spacemesh.v1.TransactionStateR\atxstate\"\x91\x01\n" +
	"\x18TransactionsStateRequest\x12B\n" +
	"\x0etransaction_id\x18\x01 \x03(\v2\x1b.spacemesh.v1.TransactionIdR\rtransactionId\x121\n" +
	"\x14include_transactions\x18\x02 \x01(\bR\x13includeTransactions\"\xa9\x01\n" +
	"\x19TransactionsStateResponse\x12M\n" +
	"\x12transactions_state\x18\x01 \x03(\v2\x1e.spacemesh.v1.TransactionStateR\x11transactionsState\x12=\n" +
	"\ftransactions\x18\x02 \x03(\v2\x19.spacemesh.v1.TransactionR\ftransactions\"\x97\x01\n" +
	"\x1eTransactionsStateStreamRequest\x12B\n" +
	"\x0etransaction_id\x18\x01 \x03(\v2\x1b.spacemesh.v1.TransactionIdR\rtransactionId\x121\n" +
	"\x14include_transactions\x18\x02 \x01(\bR\x13includeTransactions\"\xab\x01\n" +
	"\x1fTransactionsStateStreamResponse\x12K\n" +
	"\x11transaction_state\x18\x01 \x01(\v2\x1e.spacemesh.v1.TransactionStateR\x10transactionState\x12;\n" +
	"\vtransaction\x18\x02 \x01(\v2\x19.spacemesh.v1.TransactionR\vtransaction\"\xaa\x03\n" +
	"\x10TransactionState\x12+\n" +
	"\x02id\x18\x01 \x01(\v2\x1b.spacemesh.v1.TransactionIdR\x02id\x12E\n" +
	"\x05state\x18\x02 \x01(\x0e2/.spacemesh.v1.TransactionState.TransactionStateR\x05state\"\xa1\x02\n" +
	"\x10TransactionState\x12!\n" +
	"\x1dTRANSACTION_STATE_UNSPECIFIED\x10\x00\x12\x1e\n" +
	"\x1aTRANSACTION_STATE_REJECTED\x10\x01\x12(\n" +
	"$TRANSACTION_STATE_INSUFFICIENT_FUNDS\x10\x02\x12!\n" +
	"\x1dTRANSACTION_STATE_CONFLICTING\x10\x03\x12\x1d\n" +
	"\x19TRANSACTION_STATE_MEMPOOL\x10\x04\x12\x1a\n" +
	"\x16TRANSACTION_STATE_MESH\x10\x05\x12\x1f\n" +
	"\x1bTRANSACTION_STATE_PROCESSED\x10\x06\x12!\n" +
	"\x1dTRANSACTION_STATE_INEFFECTUAL\x10\a\"\x83\x01\n" +
	"\x19TransactionResultsRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\fR\x02id\x12\x18\n" +
	"\aaddress\x18\x02 \x01(\tR\aaddress\x12\x14\n" +
	"\x05start\x18\x03 \x01(\rR\x05start\x12\x10\n" +
	"\x03end\x18\x04 \x01(\rR\x03end\x12\x14\n" +
	"\x05watch\x18\x05 \x01(\bR\x05watch\"\xd7\x02\n" +
	"\x11TransactionResult\x12)\n" +
	"\x02tx\x18\x01 \x01(\v2\x19.spacemesh.v1.TransactionR\x02tx\x12>\n" +
	"\x06status\x18\x02 \x01(\x0e2&.spacemesh.v1.TransactionResult.StatusR\x06status\x12\x18\n" +
	"\amessage\x18\x03 \x01(\tR\amessage\x12!\n" +
	"\fgas_consumed\x18\x04 \x01(\x04R\vgasConsumed\x12\x10\n" +
	"\x03fee\x18\x05 \x01(\x04R\x03fee\x12\x14\n" +
	"\x05block\x18\x06 \x01(\fR\x05block\x12\x14\n" +
	"\x05layer\x18\a \x01(\rR\x05layer\x12+\n" +
	"\x11touched_addresses\x18\b \x03(\tR\x10touchedAddresses\"/\n" +
	"\x06Status\x12\v\n" +
	"\aSUCCESS\x10\x00\x12\v\n" +
	"\aFAILURE\x10\x01\x12\v\n" +
	"\aINVALID\x10\x02\"S\n" +
	"\x17ParseTransactionRequest\x12 \n" +
	"\vtransaction\x18\x01 \x01(\fR\vtransaction\x12\x16\n" +
	"\x06verify\x18\x02 \x01(\bR\x06verify\"E\n" +
	"\x18ParseTransactionResponse\x12)\n" +
	"\x02tx\x18\x01 \x01(\v2\x19.spacemesh.v1.TransactionR\x02txB\xb1\x01\n" +
	"\x10com.spacemesh.v1B\fTxTypesProtoP\x01Z>github.com/spacemeshos/api/release/go/spacemesh/v1;spacemeshv1\xa2\x02\x03SXX\xaa\x02\fSpacemesh.V1\xca\x02\fSpacemesh\\V1\xe2\x02\x18Spacemesh\\V1\\GPBMetadata\xea\x02\rSpacemesh::V1b\x06proto3"

var (
	file_spacemesh_v1_tx_types_proto_rawDescOnce sync.Once
	file_spacemesh_v1_tx_types_proto_rawDescData []byte
)

func file_spacemesh_v1_tx_types_proto_rawDescGZIP() []byte {
	file_spacemesh_v1_tx_types_proto_rawDescOnce.Do(func() {
		file_spacemesh_v1_tx_types_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spacemesh_v1_tx_types_proto_rawDesc), len(file_spacemesh_v1_tx_types_proto_rawDesc)))
	})
	return file_spacemesh_v1_tx_types_proto_rawDescData
}

var file_spacemesh_v1_tx_types_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_spacemesh_v1_tx_types_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_spacemesh_v1_tx_types_proto_goTypes = []any{
	(TransactionState_TransactionState)(0),  // 0: spacemesh.v1.TransactionState.TransactionState
	(TransactionResult_Status)(0),           // 1: spacemesh.v1.TransactionResult.Status
	(*TransactionsIds)(nil),                 // 2: spacemesh.v1.TransactionsIds
	(*SubmitTransactionRequest)(nil),        // 3: spacemesh.v1.SubmitTransactionRequest
	(*SubmitTransactionResponse)(nil),       // 4: spacemesh.v1.SubmitTransactionResponse
	(*TransactionsStateRequest)(nil),        // 5: spacemesh.v1.TransactionsStateRequest
	(*TransactionsStateResponse)(nil),       // 6: spacemesh.v1.TransactionsStateResponse
	(*TransactionsStateStreamRequest)(nil),  // 7: spacemesh.v1.TransactionsStateStreamRequest
	(*TransactionsStateStreamResponse)(nil), // 8: spacemesh.v1.TransactionsStateStreamResponse
	(*TransactionState)(nil),                // 9: spacemesh.v1.TransactionState
	(*TransactionResultsRequest)(nil),       // 10: spacemesh.v1.TransactionResultsRequest
	(*TransactionResult)(nil),               // 11: spacemesh.v1.TransactionResult
	(*ParseTransactionRequest)(nil),         // 12: spacemesh.v1.ParseTransactionRequest
	(*ParseTransactionResponse)(nil),        // 13: spacemesh.v1.ParseTransactionResponse
	(*TransactionId)(nil),                   // 14: spacemesh.v1.TransactionId
	(*status.Status)(nil),                   // 15: google.rpc.Status
	(*Transaction)(nil),                     // 16: spacemesh.v1.Transaction
}
var file_spacemesh_v1_tx_types_proto_depIdxs = []int32{
	14, // 0: spacemesh.v1.TransactionsIds.transaction_id:type_name -> spacemesh.v1.TransactionId
	15, // 1: spacemesh.v1.SubmitTransactionResponse.status:type_name -> google.rpc.Status
	9,  // 2: spacemesh.v1.SubmitTransactionResponse.txstate:type_name -> spacemesh.v1.TransactionState
	14, // 3: spacemesh.v1.TransactionsStateRequest.transaction_id:type_name -> spacemesh.v1.TransactionId
	9,  // 4: spacemesh.v1.TransactionsStateResponse.transactions_state:type_name -> spacemesh.v1.TransactionState
	16, // 5: spacemesh.v1.TransactionsStateResponse.transactions:type_name -> spacemesh.v1.Transaction
	14, // 6: spacemesh.v1.TransactionsStateStreamRequest.transaction_id:type_name -> spacemesh.v1.TransactionId
	9,  // 7: spacemesh.v1.TransactionsStateStreamResponse.transaction_state:type_name -> spacemesh.v1.TransactionState
	16, // 8: spacemesh.v1.TransactionsStateStreamResponse.transaction:type_name -> spacemesh.v1.Transaction
	14, // 9: spacemesh.v1.TransactionState.id:type_name -> spacemesh.v1.TransactionId
	0,  // 10: spacemesh.v1.TransactionState.state:type_name -> spacemesh.v1.TransactionState.TransactionState
	16, // 11: spacemesh.v1.TransactionResult.tx:type_name -> spacemesh.v1.Transaction
	1,  // 12: spacemesh.v1.TransactionResult.status:type_name -> spacemesh.v1.TransactionResult.Status
	16, // 13: spacemesh.v1.ParseTransactionResponse.tx:type_name -> spacemesh.v1.Transaction
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_spacemesh_v1_tx_types_proto_init() }
func file_spacemesh_v1_tx_types_proto_init() {
	if File_spacemesh_v1_tx_types_proto != nil {
		return
	}
	file_spacemesh_v1_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spacemesh_v1_tx_types_proto_rawDesc), len(file_spacemesh_v1_tx_types_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spacemesh_v1_tx_types_proto_goTypes,
		DependencyIndexes: file_spacemesh_v1_tx_types_proto_depIdxs,
		EnumInfos:         file_spacemesh_v1_tx_types_proto_enumTypes,
		MessageInfos:      file_spacemesh_v1_tx_types_proto_msgTypes,
	}.Build()
	File_spacemesh_v1_tx_types_proto = out.File
	file_spacemesh_v1_tx_types_proto_goTypes = nil
	file_spacemesh_v1_tx_types_proto_depIdxs = nil
}
