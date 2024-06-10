/*
Copyright [2024] [xiexianbin.cn]

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/xiexianbin/go-grpc-demo/pkg/demo"
	pb "github.com/xiexianbin/go-grpc-demo/proto"
)

type Token struct {
	key, secret string
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = status.Errorf(codes.Unauthenticated, "bad key or secret")
		return
	}

	var token Token
	if value, ok := md["key"]; ok {
		token.key = value[0]
	}
	if value, ok := md["secret"]; ok {
		token.secret = value[0]
	}

	fmt.Printf("token: %v", token)

	if token.key != "foo" || token.secret != "bar" {
		err = status.Errorf(codes.Unauthenticated, "bad key or secret")
		return
	}

	resp, err = handler(ctx, req)
	return
}

func main() {
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			AuthInterceptor,
		),
	}
	server := grpc.NewServer(opts...)
	pb.RegisterDemoServiceServer(server, &demo.DemoServiceServer{})

	// Listener
	addr := "0.0.0.0:8000"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.Println("grpc server closed.")
	}(listener)
	log.Printf("listen at %s", addr)

	server.Serve(listener)
}
