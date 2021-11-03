// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rtc

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

// RTCClient is the client API for RTC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RTCClient interface {
	Signal(ctx context.Context, opts ...grpc.CallOption) (RTC_SignalClient, error)
}

type rTCClient struct {
	cc grpc.ClientConnInterface
}

func NewRTCClient(cc grpc.ClientConnInterface) RTCClient {
	return &rTCClient{cc}
}

func (c *rTCClient) Signal(ctx context.Context, opts ...grpc.CallOption) (RTC_SignalClient, error) {
	stream, err := c.cc.NewStream(ctx, &RTC_ServiceDesc.Streams[0], "/rtc.RTC/Signal", opts...)
	if err != nil {
		return nil, err
	}
	x := &rTCSignalClient{stream}
	return x, nil
}

type RTC_SignalClient interface {
	Send(*Request) error
	Recv() (*Reply, error)
	grpc.ClientStream
}

type rTCSignalClient struct {
	grpc.ClientStream
}

func (x *rTCSignalClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rTCSignalClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RTCServer is the server API for RTC service.
// All implementations must embed UnimplementedRTCServer
// for forward compatibility
type RTCServer interface {
	Signal(RTC_SignalServer) error
	mustEmbedUnimplementedRTCServer()
}

// UnimplementedRTCServer must be embedded to have forward compatible implementations.
type UnimplementedRTCServer struct {
}

func (UnimplementedRTCServer) Signal(RTC_SignalServer) error {
	return status.Errorf(codes.Unimplemented, "method Signal not implemented")
}
func (UnimplementedRTCServer) mustEmbedUnimplementedRTCServer() {}

// UnsafeRTCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RTCServer will
// result in compilation errors.
type UnsafeRTCServer interface {
	mustEmbedUnimplementedRTCServer()
}

func RegisterRTCServer(s grpc.ServiceRegistrar, srv RTCServer) {
	s.RegisterService(&RTC_ServiceDesc, srv)
}

func _RTC_Signal_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RTCServer).Signal(&rTCSignalServer{stream})
}

type RTC_SignalServer interface {
	Send(*Reply) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type rTCSignalServer struct {
	grpc.ServerStream
}

func (x *rTCSignalServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rTCSignalServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RTC_ServiceDesc is the grpc.ServiceDesc for RTC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RTC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rtc.RTC",
	HandlerType: (*RTCServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Signal",
			Handler:       _RTC_Signal_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/rtc/rtc.proto",
}
