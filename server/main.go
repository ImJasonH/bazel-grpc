package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/imjasonh/bazel-grpc/proto/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct{}

func (s *server) Foo(ctx context.Context, req *pb.FooRequest) (*pb.FooReply, error) {
	return &pb.FooReply{
		Message: fmt.Sprintf("Hello, %s", req.Name),
	}, nil
}

func (s *server) Bar(ctx context.Context, req *pb.BarRequest) (*pb.BarReply, error) {
	return &pb.BarReply{
		Message: fmt.Sprintf("Farewell, sweet %s", req.Name),
	}, nil
}
