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
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/xiexianbin/go-grpc-demo/pkg/stream"

	demov1 "github.com/xiexianbin/go-grpc-demo/gen/go/demo/v1"
)

func main() {
	server := grpc.NewServer()
	demov1.RegisterStreamServiceServer(server, &stream.StreamServiceServer{})

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

	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
