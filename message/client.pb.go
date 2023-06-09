// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.4
// source: client.proto

package message

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

// 登录账号
type LoginAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PacketHead  *IPacket `protobuf:"bytes,1,opt,name=PacketHead,proto3" json:"PacketHead,omitempty"`
	AccountName string   `protobuf:"bytes,2,opt,name=AccountName,proto3" json:"AccountName,omitempty"`
	Password    string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	BuildNo     string   `protobuf:"bytes,5,opt,name=BuildNo,proto3" json:"BuildNo,omitempty"`
	Key         int64    `protobuf:"varint,6,opt,name=Key,proto3" json:"Key,omitempty"` //uint32 Crc = 7;
}

func (x *LoginAccountRequest) Reset() {
	*x = LoginAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginAccountRequest) ProtoMessage() {}

func (x *LoginAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginAccountRequest.ProtoReflect.Descriptor instead.
func (*LoginAccountRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0}
}

func (x *LoginAccountRequest) GetPacketHead() *IPacket {
	if x != nil {
		return x.PacketHead
	}
	return nil
}

func (x *LoginAccountRequest) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *LoginAccountRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginAccountRequest) GetBuildNo() string {
	if x != nil {
		return x.BuildNo
	}
	return ""
}

func (x *LoginAccountRequest) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

// 登录账号反馈
type LoginAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PacketHead  *IPacket `protobuf:"bytes,1,opt,name=PacketHead,proto3" json:"PacketHead,omitempty"`
	Error       int32    `protobuf:"varint,2,opt,name=Error,proto3" json:"Error,omitempty"`
	AccountName string   `protobuf:"bytes,4,opt,name=AccountName,proto3" json:"AccountName,omitempty"`
}

func (x *LoginAccountResponse) Reset() {
	*x = LoginAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginAccountResponse) ProtoMessage() {}

func (x *LoginAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginAccountResponse.ProtoReflect.Descriptor instead.
func (*LoginAccountResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{1}
}

func (x *LoginAccountResponse) GetPacketHead() *IPacket {
	if x != nil {
		return x.PacketHead
	}
	return nil
}

func (x *LoginAccountResponse) GetError() int32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *LoginAccountResponse) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

// 创角
type CreatePlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PacketHead *IPacket `protobuf:"bytes,1,opt,name=PacketHead,proto3" json:"PacketHead,omitempty"`
	PlayerName string   `protobuf:"bytes,2,opt,name=PlayerName,proto3" json:"PlayerName,omitempty"`
	Sex        int32    `protobuf:"varint,3,opt,name=Sex,proto3" json:"Sex,omitempty"`
}

func (x *CreatePlayerRequest) Reset() {
	*x = CreatePlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlayerRequest) ProtoMessage() {}

func (x *CreatePlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlayerRequest.ProtoReflect.Descriptor instead.
func (*CreatePlayerRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePlayerRequest) GetPacketHead() *IPacket {
	if x != nil {
		return x.PacketHead
	}
	return nil
}

func (x *CreatePlayerRequest) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *CreatePlayerRequest) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

// 登录游戏
type LoginPlayerRequset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PacketHead *IPacket `protobuf:"bytes,1,opt,name=PacketHead,proto3" json:"PacketHead,omitempty"`
	PlayerId   int64    `protobuf:"varint,2,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	Key        int64    `protobuf:"varint,3,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (x *LoginPlayerRequset) Reset() {
	*x = LoginPlayerRequset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginPlayerRequset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginPlayerRequset) ProtoMessage() {}

func (x *LoginPlayerRequset) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginPlayerRequset.ProtoReflect.Descriptor instead.
func (*LoginPlayerRequset) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{3}
}

func (x *LoginPlayerRequset) GetPacketHead() *IPacket {
	if x != nil {
		return x.PacketHead
	}
	return nil
}

func (x *LoginPlayerRequset) GetPlayerId() int64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *LoginPlayerRequset) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

// 选角反馈
type SelectPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PacketHead *IPacket      `protobuf:"bytes,1,opt,name=PacketHead,proto3" json:"PacketHead,omitempty"`
	AccountId  int64         `protobuf:"varint,2,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	Key        int64         `protobuf:"varint,3,opt,name=Key,proto3" json:"Key,omitempty"`
	PlayerData []*PlayerData `protobuf:"bytes,5,rep,name=PlayerData,proto3" json:"PlayerData,omitempty"`
}

func (x *SelectPlayerResponse) Reset() {
	*x = SelectPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectPlayerResponse) ProtoMessage() {}

func (x *SelectPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectPlayerResponse.ProtoReflect.Descriptor instead.
func (*SelectPlayerResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{4}
}

func (x *SelectPlayerResponse) GetPacketHead() *IPacket {
	if x != nil {
		return x.PacketHead
	}
	return nil
}

func (x *SelectPlayerResponse) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *SelectPlayerResponse) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *SelectPlayerResponse) GetPlayerData() []*PlayerData {
	if x != nil {
		return x.PlayerData
	}
	return nil
}

// 聊天
type ChatMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PacketHead  *IPacket `protobuf:"bytes,1,opt,name=PacketHead,proto3" json:"PacketHead,omitempty"`
	Sender      int64    `protobuf:"varint,2,opt,name=Sender,proto3" json:"Sender,omitempty"`
	Recver      int64    `protobuf:"varint,3,opt,name=Recver,proto3" json:"Recver,omitempty"`
	MessageType int32    `protobuf:"varint,4,opt,name=MessageType,proto3" json:"MessageType,omitempty"`
	Message     string   `protobuf:"bytes,5,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *ChatMessageRequest) Reset() {
	*x = ChatMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessageRequest) ProtoMessage() {}

