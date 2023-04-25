package main

import (
	pb "gRPC_protoc/bothway_stream/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
)

type BothWayStream struct {
	pb.UnimplementedBothWayTaleServer
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("net.listen err:%v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBothWayTaleServer(grpcServer, &BothWayStream{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer serve err:%v", err)
	}
}

func (b *BothWayStream) Cover(srv pb.BothWayTale_CoverServer) error {
	n := 1
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("srv.recv err:%v", err)
		}
		err = srv.Send(&pb.BothwayResponse{
			Answer: "service stream answer: the" + strconv.Itoa(n) + "question is " + req.Question,
		})
		if err != nil {
			log.Fatalf("srv.Send err:%v", err)
		}
		n++
		log.Printf("from stream client question: %s", req.Question)
	}
}
