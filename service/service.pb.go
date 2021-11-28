// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: service/service.proto

// Compile protofile with:
// protoc -I="." --go_out="." service/service.proto
// protoc --go-grpc_out="." service/service.proto

// go mod init packageName
// go mod tidy

package service

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

type Ack_Status int32

const (
	Ack_FAIL    Ack_Status = 0
	Ack_SUCCESS Ack_Status = 1
)

// Enum value maps for Ack_Status.
var (
	Ack_Status_name = map[int32]string{
		0: "FAIL",
		1: "SUCCESS",
	}
	Ack_Status_value = map[string]int32{
		"FAIL":    0,
		"SUCCESS": 1,
	}
)

func (x Ack_Status) Enum() *Ack_Status {
	p := new(Ack_Status)
	*p = x
	return p
}

func (x Ack_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ack_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_service_service_proto_enumTypes[0].Descriptor()
}

func (Ack_Status) Type() protoreflect.EnumType {
	return &file_service_service_proto_enumTypes[0]
}

func (x Ack_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ack_Status.Descriptor instead.
func (Ack_Status) EnumDescriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{9, 0}
}

type RequestPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Port string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *RequestPort) Reset() {
	*x = RequestPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPort) ProtoMessage() {}

func (x *RequestPort) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPort.ProtoReflect.Descriptor instead.
func (*RequestPort) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{0}
}

func (x *RequestPort) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RequestPort) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type ReturnPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ReturnPort) Reset() {
	*x = ReturnPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnPort) ProtoMessage() {}

func (x *ReturnPort) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnPort.ProtoReflect.Descriptor instead.
func (*ReturnPort) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReturnPort) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestName string `protobuf:"bytes,1,opt,name=requestName,proto3" json:"requestName,omitempty"`
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{2}
}

func (x *InfoRequest) GetRequestName() string {
	if x != nil {
		return x.RequestName
	}
	return ""
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseName string `protobuf:"bytes,1,opt,name=responseName,proto3" json:"responseName,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{3}
}

func (x *InfoResponse) GetResponseName() string {
	if x != nil {
		return x.ResponseName
	}
	return ""
}

type ReplicateClientAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addresses []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *ReplicateClientAddress) Reset() {
	*x = ReplicateClientAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicateClientAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicateClientAddress) ProtoMessage() {}

func (x *ReplicateClientAddress) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicateClientAddress.ProtoReflect.Descriptor instead.
func (*ReplicateClientAddress) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{4}
}

func (x *ReplicateClientAddress) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type RequestStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultComando string `protobuf:"bytes,1,opt,name=resultComando,proto3" json:"resultComando,omitempty"`
}

func (x *RequestStatus) Reset() {
	*x = RequestStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestStatus) ProtoMessage() {}

func (x *RequestStatus) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestStatus.ProtoReflect.Descriptor instead.
func (*RequestStatus) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{5}
}

func (x *RequestStatus) GetResultComando() string {
	if x != nil {
		return x.ResultComando
	}
	return ""
}

type AuctionResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ongoing bool   `protobuf:"varint,1,opt,name=ongoing,proto3" json:"ongoing,omitempty"`
	NodeId  string `protobuf:"bytes,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Price   int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *AuctionResults) Reset() {
	*x = AuctionResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuctionResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuctionResults) ProtoMessage() {}

func (x *AuctionResults) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuctionResults.ProtoReflect.Descriptor instead.
func (*AuctionResults) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{6}
}

func (x *AuctionResults) GetOngoing() bool {
	if x != nil {
		return x.Ongoing
	}
	return false
}

