package main

import (
	"an-overengineered-app/internal/config"
	"an-overengineered-app/internal/helpers"
	"an-overengineered-app/internal/logger"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	helpers.LoadEnv()
	config.SetupServerConfig()
	err := config.SetupDB()

	if err != nil {
		logger.PrintFatal(context.TODO(), "Failed to setup db connection", err)
	}

	gin.SetMode(config.AppConfig.RunMode)
}

func main() {
	routes := InitRouter()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validateBirthDate", helpers.IsDateBefore)
	}
	endPoint := fmt.Sprintf(":%d", config.AppConfig.HttpPort)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routes,
	}

	logger.PrintInfo(context.TODO(), fmt.Sprintf("Server is running at: %s:%d",
		config.AppConfig.AppUrl,
		config.AppConfig.HttpPort),
		nil)

	err := server.ListenAndServe()

	if err != nil {
		logger.PrintFatal(context.TODO(), "Failed to start server, error:", err)
	}
}
