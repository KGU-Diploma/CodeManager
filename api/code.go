package api

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) RunAndAnalyzeHandler(c *gin.Context) {
	
}


func (h *Handler) GetRuntimesHandler(c *gin.Context) {
	runtimes, err := h.usecases.GetRuntimesUsecase.Handle()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get runtimes"})
		return
	}
	c.JSON(200, runtimes)
}