// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: chain.proto

package pro

import(
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

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version          uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	PreviousHash     string `protobuf:"bytes,2,opt,name=previous_hash,json=previousHash,proto3" json:"previous_hash,omitempty"`
	MerkleRoot       string `protobuf:"bytes,3,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	DifficultyTarget string `protobuf:"bytes,4,opt,name=difficulty_target,json=difficultyTarget,proto3" json:"difficulty_target,omitempty"`
	Nonce            uint32 `protobuf:"varint,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Timestamp        uint32 `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_chain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_chain_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Header) GetPreviousHash() string {
	if x != nil {
		return x.PreviousHash
	}
	return ""
}

func (x *Header) GetMerkleRoot() string {
	if x != nil {
		return x.MerkleRoot
	}
	return ""
}

func (x *Header) GetDifficultyTarget() string {
	if x != nil {
		return x.DifficultyTarget
	}
	return ""
}

func (x *Header) GetNonce() uint32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Header) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type TransactionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReferenceTransactionHash string `protobuf:"bytes,1,opt,name=reference_transaction_hash,json=referenceTransactionHash,proto3" json:"reference_transaction_hash,omitempty"`
	OutputIndex              uint32 `protobuf:"varint,2,opt,name=output_index,json=outputIndex,proto3" json:"output_index,omitempty"`
	UnlockingScript          string `protobuf:"bytes,3,opt,name=unlocking_script,json=unlockingScript,proto3" json:"unlocking_script,omitempty"`
}

func (x *TransactionInput) Reset() {
	*x = TransactionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionInput) ProtoMessage() {}

func (x *TransactionInput) ProtoReflect() protoreflect.Message {
	mi := &file_chain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionInput.ProtoReflect.Descriptor instead.
func (*TransactionInput) Descriptor() ([]byte, []int) {
	return file_chain_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionInput) GetReferenceTransactionHash() string {
	if x != nil {
		return x.ReferenceTransactionHash
	}
	return ""
}

func (x *TransactionInput) GetOutputIndex() uint32 {
	if x != nil {
		return x.OutputIndex
	}
	return 0
}

func (x *TransactionInput) GetUnlockingScript() string {
	if x != nil {
		return x.UnlockingScript
	}
	return ""
}

type TransactionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount        uint32 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	LockingScript string `protobuf:"bytes,2,opt,name=locking_script,json=lockingScript,proto3" json:"locking_script,omitempty"`
}

func (x *TransactionOutput) Reset() {
	*x = TransactionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionOutput) ProtoMessage() {}