func (x *AuctionResults) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *AuctionResults) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type RequestBid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId  string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	MyBid   int32  `protobuf:"varint,2,opt,name=myBid,proto3" json:"myBid,omitempty"`
	Lamport int32  `protobuf:"varint,3,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *RequestBid) Reset() {
	*x = RequestBid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestBid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestBid) ProtoMessage() {}

func (x *RequestBid) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestBid.ProtoReflect.Descriptor instead.
func (*RequestBid) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{7}
}

func (x *RequestBid) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *RequestBid) GetMyBid() int32 {
	if x != nil {
		return x.MyBid
	}
	return 0
}

func (x *RequestBid) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

type HighestBidInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId     string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	HighestBid int32  `protobuf:"varint,2,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
	Lamport    int32  `protobuf:"varint,3,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *HighestBidInfo) Reset() {
	*x = HighestBidInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HighestBidInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HighestBidInfo) ProtoMessage() {}

func (x *HighestBidInfo) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HighestBidInfo.ProtoReflect.Descriptor instead.
func (*HighestBidInfo) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{8}
}

func (x *HighestBidInfo) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *HighestBidInfo) GetHighestBid() int32 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

func (x *HighestBidInfo) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Ack_Status `protobuf:"varint,1,opt,name=status,proto3,enum=service.Ack_Status" json:"status,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{9}
}

func (x *Ack) GetStatus() Ack_Status {
	if x != nil {
		return x.Status
	}
	return Ack_FAIL
}

type Replica struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionReplica *AuctionReplica `protobuf:"bytes,1,opt,name=auctionReplica,proto3" json:"auctionReplica,omitempty"`
	Addresses      []string        `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *Replica) Reset() {
	*x = Replica{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Replica) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Replica) ProtoMessage() {}

func (x *Replica) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Replica.ProtoReflect.Descriptor instead.
func (*Replica) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{10}
}

func (x *Replica) GetAuctionReplica() *AuctionReplica {
	if x != nil {
		return x.AuctionReplica
	}
	return nil
}

func (x *Replica) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type AuctionReplica struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HighestBidderId string `protobuf:"bytes,1,opt,name=highestBidderId,proto3" json:"highestBidderId,omitempty"`
	HighestBid      int32  `protobuf:"varint,2,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
	OnGoing         bool   `protobuf:"varint,3,opt,name=onGoing,proto3" json:"onGoing,omitempty"`
	TimeRemaining   int32  `protobuf:"varint,4,opt,name=timeRemaining,proto3" json:"timeRemaining,omitempty"`
}

func (x *AuctionReplica) Reset() {
	*x = AuctionReplica{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuctionReplica) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuctionReplica) ProtoMessage() {}

func (x *AuctionReplica) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuctionReplica.ProtoReflect.Descriptor instead.
func (*AuctionReplica) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{11}
}

func (x *AuctionReplica) GetHighestBidderId() string {
	if x != nil {
		return x.HighestBidderId
	}
	return ""
}

func (x *AuctionReplica) GetHighestBid() int32 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

func (x *AuctionReplica) GetOnGoing() bool {
	if x != nil {
		return x.OnGoing
	}
	return false
}

func (x *AuctionReplica) GetTimeRemaining() int32 {
	if x != nil {
		return x.TimeRemaining
	}
	return 0
}

var File_service_service_proto protoreflect.FileDescriptor

var file_service_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x35, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x26, 0x0a, 0x0a, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x2f, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x32, 0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x16, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x0d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6d, 0x61,
	0x6e, 0x64, 0x6f, 0x22, 0x58, 0x0a, 0x0e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x54, 0x0a,
	0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x79, 0x42, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6d, 0x79, 0x42, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x61, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x22, 0x62, 0x0a, 0x0e, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x53, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x2b,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1f, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x22, 0x68, 0x0a, 0x07,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x3f, 0x0a, 0x0e, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x0e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x0e, 0x41, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x28, 0x0a, 0x0f, 0x68, 0x69, 0x67,
	0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74,
	0x42, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x6e, 0x47, 0x6f, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6f, 0x6e, 0x47, 0x6f, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a,
	0x0d, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x32, 0x98, 0x03, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4e, 0x0a, 0x21, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x50, 0x6f, 0x72, 0x74, 0x4f, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x28, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x1a, 0x0c, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x39, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x1a, 0x0c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x3c, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x12,
	0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73,
	0x74, 0x42, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x2e, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x1a, 0x0c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x36, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_service_service_proto_rawDescOnce sync.Once
	file_service_service_proto_rawDescData = file_service_service_proto_rawDesc
)

