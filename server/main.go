package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/penril0326/hello_grpc/proto/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedCalculatorServiceServer
}

func main() {
	fmt.Println("Starting to serve gRPC...")

	l, err := net.Listen("tcp", "localhost:5555")
	if err != nil {
		log.Fatalf("Failed to listen: %s\n", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("Failed to serve gRPC: %s\n", err.Error())
	}
}

func (s *Server) Sum(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	fmt.Printf("Receive...: %v\n", req)
	// a := req.GetA()
	// b := req.GetB()

	// resp := pb.CalculatorResponse{
	// 	Result: a + b,
	// }

	return nil, status.Error(codes.Aborted, "asda")
}

func (s *Server) Deletetest(ctx context.Context, req *pb.DeleteTest) (*emptypb.Empty, error) {
	grpc.SetHeader(ctx, metadata.Pairs("X-Http-Code", "211"))
	return &emptypb.Empty{}, nil
}
