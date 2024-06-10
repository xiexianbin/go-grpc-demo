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
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	pb "github.com/xiexianbin/go-grpc-demo/proto"
)

func main() {
	conn, err := grpc.NewClient(
		":8000",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithDeadline(
		context.Background(),
		time.Now().Add(time.Duration(3*time.Second)))
	defer cancel()

	client := pb.NewDemoServiceClient(conn)
	resp, err := client.Sum(ctx, &pb.NumRequest{
		Nums: []int64{1, 2},
	})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			log.Fatalf("err: %d: %s", statusErr.Code(), statusErr.Message())
		}
	}

	log.Printf("sum: %d", resp.GetResult())
}
