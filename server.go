/*
Copyright [2022] [xiexianbin.cn]

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
	googlerpc "google.golang.org/grpc"
	"log"
	"net"

	dgrpc "github.com/xiexianbin/go-rpc-demo/grpc"
	dgrpcserver "github.com/xiexianbin/go-rpc-demo/grpc/server"
)

var server = &dgrpcserver.DemoServiceServer{}

func main() {
	// Listener
	addr := ":8000"
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

	var s = googlerpc.NewServer()
	dgrpc.RegisterServiceServer(s, server)
	log.Println("grpc server listen on", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}
