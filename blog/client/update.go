package main

import (
	"context"
	pb "github.com/iugmali/golang-grpc-course/blog/proto"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id: id,
		AuthorId: "iugz",
		Title: "New Title",
		Content: "a new awesome Content",
	}
	
	_, err := c.UpdateBlog(context.Background(), newBlog)
	
	if err != nil {
		log.Fatalf("Error while updating: %v\n", err)
	}
	
	log.Println("Blog post updated")
}
