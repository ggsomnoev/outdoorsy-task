package main

import (
	"simple-rentals-api/handlers"
	"simple-rentals-api/helpers"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	helpers.SetDefaultLogger("simple-rentals-api")
}

func main() {
	r := gin.New()
	r.Use(helpers.GinFormatMiddleware())
	r.Use(helpers.CustomRecoveryWithWriter())

	r.GET("/rentals/:RENTAL_ID", handlers.HandleRental)
	r.GET("/rentals", handlers.HandleRentals)

	// Kubernetes extras
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/healthz", helpers.HandleHealthCheckRequest)

	r.Run(":8080")
}
