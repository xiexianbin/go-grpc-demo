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

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/xiexianbin/go-grpc-demo/pkg/demo"

	demov1 "github.com/xiexianbin/go-grpc-demo/gen/go/demo/v1"
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

	var err error
	var crt credentials.TransportCredentials
	if serverCrtPath != "" && serverKeyPath != "" {
		crt, err = credentials.NewServerTLSFromFile(serverCrtPath, serverKeyPath)
		if err != nil {
			log.Panicf("load tls failed, %s", err)
		}
	} else {
		crt = credentials.TransportCredentials(insecure.NewCredentials())
	}

	// gRPC method
	server := grpc.NewServer(grpc.Creds(crt))
	demov1.RegisterDemoServiceServer(server, &demo.DemoServiceServer{})

	// http method
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello word!"))
	})

	// Listener
	addr := "0.0.0.0:8000"
	log.Printf("listen at %s", addr)
	if serverCrtPath != "" && serverKeyPath != "" {
		err = http.ListenAndServeTLS(addr, serverCrtPath, serverKeyPath, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Header.Get("Content-Type"))
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				server.ServeHTTP(w, r)
			} else {
				mux.ServeHTTP(w, r)
			}
		}))
		if err != nil {
			panic(err)
		}
	} else {
		// use h2c without tls by HTTP/2
		err = http.ListenAndServe(addr, h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				server.ServeHTTP(w, r)
			} else {
				mux.ServeHTTP(w, r)
			}
		}), &http2.Server{}))
		if err != nil {
			panic(err)
		}
	}
}
