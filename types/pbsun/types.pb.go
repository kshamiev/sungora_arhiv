// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: types/pbsun/types.proto

package pbsun

import (
	reflect "reflect"
	sync "sync"

	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text      string               `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Texts     []string             `protobuf:"bytes,2,rep,name=texts,proto3" json:"texts,omitempty"`
	Tests     map[string]*Test     `protobuf:"bytes,3,rep,name=tests,proto3" json:"tests,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Empty     *empty.Empty         `protobuf:"bytes,4,opt,name=empty,proto3" json:"empty,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_pbsun_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_types_pbsun_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_types_pbsun_types_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Test) GetTexts() []string {
	if x != nil {
		return x.Texts
	}
	return nil
}

func (x *Test) GetTests() map[string]*Test {
	if x != nil {
		return x.Tests
	}
	return nil
}

func (x *Test) GetEmpty() *empty.Empty {
	if x != nil {
		return x.Empty
	}
	return nil
}

func (x *Test) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type GooseDBVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	VersionId int64                `protobuf:"varint,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	IsApplied bool                 `protobuf:"varint,3,opt,name=is_applied,json=isApplied,proto3" json:"is_applied,omitempty"`
	Tstamp    *timestamp.Timestamp `protobuf:"bytes,4,opt,name=tstamp,proto3" json:"tstamp,omitempty"`
}

func (x *GooseDBVersion) Reset() {
	*x = GooseDBVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_pbsun_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GooseDBVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GooseDBVersion) ProtoMessage() {}

func (x *GooseDBVersion) ProtoReflect() protoreflect.Message {
	mi := &file_types_pbsun_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GooseDBVersion.ProtoReflect.Descriptor instead.
func (*GooseDBVersion) Descriptor() ([]byte, []int) {
	return file_types_pbsun_types_proto_rawDescGZIP(), []int{1}
}

func (x *GooseDBVersion) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GooseDBVersion) GetVersionId() int64 {
	if x != nil {
		return x.VersionId
	}
	return 0
}

func (x *GooseDBVersion) GetIsApplied() bool {
	if x != nil {
		return x.IsApplied
	}
	return false
}

func (x *GooseDBVersion) GetTstamp() *timestamp.Timestamp {
	if x != nil {
		return x.Tstamp
	}
	return nil
}

type GooseDBVersionSlice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slice []*GooseDBVersion `protobuf:"bytes,1,rep,name=slice,proto3" json:"slice,omitempty"`
}

func (x *GooseDBVersionSlice) Reset() {
	*x = GooseDBVersionSlice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_pbsun_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GooseDBVersionSlice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GooseDBVersionSlice) ProtoMessage() {}

func (x *GooseDBVersionSlice) ProtoReflect() protoreflect.Message {
	mi := &file_types_pbsun_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GooseDBVersionSlice.ProtoReflect.Descriptor instead.
func (*GooseDBVersionSlice) Descriptor() ([]byte, []int) {
	return file_types_pbsun_types_proto_rawDescGZIP(), []int{2}
}

func (x *GooseDBVersionSlice) GetSlice() []*GooseDBVersion {
	if x != nil {
		return x.Slice
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string               `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Number    int64                `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	Status    string               `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_pbsun_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_types_pbsun_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_types_pbsun_types_proto_rawDescGZIP(), []int{3}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Order) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Order) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type OrderSlice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slice []*Order `protobuf:"bytes,1,rep,name=slice,proto3" json:"slice,omitempty"`
}

func (x *OrderSlice) Reset() {
	*x = OrderSlice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_pbsun_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderSlice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderSlice) ProtoMessage() {}

func (x *OrderSlice) ProtoReflect() protoreflect.Message {
	mi := &file_types_pbsun_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderSlice.ProtoReflect.Descriptor instead.
func (*OrderSlice) Descriptor() ([]byte, []int) {
	return file_types_pbsun_types_proto_rawDescGZIP(), []int{4}
}

func (x *OrderSlice) GetSlice() []*Order {
	if x != nil {
		return x.Slice
	}
	return nil
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code        string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_pbsun_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_types_pbsun_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_types_pbsun_types_proto_rawDescGZIP(), []int{5}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Role) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type RoleSlice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slice []*Role `protobuf:"bytes,1,rep,name=slice,proto3" json:"slice,omitempty"`
}

func (x *RoleSlice) Reset() {
	*x = RoleSlice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_pbsun_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleSlice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleSlice) ProtoMessage() {}

