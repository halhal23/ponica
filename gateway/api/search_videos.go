package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

func SearchVideos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		service, _ := ctx.Get("yts")
		yts := service.(*youtube.Service)
		query := ctx.Query("q")
		call := yts.Search.List([]string{"id,snippet"}).Q(query).MaxResults(3)

		pageToken := ctx.Query("pageToken")
		if len(pageToken) > 0 {
			call = call.PageToken(pageToken)
		}

		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Failed to call search youtube api: %v", err)
		}

		ctx.JSON(fasthttp.StatusOK, res)
	}
}
