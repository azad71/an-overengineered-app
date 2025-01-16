package users

import (
	"an-overengineered-app/internal/db"
	"an-overengineered-app/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup, dbInstance *db.DB) {
	// userController := NewUserController(dbInstance)
	// NewControllers(userController)

	authRoutes := rg.Group("/auth")
	authRoutes.POST("/signup", middleware.Validation[SignupBody](), Controller.SignupUser)
	authRoutes.POST("/signup/verify-otp", Controller.VerifySignupOTP)
}
