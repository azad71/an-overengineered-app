package middleware

import (
	"an-overengineered-social-media-app/internal/httpError"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case httpError.CustomError:
				c.AbortWithStatusJSON(e.StatusCode, e)
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Service Unavailable",
				})
			}
		}
	}
}
