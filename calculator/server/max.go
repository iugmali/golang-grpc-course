package main

import (
	pb "github.com/iugmali/golang-grpc-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")
	var maximum int64 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		if number := req.Number; number > maximum {
			maximum = number
			err = stream.Send(&pb.MaxResponse{
				Result: maximum,
			})
			log.Printf("Sent %d to client\n", maximum)
			if err != nil {
				log.Fatalf("Error while sending data to client")
			}
		}
	}
}