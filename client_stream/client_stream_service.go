package main

import (
	pb "gRPC_protoc/client_stream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"io"
	"log"
	"net"
	"time"
)

type ClientStreamService struct {
	pb.UnimplementedClientStreamGoServer
}

func main() {
	listenner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("net.Listen err:%V", err)
	}

	keepAlivePar := keepalive.ServerParameters{
		Time:    15 * time.Second,
		Timeout: 2 * time.Second,
	}
	
	grpcServer := grpc.NewServer(grpc.KeepaliveParams(keepAlivePar))
	pb.RegisterClientStreamGoServer(grpcServer, &ClientStreamService{})
	err = grpcServer.Serve(listenner)
	if err != nil {

		log.Fatalf("grpcServer serve err；%V", err)
	}

}

func (c *ClientStreamService) RouteList(srv pb.ClientStreamGo_RouteListServer) error {
	for {
		// 从流中 获取信息
		res, err := srv.Recv()
		if err == io.EOF {
			return srv.SendAndClose(&pb.ClientStreamResponse{
				Code:  int32(1),
				Value: "ok",
			})
		}
		if err != nil {
			log.Fatalf("srv.Recv() err:%v", err)
		}
		log.Println(res.Data)
	}
	return nil
}
