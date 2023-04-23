package main

import (
	"context"
	pb "gRPC_protoc/server_stream/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

var client pb.ServerStreamTalkClient

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc dial err:%v", err)
	}
	defer conn.Close()

	client = pb.NewServerStreamTalkClient(conn)

}

func listValue() {
	req := pb.ServerStreamRequest{
		Data: "server stream gRPC",
	}
	// 调用自己的服务
	stream, err := client.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("client.ListValue err:%v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF { // 判断io流是否输出完
			break
		}
		if err != nil {
			log.Fatalf("stream.Recv err:%v", err)
		}
		log.Println("data:"+res.Value, "code:", res.Code)
	}

}
