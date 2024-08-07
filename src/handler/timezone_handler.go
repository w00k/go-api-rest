package handler

import (
	"go-api-rest/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TimezoneHandler(c *gin.Context) {

	data, error := service.TimezoneService()
	if error.Status != 0 {
		c.JSON(error.Status, error)
		return
	}
	c.JSON(http.StatusOK, data)
}
