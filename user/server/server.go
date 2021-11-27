package main

import (
	"fmt"
	"log"
	"net"
	"ponica/user/gen/api"
	"ponica/user/service"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello ponica")

	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Faild to listen: %v\n", err)
	}

	server := grpc.NewServer()
	service := service.NewUserService()
	api.RegisterUserServiceServer(server, &service)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Faild to server: %v\n", err)
	}
}
