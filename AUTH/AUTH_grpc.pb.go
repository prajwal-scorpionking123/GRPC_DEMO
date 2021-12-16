// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package AUTH

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

// AUTHENTICATIONClient is the client API for AUTHENTICATION service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AUTHENTICATIONClient interface {
	Authenticate(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
}

type aUTHENTICATIONClient struct {
	cc grpc.ClientConnInterface
}

func NewAUTHENTICATIONClient(cc grpc.ClientConnInterface) AUTHENTICATIONClient {
	return &aUTHENTICATIONClient{cc}
}

func (c *aUTHENTICATIONClient) Authenticate(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/AUTH.AUTHENTICATION/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aUTHENTICATIONClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/AUTH.AUTHENTICATION/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AUTHENTICATIONServer is the server API for AUTHENTICATION service.
// All implementations must embed UnimplementedAUTHENTICATIONServer
// for forward compatibility
type AUTHENTICATIONServer interface {
	Authenticate(context.Context, *LoginRequest) (*LoginReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	mustEmbedUnimplementedAUTHENTICATIONServer()
}

// UnimplementedAUTHENTICATIONServer must be embedded to have forward compatible implementations.
type UnimplementedAUTHENTICATIONServer struct {
}

func (UnimplementedAUTHENTICATIONServer) Authenticate(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedAUTHENTICATIONServer) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAUTHENTICATIONServer) mustEmbedUnimplementedAUTHENTICATIONServer() {}

// UnsafeAUTHENTICATIONServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AUTHENTICATIONServer will
// result in compilation errors.
type UnsafeAUTHENTICATIONServer interface {
	mustEmbedUnimplementedAUTHENTICATIONServer()
}

func RegisterAUTHENTICATIONServer(s grpc.ServiceRegistrar, srv AUTHENTICATIONServer) {
	s.RegisterService(&AUTHENTICATION_ServiceDesc, srv)
}

func _AUTHENTICATION_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AUTHENTICATIONServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AUTH.AUTHENTICATION/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AUTHENTICATIONServer).Authenticate(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AUTHENTICATION_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AUTHENTICATIONServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AUTH.AUTHENTICATION/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AUTHENTICATIONServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AUTHENTICATION_ServiceDesc is the grpc.ServiceDesc for AUTHENTICATION service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AUTHENTICATION_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AUTH.AUTHENTICATION",
	HandlerType: (*AUTHENTICATIONServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _AUTHENTICATION_Authenticate_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _AUTHENTICATION_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AUTH.proto",
}

// GuploadServiceClient is the client API for GuploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GuploadServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (GuploadService_UploadClient, error)
}

type guploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGuploadServiceClient(cc grpc.ClientConnInterface) GuploadServiceClient {
	return &guploadServiceClient{cc}
}

func (c *guploadServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (GuploadService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &GuploadService_ServiceDesc.Streams[0], "/AUTH.GuploadService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &guploadServiceUploadClient{stream}
	return x, nil
}

type GuploadService_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type guploadServiceUploadClient struct {
	grpc.ClientStream
}

func (x *guploadServiceUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *guploadServiceUploadClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GuploadServiceServer is the server API for GuploadService service.
// All implementations must embed UnimplementedGuploadServiceServer
// for forward compatibility
type GuploadServiceServer interface {
	Upload(GuploadService_UploadServer) error
	mustEmbedUnimplementedGuploadServiceServer()
}

// UnimplementedGuploadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGuploadServiceServer struct {
}

func (UnimplementedGuploadServiceServer) Upload(GuploadService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedGuploadServiceServer) mustEmbedUnimplementedGuploadServiceServer() {}

// UnsafeGuploadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GuploadServiceServer will
// result in compilation errors.
type UnsafeGuploadServiceServer interface {
	mustEmbedUnimplementedGuploadServiceServer()
}

func RegisterGuploadServiceServer(s grpc.ServiceRegistrar, srv GuploadServiceServer) {
	s.RegisterService(&GuploadService_ServiceDesc, srv)
}

func _GuploadService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GuploadServiceServer).Upload(&guploadServiceUploadServer{stream})
}

type GuploadService_UploadServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type guploadServiceUploadServer struct {
	grpc.ServerStream
}

func (x *guploadServiceUploadServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *guploadServiceUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GuploadService_ServiceDesc is the grpc.ServiceDesc for GuploadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GuploadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AUTH.GuploadService",
	HandlerType: (*GuploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _GuploadService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "AUTH.proto",
}
