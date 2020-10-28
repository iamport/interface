// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: v2/authenticate/authenticate.proto

package authenticate

import (
	proto "github.com/golang/protobuf/proto"
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

type TokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImpKey    string `protobuf:"bytes,1,opt,name=imp_key,json=impKey,proto3" json:"imp_key,omitempty"`
	ImpSecret string `protobuf:"bytes,2,opt,name=imp_secret,json=impSecret,proto3" json:"imp_secret,omitempty"`
}

func (x *TokenRequest) Reset() {
	*x = TokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_authenticate_authenticate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRequest) ProtoMessage() {}

func (x *TokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_authenticate_authenticate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRequest.ProtoReflect.Descriptor instead.
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return file_v2_authenticate_authenticate_proto_rawDescGZIP(), []int{0}
}

func (x *TokenRequest) GetImpKey() string {
	if x != nil {
		return x.ImpKey
	}
	return ""
}

func (x *TokenRequest) GetImpSecret() string {
	if x != nil {
		return x.ImpSecret
	}
	return ""
}

type TokenData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ExpiredAt   int32  `protobuf:"varint,2,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	Now         int32  `protobuf:"varint,3,opt,name=now,proto3" json:"now,omitempty"`
}

func (x *TokenData) Reset() {
	*x = TokenData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_authenticate_authenticate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenData) ProtoMessage() {}

func (x *TokenData) ProtoReflect() protoreflect.Message {
	mi := &file_v2_authenticate_authenticate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenData.ProtoReflect.Descriptor instead.
func (*TokenData) Descriptor() ([]byte, []int) {
	return file_v2_authenticate_authenticate_proto_rawDescGZIP(), []int{1}
}

func (x *TokenData) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *TokenData) GetExpiredAt() int32 {
	if x != nil {
		return x.ExpiredAt
	}
	return 0
}

func (x *TokenData) GetNow() int32 {
	if x != nil {
		return x.Now
	}
	return 0
}

type TokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message  string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Response *TokenData `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *TokenResponse) Reset() {
	*x = TokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_authenticate_authenticate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenResponse) ProtoMessage() {}

func (x *TokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_authenticate_authenticate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenResponse.ProtoReflect.Descriptor instead.
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return file_v2_authenticate_authenticate_proto_rawDescGZIP(), []int{2}
}

func (x *TokenResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *TokenResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TokenResponse) GetResponse() *TokenData {
	if x != nil {
		return x.Response
	}
	return nil
}

type PubKeyRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImpKey    string `protobuf:"bytes,1,opt,name=imp_key,json=impKey,proto3" json:"imp_key,omitempty"`
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"` // want to make it work without session dependency
}

func (x *PubKeyRegisterRequest) Reset() {
	*x = PubKeyRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_authenticate_authenticate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubKeyRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubKeyRegisterRequest) ProtoMessage() {}

