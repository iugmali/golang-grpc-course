package main

import (
	"context"
	"fmt"
	"github.com/iugmali/golang-grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type server struct {}

func String(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
func (*server) Sum(ctx context.Context, req *calculatorpb.TwoNumbersRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Sum fnc invoked with %v\n", req)
	numberOne := req.GetNumberOne()
	numberTwo := req.GetNumberTwo()
	result :=  String(numberOne + numberTwo)
	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}
	return res, nil
}
func (*server) Multiply(ctx context.Context, req *calculatorpb.TwoNumbersRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Multiply fnc invoked with %v\n", req)
	numberOne := req.GetNumberOne()
	numberTwo := req.GetNumberTwo()
	result :=  String(numberOne * numberTwo)
	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}
	return res, nil
}
func (*server) Divide(ctx context.Context, req *calculatorpb.TwoNumbersRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Divide fnc invoked with %v\n", req)
	numberOne := req.GetNumberOne()
	numberTwo := req.GetNumberTwo()
	result := String(numberOne / numberTwo)
	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}
	return res, nil
}
func (*server) PrimeNumberDecomposition(req *calculatorpb.OneNumberRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("PrimeNumberDecomposition fnc invoked with %v\n", req)
	number := req.GetNumber()
	divisor := int64(2)
	for number > 1 {
		if number % divisor == 0 {
			stream.Send(&calculatorpb.CalculatorResponse{
				Result: strconv.Itoa(int(divisor)) + "\n",
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor increased to %v \n", divisor)
		}
	}
	return nil
}
func main()  {
	fmt.Println("Server Starting...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
