// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: data.proto

package pb

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

// ViewStorageServiceClient is the client API for ViewStorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ViewStorageServiceClient interface {
	ReadFile(ctx context.Context, in *RequestReadFile, opts ...grpc.CallOption) (*File, error)
}

type viewStorageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewViewStorageServiceClient(cc grpc.ClientConnInterface) ViewStorageServiceClient {
	return &viewStorageServiceClient{cc}
}

func (c *viewStorageServiceClient) ReadFile(ctx context.Context, in *RequestReadFile, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := c.cc.Invoke(ctx, "/cel_sandbox.ViewStorageService/ReadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ViewStorageServiceServer is the server API for ViewStorageService service.
// All implementations must embed UnimplementedViewStorageServiceServer
// for forward compatibility
type ViewStorageServiceServer interface {
	ReadFile(context.Context, *RequestReadFile) (*File, error)
	mustEmbedUnimplementedViewStorageServiceServer()
}

// UnimplementedViewStorageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedViewStorageServiceServer struct {
}

func (UnimplementedViewStorageServiceServer) ReadFile(context.Context, *RequestReadFile) (*File, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadFile not implemented")
}
func (UnimplementedViewStorageServiceServer) mustEmbedUnimplementedViewStorageServiceServer() {}

// UnsafeViewStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ViewStorageServiceServer will
// result in compilation errors.
type UnsafeViewStorageServiceServer interface {
	mustEmbedUnimplementedViewStorageServiceServer()
}

func RegisterViewStorageServiceServer(s grpc.ServiceRegistrar, srv ViewStorageServiceServer) {
	s.RegisterService(&ViewStorageService_ServiceDesc, srv)
}

func _ViewStorageService_ReadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestReadFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewStorageServiceServer).ReadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cel_sandbox.ViewStorageService/ReadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewStorageServiceServer).ReadFile(ctx, req.(*RequestReadFile))
	}
	return interceptor(ctx, in, info, handler)
}

// ViewStorageService_ServiceDesc is the grpc.ServiceDesc for ViewStorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ViewStorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cel_sandbox.ViewStorageService",
	HandlerType: (*ViewStorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadFile",
			Handler:    _ViewStorageService_ReadFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}