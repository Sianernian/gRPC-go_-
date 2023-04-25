package main

import (
	"context"
	pb "gRPC_protoc/client_stream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"strconv"
	"time"
)

var client pb.ClientStreamGoClient

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                15 * time.Second,
		Timeout:             2 * time.Second,
		PermitWithoutStream: false,
	}))
	if err != nil {
		log.Fatalf("grpc dial err:%v", err)
	}
	defer conn.Close()
	client = pb.NewClientStreamGoClient(conn)
	routeList()
}

func routeList() {
	stream, err := client.RouteList(context.Background())
	if err != nil {
		log.Fatalf("client routeList err:%v", err)
	}
	for i := 0; i < 10; i++ {
		err = stream.Send(&pb.ClientStreamRequest{
			Data: "client stream grpc" + strconv.Itoa(i),
		})
		if err != nil {
			log.Fatalf("stream.Send err:%v", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("stream.CloseAndRecv() err:%v", err)
	}
	log.Println(res)
}
