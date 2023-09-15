package helpers

import (
	"os"
	"fmt"
	"time"
	"log/slog"

	"github.com/gin-gonic/gin"
)


// Note: slog requires GO Version 1.21
func SetDefaultLogger(serviceName string) {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		// AddSource: true, // TMI
	}).WithAttrs([]slog.Attr{slog.String("service", serviceName)})
	
	logger := slog.New(logHandler)

	slog.SetDefault(logger)
}

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
