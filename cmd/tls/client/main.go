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
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/xiexianbin/go-grpc-demo/proto"
)

var (
	ch            bool
	clientCrtPath string
	clientKeyPath string
)

func init() {
	flag.BoolVar(&ch, "help", false, "show help message")
	flag.StringVar(&clientCrtPath, "client-crt", "", "client crt file path")
	flag.StringVar(&clientKeyPath, "client-key", "", "client key file path")

	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	if ch {
		flag.Usage()
		return
	}

	crt, err := credentials.NewClientTLSFromFile(clientCrtPath, "go-grpc-demo")
	if err != nil {
		log.Panicf("load tls failed, %s", err)
	}

	conn, err := grpc.NewClient("127.0.0.1:8000", grpc.WithTransportCredentials(crt))
	if err != nil {
		log.Fatalf("cannot dial server %v", err)
	}
	defer conn.Close()

	rpcClient := pb.NewDemoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// version
	version, err := rpcClient.Version(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("error happen when call gRPC client: %s", err.Error())
	}
	log.Printf("version: %s", version)

	// sum
	nums := &pb.NumRequest{
		Nums: []int64{1, 2},
	}
	sum, err := rpcClient.Sum(ctx, nums)
	if err != nil {
		log.Fatalf("error happen when call gRPC client: %s", err.Error())
	}
	log.Printf("sum: %s", sum)

	// diff
	diff, err := rpcClient.Diff(ctx, nums)
	if err != nil {
		log.Fatalf("error happen when call gRPC client: %s", err.Error())
	}
	log.Printf("diff: %s", diff)

	// read file
	filePath := &pb.FilePath{
		Path: "/etc/hosts",
	}
	fileContent, err := rpcClient.ReadFile(ctx, filePath)
	if err != nil {
		log.Fatalf("error happen when call gRPC client: %s", err.Error())
	}
	log.Printf("fileContent: %s", fileContent)
}
