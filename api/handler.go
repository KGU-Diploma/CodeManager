package api

import (
	"CodeManager/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{
	usecases *usecases.Usecase
}

func NewHandler(usecases *usecases.Usecase) *Handler {
	return &Handler{
		usecases: usecases,
	}
}


func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// RegisterRoutes registers all routes for the API
func (h *Handler) SetupRoutes() *gin.Engine {
	router := gin.Default()

	health := router.Group("/health")
	{
		health.GET("/ping", h.Ping)
	}

	domain := router.Group("/api/v1")
	{
		domain.POST("/run-and-analyze", h.RunAndAnalyzeHandler)
		domain.GET("/runtimes", h.GetRuntimesHandler)
	}

	return router
}
