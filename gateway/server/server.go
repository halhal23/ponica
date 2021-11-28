package main

import (
	"fmt"
	"log"
	"net/http"

	"ponica/gateway/gen/api"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello rest gateway")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Faild to connect user client: %v\n", err)
	}
	userClient := api.NewUserServiceClient(conn)

	e := gin.Default()
	e.GET("/check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "this is rest gateway.")
	})
	e.GET("/user", func(ctx *gin.Context) {
		res, err := userClient.ShowUser(ctx, &api.ShowUserRequest{Id: 40})
		if err != nil {
			log.Fatalf("Failed to showUser: %v\n", err)
		}
		ctx.JSON(http.StatusOK, res)
	})
	e.Run()
}