func (x *TransactionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_chain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionOutput.ProtoReflect.Descriptor instead.
func (*TransactionOutput) Descriptor() ([]byte, []int) {
	return file_chain_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionOutput) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransactionOutput) GetLockingScript() string {
	if x != nil {
		return x.LockingScript
	}
	return ""
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  uint32               `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Inputs   []*TransactionInput  `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs  []*TransactionOutput `protobuf:"bytes,3,rep,name=outputs,proto3" json:"outputs,omitempty"`
	LockTime uint32               `protobuf:"varint,4,opt,name=lock_time,json=lockTime,proto3" json:"lock_time,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_chain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_chain_proto_rawDescGZIP(), []int{3}
}

func (x *Transaction) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Transaction) GetInputs() []*TransactionInput {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *Transaction) GetOutputs() []*TransactionOutput {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *Transaction) GetLockTime() uint32 {
	if x != nil {
		return x.LockTime
	}
	return 0
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *Header        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_chain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_chain_proto_rawDescGZIP(), []int{4}
}

func (x *Block) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type BlockRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header               *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Height               uint32  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	NumberOfTransactions uint32  `protobuf:"varint,3,opt,name=number_of_transactions,json=numberOfTransactions,proto3" json:"number_of_transactions,omitempty"`
	BlockFile            string  `protobuf:"bytes,4,opt,name=block_file,json=blockFile,proto3" json:"block_file,omitempty"`
	BlockStartOffset     uint32  `protobuf:"varint,5,opt,name=block_start_offset,json=blockStartOffset,proto3" json:"block_start_offset,omitempty"`
	BlockEndOffset       uint32  `protobuf:"varint,6,opt,name=block_end_offset,json=blockEndOffset,proto3" json:"block_end_offset,omitempty"`
	UndoFile             string  `protobuf:"bytes,7,opt,name=undo_file,json=undoFile,proto3" json:"undo_file,omitempty"`
	UndoStartOffset      uint32  `protobuf:"varint,8,opt,name=undo_start_offset,json=undoStartOffset,proto3" json:"undo_start_offset,omitempty"`
	UndoEndOffset        uint32  `protobuf:"varint,9,opt,name=undo_end_offset,json=undoEndOffset,proto3" json:"undo_end_offset,omitempty"`
}

func (x *BlockRecord) Reset() {
	*x = BlockRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRecord) ProtoMessage() {}

func (x *BlockRecord) ProtoReflect() protoreflect.Message {
	mi := &file_chain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRecord.ProtoReflect.Descriptor instead.
func (*BlockRecord) Descriptor() ([]byte, []int) {
	return file_chain_proto_rawDescGZIP(), []int{5}
}

func (x *BlockRecord) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BlockRecord) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BlockRecord) GetNumberOfTransactions() uint32 {
	if x != nil {
		return x.NumberOfTransactions
	}
	return 0
}

func (x *BlockRecord) GetBlockFile() string {
	if x != nil {
		return x.BlockFile
	}
	return ""
}

func (x *BlockRecord) GetBlockStartOffset() uint32 {
	if x != nil {
		return x.BlockStartOffset
	}
	return 0
}

func (x *BlockRecord) GetBlockEndOffset() uint32 {
	if x != nil {
		return x.BlockEndOffset
	}
	return 0
}

func (x *BlockRecord) GetUndoFile() string {
	if x != nil {
		return x.UndoFile
	}
	return ""
}

func (x *BlockRecord) GetUndoStartOffset() uint32 {
	if x != nil {
		return x.UndoStartOffset
	}
	return 0
}

func (x *BlockRecord) GetUndoEndOffset() uint32 {
	if x != nil {
		return x.UndoEndOffset
	}
	return 0
}

type CoinRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version        uint32   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Active         bool     `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	OutputIndexes  []uint32 `protobuf:"varint,3,rep,packed,name=output_indexes,json=outputIndexes,proto3" json:"output_indexes,omitempty"`
	Amounts        []uint32 `protobuf:"varint,4,rep,packed,name=amounts,proto3" json:"amounts,omitempty"`
	LockingScripts []string `protobuf:"bytes,5,rep,name=locking_scripts,json=lockingScripts,proto3" json:"locking_scripts,omitempty"`
}

func (x *CoinRecord) Reset() {
	*x = CoinRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinRecord) ProtoMessage() {}

func (x *CoinRecord) ProtoReflect() protoreflect.Message {
	mi := &file_chain_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinRecord.ProtoReflect.Descriptor instead.
func (*CoinRecord) Descriptor() ([]byte, []int) {
	return file_chain_proto_rawDescGZIP(), []int{6}
}

func (x *CoinRecord) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *CoinRecord) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *CoinRecord) GetOutputIndexes() []uint32 {
	if x != nil {
		return x.OutputIndexes
	}
	return nil
}

func (x *CoinRecord) GetAmounts() []uint32 {
	if x != nil {
		return x.Amounts
	}
	return nil
}

func (x *CoinRecord) GetLockingScripts() []string {
	if x != nil {
		return x.LockingScripts
	}
	return nil
}

type UndoBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionInputHashes []string `protobuf:"bytes,1,rep,name=transaction_input_hashes,json=transactionInputHashes,proto3" json:"transaction_input_hashes,omitempty"`
	OutputIndexes          []uint32 `protobuf:"varint,2,rep,packed,name=output_indexes,json=outputIndexes,proto3" json:"output_indexes,omitempty"`
	Amounts                []uint32 `protobuf:"varint,3,rep,packed,name=amounts,proto3" json:"amounts,omitempty"`
	LockingScripts         []string `protobuf:"bytes,4,rep,name=locking_scripts,json=lockingScripts,proto3" json:"locking_scripts,omitempty"`
}

