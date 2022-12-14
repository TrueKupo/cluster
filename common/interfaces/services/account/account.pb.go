// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: account/account.proto

package account

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

type WordLang int32

const (
	WordLang_ENGLISH            WordLang = 0
	WordLang_FRENCH             WordLang = 1
	WordLang_ITALIAN            WordLang = 2
	WordLang_SPANISH            WordLang = 3
	WordLang_CZECH              WordLang = 4
	WordLang_JAPANESE           WordLang = 5
	WordLang_KOREAN             WordLang = 6
	WordLang_CHINESESIMPLIFIED  WordLang = 7
	WordLang_CHINESETRADITIONAL WordLang = 8
)

// Enum value maps for WordLang.
var (
	WordLang_name = map[int32]string{
		0: "ENGLISH",
		1: "FRENCH",
		2: "ITALIAN",
		3: "SPANISH",
		4: "CZECH",
		5: "JAPANESE",
		6: "KOREAN",
		7: "CHINESESIMPLIFIED",
		8: "CHINESETRADITIONAL",
	}
	WordLang_value = map[string]int32{
		"ENGLISH":            0,
		"FRENCH":             1,
		"ITALIAN":            2,
		"SPANISH":            3,
		"CZECH":              4,
		"JAPANESE":           5,
		"KOREAN":             6,
		"CHINESESIMPLIFIED":  7,
		"CHINESETRADITIONAL": 8,
	}
)

func (x WordLang) Enum() *WordLang {
	p := new(WordLang)
	*p = x
	return p
}

func (x WordLang) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WordLang) Descriptor() protoreflect.EnumDescriptor {
	return file_account_account_proto_enumTypes[0].Descriptor()
}

func (WordLang) Type() protoreflect.EnumType {
	return &file_account_account_proto_enumTypes[0]
}

func (x WordLang) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WordLang.Descriptor instead.
func (WordLang) EnumDescriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{0}
}

type NewMnemonicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entropy int32    `protobuf:"varint,1,opt,name=entropy,proto3" json:"entropy,omitempty"`
	Lang    WordLang `protobuf:"varint,2,opt,name=lang,proto3,enum=account.WordLang" json:"lang,omitempty"`
}

func (x *NewMnemonicRequest) Reset() {
	*x = NewMnemonicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMnemonicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMnemonicRequest) ProtoMessage() {}

func (x *NewMnemonicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMnemonicRequest.ProtoReflect.Descriptor instead.
func (*NewMnemonicRequest) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{0}
}

func (x *NewMnemonicRequest) GetEntropy() int32 {
	if x != nil {
		return x.Entropy
	}
	return 0
}

func (x *NewMnemonicRequest) GetLang() WordLang {
	if x != nil {
		return x.Lang
	}
	return WordLang_ENGLISH
}

type NewMnemonicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mnemonic  string           `protobuf:"bytes,1,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	RetStatus *response.Status `protobuf:"bytes,2,opt,name=ret_status,json=retStatus,proto3" json:"ret_status,omitempty"`
}

func (x *NewMnemonicResponse) Reset() {
	*x = NewMnemonicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMnemonicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMnemonicResponse) ProtoMessage() {}

func (x *NewMnemonicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMnemonicResponse.ProtoReflect.Descriptor instead.
func (*NewMnemonicResponse) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{1}
}

func (x *NewMnemonicResponse) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

func (x *NewMnemonicResponse) GetRetStatus() *response.Status {
	if x != nil {
		return x.RetStatus
	}
	return nil
}

type GetSeedFromMnemonicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mnemonic string `protobuf:"bytes,1,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *GetSeedFromMnemonicRequest) Reset() {
	*x = GetSeedFromMnemonicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSeedFromMnemonicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSeedFromMnemonicRequest) ProtoMessage() {}

func (x *GetSeedFromMnemonicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSeedFromMnemonicRequest.ProtoReflect.Descriptor instead.
func (*GetSeedFromMnemonicRequest) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{2}
}

