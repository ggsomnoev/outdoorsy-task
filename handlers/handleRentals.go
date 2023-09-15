package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"log/slog"
	"net/http"
	"simple-rentals-api/helpers"

	"github.com/gin-gonic/gin"
)

func HandleRental(c *gin.Context) {

	rental := c.Param("RENTAL_ID")
	rentalID, err := strconv.Atoi(rental)

	if err != nil {
		slog.Error("Could not parse rentalID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rentalID"})
		return
	}

	query := fmt.Sprintf("%v AND r.id = %v", helpers.BaseQuery, rentalID)

	results, err := helpers.ExecuteDatabaseQuery(query)

	if err != nil {
		slog.Error("Failed to fetch data from the DB", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from the DB"})
		return
	}

	c.JSON(http.StatusOK, results)
}

func HandleRentals(c *gin.Context) {
	var queryParams helpers.QueryParams

	if err := c.BindQuery(&queryParams); err != nil {
		slog.Error("Invalid query parameters", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	if queryParams.Near != "" {
		var errF error
		queryParams.NearF, errF = parseCoordinates(queryParams.Near)

		if errF != nil {
			slog.Error("Could not parse coordinates", slog.String("error", errF.Error()))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
			return
		}
	}

	query := helpers.GenerateSQLQuery(queryParams)

	results, err := helpers.ExecuteDatabaseQuery(query)

	if err != nil {
		slog.Error("Failed to fetch data from the DB", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from the DB"})
		return
	}

	c.JSON(http.StatusOK, results)
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
