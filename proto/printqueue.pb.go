// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: printqueue.proto

package proto

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

type Destination int32

const (
	Destination_DESTINATION_UNKNOWN Destination = 0
	Destination_DESTINATION_RECEIPT Destination = 1
)

// Enum value maps for Destination.
var (
	Destination_name = map[int32]string{
		0: "DESTINATION_UNKNOWN",
		1: "DESTINATION_RECEIPT",
	}
	Destination_value = map[string]int32{
		"DESTINATION_UNKNOWN": 0,
		"DESTINATION_RECEIPT": 1,
	}
)

func (x Destination) Enum() *Destination {
	p := new(Destination)
	*p = x
	return p
}

func (x Destination) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Destination) Descriptor() protoreflect.EnumDescriptor {
	return file_printqueue_proto_enumTypes[0].Descriptor()
}

func (Destination) Type() protoreflect.EnumType {
	return &file_printqueue_proto_enumTypes[0]
}

func (x Destination) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Destination.Descriptor instead.
func (Destination) EnumDescriptor() ([]byte, []int) {
	return file_printqueue_proto_rawDescGZIP(), []int{0}
}

type Urgency int32

const (
	Urgency_URGENCY_UNKNOWN   Urgency = 0
	Urgency_URGENCY_REGULAR   Urgency = 1
	Urgency_URGENCY_IMMEDIATE Urgency = 2
)

// Enum value maps for Urgency.
var (
	Urgency_name = map[int32]string{
		0: "URGENCY_UNKNOWN",
		1: "URGENCY_REGULAR",
		2: "URGENCY_IMMEDIATE",
	}
	Urgency_value = map[string]int32{
		"URGENCY_UNKNOWN":   0,
		"URGENCY_REGULAR":   1,
		"URGENCY_IMMEDIATE": 2,
	}
)

func (x Urgency) Enum() *Urgency {
	p := new(Urgency)
	*p = x
	return p
}

func (x Urgency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Urgency) Descriptor() protoreflect.EnumDescriptor {
	return file_printqueue_proto_enumTypes[1].Descriptor()
}

func (Urgency) Type() protoreflect.EnumType {
	return &file_printqueue_proto_enumTypes[1]
}

func (x Urgency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Urgency.Descriptor instead.
func (Urgency) EnumDescriptor() ([]byte, []int) {
	return file_printqueue_proto_rawDescGZIP(), []int{1}
}

type Fanout int32

const (
	Fanout_FANOUT_UNKNOWN Fanout = 0
	Fanout_FANOUT_ONE     Fanout = 1
	Fanout_FANOUT_ALL     Fanout = 2
)

// Enum value maps for Fanout.
var (
	Fanout_name = map[int32]string{
		0: "FANOUT_UNKNOWN",
		1: "FANOUT_ONE",
		2: "FANOUT_ALL",
	}
	Fanout_value = map[string]int32{
		"FANOUT_UNKNOWN": 0,
		"FANOUT_ONE":     1,
		"FANOUT_ALL":     2,
	}
)

func (x Fanout) Enum() *Fanout {
	p := new(Fanout)
	*p = x
	return p
}

func (x Fanout) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Fanout) Descriptor() protoreflect.EnumDescriptor {
	return file_printqueue_proto_enumTypes[2].Descriptor()
}

func (Fanout) Type() protoreflect.EnumType {
	return &file_printqueue_proto_enumTypes[2]
}

func (x Fanout) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Fanout.Descriptor instead.
func (Fanout) EnumDescriptor() ([]byte, []int) {
	return file_printqueue_proto_rawDescGZIP(), []int{2}
}

type PrintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lines       []string    `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines,omitempty"`
	Origin      string      `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Urgency     Urgency     `protobuf:"varint,3,opt,name=urgency,proto3,enum=printqueue.Urgency" json:"urgency,omitempty"`
	Destination Destination `protobuf:"varint,4,opt,name=destination,proto3,enum=printqueue.Destination" json:"destination,omitempty"`
	Fanout      Fanout      `protobuf:"varint,5,opt,name=fanout,proto3,enum=printqueue.Fanout" json:"fanout,omitempty"`
}

func (x *PrintRequest) Reset() {
	*x = PrintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_printqueue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintRequest) ProtoMessage() {}

func (x *PrintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_printqueue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintRequest.ProtoReflect.Descriptor instead.
func (*PrintRequest) Descriptor() ([]byte, []int) {
	return file_printqueue_proto_rawDescGZIP(), []int{0}
}

func (x *PrintRequest) GetLines() []string {
	if x != nil {
		return x.Lines
	}
	return nil
}

func (x *PrintRequest) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *PrintRequest) GetUrgency() Urgency {
	if x != nil {
		return x.Urgency
	}
	return Urgency_URGENCY_UNKNOWN
}

func (x *PrintRequest) GetDestination() Destination {
	if x != nil {
		return x.Destination
	}
	return Destination_DESTINATION_UNKNOWN
}

func (x *PrintRequest) GetFanout() Fanout {
	if x != nil {
		return x.Fanout
	}
	return Fanout_FANOUT_UNKNOWN
}

type StoredPrintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lines       []string    `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines,omitempty"`
	Origin      string      `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Urgency     Urgency     `protobuf:"varint,3,opt,name=urgency,proto3,enum=printqueue.Urgency" json:"urgency,omitempty"`
	Destination Destination `protobuf:"varint,4,opt,name=destination,proto3,enum=printqueue.Destination" json:"destination,omitempty"`
	Fanout      Fanout      `protobuf:"varint,5,opt,name=fanout,proto3,enum=printqueue.Fanout" json:"fanout,omitempty"`
}

func (x *StoredPrintRequest) Reset() {
	*x = StoredPrintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_printqueue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoredPrintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredPrintRequest) ProtoMessage() {}

func (x *StoredPrintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_printqueue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredPrintRequest.ProtoReflect.Descriptor instead.
func (*StoredPrintRequest) Descriptor() ([]byte, []int) {
	return file_printqueue_proto_rawDescGZIP(), []int{1}
}

func (x *StoredPrintRequest) GetLines() []string {
	if x != nil {
		return x.Lines
	}
	return nil
}

func (x *StoredPrintRequest) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *StoredPrintRequest) GetUrgency() Urgency {
	if x != nil {
		return x.Urgency
	}
	return Urgency_URGENCY_UNKNOWN
}

func (x *StoredPrintRequest) GetDestination() Destination {
	if x != nil {
		return x.Destination
	}
	return Destination_DESTINATION_UNKNOWN
}

func (x *StoredPrintRequest) GetFanout() Fanout {
	if x != nil {
		return x.Fanout
	}
	return Fanout_FANOUT_UNKNOWN
}

type PrintResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PrintResponse) Reset() {
	*x = PrintResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_printqueue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrintResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintResponse) ProtoMessage() {}

func (x *PrintResponse) ProtoReflect() protoreflect.Message {
	mi := &file_printqueue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintResponse.ProtoReflect.Descriptor instead.
func (*PrintResponse) Descriptor() ([]byte, []int) {
	return file_printqueue_proto_rawDescGZIP(), []int{2}
}

func (x *PrintResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ClientPrintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lines   []string `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines,omitempty"`
	Origin  string   `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Urgency Urgency  `protobuf:"varint,3,opt,name=urgency,proto3,enum=printqueue.Urgency" json:"urgency,omitempty"`
}

func (x *ClientPrintRequest) Reset() {
	*x = ClientPrintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_printqueue_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientPrintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientPrintRequest) ProtoMessage() {}

func (x *ClientPrintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_printqueue_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientPrintRequest.ProtoReflect.Descriptor instead.
func (*ClientPrintRequest) Descriptor() ([]byte, []int) {
	return file_printqueue_proto_rawDescGZIP(), []int{3}
}

func (x *ClientPrintRequest) GetLines() []string {
	if x != nil {
		return x.Lines
	}
	return nil
}

func (x *ClientPrintRequest) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *ClientPrintRequest) GetUrgency() Urgency {
	if x != nil {
		return x.Urgency
	}
	return Urgency_URGENCY_UNKNOWN
}

type ClientPrintResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientPrintResponse) Reset() {
	*x = ClientPrintResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_printqueue_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientPrintResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientPrintResponse) ProtoMessage() {}

func (x *ClientPrintResponse) ProtoReflect() protoreflect.Message {
	mi := &file_printqueue_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientPrintResponse.ProtoReflect.Descriptor instead.
func (*ClientPrintResponse) Descriptor() ([]byte, []int) {
	return file_printqueue_proto_rawDescGZIP(), []int{4}
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_printqueue_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_printqueue_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_printqueue_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_printqueue_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_printqueue_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_printqueue_proto_rawDescGZIP(), []int{6}
}

type RegisterPrinterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CallbackAddress string      `protobuf:"bytes,2,opt,name=callback_address,json=callbackAddress,proto3" json:"callback_address,omitempty"`
	ReceiverType    Destination `protobuf:"varint,3,opt,name=receiver_type,json=receiverType,proto3,enum=printqueue.Destination" json:"receiver_type,omitempty"`
}

func (x *RegisterPrinterRequest) Reset() {
	*x = RegisterPrinterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_printqueue_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPrinterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPrinterRequest) ProtoMessage() {}

func (x *RegisterPrinterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_printqueue_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPrinterRequest.ProtoReflect.Descriptor instead.
func (*RegisterPrinterRequest) Descriptor() ([]byte, []int) {
	return file_printqueue_proto_rawDescGZIP(), []int{7}
}

func (x *RegisterPrinterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegisterPrinterRequest) GetCallbackAddress() string {
	if x != nil {
		return x.CallbackAddress
	}
	return ""
}

func (x *RegisterPrinterRequest) GetReceiverType() Destination {
	if x != nil {
		return x.ReceiverType
	}
	return Destination_DESTINATION_UNKNOWN
}

type RegisterPrinterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterPrinterResponse) Reset() {
	*x = RegisterPrinterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_printqueue_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPrinterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPrinterResponse) ProtoMessage() {}

func (x *RegisterPrinterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_printqueue_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPrinterResponse.ProtoReflect.Descriptor instead.
func (*RegisterPrinterResponse) Descriptor() ([]byte, []int) {
	return file_printqueue_proto_rawDescGZIP(), []int{8}
}

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_printqueue_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_printqueue_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_printqueue_proto_rawDescGZIP(), []int{9}
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_printqueue_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_printqueue_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_printqueue_proto_rawDescGZIP(), []int{10}
}

var File_printqueue_proto protoreflect.FileDescriptor

var file_printqueue_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x71, 0x75, 0x65, 0x75, 0x65, 0x22, 0xd2,
	0x01, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x2d, 0x0a,
	0x07, 0x75, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x55, 0x72, 0x67, 0x65,
	0x6e, 0x63, 0x79, 0x52, 0x07, 0x75, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x39, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x06, 0x66, 0x61, 0x6e, 0x6f, 0x75,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x2e, 0x46, 0x61, 0x6e, 0x6f, 0x75, 0x74, 0x52, 0x06, 0x66, 0x61, 0x6e,
	0x6f, 0x75, 0x74, 0x22, 0xd8, 0x01, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x72,
	0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x2d, 0x0a, 0x07, 0x75, 0x72, 0x67, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x55, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x07,
	0x75, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x39, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x06, 0x66, 0x61, 0x6e, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e,
	0x46, 0x61, 0x6e, 0x6f, 0x75, 0x74, 0x52, 0x06, 0x66, 0x61, 0x6e, 0x6f, 0x75, 0x74, 0x22, 0x1f,
	0x0a, 0x0d, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x71, 0x0a, 0x12, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x12, 0x2d, 0x0a, 0x07, 0x75, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x2e, 0x55, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x07, 0x75, 0x72, 0x67, 0x65, 0x6e,
	0x63, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x91, 0x01, 0x0a,
	0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x3c, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x19, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x13, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x3f, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13,
	0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x45,
	0x49, 0x50, 0x54, 0x10, 0x01, 0x2a, 0x4a, 0x0a, 0x07, 0x55, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x13, 0x0a, 0x0f, 0x55, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59,
	0x5f, 0x52, 0x45, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x52,
	0x47, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x49, 0x4d, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x45, 0x10,
	0x02, 0x2a, 0x3c, 0x0a, 0x06, 0x46, 0x61, 0x6e, 0x6f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x46,
	0x41, 0x4e, 0x4f, 0x55, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x46, 0x41, 0x4e, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12,
	0x0e, 0x0a, 0x0a, 0x46, 0x41, 0x4e, 0x4f, 0x55, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x32,
	0xef, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3e, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x32, 0xb2, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x09, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x69,
	0x63, 0x2f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_printqueue_proto_rawDescOnce sync.Once
	file_printqueue_proto_rawDescData = file_printqueue_proto_rawDesc
)

func file_printqueue_proto_rawDescGZIP() []byte {
	file_printqueue_proto_rawDescOnce.Do(func() {
		file_printqueue_proto_rawDescData = protoimpl.X.CompressGZIP(file_printqueue_proto_rawDescData)
	})
	return file_printqueue_proto_rawDescData
}

var file_printqueue_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_printqueue_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_printqueue_proto_goTypes = []interface{}{
	(Destination)(0),                // 0: printqueue.Destination
	(Urgency)(0),                    // 1: printqueue.Urgency
	(Fanout)(0),                     // 2: printqueue.Fanout
	(*PrintRequest)(nil),            // 3: printqueue.PrintRequest
	(*StoredPrintRequest)(nil),      // 4: printqueue.StoredPrintRequest
	(*PrintResponse)(nil),           // 5: printqueue.PrintResponse
	(*ClientPrintRequest)(nil),      // 6: printqueue.ClientPrintRequest
	(*ClientPrintResponse)(nil),     // 7: printqueue.ClientPrintResponse
	(*DeleteRequest)(nil),           // 8: printqueue.DeleteRequest
	(*DeleteResponse)(nil),          // 9: printqueue.DeleteResponse
	(*RegisterPrinterRequest)(nil),  // 10: printqueue.RegisterPrinterRequest
	(*RegisterPrinterResponse)(nil), // 11: printqueue.RegisterPrinterResponse
	(*HeartbeatRequest)(nil),        // 12: printqueue.HeartbeatRequest
	(*HeartbeatResponse)(nil),       // 13: printqueue.HeartbeatResponse
}
var file_printqueue_proto_depIdxs = []int32{
	1,  // 0: printqueue.PrintRequest.urgency:type_name -> printqueue.Urgency
	0,  // 1: printqueue.PrintRequest.destination:type_name -> printqueue.Destination
	2,  // 2: printqueue.PrintRequest.fanout:type_name -> printqueue.Fanout
	1,  // 3: printqueue.StoredPrintRequest.urgency:type_name -> printqueue.Urgency
	0,  // 4: printqueue.StoredPrintRequest.destination:type_name -> printqueue.Destination
	2,  // 5: printqueue.StoredPrintRequest.fanout:type_name -> printqueue.Fanout
	1,  // 6: printqueue.ClientPrintRequest.urgency:type_name -> printqueue.Urgency
	0,  // 7: printqueue.RegisterPrinterRequest.receiver_type:type_name -> printqueue.Destination
	3,  // 8: printqueue.PrintService.Print:input_type -> printqueue.PrintRequest
	8,  // 9: printqueue.PrintService.Delete:input_type -> printqueue.DeleteRequest
	10, // 10: printqueue.PrintService.RegisterPrinter:input_type -> printqueue.RegisterPrinterRequest
	6,  // 11: printqueue.PrintClientService.ClientPrint:input_type -> printqueue.ClientPrintRequest
	12, // 12: printqueue.PrintClientService.Heartbeat:input_type -> printqueue.HeartbeatRequest
	5,  // 13: printqueue.PrintService.Print:output_type -> printqueue.PrintResponse
	9,  // 14: printqueue.PrintService.Delete:output_type -> printqueue.DeleteResponse
	11, // 15: printqueue.PrintService.RegisterPrinter:output_type -> printqueue.RegisterPrinterResponse
	7,  // 16: printqueue.PrintClientService.ClientPrint:output_type -> printqueue.ClientPrintResponse
	13, // 17: printqueue.PrintClientService.Heartbeat:output_type -> printqueue.HeartbeatResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_printqueue_proto_init() }
func file_printqueue_proto_init() {
	if File_printqueue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_printqueue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrintRequest); i {
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
		file_printqueue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoredPrintRequest); i {
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
		file_printqueue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrintResponse); i {
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
		file_printqueue_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientPrintRequest); i {
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
		file_printqueue_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientPrintResponse); i {
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
		file_printqueue_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_printqueue_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_printqueue_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPrinterRequest); i {
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
		file_printqueue_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPrinterResponse); i {
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
		file_printqueue_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
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
		file_printqueue_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
			RawDescriptor: file_printqueue_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_printqueue_proto_goTypes,
		DependencyIndexes: file_printqueue_proto_depIdxs,
		EnumInfos:         file_printqueue_proto_enumTypes,
		MessageInfos:      file_printqueue_proto_msgTypes,
	}.Build()
	File_printqueue_proto = out.File
	file_printqueue_proto_rawDesc = nil
	file_printqueue_proto_goTypes = nil
	file_printqueue_proto_depIdxs = nil
}
