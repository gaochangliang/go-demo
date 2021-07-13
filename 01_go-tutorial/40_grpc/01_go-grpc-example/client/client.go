package main

import (
	"context"
	pb "example/go-grpc/proto"
	"google.golang.org/grpc"
	"log"
)

const PORT = "9001"

func main() {

	//连接客户端
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("grpc.Dial err %v", err)
	}
	defer conn.Close()

	//创建 SearchService 的客户端对象
	client := pb.NewSearchServiceClient(conn)

	//发送 RPC 请求，等待同步响应，得到回调后返回响应结果
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "GRPC",
	})

	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	//输出响应结果
	log.Printf("resp: %s", resp.GetResponse())
}
