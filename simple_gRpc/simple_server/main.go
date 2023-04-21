package main

import (
	"context"
	"errors"
	"fmt"
	pb "gRPC_protoc/simple_gRpc/simple_server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

type SimpleService struct {
	pb.UnimplementedSimpleSayServer
	//没有将`pb.UnimplementedXXXXServer`结构体作为匿名字段嵌入到自定义的服务器结构体中，那么在注册您的服务时会遇到错误。
	//
	//具体来说，当使用gRPC服务器的`RegisterService`函数注册您的服务时，gRPC运行时会检查服务是否实现了它所声明的所有gRPC接口方法。
	//如果没有实现某个方法，gRPC将返回一个错误，指示该方法未实现。而在Go中，为了避免在自定义的服务器结构体中实现每个未实现的方法，
	//可以使用`pb.UnimplementedXXXXServer`作为匿名字段。
	//
	//当使用`pb.UnimplementedXXXXServer`作为匿名字段时，您只需要实现您所感兴趣的方法即可，其他方法将自动转发到未实现的默认实现中。
	//这有助于避免非常冗长的代码，同时确保您已经正确实现了期望的gRPC接口方法，可以更加高效地实现自己的gRPC服务器。
}

func main() {
	// 加密
	//Creds, err := credentials.NewServerTLSFromFile("E:\\GoProject\\src\\gRPC_protoc\\key\\an.pem",
	//	"E:\\GoProject\\src\\gRPC_protoc\\key\\an.key")
	//if err != nil {
	//	log.Fatalf("credentials.NewServerTLSFromFile err:%v", err)
	//}
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("net.listen err:%v", err)
	}
	log.Println(":8080 net listening...")
	// 创建 grpc 服务
	//grpcServer := grpc.NewServer(grpc.Creds(Creds))
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	// 在gRPC服务上 注册 自己的服务
	pb.RegisterSimpleSayServer(grpcServer, &SimpleService{})
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.serve err:%v", err)
	}
}

func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	// 获取元数据信息
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Printf("%+v\n", md)
	if !ok {
		return nil, errors.New("为传输token")
	}
	var appID string
	var appKey string
	if v, ok := md["appid"]; ok {
		appID = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}
	// 根据用户 id  判断 appID
	if appID != "AN" || appKey != "1" {
		return nil, errors.New("toekn 不正确")
	}

	a := []string{"AN", "quan"}
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello" + req.Data,
		Name:  a,
	}
	return &res, nil
}
