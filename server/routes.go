package main

import (
	"an-overengineered-app/internal/db"
	"an-overengineered-app/internal/logger"
	"an-overengineered-app/internal/middleware"
	users "an-overengineered-app/modules/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(dbInstance *db.DB) *gin.Engine {
	logInstance := logger.GetLogger()

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.HttpLogger(logInstance))
	router.Use(middleware.ErrorHandler())

	apiV1 := router.Group("/api/v1")

	users.InitRoutes(apiV1, dbInstance)

	return router
}
