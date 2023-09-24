package service

import (
	"fmt"
	"errors"
	"strconv"
	"strings"
	"simple-rentals-api/entity"
	"simple-rentals-api/repository"

	"github.com/gin-gonic/gin"
)

type RentalService interface {
	GetRental(ctx *gin.Context) ([]entity.Rental, error)
	GetRentals(ctx *gin.Context) ([]entity.Rental, error)
}

type service struct{
	repo repository.RentalRepository
}

//NewRentalsRepository creates a new rentals repository
func NewRentalService(r repository.RentalRepository) RentalService {	
	return &service{repo: r}
}

//GetRental returns a single rental
func (s *service) GetRental(c *gin.Context) ([]entity.Rental, error) {	
	rental := c.Param("RENTAL_ID")
	rentalID, err := strconv.Atoi(rental)

	if err != nil {
		return []entity.Rental{}, errors.New(fmt.Sprintf("Could not parse rentalID %v", err.Error()))
	}

	return s.repo.GetRental(rentalID)
}

//GetRentals returns multiple rentals
func (s *service) GetRentals(c *gin.Context) ([]entity.Rental, error) {
	var queryParams entity.QueryParams

	if err := c.BindQuery(&queryParams); err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid query parameters %v", err.Error()))
	}

	if queryParams.Near != "" {
		var errF error
		queryParams.NearF, errF = parseCoordinates(queryParams.Near)

		if errF != nil {			
			return nil, errors.New(fmt.Sprintf("Could not parse coordinates %v", errF.Error()))
		}
	}

	return s.repo.GetRentals(queryParams)
}

func parseCoordinates(coordStr string) ([]float64, error) {

	parts := strings.Split(coordStr, ",")

	coordinates := make([]float64, len(parts))

	for i, part := range parts {
		coord, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return nil, err
		}
		coordinates[i] = coord
	}

	return coordinates, nil
}