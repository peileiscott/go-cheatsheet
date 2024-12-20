package main

import (
	"context"
	"log"
	"time"

	pb "github.com/peileiscott/go-cheatsheet/grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Bob"})
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}
	log.Printf("Greeting: %s", response.GetMessage())
}
