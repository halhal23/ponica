package api

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

func FetchMostPopularVideos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := os.Getenv("API_KEY")

		c := context.Background()
		youtube.NewService(c)
		ctx.JSON(fasthttp.StatusOK, "successfully fetch most popular videos"+apiKey)
	}
}
