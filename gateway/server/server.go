package main

import (
	"fmt"
	"log"

	api_pb "ponica/gateway/gen/api"
	"ponica/gateway/handler"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello rest gateway")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Faild to connect user client: %v\n", err)
	}
	userClient := api_pb.NewUserServiceClient(conn)

	hdl := handler.NewRestHandler(userClient)

	e := gin.Default()
	e.GET("/check", hdl.HeathCheck)
	e.GET("/user", hdl.ShowUser)
	e.Run()
}
