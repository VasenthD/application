// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/application.proto

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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	ChannelId string `protobuf:"bytes,2,opt,name=ChannelId,proto3" json:"ChannelId,omitempty"`
	CreatedBy string `protobuf:"bytes,3,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_application_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *CreateRequest) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=ChannelId,proto3" json:"ChannelId,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_application_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_application_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_application_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_proto_application_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	ChannelId string `protobuf:"bytes,3,opt,name=ChannelId,proto3" json:"ChannelId,omitempty"`
	CreatedBy string `protobuf:"bytes,4,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_application_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_application_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_application_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListResponse) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *ListResponse) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *ListResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ListResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_application_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_application_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_application_proto_rawDescGZIP(), []int{4}
}

type ListAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListAll []*ListResponse `protobuf:"bytes,1,rep,name=ListAll,proto3" json:"ListAll,omitempty"`
}

func (x *ListAllResponse) Reset() {
	*x = ListAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_application_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllResponse) ProtoMessage() {}

func (x *ListAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_application_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllResponse.ProtoReflect.Descriptor instead.
func (*ListAllResponse) Descriptor() ([]byte, []int) {
	return file_proto_application_proto_rawDescGZIP(), []int{5}
}

func (x *ListAllResponse) GetListAll() []*ListResponse {
	if x != nil {
		return x.ListAll
	}
	return nil
}

type PutFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID        string `protobuf:"bytes,1,opt,name=AppID,proto3" json:"AppID,omitempty"`
	FoldID       string `protobuf:"bytes,2,opt,name=FoldID,proto3" json:"FoldID,omitempty"`
	FileName     string `protobuf:"bytes,3,opt,name=FileName,proto3" json:"FileName,omitempty"`
	Chunk        []byte `protobuf:"bytes,4,opt,name=Chunk,proto3" json:"Chunk,omitempty"`
	RequiredSize int32  `protobuf:"varint,5,opt,name=RequiredSize,proto3" json:"RequiredSize,omitempty"`
	StreamedSize int32  `protobuf:"varint,6,opt,name=StreamedSize,proto3" json:"StreamedSize,omitempty"`
}

func (x *PutFileReq) Reset() {
	*x = PutFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_application_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutFileReq) ProtoMessage() {}

func (x *PutFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_application_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutFileReq.ProtoReflect.Descriptor instead.
func (*PutFileReq) Descriptor() ([]byte, []int) {
	return file_proto_application_proto_rawDescGZIP(), []int{6}
}

func (x *PutFileReq) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *PutFileReq) GetFoldID() string {
	if x != nil {
		return x.FoldID
	}
	return ""
}

func (x *PutFileReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *PutFileReq) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *PutFileReq) GetRequiredSize() int32 {
	if x != nil {
		return x.RequiredSize
	}
	return 0
}

func (x *PutFileReq) GetStreamedSize() int32 {
	if x != nil {
		return x.StreamedSize
	}
	return 0
}

type PutFileResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *PutFileResp) Reset() {
	*x = PutFileResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_application_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutFileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutFileResp) ProtoMessage() {}

