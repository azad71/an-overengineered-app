package middleware

import (
	"an-overengineered-app/internal/httpResponse"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case httpResponse.CustomError:
				c.AbortWithStatusJSON(e.StatusCode, e)
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Service Unavailable",
				})
			}
		}
	}
}
