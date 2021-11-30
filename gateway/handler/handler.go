package handler

import (
	"net/http"
	"ponica/gateway/gen/api"

	"github.com/gin-gonic/gin"
)

type RestHandler struct {
	UserClient api.UserServiceClient
}

func NewRestHandler(userClient api.UserServiceClient) RestHandler {
	return RestHandler{
		UserClient: userClient,
	}
}

func (handler *RestHandler) HeathCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Status Ok, keep going man.")
}

func (handler *RestHandler) ShowUser(ctx *gin.Context) {
	res, err := handler.UserClient.ShowUser(ctx, &api.ShowUserRequest{Id: 20})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, res)
}
