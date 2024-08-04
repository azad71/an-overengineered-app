package main

import (
	"an-overengineered-social-media-app/internal/config"
	"an-overengineered-social-media-app/internal/helpers"
	users "an-overengineered-social-media-app/modules/user"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	helpers.LoadEnv()
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

	log.Printf("[info] Server is running at: %s:%d", config.AppConfig.AppUrl, config.AppConfig.HttpPort)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to start server, error: %v", err)
	}
}
