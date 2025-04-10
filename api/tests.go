package api

import (
	"SolutionService/internal/dto"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

func (h *Handler) CreateTestAnswerHandler(c *gin.Context) {
	taskId := c.Param("taskId")

	// Преобразуем taskId в uuid, если необходимо
	taskUUID, err := uuid.Parse(taskId)
	if err != nil {
		slog.Error("Invalid taskId format", "taskId", taskId, "error", err)
		c.JSON(http.StatusBadRequest, "Invalid taskId format")
		return
	}
	var request dto.CreateTestAnswerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		slog.Error("Invalid input", "error", err)
		c.JSON(http.StatusBadRequest, "Invalid input")
		return
	}

	result, err := h.usecases.CreateAnswerUsecase.Handle(taskUUID, request)
	if err != nil {
		slog.Error("Failed to create test answer", "error", err)
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, result)
}
