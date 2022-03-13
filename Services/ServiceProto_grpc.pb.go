// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: ServiceProto.proto

package Services

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

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	// 增加书籍库存
	AddBookNum(ctx context.Context, in *AddBookNumRequest, opts ...grpc.CallOption) (*AddBookNumResponse, error)
	// 减少书籍库存
	DecBookNum(ctx context.Context, in *DecBookRequest, opts ...grpc.CallOption) (*DecBookResponse, error)
	// 上架新书
	AddNewBook(ctx context.Context, in *AddNewBookRequest, opts ...grpc.CallOption) (*AddNewBookResponse, error)
	// 下架书籍
	DelBook(ctx context.Context, in *DelBookRequest, opts ...grpc.CallOption) (*DelBookResponse, error)
	// 删除书籍
	RemoveBook(ctx context.Context, in *RemoveBookRequest, opts ...grpc.CallOption) (*RemoveBookResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) AddBookNum(ctx context.Context, in *AddBookNumRequest, opts ...grpc.CallOption) (*AddBookNumResponse, error) {
	out := new(AddBookNumResponse)
	err := c.cc.Invoke(ctx, "/Services.AdminService/AddBookNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DecBookNum(ctx context.Context, in *DecBookRequest, opts ...grpc.CallOption) (*DecBookResponse, error) {
	out := new(DecBookResponse)
	err := c.cc.Invoke(ctx, "/Services.AdminService/DecBookNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AddNewBook(ctx context.Context, in *AddNewBookRequest, opts ...grpc.CallOption) (*AddNewBookResponse, error) {
	out := new(AddNewBookResponse)
	err := c.cc.Invoke(ctx, "/Services.AdminService/AddNewBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DelBook(ctx context.Context, in *DelBookRequest, opts ...grpc.CallOption) (*DelBookResponse, error) {
	out := new(DelBookResponse)
	err := c.cc.Invoke(ctx, "/Services.AdminService/DelBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) RemoveBook(ctx context.Context, in *RemoveBookRequest, opts ...grpc.CallOption) (*RemoveBookResponse, error) {
	out := new(RemoveBookResponse)
	err := c.cc.Invoke(ctx, "/Services.AdminService/RemoveBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	// 增加书籍库存
	AddBookNum(context.Context, *AddBookNumRequest) (*AddBookNumResponse, error)
	// 减少书籍库存
	DecBookNum(context.Context, *DecBookRequest) (*DecBookResponse, error)
	// 上架新书
	AddNewBook(context.Context, *AddNewBookRequest) (*AddNewBookResponse, error)
	// 下架书籍
	DelBook(context.Context, *DelBookRequest) (*DelBookResponse, error)
	// 删除书籍
	RemoveBook(context.Context, *RemoveBookRequest) (*RemoveBookResponse, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) AddBookNum(context.Context, *AddBookNumRequest) (*AddBookNumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBookNum not implemented")
}
func (UnimplementedAdminServiceServer) DecBookNum(context.Context, *DecBookRequest) (*DecBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecBookNum not implemented")
}
func (UnimplementedAdminServiceServer) AddNewBook(context.Context, *AddNewBookRequest) (*AddNewBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewBook not implemented")
}
func (UnimplementedAdminServiceServer) DelBook(context.Context, *DelBookRequest) (*DelBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelBook not implemented")
}
func (UnimplementedAdminServiceServer) RemoveBook(context.Context, *RemoveBookRequest) (*RemoveBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBook not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_AddBookNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBookNumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AddBookNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Services.AdminService/AddBookNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AddBookNum(ctx, req.(*AddBookNumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DecBookNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DecBookNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Services.AdminService/DecBookNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DecBookNum(ctx, req.(*DecBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AddNewBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNewBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AddNewBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Services.AdminService/AddNewBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AddNewBook(ctx, req.(*AddNewBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DelBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DelBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Services.AdminService/DelBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DelBook(ctx, req.(*DelBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_RemoveBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).RemoveBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Services.AdminService/RemoveBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).RemoveBook(ctx, req.(*RemoveBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Services.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBookNum",
			Handler:    _AdminService_AddBookNum_Handler,
		},
		{
			MethodName: "DecBookNum",
			Handler:    _AdminService_DecBookNum_Handler,
		},
		{
			MethodName: "AddNewBook",
			Handler:    _AdminService_AddNewBook_Handler,
		},
		{
			MethodName: "DelBook",
			Handler:    _AdminService_DelBook_Handler,
		},
		{
			MethodName: "RemoveBook",
			Handler:    _AdminService_RemoveBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ServiceProto.proto",
}

// WorkerServiceClient is the client API for WorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerServiceClient interface {
	// 查询所有书籍数量
	GetAllBook(ctx context.Context, opts ...grpc.CallOption) (WorkerService_GetAllBookClient, error)
	// 查询指定书籍数量
	GetBookNum(ctx context.Context, in *GetBookNumRequest, opts ...grpc.CallOption) (*GetBookNumResponse, error)
}

type workerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerServiceClient(cc grpc.ClientConnInterface) WorkerServiceClient {
	return &workerServiceClient{cc}
}

func (c *workerServiceClient) GetAllBook(ctx context.Context, opts ...grpc.CallOption) (WorkerService_GetAllBookClient, error) {
	stream, err := c.cc.NewStream(ctx, &WorkerService_ServiceDesc.Streams[0], "/Services.WorkerService/GetAllBook", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerServiceGetAllBookClient{stream}
	return x, nil
}

type WorkerService_GetAllBookClient interface {
	Send(*GetAllBookRequest) error
	Recv() (*GetAllBookResponse, error)
	grpc.ClientStream
}

type workerServiceGetAllBookClient struct {
	grpc.ClientStream
}

func (x *workerServiceGetAllBookClient) Send(m *GetAllBookRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workerServiceGetAllBookClient) Recv() (*GetAllBookResponse, error) {
	m := new(GetAllBookResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workerServiceClient) GetBookNum(ctx context.Context, in *GetBookNumRequest, opts ...grpc.CallOption) (*GetBookNumResponse, error) {
	out := new(GetBookNumResponse)
	err := c.cc.Invoke(ctx, "/Services.WorkerService/GetBookNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServiceServer is the server API for WorkerService service.
// All implementations must embed UnimplementedWorkerServiceServer
// for forward compatibility
type WorkerServiceServer interface {
	// 查询所有书籍数量
	GetAllBook(WorkerService_GetAllBookServer) error
	// 查询指定书籍数量
	GetBookNum(context.Context, *GetBookNumRequest) (*GetBookNumResponse, error)
	mustEmbedUnimplementedWorkerServiceServer()
}

// UnimplementedWorkerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkerServiceServer struct {
}

func (UnimplementedWorkerServiceServer) GetAllBook(WorkerService_GetAllBookServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllBook not implemented")
}
func (UnimplementedWorkerServiceServer) GetBookNum(context.Context, *GetBookNumRequest) (*GetBookNumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookNum not implemented")
}
func (UnimplementedWorkerServiceServer) mustEmbedUnimplementedWorkerServiceServer() {}

// UnsafeWorkerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServiceServer will
// result in compilation errors.
type UnsafeWorkerServiceServer interface {
	mustEmbedUnimplementedWorkerServiceServer()
}

func RegisterWorkerServiceServer(s grpc.ServiceRegistrar, srv WorkerServiceServer) {
	s.RegisterService(&WorkerService_ServiceDesc, srv)
}

func _WorkerService_GetAllBook_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkerServiceServer).GetAllBook(&workerServiceGetAllBookServer{stream})
}

type WorkerService_GetAllBookServer interface {
	Send(*GetAllBookResponse) error
	Recv() (*GetAllBookRequest, error)
	grpc.ServerStream
}

type workerServiceGetAllBookServer struct {
	grpc.ServerStream
}

func (x *workerServiceGetAllBookServer) Send(m *GetAllBookResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workerServiceGetAllBookServer) Recv() (*GetAllBookRequest, error) {
	m := new(GetAllBookRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _WorkerService_GetBookNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookNumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).GetBookNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Services.WorkerService/GetBookNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).GetBookNum(ctx, req.(*GetBookNumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkerService_ServiceDesc is the grpc.ServiceDesc for WorkerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Services.WorkerService",
	HandlerType: (*WorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBookNum",
			Handler:    _WorkerService_GetBookNum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllBook",
			Handler:       _WorkerService_GetAllBook_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ServiceProto.proto",
}