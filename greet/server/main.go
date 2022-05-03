package main

import (
	pb "github.com/iugmali/golang-grpc-course/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)
	greet_server := grpc.NewServer()
	pb.RegisterGreetServiceServer(greet_server, &Server{})
	if err = greet_server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
