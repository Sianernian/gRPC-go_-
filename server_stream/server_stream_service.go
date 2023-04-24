package main

import (
	"context"
	pb "gRPC_protoc/server_stream/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type ServerStreamService struct {
	pb.UnimplementedServerStreamTalkServer
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("net listen err:%v", err)
	}
	gRPCServer := grpc.NewServer()
	pb.RegisterServerStreamTalkServer(gRPCServer, &ServerStreamService{})
	err = gRPCServer.Serve(listener)
	if err != nil {
		log.Fatalf("gRPCServer serve err:%v", err)
	}
}

func (s *ServerStreamService) ListValue(req *pb.ServerStreamRequest, stv pb.ServerStreamTalk_ListValueServer) error {
	for i := 0; i < 15; i++ {
		err := stv.Send(&pb.ServerStreamResponse{
			Value: "hello" + req.Data,
			Code:  int32(i),
		})
		if err != nil {
			log.Fatalf("stv.Send pb.ServerStreamResponse err:%v", err)
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s *ServerStreamService) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Value: "hello" + req.Data,
	}, nil
}
