// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: api/pb/proxy/users/user.proto

package users_service_contracts

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

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Login string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Mail  string `protobuf:"bytes,4,opt,name=mail,proto3" json:"mail,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_api_pb_proxy_users_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_proxy_users_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_api_pb_proxy_users_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateUserRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *CreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserRequest) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_api_pb_proxy_users_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_proxy_users_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_api_pb_proxy_users_user_proto_rawDescGZIP(), []int{1}
}

type UserProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Login string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Mail  string `protobuf:"bytes,4,opt,name=mail,proto3" json:"mail,omitempty"`
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	mi := &file_api_pb_proxy_users_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_proxy_users_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_api_pb_proxy_users_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserProfile) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserProfile) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UserProfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserProfile) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

type UsersListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     []uint64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
	Login  []string `protobuf:"bytes,2,rep,name=login,proto3" json:"login,omitempty"`
	Offset uint64   `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  uint64   `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *UsersListRequest) Reset() {
	*x = UsersListRequest{}
	mi := &file_api_pb_proxy_users_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersListRequest) ProtoMessage() {}

func (x *UsersListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_proxy_users_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersListRequest.ProtoReflect.Descriptor instead.
func (*UsersListRequest) Descriptor() ([]byte, []int) {
	return file_api_pb_proxy_users_user_proto_rawDescGZIP(), []int{3}
}

func (x *UsersListRequest) GetId() []uint64 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UsersListRequest) GetLogin() []string {
	if x != nil {
		return x.Login
	}
	return nil
}

func (x *UsersListRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *UsersListRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type UsersListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserProfile `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UsersListResponse) Reset() {
	*x = UsersListResponse{}
	mi := &file_api_pb_proxy_users_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersListResponse) ProtoMessage() {}

func (x *UsersListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_proxy_users_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersListResponse.ProtoReflect.Descriptor instead.
func (*UsersListResponse) Descriptor() ([]byte, []int) {
	return file_api_pb_proxy_users_user_proto_rawDescGZIP(), []int{4}
}

func (x *UsersListResponse) GetUsers() []*UserProfile {
	if x != nil {
		return x.Users
	}
	return nil
}

type UserProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserProfileRequest) Reset() {
	*x = UserProfileRequest{}
	mi := &file_api_pb_proxy_users_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfileRequest) ProtoMessage() {}

func (x *UserProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_proxy_users_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfileRequest.ProtoReflect.Descriptor instead.
func (*UserProfileRequest) Descriptor() ([]byte, []int) {
	return file_api_pb_proxy_users_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserProfileRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserProfile `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserProfileResponse) Reset() {
	*x = UserProfileResponse{}
	mi := &file_api_pb_proxy_users_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfileResponse) ProtoMessage() {}

func (x *UserProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_proxy_users_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfileResponse.ProtoReflect.Descriptor instead.
func (*UserProfileResponse) Descriptor() ([]byte, []int) {
	return file_api_pb_proxy_users_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserProfileResponse) GetUser() *UserProfile {
	if x != nil {
		return x.User
	}
	return nil
}

var File_api_pb_proxy_users_user_proto protoreflect.FileDescriptor

var file_api_pb_proxy_users_user_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x22, 0x61, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x14, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x5b, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x66,
	0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x4f, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4f, 0x0a,
	0x13, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xca,
	0x02, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x65, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2a, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x69, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67,
	0x65, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_pb_proxy_users_user_proto_rawDescOnce sync.Once
	file_api_pb_proxy_users_user_proto_rawDescData = file_api_pb_proxy_users_user_proto_rawDesc
)

func file_api_pb_proxy_users_user_proto_rawDescGZIP() []byte {
	file_api_pb_proxy_users_user_proto_rawDescOnce.Do(func() {
		file_api_pb_proxy_users_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_pb_proxy_users_user_proto_rawDescData)
	})
	return file_api_pb_proxy_users_user_proto_rawDescData
}

var file_api_pb_proxy_users_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_pb_proxy_users_user_proto_goTypes = []any{
	(*CreateUserRequest)(nil),   // 0: users_service_contracts.CreateUserRequest
	(*CreateUserResponse)(nil),  // 1: users_service_contracts.CreateUserResponse
	(*UserProfile)(nil),         // 2: users_service_contracts.UserProfile
	(*UsersListRequest)(nil),    // 3: users_service_contracts.UsersListRequest
	(*UsersListResponse)(nil),   // 4: users_service_contracts.UsersListResponse
	(*UserProfileRequest)(nil),  // 5: users_service_contracts.UserProfileRequest
	(*UserProfileResponse)(nil), // 6: users_service_contracts.UserProfileResponse
}
var file_api_pb_proxy_users_user_proto_depIdxs = []int32{
	2, // 0: users_service_contracts.UsersListResponse.users:type_name -> users_service_contracts.UserProfile
	2, // 1: users_service_contracts.UserProfileResponse.user:type_name -> users_service_contracts.UserProfile
	0, // 2: users_service_contracts.UsersManager.CreateUser:input_type -> users_service_contracts.CreateUserRequest
	5, // 3: users_service_contracts.UsersManager.UserProfile:input_type -> users_service_contracts.UserProfileRequest
	3, // 4: users_service_contracts.UsersManager.GetUsersByFilter:input_type -> users_service_contracts.UsersListRequest
	1, // 5: users_service_contracts.UsersManager.CreateUser:output_type -> users_service_contracts.CreateUserResponse
	6, // 6: users_service_contracts.UsersManager.UserProfile:output_type -> users_service_contracts.UserProfileResponse
	4, // 7: users_service_contracts.UsersManager.GetUsersByFilter:output_type -> users_service_contracts.UsersListResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_pb_proxy_users_user_proto_init() }
func file_api_pb_proxy_users_user_proto_init() {
	if File_api_pb_proxy_users_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_pb_proxy_users_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_pb_proxy_users_user_proto_goTypes,
		DependencyIndexes: file_api_pb_proxy_users_user_proto_depIdxs,
		MessageInfos:      file_api_pb_proxy_users_user_proto_msgTypes,
	}.Build()
	File_api_pb_proxy_users_user_proto = out.File
	file_api_pb_proxy_users_user_proto_rawDesc = nil
	file_api_pb_proxy_users_user_proto_goTypes = nil
	file_api_pb_proxy_users_user_proto_depIdxs = nil
}
