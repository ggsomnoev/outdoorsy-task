package controller

import (
	"log/slog"
	"net/http"	
	"simple-rentals-api/service"

	"github.com/gin-gonic/gin"
)


type RentalController interface{
	HandleRentalRequest(c *gin.Context)
	HandleRentalsRequest(c *gin.Context)
	HandleMetricsRequest(c *gin.Context)
	HandleHealthCheckRequest(c *gin.Context)
}

type controller struct{
	rentalService service.RentalService
}

func NewRentalController(s service.RentalService) RentalController {
	return &controller{rentalService : s}
}

func (r *controller) HandleRentalRequest(c *gin.Context) {

	result, err := r.rentalService.GetRental(c)

	if err != nil {
		slog.Error("An error occured trying to get rental data", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (r *controller) HandleRentalsRequest(c *gin.Context) {

	results, err := r.rentalService.GetRentals(c)

	if err != nil {
		slog.Error("An error occured trying to get rental data", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}