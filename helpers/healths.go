package helpers

import (
	"net/http"
	"github.com/gin-gonic/gin"
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


func HandleHealthCheckRequest(c *gin.Context) {

	dd := HealthStatus{}
	dd.Status = true

	if err := RentalsDB.Ping(); err != nil {
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