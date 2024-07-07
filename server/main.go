package main

import (
	"an-overengineered-social-media-app/pkg/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	config.SetupServerConfig()
	config.SetupDB()
	gin.SetMode(config.AppConfig.RunMode)
}

func main() {
	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"message": "Hello, world",
		})
	})

	err := app.Run(fmt.Sprintf(":%d", config.AppConfig.HttpPort))

	if err != nil {
		log.Panicf("Failed to start server, err: %v", err)
	}

}
