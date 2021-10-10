package main

import (
	"context"
	//"google.golang.org/grpc/examples/helloworld/helloworld"

	//"google.golang.org/grpc/examples/helloworld/helloworld"
	"log"
	"net"


	"google.golang.org/grpc"
	pb "helloworld/helloworld"
)

const (
	port = ":50051"
)

// server is used to implement example.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements example.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}