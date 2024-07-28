package main

import (
	users "an-overengineered-social-media-app/modules/user"
	"an-overengineered-social-media-app/pkg/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	config.SetupServerConfig()
	config.SetupDB()
	gin.SetMode(config.AppConfig.RunMode)
}

func main() {
	routes := InitRouter()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validateBirthDate", users.ValidateBirthDate)
	}
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
