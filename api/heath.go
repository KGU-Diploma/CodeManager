package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary Ping the server
// @Description Simple health check to see if the server is running.
// @Tags health
// @Success 200 {object} string "Pong response"
// @Router /health/ping [get]
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