func (x *PutFileResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_application_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutFileResp.ProtoReflect.Descriptor instead.
func (*PutFileResp) Descriptor() ([]byte, []int) {
	return file_proto_application_proto_rawDescGZIP(), []int{7}
}

func (x *PutFileResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID    string `protobuf:"bytes,1,opt,name=AppID,proto3" json:"AppID,omitempty"`
	FoldID   string `protobuf:"bytes,2,opt,name=FoldID,proto3" json:"FoldID,omitempty"`
	FileName string `protobuf:"bytes,3,opt,name=FileName,proto3" json:"FileName,omitempty"`
}

func (x *GetFileReq) Reset() {
	*x = GetFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_application_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileReq) ProtoMessage() {}

func (x *GetFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_application_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileReq.ProtoReflect.Descriptor instead.
func (*GetFileReq) Descriptor() ([]byte, []int) {
	return file_proto_application_proto_rawDescGZIP(), []int{8}
}

func (x *GetFileReq) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetFileReq) GetFoldID() string {
	if x != nil {
		return x.FoldID
	}
	return ""
}

func (x *GetFileReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type GetFileResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Chunk   []byte `protobuf:"bytes,2,opt,name=Chunk,proto3" json:"Chunk,omitempty"`
}

func (x *GetFileResp) Reset() {
	*x = GetFileResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_application_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileResp) ProtoMessage() {}

func (x *GetFileResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_application_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileResp.ProtoReflect.Descriptor instead.
func (*GetFileResp) Descriptor() ([]byte, []int) {
	return file_proto_application_proto_rawDescGZIP(), []int{9}
}

func (x *GetFileResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetFileResp) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

var File_proto_application_proto protoreflect.FileDescriptor

var file_proto_application_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x2e, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x1d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0xaa, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x46, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x33, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x6c, 0x22, 0xb4, 0x01, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x6f, 0x6c,
	0x64, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x6f, 0x6c, 0x64, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x50,
	0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x56, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x64,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x6f, 0x6c, 0x64, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0xd1, 0x02, 0x0a, 0x12,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x12, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1c, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x07, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x28, 0x01, 0x12,
	0x3e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x30, 0x01, 0x42,
	0x14, 0x5a, 0x12, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61,
	0x4c, 0x61, 0x30, 0x37, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_application_proto_rawDescOnce sync.Once
	file_proto_application_proto_rawDescData = file_proto_application_proto_rawDesc
)

func file_proto_application_proto_rawDescGZIP() []byte {
	file_proto_application_proto_rawDescOnce.Do(func() {
		file_proto_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_application_proto_rawDescData)
	})
	return file_proto_application_proto_rawDescData
}

var file_proto_application_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_application_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),   // 0: Application.CreateRequest
	(*CreateResponse)(nil),  // 1: Application.CreateResponse
	(*ListRequest)(nil),     // 2: Application.ListRequest
	(*ListResponse)(nil),    // 3: Application.ListResponse
	(*Empty)(nil),           // 4: Application.empty
	(*ListAllResponse)(nil), // 5: Application.ListAllResponse
	(*PutFileReq)(nil),      // 6: Application.PutFileReq
	(*PutFileResp)(nil),     // 7: Application.PutFileResp
	(*GetFileReq)(nil),      // 8: Application.GetFileReq
	(*GetFileResp)(nil),     // 9: Application.GetFileResp
}
var file_proto_application_proto_depIdxs = []int32{
	3, // 0: Application.ListAllResponse.ListAll:type_name -> Application.ListResponse
	0, // 1: Application.ApplicationService.Create:input_type -> Application.CreateRequest
	2, // 2: Application.ApplicationService.List:input_type -> Application.ListRequest
	4, // 3: Application.ApplicationService.ListAll:input_type -> Application.empty
	6, // 4: Application.ApplicationService.PutFile:input_type -> Application.PutFileReq
	8, // 5: Application.ApplicationService.GetFile:input_type -> Application.GetFileReq
	1, // 6: Application.ApplicationService.Create:output_type -> Application.CreateResponse
	3, // 7: Application.ApplicationService.List:output_type -> Application.ListResponse
	5, // 8: Application.ApplicationService.ListAll:output_type -> Application.ListAllResponse
	7, // 9: Application.ApplicationService.PutFile:output_type -> Application.PutFileResp
	9, // 10: Application.ApplicationService.GetFile:output_type -> Application.GetFileResp
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_application_proto_init() }
func file_proto_application_proto_init() {
	if File_proto_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_proto_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_proto_application_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_proto_application_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_proto_application_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_application_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllResponse); i {
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
		file_proto_application_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutFileReq); i {
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
		file_proto_application_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutFileResp); i {
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
		file_proto_application_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileReq); i {
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
		file_proto_application_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileResp); i {
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
			RawDescriptor: file_proto_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_application_proto_goTypes,
		DependencyIndexes: file_proto_application_proto_depIdxs,
		MessageInfos:      file_proto_application_proto_msgTypes,
	}.Build()
	File_proto_application_proto = out.File
	file_proto_application_proto_rawDesc = nil
	file_proto_application_proto_goTypes = nil
	file_proto_application_proto_depIdxs = nil
}
