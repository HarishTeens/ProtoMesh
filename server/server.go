package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/route_guide2/hello"
)

type helloServer struct {
	pb.UnimplementedHelloServer
}

func (s *helloServer) SayHello(ctx context.Context, helloReq *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", helloReq.GetName())
	return &pb.HelloResponse{Message: "Hello " + helloReq.GetName()}, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServer(grpcServer, &helloServer{})
	grpcServer.Serve(lis)
}