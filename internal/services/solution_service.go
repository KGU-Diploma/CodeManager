package services

import (
	"SolutionService/internal/dto"
	"SolutionService/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type SolutionServiceImpl struct {
	SolutionRepository repositories.SolutionRepository
}

func (s *SolutionServiceImpl) CreateSolution(taskId, userId uuid.UUID, testResults []dto.TestCaseResult, code, language string, lintingResult bool) error {
    isCorrect := true
    for _, testResult := range testResults {
        isCorrect = testResult.Passed
    }

    return s.SolutionRepository.CreateSolution(
        taskId,
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
