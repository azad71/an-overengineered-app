package httpResponse

import (
	"an-overengineered-app/internal/logger"
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

func getRequestId(ctx *gin.Context) string {
	requestId, ok := ctx.Request.
		Context().
		Value(logger.RequestIdKey).(string)

	if !ok {
		return ""
	}

	return requestId
}

func Created(ctx *gin.Context, msg string, data gin.H) {

	requestId := getRequestId(ctx)

	ctx.JSON(http.StatusCreated, buildResponseObj(msg, data, requestId))
}

func Success(ctx *gin.Context, msg string, data gin.H) {
	requestId := getRequestId(ctx)

	ctx.JSON(http.StatusOK, buildResponseObj(msg, data, requestId))
}
