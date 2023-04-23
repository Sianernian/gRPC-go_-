// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: server_stream.proto

package service

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

const (
	ServerStreamTalk_ListValue_FullMethodName = "/ServerStreamTalk/listValue"
)

// ServerStreamTalkClient is the client API for ServerStreamTalk service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerStreamTalkClient interface {
	ListValue(ctx context.Context, in *ServerStreamRequest, opts ...grpc.CallOption) (ServerStreamTalk_ListValueClient, error)
}

type serverStreamTalkClient struct {
	cc grpc.ClientConnInterface
}

func NewServerStreamTalkClient(cc grpc.ClientConnInterface) ServerStreamTalkClient {
	return &serverStreamTalkClient{cc}
}

func (c *serverStreamTalkClient) ListValue(ctx context.Context, in *ServerStreamRequest, opts ...grpc.CallOption) (ServerStreamTalk_ListValueClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServerStreamTalk_ServiceDesc.Streams[0], ServerStreamTalk_ListValue_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serverStreamTalkListValueClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServerStreamTalk_ListValueClient interface {
	Recv() (*ServerStreamResponse, error)
	grpc.ClientStream
}

type serverStreamTalkListValueClient struct {
	grpc.ClientStream
}

func (x *serverStreamTalkListValueClient) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerStreamTalkServer is the server API for ServerStreamTalk service.
// All implementations must embed UnimplementedServerStreamTalkServer
// for forward compatibility
type ServerStreamTalkServer interface {
	ListValue(*ServerStreamRequest, ServerStreamTalk_ListValueServer) error
	mustEmbedUnimplementedServerStreamTalkServer()
}

// UnimplementedServerStreamTalkServer must be embedded to have forward compatible implementations.
type UnimplementedServerStreamTalkServer struct {
}

func (UnimplementedServerStreamTalkServer) ListValue(*ServerStreamRequest, ServerStreamTalk_ListValueServer) error {
	return status.Errorf(codes.Unimplemented, "method ListValue not implemented")
}
func (UnimplementedServerStreamTalkServer) mustEmbedUnimplementedServerStreamTalkServer() {}

// UnsafeServerStreamTalkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerStreamTalkServer will
// result in compilation errors.
type UnsafeServerStreamTalkServer interface {
	mustEmbedUnimplementedServerStreamTalkServer()
}

func RegisterServerStreamTalkServer(s grpc.ServiceRegistrar, srv ServerStreamTalkServer) {
	s.RegisterService(&ServerStreamTalk_ServiceDesc, srv)
}

func _ServerStreamTalk_ListValue_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServerStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServerStreamTalkServer).ListValue(m, &serverStreamTalkListValueServer{stream})
}

type ServerStreamTalk_ListValueServer interface {
	Send(*ServerStreamResponse) error
	grpc.ServerStream
}

type serverStreamTalkListValueServer struct {
	grpc.ServerStream
}

func (x *serverStreamTalkListValueServer) Send(m *ServerStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ServerStreamTalk_ServiceDesc is the grpc.ServiceDesc for ServerStreamTalk service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerStreamTalk_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ServerStreamTalk",
	HandlerType: (*ServerStreamTalkServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "listValue",
			Handler:       _ServerStreamTalk_ListValue_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server_stream.proto",
}