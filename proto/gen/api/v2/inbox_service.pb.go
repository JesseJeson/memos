// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v2/inbox_service.proto

package apiv2

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Inbox_Status int32

const (
	Inbox_STATUS_UNSPECIFIED Inbox_Status = 0
	Inbox_UNREAD             Inbox_Status = 1
	Inbox_ARCHIVED           Inbox_Status = 2
)

// Enum value maps for Inbox_Status.
var (
	Inbox_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "UNREAD",
		2: "ARCHIVED",
	}
	Inbox_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"UNREAD":             1,
		"ARCHIVED":           2,
	}
)

func (x Inbox_Status) Enum() *Inbox_Status {
	p := new(Inbox_Status)
	*p = x
	return p
}

func (x Inbox_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Inbox_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v2_inbox_service_proto_enumTypes[0].Descriptor()
}

func (Inbox_Status) Type() protoreflect.EnumType {
	return &file_api_v2_inbox_service_proto_enumTypes[0]
}

func (x Inbox_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Inbox_Status.Descriptor instead.
func (Inbox_Status) EnumDescriptor() ([]byte, []int) {
	return file_api_v2_inbox_service_proto_rawDescGZIP(), []int{0, 0}
}

type Inbox_Type int32

const (
	Inbox_TYPE_UNSPECIFIED    Inbox_Type = 0
	Inbox_TYPE_MEMO_COMMENT   Inbox_Type = 1
	Inbox_TYPE_VERSION_UPDATE Inbox_Type = 2
)

// Enum value maps for Inbox_Type.
var (
	Inbox_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_MEMO_COMMENT",
		2: "TYPE_VERSION_UPDATE",
	}
	Inbox_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":    0,
		"TYPE_MEMO_COMMENT":   1,
		"TYPE_VERSION_UPDATE": 2,
	}
)

func (x Inbox_Type) Enum() *Inbox_Type {
	p := new(Inbox_Type)
	*p = x
	return p
}

func (x Inbox_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Inbox_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v2_inbox_service_proto_enumTypes[1].Descriptor()
}

func (Inbox_Type) Type() protoreflect.EnumType {
	return &file_api_v2_inbox_service_proto_enumTypes[1]
}

func (x Inbox_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Inbox_Type.Descriptor instead.
func (Inbox_Type) EnumDescriptor() ([]byte, []int) {
	return file_api_v2_inbox_service_proto_rawDescGZIP(), []int{0, 1}
}

type Inbox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the inbox.
	// Format: inboxes/{id}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Format: users/{username}
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	// Format: users/{username}
	Receiver   string                 `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Status     Inbox_Status           `protobuf:"varint,4,opt,name=status,proto3,enum=memos.api.v2.Inbox_Status" json:"status,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Type       Inbox_Type             `protobuf:"varint,6,opt,name=type,proto3,enum=memos.api.v2.Inbox_Type" json:"type,omitempty"`
	ActivityId *int32                 `protobuf:"varint,7,opt,name=activity_id,json=activityId,proto3,oneof" json:"activity_id,omitempty"`
}

func (x *Inbox) Reset() {
	*x = Inbox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_inbox_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inbox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inbox) ProtoMessage() {}

func (x *Inbox) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_inbox_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inbox.ProtoReflect.Descriptor instead.
func (*Inbox) Descriptor() ([]byte, []int) {
	return file_api_v2_inbox_service_proto_rawDescGZIP(), []int{0}
}

func (x *Inbox) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Inbox) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Inbox) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *Inbox) GetStatus() Inbox_Status {
	if x != nil {
		return x.Status
	}
	return Inbox_STATUS_UNSPECIFIED
}

func (x *Inbox) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Inbox) GetType() Inbox_Type {
	if x != nil {
		return x.Type
	}
	return Inbox_TYPE_UNSPECIFIED
}

func (x *Inbox) GetActivityId() int32 {
	if x != nil && x.ActivityId != nil {
		return *x.ActivityId
	}
	return 0
}

type ListInboxesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Format: users/{username}
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ListInboxesRequest) Reset() {
	*x = ListInboxesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_inbox_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInboxesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInboxesRequest) ProtoMessage() {}

