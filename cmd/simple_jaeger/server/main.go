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

// ref https://github.com/grpc-ecosystem/go-grpc-middleware/blob/main/examples

package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"google.golang.org/grpc"

	simplejaeger "github.com/xiexianbin/go-grpc-demo/cmd/simple_jaeger"
	demov1 "github.com/xiexianbin/go-grpc-demo/gen/go/demo/v1"
	"github.com/xiexianbin/go-grpc-demo/pkg/demo"
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

	// tracer := otel.Tracer("test-tracer")
	// meter := otel.Meter("test-meter")

	// grpc server
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			// Order matters e.g. tracing interceptor have to create span first for the later exemplars to work.
			otelgrpc.UnaryServerInterceptor(),
		),
	}
	server := grpc.NewServer(opts...)
	demov1.RegisterDemoServiceServer(server, &demo.DemoServiceServer{})

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
	_ = server.Serve(listener)
}
