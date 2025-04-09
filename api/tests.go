package api

import (
	"SolutionService/internal/dto"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"github.com/google/uuid"
)

func (h *Handler) CreateTestAnswerHandler(c *gin.Context) {
	taskId := c.Param("taskId")
	
	// Преобразуем taskId в uuid, если необходимо
	taskUUID, err := uuid.Parse(taskId)
	if err != nil {
		c.JSON(400, "Invalid taskId format")
		return
	}
	var request dto.CreateTestAnswerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, "Invalid input")
		return
	}

	result, err := h.usecases.CreateAnswerUsecase.Handle(taskUUID, request)
	if err != nil {
		slog.Error("Failed to create test answer", "error", err)
		c.JSON(500, "Internal server error")
		return
	}
	
	c.JSON(200, result)
}