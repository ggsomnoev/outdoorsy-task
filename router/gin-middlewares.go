package router

import (
	"fmt"
	"time"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func GinFormatMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
			c.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		slog.Info("req_info",
			slog.Int("status_code", c.Writer.Status()),
			slog.String("client_ip", c.ClientIP()),
			slog.String("req_method", c.Request.Method),
			slog.String("req_uri", c.Request.RequestURI),
			slog.Duration("latency_time", latencyTime),
		)
	}
}

func CustomRecoveryWithWriter() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("gin_recovery",
					slog.String("client_ip", c.ClientIP()),
					slog.String("req_method", c.Request.Method),
					slog.String("headers", fmt.Sprint(c.Request.Header)),
					slog.String("url_params", c.Request.URL.RawQuery),
					slog.String("panic_message", fmt.Sprint(err)),
				)
			}
		}()

		c.Next()
	}
}