package services

import (
	"SolutionService/internal/dto"
	"SolutionService/internal/repositories/models"
	"SolutionService/internal/repositories"
	"time"
    "log/slog"
	"github.com/google/uuid"
)

type SolutionServiceImpl struct {
    testsAnswerRepository repositories.TestsAnswerRepository
	solutionRepository repositories.SolutionRepository
}

func NewSolutionService(testAnswerRepo repositories.TestsAnswerRepository, solutionRepo repositories.SolutionRepository) *SolutionServiceImpl {
	return &SolutionServiceImpl{
        testsAnswerRepository: testAnswerRepo,
		solutionRepository: solutionRepo,
	}
}

func (s *SolutionServiceImpl) CreateCodingSolution(taskId string, userId uuid.UUID, testResults []dto.TestCaseResult, code, language string, lintingIssues []string) error {
    isCorrect := true
    for _, testResult := range testResults {
        isCorrect = testResult.Passed
    }

    taskUUID, err := uuid.Parse(taskId)
    if err != nil {
        slog.Error("Occured error while parsing taskId", "taskId", taskId, "error", err)
        return err
    }
    
    // If no issues are found, then we consider it passed
    lintingResult := len(lintingIssues) == 0

    return s.solutionRepository.CreateSolution(
        taskUUID,
        userId,
        nil,
        isCorrect,
        time.Now(),
        &code,
        &language,
        nil,
        &lintingResult,
    )
}

func (s *SolutionServiceImpl) CreateTestSolution(taskId uuid.UUID, request dto.CreateTestAnswerRequest) (dto.CreateTestAnswerResponse, error) {
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

func (s *SolutionServiceImpl) GetCorrectAnswer(testAnswers []models.TestAnswer) models.TestAnswer {
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
