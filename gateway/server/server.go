package main

import (
	"fmt"
	"log"

	"ponica/gateway/api"
	api_pb "ponica/gateway/gen/api"
	"ponica/gateway/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Faild to load .env by godotenv: %v\n", err)
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

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
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
	}
	e.Run()
}
