package router

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

var (
	ginDispatcher = gin.New()
)

type GinRouter interface {
	GET(uri string, f gin.HandlerFunc)
	SERVE(port string)
}

type ginRouter struct{}

func NewGinRouter() GinRouter {	
	ginDispatcher.Use(GinFormatMiddleware())
	ginDispatcher.Use(CustomRecoveryWithWriter())
	return &ginRouter{}
}

func (*ginRouter) GET(uri string, f gin.HandlerFunc) {
	ginDispatcher.GET(uri, f)
}

func (*ginRouter) SERVE(port string) {
	slog.Info("GIN server running on...", slog.String("port", port))
	ginDispatcher.Run(port)
}