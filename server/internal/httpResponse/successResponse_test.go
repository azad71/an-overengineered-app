package httpResponse

import (
	"an-overengineered-app/internal/logger"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var requestId = "6dMkvTsUY8PN3eQ"

func TestBuildResponseObj(t *testing.T) {
	t.Run("Should return formatted response", func(t *testing.T) {

		msg := "User logged in successfully"
		data := gin.H{
			"userId":   1,
			"username": "testuser",
		}

		result := buildResponseObj(msg, data, requestId)

		expectedOutput := gin.H{
			"message":   msg,
			"data":      data,
			"requestId": requestId,
		}

		assert.Equal(t, expectedOutput, result)
	})
}

func getCreatedParams(requestId string) (*httptest.ResponseRecorder, *gin.Context, string, gin.H) {
	gin.SetMode(gin.TestMode)

	// Create a new response recorder and context
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	// Create a mock HTTP request and associate it with the context
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	ctx.Request = req

	if requestId != "" {
		// Set up a request ID in the context
		ctx.Request = ctx.Request.
			WithContext(context.WithValue(ctx.Request.Context(), logger.RequestIdKey, requestId))
	}

	msg := "Resource created successfully"
	data := gin.H{"id": 1}

	return w, ctx, msg, data
}

func TestCreated(t *testing.T) {
	t.Run("Should return http 201", func(t *testing.T) {
		w, ctx, msg, data := getCreatedParams(requestId)
		Created(ctx, msg, data)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("Should return defined json response", func(t *testing.T) {
		w, ctx, msg, data := getCreatedParams(requestId)
		Created(ctx, msg, data)
		expectedResponse := `{"message":"Resource created successfully","data":{"id":1},"requestId":"6dMkvTsUY8PN3eQ"}`

		assert.JSONEq(t, expectedResponse, w.Body.String())
	})

	t.Run("Should return empty requestId if not attached to context", func(t *testing.T) {
		w, ctx, msg, data := getCreatedParams("")
		Created(ctx, msg, data)
		expectedResponse := `{"message":"Resource created successfully","data":{"id":1},"requestId":""}`

		assert.JSONEq(t, expectedResponse, w.Body.String())

	})

}
