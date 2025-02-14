package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHealth godoc
// @Summary Get the health status of the server
// @Description Returns the server health status as a simple "OK" message.
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Failure 500 {string} string "Internal Server Error"
// @Router /health [get]
func (controller *Controller) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
