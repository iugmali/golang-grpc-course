package main

import (
	"context"
	pb "github.com/iugmali/golang-grpc-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func doSqrt(c pb.CalculatorServiceClient) {
	log.Println("doSqrt was invoked")
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number")
			}
			return
		} else {
			log.Fatalf("non gRPC error: %v\n", err)
		}
	}
	log.Printf("Sqrt: %f\n", res.Result)
}