func (x *ListInboxesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_inbox_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInboxesRequest.ProtoReflect.Descriptor instead.
func (*ListInboxesRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_inbox_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListInboxesRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type ListInboxesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inboxes []*Inbox `protobuf:"bytes,1,rep,name=inboxes,proto3" json:"inboxes,omitempty"`
}

func (x *ListInboxesResponse) Reset() {
	*x = ListInboxesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_inbox_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInboxesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInboxesResponse) ProtoMessage() {}

func (x *ListInboxesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_inbox_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInboxesResponse.ProtoReflect.Descriptor instead.
func (*ListInboxesResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_inbox_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListInboxesResponse) GetInboxes() []*Inbox {
	if x != nil {
		return x.Inboxes
	}
	return nil
}

type UpdateInboxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the inbox to delete.
	// Format: inboxes/{inbox}
	Name       string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Inbox      *Inbox                 `protobuf:"bytes,2,opt,name=inbox,proto3" json:"inbox,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateInboxRequest) Reset() {
	*x = UpdateInboxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_inbox_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInboxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInboxRequest) ProtoMessage() {}

func (x *UpdateInboxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_inbox_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInboxRequest.ProtoReflect.Descriptor instead.
func (*UpdateInboxRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_inbox_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateInboxRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateInboxRequest) GetInbox() *Inbox {
	if x != nil {
		return x.Inbox
	}
	return nil
}

func (x *UpdateInboxRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type UpdateInboxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inbox *Inbox `protobuf:"bytes,1,opt,name=inbox,proto3" json:"inbox,omitempty"`
}

func (x *UpdateInboxResponse) Reset() {
	*x = UpdateInboxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_inbox_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInboxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInboxResponse) ProtoMessage() {}

func (x *UpdateInboxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_inbox_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInboxResponse.ProtoReflect.Descriptor instead.
func (*UpdateInboxResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_inbox_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateInboxResponse) GetInbox() *Inbox {
	if x != nil {
		return x.Inbox
	}
	return nil
}

type DeleteInboxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the inbox to delete.
	// Format: inboxes/{inbox}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteInboxRequest) Reset() {
	*x = DeleteInboxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_inbox_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInboxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInboxRequest) ProtoMessage() {}

func (x *DeleteInboxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_inbox_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInboxRequest.ProtoReflect.Descriptor instead.
func (*DeleteInboxRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_inbox_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteInboxRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteInboxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteInboxResponse) Reset() {
	*x = DeleteInboxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_inbox_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInboxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInboxResponse) ProtoMessage() {}

func (x *DeleteInboxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_inbox_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInboxResponse.ProtoReflect.Descriptor instead.
func (*DeleteInboxResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_inbox_service_proto_rawDescGZIP(), []int{6}
}

var File_api_v2_inbox_service_proto protoreflect.FileDescriptor

var file_api_v2_inbox_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65,
	0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x03, 0x0a, 0x05, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01, 0x22, 0x3a, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x55, 0x4e, 0x52, 0x45, 0x41, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x52, 0x43,
	0x48, 0x49, 0x56, 0x45, 0x44, 0x10, 0x02, 0x22, 0x4c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45,
	0x4d, 0x4f, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x10, 0x02, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x62,
	0x6f, 0x78, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x44, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x07, 0x69, 0x6e,
	0x62, 0x6f, 0x78, 0x65, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x29, 0x0a, 0x05, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49,
	0x6e, 0x62, 0x6f, 0x78, 0x52, 0x05, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x40, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x05, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e,
	0x62, 0x6f, 0x78, 0x52, 0x05, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x22, 0x28, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e,
	0x62, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8f, 0x03, 0x0a, 0x0c,
	0x49, 0x6e, 0x62, 0x6f, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x6d, 0x65,
	0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x32, 0x2f, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x12, 0x94, 0x01, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x12, 0x20, 0x2e, 0x6d, 0x65, 0x6d, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x65,
	0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40,
	0xda, 0x41, 0x16, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x2c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a,
	0x05, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x32, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x2f, 0x2a, 0x7d,
	0x12, 0x7b, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x12,
	0x20, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x42, 0xa9, 0x01,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x42, 0x11, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x73, 0x65, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x6d,
	0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x32, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x4d, 0x41, 0x58, 0xaa,
	0x02, 0x0c, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x32, 0xca, 0x02,
	0x0c, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x18,
	0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x4d, 0x65, 0x6d, 0x6f, 0x73,
	0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_v2_inbox_service_proto_rawDescOnce sync.Once
	file_api_v2_inbox_service_proto_rawDescData = file_api_v2_inbox_service_proto_rawDesc
)

func file_api_v2_inbox_service_proto_rawDescGZIP() []byte {
	file_api_v2_inbox_service_proto_rawDescOnce.Do(func() {
		file_api_v2_inbox_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v2_inbox_service_proto_rawDescData)
	})
	return file_api_v2_inbox_service_proto_rawDescData
}

var file_api_v2_inbox_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_v2_inbox_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_v2_inbox_service_proto_goTypes = []interface{}{
	(Inbox_Status)(0),             // 0: memos.api.v2.Inbox.Status
	(Inbox_Type)(0),               // 1: memos.api.v2.Inbox.Type
	(*Inbox)(nil),                 // 2: memos.api.v2.Inbox
	(*ListInboxesRequest)(nil),    // 3: memos.api.v2.ListInboxesRequest
	(*ListInboxesResponse)(nil),   // 4: memos.api.v2.ListInboxesResponse
	(*UpdateInboxRequest)(nil),    // 5: memos.api.v2.UpdateInboxRequest
	(*UpdateInboxResponse)(nil),   // 6: memos.api.v2.UpdateInboxResponse
	(*DeleteInboxRequest)(nil),    // 7: memos.api.v2.DeleteInboxRequest
	(*DeleteInboxResponse)(nil),   // 8: memos.api.v2.DeleteInboxResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil), // 10: google.protobuf.FieldMask
}
var file_api_v2_inbox_service_proto_depIdxs = []int32{
	0,  // 0: memos.api.v2.Inbox.status:type_name -> memos.api.v2.Inbox.Status
	9,  // 1: memos.api.v2.Inbox.create_time:type_name -> google.protobuf.Timestamp
	1,  // 2: memos.api.v2.Inbox.type:type_name -> memos.api.v2.Inbox.Type
	2,  // 3: memos.api.v2.ListInboxesResponse.inboxes:type_name -> memos.api.v2.Inbox
	2,  // 4: memos.api.v2.UpdateInboxRequest.inbox:type_name -> memos.api.v2.Inbox
	10, // 5: memos.api.v2.UpdateInboxRequest.update_mask:type_name -> google.protobuf.FieldMask
	2,  // 6: memos.api.v2.UpdateInboxResponse.inbox:type_name -> memos.api.v2.Inbox
	3,  // 7: memos.api.v2.InboxService.ListInboxes:input_type -> memos.api.v2.ListInboxesRequest
	5,  // 8: memos.api.v2.InboxService.UpdateInbox:input_type -> memos.api.v2.UpdateInboxRequest
	7,  // 9: memos.api.v2.InboxService.DeleteInbox:input_type -> memos.api.v2.DeleteInboxRequest
	4,  // 10: memos.api.v2.InboxService.ListInboxes:output_type -> memos.api.v2.ListInboxesResponse
	6,  // 11: memos.api.v2.InboxService.UpdateInbox:output_type -> memos.api.v2.UpdateInboxResponse
	8,  // 12: memos.api.v2.InboxService.DeleteInbox:output_type -> memos.api.v2.DeleteInboxResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_v2_inbox_service_proto_init() }
func file_api_v2_inbox_service_proto_init() {
	if File_api_v2_inbox_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v2_inbox_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inbox); i {
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
		file_api_v2_inbox_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInboxesRequest); i {
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
		file_api_v2_inbox_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInboxesResponse); i {
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
		file_api_v2_inbox_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInboxRequest); i {
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
		file_api_v2_inbox_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInboxResponse); i {
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
		file_api_v2_inbox_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInboxRequest); i {
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
		file_api_v2_inbox_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInboxResponse); i {
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
	file_api_v2_inbox_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v2_inbox_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v2_inbox_service_proto_goTypes,
		DependencyIndexes: file_api_v2_inbox_service_proto_depIdxs,
		EnumInfos:         file_api_v2_inbox_service_proto_enumTypes,
		MessageInfos:      file_api_v2_inbox_service_proto_msgTypes,
	}.Build()
	File_api_v2_inbox_service_proto = out.File
	file_api_v2_inbox_service_proto_rawDesc = nil
	file_api_v2_inbox_service_proto_goTypes = nil
	file_api_v2_inbox_service_proto_depIdxs = nil
}
