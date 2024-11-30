package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rateitorg/chatrooms/handler"
	"github.com/rateitorg/chatrooms/service"
)

// Entry point of the application
func main() {
	// Set up gin engine
	engine := setUpEngine()

	// Create a new hub
	hub := service.NewHub()

	// Start the hub
	go hub.Run()

	// Set up routes
	// Default route
	engine.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// WebSocket route
	engine.GET("/ws", func(ctx *gin.Context) {
		handler.WebSocketHandler(hub, ctx.Writer, ctx.Request)
	})

	// Start engine
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}

// Create a new gin engine
func setUpEngine() *gin.Engine {
	engine := gin.Default()
	err := engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}
	return engine
}
