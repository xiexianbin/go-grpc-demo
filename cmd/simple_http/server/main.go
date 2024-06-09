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
	"flag"
	"log"
	"net/http"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/xiexianbin/go-grpc-demo/pkg/demo"
	pb "github.com/xiexianbin/go-grpc-demo/proto"
)

var (
	h             bool
	serverCrtPath string
	serverKeyPath string
)

func init() {
	flag.BoolVar(&h, "help", false, "show help message")
	flag.StringVar(&serverCrtPath, "server-crt", "", "server crt file path")
	flag.StringVar(&serverKeyPath, "server-key", "", "server key file path")

	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	if h {
		flag.Usage()
		return
	}

	crt, err := credentials.NewServerTLSFromFile(serverCrtPath, serverKeyPath)
	if err != nil {
		log.Panicf("load tls failed, %s", err)
	}

	// gRPC method
	server := grpc.NewServer(grpc.Creds(crt))
	pb.RegisterDemoServiceServer(server, &demo.DemoServiceServer{})

	// http method
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello word!"))
	})

	// Listener
	addr := "0.0.0.0:8000"
	log.Printf("listen at %s", addr)
	http.ListenAndServeTLS(addr, serverCrtPath, serverKeyPath, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header.Get("Content-Type"))
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			server.ServeHTTP(w, r)
		} else {
			mux.ServeHTTP(w, r)
		}
	}))
}
