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
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/xiexianbin/go-grpc-demo/proto"
)

var (
	h             bool
	clientCrtPath string
	clientKeyPath string
)

func init() {
	flag.BoolVar(&h, "help", false, "show help message")
	flag.StringVar(&clientCrtPath, "client-crt", "", "client crt file path")
	flag.StringVar(&clientKeyPath, "client-key", "", "client key file path")

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

	crt, err := credentials.NewClientTLSFromFile(clientCrtPath, "go-grpc-demo")
	if err != nil {
		log.Panicf("load tls failed, %s", err)
	}

	conn, err := grpc.NewClient("127.0.0.1:8000", grpc.WithTransportCredentials(crt))
	// conn, err := grpc.NewClient("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server %v", err)
	}
	defer conn.Close()

	client := pb.NewDemoServiceClient(conn)
	resp, err := client.Sum(context.Background(), &pb.NumRequest{
		Nums: []int64{1, 2},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("sum: %d", resp.GetResult())
}
