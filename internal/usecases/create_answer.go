package usecases

import (
	"SolutionService/internal/dto"
	"SolutionService/internal/services"
	"log/slog"

	"github.com/google/uuid"
)

type CreateAnswerUsecaseImpl struct {
	services *services.Service
}

func NewCreateAnswerUsecase(services *services.Service) CreateAnswerUsecase {
	return &CreateAnswerUsecaseImpl{
		services: services,
	}
}

func (u *CreateAnswerUsecaseImpl) Handle(taskId uuid.UUID, request dto.CreateTestAnswerRequest) (dto.CreateTestAnswerResponse, error) {
	response, err := u.services.SolutionService.CreateTestSolution(taskId, request)
	if err != nil {
		slog.Error("Error creating test answer", "error", err)
		return dto.CreateTestAnswerResponse{}, err
	}

	return response, nil
}
