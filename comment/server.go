package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"comment/pb"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("hello comment server: 50052")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Faild to listen: %v\n", err)
	}

	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Faild to log: %v\n", err)
	}
	grpc_zap.ReplaceGrpcLogger(zapLogger)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_zap.UnaryServerInterceptor(zapLogger),
		),
	)
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
