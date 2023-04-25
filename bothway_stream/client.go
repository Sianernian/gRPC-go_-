package main

import (
	"context"
	pb "gRPC_protoc/bothway_stream/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
)

var streamClient pb.BothWayTaleClient

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err:%v", err)
	}
	defer conn.Close()
	streamClient = pb.NewBothWayTaleClient(conn)
	conver()
}

func conver() {
	stream, err := streamClient.Cover(context.Background())
	if err != nil {
		log.Fatalf("client.Cover err:%v", err)
	}

	for i := 0; i < 10; i++ {
		err = stream.Send(&pb.BothwayRequest{Question: "stream client rpc" + strconv.Itoa(i)})
		if err != nil {
			log.Fatalf("stream.Send err:%v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream.recv err:%v", err)
		}
		log.Println(res.Answer)
	}
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("stream.CloseSend() err:%v", err)
	}
}
