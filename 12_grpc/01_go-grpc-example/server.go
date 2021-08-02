package main

import (
	"context"
	pb "example/go-grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SearchService struct {
}

const PORT = "9001"

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " api"}, nil
}

func main() {
	//创建 gRPC Server 对象，你可以理解为它是 Server 端的抽象对象
	server := grpc.NewServer()

	/*
		SearchService（其包含需要被调用的服务端接口）注册到 gRPC Server 的内部注册中心。
		这样可以在接受到请求时，通过内部的服务发现，发现该服务端接口并转接进行逻辑处理
	*/
	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)

	if err != nil {
		log.Fatalf("net.listen err:%v", err)
	}

	server.Serve(lis)
}

//context
/*
withCancel

WithCancel返回一个带有新的Done通道的父类副本。

当返回的取消函数被调用时，上下文的Done通道被关闭。或者当父级上下文的Done通道被关闭时（以先发生者为准）。

取消这个上下文会释放与之相关的资源，一旦在此上下文中运行的操作完成，就应立即调用cancel。
*/
