package middleware

import (
	"an-overengineered-app/internal/httpResponse"
	"an-overengineered-app/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		requestId, ok := c.Request.Context().Value(logger.RequestIdKey).(string)

		if !ok {
			requestId = ""
		}

		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case httpResponse.CustomError:
				e.RequestID = requestId
				c.AbortWithStatusJSON(e.StatusCode, e)
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message":   "Service Unavailable",
					"requestId": requestId,
				})
			}
		}
	}
}
