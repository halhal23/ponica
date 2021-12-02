package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

type VideoResponse struct {
	VideoList *youtube.VideoListResponse `json:"video_list"`
}

func GetVideo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		service, _ := ctx.Get("yts")
		yts := service.(*youtube.Service)

		videoId := ctx.Params.ByName("id")

		call := yts.Videos.List([]string{"id", "snippet"}).Id(videoId)

		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Faild to get youtube api calling: %v", err)
		}

		v := VideoResponse{
			VideoList: res,
		}

		ctx.JSON(fasthttp.StatusOK, v)
	}
}
