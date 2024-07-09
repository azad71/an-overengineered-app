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
	routes := InitRouter()
	endPoint := fmt.Sprintf(":%d", config.AppConfig.HttpPort)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routes,
	}

	log.Printf("[info] Server is running at: %s", config.AppConfig.AppUrl)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to start server, error: %v", err)
	}
}
