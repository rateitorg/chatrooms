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
	config.Init() // Add all .env values to app environment
	app := router.CreateRouter()

	app.Run(":" + config.PORT)
}