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

	SolutionService interface {
		CreateCodingSolution(taskId string, userId uuid.UUID, testResults []dto.TestCaseResult, code, language string, lintingIssues []string) error
		CreateTestSolution(taskId uuid.UUID, request dto.CreateTestAnswerRequest) (dto.CreateTestAnswerResponse, error)
		GetCorrectAnswer(testAnswers []models.TestAnswer) models.TestAnswer
	}

	Service struct {
		Piston PistonService
		SolutionService SolutionService
	}
)


func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Piston: NewPistonService(),
		SolutionService: NewSolutionService(repos.TestsAnswerRepository, repos.SolutionRepository),
	}
}