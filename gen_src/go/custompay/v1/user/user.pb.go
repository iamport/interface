// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: custompay/v1/user/user.proto

package user

import (
	proto "github.com/golang/protobuf/proto"
	basic "github.com/iamport/interface/gen_src/go/custompay/v1/basic"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Tel      string `protobuf:"bytes,4,opt,name=tel,proto3" json:"tel,omitempty"`
	Address  string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Postcode string `protobuf:"bytes,6,opt,name=postcode,proto3" json:"postcode,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custompay_v1_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_custompay_v1_user_user_proto_msgTypes[0]
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
	return file_custompay_v1_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *User) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *User) GetPostcode() string {
	if x != nil {
		return x.Postcode
	}
	return ""
}

type UserRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Tel      string `protobuf:"bytes,3,opt,name=tel,proto3" json:"tel,omitempty"`
	Address  string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Postcode string `protobuf:"bytes,5,opt,name=postcode,proto3" json:"postcode,omitempty"`
}

func (x *UserRegisterRequest) Reset() {
	*x = UserRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custompay_v1_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterRequest) ProtoMessage() {}

func (x *UserRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_custompay_v1_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterRequest.ProtoReflect.Descriptor instead.
func (*UserRegisterRequest) Descriptor() ([]byte, []int) {
	return file_custompay_v1_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserRegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserRegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRegisterRequest) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *UserRegisterRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UserRegisterRequest) GetPostcode() string {
	if x != nil {
		return x.Postcode
	}
	return ""
}

type UserEditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Tel      string `protobuf:"bytes,4,opt,name=tel,proto3" json:"tel,omitempty"`
	Address  string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Postcode string `protobuf:"bytes,6,opt,name=postcode,proto3" json:"postcode,omitempty"`
}

func (x *UserEditRequest) Reset() {
	*x = UserEditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custompay_v1_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEditRequest) ProtoMessage() {}

func (x *UserEditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_custompay_v1_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEditRequest.ProtoReflect.Descriptor instead.
func (*UserEditRequest) Descriptor() ([]byte, []int) {
	return file_custompay_v1_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserEditRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserEditRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserEditRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserEditRequest) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *UserEditRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UserEditRequest) GetPostcode() string {
	if x != nil {
		return x.Postcode
	}
	return ""
}

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custompay_v1_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_custompay_v1_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_custompay_v1_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Response *User  `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custompay_v1_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_custompay_v1_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_custompay_v1_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserResponse) GetResponse() *User {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_custompay_v1_user_user_proto protoreflect.FileDescriptor

var file_custompay_v1_user_user_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x70, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x70, 0x61, 0x79, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x70, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x65, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x93, 0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6e, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x70, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x92, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x50, 0x43, 0x12, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x70, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x70, 0x61, 0x79, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0a, 0x22, 0x05, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x5c,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x50, 0x43, 0x12, 0x1b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x70, 0x61, 0x79, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x70, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c,
	0x12, 0x0a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5b, 0x0a, 0x0d,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x50, 0x43, 0x12, 0x1b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x70, 0x61, 0x79, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x70, 0x61, 0x79, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x2a, 0x0a, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x60, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x64, 0x69, 0x74, 0x52, 0x50, 0x43, 0x12, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x70, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x64,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x70, 0x61, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x32, 0x0a, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x42, 0x4f, 0x5a, 0x39, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x61, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x5f,
	0x73, 0x72, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x70, 0x61, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0xaa, 0x02, 0x11, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x70, 0x61, 0x79, 0x2e, 0x56, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_custompay_v1_user_user_proto_rawDescOnce sync.Once
	file_custompay_v1_user_user_proto_rawDescData = file_custompay_v1_user_user_proto_rawDesc
)

func file_custompay_v1_user_user_proto_rawDescGZIP() []byte {
	file_custompay_v1_user_user_proto_rawDescOnce.Do(func() {
		file_custompay_v1_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_custompay_v1_user_user_proto_rawDescData)
	})
	return file_custompay_v1_user_user_proto_rawDescData
}

var file_custompay_v1_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_custompay_v1_user_user_proto_goTypes = []interface{}{
	(*User)(nil),                // 0: user_custompay.User
	(*UserRegisterRequest)(nil), // 1: user_custompay.UserRegisterRequest
	(*UserEditRequest)(nil),     // 2: user_custompay.UserEditRequest
	(*UserRequest)(nil),         // 3: user_custompay.UserRequest
	(*UserResponse)(nil),        // 4: user_custompay.UserResponse
	(*basic.Response)(nil),      // 5: basic_custompay.Response
}
var file_custompay_v1_user_user_proto_depIdxs = []int32{
	0, // 0: user_custompay.UserResponse.response:type_name -> user_custompay.User
	1, // 1: user_custompay.UserService.UserRegisterRPC:input_type -> user_custompay.UserRegisterRequest
	3, // 2: user_custompay.UserService.UserInfoRPC:input_type -> user_custompay.UserRequest
	3, // 3: user_custompay.UserService.UserDeleteRPC:input_type -> user_custompay.UserRequest
	2, // 4: user_custompay.UserService.UserEditRPC:input_type -> user_custompay.UserEditRequest
	4, // 5: user_custompay.UserService.UserRegisterRPC:output_type -> user_custompay.UserResponse
	4, // 6: user_custompay.UserService.UserInfoRPC:output_type -> user_custompay.UserResponse
	5, // 7: user_custompay.UserService.UserDeleteRPC:output_type -> basic_custompay.Response
	5, // 8: user_custompay.UserService.UserEditRPC:output_type -> basic_custompay.Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_custompay_v1_user_user_proto_init() }
func file_custompay_v1_user_user_proto_init() {
	if File_custompay_v1_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_custompay_v1_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_custompay_v1_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterRequest); i {
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
		file_custompay_v1_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEditRequest); i {
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
		file_custompay_v1_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_custompay_v1_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
			RawDescriptor: file_custompay_v1_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_custompay_v1_user_user_proto_goTypes,
		DependencyIndexes: file_custompay_v1_user_user_proto_depIdxs,
		MessageInfos:      file_custompay_v1_user_user_proto_msgTypes,
	}.Build()
	File_custompay_v1_user_user_proto = out.File
	file_custompay_v1_user_user_proto_rawDesc = nil
	file_custompay_v1_user_user_proto_goTypes = nil
	file_custompay_v1_user_user_proto_depIdxs = nil
}
