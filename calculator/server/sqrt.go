package main

import (
	"context"
	"fmt"
	"github.com/iugmali/golang-grpc-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt was invoked with: %v\n", in)
	number := in.Number
	
	if number < 0 {
		return nil, status.Errorf(
				codes.InvalidArgument,
				fmt.Sprintf("Received a negative number: %d\n", number)
			)
	}
	return &pb.SqrtResponse {
		Result: math.Sqrt(float64(number)),
	}, nil
}
