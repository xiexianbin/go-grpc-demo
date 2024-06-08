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

package stream

import (
	"io"
	"log"

	pb "github.com/xiexianbin/go-grpc-demo/proto"
)

type StreamServiceServer struct {
	pb.UnimplementedStreamServiceServer
}

// List 服务器端流式 RPC
func (s *StreamServiceServer) List(r *pb.StreamRequest, stream pb.StreamService_ListServer) error {
	for i := 0; i < 10; i++ {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(i),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// Record 客户端流式 RPC
func (s *StreamServiceServer) Record(stream pb.StreamService_RecordServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.StreamResponse{
				Pt: &pb.StreamPoint{
					Name:  "gRPC Stream Server: Record",
					Value: -1,
				},
			})
		}
		if err != nil {
			return err
		}
		log.Printf("StreamService.Record resp pt: %v", resp.Pt)
	}
}

// Route 双向流式 RPC
func (s *StreamServiceServer) Route(stream pb.StreamService_RouteServer) error {
	i := 0
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("StreamService.Route pt: %v", r.Pt)

		err = stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  "gPRC StreamService: Route",
				Value: int32(i),
			},
		})
		if err != nil {
			return err
		}

		i++
	}
}
