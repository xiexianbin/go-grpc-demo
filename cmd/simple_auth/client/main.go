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

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	demov1 "github.com/xiexianbin/go-grpc-demo/gen/go/demo/v1"
)

type Token struct {
	Key, Secret string
}

func (t *Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"key": t.Key, "secret": t.Secret}, nil
}

func (t *Token) RequireTransportSecurity() bool {
	return false
}

func main() {
	conn, err := grpc.NewClient(
		":8000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(&Token{
			Key:    "foo",
			Secret: "bar",
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := demov1.NewDemoServiceClient(conn)
	resp, err := client.Sum(context.Background(), &demov1.SumRequest{
		Nums: []int64{1, 2},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("sum: %d", resp.GetResult())
}
