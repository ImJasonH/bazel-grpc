// Copyright 2016 Google, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	"log"
	"net"
	"context"

	pb "github.com/imjasonh/bazel-grpc/proto/proto"
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
