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
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/xiexianbin/go-grpc-demo/proto"
)

type ServerError struct {
	message string
}

func (e *ServerError) Error() string {
	return e.message
}

type DemoService struct {
	pb.UnimplementedServiceServer
}

func (s *DemoService) Sum(ctx context.Context, numRequest *pb.NumRequest) (*pb.NumResponse, error) {
	numResponse := &pb.NumResponse{
		Result: numRequest.Nums[0] + numRequest.Nums[1],
	}
	return numResponse, nil
}

func (s *DemoService) Diff(ctx context.Context, numRequest *pb.NumRequest) (*pb.NumResponse, error) {
	numResponse := &pb.NumResponse{
		Result: numRequest.Nums[0] - numRequest.Nums[1],
	}
	return numResponse, nil
}

func (s *DemoService) Version(ctx context.Context, empty *emptypb.Empty) (*pb.VersionResponse, error) {
	version := &pb.VersionResponse{
		Version: "v0.1.0",
	}
	return version, nil
}

func (s *DemoService) ReadFile(ctx context.Context, filePath *pb.FilePath) (*pb.FileResponse, error) {
	_, err := os.Stat(filePath.GetPath())
	if err != nil {
		// os.IsNotExist(err)
		message := fmt.Sprintf("file path %s not exist", filePath.GetPath())
		log.Print(message)
		return nil, &ServerError{message: message}
	}

	fileContent, err := os.ReadFile(filePath.GetPath())
	if err != nil {
		message := fmt.Sprintf("read file error %v", err)
		log.Print(message)
		return nil, &ServerError{message: message}
	}

	return &pb.FileResponse{Content: fileContent}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterServiceServer(server, &DemoService{})

	// Listener
	addr := "0.0.0.0:8000"
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
	log.Printf("listen at %s", addr)

	server.Serve(listener)
}
