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

	demov1 "github.com/xiexianbin/go-grpc-demo/gen/go/demo/v1"
	"github.com/xiexianbin/go-grpc-demo/pkg/util"
)

type DemoServiceServer struct {
	demov1.UnimplementedDemoServiceServer
}

func (s *DemoServiceServer) Sum(ctx context.Context, sumRequest *demov1.SumRequest) (*demov1.SumResponse, error) {
	for i := 0; i < 5; i++ {
		if ctx.Err() == context.Canceled {
			return nil, status.Errorf(codes.Canceled, "canceled by client")
		}
		time.Sleep(1 * time.Second)
	}

	return &demov1.SumResponse{
		Result: sumRequest.Nums[0] + sumRequest.Nums[1],
	}, nil
}

func (s *DemoServiceServer) Diff(ctx context.Context, diffRequest *demov1.DiffRequest) (*demov1.DiffResponse, error) {
	return &demov1.DiffResponse{
		Result: diffRequest.Nums[0] - diffRequest.Nums[1],
	}, nil
}

func (s *DemoServiceServer) Version(ctx context.Context, versionRequest *demov1.VersionRequest) (*demov1.VersionResponse, error) {
	return &demov1.VersionResponse{
		Version: "v0.1.0",
	}, nil
}

func (s *DemoServiceServer) ReadFile(ctx context.Context, readFileRequest *demov1.ReadFileRequest) (*demov1.ReadFileResponse, error) {
	_, err := os.Stat(readFileRequest.GetPath())
	if err != nil {
		// os.IsNotExist(err)
		message := fmt.Sprintf("file path %s not exist", readFileRequest.GetPath())
		log.Print(message)
		return nil, &util.ServerError{Message: message}
	}

	fileContent, err := os.ReadFile(readFileRequest.GetPath())
	if err != nil {
		message := fmt.Sprintf("read file error %v", err)
		log.Print(message)
		return nil, &util.ServerError{Message: message}
	}

	return &demov1.ReadFileResponse{Content: fileContent}, nil
}
