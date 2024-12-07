package main

import (
	"github.com/joho/godotenv"
	"github.com/rateitorg/chatrooms/api/router"
	"github.com/rateitorg/chatrooms/config"
)

func init () {
	godotenv.Load()
}

// API entry point
func main () {
	cfg := config.GetConfig()
	app := router.CreateRouter()

	app.Run(":" + cfg.Port)
}