package services

import (
	"SolutionService/internal/dto"
	"SolutionService/internal/repositories"
	"SolutionService/internal/repositories/models"
	"github.com/google/uuid"
)

type (
	PistonService interface {
		ExecuteCode(req dto.PistonExecuteRequest) (*dto.PistonExecuteResponse, error)
		GetRuntimes() ([]dto.RuntimeResponse, error)
	}

	Linter interface {
		Lint(source string) ([]string, error)
		ExtractLinterResult(out string) []string
	}

	TestsService interface{
		CreateTestAnswer(taskId uuid.UUID, request dto.CreateTestAnswerRequest) (dto.CreateTestAnswerResponse, error)
		GetCorrectAnswer(testAnswers []models.TestAnswer) models.TestAnswer
	}

	Service struct {
		Piston PistonService
		TestsService TestsService
	}
)


func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Piston: NewPistonService(),
		TestsService: NewTestsService(repos.TestsAnswerRepository, repos.SolutionRepository),
	}
}