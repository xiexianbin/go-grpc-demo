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
	"embed"
	_ "embed"
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	demov1 "github.com/xiexianbin/go-grpc-demo/gen/go/demo/v1"
	"github.com/xiexianbin/go-grpc-demo/pkg/demo"
)

const addr = "0.0.0.0:8000"

var (
	h             bool
	serverCrtPath string
	serverCrtName string
	serverKeyPath string
	//go:embed swagger-ui/*
	swaggerFS embed.FS
)

func init() {
	flag.BoolVar(&h, "help", false, "show help message")
	flag.StringVar(&serverCrtPath, "server-crt", "", "server crt file path")
	flag.StringVar(&serverCrtName, "crt-name", "go-grpc-demo", "server crt name, default is go-grpc-demo")
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

	// gRPC server
	grpcServer := grpc.NewServer(grpc.Creds(crt))
	// register grpc server
	demov1.RegisterDemoServiceServer(grpcServer, &demo.DemoServiceServer{})

	// gateway server
	ctx := context.Background()
	dopts := []grpc.DialOption{
		grpc.WithTransportCredentials(crt),
	}
	gwmux := runtime.NewServeMux()
	// register grpc-gateway service
	if err := demov1.RegisterDemoServiceHandlerFromEndpoint(ctx, gwmux, addr, dopts); err != nil {
		panic(err)
	}
	// handle grpc-gateway http service
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	// swagger-ui
	mux.Handle("/swagger-ui/", http.FileServer(http.FS(swaggerFS)))
	// prefix := "/swagger-ui/"
	// mux.Handle(prefix, http.StripPrefix(prefix, fileServer))

	log.Printf("listen at %s", addr)
	if serverCrtPath != "" && serverKeyPath != "" {
		err = http.ListenAndServeTLS(addr, serverCrtPath, serverKeyPath, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Header.Get("Content-Type"))
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r)
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
				grpcServer.ServeHTTP(w, r)
			} else {
				mux.ServeHTTP(w, r)
			}
		}), &http2.Server{}))
		if err != nil {
			panic(err)
		}
	}
}
