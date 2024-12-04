package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rateitorg/chatrooms/api/handler"
	"github.com/rateitorg/chatrooms/api/service"
)

func CreateRouter() *gin.Engine{
	// Create a gin router
	router := gin.Default()

	// Create handlers
	versionHandler := handler.VersionHandler{
		Service:&service.VersionService{},
	}

	// Define routes
	defineRoute(router, "version", "GET", versionHandler.GetVersion)

	return router
}

func defineRoute(router *gin.Engine, path string, method string, handler gin.HandlerFunc) {
	switch method {
	case "GET":
		router.GET("/"+path, handler)
	case "POST":
		router.POST("/"+path, handler)
	case "PUT":
		router.PUT("/"+path, handler)
	case "DELETE":
		router.DELETE("/"+path, handler)
	default:
		panic("unsupported method: " + method)
	}
}