package main

import (
	pb "github.com/iugmali/golang-grpc-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	//"time"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}
	
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		
		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}
		
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	
	conn, err := grpc.Dial(addr, opts...)
	
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	
	defer conn.Close()
	greet_client := pb.NewGreetServiceClient(conn)

	doGreet(greet_client)
	//doGreetManyTimes(greet_client)
	//doLongGreet(greet_client)
	//doGreetEveryone(greet_client)
	//doGreetWithDeadline(greet_client, 1 * time.Second)
}
