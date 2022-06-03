package main

import (
	"context"
	pb "github.com/iugmali/golang-grpc-course/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient)  {
	log.Println("doLongGreet was invoked")
	reqs := []*pb.GreetRequest {
		{FirstName: "Clement"},
		{FirstName: "Aline"},
		{FirstName: "Guilherme"},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}
	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}
	log.Printf("Long Greet: %s\n", res.Result)
}