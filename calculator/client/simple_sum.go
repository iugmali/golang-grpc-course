package main

import (
	"context"
	pb "github.com/iugmali/golang-grpc-course/calculator/proto"
	"log"
)

func doSimpleSum(c pb.CalculatorServiceClient) {
	log.Printf("doSimpleSum was invoked")
	res, err := c.SimpleSum(context.Background(), &pb.SimpleSumRequest{
		FirstNumber:  1,
		SecondNumber: 1,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}
