package api

import (
	"github.com/gin-gonic/gin"
	"CodeManager/internal/dto"
)

// RunAndAnalyzeHandler godoc
// @Summary Run and analyze the provided input
// @Description This endpoint runs and analyzes the input provided in the request body.
// @Tags analysis
// @Accept  json
// @Produce  json
// @Param request body dto.ExecuteRequest true "Input data for analysis"
// @Success 200 {object} dto.ExecuteResponse "Successfully run and analyzed the data"
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/run-and-analyze [post]
func (h *Handler) RunAndAnalyzeHandler(c *gin.Context) {
	var request dto.ExecuteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, "Invalid input")
		return
	}

	// Здесь реализуйте логику для анализа данных
	// Примерный ответ
	// c.JSON(200, ResultData{
	// 	Status:  "success",
	// 	Message: "Data successfully analyzed",
	// })
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
		c.JSON(500, "Failed to fetch runtimes")
		return
	}
	c.JSON(200, runtimes)
}
