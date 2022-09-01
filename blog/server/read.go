package main

import (
	"context"
	pb "github.com/iugmali/golang-grpc-course/blog/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked with: %v\n", in)
	oid, err := primitive.ObjectIDFromHex(in.Id)
	
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
			)
	}
	
	data := &BlogItem{}
	filter := bson.M{"_id": oid}
	res := collection.FindOne(ctx, filter)
	
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog post with the ID provided",
			)
	}
	
	return documentToBlog(data), nil
}