func (x *GetSeedFromMnemonicRequest) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

func (x *GetSeedFromMnemonicRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetMasterKeyFromMnemonicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mnemonic string `protobuf:"bytes,1,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *GetMasterKeyFromMnemonicRequest) Reset() {
	*x = GetMasterKeyFromMnemonicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMasterKeyFromMnemonicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMasterKeyFromMnemonicRequest) ProtoMessage() {}

func (x *GetMasterKeyFromMnemonicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMasterKeyFromMnemonicRequest.ProtoReflect.Descriptor instead.
func (*GetMasterKeyFromMnemonicRequest) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{3}
}

func (x *GetMasterKeyFromMnemonicRequest) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

func (x *GetMasterKeyFromMnemonicRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SeedFromMnemonicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seed      []byte           `protobuf:"bytes,1,opt,name=seed,proto3" json:"seed,omitempty"`
	RetStatus *response.Status `protobuf:"bytes,2,opt,name=ret_status,json=retStatus,proto3" json:"ret_status,omitempty"`
}

func (x *SeedFromMnemonicResponse) Reset() {
	*x = SeedFromMnemonicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeedFromMnemonicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeedFromMnemonicResponse) ProtoMessage() {}

func (x *SeedFromMnemonicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeedFromMnemonicResponse.ProtoReflect.Descriptor instead.
func (*SeedFromMnemonicResponse) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{4}
}

func (x *SeedFromMnemonicResponse) GetSeed() []byte {
	if x != nil {
		return x.Seed
	}
	return nil
}

func (x *SeedFromMnemonicResponse) GetRetStatus() *response.Status {
	if x != nil {
		return x.RetStatus
	}
	return nil
}

type MasterKeyFromMnemonicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MasterKey []byte           `protobuf:"bytes,1,opt,name=master_key,json=masterKey,proto3" json:"master_key,omitempty"`
	RetStatus *response.Status `protobuf:"bytes,2,opt,name=ret_status,json=retStatus,proto3" json:"ret_status,omitempty"`
}

func (x *MasterKeyFromMnemonicResponse) Reset() {
	*x = MasterKeyFromMnemonicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterKeyFromMnemonicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterKeyFromMnemonicResponse) ProtoMessage() {}

func (x *MasterKeyFromMnemonicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterKeyFromMnemonicResponse.ProtoReflect.Descriptor instead.
func (*MasterKeyFromMnemonicResponse) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{5}
}

func (x *MasterKeyFromMnemonicResponse) GetMasterKey() []byte {
	if x != nil {
		return x.MasterKey
	}
	return nil
}

func (x *MasterKeyFromMnemonicResponse) GetRetStatus() *response.Status {
	if x != nil {
		return x.RetStatus
	}
	return nil
}

type SeedDeriveToAddressHexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol types.CoinCode `protobuf:"varint,1,opt,name=symbol,proto3,enum=types.CoinCode" json:"symbol,omitempty"`
	Seed   []byte         `protobuf:"bytes,2,opt,name=seed,proto3" json:"seed,omitempty"`
	Path   string         `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *SeedDeriveToAddressHexRequest) Reset() {
	*x = SeedDeriveToAddressHexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeedDeriveToAddressHexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeedDeriveToAddressHexRequest) ProtoMessage() {}

func (x *SeedDeriveToAddressHexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeedDeriveToAddressHexRequest.ProtoReflect.Descriptor instead.
func (*SeedDeriveToAddressHexRequest) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{6}
}

func (x *SeedDeriveToAddressHexRequest) GetSymbol() types.CoinCode {
	if x != nil {
		return x.Symbol
	}
	return types.CoinCode_BTC
}

func (x *SeedDeriveToAddressHexRequest) GetSeed() []byte {
	if x != nil {
		return x.Seed
	}
	return nil
}

func (x *SeedDeriveToAddressHexRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type SeedDeriveToAddressHexResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      string           `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	RetStatus *response.Status `protobuf:"bytes,2,opt,name=ret_status,json=retStatus,proto3" json:"ret_status,omitempty"`
}

func (x *SeedDeriveToAddressHexResponse) Reset() {
	*x = SeedDeriveToAddressHexResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeedDeriveToAddressHexResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeedDeriveToAddressHexResponse) ProtoMessage() {}

func (x *SeedDeriveToAddressHexResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeedDeriveToAddressHexResponse.ProtoReflect.Descriptor instead.
func (*SeedDeriveToAddressHexResponse) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{7}
}

func (x *SeedDeriveToAddressHexResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *SeedDeriveToAddressHexResponse) GetRetStatus() *response.Status {
	if x != nil {
		return x.RetStatus
	}
	return nil
}

type SeedDeriveToAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol types.CoinCode `protobuf:"varint,1,opt,name=symbol,proto3,enum=types.CoinCode" json:"symbol,omitempty"`
	Seed   []byte         `protobuf:"bytes,2,opt,name=seed,proto3" json:"seed,omitempty"`
	Path   string         `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *SeedDeriveToAccountRequest) Reset() {
	*x = SeedDeriveToAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeedDeriveToAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeedDeriveToAccountRequest) ProtoMessage() {}

func (x *SeedDeriveToAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeedDeriveToAccountRequest.ProtoReflect.Descriptor instead.
func (*SeedDeriveToAccountRequest) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{8}
}

func (x *SeedDeriveToAccountRequest) GetSymbol() types.CoinCode {
	if x != nil {
		return x.Symbol
	}
	return types.CoinCode_BTC
}

func (x *SeedDeriveToAccountRequest) GetSeed() []byte {
	if x != nil {
		return x.Seed
	}
	return nil
}

func (x *SeedDeriveToAccountRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type SeedDeriveToAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivateKey string           `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	PublicKey  string           `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PublicAddr string           `protobuf:"bytes,3,opt,name=public_addr,json=publicAddr,proto3" json:"public_addr,omitempty"`
	RetStatus  *response.Status `protobuf:"bytes,4,opt,name=ret_status,json=retStatus,proto3" json:"ret_status,omitempty"`
}

func (x *SeedDeriveToAccountResponse) Reset() {
	*x = SeedDeriveToAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeedDeriveToAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeedDeriveToAccountResponse) ProtoMessage() {}

func (x *SeedDeriveToAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeedDeriveToAccountResponse.ProtoReflect.Descriptor instead.
func (*SeedDeriveToAccountResponse) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{9}
}

func (x *SeedDeriveToAccountResponse) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *SeedDeriveToAccountResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *SeedDeriveToAccountResponse) GetPublicAddr() string {
	if x != nil {
		return x.PublicAddr
	}
	return ""
}

func (x *SeedDeriveToAccountResponse) GetRetStatus() *response.Status {
	if x != nil {
		return x.RetStatus
	}
	return nil
}

var File_account_account_proto protoreflect.FileDescriptor

