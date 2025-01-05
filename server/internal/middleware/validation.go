package middleware

import (
	"an-overengineered-app/internal/httpResponse"
	"an-overengineered-app/internal/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validation[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger.Info(ctx, "Executing validation middleware", nil)
		var payload T

		if err := ctx.ShouldBindJSON(&payload); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				logger.ErrorStack(ctx, "Failed to validate req body", err)
				ctx.Error(httpResponse.ValidationError("", errs))
				ctx.Abort()
				return
			}
			logger.ErrorStack(ctx, "Error binding request body to struct", err)
			ctx.Error(httpResponse.InternalServerError(""))
			ctx.Abort()
			return
		}

		logger.Info(ctx, "Setting validated request body to context", nil)
		ctx.Set("body", payload)

		ctx.Next()
	}
}
