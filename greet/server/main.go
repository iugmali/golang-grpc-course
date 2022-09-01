package main

import (
	pb "github.com/iugmali/golang-grpc-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	
	opts := []grpc.ServerOption{}
	tls := true
	
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	
	greet_server := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(greet_server, &Server{})
	if err = greet_server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
