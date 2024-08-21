package httpResponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func buildResponseObj(msg string, data gin.H, requestId string) gin.H {
	return gin.H{
		"message":   msg,
		"data":      data,
		"requestId": requestId,
	}
}

func Created(ctx *gin.Context, msg string, data gin.H) {
	requestId := ctx.Request.Context().Value("requestId").(string)

	ctx.JSON(http.StatusCreated, buildResponseObj(msg, data, requestId))
}
