package services

import (
	"SolutionService/internal/dto"
	"SolutionService/internal/repositories/models"
	"SolutionService/internal/repositories"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type TestsServiceImpl struct {
	testsAnswerRepository repositories.TestsAnswerRepository
	solutionRepository repositories.SolutionRepository
}

func NewTestsService(testAnswerRepo repositories.TestsAnswerRepository, solutionRepo repositories.SolutionRepository) *TestsServiceImpl {
	return &TestsServiceImpl{
		testsAnswerRepository: testAnswerRepo,
		solutionRepository: solutionRepo,
	}
}

func (s *TestsServiceImpl) CreateTestAnswer(taskId uuid.UUID, request dto.CreateTestAnswerRequest) (dto.CreateTestAnswerResponse, error) {
	testsAnswers, err := s.testsAnswerRepository.GetAllByTaskId(taskId)
	if err != nil {
		slog.Info("Failed to get all test answers", "error", err)
		return dto.CreateTestAnswerResponse{}, err
	}

	correctAnswer := s.GetCorrectAnswer(testsAnswers)

	isCorrect := correctAnswer.Id == request.AnswerId

	err = s.solutionRepository.CreateSolution(taskId, request.UserId, &request.AnswerId, isCorrect, time.Now(), nil, nil, nil, nil)
	if err != nil {
		slog.Info("Failed to create solution", "error", err)
		return dto.CreateTestAnswerResponse{}, err
	}
	
	return dto.CreateTestAnswerResponse{
		Success: isCorrect,
	}, nil
}

func (s *TestsServiceImpl) GetCorrectAnswer(testAnswers []models.TestAnswer) models.TestAnswer {
	if len(testAnswers) == 0 {
		return models.TestAnswer{}
	}

	correctAnswer := testAnswers[0]
	for _, answer := range testAnswers {
		if answer.IsCorrect {
			correctAnswer = answer
			break
		}
	}

	return correctAnswer
}