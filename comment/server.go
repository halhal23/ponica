package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"comment/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("hello comment server: 50052")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Faild to listen: %v\n", err)
	}

	server := grpc.NewServer()
	pb.RegisterCommentServiceServer(server, &CommentService{})

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Faild to serve: %v\n", err)
	}
}

type CommentService struct{}

func (service *CommentService) ShowComment(ctx context.Context, req *pb.ShowCommentRequest) (*pb.ShowCommentResponse, error) {
	return &pb.ShowCommentResponse{
		CommentId: 234,
	}, nil
}
