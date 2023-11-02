// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: subscriber/subscriber.proto

package subscriber

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

// PushClient is the client API for Push service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushClient interface {
	// Sends a greeting
	PushMessage(ctx context.Context, in *PushMessageRequest, opts ...grpc.CallOption) (*PushMessageReply, error)
}

type pushClient struct {
	cc grpc.ClientConnInterface
}

func NewPushClient(cc grpc.ClientConnInterface) PushClient {
	return &pushClient{cc}
}

func (c *pushClient) PushMessage(ctx context.Context, in *PushMessageRequest, opts ...grpc.CallOption) (*PushMessageReply, error) {
	out := new(PushMessageReply)
	err := c.cc.Invoke(ctx, "/Push/PushMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushServer is the server API for Push service.
// All implementations must embed UnimplementedPushServer
// for forward compatibility
type PushServer interface {
	// Sends a greeting
	PushMessage(context.Context, *PushMessageRequest) (*PushMessageReply, error)
	mustEmbedUnimplementedPushServer()
}

// UnimplementedPushServer must be embedded to have forward compatible implementations.
type UnimplementedPushServer struct {
}

func (UnimplementedPushServer) PushMessage(context.Context, *PushMessageRequest) (*PushMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMessage not implemented")
}
func (UnimplementedPushServer) mustEmbedUnimplementedPushServer() {}

// UnsafePushServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushServer will
// result in compilation errors.
type UnsafePushServer interface {
	mustEmbedUnimplementedPushServer()
}

func RegisterPushServer(s grpc.ServiceRegistrar, srv PushServer) {
	s.RegisterService(&Push_ServiceDesc, srv)
}

func _Push_PushMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).PushMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Push/PushMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).PushMessage(ctx, req.(*PushMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Push_ServiceDesc is the grpc.ServiceDesc for Push service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Push_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Push",
	HandlerType: (*PushServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushMessage",
			Handler:    _Push_PushMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscriber/subscriber.proto",
}
