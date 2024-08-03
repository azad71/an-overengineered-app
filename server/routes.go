package main

import (
	users "an-overengineered-social-media-app/modules/user"
	"an-overengineered-social-media-app/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.ErrorHandler())

	apiV1 := router.Group("/api/v1")
	authRoutes := apiV1.Group("/auth")
	{
		authRoutes.POST("/signup", users.SignupUser)
	}

	return router
}
