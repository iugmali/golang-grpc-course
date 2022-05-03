package main

import (
	"context"
	pb "github.com/iugmali/golang-grpc-course/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was involved with %v\n", in)
	return &pb.GreetResponse{
		Result: in.FirstName + " is the best developer I've known.",
	}, nil
}