func (x *UndoBlock) Reset() {
	*x = UndoBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UndoBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UndoBlock) ProtoMessage() {}

func (x *UndoBlock) ProtoReflect() protoreflect.Message {
	mi := &file_chain_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UndoBlock.ProtoReflect.Descriptor instead.
func (*UndoBlock) Descriptor() ([]byte, []int) {
	return file_chain_proto_rawDescGZIP(), []int{7}
}

func (x *UndoBlock) GetTransactionInputHashes() []string {
	if x != nil {
		return x.TransactionInputHashes
	}
	return nil
}

func (x *UndoBlock) GetOutputIndexes() []uint32 {
	if x != nil {
		return x.OutputIndexes
	}
	return nil
}

func (x *UndoBlock) GetAmounts() []uint32 {
	if x != nil {
		return x.Amounts
	}
	return nil
}

func (x *UndoBlock) GetLockingScripts() []string {
	if x != nil {
		return x.LockingScripts
	}
	return nil
}

var File_chain_proto protoreflect.FileDescriptor

var file_chain_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01,
	0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x6b, 0x6c,
	0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65,
	0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x9e, 0x01, 0x0a, 0x10, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x3c,
	0x0a, 0x1a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x18, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x29, 0x0a, 0x10, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x52, 0x0a, 0x11, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x9d,
	0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x06, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x5a,
	0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xe4, 0x02, 0x0a, 0x0b, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x14, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x64, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x64, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x64, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2a, 0x0a,
	0x11, 0x75, 0x6e, 0x64, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x75, 0x6e, 0x64, 0x6f, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x6e, 0x64,
	0x6f, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x75, 0x6e, 0x64, 0x6f, 0x45, 0x6e, 0x64, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x22, 0xa8, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x22, 0xaf, 0x01, 0x0a,
	0x09, 0x55, 0x6e, 0x64, 0x6f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x16, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x48, 0x61,
	0x73, 0x68, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e,
	0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chain_proto_rawDescOnce sync.Once
	file_chain_proto_rawDescData = file_chain_proto_rawDesc
)

func file_chain_proto_rawDescGZIP() []byte {
	file_chain_proto_rawDescOnce.Do(func() {
		file_chain_proto_rawDescData = protoimpl.X.CompressGZIP(file_chain_proto_rawDescData)
	})
	return file_chain_proto_rawDescData
}

var file_chain_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_chain_proto_goTypes = []interface{}{
	(*Header)(nil),            // 0: Header
	(*TransactionInput)(nil),  // 1: TransactionInput
	(*TransactionOutput)(nil), // 2: TransactionOutput
	(*Transaction)(nil),       // 3: Transaction
	(*Block)(nil),             // 4: Block
	(*BlockRecord)(nil),       // 5: BlockRecord
	(*CoinRecord)(nil),        // 6: CoinRecord
	(*UndoBlock)(nil),         // 7: UndoBlock
}
var file_chain_proto_depIdxs = []int32{
	1, // 0: Transaction.inputs:type_name -> TransactionInput
	2, // 1: Transaction.outputs:type_name -> TransactionOutput
	0, // 2: Block.header:type_name -> Header
	3, // 3: Block.transactions:type_name -> Transaction
	0, // 4: BlockRecord.header:type_name -> Header
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_chain_proto_init() }
func file_chain_proto_init() {
	if File_chain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_chain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionInput); i {
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
		file_chain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionOutput); i {
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
		file_chain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_chain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_chain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRecord); i {
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
		file_chain_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinRecord); i {
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
		file_chain_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UndoBlock); i {
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
			RawDescriptor: file_chain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chain_proto_goTypes,
		DependencyIndexes: file_chain_proto_depIdxs,
		MessageInfos:      file_chain_proto_msgTypes,
	}.Build()
	File_chain_proto = out.File
	file_chain_proto_rawDesc = nil
	file_chain_proto_goTypes = nil
	file_chain_proto_depIdxs = nil
}
