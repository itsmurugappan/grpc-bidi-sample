// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: pp.proto

package pp

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PingPongClient is the client API for PingPong service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingPongClient interface {
	// PP
	PingPong(ctx context.Context, opts ...grpc.CallOption) (PingPong_PingPongClient, error)
}

type pingPongClient struct {
	cc grpc.ClientConnInterface
}

func NewPingPongClient(cc grpc.ClientConnInterface) PingPongClient {
	return &pingPongClient{cc}
}

func (c *pingPongClient) PingPong(ctx context.Context, opts ...grpc.CallOption) (PingPong_PingPongClient, error) {
	stream, err := c.cc.NewStream(ctx, &PingPong_ServiceDesc.Streams[0], "/pp.PingPong/PingPong", opts...)
	if err != nil {
		return nil, err
	}
	x := &pingPongPingPongClient{stream}
	return x, nil
}

type PingPong_PingPongClient interface {
	Send(*anypb.Any) error
	Recv() (*anypb.Any, error)
	grpc.ClientStream
}

type pingPongPingPongClient struct {
	grpc.ClientStream
}

func (x *pingPongPingPongClient) Send(m *anypb.Any) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pingPongPingPongClient) Recv() (*anypb.Any, error) {
	m := new(anypb.Any)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PingPongServer is the server API for PingPong service.
// All implementations must embed UnimplementedPingPongServer
// for forward compatibility
type PingPongServer interface {
	// PP
	PingPong(PingPong_PingPongServer) error
	mustEmbedUnimplementedPingPongServer()
}

// UnimplementedPingPongServer must be embedded to have forward compatible implementations.
type UnimplementedPingPongServer struct {
}

func (UnimplementedPingPongServer) PingPong(PingPong_PingPongServer) error {
	return status.Errorf(codes.Unimplemented, "method PingPong not implemented")
}
func (UnimplementedPingPongServer) mustEmbedUnimplementedPingPongServer() {}

// UnsafePingPongServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PingPongServer will
// result in compilation errors.
type UnsafePingPongServer interface {
	mustEmbedUnimplementedPingPongServer()
}

func RegisterPingPongServer(s grpc.ServiceRegistrar, srv PingPongServer) {
	s.RegisterService(&PingPong_ServiceDesc, srv)
}

func _PingPong_PingPong_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PingPongServer).PingPong(&pingPongPingPongServer{stream})
}

type PingPong_PingPongServer interface {
	Send(*anypb.Any) error
	Recv() (*anypb.Any, error)
	grpc.ServerStream
}

type pingPongPingPongServer struct {
	grpc.ServerStream
}

func (x *pingPongPingPongServer) Send(m *anypb.Any) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pingPongPingPongServer) Recv() (*anypb.Any, error) {
	m := new(anypb.Any)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PingPong_ServiceDesc is the grpc.ServiceDesc for PingPong service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PingPong_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pp.PingPong",
	HandlerType: (*PingPongServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingPong",
			Handler:       _PingPong_PingPong_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pp.proto",
}
