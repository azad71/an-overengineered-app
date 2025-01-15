package main

import (
	"an-overengineered-app/internal/config"
	"an-overengineered-app/internal/db"
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

	gin.SetMode(config.AppConfig.RunMode)
}

func main() {
	dbConn, err := db.ConnectDB()

	if err != nil {
		logger.Fatal(context.Background(), "Failed to connect to Database. Closing server...", err)
	}

	routes := InitRouter(dbConn)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validateBirthDate", helpers.IsDateBefore)
	}
	endPoint := fmt.Sprintf(":%d", config.AppConfig.HttpPort)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routes,
	}

	logger.Info(context.Background(), fmt.Sprintf("Server is running at: %s:%d",
		config.AppConfig.AppUrl,
		config.AppConfig.HttpPort),
		nil)

	err = server.ListenAndServe()

	if err != nil {
		logger.Fatal(context.Background(), "Failed to start server, error:", err)
	}
}
