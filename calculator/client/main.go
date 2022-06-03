package main

import (
	pb "github.com/iugmali/golang-grpc-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()
	calculatorClient := pb.NewCalculatorServiceClient(conn)

	//doSimpleSum(calculator_client)
	//doPrimes(calculatorClient)
	//doAvg(calculatorClient)
	doMax(calculatorClient)
}
