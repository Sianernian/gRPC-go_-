package main

import (
	"context"
	pb "gRPC_protoc/server_stream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"io"
	"log"
	"time"
)

const (
	address = ":8080"
)

func main() {
	var conn *grpc.ClientConn
	var err error
	conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                15 * time.Second,
		Timeout:             2 * time.Second,
		PermitWithoutStream: false, // 如果为true，则即使客户端没有活动流，也会发送keepalive
	}))
	if err != nil {
		log.Fatalf("grpc.Dial err:%v", err)
	}

	defer conn.Close()

	client := pb.NewServerStreamTalkClient(conn)
	//心跳检测
	//ticker := time.NewTicker(time.Second * 5)
	//defer ticker.Stop()
	//go func() {
	//	for range ticker.C {
	//		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	//		defer cancel()
	//
	//		if _, err := client.Ping(ctx, &pb.PingRequest{Data: "ping"}); err != nil {
	//			log.Printf("ping error: %v", err)
	//		}
	//
	//	}
	//}()

	// 执行其他操作
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		stream, err := client.ListValue(ctx, &pb.ServerStreamRequest{})
		defer stream.CloseSend()
		if err != nil {
			log.Fatalf("client.ListValue(ctx, &pb.ServerStreamRequest{}):%v", err)
		}
		for {
			res, err := stream.Recv()
			if err == io.EOF { // 判断io流是否输出完
				break
			}
			if err != nil {
				log.Fatalf("stream.Recv err:%v", err)
			}
			if res.Code == int32(5) {
				stream.CloseSend()
				time.Sleep(16 * time.Second)
				log.Println("stream paused")
				//	进行重连
				for {
					conn, err = grpc.Dial(address, grpc.WithInsecure())
					if err != nil {
						log.Printf("failed to reconnect: %v", err)
						time.Sleep(time.Second)
						continue
					}
					client = pb.NewServerStreamTalkClient(conn)
					break
				}
				continue
				//break
			}
			log.Println("data:"+res.Value, "code:", res.Code)
		}
		time.Sleep(time.Second)
	}
}
