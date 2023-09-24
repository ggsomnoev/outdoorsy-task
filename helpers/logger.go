package helpers

import (
	"os"
	"log/slog"
)


// Note: slog requires GO Version 1.21
func SetDefaultLogger(serviceName string) {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		// AddSource: true, // TMI
	}).WithAttrs([]slog.Attr{slog.String("service", serviceName)})
	
	logger := slog.New(logHandler)

	slog.SetDefault(logger)
}