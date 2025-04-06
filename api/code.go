package api

import (
	"CodeManager/internal/dto"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

// RunAndAnalyzeHandler godoc
// @Summary Run and analyze the provided input
// @Description This endpoint runs and analyzes the input provided in the request body.
// @Tags analysis
// @Accept  json
// @Produce  json
// @Param request body dto.ExecuteRequest true "Input data for analysis"
// @Success 200 {object} dto.MultiExecuteResponse "Successfully run and analyzed the data"
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/run-and-analyze [post]
func (h *Handler) RunAndAnalyzeHandler(c *gin.Context) {
	var request dto.ExecuteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, "Invalid input")
		return
	}

	result, err := h.usecases.ExecuteCodeUsecase.Handle(request)
	if err != nil {
		slog.Error("Failed to execute code", "error", err)
		c.JSON(500, "Internal server error")
		return
	}
	
	c.JSON(200, result)
}

// GetRuntimesHandler godoc
// @Summary Get all available runtimes
// @Description Fetches a list of all available runtimes from the usecase layer.
// @Tags runtimes
// @Produce  json
// @Success 200 {array} dto.RuntimeResponse "List of runtimes"
// @Failure 500 {object} string "Failed to fetch runtimes"
// @Router /api/v1/runtimes [get]
func (h *Handler) GetRuntimesHandler(c *gin.Context) {
	runtimes, err := h.usecases.GetRuntimesUsecase.Handle()
	if err != nil {
		slog.Error("Failed to fetch runtimes", "error", err)
		c.JSON(500, "Failed to fetch runtimes")
		return
	}
	c.JSON(200, runtimes)
}
