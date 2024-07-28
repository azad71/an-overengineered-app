package users

import (
	"an-overengineered-social-media-app/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignupUser(c *gin.Context) {

	var userData SignupBody
	var err error

	if err = c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":   "Something went wrong",
			"errorData": helpers.FormatValidationError(err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully signed up",
		"data":    userData,
	})
}
