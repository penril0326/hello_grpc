package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gw "github.com/penril0326/hello_grpc/proto/calculator"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
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

	mux := runtime.NewServeMux(runtime.WithForwardResponseOption(httpResponseModifier),
		emitDefault())
	// mux := runtime.NewServeMux(runtime.WithForwardResponseOption(httpResponseModifier))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterCalculatorServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":5656", mux)
}

func httpResponseModifier(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	// set http status code
	if vals := md.HeaderMD.Get("x-http-code"); len(vals) > 0 {
		code, err := strconv.Atoi(vals[0])
		if err != nil {
			return err
		}
		// delete the headers to not expose any grpc-metadata in http response
		delete(md.HeaderMD, "x-http-code")
		delete(w.Header(), "Grpc-Metadata-X-Http-Code")
		w.WriteHeader(code)
	}

	return nil
}

func emitDefault() runtime.ServeMuxOption {
	opt := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
		Marshaler: &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: false,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	})

	return opt
}
