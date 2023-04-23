package main

import (
	pb "gRPC_protoc/server_stream/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
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

func (s *ServerStreamService) Route(req *pb.ServerStreamRequest, stv pb.ServerStreamTalk_ListValueServer) error {
	for i := 0; i < 5; i++ {
		err := stv.Send(&pb.ServerStreamResponse{
			Value: "hello" + req.Data + strconv.Itoa(i),
		})
		if err != nil {
			log.Fatalf("stv.Send pb.ServerStreamResponse err:%v", err)
		}
	}
	return nil
}
