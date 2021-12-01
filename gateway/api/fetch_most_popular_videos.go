package api

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func FetchMostPopularVideos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := os.Getenv("API_KEY")

		c := context.Background()
		yts, err := youtube.NewService(c, option.WithAPIKey(apiKey))
		if err != nil {
			logrus.Fatalf("Faild to create youtube service: %v", err)
		}
		call := yts.Videos.List([]string{
			"id",
			"snippet",
		}).Chart("mostPopular").MaxResults(3)
		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Faild to youtube api call: %v", err)
		}

		ctx.JSON(fasthttp.StatusOK, res)
	}
}
