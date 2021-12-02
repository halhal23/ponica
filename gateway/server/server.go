package main

import (
	"fmt"
	"log"
	"time"

	"ponica/gateway/api"
	api_pb "ponica/gateway/gen/api"
	"ponica/gateway/handler"
	"ponica/gateway/middleware"

	"github.com/gin-contrib/cors"
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
	e.Use(middleware.YoutbeService())
	e.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PATCH", "OPTIONS"},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
	e.GET("/check", hdl.HeathCheck)
	e.GET("/user", hdl.ShowUser)
	g := e.Group("/api")
	{
		g.GET("/populars", api.FetchMostPopularVideos())
		g.GET("/video/:id", api.GetVideo())
		g.GET("/related/:id", api.FetchRelatedVideos())
	}
	e.Run()
}
