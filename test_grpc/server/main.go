package main

import (
	"context"
	"log"
	"net"
	"strconv"

	pb "test_grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(in *pb.HelloRequest, stream pb.Greeter_SayHelloAgainServer) error {
	for i := 0; i < 10; i++ {
		reply := &pb.HelloReply{
			Message: "Hello" + in.GetName() + ", " + strconv.Itoa(i+1),
		}
		stream.Send(reply)
	}
	return nil
}
