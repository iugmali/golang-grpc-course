package main

import (
	"context"
	pb "github.com/iugmali/golang-grpc-course/calculator/proto"
	"log"
)

func (s *Server) SimpleSum(ctx context.Context, in *pb.SimpleSumRequest) (*pb.SimpleSumResponse, error) {
	log.Printf("SimpleSum function was invoked with %v\n", in)
	return &pb.SimpleSumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
