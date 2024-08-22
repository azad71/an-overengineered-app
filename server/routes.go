package main

import (
	"an-overengineered-app/internal/middleware"
	users "an-overengineered-app/modules/user"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.DefaultStructuredLogger())
	router.Use(middleware.ErrorHandler())

	apiV1 := router.Group("/api/v1")
	authRoutes := apiV1.Group("/auth")
	{
		authRoutes.POST("/signup", users.SignupUser)
	}

	return router
}
