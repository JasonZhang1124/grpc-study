package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	pb "grpc-study/hello-server/proto"
	"net"
)

// hello server
type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 获取元素信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("没有传输token")
	}
	var appId string
	var appKey string
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}

	if appId != "kuangshen" || appKey != "123123" {
		return nil, errors.New("token不正确")
	}

	fmt.Printf("hello" + req.RequestName)
	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func main() {
	//========== ==== ====
	//TSL认证
	//========== ==== ====
	// 两个参数分别是 cretFile ，keyFile
	// 自签名证书文件和私钥文件
	//creds, _ := credentials.NewServerTLSFromFile("D:绝对路径.test.pem", "D:绝对路径.test.key")

	//	开启端口
	listen, _ := net.Listen("tcp", ":9090")
	// 创建grpc服务
	//grpcServer := grpc.NewServer(grpc.Creds(creds))
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	//	在grpc服务端中去注册我们自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})

	//	启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve:%v", err)
		return
	}
}
