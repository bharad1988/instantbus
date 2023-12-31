// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: busproto/bus.proto

package busproto

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

// BusClient is the client API for Bus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BusClient interface {
	SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error)
	SubTopic(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeReply, error)
}

type busClient struct {
	cc grpc.ClientConnInterface
}

func NewBusClient(cc grpc.ClientConnInterface) BusClient {
	return &busClient{cc}
}

func (c *busClient) SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error) {
	out := new(MessageReply)
	err := c.cc.Invoke(ctx, "/Bus/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busClient) SubTopic(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeReply, error) {
	out := new(SubscribeReply)
	err := c.cc.Invoke(ctx, "/Bus/SubTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BusServer is the server API for Bus service.
// All implementations must embed UnimplementedBusServer
// for forward compatibility
type BusServer interface {
	SendMessage(context.Context, *MessageRequest) (*MessageReply, error)
	SubTopic(context.Context, *SubscribeRequest) (*SubscribeReply, error)
	mustEmbedUnimplementedBusServer()
}

// UnimplementedBusServer must be embedded to have forward compatible implementations.
type UnimplementedBusServer struct {
}

func (UnimplementedBusServer) SendMessage(context.Context, *MessageRequest) (*MessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedBusServer) SubTopic(context.Context, *SubscribeRequest) (*SubscribeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubTopic not implemented")
}
func (UnimplementedBusServer) mustEmbedUnimplementedBusServer() {}

// UnsafeBusServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BusServer will
// result in compilation errors.
type UnsafeBusServer interface {
	mustEmbedUnimplementedBusServer()
}

func RegisterBusServer(s grpc.ServiceRegistrar, srv BusServer) {
	s.RegisterService(&Bus_ServiceDesc, srv)
}

func _Bus_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bus/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusServer).SendMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bus_SubTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusServer).SubTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bus/SubTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusServer).SubTopic(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bus_ServiceDesc is the grpc.ServiceDesc for Bus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Bus",
	HandlerType: (*BusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _Bus_SendMessage_Handler,
		},
		{
			MethodName: "SubTopic",
			Handler:    _Bus_SubTopic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "busproto/bus.proto",
}
