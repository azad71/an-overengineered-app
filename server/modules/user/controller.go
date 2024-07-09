package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignupUser(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Successfully signed up",
	})
}
