// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package subscribe

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SubscribeServiceClient is the client API for SubscribeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscribeServiceClient interface {
	OnetimePaymentRPC(ctx context.Context, in *OnetimePaymentRequest, opts ...grpc.CallOption) (*OnetimePaymentResponse, error)
	AgainPaymentRPC(ctx context.Context, in *AgainPaymentRequest, opts ...grpc.CallOption) (*AgainPaymentResponse, error)
	SchedulePaymentRPC(ctx context.Context, in *SchedulePayemntRequest, opts ...grpc.CallOption) (*SchedulePaymentResponse, error)
	UnschedulePaymentRPC(ctx context.Context, in *UnschedulePaymentRequest, opts ...grpc.CallOption) (*UnschedulePaymentResponse, error)
	GetScheduledPaymentRPC(ctx context.Context, in *GetPaymentScheduleRequest, opts ...grpc.CallOption) (*GetPaymentScheduleResponse, error)
	GetScheduledPaymentByCustomerUidRPC(ctx context.Context, in *GetPaymentScheduleByCustomerRequest, opts ...grpc.CallOption) (*GetPaymentScheduleByCustomerResponse, error)
}

type subscribeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscribeServiceClient(cc grpc.ClientConnInterface) SubscribeServiceClient {
	return &subscribeServiceClient{cc}
}

