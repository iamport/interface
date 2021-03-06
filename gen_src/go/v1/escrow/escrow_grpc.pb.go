// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package escrow

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// EscrowServiceClient is the client API for EscrowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EscrowServiceClient interface {
	EscrowPostRPC(ctx context.Context, in *EscrowRequest, opts ...grpc.CallOption) (*EscrowResponse, error)
	EscrowPutRPC(ctx context.Context, in *EscrowRequest, opts ...grpc.CallOption) (*EscrowResponse, error)
}

type escrowServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEscrowServiceClient(cc grpc.ClientConnInterface) EscrowServiceClient {
	return &escrowServiceClient{cc}
}

func (c *escrowServiceClient) EscrowPostRPC(ctx context.Context, in *EscrowRequest, opts ...grpc.CallOption) (*EscrowResponse, error) {
	out := new(EscrowResponse)
	err := c.cc.Invoke(ctx, "/escrow.EscrowService/EscrowPostRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *escrowServiceClient) EscrowPutRPC(ctx context.Context, in *EscrowRequest, opts ...grpc.CallOption) (*EscrowResponse, error) {
	out := new(EscrowResponse)
	err := c.cc.Invoke(ctx, "/escrow.EscrowService/EscrowPutRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EscrowServiceServer is the server API for EscrowService service.
// All implementations must embed UnimplementedEscrowServiceServer
// for forward compatibility
type EscrowServiceServer interface {
	EscrowPostRPC(context.Context, *EscrowRequest) (*EscrowResponse, error)
	EscrowPutRPC(context.Context, *EscrowRequest) (*EscrowResponse, error)
	mustEmbedUnimplementedEscrowServiceServer()
}

// UnimplementedEscrowServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEscrowServiceServer struct {
}

func (UnimplementedEscrowServiceServer) EscrowPostRPC(context.Context, *EscrowRequest) (*EscrowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EscrowPostRPC not implemented")
}
func (UnimplementedEscrowServiceServer) EscrowPutRPC(context.Context, *EscrowRequest) (*EscrowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EscrowPutRPC not implemented")
}
func (UnimplementedEscrowServiceServer) mustEmbedUnimplementedEscrowServiceServer() {}

// UnsafeEscrowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EscrowServiceServer will
// result in compilation errors.
type UnsafeEscrowServiceServer interface {
	mustEmbedUnimplementedEscrowServiceServer()
}

func RegisterEscrowServiceServer(s grpc.ServiceRegistrar, srv EscrowServiceServer) {
	s.RegisterService(&_EscrowService_serviceDesc, srv)
}

func _EscrowService_EscrowPostRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EscrowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EscrowServiceServer).EscrowPostRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/escrow.EscrowService/EscrowPostRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EscrowServiceServer).EscrowPostRPC(ctx, req.(*EscrowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EscrowService_EscrowPutRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EscrowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EscrowServiceServer).EscrowPutRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/escrow.EscrowService/EscrowPutRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EscrowServiceServer).EscrowPutRPC(ctx, req.(*EscrowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EscrowService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "escrow.EscrowService",
	HandlerType: (*EscrowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EscrowPostRPC",
			Handler:    _EscrowService_EscrowPostRPC_Handler,
		},
		{
			MethodName: "EscrowPutRPC",
			Handler:    _EscrowService_EscrowPutRPC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/escrow/escrow.proto",
}
