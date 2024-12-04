package config

import "os"

var PORT string
var LOG_LEVEL string
var API_VERSION string
var API_NAME string

func Init() {
	PORT = os.Getenv("PORT")
	LOG_LEVEL = os.Getenv("LOG_LEVEL")
	API_VERSION = os.Getenv("API_VERSION")
	API_NAME = os.Getenv("API_NAME")
}
