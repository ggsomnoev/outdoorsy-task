package main

import (
	"fmt"
	"testing"
	"simple-rentals-api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/appleboy/gofight/v2"
	"github.com/stretchr/testify/assert"
)

func initEngine() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.GET("rentals/:RENTAL_ID", func(c *gin.Context) {
		handlers.HandleRental(c)
	})

	r.GET("rentals", func(c *gin.Context) {
		handlers.HandleRentals(c)
	})

	return r
}

func TestRentals(t *testing.T) {
	r := gofight.New()

	r.GET("/rentals/12").
		SetDebug(true).
		SetHeader(gofight.H{
			"Host": "api.unnittest.com",
		}).
		Run(initEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, 200, r.Code)

			fmt.Println(r.Body.String())
			assert.Contains(t, r.Body.String(), "Kihei")
		})

	r.GET("/rentals?price_min=9000&price_max=9200").
		SetDebug(true).
		SetHeader(gofight.H{
			"Host": "api.unnittest.com",
		}).
		Run(initEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, 200, r.Code)

			fmt.Println(r.Body.String())
			assert.Contains(t, r.Body.String(), "Cumbria")
		})
}