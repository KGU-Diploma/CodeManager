package usecases

import (
	"SolutionService/internal/dto"
	"SolutionService/internal/services"

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
	response, err := u.services.TestsService.CreateTestAnswer(taskId, request)
	if err != nil {
		return dto.CreateTestAnswerResponse{}, err
	}

	return response, nil
}