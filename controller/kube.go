package controller

import (
	"net/http"
	
	rep "simple-rentals-api/repository"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type HealthStatus struct {
	Status  bool
	Entries []HealthEntry
}

type HealthEntry struct {
	Status        bool   `json:"status"`
	ErrorMessage  string `json:"error_message"`
	ComponentName string `json:"component_name"`
}


func (r *controller) HandleHealthCheckRequest(c *gin.Context) {

	dd := HealthStatus{}
	dd.Status = true

	if err := rep.RentalsDB.Ping(); err != nil {
		dd.Status = false
		dd.Entries = append(dd.Entries, HealthEntry{
			ComponentName: "DB",
			Status:        false,
			ErrorMessage:  err.Error(),
		})
	} else {
		dd.Entries = append(dd.Entries, HealthEntry{
			ComponentName: "DB",
			Status:        true,
		})
	}

	c.JSON(http.StatusOK, dd)
}

func (r *controller) HandleMetricsRequest(c *gin.Context) {
	gin.WrapH(promhttp.Handler())
}