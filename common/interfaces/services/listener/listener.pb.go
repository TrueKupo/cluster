// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: listener/listener.proto

package listener

import (
	response "github.com/truekupo/cluster/common/interfaces/messages/response"
	types "github.com/truekupo/cluster/common/interfaces/messages/types"
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

type AddAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountUuid string `protobuf:"bytes,1,opt,name=account_uuid,json=accountUuid,proto3" json:"account_uuid,omitempty"`
	Address     string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *AddAddressRequest) Reset() {
	*x = AddAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listener_listener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAddressRequest) ProtoMessage() {}

func (x *AddAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listener_listener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAddressRequest.ProtoReflect.Descriptor instead.
func (*AddAddressRequest) Descriptor() ([]byte, []int) {
	return file_listener_listener_proto_rawDescGZIP(), []int{0}
}

func (x *AddAddressRequest) GetAccountUuid() string {
	if x != nil {
		return x.AccountUuid
	}
	return ""
}

func (x *AddAddressRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type AddAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetStatus *response.Status `protobuf:"bytes,1,opt,name=ret_status,json=retStatus,proto3" json:"ret_status,omitempty"`
}

func (x *AddAddressResponse) Reset() {
	*x = AddAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listener_listener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAddressResponse) ProtoMessage() {}

func (x *AddAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listener_listener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAddressResponse.ProtoReflect.Descriptor instead.
func (*AddAddressResponse) Descriptor() ([]byte, []int) {
	return file_listener_listener_proto_rawDescGZIP(), []int{1}
}

func (x *AddAddressResponse) GetRetStatus() *response.Status {
	if x != nil {
		return x.RetStatus
	}
	return nil
}

type AddTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *types.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *AddTransactionRequest) Reset() {
	*x = AddTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listener_listener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTransactionRequest) ProtoMessage() {}

func (x *AddTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listener_listener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTransactionRequest.ProtoReflect.Descriptor instead.
func (*AddTransactionRequest) Descriptor() ([]byte, []int) {
	return file_listener_listener_proto_rawDescGZIP(), []int{2}
}

func (x *AddTransactionRequest) GetTransaction() *types.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type AddTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetStatus *response.Status `protobuf:"bytes,1,opt,name=ret_status,json=retStatus,proto3" json:"ret_status,omitempty"`
}

func (x *AddTransactionResponse) Reset() {
	*x = AddTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listener_listener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTransactionResponse) ProtoMessage() {}

func (x *AddTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listener_listener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTransactionResponse.ProtoReflect.Descriptor instead.
func (*AddTransactionResponse) Descriptor() ([]byte, []int) {
	return file_listener_listener_proto_rawDescGZIP(), []int{3}
}

func (x *AddTransactionResponse) GetRetStatus() *response.Status {
	if x != nil {
		return x.RetStatus
	}
	return nil
}

type TransactionsByAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	From    int32  `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	Limit   int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *TransactionsByAddressRequest) Reset() {
	*x = TransactionsByAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listener_listener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionsByAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsByAddressRequest) ProtoMessage() {}

func (x *TransactionsByAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listener_listener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsByAddressRequest.ProtoReflect.Descriptor instead.
func (*TransactionsByAddressRequest) Descriptor() ([]byte, []int) {
	return file_listener_listener_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionsByAddressRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *TransactionsByAddressRequest) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *TransactionsByAddressRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type TransactionsByAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*types.Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	RetStatus    *response.Status     `protobuf:"bytes,2,opt,name=ret_status,json=retStatus,proto3" json:"ret_status,omitempty"`
}

func (x *TransactionsByAddressResponse) Reset() {
	*x = TransactionsByAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listener_listener_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionsByAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsByAddressResponse) ProtoMessage() {}

func (x *TransactionsByAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listener_listener_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsByAddressResponse.ProtoReflect.Descriptor instead.
func (*TransactionsByAddressResponse) Descriptor() ([]byte, []int) {
	return file_listener_listener_proto_rawDescGZIP(), []int{5}
}

func (x *TransactionsByAddressResponse) GetTransactions() []*types.Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *TransactionsByAddressResponse) GetRetStatus() *response.Status {
	if x != nil {
		return x.RetStatus
	}
	return nil
}

type TransactionsByAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountUuid string `protobuf:"bytes,1,opt,name=account_uuid,json=accountUuid,proto3" json:"account_uuid,omitempty"`
	From        int32  `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	Limit       int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *TransactionsByAccountRequest) Reset() {
	*x = TransactionsByAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listener_listener_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionsByAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsByAccountRequest) ProtoMessage() {}

func (x *TransactionsByAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listener_listener_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsByAccountRequest.ProtoReflect.Descriptor instead.
func (*TransactionsByAccountRequest) Descriptor() ([]byte, []int) {
	return file_listener_listener_proto_rawDescGZIP(), []int{6}
}

func (x *TransactionsByAccountRequest) GetAccountUuid() string {
	if x != nil {
		return x.AccountUuid
	}
	return ""
}

func (x *TransactionsByAccountRequest) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *TransactionsByAccountRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type TransactionsByAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*types.Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	RetStatus    *response.Status     `protobuf:"bytes,2,opt,name=ret_status,json=retStatus,proto3" json:"ret_status,omitempty"`
}

func (x *TransactionsByAccountResponse) Reset() {
	*x = TransactionsByAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listener_listener_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionsByAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsByAccountResponse) ProtoMessage() {}

func (x *TransactionsByAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listener_listener_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsByAccountResponse.ProtoReflect.Descriptor instead.
func (*TransactionsByAccountResponse) Descriptor() ([]byte, []int) {
	return file_listener_listener_proto_rawDescGZIP(), []int{7}
}

func (x *TransactionsByAccountResponse) GetTransactions() []*types.Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *TransactionsByAccountResponse) GetRetStatus() *response.Status {
	if x != nil {
		return x.RetStatus
	}
	return nil
}

type GetTxByHashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol types.CoinCode `protobuf:"varint,1,opt,name=symbol,proto3,enum=types.CoinCode" json:"symbol,omitempty"`
	TxHash string         `protobuf:"bytes,2,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
}

func (x *GetTxByHashRequest) Reset() {
	*x = GetTxByHashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listener_listener_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxByHashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxByHashRequest) ProtoMessage() {}

func (x *GetTxByHashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listener_listener_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTxByHashRequest.ProtoReflect.Descriptor instead.
func (*GetTxByHashRequest) Descriptor() ([]byte, []int) {
	return file_listener_listener_proto_rawDescGZIP(), []int{8}
}

func (x *GetTxByHashRequest) GetSymbol() types.CoinCode {
	if x != nil {
		return x.Symbol
	}
	return types.CoinCode_BTC
}

func (x *GetTxByHashRequest) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

type GetTxByHashResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *types.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	RetStatus   *response.Status   `protobuf:"bytes,2,opt,name=ret_status,json=retStatus,proto3" json:"ret_status,omitempty"`
}

func (x *GetTxByHashResponse) Reset() {
	*x = GetTxByHashResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listener_listener_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxByHashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxByHashResponse) ProtoMessage() {}

func (x *GetTxByHashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listener_listener_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTxByHashResponse.ProtoReflect.Descriptor instead.
func (*GetTxByHashResponse) Descriptor() ([]byte, []int) {
	return file_listener_listener_proto_rawDescGZIP(), []int{9}
}

func (x *GetTxByHashResponse) GetTransaction() *types.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *GetTxByHashResponse) GetRetStatus() *response.Status {
	if x != nil {
		return x.RetStatus
	}
	return nil
}

var File_listener_listener_proto protoreflect.FileDescriptor

var file_listener_listener_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x1a, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x50, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x45, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x72, 0x65, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x09, 0x72, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4d, 0x0a, 0x15, 0x41, 0x64,
	0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x49, 0x0a, 0x16, 0x41, 0x64, 0x64,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x72, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x62, 0x0a, 0x1c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x88, 0x01, 0x0a, 0x1d, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x72, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x6b, 0x0a, 0x1c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x88, 0x01, 0x0a, 0x1d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x0a, 0x72, 0x65,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x09, 0x72, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x56, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x54, 0x78, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48,
	0x61, 0x73, 0x68, 0x22, 0x7c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x78, 0x42, 0x79, 0x48, 0x61,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2f, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x72, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0xd9, 0x03, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x41,
	0x64, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x55, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x41, 0x64,
	0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x41,
	0x64, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x26, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x78, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x42,
	0x79, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x42, 0x79, 0x48,
	0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_listener_listener_proto_rawDescOnce sync.Once
	file_listener_listener_proto_rawDescData = file_listener_listener_proto_rawDesc
)

func file_listener_listener_proto_rawDescGZIP() []byte {
	file_listener_listener_proto_rawDescOnce.Do(func() {
		file_listener_listener_proto_rawDescData = protoimpl.X.CompressGZIP(file_listener_listener_proto_rawDescData)
	})
	return file_listener_listener_proto_rawDescData
}

var file_listener_listener_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_listener_listener_proto_goTypes = []interface{}{
	(*AddAddressRequest)(nil),             // 0: listener.AddAddressRequest
	(*AddAddressResponse)(nil),            // 1: listener.AddAddressResponse
	(*AddTransactionRequest)(nil),         // 2: listener.AddTransactionRequest
	(*AddTransactionResponse)(nil),        // 3: listener.AddTransactionResponse
	(*TransactionsByAddressRequest)(nil),  // 4: listener.TransactionsByAddressRequest
	(*TransactionsByAddressResponse)(nil), // 5: listener.TransactionsByAddressResponse
	(*TransactionsByAccountRequest)(nil),  // 6: listener.TransactionsByAccountRequest
	(*TransactionsByAccountResponse)(nil), // 7: listener.TransactionsByAccountResponse
	(*GetTxByHashRequest)(nil),            // 8: listener.GetTxByHashRequest
	(*GetTxByHashResponse)(nil),           // 9: listener.GetTxByHashResponse
	(*response.Status)(nil),               // 10: response.Status
	(*types.Transaction)(nil),             // 11: types.Transaction
	(types.CoinCode)(0),                   // 12: types.CoinCode
}
var file_listener_listener_proto_depIdxs = []int32{
	10, // 0: listener.AddAddressResponse.ret_status:type_name -> response.Status
	11, // 1: listener.AddTransactionRequest.transaction:type_name -> types.Transaction
	10, // 2: listener.AddTransactionResponse.ret_status:type_name -> response.Status
	11, // 3: listener.TransactionsByAddressResponse.transactions:type_name -> types.Transaction
	10, // 4: listener.TransactionsByAddressResponse.ret_status:type_name -> response.Status
	11, // 5: listener.TransactionsByAccountResponse.transactions:type_name -> types.Transaction
	10, // 6: listener.TransactionsByAccountResponse.ret_status:type_name -> response.Status
	12, // 7: listener.GetTxByHashRequest.symbol:type_name -> types.CoinCode
	11, // 8: listener.GetTxByHashResponse.transaction:type_name -> types.Transaction
	10, // 9: listener.GetTxByHashResponse.ret_status:type_name -> response.Status
	0,  // 10: listener.ListenerService.AddAddress:input_type -> listener.AddAddressRequest
	2,  // 11: listener.ListenerService.AddTransaction:input_type -> listener.AddTransactionRequest
	4,  // 12: listener.ListenerService.TransactionsByAddress:input_type -> listener.TransactionsByAddressRequest
	6,  // 13: listener.ListenerService.TransactionsByAccount:input_type -> listener.TransactionsByAccountRequest
	8,  // 14: listener.ListenerService.GetTxByHash:input_type -> listener.GetTxByHashRequest
	1,  // 15: listener.ListenerService.AddAddress:output_type -> listener.AddAddressResponse
	3,  // 16: listener.ListenerService.AddTransaction:output_type -> listener.AddTransactionResponse
	5,  // 17: listener.ListenerService.TransactionsByAddress:output_type -> listener.TransactionsByAddressResponse
	7,  // 18: listener.ListenerService.TransactionsByAccount:output_type -> listener.TransactionsByAccountResponse
	9,  // 19: listener.ListenerService.GetTxByHash:output_type -> listener.GetTxByHashResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_listener_listener_proto_init() }
func file_listener_listener_proto_init() {
	if File_listener_listener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_listener_listener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAddressRequest); i {
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
		file_listener_listener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAddressResponse); i {
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
		file_listener_listener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTransactionRequest); i {
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
		file_listener_listener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTransactionResponse); i {
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
		file_listener_listener_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionsByAddressRequest); i {
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
		file_listener_listener_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionsByAddressResponse); i {
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
		file_listener_listener_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionsByAccountRequest); i {
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
		file_listener_listener_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionsByAccountResponse); i {
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
		file_listener_listener_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxByHashRequest); i {
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
		file_listener_listener_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxByHashResponse); i {
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
			RawDescriptor: file_listener_listener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_listener_listener_proto_goTypes,
		DependencyIndexes: file_listener_listener_proto_depIdxs,
		MessageInfos:      file_listener_listener_proto_msgTypes,
	}.Build()
	File_listener_listener_proto = out.File
	file_listener_listener_proto_rawDesc = nil
	file_listener_listener_proto_goTypes = nil
	file_listener_listener_proto_depIdxs = nil
}
