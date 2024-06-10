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

package demo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/xiexianbin/go-grpc-demo/pkg/util"
	pb "github.com/xiexianbin/go-grpc-demo/proto"
)

type DemoServiceServer struct {
	pb.UnimplementedDemoServiceServer
}

func (s *DemoServiceServer) Sum(ctx context.Context, numRequest *pb.NumRequest) (*pb.NumResponse, error) {
	for i := 0; i < 5; i++ {
		if ctx.Err() == context.Canceled {
			return nil, status.Errorf(codes.Canceled, "canceled by client")
		}
		time.Sleep(1 * time.Second)
	}

	numResponse := &pb.NumResponse{
		Result: numRequest.Nums[0] + numRequest.Nums[1],
	}
	return numResponse, nil
}

func (s *DemoServiceServer) Diff(ctx context.Context, numRequest *pb.NumRequest) (*pb.NumResponse, error) {
	numResponse := &pb.NumResponse{
		Result: numRequest.Nums[0] - numRequest.Nums[1],
	}
	return numResponse, nil
}

func (s *DemoServiceServer) Version(ctx context.Context, empty *emptypb.Empty) (*pb.VersionResponse, error) {
	version := &pb.VersionResponse{
		Version: "v0.1.0",
	}
	return version, nil
}

func (s *DemoServiceServer) ReadFile(ctx context.Context, filePath *pb.FilePath) (*pb.FileResponse, error) {
	_, err := os.Stat(filePath.GetPath())
	if err != nil {
		// os.IsNotExist(err)
		message := fmt.Sprintf("file path %s not exist", filePath.GetPath())
		log.Print(message)
		return nil, &util.ServerError{Message: message}
	}

	fileContent, err := os.ReadFile(filePath.GetPath())
	if err != nil {
		message := fmt.Sprintf("read file error %v", err)
		log.Print(message)
		return nil, &util.ServerError{Message: message}
	}

	return &pb.FileResponse{Content: fileContent}, nil
}
