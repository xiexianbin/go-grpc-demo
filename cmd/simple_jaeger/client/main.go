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
	"os"
	"os/signal"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	simplejaeger "github.com/xiexianbin/go-grpc-demo/cmd/simple_jaeger"
	pb "github.com/xiexianbin/go-grpc-demo/proto"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	conn, err := simplejaeger.InitConn()
	if err != nil {
		log.Fatal(err)
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// The service name used to display traces in backends
			simplejaeger.ServiceName,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	shutdownTracerProvider, err := simplejaeger.InitTracerProvider(ctx, res, conn)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdownTracerProvider(ctx); err != nil {
			log.Fatalf("failed to shutdown TracerProvider: %s", err)
		}
	}()

	// shutdownMeterProvider, err := simplejaeger.InitMeterProvider(ctx, res, conn)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer func() {
	// 	if err := shutdownMeterProvider(ctx); err != nil {
	// 		log.Fatalf("failed to shutdown MeterProvider: %s", err)
	// 	}
	// }()

	conn, err = grpc.NewClient(
		":8000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			otelgrpc.UnaryClientInterceptor(),
		),
	)
	if err != nil {
		log.Fatal(err)
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
