// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: limit.proto

package limit

import (
	meta "github.com/Rezky08/microservices-go/pb/meta"
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

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LimitId           uint64  `protobuf:"varint,12,opt,name=LimitId,proto3" json:"LimitId,omitempty"`
	KeyType           string  `protobuf:"bytes,1,opt,name=KeyType,proto3" json:"KeyType,omitempty"`
	KeyValue          string  `protobuf:"bytes,2,opt,name=KeyValue,proto3" json:"KeyValue,omitempty"`
	AvailableAmount   float64 `protobuf:"fixed64,3,opt,name=AvailableAmount,proto3" json:"AvailableAmount,omitempty"`
	UsableAmount      float64 `protobuf:"fixed64,4,opt,name=UsableAmount,proto3" json:"UsableAmount,omitempty"`
	OutstandingAmount float64 `protobuf:"fixed64,5,opt,name=OutstandingAmount,proto3" json:"OutstandingAmount,omitempty"`
	HoldAmount        float64 `protobuf:"fixed64,6,opt,name=HoldAmount,proto3" json:"HoldAmount,omitempty"`
	SharedAmount      float64 `protobuf:"fixed64,7,opt,name=SharedAmount,proto3" json:"SharedAmount,omitempty"`
	CcyId             float64 `protobuf:"fixed64,8,opt,name=CcyId,proto3" json:"CcyId,omitempty"`
	ExpiryDate        float64 `protobuf:"fixed64,9,opt,name=ExpiryDate,proto3" json:"ExpiryDate,omitempty"`
	ParentId          float64 `protobuf:"fixed64,10,opt,name=ParentId,proto3" json:"ParentId,omitempty"`
	IsInherit         float64 `protobuf:"fixed64,11,opt,name=IsInherit,proto3" json:"IsInherit,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_limit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_limit_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetLimitId() uint64 {
	if x != nil {
		return x.LimitId
	}
	return 0
}

func (x *Data) GetKeyType() string {
	if x != nil {
		return x.KeyType
	}
	return ""
}

func (x *Data) GetKeyValue() string {
	if x != nil {
		return x.KeyValue
	}
	return ""
}

func (x *Data) GetAvailableAmount() float64 {
	if x != nil {
		return x.AvailableAmount
	}
	return 0
}

func (x *Data) GetUsableAmount() float64 {
	if x != nil {
		return x.UsableAmount
	}
	return 0
}

func (x *Data) GetOutstandingAmount() float64 {
	if x != nil {
		return x.OutstandingAmount
	}
	return 0
}

func (x *Data) GetHoldAmount() float64 {
	if x != nil {
		return x.HoldAmount
	}
	return 0
}

func (x *Data) GetSharedAmount() float64 {
	if x != nil {
		return x.SharedAmount
	}
	return 0
}

func (x *Data) GetCcyId() float64 {
	if x != nil {
		return x.CcyId
	}
	return 0
}

func (x *Data) GetExpiryDate() float64 {
	if x != nil {
		return x.ExpiryDate
	}
	return 0
}

func (x *Data) GetParentId() float64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Data) GetIsInherit() float64 {
	if x != nil {
		return x.IsInherit
	}
	return 0
}

type LimitCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyType           string  `protobuf:"bytes,1,opt,name=KeyType,proto3" json:"KeyType,omitempty"`
	KeyValue          string  `protobuf:"bytes,2,opt,name=KeyValue,proto3" json:"KeyValue,omitempty"`
	AvailableAmount   float64 `protobuf:"fixed64,3,opt,name=AvailableAmount,proto3" json:"AvailableAmount,omitempty"`
	UsableAmount      float64 `protobuf:"fixed64,4,opt,name=UsableAmount,proto3" json:"UsableAmount,omitempty"`
	OutstandingAmount float64 `protobuf:"fixed64,5,opt,name=OutstandingAmount,proto3" json:"OutstandingAmount,omitempty"`
	HoldAmount        float64 `protobuf:"fixed64,6,opt,name=HoldAmount,proto3" json:"HoldAmount,omitempty"`
	SharedAmount      float64 `protobuf:"fixed64,7,opt,name=SharedAmount,proto3" json:"SharedAmount,omitempty"`
	CcyId             float64 `protobuf:"fixed64,8,opt,name=CcyId,proto3" json:"CcyId,omitempty"`
	ExpiryDate        float64 `protobuf:"fixed64,9,opt,name=ExpiryDate,proto3" json:"ExpiryDate,omitempty"`
	ParentId          float64 `protobuf:"fixed64,10,opt,name=ParentId,proto3" json:"ParentId,omitempty"`
	IsInherit         float64 `protobuf:"fixed64,11,opt,name=IsInherit,proto3" json:"IsInherit,omitempty"`
}

func (x *LimitCreate) Reset() {
	*x = LimitCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitCreate) ProtoMessage() {}

func (x *LimitCreate) ProtoReflect() protoreflect.Message {
	mi := &file_limit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitCreate.ProtoReflect.Descriptor instead.
func (*LimitCreate) Descriptor() ([]byte, []int) {
	return file_limit_proto_rawDescGZIP(), []int{1}
}

func (x *LimitCreate) GetKeyType() string {
	if x != nil {
		return x.KeyType
	}
	return ""
}

func (x *LimitCreate) GetKeyValue() string {
	if x != nil {
		return x.KeyValue
	}
	return ""
}

func (x *LimitCreate) GetAvailableAmount() float64 {
	if x != nil {
		return x.AvailableAmount
	}
	return 0
}

func (x *LimitCreate) GetUsableAmount() float64 {
	if x != nil {
		return x.UsableAmount
	}
	return 0
}

func (x *LimitCreate) GetOutstandingAmount() float64 {
	if x != nil {
		return x.OutstandingAmount
	}
	return 0
}

func (x *LimitCreate) GetHoldAmount() float64 {
	if x != nil {
		return x.HoldAmount
	}
	return 0
}

func (x *LimitCreate) GetSharedAmount() float64 {
	if x != nil {
		return x.SharedAmount
	}
	return 0
}

func (x *LimitCreate) GetCcyId() float64 {
	if x != nil {
		return x.CcyId
	}
	return 0
}

func (x *LimitCreate) GetExpiryDate() float64 {
	if x != nil {
		return x.ExpiryDate
	}
	return 0
}

func (x *LimitCreate) GetParentId() float64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *LimitCreate) GetIsInherit() float64 {
	if x != nil {
		return x.IsInherit
	}
	return 0
}

type LimitCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *meta.Meta `protobuf:"bytes,1,opt,name=Meta,proto3" json:"Meta,omitempty"`
}

func (x *LimitCreateResponse) Reset() {
	*x = LimitCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitCreateResponse) ProtoMessage() {}

func (x *LimitCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_limit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitCreateResponse.ProtoReflect.Descriptor instead.
func (*LimitCreateResponse) Descriptor() ([]byte, []int) {
	return file_limit_proto_rawDescGZIP(), []int{2}
}

func (x *LimitCreateResponse) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type LimitUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyType           string  `protobuf:"bytes,1,opt,name=KeyType,proto3" json:"KeyType,omitempty"`
	KeyValue          string  `protobuf:"bytes,2,opt,name=KeyValue,proto3" json:"KeyValue,omitempty"`
	AvailableAmount   float64 `protobuf:"fixed64,3,opt,name=AvailableAmount,proto3" json:"AvailableAmount,omitempty"`
	UsableAmount      float64 `protobuf:"fixed64,4,opt,name=UsableAmount,proto3" json:"UsableAmount,omitempty"`
	OutstandingAmount float64 `protobuf:"fixed64,5,opt,name=OutstandingAmount,proto3" json:"OutstandingAmount,omitempty"`
	HoldAmount        float64 `protobuf:"fixed64,6,opt,name=HoldAmount,proto3" json:"HoldAmount,omitempty"`
	SharedAmount      float64 `protobuf:"fixed64,7,opt,name=SharedAmount,proto3" json:"SharedAmount,omitempty"`
	CcyId             float64 `protobuf:"fixed64,8,opt,name=CcyId,proto3" json:"CcyId,omitempty"`
	ExpiryDate        float64 `protobuf:"fixed64,9,opt,name=ExpiryDate,proto3" json:"ExpiryDate,omitempty"`
	ParentId          float64 `protobuf:"fixed64,10,opt,name=ParentId,proto3" json:"ParentId,omitempty"`
	IsInherit         float64 `protobuf:"fixed64,11,opt,name=IsInherit,proto3" json:"IsInherit,omitempty"`
}

func (x *LimitUpdate) Reset() {
	*x = LimitUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitUpdate) ProtoMessage() {}

func (x *LimitUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_limit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitUpdate.ProtoReflect.Descriptor instead.
func (*LimitUpdate) Descriptor() ([]byte, []int) {
	return file_limit_proto_rawDescGZIP(), []int{3}
}

func (x *LimitUpdate) GetKeyType() string {
	if x != nil {
		return x.KeyType
	}
	return ""
}

func (x *LimitUpdate) GetKeyValue() string {
	if x != nil {
		return x.KeyValue
	}
	return ""
}

func (x *LimitUpdate) GetAvailableAmount() float64 {
	if x != nil {
		return x.AvailableAmount
	}
	return 0
}

func (x *LimitUpdate) GetUsableAmount() float64 {
	if x != nil {
		return x.UsableAmount
	}
	return 0
}

func (x *LimitUpdate) GetOutstandingAmount() float64 {
	if x != nil {
		return x.OutstandingAmount
	}
	return 0
}

func (x *LimitUpdate) GetHoldAmount() float64 {
	if x != nil {
		return x.HoldAmount
	}
	return 0
}

func (x *LimitUpdate) GetSharedAmount() float64 {
	if x != nil {
		return x.SharedAmount
	}
	return 0
}

func (x *LimitUpdate) GetCcyId() float64 {
	if x != nil {
		return x.CcyId
	}
	return 0
}

func (x *LimitUpdate) GetExpiryDate() float64 {
	if x != nil {
		return x.ExpiryDate
	}
	return 0
}

func (x *LimitUpdate) GetParentId() float64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *LimitUpdate) GetIsInherit() float64 {
	if x != nil {
		return x.IsInherit
	}
	return 0
}

type LimitUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *meta.Meta `protobuf:"bytes,1,opt,name=Meta,proto3" json:"Meta,omitempty"`
}

func (x *LimitUpdateResponse) Reset() {
	*x = LimitUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitUpdateResponse) ProtoMessage() {}

func (x *LimitUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_limit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitUpdateResponse.ProtoReflect.Descriptor instead.
func (*LimitUpdateResponse) Descriptor() ([]byte, []int) {
	return file_limit_proto_rawDescGZIP(), []int{4}
}

func (x *LimitUpdateResponse) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type LimitDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *LimitDetail) Reset() {
	*x = LimitDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitDetail) ProtoMessage() {}

func (x *LimitDetail) ProtoReflect() protoreflect.Message {
	mi := &file_limit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitDetail.ProtoReflect.Descriptor instead.
func (*LimitDetail) Descriptor() ([]byte, []int) {
	return file_limit_proto_rawDescGZIP(), []int{5}
}

func (x *LimitDetail) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type LimitDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *meta.Meta `protobuf:"bytes,1,opt,name=Meta,proto3" json:"Meta,omitempty"`
	Data *Data      `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *LimitDetailResponse) Reset() {
	*x = LimitDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limit_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitDetailResponse) ProtoMessage() {}