func (x *RoleSlice) ProtoReflect() protoreflect.Message {
	mi := &file_types_pbsun_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleSlice.ProtoReflect.Descriptor instead.
func (*RoleSlice) Descriptor() ([]byte, []int) {
	return file_types_pbsun_types_proto_rawDescGZIP(), []int{6}
}

func (x *RoleSlice) GetSlice() []*Role {
	if x != nil {
		return x.Slice
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Login     string               `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Email     string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Price     string               `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	SummaOne  float32              `protobuf:"fixed32,5,opt,name=summa_one,json=summaOne,proto3" json:"summa_one,omitempty"`
	SummaTwo  float64              `protobuf:"fixed64,6,opt,name=summa_two,json=summaTwo,proto3" json:"summa_two,omitempty"`
	Cnt2      int32                `protobuf:"varint,7,opt,name=cnt2,proto3" json:"cnt2,omitempty"`
	Cnt4      int64                `protobuf:"varint,8,opt,name=cnt4,proto3" json:"cnt4,omitempty"`
	Cnt8      int64                `protobuf:"varint,9,opt,name=cnt8,proto3" json:"cnt8,omitempty"`
	IsOnline  bool                 `protobuf:"varint,10,opt,name=is_online,json=isOnline,proto3" json:"is_online,omitempty"`
	Metrika   []byte               `protobuf:"bytes,11,opt,name=metrika,proto3" json:"metrika,omitempty"`
	Duration  int64                `protobuf:"varint,12,opt,name=duration,proto3" json:"duration,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamp.Timestamp `protobuf:"bytes,15,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_pbsun_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_types_pbsun_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_types_pbsun_types_proto_rawDescGZIP(), []int{7}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *User) GetSummaOne() float32 {
	if x != nil {
		return x.SummaOne
	}
	return 0
}

func (x *User) GetSummaTwo() float64 {
	if x != nil {
		return x.SummaTwo
	}
	return 0
}

func (x *User) GetCnt2() int32 {
	if x != nil {
		return x.Cnt2
	}
	return 0
}

func (x *User) GetCnt4() int64 {
	if x != nil {
		return x.Cnt4
	}
	return 0
}

func (x *User) GetCnt8() int64 {
	if x != nil {
		return x.Cnt8
	}
	return 0
}

func (x *User) GetIsOnline() bool {
	if x != nil {
		return x.IsOnline
	}
	return false
}

func (x *User) GetMetrika() []byte {
	if x != nil {
		return x.Metrika
	}
	return nil
}

func (x *User) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *User) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *User) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type UserSlice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slice []*User `protobuf:"bytes,1,rep,name=slice,proto3" json:"slice,omitempty"`
}

func (x *UserSlice) Reset() {
	*x = UserSlice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_pbsun_types_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSlice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSlice) ProtoMessage() {}

func (x *UserSlice) ProtoReflect() protoreflect.Message {
	mi := &file_types_pbsun_types_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSlice.ProtoReflect.Descriptor instead.
func (*UserSlice) Descriptor() ([]byte, []int) {
	return file_types_pbsun_types_proto_rawDescGZIP(), []int{8}
}

func (x *UserSlice) GetSlice() []*User {
	if x != nil {
		return x.Slice
	}
	return nil
}

var File_types_pbsun_types_proto protoreflect.FileDescriptor

var file_types_pbsun_types_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x73, 0x75, 0x6e, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x62, 0x73, 0x75, 0x6e,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e,
	0x02, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x65, 0x78, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x65, 0x78, 0x74,
	0x73, 0x12, 0x2c, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x70, 0x62, 0x73, 0x75, 0x6e, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x12,
	0x2c, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x45, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x73, 0x75, 0x6e, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x92, 0x01, 0x0a, 0x0e, 0x47, 0x6f, 0x6f, 0x73, 0x65, 0x44, 0x42, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64,
	0x12, 0x32, 0x0a, 0x06, 0x74, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x74, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x42, 0x0a, 0x13, 0x47, 0x6f, 0x6f, 0x73, 0x65, 0x44, 0x42, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x73,
	0x6c, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x73,
	0x75, 0x6e, 0x2e, 0x47, 0x6f, 0x6f, 0x73, 0x65, 0x44, 0x42, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x05, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x22, 0x91, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x30, 0x0a, 0x0a,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x6c,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x73, 0x75,
	0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x22, 0x4c,
	0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2e, 0x0a, 0x09,
	0x52, 0x6f, 0x6c, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x6c, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x73, 0x75, 0x6e,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x22, 0xd2, 0x03, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x4f, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x5f, 0x74, 0x77,
	0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x54, 0x77,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6e, 0x74, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6e, 0x74, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6e, 0x74, 0x34, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6e, 0x74, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6e, 0x74,
	0x38, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6e, 0x74, 0x38, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x6b, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x6b, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x2e, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x21,
	0x0a, 0x05, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x62, 0x73, 0x75, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x73, 0x6c, 0x69, 0x63,
	0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x73, 0x75, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_pbsun_types_proto_rawDescOnce sync.Once
	file_types_pbsun_types_proto_rawDescData = file_types_pbsun_types_proto_rawDesc
)

