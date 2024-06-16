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
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	demov1 "github.com/xiexianbin/go-grpc-demo/gen/go/demo/v1"
)

func main() {
	var err error
	conn, err := grpc.NewClient(
		":8000",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := demov1.NewStreamServiceClient(conn)
	err = callList(client, &demov1.ListRequest{Pt: &demov1.StreamPoint{Name: "callList", Value: 0}})
	if err != nil {
		log.Fatalf("callList err: %s", err)
	}

	err = callRecord(client, &demov1.RecordRequest{Pt: &demov1.StreamPoint{Name: "callRecord", Value: 0}})
	if err != nil {
		log.Fatalf("callRecord err: %s", err)
	}

	err = callRoute(client, &demov1.RouteRequest{Pt: &demov1.StreamPoint{Name: "callRoute", Value: 0}})
	if err != nil {
		log.Fatalf("callRoute err: %s", err)
	}
}

func callList(client demov1.StreamServiceClient, r *demov1.ListRequest) error {
	stream, err := client.List(context.Background(), r)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("callList resp pt: %v", resp.Pt)
	}
	return nil
}

func callRecord(client demov1.StreamServiceClient, r *demov1.RecordRequest) error {
	stream, err := client.Record(context.Background())
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("callRecord resp pt: %v", resp.Pt)

	return nil
}

func callRoute(client demov1.StreamServiceClient, r *demov1.RouteRequest) error {
	stream, err := client.Route(context.Background())
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}

		resp := demov1.RecordResponse{}
		err = stream.RecvMsg(&resp)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("callRoute resp pt: %v", resp.Pt)
	}
	return stream.CloseSend()
}
