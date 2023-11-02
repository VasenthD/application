// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/channels.proto

package BaLa071

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

type AddChannelReq_ChanType int32

const (
	AddChannelReq_SFTP AddChannelReq_ChanType = 0
	AddChannelReq_S3   AddChannelReq_ChanType = 1
)

// Enum value maps for AddChannelReq_ChanType.
var (
	AddChannelReq_ChanType_name = map[int32]string{
		0: "SFTP",
		1: "S3",
	}
	AddChannelReq_ChanType_value = map[string]int32{
		"SFTP": 0,
		"S3":   1,
	}
)

func (x AddChannelReq_ChanType) Enum() *AddChannelReq_ChanType {
	p := new(AddChannelReq_ChanType)
	*p = x
	return p
}

func (x AddChannelReq_ChanType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddChannelReq_ChanType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_channels_proto_enumTypes[0].Descriptor()
}

func (AddChannelReq_ChanType) Type() protoreflect.EnumType {
	return &file_proto_channels_proto_enumTypes[0]
}

func (x AddChannelReq_ChanType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddChannelReq_ChanType.Descriptor instead.
func (AddChannelReq_ChanType) EnumDescriptor() ([]byte, []int) {
	return file_proto_channels_proto_rawDescGZIP(), []int{0, 0}
}

type AddChannelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string                 `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	ChannelType        AddChannelReq_ChanType `protobuf:"varint,2,opt,name=ChannelType,proto3,enum=channels.AddChannelReq_ChanType" json:"ChannelType,omitempty"`
	ServerIP           string                 `protobuf:"bytes,3,opt,name=ServerIP,proto3" json:"ServerIP,omitempty"`
	AuthenticationType string                 `protobuf:"bytes,4,opt,name=AuthenticationType,proto3" json:"AuthenticationType,omitempty"`
	Username           string                 `protobuf:"bytes,5,opt,name=Username,proto3" json:"Username,omitempty"`
	Password           string                 `protobuf:"bytes,6,opt,name=Password,proto3" json:"Password,omitempty"`
	PrivateKey         string                 `protobuf:"bytes,7,opt,name=PrivateKey,proto3" json:"PrivateKey,omitempty"`
	Folders            []*Fold                `protobuf:"bytes,8,rep,name=Folders,proto3" json:"Folders,omitempty"`
}

func (x *AddChannelReq) Reset() {
	*x = AddChannelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddChannelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddChannelReq) ProtoMessage() {}

func (x *AddChannelReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddChannelReq.ProtoReflect.Descriptor instead.
func (*AddChannelReq) Descriptor() ([]byte, []int) {
	return file_proto_channels_proto_rawDescGZIP(), []int{0}
}

func (x *AddChannelReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddChannelReq) GetChannelType() AddChannelReq_ChanType {
	if x != nil {
		return x.ChannelType
	}
	return AddChannelReq_SFTP
}

func (x *AddChannelReq) GetServerIP() string {
	if x != nil {
		return x.ServerIP
	}
	return ""
}

func (x *AddChannelReq) GetAuthenticationType() string {
	if x != nil {
		return x.AuthenticationType
	}
	return ""
}

func (x *AddChannelReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddChannelReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddChannelReq) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *AddChannelReq) GetFolders() []*Fold {
	if x != nil {
		return x.Folders
	}
	return nil
}

type Fold struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`
	GPGKeyID string `protobuf:"bytes,2,opt,name=GPGKeyID,proto3" json:"GPGKeyID,omitempty"`
}

func (x *Fold) Reset() {
	*x = Fold{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fold) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fold) ProtoMessage() {}

func (x *Fold) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fold.ProtoReflect.Descriptor instead.
func (*Fold) Descriptor() ([]byte, []int) {
	return file_proto_channels_proto_rawDescGZIP(), []int{1}
}

func (x *Fold) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Fold) GetGPGKeyID() string {
	if x != nil {
		return x.GPGKeyID
	}
	return ""
}

type ViewChannelRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                 string      `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name               string      `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	ChannelType        string      `protobuf:"bytes,3,opt,name=ChannelType,proto3" json:"ChannelType,omitempty"`
	ServerIP           string      `protobuf:"bytes,4,opt,name=ServerIP,proto3" json:"ServerIP,omitempty"`
	AuthenticationType string      `protobuf:"bytes,5,opt,name=AuthenticationType,proto3" json:"AuthenticationType,omitempty"`
	Username           string      `protobuf:"bytes,6,opt,name=Username,proto3" json:"Username,omitempty"`
	Password           string      `protobuf:"bytes,7,opt,name=Password,proto3" json:"Password,omitempty"`
	PrivateKey         string      `protobuf:"bytes,8,opt,name=PrivateKey,proto3" json:"PrivateKey,omitempty"`
	FoldersResp        []*FoldResp `protobuf:"bytes,9,rep,name=FoldersResp,proto3" json:"FoldersResp,omitempty"`
}

func (x *ViewChannelRes) Reset() {
	*x = ViewChannelRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewChannelRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewChannelRes) ProtoMessage() {}

func (x *ViewChannelRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channels_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewChannelRes.ProtoReflect.Descriptor instead.
func (*ViewChannelRes) Descriptor() ([]byte, []int) {
	return file_proto_channels_proto_rawDescGZIP(), []int{2}
}

func (x *ViewChannelRes) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ViewChannelRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ViewChannelRes) GetChannelType() string {
	if x != nil {
		return x.ChannelType
	}
	return ""
}

func (x *ViewChannelRes) GetServerIP() string {
	if x != nil {
		return x.ServerIP
	}
	return ""
}

func (x *ViewChannelRes) GetAuthenticationType() string {
	if x != nil {
		return x.AuthenticationType
	}
	return ""
}

func (x *ViewChannelRes) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ViewChannelRes) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ViewChannelRes) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *ViewChannelRes) GetFoldersResp() []*FoldResp {
	if x != nil {
		return x.FoldersResp
	}
	return nil
}

type FoldResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoldID   string `protobuf:"bytes,1,opt,name=FoldID,proto3" json:"FoldID,omitempty"`
	Path     string `protobuf:"bytes,2,opt,name=Path,proto3" json:"Path,omitempty"`
	GPGKeyID string `protobuf:"bytes,3,opt,name=GPGKeyID,proto3" json:"GPGKeyID,omitempty"`
}

func (x *FoldResp) Reset() {
	*x = FoldResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channels_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoldResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoldResp) ProtoMessage() {}

func (x *FoldResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channels_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoldResp.ProtoReflect.Descriptor instead.
func (*FoldResp) Descriptor() ([]byte, []int) {
	return file_proto_channels_proto_rawDescGZIP(), []int{3}
}

func (x *FoldResp) GetFoldID() string {
	if x != nil {
		return x.FoldID
	}
	return ""
}

func (x *FoldResp) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FoldResp) GetGPGKeyID() string {
	if x != nil {
		return x.GPGKeyID
	}
	return ""
}

type EmptyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReq) Reset() {
	*x = EmptyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channels_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReq) ProtoMessage() {}

