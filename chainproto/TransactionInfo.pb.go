// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: transaction.proto

package chainproto

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

type TransactionType int32

const (
	TransactionType_NORMAL         TransactionType = 0
	TransactionType_MULI_SIGN      TransactionType = 1
	TransactionType_RC20_CONTRACT  TransactionType = 2
	TransactionType_RC721_CONTRACT TransactionType = 3
	TransactionType_XVM_CONTRACT   TransactionType = 4
	TransactionType_JSVM_CONTRACT  TransactionType = 5
	TransactionType_EVFS_CONTRACT  TransactionType = 6 //evfs
	TransactionType_CHAIN_CONFIG   TransactionType = 255
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0:   "NORMAL",
		1:   "MULI_SIGN",
		2:   "RC20_CONTRACT",
		3:   "RC721_CONTRACT",
		4:   "XVM_CONTRACT",
		5:   "JSVM_CONTRACT",
		6:   "EVFS_CONTRACT",
		255: "CHAIN_CONFIG",
	}
	TransactionType_value = map[string]int32{
		"NORMAL":         0,
		"MULI_SIGN":      1,
		"RC20_CONTRACT":  2,
		"RC721_CONTRACT": 3,
		"XVM_CONTRACT":   4,
		"JSVM_CONTRACT":  5,
		"EVFS_CONTRACT":  6,
		"CHAIN_CONFIG":   255,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_transaction_proto_enumTypes[0].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_transaction_proto_enumTypes[0]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{0}
}

//函数操作类型：{
//1.construct(address managers[],int voteCount) public returns (bool);创建合约
//2.addManagers(address to[],int voteCount);
//3.rmManagers(address to[]),int voteCount);
//}
type ContractMultiSig_FunctionMultiSig int32

const (
	ContractMultiSig_UNKNOW      ContractMultiSig_FunctionMultiSig = 0
	ContractMultiSig_CONSTRUCT   ContractMultiSig_FunctionMultiSig = 1
	ContractMultiSig_ADDMANAGERS ContractMultiSig_FunctionMultiSig = 2
	ContractMultiSig_RMMANAGERS  ContractMultiSig_FunctionMultiSig = 3
)

// Enum value maps for ContractMultiSig_FunctionMultiSig.
var (
	ContractMultiSig_FunctionMultiSig_name = map[int32]string{
		0: "UNKNOW",
		1: "CONSTRUCT",
		2: "ADDMANAGERS",
		3: "RMMANAGERS",
	}
	ContractMultiSig_FunctionMultiSig_value = map[string]int32{
		"UNKNOW":      0,
		"CONSTRUCT":   1,
		"ADDMANAGERS": 2,
		"RMMANAGERS":  3,
	}
)

func (x ContractMultiSig_FunctionMultiSig) Enum() *ContractMultiSig_FunctionMultiSig {
	p := new(ContractMultiSig_FunctionMultiSig)
	*p = x
	return p
}

func (x ContractMultiSig_FunctionMultiSig) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContractMultiSig_FunctionMultiSig) Descriptor() protoreflect.EnumDescriptor {
	return file_transaction_proto_enumTypes[1].Descriptor()
}

func (ContractMultiSig_FunctionMultiSig) Type() protoreflect.EnumType {
	return &file_transaction_proto_enumTypes[1]
}

func (x ContractMultiSig_FunctionMultiSig) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContractMultiSig_FunctionMultiSig.Descriptor instead.
func (ContractMultiSig_FunctionMultiSig) EnumDescriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{2, 0}
}

type TransactionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash      []byte           `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Body      *TransactionBody `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Signature []byte           `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	//	TransactionStatus status = 4; @deprecated,status apply in result trie
	Node            *TransactionNode `protobuf:"bytes,5,opt,name=node,proto3" json:"node,omitempty"`
	Accepttimestamp int64            `protobuf:"varint,6,opt,name=accepttimestamp,proto3" json:"accepttimestamp,omitempty"`
	Droptimestamp   int64            `protobuf:"varint,7,opt,name=droptimestamp,proto3" json:"droptimestamp,omitempty"`
	Dropreason      string           `protobuf:"bytes,8,opt,name=dropreason,proto3" json:"dropreason,omitempty"`
}

func (x *TransactionInfo) Reset() {
	*x = TransactionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionInfo) ProtoMessage() {}

func (x *TransactionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionInfo.ProtoReflect.Descriptor instead.
func (*TransactionInfo) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionInfo) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *TransactionInfo) GetBody() *TransactionBody {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *TransactionInfo) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *TransactionInfo) GetNode() *TransactionNode {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *TransactionInfo) GetAccepttimestamp() int64 {
	if x != nil {
		return x.Accepttimestamp
	}
	return 0
}

func (x *TransactionInfo) GetDroptimestamp() int64 {
	if x != nil {
		return x.Droptimestamp
	}
	return 0
}

func (x *TransactionInfo) GetDropreason() string {
	if x != nil {
		return x.Dropreason
	}
	return ""
}

type TransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce   int64  `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Address []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	//	repeated TransactionOutput outputs = 3;
	Recipient     []byte `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount        []byte `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Fees          []byte `protobuf:"bytes,5,opt,name=fees,proto3" json:"fees,omitempty"`                                         //手续费高位
	InnerCodetype int32  `protobuf:"varint,6,opt,name=inner_codetype,json=innerCodetype,proto3" json:"inner_codetype,omitempty"` //内置指令交易[0=普通交易,1=多重签名交易(取消),2=类ETH交易，，3=RC721交易(取消)，4=XVM合约调用,5=JSVM合约调用(暂无),6=evfs交易,7=链委员会(暂无),8=链管理员组(暂无)
	CodeData      []byte `protobuf:"bytes,7,opt,name=code_data,json=codeData,proto3" json:"code_data,omitempty"`                 //指令数据
	ExtData       []byte `protobuf:"bytes,8,opt,name=ext_data,json=extData,proto3" json:"ext_data,omitempty"`
	Timestamp     int64  `protobuf:"varint,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	BizCode       []byte `protobuf:"bytes,10,opt,name=biz_code,json=bizCode,proto3" json:"biz_code,omitempty"` //业务代码
	Version       []byte `protobuf:"bytes,11,opt,name=version,proto3" json:"version,omitempty"`                //版本
	ChainId       int32  `protobuf:"varint,12,opt,name=chainId,proto3" json:"chainId,omitempty"`
	Deadline      int64  `protobuf:"varint,13,opt,name=deadline,proto3" json:"deadline,omitempty"` // 什么时候丢弃该交易，0表示永远不丢弃，除非nonce超过很多次
	// for web3 compatible
	To    []byte `protobuf:"bytes,20,opt,name=to,proto3" json:"to,omitempty"`
	Value []byte `protobuf:"bytes,21,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TransactionBody) Reset() {
	*x = TransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionBody) ProtoMessage() {}

func (x *TransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionBody.ProtoReflect.Descriptor instead.
func (*TransactionBody) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionBody) GetNonce() int64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *TransactionBody) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *TransactionBody) GetRecipient() []byte {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *TransactionBody) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *TransactionBody) GetFees() []byte {
	if x != nil {
		return x.Fees
	}
	return nil
}

func (x *TransactionBody) GetInnerCodetype() int32 {
	if x != nil {
		return x.InnerCodetype
	}
	return 0
}

func (x *TransactionBody) GetCodeData() []byte {
	if x != nil {
		return x.CodeData
	}
	return nil
}

func (x *TransactionBody) GetExtData() []byte {
	if x != nil {
		return x.ExtData
	}
	return nil
}

func (x *TransactionBody) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *TransactionBody) GetBizCode() []byte {
	if x != nil {
		return x.BizCode
	}
	return nil
}

func (x *TransactionBody) GetVersion() []byte {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *TransactionBody) GetChainId() int32 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *TransactionBody) GetDeadline() int64 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

