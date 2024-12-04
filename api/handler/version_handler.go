package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rateitorg/chatrooms/api/service"
)

type VersionHandler struct {
	Service *service.VersionService
}

func (vh * VersionHandler) GetVersion(ctx *gin.Context) {
	response := vh.Service.GetVersionService()

	ctx.JSON(response.Code, response.Data)
}