func (x *PubKeyRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_authenticate_authenticate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubKeyRegisterRequest.ProtoReflect.Descriptor instead.
func (*PubKeyRegisterRequest) Descriptor() ([]byte, []int) {
	return file_v2_authenticate_authenticate_proto_rawDescGZIP(), []int{3}
}

func (x *PubKeyRegisterRequest) GetImpKey() string {
	if x != nil {
		return x.ImpKey
	}
	return ""
}

func (x *PubKeyRegisterRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *PubKeyRegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type PubKeyRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PubKeyRegisterResponse) Reset() {
	*x = PubKeyRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_authenticate_authenticate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubKeyRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubKeyRegisterResponse) ProtoMessage() {}

func (x *PubKeyRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_authenticate_authenticate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubKeyRegisterResponse.ProtoReflect.Descriptor instead.
func (*PubKeyRegisterResponse) Descriptor() ([]byte, []int) {
	return file_v2_authenticate_authenticate_proto_rawDescGZIP(), []int{4}
}

func (x *PubKeyRegisterResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PubKeyRegisterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_v2_authenticate_authenticate_proto protoreflect.FileDescriptor

var file_v2_authenticate_authenticate_proto_rawDesc = []byte{
	0x0a, 0x22, 0x76, 0x32, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x5f, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6d, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x6d, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6d, 0x70, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x5f, 0x0a, 0x09, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x6f,
	0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x6f, 0x77, 0x22, 0x75, 0x0a, 0x0d,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x32, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x6b, 0x0a, 0x15, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x6d, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x6d, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x46, 0x0a, 0x16, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x91, 0x02, 0x0a, 0x13, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x6e, 0x0a, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x32, 0x52, 0x50, 0x43, 0x12, 0x1d,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x32,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x32, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x01, 0x2a,
	0x12, 0x89, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x75, 0x62,
	0x4b, 0x65, 0x79, 0x52, 0x50, 0x43, 0x12, 0x26, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x32, 0x2e, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x32,
	0x2e, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22,
	0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x3a, 0x01, 0x2a, 0x42, 0x4b, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x61, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e,
	0x5f, 0x73, 0x72, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0xaa, 0x02, 0x0f, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v2_authenticate_authenticate_proto_rawDescOnce sync.Once
	file_v2_authenticate_authenticate_proto_rawDescData = file_v2_authenticate_authenticate_proto_rawDesc
)

func file_v2_authenticate_authenticate_proto_rawDescGZIP() []byte {
	file_v2_authenticate_authenticate_proto_rawDescOnce.Do(func() {
		file_v2_authenticate_authenticate_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_authenticate_authenticate_proto_rawDescData)
	})
	return file_v2_authenticate_authenticate_proto_rawDescData
}

var file_v2_authenticate_authenticate_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v2_authenticate_authenticate_proto_goTypes = []interface{}{
	(*TokenRequest)(nil),           // 0: authenticate_v2.TokenRequest
	(*TokenData)(nil),              // 1: authenticate_v2.TokenData
	(*TokenResponse)(nil),          // 2: authenticate_v2.TokenResponse
	(*PubKeyRegisterRequest)(nil),  // 3: authenticate_v2.PubKeyRegisterRequest
	(*PubKeyRegisterResponse)(nil), // 4: authenticate_v2.PubKeyRegisterResponse
}
var file_v2_authenticate_authenticate_proto_depIdxs = []int32{
	1, // 0: authenticate_v2.TokenResponse.response:type_name -> authenticate_v2.TokenData
	0, // 1: authenticate_v2.AuthenticateService.TokenV2RPC:input_type -> authenticate_v2.TokenRequest
	3, // 2: authenticate_v2.AuthenticateService.RegisterPubKeyRPC:input_type -> authenticate_v2.PubKeyRegisterRequest
	2, // 3: authenticate_v2.AuthenticateService.TokenV2RPC:output_type -> authenticate_v2.TokenResponse
	4, // 4: authenticate_v2.AuthenticateService.RegisterPubKeyRPC:output_type -> authenticate_v2.PubKeyRegisterResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v2_authenticate_authenticate_proto_init() }
func file_v2_authenticate_authenticate_proto_init() {
	if File_v2_authenticate_authenticate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_authenticate_authenticate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRequest); i {
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
		file_v2_authenticate_authenticate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenData); i {
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
		file_v2_authenticate_authenticate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenResponse); i {
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
		file_v2_authenticate_authenticate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubKeyRegisterRequest); i {
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
		file_v2_authenticate_authenticate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubKeyRegisterResponse); i {
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
			RawDescriptor: file_v2_authenticate_authenticate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v2_authenticate_authenticate_proto_goTypes,
		DependencyIndexes: file_v2_authenticate_authenticate_proto_depIdxs,
		MessageInfos:      file_v2_authenticate_authenticate_proto_msgTypes,
	}.Build()
	File_v2_authenticate_authenticate_proto = out.File
	file_v2_authenticate_authenticate_proto_rawDesc = nil
	file_v2_authenticate_authenticate_proto_goTypes = nil
	file_v2_authenticate_authenticate_proto_depIdxs = nil
}
