package handlers

import (
	//"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Rental struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Type            string `json:"type"`
	Make            string `json:"make"`
	Model           string `json:"model"`
	Year            string `json:"year"`
	Length          string `json:"length"`
	Sleeps          string `json:"sleeps"`
	PrimaryImageURL string `json:"primary_image_url"`
	Price           struct {
		Day string `json:"day"`
	} `json:"price"`
	Location struct {
		City    string `json:"city"`
		State   string `json:"state"`
		Zip     string `json:"zip"`
		Country string `json:"country"`
		Lat     string `json:"lat"`
		Lng     string `json:"lng"`
	} `json:"location"`
	User struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"user"`
}

func HandleRental(c *gin.Context) {

	rental_id, _ := c.Params.Get("RENTAL_ID")

	c.JSON(http.StatusOK, rental_id)
}

func HandleRentals(c *gin.Context) {

	c.JSON(http.StatusOK, c.Request.URL.RawQuery)
}
