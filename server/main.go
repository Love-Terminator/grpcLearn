package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpcLearn/helloworld"
	"io"
	"log"
	"net"
	"os"
)

const port = ":8080"

type server struct {
	pb.GreetServer
	people *People
	book   []byte
}

type People struct {
	name    string
	age     int
	high    float64
	hobbies []string
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (out *pb.HelloReply, err error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SearchPeople(ctx context.Context, in *pb.SearchRequest) (out *pb.SearchReponse, err error) {
	return &pb.SearchReponse{
		People_Name: in.People_Name,
		People_Age:  in.People_Age,
		People_High: s.people.high,
		Hobbies:     s.people.hobbies,
		People_Read: s.book}, nil
}

func main() {
	// 配置端口
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed run server: %v\n", err)
	}

	var buff [128]byte
	var content []byte
	file, err := os.Open("book.txt")
	if err != nil {
		return
	}
	for {
		n, err := file.Read(buff[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		content = append(content, buff[:n]...)
	}

	people := People{
		high:    175.6,
		hobbies: []string{"swim", "basketball", "run"},
	}

	// 创建服务端并注册
	s := grpc.NewServer()
	pb.RegisterGreetServer(s, &server{people: &people, book: content})

	// 监听端口
	log.Println("Server is running~~~")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("fail to serve: %v", err)
	}
}