func (c *subscribeServiceClient) OnetimePaymentRPC(ctx context.Context, in *OnetimePaymentRequest, opts ...grpc.CallOption) (*OnetimePaymentResponse, error) {
	out := new(OnetimePaymentResponse)
	err := c.cc.Invoke(ctx, "/subscribe.SubscribeService/OnetimePaymentRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeServiceClient) AgainPaymentRPC(ctx context.Context, in *AgainPaymentRequest, opts ...grpc.CallOption) (*AgainPaymentResponse, error) {
	out := new(AgainPaymentResponse)
	err := c.cc.Invoke(ctx, "/subscribe.SubscribeService/AgainPaymentRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeServiceClient) SchedulePaymentRPC(ctx context.Context, in *SchedulePayemntRequest, opts ...grpc.CallOption) (*SchedulePaymentResponse, error) {
	out := new(SchedulePaymentResponse)
	err := c.cc.Invoke(ctx, "/subscribe.SubscribeService/SchedulePaymentRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeServiceClient) UnschedulePaymentRPC(ctx context.Context, in *UnschedulePaymentRequest, opts ...grpc.CallOption) (*UnschedulePaymentResponse, error) {
	out := new(UnschedulePaymentResponse)
	err := c.cc.Invoke(ctx, "/subscribe.SubscribeService/UnschedulePaymentRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeServiceClient) GetScheduledPaymentRPC(ctx context.Context, in *GetPaymentScheduleRequest, opts ...grpc.CallOption) (*GetPaymentScheduleResponse, error) {
	out := new(GetPaymentScheduleResponse)
	err := c.cc.Invoke(ctx, "/subscribe.SubscribeService/GetScheduledPaymentRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeServiceClient) GetScheduledPaymentByCustomerUidRPC(ctx context.Context, in *GetPaymentScheduleByCustomerRequest, opts ...grpc.CallOption) (*GetPaymentScheduleByCustomerResponse, error) {
	out := new(GetPaymentScheduleByCustomerResponse)
	err := c.cc.Invoke(ctx, "/subscribe.SubscribeService/GetScheduledPaymentByCustomerUidRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscribeServiceServer is the server API for SubscribeService service.
// All implementations must embed UnimplementedSubscribeServiceServer
// for forward compatibility
type SubscribeServiceServer interface {
	OnetimePaymentRPC(context.Context, *OnetimePaymentRequest) (*OnetimePaymentResponse, error)
	AgainPaymentRPC(context.Context, *AgainPaymentRequest) (*AgainPaymentResponse, error)
	SchedulePaymentRPC(context.Context, *SchedulePayemntRequest) (*SchedulePaymentResponse, error)
	UnschedulePaymentRPC(context.Context, *UnschedulePaymentRequest) (*UnschedulePaymentResponse, error)
	GetScheduledPaymentRPC(context.Context, *GetPaymentScheduleRequest) (*GetPaymentScheduleResponse, error)
	GetScheduledPaymentByCustomerUidRPC(context.Context, *GetPaymentScheduleByCustomerRequest) (*GetPaymentScheduleByCustomerResponse, error)
	mustEmbedUnimplementedSubscribeServiceServer()
}

// UnimplementedSubscribeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubscribeServiceServer struct {
}

func (UnimplementedSubscribeServiceServer) OnetimePaymentRPC(context.Context, *OnetimePaymentRequest) (*OnetimePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnetimePaymentRPC not implemented")
}
func (UnimplementedSubscribeServiceServer) AgainPaymentRPC(context.Context, *AgainPaymentRequest) (*AgainPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AgainPaymentRPC not implemented")
}
func (UnimplementedSubscribeServiceServer) SchedulePaymentRPC(context.Context, *SchedulePayemntRequest) (*SchedulePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchedulePaymentRPC not implemented")
}
func (UnimplementedSubscribeServiceServer) UnschedulePaymentRPC(context.Context, *UnschedulePaymentRequest) (*UnschedulePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnschedulePaymentRPC not implemented")
}
func (UnimplementedSubscribeServiceServer) GetScheduledPaymentRPC(context.Context, *GetPaymentScheduleRequest) (*GetPaymentScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScheduledPaymentRPC not implemented")
}
func (UnimplementedSubscribeServiceServer) GetScheduledPaymentByCustomerUidRPC(context.Context, *GetPaymentScheduleByCustomerRequest) (*GetPaymentScheduleByCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScheduledPaymentByCustomerUidRPC not implemented")
}
func (UnimplementedSubscribeServiceServer) mustEmbedUnimplementedSubscribeServiceServer() {}

// UnsafeSubscribeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscribeServiceServer will
// result in compilation errors.
type UnsafeSubscribeServiceServer interface {
	mustEmbedUnimplementedSubscribeServiceServer()
}

func RegisterSubscribeServiceServer(s *grpc.Server, srv SubscribeServiceServer) {
	s.RegisterService(&_SubscribeService_serviceDesc, srv)
}

func _SubscribeService_OnetimePaymentRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnetimePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServiceServer).OnetimePaymentRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscribe.SubscribeService/OnetimePaymentRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServiceServer).OnetimePaymentRPC(ctx, req.(*OnetimePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscribeService_AgainPaymentRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgainPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServiceServer).AgainPaymentRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscribe.SubscribeService/AgainPaymentRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServiceServer).AgainPaymentRPC(ctx, req.(*AgainPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscribeService_SchedulePaymentRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulePayemntRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServiceServer).SchedulePaymentRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscribe.SubscribeService/SchedulePaymentRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServiceServer).SchedulePaymentRPC(ctx, req.(*SchedulePayemntRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscribeService_UnschedulePaymentRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnschedulePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServiceServer).UnschedulePaymentRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscribe.SubscribeService/UnschedulePaymentRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServiceServer).UnschedulePaymentRPC(ctx, req.(*UnschedulePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscribeService_GetScheduledPaymentRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServiceServer).GetScheduledPaymentRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscribe.SubscribeService/GetScheduledPaymentRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServiceServer).GetScheduledPaymentRPC(ctx, req.(*GetPaymentScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscribeService_GetScheduledPaymentByCustomerUidRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentScheduleByCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServiceServer).GetScheduledPaymentByCustomerUidRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscribe.SubscribeService/GetScheduledPaymentByCustomerUidRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServiceServer).GetScheduledPaymentByCustomerUidRPC(ctx, req.(*GetPaymentScheduleByCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SubscribeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "subscribe.SubscribeService",
	HandlerType: (*SubscribeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnetimePaymentRPC",
			Handler:    _SubscribeService_OnetimePaymentRPC_Handler,
		},
		{
			MethodName: "AgainPaymentRPC",
			Handler:    _SubscribeService_AgainPaymentRPC_Handler,
		},
		{
			MethodName: "SchedulePaymentRPC",
			Handler:    _SubscribeService_SchedulePaymentRPC_Handler,
		},
		{
			MethodName: "UnschedulePaymentRPC",
			Handler:    _SubscribeService_UnschedulePaymentRPC_Handler,
		},
		{
			MethodName: "GetScheduledPaymentRPC",
			Handler:    _SubscribeService_GetScheduledPaymentRPC_Handler,
		},
		{
			MethodName: "GetScheduledPaymentByCustomerUidRPC",
			Handler:    _SubscribeService_GetScheduledPaymentByCustomerUidRPC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/subscribe/subscribe.proto",
}
