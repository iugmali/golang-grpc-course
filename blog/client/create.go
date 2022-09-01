package main

import (
	"context"
	pb "github.com/iugmali/golang-grpc-course/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---creteBlog was invoked---")
	blog := &pb.Blog{
		AuthorId: "iugmali",
		Title: "My first blog",
		Content: "My first blog content",
	}
	
	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unespected error: %v\n", err)
	}
	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}
