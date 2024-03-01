package main

import (
	"context"
	"fmt"
	"hello_server/pb"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Reply: "Hello " + in.Name}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Println("failed to listen: ", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		fmt.Println("failed to server", err)
		return
	}

}
