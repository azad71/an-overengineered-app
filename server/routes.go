package main

import (
	"an-overengineered-app/internal/logger"
	"an-overengineered-app/internal/middleware"
	users "an-overengineered-app/modules/user"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	logInstance := logger.GetLogger()

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.HttpLogger(logInstance))
	router.Use(middleware.ErrorHandler())

	apiV1 := router.Group("/api/v1")
	authRoutes := apiV1.Group("/auth")
	{
		authRoutes.POST("/signup", middleware.Validation[users.SignupBody](), users.SignupUser)
		authRoutes.POST("/signup/verify-otp", users.VerifySignupOTP)
	}

	return router
}