func (x *TransactionBody) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *TransactionBody) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type ContractMultiSig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Function ContractMultiSig_FunctionMultiSig `protobuf:"varint,1,opt,name=function,proto3,enum=org.brewchain.mcore.model.ContractMultiSig_FunctionMultiSig" json:"function,omitempty"` //
	Managers [][]byte                          `protobuf:"bytes,2,rep,name=managers,proto3" json:"managers,omitempty"`
	ExtDatas []byte                            `protobuf:"bytes,3,opt,name=ext_datas,json=extDatas,proto3" json:"ext_datas,omitempty"` //扩展信息
	Name     string                            `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`                         //构建函数时候使用
	Symbol   string                            `protobuf:"bytes,5,opt,name=symbol,proto3" json:"symbol,omitempty"`                     //构建函数时候使用
	MinVote  int32                             `protobuf:"varint,6,opt,name=min_vote,json=minVote,proto3" json:"min_vote,omitempty"`   //最少几个人投票
}

func (x *ContractMultiSig) Reset() {
	*x = ContractMultiSig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractMultiSig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractMultiSig) ProtoMessage() {}

func (x *ContractMultiSig) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractMultiSig.ProtoReflect.Descriptor instead.
func (*ContractMultiSig) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *ContractMultiSig) GetFunction() ContractMultiSig_FunctionMultiSig {
	if x != nil {
		return x.Function
	}
	return ContractMultiSig_UNKNOW
}

func (x *ContractMultiSig) GetManagers() [][]byte {
	if x != nil {
		return x.Managers
	}
	return nil
}

func (x *ContractMultiSig) GetExtDatas() []byte {
	if x != nil {
		return x.ExtDatas
	}
	return nil
}

func (x *ContractMultiSig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContractMultiSig) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *ContractMultiSig) GetMinVote() int32 {
	if x != nil {
		return x.MinVote
	}
	return 0
}

type PBMultsigAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creator  []byte   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Managers [][]byte `protobuf:"bytes,2,rep,name=managers,proto3" json:"managers,omitempty"`
	MinVote  int32    `protobuf:"varint,3,opt,name=min_vote,json=minVote,proto3" json:"min_vote,omitempty"` //至少多少个人签名
}

func (x *PBMultsigAccount) Reset() {
	*x = PBMultsigAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBMultsigAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBMultsigAccount) ProtoMessage() {}

func (x *PBMultsigAccount) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBMultsigAccount.ProtoReflect.Descriptor instead.
func (*PBMultsigAccount) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *PBMultsigAccount) GetCreator() []byte {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *PBMultsigAccount) GetManagers() [][]byte {
	if x != nil {
		return x.Managers
	}
	return nil
}

func (x *PBMultsigAccount) GetMinVote() int32 {
	if x != nil {
		return x.MinVote
	}
	return 0
}

type TransactionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address     []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Amount      []byte   `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Token       []byte   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	TokenAmount []byte   `protobuf:"bytes,4,opt,name=tokenAmount,proto3" json:"tokenAmount,omitempty"`
	Symbol      []byte   `protobuf:"bytes,5,opt,name=symbol,proto3" json:"symbol,omitempty"`
	CryptoToken [][]byte `protobuf:"bytes,6,rep,name=cryptoToken,proto3" json:"cryptoToken,omitempty"`
}

func (x *TransactionOutput) Reset() {
	*x = TransactionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionOutput) ProtoMessage() {}

func (x *TransactionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[4]
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
	return file_transaction_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionOutput) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *TransactionOutput) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *TransactionOutput) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *TransactionOutput) GetTokenAmount() []byte {
	if x != nil {
		return x.TokenAmount
	}
	return nil
}

func (x *TransactionOutput) GetSymbol() []byte {
	if x != nil {
		return x.Symbol
	}
	return nil
}

func (x *TransactionOutput) GetCryptoToken() [][]byte {
	if x != nil {
		return x.CryptoToken
	}
	return nil
}

type TransactionNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nid     string `protobuf:"bytes,1,opt,name=nid,proto3" json:"nid,omitempty"`
	Address []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *TransactionNode) Reset() {
	*x = TransactionNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionNode) ProtoMessage() {}

func (x *TransactionNode) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionNode.ProtoReflect.Descriptor instead.
func (*TransactionNode) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{5}
}

func (x *TransactionNode) GetNid() string {
	if x != nil {
		return x.Nid
	}
	return ""
}

func (x *TransactionNode) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

type TransactionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Result []byte `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	//	bytes hash = 3;
	Height int64 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	//	int64 timestamp = 5; //for saving storage
	//support for eth like chain client
	Logs      [][]byte `protobuf:"bytes,6,rep,name=logs,proto3" json:"logs,omitempty"`
	LogsBloom []byte   `protobuf:"bytes,7,opt,name=logsBloom,proto3" json:"logsBloom,omitempty"`
	//	int32 blockIndex = 8; //for saving storage
	GasUsed           []byte `protobuf:"bytes,9,opt,name=gasUsed,proto3" json:"gasUsed,omitempty"`
	CumulativeGasUsed []byte `protobuf:"bytes,10,opt,name=cumulativeGasUsed,proto3" json:"cumulativeGasUsed,omitempty"`
	ContractAddress   []byte `protobuf:"bytes,11,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
}

func (x *TransactionStatus) Reset() {
	*x = TransactionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionStatus) ProtoMessage() {}

func (x *TransactionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionStatus.ProtoReflect.Descriptor instead.
func (*TransactionStatus) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{6}
}

func (x *TransactionStatus) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TransactionStatus) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *TransactionStatus) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *TransactionStatus) GetLogs() [][]byte {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *TransactionStatus) GetLogsBloom() []byte {
	if x != nil {
		return x.LogsBloom
	}
	return nil
}

func (x *TransactionStatus) GetGasUsed() []byte {
	if x != nil {
		return x.GasUsed
	}
	return nil
}

func (x *TransactionStatus) GetCumulativeGasUsed() []byte {
	if x != nil {
		return x.CumulativeGasUsed
	}
	return nil
}

func (x *TransactionStatus) GetContractAddress() []byte {
	if x != nil {
		return x.ContractAddress
	}
	return nil
}

type BroadcastTransactionMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash  [][]byte `protobuf:"bytes,1,rep,name=txHash,proto3" json:"txHash,omitempty"`
	TxDatas [][]byte `protobuf:"bytes,2,rep,name=txDatas,proto3" json:"txDatas,omitempty"`
}

func (x *BroadcastTransactionMsg) Reset() {
	*x = BroadcastTransactionMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastTransactionMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastTransactionMsg) ProtoMessage() {}

func (x *BroadcastTransactionMsg) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastTransactionMsg.ProtoReflect.Descriptor instead.
func (*BroadcastTransactionMsg) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{7}
}

func (x *BroadcastTransactionMsg) GetTxHash() [][]byte {
	if x != nil {
		return x.TxHash
	}
	return nil
}

func (x *BroadcastTransactionMsg) GetTxDatas() [][]byte {
	if x != nil {
		return x.TxDatas
	}
	return nil
}

var File_transaction_proto protoreflect.FileDescriptor

var file_transaction_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x72, 0x65, 0x77, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x6d, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xb3,
	0x02, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x3e, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x72, 0x65, 0x77, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x72, 0x65, 0x77, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x6d, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x6e, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x61,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x24,
	0x0a, 0x0d, 0x64, 0x72, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x72, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x72, 0x6f, 0x70, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x72, 0x6f, 0x70, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x22, 0x99, 0x03, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x65, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x65,
	0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x64,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6f,
	0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x19, 0x0a, 0x08, 0x62, 0x69, 0x7a, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x62, 0x69, 0x7a, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0xbc, 0x02, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x53, 0x69, 0x67, 0x12, 0x58, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x72,
	0x65, 0x77, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x53, 0x69, 0x67, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x53, 0x69, 0x67, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x08, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x65,
	0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x6f, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x56, 0x6f, 0x74, 0x65, 0x22,
	0x4e, 0x0a, 0x10, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x53, 0x69, 0x67, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x43, 0x54, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x41, 0x44, 0x44, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x53, 0x10, 0x02, 0x12,
	0x0e, 0x0a, 0x0a, 0x52, 0x4d, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x53, 0x10, 0x03, 0x22,
	0x63, 0x0a, 0x10, 0x50, 0x42, 0x4d, 0x75, 0x6c, 0x74, 0x73, 0x69, 0x67, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x08, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e,
	0x5f, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x69, 0x6e,
	0x56, 0x6f, 0x74, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x0b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3d,
	0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6e, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xff, 0x01,
	0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x6f, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x73, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x73, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x12, 0x18, 0x0a,
	0x07, 0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x75, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x47, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x11, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x47, 0x61,
	0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x4b, 0x0a, 0x17, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x78, 0x44, 0x61, 0x74, 0x61, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x78, 0x44, 0x61, 0x74, 0x61, 0x73, 0x2a, 0x9e, 0x01, 0x0a,
	0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x4d, 0x55, 0x4c, 0x49, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x52,
	0x43, 0x32, 0x30, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x41, 0x43, 0x54, 0x10, 0x02, 0x12, 0x12,
	0x0a, 0x0e, 0x52, 0x43, 0x37, 0x32, 0x31, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x41, 0x43, 0x54,
	0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x58, 0x56, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x41,
	0x43, 0x54, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x4a, 0x53, 0x56, 0x4d, 0x5f, 0x43, 0x4f, 0x4e,
	0x54, 0x52, 0x41, 0x43, 0x54, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x56, 0x46, 0x53, 0x5f,
	0x43, 0x4f, 0x4e, 0x54, 0x52, 0x41, 0x43, 0x54, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0c, 0x43, 0x48,
	0x41, 0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0xff, 0x01, 0x42, 0x14, 0x5a,
	0x12, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_proto_rawDescOnce sync.Once
	file_transaction_proto_rawDescData = file_transaction_proto_rawDesc
)

func file_transaction_proto_rawDescGZIP() []byte {
	file_transaction_proto_rawDescOnce.Do(func() {
		file_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_proto_rawDescData)
	})
	return file_transaction_proto_rawDescData
}

var file_transaction_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_transaction_proto_goTypes = []interface{}{
	(TransactionType)(0),                   // 0: org.brewchain.mcore.model.TransactionType
	(ContractMultiSig_FunctionMultiSig)(0), // 1: org.brewchain.mcore.model.ContractMultiSig.FunctionMultiSig
	(*TransactionInfo)(nil),                // 2: org.brewchain.mcore.model.TransactionInfo
	(*TransactionBody)(nil),                // 3: org.brewchain.mcore.model.TransactionBody
	(*ContractMultiSig)(nil),               // 4: org.brewchain.mcore.model.ContractMultiSig
	(*PBMultsigAccount)(nil),               // 5: org.brewchain.mcore.model.PBMultsigAccount
	(*TransactionOutput)(nil),              // 6: org.brewchain.mcore.model.TransactionOutput
	(*TransactionNode)(nil),                // 7: org.brewchain.mcore.model.TransactionNode
	(*TransactionStatus)(nil),              // 8: org.brewchain.mcore.model.TransactionStatus
	(*BroadcastTransactionMsg)(nil),        // 9: org.brewchain.mcore.model.BroadcastTransactionMsg
}
var file_transaction_proto_depIdxs = []int32{
	3, // 0: org.brewchain.mcore.model.TransactionInfo.body:type_name -> org.brewchain.mcore.model.TransactionBody
	7, // 1: org.brewchain.mcore.model.TransactionInfo.node:type_name -> org.brewchain.mcore.model.TransactionNode
	1, // 2: org.brewchain.mcore.model.ContractMultiSig.function:type_name -> org.brewchain.mcore.model.ContractMultiSig.FunctionMultiSig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_transaction_proto_init() }
func file_transaction_proto_init() {
	if File_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionInfo); i {
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
		file_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionBody); i {
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
		file_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractMultiSig); i {
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
		file_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBMultsigAccount); i {
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
		file_transaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_transaction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionNode); i {
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
		file_transaction_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionStatus); i {
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
		file_transaction_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastTransactionMsg); i {
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
			RawDescriptor: file_transaction_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transaction_proto_goTypes,
		DependencyIndexes: file_transaction_proto_depIdxs,
		EnumInfos:         file_transaction_proto_enumTypes,
		MessageInfos:      file_transaction_proto_msgTypes,
	}.Build()
	File_transaction_proto = out.File
	file_transaction_proto_rawDesc = nil
	file_transaction_proto_goTypes = nil
	file_transaction_proto_depIdxs = nil
}