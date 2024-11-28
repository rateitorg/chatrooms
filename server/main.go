package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := setUpEngine()

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