package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "hello_grpc_gateway/proto/test"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

type Server struct {
	pb.UnimplementedTestServiceServer
}

func main() {
	fmt.Println("Starting to serve gRPC...")

	l, err := net.Listen("tcp", "localhost:5555")
	if err != nil {
		log.Fatalf("Failed to listen: %s\n", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTestServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("Failed to serve gRPC: %s\n", err.Error())
	}
}

func (s *Server) Test(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	allow := []string{
		"b_i",
		"n",
	}

	allowMask, _ := fieldmaskpb.New(&pb.TestRequest{}, allow...)
	fmt.Println("allow = ", allowMask.GetPaths())
	intersect := fieldmaskpb.Intersect(allowMask, req.GetFields())
	intersect.Normalize()

	valid := intersect.GetPaths()

	fmt.Println(valid)

	return &pb.TestResponse{
		Code:   200,
		Result: 10,
	}, nil
}
