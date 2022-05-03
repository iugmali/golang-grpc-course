package main

import (
	"context"
	pb "github.com/iugmali/golang-grpc-course/greet/proto"
	"log"
)

func doGreet(greet_client pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := greet_client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "iugmali",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s", res.Result)
}
