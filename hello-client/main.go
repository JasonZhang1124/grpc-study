package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-study/hello-server/proto"
	"log"
)

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "kuangshen",
		"appKey": "123123",
	}, nil
}
func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	//creds,_ :=credentials.NewClientTLSFromFile("密钥文件绝对路径","*.xxx.com")
	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds))

	//连接到server端，此处禁用安全传输，没有加密和验证
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	conn, err := grpc.Dial("127.0.0.1:9090", opts...)
	if err != nil {
		log.Fatalf("did not commect:%v", err)
		return
	}
	defer conn.Close()

	//建立链接
	client := pb.NewSayHelloClient(conn)

	//执行rpc的调用（这个方法在服务端来实现返回结果）
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "kuangshen"})
	if err != nil {
		return
	}
	fmt.Println(resp.GetResponseMsg())
}
