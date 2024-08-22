package middleware

import (
	"context"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	logInstance "an-overengineered-social-media-app/internal/logger"
)

func DefaultStructuredLogger() gin.HandlerFunc {
	return HttpLogger(&log.Logger)
}

func HttpLogger(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		log := *logInstance.GetLogger()
		requestId := xid.New().String()

		log = log.With().Str("requestId", requestId).Logger()

		ctx := log.WithContext(c.Request.Context())
		ctx = context.WithValue(ctx, logInstance.RequestIdKey, requestId)
		c.Request = c.Request.WithContext(ctx)

		c.Header("requestId", requestId)

		latency := time.Since(start).String()
		statusCode := c.Writer.Status()
		buildInfo, _ := debug.ReadBuildInfo()

		c.Next()

		defer func() {
			if panicValue := recover(); panicValue != nil {
				statusCode = http.StatusInternalServerError
				panic(panicValue)
			}

			log.Info().
				Interface("header", c.Request.Header).
				Str("user_agent", c.Request.UserAgent()).
				Str("method", c.Request.Method).
				Int("status_code", statusCode).
				Int64("body_size", c.Request.ContentLength).
				Str("url", c.Request.URL.RequestURI()).
				Str("remote_address", c.Request.RemoteAddr).
				Str("latency", latency).
				Str("go_version", buildInfo.GoVersion).
				Send()
		}()

	}
}
