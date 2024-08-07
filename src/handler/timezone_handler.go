package handler

import (
	"go-api-rest/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /timezone

// PingExample godoc
// @Summary listing of all timezones.
// @Schemes
// @Description listing of all timezones from the backend.
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {array} string Data "List of timezones"
// @Failure 400 {object} model.Exception
// @Failure 409 {object} model.Exception
// @Failure 500 {object} model.Exception
// @Router /timezone [get]
func TimezoneHandler(c *gin.Context) {

	data, error := service.TimezoneService()
	if error.Status != 0 {
		c.JSON(error.Status, error)
		return
	}
	c.JSON(http.StatusOK, data)
}