func file_service_service_proto_rawDescGZIP() []byte {
	file_service_service_proto_rawDescOnce.Do(func() {
		file_service_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_service_proto_rawDescData)
	})
	return file_service_service_proto_rawDescData
}

var file_service_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_service_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_service_service_proto_goTypes = []interface{}{
	(Ack_Status)(0),                // 0: service.Ack.Status
	(*RequestPort)(nil),            // 1: service.RequestPort
	(*ReturnPort)(nil),             // 2: service.ReturnPort
	(*InfoRequest)(nil),            // 3: service.InfoRequest
	(*InfoResponse)(nil),           // 4: service.InfoResponse
	(*ReplicateClientAddress)(nil), // 5: service.ReplicateClientAddress
	(*RequestStatus)(nil),          // 6: service.RequestStatus
	(*AuctionResults)(nil),         // 7: service.AuctionResults
	(*RequestBid)(nil),             // 8: service.RequestBid
	(*HighestBidInfo)(nil),         // 9: service.HighestBidInfo
	(*Ack)(nil),                    // 10: service.Ack
	(*Replica)(nil),                // 11: service.Replica
	(*AuctionReplica)(nil),         // 12: service.AuctionReplica
}
var file_service_service_proto_depIdxs = []int32{
	0,  // 0: service.Ack.status:type_name -> service.Ack.Status
	12, // 1: service.Replica.auctionReplica:type_name -> service.AuctionReplica
	1,  // 2: service.Service.RegisterClientPortOnPrimaryServer:input_type -> service.RequestPort
	8,  // 3: service.Service.Bid:input_type -> service.RequestBid
	6,  // 4: service.Service.Result:input_type -> service.RequestStatus
	11, // 5: service.Service.PublishReplicate:input_type -> service.Replica
	9,  // 6: service.Service.ReplicateHighestBid:input_type -> service.HighestBidInfo
	8,  // 7: service.Service.Broadcast:input_type -> service.RequestBid
	3,  // 8: service.Service.GetName:input_type -> service.InfoRequest
	2,  // 9: service.Service.RegisterClientPortOnPrimaryServer:output_type -> service.ReturnPort
	10, // 10: service.Service.Bid:output_type -> service.Ack
	7,  // 11: service.Service.Result:output_type -> service.AuctionResults
	10, // 12: service.Service.PublishReplicate:output_type -> service.Ack
	10, // 13: service.Service.ReplicateHighestBid:output_type -> service.Ack
	10, // 14: service.Service.Broadcast:output_type -> service.Ack
	4,  // 15: service.Service.GetName:output_type -> service.InfoResponse
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_service_service_proto_init() }
func file_service_service_proto_init() {
	if File_service_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPort); i {
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
		file_service_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReturnPort); i {
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
		file_service_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
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
		file_service_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
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
		file_service_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicateClientAddress); i {
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
		file_service_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestStatus); i {
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
		file_service_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuctionResults); i {
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
		file_service_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestBid); i {
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
		file_service_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HighestBidInfo); i {
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
		file_service_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
		file_service_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Replica); i {
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
		file_service_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuctionReplica); i {
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
			RawDescriptor: file_service_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_service_proto_goTypes,
		DependencyIndexes: file_service_service_proto_depIdxs,
		EnumInfos:         file_service_service_proto_enumTypes,
		MessageInfos:      file_service_service_proto_msgTypes,
	}.Build()
	File_service_service_proto = out.File
	file_service_service_proto_rawDesc = nil
	file_service_service_proto_goTypes = nil
	file_service_service_proto_depIdxs = nil
}
