package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler struct to group all handlers
type Handler struct{}

// NewHandler initializes a new Handler
func NewHandler() *Handler {
	return &Handler{}
}


// Ping is a simple health check endpoint
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// RegisterRoutes registers all routes for the API
func (h *Handler) SetupRoutes(router *gin.Engine) {
    router.GET("/ping", h.Ping)
}
