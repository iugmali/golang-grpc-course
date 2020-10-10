package main

import (
	"context"
	"fmt"
	"github.com/iugmali/golang-grpc-course/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main()  {
	fmt.Println("Client Starting...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("nao foi possivel conectar: %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Client criado: %f", c)
	doUnary(c)
	doStream(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Guilherme",
			LastName: "Almeida",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("erro ao chamar Greet RPC: %v", err)
	}
	log.Printf("Resposta de Greet: %v", res.Result)
}

func doStream(c greetpb.GreetServiceClient) {
	fmt.Println("Starting stream RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Guilherme",
			LastName: "Almeida",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("erro ao chamar GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			fmt.Println("End Of Streaming")
			break
		}
		if err != nil {
			log.Fatalf("erro na stream RPC: %v", err)
		}
		log.Printf("Resposta de GreetManyTimes: %v", msg.GetResult())
	}
}

