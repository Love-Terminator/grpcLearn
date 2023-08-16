package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpcLearn/helloworld"
	"log"
	"net"
)

const port = ":8080"

type server struct {
	pb.GreetServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (out *pb.HelloReply, err error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	// 配置端口
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed run server: %v\n", err)
	}

	// 创建服务端并注册
	s := grpc.NewServer()
	pb.RegisterGreetServer(s, &server{})

	// 监听端口
	log.Println("Server is running~~~")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("fail to serve: %v", err)
	}
}
