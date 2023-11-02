// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/application.proto

package BaLa071

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ApplicationServiceClient is the client API for ApplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApplicationServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	ListAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListAllResponse, error)
	PutFile(ctx context.Context, opts ...grpc.CallOption) (ApplicationService_PutFileClient, error)
	GetFile(ctx context.Context, in *GetFileReq, opts ...grpc.CallOption) (ApplicationService_GetFileClient, error)
}

type applicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationServiceClient(cc grpc.ClientConnInterface) ApplicationServiceClient {
	return &applicationServiceClient{cc}
}

func (c *applicationServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/Application.ApplicationService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/Application.ApplicationService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) ListAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListAllResponse, error) {
	out := new(ListAllResponse)
	err := c.cc.Invoke(ctx, "/Application.ApplicationService/ListAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) PutFile(ctx context.Context, opts ...grpc.CallOption) (ApplicationService_PutFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &ApplicationService_ServiceDesc.Streams[0], "/Application.ApplicationService/PutFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &applicationServicePutFileClient{stream}
	return x, nil
}

type ApplicationService_PutFileClient interface {
	Send(*PutFileReq) error
	CloseAndRecv() (*PutFileResp, error)
	grpc.ClientStream
}

type applicationServicePutFileClient struct {
	grpc.ClientStream
}

func (x *applicationServicePutFileClient) Send(m *PutFileReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *applicationServicePutFileClient) CloseAndRecv() (*PutFileResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PutFileResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *applicationServiceClient) GetFile(ctx context.Context, in *GetFileReq, opts ...grpc.CallOption) (ApplicationService_GetFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &ApplicationService_ServiceDesc.Streams[1], "/Application.ApplicationService/GetFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &applicationServiceGetFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ApplicationService_GetFileClient interface {
	Recv() (*GetFileResp, error)
	grpc.ClientStream
}

type applicationServiceGetFileClient struct {
	grpc.ClientStream
}

func (x *applicationServiceGetFileClient) Recv() (*GetFileResp, error) {
	m := new(GetFileResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ApplicationServiceServer is the server API for ApplicationService service.
// All implementations must embed UnimplementedApplicationServiceServer
// for forward compatibility
type ApplicationServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	ListAll(context.Context, *Empty) (*ListAllResponse, error)
	PutFile(ApplicationService_PutFileServer) error
	GetFile(*GetFileReq, ApplicationService_GetFileServer) error
	mustEmbedUnimplementedApplicationServiceServer()
}

// UnimplementedApplicationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApplicationServiceServer struct {
}

func (UnimplementedApplicationServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedApplicationServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedApplicationServiceServer) ListAll(context.Context, *Empty) (*ListAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAll not implemented")
}
func (UnimplementedApplicationServiceServer) PutFile(ApplicationService_PutFileServer) error {
	return status.Errorf(codes.Unimplemented, "method PutFile not implemented")
}
func (UnimplementedApplicationServiceServer) GetFile(*GetFileReq, ApplicationService_GetFileServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedApplicationServiceServer) mustEmbedUnimplementedApplicationServiceServer() {}

// UnsafeApplicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApplicationServiceServer will
// result in compilation errors.
type UnsafeApplicationServiceServer interface {
	mustEmbedUnimplementedApplicationServiceServer()
}

func RegisterApplicationServiceServer(s grpc.ServiceRegistrar, srv ApplicationServiceServer) {
	s.RegisterService(&ApplicationService_ServiceDesc, srv)
}

func _ApplicationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Application.ApplicationService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Application.ApplicationService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_ListAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).ListAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Application.ApplicationService/ListAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).ListAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_PutFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ApplicationServiceServer).PutFile(&applicationServicePutFileServer{stream})
}

type ApplicationService_PutFileServer interface {
	SendAndClose(*PutFileResp) error
	Recv() (*PutFileReq, error)
	grpc.ServerStream
}

type applicationServicePutFileServer struct {
	grpc.ServerStream
}

func (x *applicationServicePutFileServer) SendAndClose(m *PutFileResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *applicationServicePutFileServer) Recv() (*PutFileReq, error) {
	m := new(PutFileReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ApplicationService_GetFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetFileReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApplicationServiceServer).GetFile(m, &applicationServiceGetFileServer{stream})
}

type ApplicationService_GetFileServer interface {
	Send(*GetFileResp) error
	grpc.ServerStream
}

type applicationServiceGetFileServer struct {
	grpc.ServerStream
}

func (x *applicationServiceGetFileServer) Send(m *GetFileResp) error {
	return x.ServerStream.SendMsg(m)
}

// ApplicationService_ServiceDesc is the grpc.ServiceDesc for ApplicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApplicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Application.ApplicationService",
	HandlerType: (*ApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ApplicationService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ApplicationService_List_Handler,
		},
		{
			MethodName: "ListAll",
			Handler:    _ApplicationService_ListAll_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PutFile",
			Handler:       _ApplicationService_PutFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetFile",
			Handler:       _ApplicationService_GetFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/application.proto",
}
