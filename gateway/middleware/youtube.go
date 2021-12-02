package middleware

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func YoutbeService() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := os.Getenv("API_KEY")

		c := context.Background()
		yts, err := youtube.NewService(c, option.WithAPIKey(apiKey))
		if err != nil {
			logrus.Fatalf("Faild to create youtube service: %v", err)
		}

		ctx.Set("yts", yts)

		ctx.Next()
	}
}
