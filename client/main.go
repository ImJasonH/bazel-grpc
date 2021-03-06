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
	"context"
	"flag"
	"log"

	pb "github.com/imjasonh/bazel-grpc/proto/proto"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost:50051", "Server address to connect to")
	name = flag.String("name", "Bob", "Name to send")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewSimpleClient(conn)

	fooRep, err := client.Foo(ctx, &pb.FooRequest{
		Name: *name,
	})
	if err != nil {
		log.Fatalf("Foo: %v", err)
	}
	log.Println("Foo: %s", fooRep.Message)

	barRep, err := client.Bar(ctx, &pb.BarRequest{
		Name: *name,
	})
	if err != nil {
		log.Fatalf("Bar: %v", err)
	}
	log.Println("Bar: %s", barRep.Message)
}
