package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

func FetchMostPopularVideos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		service, _ := ctx.Get("yts")
		yts := service.(*youtube.Service)

		call := yts.Videos.List([]string{
			"id",
			"snippet",
		}).Chart("mostPopular").MaxResults(3)

		pageToken := ctx.Query("pageToken")
		if len(pageToken) > 0 {
			call = call.PageToken(pageToken)
		}

		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Faild to youtube api call: %v", err)
		}

		ctx.JSON(fasthttp.StatusOK, res)
	}
}
