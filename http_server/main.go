package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gw "github.com/penril0326/hello_grpc/proto/calculator"
	"google.golang.org/grpc"
)

var (
	// command-line options:
	// gRPC server endpoint
	// 可以指定grpc
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:5555", "gRPC server endpoint")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	fmt.Printf("Starting serve http server...")
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterCalculatorServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":5656", mux)
}