func (x *EmptyReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channels_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReq.ProtoReflect.Descriptor instead.
func (*EmptyReq) Descriptor() ([]byte, []int) {
	return file_proto_channels_proto_rawDescGZIP(), []int{4}
}

type ChanReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelID string `protobuf:"bytes,1,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
}

func (x *ChanReq) Reset() {
	*x = ChanReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channels_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChanReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChanReq) ProtoMessage() {}

func (x *ChanReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channels_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChanReq.ProtoReflect.Descriptor instead.
func (*ChanReq) Descriptor() ([]byte, []int) {
	return file_proto_channels_proto_rawDescGZIP(), []int{5}
}

func (x *ChanReq) GetChannelID() string {
	if x != nil {
		return x.ChannelID
	}
	return ""
}

type ListChannelRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListResp []*ViewChannelRes `protobuf:"bytes,1,rep,name=ListResp,proto3" json:"ListResp,omitempty"`
}

func (x *ListChannelRes) Reset() {
	*x = ListChannelRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channels_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChannelRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChannelRes) ProtoMessage() {}

func (x *ListChannelRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channels_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChannelRes.ProtoReflect.Descriptor instead.
func (*ListChannelRes) Descriptor() ([]byte, []int) {
	return file_proto_channels_proto_rawDescGZIP(), []int{6}
}

func (x *ListChannelRes) GetListResp() []*ViewChannelRes {
	if x != nil {
		return x.ListResp
	}
	return nil
}

var File_proto_channels_proto protoreflect.FileDescriptor

var file_proto_channels_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73,
	0x22, 0xd3, 0x02, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x50, 0x12, 0x2e, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x28,
	0x0a, 0x07, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x52,
	0x07, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x22, 0x1c, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x46, 0x54, 0x50, 0x10, 0x00, 0x12, 0x06,
	0x0a, 0x02, 0x53, 0x33, 0x10, 0x01, 0x22, 0x36, 0x0a, 0x04, 0x46, 0x6f, 0x6c, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x50, 0x47, 0x4b, 0x65, 0x79, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x47, 0x50, 0x47, 0x4b, 0x65, 0x79, 0x49, 0x44, 0x22, 0xb0,
	0x02, 0x0a, 0x0e, 0x56, 0x69, 0x65, 0x77, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x50, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x50, 0x12, 0x2e, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x0b, 0x46,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x46, 0x6f, 0x6c, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x52, 0x0b, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x52, 0x0a, 0x08, 0x46, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x46, 0x6f, 0x6c, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46,
	0x6f, 0x6c, 0x64, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x50, 0x47,
	0x4b, 0x65, 0x79, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x47, 0x50, 0x47,
	0x4b, 0x65, 0x79, 0x49, 0x44, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x22, 0x27, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x22, 0x46, 0x0a, 0x0e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x08,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x52, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x32, 0xca, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x41,
	0x64, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x11, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x73, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x12, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x42,
	0x14, 0x5a, 0x12, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61,
	0x4c, 0x61, 0x30, 0x37, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_channels_proto_rawDescOnce sync.Once
	file_proto_channels_proto_rawDescData = file_proto_channels_proto_rawDesc
)

func file_proto_channels_proto_rawDescGZIP() []byte {
	file_proto_channels_proto_rawDescOnce.Do(func() {
		file_proto_channels_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_channels_proto_rawDescData)
	})
	return file_proto_channels_proto_rawDescData
}

var file_proto_channels_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_channels_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_channels_proto_goTypes = []interface{}{
	(AddChannelReq_ChanType)(0), // 0: channels.AddChannelReq.ChanType
	(*AddChannelReq)(nil),       // 1: channels.AddChannelReq
	(*Fold)(nil),                // 2: channels.Fold
	(*ViewChannelRes)(nil),      // 3: channels.ViewChannelRes
	(*FoldResp)(nil),            // 4: channels.FoldResp
	(*EmptyReq)(nil),            // 5: channels.EmptyReq
	(*ChanReq)(nil),             // 6: channels.ChanReq
	(*ListChannelRes)(nil),      // 7: channels.ListChannelRes
}
var file_proto_channels_proto_depIdxs = []int32{
	0, // 0: channels.AddChannelReq.ChannelType:type_name -> channels.AddChannelReq.ChanType
	2, // 1: channels.AddChannelReq.Folders:type_name -> channels.Fold
	4, // 2: channels.ViewChannelRes.FoldersResp:type_name -> channels.FoldResp
	3, // 3: channels.ListChannelRes.ListResp:type_name -> channels.ViewChannelRes
	1, // 4: channels.ChannelService.AddChannel:input_type -> channels.AddChannelReq
	6, // 5: channels.ChannelService.ViewChannel:input_type -> channels.ChanReq
	5, // 6: channels.ChannelService.ListChannel:input_type -> channels.EmptyReq
	3, // 7: channels.ChannelService.AddChannel:output_type -> channels.ViewChannelRes
	3, // 8: channels.ChannelService.ViewChannel:output_type -> channels.ViewChannelRes
	7, // 9: channels.ChannelService.ListChannel:output_type -> channels.ListChannelRes
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_channels_proto_init() }
func file_proto_channels_proto_init() {
	if File_proto_channels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_channels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddChannelReq); i {
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
		file_proto_channels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fold); i {
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
		file_proto_channels_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewChannelRes); i {
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
		file_proto_channels_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoldResp); i {
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
		file_proto_channels_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReq); i {
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
		file_proto_channels_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChanReq); i {
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
		file_proto_channels_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChannelRes); i {
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
			RawDescriptor: file_proto_channels_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_channels_proto_goTypes,
		DependencyIndexes: file_proto_channels_proto_depIdxs,
		EnumInfos:         file_proto_channels_proto_enumTypes,
		MessageInfos:      file_proto_channels_proto_msgTypes,
	}.Build()
	File_proto_channels_proto = out.File
	file_proto_channels_proto_rawDesc = nil
	file_proto_channels_proto_goTypes = nil
	file_proto_channels_proto_depIdxs = nil
}
