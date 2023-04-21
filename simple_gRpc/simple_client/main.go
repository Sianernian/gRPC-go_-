package main

import (
	"context"
	"fmt"
	pb "gRPC_protoc/simple_gRpc/simple_server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

//// Token 认证
//type PerRPCCredentials interface {
//	// 获取元数据信息 就是客户端提供的key,valu， context 用于控制和取消， url请求的url
//	GetRequestMetadata(ctx context.Context, url ...string) (map[string]string, error)
//	// 是否需要 SSL\TLS认证，ture 启用，必须要加上TLS认证
//	RequireTransportSecurity() bool
//}

type ClientTokenAuth struct {
}

func (c *ClientTokenAuth) GetRequestMetadata(ctx context.Context, url ...string) (map[string]string, error) {

	return map[string]string{
		"appID":  "AN",
		"appKey": "1",
	}, nil
}
func (c *ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	//启用TLS 安全认证
	//creds, err := credentials.NewClientTLSFromFile("E:\\GoProject\\src\\gRPC_protoc\\key\\an.pem",
	//	"http://127.0.0.1:8090")
	//if err != nil {
	//	log.Fatalf("credentials.NewClientTLSFromFile err:%v", err)
	//}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))
	conn, err := grpc.Dial(":8080", opts...)
	//conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(creds))
	//conn, err := grpc.Dial(":8080", grpc.WithInsecure()) // 没用进行加密传输，不安全
	if err != nil {
		log.Fatalf("grpc dial err:%v", err)
	}
	defer conn.Close()

	client := pb.NewSimpleSayClient(conn)
	req := pb.SimpleRequest{
		Data: "grpc",
	}
	res, err := client.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("client.Route err:", err)
	}
	fmt.Println(res)
}