func file_types_pbsun_types_proto_rawDescGZIP() []byte {
	file_types_pbsun_types_proto_rawDescOnce.Do(func() {
		file_types_pbsun_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_pbsun_types_proto_rawDescData)
	})
	return file_types_pbsun_types_proto_rawDescData
}

var file_types_pbsun_types_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_types_pbsun_types_proto_goTypes = []interface{}{
	(*Test)(nil),                // 0: pbsun.Test
	(*GooseDBVersion)(nil),      // 1: pbsun.GooseDBVersion
	(*GooseDBVersionSlice)(nil), // 2: pbsun.GooseDBVersionSlice
	(*Order)(nil),               // 3: pbsun.Order
	(*OrderSlice)(nil),          // 4: pbsun.OrderSlice
	(*Role)(nil),                // 5: pbsun.Role
	(*RoleSlice)(nil),           // 6: pbsun.RoleSlice
	(*User)(nil),                // 7: pbsun.User
	(*UserSlice)(nil),           // 8: pbsun.UserSlice
	nil,                         // 9: pbsun.Test.TestsEntry
	(*empty.Empty)(nil),         // 10: google.protobuf.Empty
	(*timestamp.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_types_pbsun_types_proto_depIdxs = []int32{
	9,  // 0: pbsun.Test.tests:type_name -> pbsun.Test.TestsEntry
	10, // 1: pbsun.Test.empty:type_name -> google.protobuf.Empty
	11, // 2: pbsun.Test.created_at:type_name -> google.protobuf.Timestamp
	11, // 3: pbsun.GooseDBVersion.tstamp:type_name -> google.protobuf.Timestamp
	1,  // 4: pbsun.GooseDBVersionSlice.slice:type_name -> pbsun.GooseDBVersion
	11, // 5: pbsun.Order.created_at:type_name -> google.protobuf.Timestamp
	11, // 6: pbsun.Order.updated_at:type_name -> google.protobuf.Timestamp
	11, // 7: pbsun.Order.deleted_at:type_name -> google.protobuf.Timestamp
	3,  // 8: pbsun.OrderSlice.slice:type_name -> pbsun.Order
	5,  // 9: pbsun.RoleSlice.slice:type_name -> pbsun.Role
	11, // 10: pbsun.User.created_at:type_name -> google.protobuf.Timestamp
	11, // 11: pbsun.User.updated_at:type_name -> google.protobuf.Timestamp
	11, // 12: pbsun.User.deleted_at:type_name -> google.protobuf.Timestamp
	7,  // 13: pbsun.UserSlice.slice:type_name -> pbsun.User
	0,  // 14: pbsun.Test.TestsEntry.value:type_name -> pbsun.Test
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_types_pbsun_types_proto_init() }
func file_types_pbsun_types_proto_init() {
	if File_types_pbsun_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_pbsun_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
		file_types_pbsun_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GooseDBVersion); i {
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
		file_types_pbsun_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GooseDBVersionSlice); i {
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
		file_types_pbsun_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_types_pbsun_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderSlice); i {
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
		file_types_pbsun_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_types_pbsun_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleSlice); i {
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
		file_types_pbsun_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_types_pbsun_types_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSlice); i {
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
			RawDescriptor: file_types_pbsun_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_pbsun_types_proto_goTypes,
		DependencyIndexes: file_types_pbsun_types_proto_depIdxs,
		MessageInfos:      file_types_pbsun_types_proto_msgTypes,
	}.Build()
	File_types_pbsun_types_proto = out.File
	file_types_pbsun_types_proto_rawDesc = nil
	file_types_pbsun_types_proto_goTypes = nil
	file_types_pbsun_types_proto_depIdxs = nil
}