var file_account_account_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x1a, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55,
	0x0a, 0x12, 0x4e, 0x65, 0x77, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x70, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x70, 0x79, 0x12, 0x25,
	0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x4c, 0x61, 0x6e, 0x67, 0x52,
	0x04, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x62, 0x0a, 0x13, 0x4e, 0x65, 0x77, 0x4d, 0x6e, 0x65, 0x6d,
	0x6f, 0x6e, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12, 0x2f, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09,
	0x72, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x54, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f,
	0x6e, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f,
	0x6e, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x59, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x46,
	0x72, 0x6f, 0x6d, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5f, 0x0a, 0x18, 0x53, 0x65,
	0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x0a, 0x72, 0x65,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x09, 0x72, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x6f, 0x0a, 0x1d, 0x4d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6e, 0x65, 0x6d,
	0x6f, 0x6e, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x0a, 0x72,
	0x65, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x09, 0x72, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x70, 0x0a, 0x1d,
	0x53, 0x65, 0x65, 0x64, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x48, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x65,
	0x0a, 0x1e, 0x53, 0x65, 0x65, 0x64, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x72, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x6d, 0x0a, 0x1a, 0x53, 0x65, 0x65, 0x64, 0x44, 0x65, 0x72,
	0x69, 0x76, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x65, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x22, 0xaf, 0x01, 0x0a, 0x1b, 0x53, 0x65, 0x65, 0x64, 0x44, 0x65, 0x72,
	0x69, 0x76, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x41, 0x64, 0x64, 0x72, 0x12, 0x2f, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x72, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x91, 0x01, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x64, 0x4c,
	0x61, 0x6e, 0x67, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x47, 0x4c, 0x49, 0x53, 0x48, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52, 0x45, 0x4e, 0x43, 0x48, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x49, 0x54, 0x41, 0x4c, 0x49, 0x41, 0x4e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x50, 0x41,
	0x4e, 0x49, 0x53, 0x48, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x5a, 0x45, 0x43, 0x48, 0x10,
	0x04, 0x12, 0x0c, 0x0a, 0x08, 0x4a, 0x41, 0x50, 0x41, 0x4e, 0x45, 0x53, 0x45, 0x10, 0x05, 0x12,
	0x0a, 0x0a, 0x06, 0x4b, 0x4f, 0x52, 0x45, 0x41, 0x4e, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x43,
	0x48, 0x49, 0x4e, 0x45, 0x53, 0x45, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x07, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x53, 0x45, 0x54, 0x52, 0x41,
	0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x08, 0x32, 0xf4, 0x04, 0x0a, 0x0e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a,
	0x0b, 0x4e, 0x65, 0x77, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12, 0x1b, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e,
	0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63,
	0x12, 0x23, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x53, 0x65, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6e,
	0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12, 0x28, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x46, 0x72, 0x6f,
	0x6d, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x4b, 0x65, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x18, 0x53, 0x65,
	0x65, 0x64, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x48, 0x65, 0x78, 0x12, 0x26, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x53, 0x65, 0x65, 0x64, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x48, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x65, 0x64, 0x44, 0x65, 0x72,
	0x69, 0x76, 0x65, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x65, 0x78, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x19, 0x53, 0x65, 0x65,
	0x64, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x48, 0x65, 0x78, 0x12, 0x26, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x53, 0x65, 0x65, 0x64, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x48, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x65, 0x64, 0x44, 0x65, 0x72,
	0x69, 0x76, 0x65, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x65, 0x78, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x17, 0x53, 0x65, 0x65,
	0x64, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x53,
	0x65, 0x65, 0x64, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x65, 0x64, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x54, 0x6f,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_account_proto_rawDescOnce sync.Once
	file_account_account_proto_rawDescData = file_account_account_proto_rawDesc
)

func file_account_account_proto_rawDescGZIP() []byte {
	file_account_account_proto_rawDescOnce.Do(func() {
		file_account_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_account_proto_rawDescData)
	})
	return file_account_account_proto_rawDescData
}

