package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/route_guide2/hello"
)


func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewHelloClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	helloRequest := &pb.HelloRequest{Name: "Harish"}

	log.Printf("Sending request to server: %s", helloRequest.Name)
	message, err := client.SayHello(ctx, helloRequest)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response from server: %s", message.Message)


}