package repositories

import (
	"SolutionService/internal/repositories/models"
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
	"time"
)

type (
	TestDataRepository interface {
		GetTestDataByTaskId(taskId string) ([]models.TestData, error)
	}

	TestsAnswerRepository interface {
		GetAllByTaskId(taskId uuid.UUID) ([]models.TestAnswer, error)
	}

	SolutionRepository interface {
		CreateSolution(taskId, userId, answerId uuid.UUID, isCorrect bool, subbmittedAt time.Time, answer *string) error	}

	Repository struct {
		TestData *PgTestDataRepository
		TestsAnswerRepository *PgTestAnswerRepository
		SolutionRepository *PgSolutionRepository
	}
)

func NewRepository(connection *sqlx.DB) *Repository {
	return &Repository{
		TestData: NewPgTestDataRepository(connection),
		TestsAnswerRepository: NewPgTestAnswerRepository(connection),
		SolutionRepository: NewPgSolutionRepository(connection),
	}
}