func (x *LimitDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_limit_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitDetailResponse.ProtoReflect.Descriptor instead.
func (*LimitDetailResponse) Descriptor() ([]byte, []int) {
	return file_limit_proto_rawDescGZIP(), []int{6}
}

func (x *LimitDetailResponse) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *LimitDetailResponse) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type LimitDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *LimitDelete) Reset() {
	*x = LimitDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limit_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitDelete) ProtoMessage() {}

func (x *LimitDelete) ProtoReflect() protoreflect.Message {
	mi := &file_limit_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitDelete.ProtoReflect.Descriptor instead.
func (*LimitDelete) Descriptor() ([]byte, []int) {
	return file_limit_proto_rawDescGZIP(), []int{7}
}

func (x *LimitDelete) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type LimitDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *meta.Meta `protobuf:"bytes,1,opt,name=Meta,proto3" json:"Meta,omitempty"`
	Data *Data      `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *LimitDeleteResponse) Reset() {
	*x = LimitDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limit_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitDeleteResponse) ProtoMessage() {}

func (x *LimitDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_limit_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitDeleteResponse.ProtoReflect.Descriptor instead.
func (*LimitDeleteResponse) Descriptor() ([]byte, []int) {
	return file_limit_proto_rawDescGZIP(), []int{8}
}

func (x *LimitDeleteResponse) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *LimitDeleteResponse) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type LimitCreateResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TempId    uint64 `protobuf:"varint,1,opt,name=TempId,proto3" json:"TempId,omitempty"`
	ChangesId uint64 `protobuf:"varint,2,opt,name=ChangesId,proto3" json:"ChangesId,omitempty"`
}

func (x *LimitCreateResponse_Data) Reset() {
	*x = LimitCreateResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limit_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitCreateResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitCreateResponse_Data) ProtoMessage() {}

func (x *LimitCreateResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_limit_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitCreateResponse_Data.ProtoReflect.Descriptor instead.
func (*LimitCreateResponse_Data) Descriptor() ([]byte, []int) {
	return file_limit_proto_rawDescGZIP(), []int{2, 0}
}

func (x *LimitCreateResponse_Data) GetTempId() uint64 {
	if x != nil {
		return x.TempId
	}
	return 0
}

func (x *LimitCreateResponse_Data) GetChangesId() uint64 {
	if x != nil {
		return x.ChangesId
	}
	return 0
}

type LimitUpdateResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TempId    uint64 `protobuf:"varint,1,opt,name=TempId,proto3" json:"TempId,omitempty"`
	ChangesId uint64 `protobuf:"varint,2,opt,name=ChangesId,proto3" json:"ChangesId,omitempty"`
}

func (x *LimitUpdateResponse_Data) Reset() {
	*x = LimitUpdateResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limit_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitUpdateResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitUpdateResponse_Data) ProtoMessage() {}

func (x *LimitUpdateResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_limit_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitUpdateResponse_Data.ProtoReflect.Descriptor instead.
func (*LimitUpdateResponse_Data) Descriptor() ([]byte, []int) {
	return file_limit_proto_rawDescGZIP(), []int{4, 0}
}

func (x *LimitUpdateResponse_Data) GetTempId() uint64 {
	if x != nil {
		return x.TempId
	}
	return 0
}

func (x *LimitUpdateResponse_Data) GetChangesId() uint64 {
	if x != nil {
		return x.ChangesId
	}
	return 0
}

var File_limit_proto protoreflect.FileDescriptor

var file_limit_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x1a, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x86, 0x03, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x49, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0f, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x55, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x55, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x4f, 0x75, 0x74, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x11, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x48, 0x6f, 0x6c, 0x64, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x48, 0x6f, 0x6c, 0x64, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x63, 0x79,
	0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x43, 0x63, 0x79, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x49,
	0x73, 0x49, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x49, 0x73, 0x49, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x22, 0xf3, 0x02, 0x0a, 0x0b, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x55, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0c, 0x55, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a,
	0x11, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x48,
	0x6f, 0x6c, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x48, 0x6f, 0x6c, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x43, 0x63, 0x79, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x43, 0x63, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x49, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x49, 0x73, 0x49, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x22,
	0x75, 0x0a, 0x13, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x3c, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x16, 0x0a, 0x06, 0x54, 0x65, 0x6d, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x54, 0x65, 0x6d, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x49, 0x64, 0x22, 0xf3, 0x02, 0x0a, 0x0b, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x55, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x55, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x4f, 0x75,
	0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x48, 0x6f, 0x6c, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x48, 0x6f,
	0x6c, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x43, 0x63, 0x79, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x43, 0x63, 0x79,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x49, 0x73, 0x49, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x49, 0x73, 0x49, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x22, 0x75, 0x0a, 0x13,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x04, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x3c, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x54, 0x65, 0x6d, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x54,
	0x65, 0x6d, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x49, 0x64, 0x22, 0x1d, 0x0a, 0x0b, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x49, 0x64, 0x22, 0x58, 0x0a, 0x13, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x4d, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x1d, 0x0a, 0x0b,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x13, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x32, 0x8a, 0x02, 0x0a, 0x0c, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x12, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x12, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x62, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_limit_proto_rawDescOnce sync.Once
	file_limit_proto_rawDescData = file_limit_proto_rawDesc
)

func file_limit_proto_rawDescGZIP() []byte {
	file_limit_proto_rawDescOnce.Do(func() {
		file_limit_proto_rawDescData = protoimpl.X.CompressGZIP(file_limit_proto_rawDescData)
	})
	return file_limit_proto_rawDescData
}

var file_limit_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_limit_proto_goTypes = []interface{}{
	(*Data)(nil),                     // 0: limit.Data
	(*LimitCreate)(nil),              // 1: limit.LimitCreate
	(*LimitCreateResponse)(nil),      // 2: limit.LimitCreateResponse
	(*LimitUpdate)(nil),              // 3: limit.LimitUpdate
	(*LimitUpdateResponse)(nil),      // 4: limit.LimitUpdateResponse
	(*LimitDetail)(nil),              // 5: limit.LimitDetail
	(*LimitDetailResponse)(nil),      // 6: limit.LimitDetailResponse
	(*LimitDelete)(nil),              // 7: limit.LimitDelete
	(*LimitDeleteResponse)(nil),      // 8: limit.LimitDeleteResponse
	(*LimitCreateResponse_Data)(nil), // 9: limit.LimitCreateResponse.Data
	(*LimitUpdateResponse_Data)(nil), // 10: limit.LimitUpdateResponse.Data
	(*meta.Meta)(nil),                // 11: format.Meta
}
var file_limit_proto_depIdxs = []int32{
	11, // 0: limit.LimitCreateResponse.Meta:type_name -> format.Meta
	11, // 1: limit.LimitUpdateResponse.Meta:type_name -> format.Meta
	11, // 2: limit.LimitDetailResponse.Meta:type_name -> format.Meta
	0,  // 3: limit.LimitDetailResponse.Data:type_name -> limit.Data
	11, // 4: limit.LimitDeleteResponse.Meta:type_name -> format.Meta
	0,  // 5: limit.LimitDeleteResponse.Data:type_name -> limit.Data
	1,  // 6: limit.LimitService.CreateLimit:input_type -> limit.LimitCreate
	3,  // 7: limit.LimitService.UpdateLimit:input_type -> limit.LimitUpdate
	5,  // 8: limit.LimitService.DetailLimit:input_type -> limit.LimitDetail
	7,  // 9: limit.LimitService.DeleteLimit:input_type -> limit.LimitDelete
	2,  // 10: limit.LimitService.CreateLimit:output_type -> limit.LimitCreateResponse
	4,  // 11: limit.LimitService.UpdateLimit:output_type -> limit.LimitUpdateResponse
	6,  // 12: limit.LimitService.DetailLimit:output_type -> limit.LimitDetailResponse
	4,  // 13: limit.LimitService.DeleteLimit:output_type -> limit.LimitUpdateResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_limit_proto_init() }
func file_limit_proto_init() {
	if File_limit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_limit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_limit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitCreate); i {
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
		file_limit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitCreateResponse); i {
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
		file_limit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitUpdate); i {
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
		file_limit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitUpdateResponse); i {
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
		file_limit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitDetail); i {
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
		file_limit_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitDetailResponse); i {
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
		file_limit_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitDelete); i {
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
		file_limit_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitDeleteResponse); i {
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
		file_limit_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitCreateResponse_Data); i {
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
		file_limit_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitUpdateResponse_Data); i {
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
			RawDescriptor: file_limit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_limit_proto_goTypes,
		DependencyIndexes: file_limit_proto_depIdxs,
		MessageInfos:      file_limit_proto_msgTypes,
	}.Build()
	File_limit_proto = out.File
	file_limit_proto_rawDesc = nil
	file_limit_proto_goTypes = nil
	file_limit_proto_depIdxs = nil
}
