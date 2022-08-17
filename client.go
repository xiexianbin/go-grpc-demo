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
	"log"
	"time"

	googlerpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	dgrpc "github.com/xiexianbin/go-rpc-demo/grpc"
)

func main() {
	//creds := credentials.NewTLS(nil)
	//cc, err := googlerpc.Dial("127.0.0.1:8000", googlerpc.WithTransportCredentials(creds))
	cc, err := googlerpc.Dial("127.0.0.1:8000", googlerpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot dial server %v", err)
	}
	defer cc.Close()

	rpcClient := dgrpc.NewServiceClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// version
	version, err := rpcClient.Version(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("error happen when call gRPC client: %s", err.Error())
	}
	log.Printf("version: %s", version)

	// sum
	nums := &dgrpc.NumRequest{
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
	filePath := &dgrpc.FilePath{
		Path: "/etc/hosts",
	}
	fileContent, err := rpcClient.ReadFile(ctx, filePath)
	if err != nil {
		log.Fatalf("error happen when call gRPC client: %s", err.Error())
	}
	log.Printf("fileContent: %s", fileContent)
}
