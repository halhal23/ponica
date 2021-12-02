package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

func FetchRelatedVideos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		service, _ := ctx.Get("yts")
		yts := service.(*youtube.Service)

		videoId := ctx.Params.ByName("id")

		call := yts.Search.List([]string{"id", "snippet"}).RelatedToVideoId(videoId).Type("video").MaxResults(3)

		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Failed to call related videos: %v", err)
		}

		ctx.JSON(fasthttp.StatusOK, res)
	}
}
