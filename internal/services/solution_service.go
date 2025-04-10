package services

import (
	"SolutionService/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type SolutionService struct {
	SolutionRepository repositories.SolutionRepository
}

func (s *SolutionService) SaveSolution(taskId, userId, answerId uuid.UUID, isCorrect bool, submittedAt time.Time, answer *string) error {
	return s.SolutionRepository.CreateSolution(taskId, userId, answerId, isCorrect, submittedAt, answer)
}
