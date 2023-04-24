package main

import (
	"context"
	pb "gRPC_protoc/server_stream/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

// var client pb.ServerStreamTalkClient
type clietn struct {
	ctx       context.Context
	client    pb.ServerStreamTalkClient
	pauseFlag bool
}

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc dial err:%v", err)
	}
	defer conn.Close()
	c := newClient()
	c.client = pb.NewServerStreamTalkClient(conn)

	for {
		if !c.pauseFlag {
			c.listValue()
		}
		time.Sleep(1 * time.Second)
	}

}

func newClient() *clietn {
	return &clietn{
		ctx:       context.Background(),
		pauseFlag: false,
	}
}

func (c *clietn) listValue() {
	req := pb.ServerStreamRequest{
		Data: "server stream gRPC",
	}
	// 调用自己的服务
	stream, err := c.client.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("client.ListValue err:%v", err)
	}
	defer stream.CloseSend()
	//stop := false
	for {
		res, err := stream.Recv()
		if err == io.EOF { // 判断io流是否输出完
			break
		}
		if err != nil {
			log.Fatalf("stream.Recv err:%v", err)
		}
		if res.Code == int32(5) {
			//stop = true
			c.pauseFlag = true
			stream.CloseSend()
			log.Println("stream paused")
			break
		}
		log.Println("data:"+res.Value, "code:", res.Code)
	}

}