var file_account_account_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_account_account_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_account_account_proto_goTypes = []interface{}{
	(WordLang)(0),                           // 0: account.WordLang
	(*NewMnemonicRequest)(nil),              // 1: account.NewMnemonicRequest
	(*NewMnemonicResponse)(nil),             // 2: account.NewMnemonicResponse
	(*GetSeedFromMnemonicRequest)(nil),      // 3: account.GetSeedFromMnemonicRequest
	(*GetMasterKeyFromMnemonicRequest)(nil), // 4: account.GetMasterKeyFromMnemonicRequest
	(*SeedFromMnemonicResponse)(nil),        // 5: account.SeedFromMnemonicResponse
	(*MasterKeyFromMnemonicResponse)(nil),   // 6: account.MasterKeyFromMnemonicResponse
	(*SeedDeriveToAddressHexRequest)(nil),   // 7: account.SeedDeriveToAddressHexRequest
	(*SeedDeriveToAddressHexResponse)(nil),  // 8: account.SeedDeriveToAddressHexResponse
	(*SeedDeriveToAccountRequest)(nil),      // 9: account.SeedDeriveToAccountRequest
	(*SeedDeriveToAccountResponse)(nil),     // 10: account.SeedDeriveToAccountResponse
	(*response.Status)(nil),                 // 11: response.Status
	(types.CoinCode)(0),                     // 12: types.CoinCode
}
var file_account_account_proto_depIdxs = []int32{
	0,  // 0: account.NewMnemonicRequest.lang:type_name -> account.WordLang
	11, // 1: account.NewMnemonicResponse.ret_status:type_name -> response.Status
	11, // 2: account.SeedFromMnemonicResponse.ret_status:type_name -> response.Status
	11, // 3: account.MasterKeyFromMnemonicResponse.ret_status:type_name -> response.Status
	12, // 4: account.SeedDeriveToAddressHexRequest.symbol:type_name -> types.CoinCode
	11, // 5: account.SeedDeriveToAddressHexResponse.ret_status:type_name -> response.Status
	12, // 6: account.SeedDeriveToAccountRequest.symbol:type_name -> types.CoinCode
	11, // 7: account.SeedDeriveToAccountResponse.ret_status:type_name -> response.Status
	1,  // 8: account.AccountService.NewMnemonic:input_type -> account.NewMnemonicRequest
	3,  // 9: account.AccountService.GetSeedFromMnemonic:input_type -> account.GetSeedFromMnemonicRequest
	4,  // 10: account.AccountService.GetMasterKeyFromMnemonic:input_type -> account.GetMasterKeyFromMnemonicRequest
	7,  // 11: account.AccountService.SeedDeriveToPublicKeyHex:input_type -> account.SeedDeriveToAddressHexRequest
	7,  // 12: account.AccountService.SeedDeriveToPrivateKeyHex:input_type -> account.SeedDeriveToAddressHexRequest
	9,  // 13: account.AccountService.SeedDeriveToAccountData:input_type -> account.SeedDeriveToAccountRequest
	2,  // 14: account.AccountService.NewMnemonic:output_type -> account.NewMnemonicResponse
	5,  // 15: account.AccountService.GetSeedFromMnemonic:output_type -> account.SeedFromMnemonicResponse
	6,  // 16: account.AccountService.GetMasterKeyFromMnemonic:output_type -> account.MasterKeyFromMnemonicResponse
	8,  // 17: account.AccountService.SeedDeriveToPublicKeyHex:output_type -> account.SeedDeriveToAddressHexResponse
	8,  // 18: account.AccountService.SeedDeriveToPrivateKeyHex:output_type -> account.SeedDeriveToAddressHexResponse
	10, // 19: account.AccountService.SeedDeriveToAccountData:output_type -> account.SeedDeriveToAccountResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_account_account_proto_init() }
func file_account_account_proto_init() {
	if File_account_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMnemonicRequest); i {
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
		file_account_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMnemonicResponse); i {
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
		file_account_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSeedFromMnemonicRequest); i {
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
		file_account_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMasterKeyFromMnemonicRequest); i {
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
		file_account_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeedFromMnemonicResponse); i {
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
		file_account_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterKeyFromMnemonicResponse); i {
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
		file_account_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeedDeriveToAddressHexRequest); i {
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
		file_account_account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeedDeriveToAddressHexResponse); i {
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
		file_account_account_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeedDeriveToAccountRequest); i {
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
		file_account_account_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeedDeriveToAccountResponse); i {
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
			RawDescriptor: file_account_account_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_account_proto_goTypes,
		DependencyIndexes: file_account_account_proto_depIdxs,
		EnumInfos:         file_account_account_proto_enumTypes,
		MessageInfos:      file_account_account_proto_msgTypes,
	}.Build()
	File_account_account_proto = out.File
	file_account_account_proto_rawDesc = nil
	file_account_account_proto_goTypes = nil
	file_account_account_proto_depIdxs = nil
}
