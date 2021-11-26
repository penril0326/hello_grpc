package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	pb "github.com/penril0326/hello_grpc/proto/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/anypb"
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
	resp := &pb.CalculatorResponse{
		Result: req.GetA() + req.GetB(),
	}

	any := &pb.TestAny{
		Str1: "test1",
		Int1: 100,
		Ints: []int64{1, 2, 3, 4, 5},
	}

	anyproto, err := anypb.New(any)
	if err != nil {
		return nil, errors.New("any marshal failed")
	}

	resp.Custom = anyproto

	return resp, nil
}

func (s *Server) Deletetest(ctx context.Context, req *pb.DeleteTest) (*emptypb.Empty, error) {
	grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "204"))
	return &emptypb.Empty{}, nil
}
