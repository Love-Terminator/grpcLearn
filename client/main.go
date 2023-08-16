package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpcLearn/helloworld"
	"log"
	"os"
	"time"
)

const (
	address     = "127.0.0.1:8080"
	defaultName = "world"
)

func main() {
	// 连接到grpc服务端
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// 结束之前关闭conn
	defer conn.Close()

	// 建立客户端
	client := pb.NewGreetClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用服务端方法
	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("couldn not greet: %v", err)
	}
	fmt.Println(r.Message)

	// 调用服务端方法
	people, err := client.SearchPeople(ctx, &pb.SearchRequest{People_Name: name, People_Age: 18})
	if err != nil {
		log.Fatalf("couldn not greet: %v", err)
	}
	fmt.Println(people)
}
