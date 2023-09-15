package main

import (
	"log/slog"
	"simple-rentals-api/helpers"
	"simple-rentals-api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	helpers.SetDefaultLogger("simple-rentals-api")

	if err := helpers.InitRentalsDbConnection(); err != nil {
		slog.Error("An error occured trying to connect to DB", slog.String("Error", err.Error()))
		defer helpers.RentalsDB.Close()
	}
}

func main() {

	// The idea here is to convert the default gin logs to JSON format and add additional info, if we need/want.
	r := gin.New()
	r.Use(helpers.GinFormatMiddleware())
	r.Use(helpers.CustomRecoveryWithWriter())

	r.GET("/rentals/:RENTAL_ID", handlers.HandleRental)
	r.GET("/rentals", handlers.HandleRentals)

	// Kubernetes extras about the status of the service
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/healthz", helpers.HandleHealthCheckRequest)

	r.Run(":8080")
}

// TODO: add some unit tests
