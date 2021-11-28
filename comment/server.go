package main

import (
	"fmt"
	"log"
	"net"

	"comment/pb"
	"comment/service"

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
	svc := service.NewCommentService()
	pb.RegisterCommentServiceServer(server, svc)

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Faild to serve: %v\n", err)
	}
}