func (x *ChatMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessageRequest.ProtoReflect.Descriptor instead.
func (*ChatMessageRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{5}
}

func (x *ChatMessageRequest) GetPacketHead() *IPacket {
	if x != nil {
		return x.PacketHead
	}
	return nil
}

func (x *ChatMessageRequest) GetSender() int64 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *ChatMessageRequest) GetRecver() int64 {
	if x != nil {
		return x.Recver
	}
	return 0
}

func (x *ChatMessageRequest) GetMessageType() int32 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

func (x *ChatMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ChatMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PacketHead  *IPacket `protobuf:"bytes,1,opt,name=PacketHead,proto3" json:"PacketHead,omitempty"`
	Sender      int64    `protobuf:"varint,2,opt,name=Sender,proto3" json:"Sender,omitempty"`
	SenderName  string   `protobuf:"bytes,3,opt,name=SenderName,proto3" json:"SenderName,omitempty"`
	Recver      int64    `protobuf:"varint,4,opt,name=Recver,proto3" json:"Recver,omitempty"`
	RecverName  string   `protobuf:"bytes,5,opt,name=RecverName,proto3" json:"RecverName,omitempty"`
	MessageType int32    `protobuf:"varint,6,opt,name=MessageType,proto3" json:"MessageType,omitempty"`
	Message     string   `protobuf:"bytes,7,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *ChatMessageResponse) Reset() {
	*x = ChatMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessageResponse) ProtoMessage() {}

func (x *ChatMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessageResponse.ProtoReflect.Descriptor instead.
func (*ChatMessageResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{6}
}

func (x *ChatMessageResponse) GetPacketHead() *IPacket {
	if x != nil {
		return x.PacketHead
	}
	return nil
}

func (x *ChatMessageResponse) GetSender() int64 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *ChatMessageResponse) GetSenderName() string {
	if x != nil {
		return x.SenderName
	}
	return ""
}

func (x *ChatMessageResponse) GetRecver() int64 {
	if x != nil {
		return x.Recver
	}
	return 0
}

func (x *ChatMessageResponse) GetRecverName() string {
	if x != nil {
		return x.RecverName
	}
	return ""
}

func (x *ChatMessageResponse) GetMessageType() int32 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

func (x *ChatMessageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_client_proto protoreflect.FileDescriptor

var file_client_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01,
	0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0a, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x49, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x0a, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x4e, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x22, 0x78, 0x0a, 0x14, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x0a, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x49, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x0a, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x71, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0a, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x49, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0a, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x48, 0x65, 0x61, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x53, 0x65, 0x78, 0x22, 0x6c, 0x0a, 0x12, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x73, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x0a,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x49, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0a, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x4b, 0x65, 0x79, 0x22, 0x9d, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x0a, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x49, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0a, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x22, 0xaa, 0x01, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0a, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x49, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0a, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x48, 0x65, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x63, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52,
	0x65, 0x63, 0x76, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0xeb, 0x01, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x49, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0a, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x53,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52,
	0x65, 0x63, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x65, 0x63,
	0x76, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65, 0x63, 0x76, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0x0b, 0x5a, 0x09, 0x2f, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_proto_rawDescOnce sync.Once
	file_client_proto_rawDescData = file_client_proto_rawDesc
)

func file_client_proto_rawDescGZIP() []byte {
	file_client_proto_rawDescOnce.Do(func() {
		file_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_proto_rawDescData)
	})
	return file_client_proto_rawDescData
}

var file_client_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_client_proto_goTypes = []interface{}{
	(*LoginAccountRequest)(nil),  // 0: LoginAccountRequest
	(*LoginAccountResponse)(nil), // 1: LoginAccountResponse
	(*CreatePlayerRequest)(nil),  // 2: CreatePlayerRequest
	(*LoginPlayerRequset)(nil),   // 3: LoginPlayerRequset
	(*SelectPlayerResponse)(nil), // 4: SelectPlayerResponse
	(*ChatMessageRequest)(nil),   // 5: ChatMessageRequest
	(*ChatMessageResponse)(nil),  // 6: ChatMessageResponse
	(*IPacket)(nil),              // 7: IPacket
	(*PlayerData)(nil),           // 8: PlayerData
}
var file_client_proto_depIdxs = []int32{
	7, // 0: LoginAccountRequest.PacketHead:type_name -> IPacket
	7, // 1: LoginAccountResponse.PacketHead:type_name -> IPacket
	7, // 2: CreatePlayerRequest.PacketHead:type_name -> IPacket
	7, // 3: LoginPlayerRequset.PacketHead:type_name -> IPacket
	7, // 4: SelectPlayerResponse.PacketHead:type_name -> IPacket
	8, // 5: SelectPlayerResponse.PlayerData:type_name -> PlayerData
	7, // 6: ChatMessageRequest.PacketHead:type_name -> IPacket
	7, // 7: ChatMessageResponse.PacketHead:type_name -> IPacket
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_client_proto_init() }
func file_client_proto_init() {
	if File_client_proto != nil {
		return
	}
	file_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginAccountRequest); i {
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
		file_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginAccountResponse); i {
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
		file_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlayerRequest); i {
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
		file_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginPlayerRequset); i {
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
		file_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectPlayerResponse); i {
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
		file_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessageRequest); i {
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
		file_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessageResponse); i {
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
			RawDescriptor: file_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_proto_goTypes,
		DependencyIndexes: file_client_proto_depIdxs,
		MessageInfos:      file_client_proto_msgTypes,
	}.Build()
	File_client_proto = out.File
	file_client_proto_rawDesc = nil
	file_client_proto_goTypes = nil
	file_client_proto_depIdxs = nil
}
