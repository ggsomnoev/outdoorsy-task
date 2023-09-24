package main

import (
	"simple-rentals-api/router"	
	"simple-rentals-api/controller"
	"simple-rentals-api/service"
	"simple-rentals-api/repository"

	"simple-rentals-api/helpers"
)

var (
	repo = repository.NewRentalsRepository()
	s 	 = service.NewRentalService(repo)
	c 	 = controller.NewRentalController(s)
	r 	 = router.NewGinRouter()
)

func init() {
	helpers.SetDefaultLogger("simple-rentals-api")

	repository.InitRentalsDbConnection()
}

func main() {
	r.GET("/rentals/:RENTAL_ID", c.HandleRentalRequest)
	r.GET("/rentals", c.HandleRentalsRequest)

	// Kubernetes extras about the status of the service
	r.GET("/metrics", c.HandleMetricsRequest)
	r.GET("/healthz", c.HandleHealthCheckRequest)

	r.SERVE(":8080")
}