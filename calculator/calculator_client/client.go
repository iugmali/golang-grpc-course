package main

import (
	"context"
	"fmt"
	"github.com/iugmali/golang-grpc-course/calculator/calculatorpb"
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
	c := calculatorpb.NewCalculatorServiceClient(cc)
	//fmt.Printf("Client criado: %f", c)
	doUnary(c)
	doStream(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting unary RPC...")
	req := &calculatorpb.TwoNumbersRequest{
		NumberOne: 45,
		NumberTwo: 10,
	}
	resSum, errSum := c.Sum(context.Background(), req)
	if errSum != nil {
		log.Fatalf("erro ao chamar Sum RPC: %v", errSum)
	}
	log.Printf("Resposta de Sum: %v", resSum.Result)

	resMultiply, errMultiply := c.Multiply(context.Background(), req)
	if errMultiply != nil {
		log.Fatalf("erro ao chamar Multiply RPC: %v", errMultiply)
	}
	log.Printf("Resposta de Multiply: %v", resMultiply.Result)

	resDivide, errDivide := c.Divide(context.Background(), req)
	if errDivide != nil {
		log.Fatalf("erro ao chamar Divide RPC: %v", errDivide)
	}
	log.Printf("Resposta de Divide: %v", resDivide.Result)
}

func doStream(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting stream RPC...")
	req := &calculatorpb.OneNumberRequest{
		Number: 63551487,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("erro ao chamar PrimeDecomposition RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("erro no stream PrimeDecomposition RPC: %v", err)
		}
		fmt.Print(res.Result)
	}

}

