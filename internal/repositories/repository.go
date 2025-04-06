package repositories

import (
	"CodeManager/internal/repositories/models"

	"github.com/jmoiron/sqlx"
)

type (
	TestDataRepository interface {
		GetTestDataByTaskId(taskId string) ([]models.TestData, error)
	}

	Repository struct {
		TestData *PgTestDataRepository
	}
)

func NewRepository(connection *sqlx.DB) *Repository {
	// Initialize the repository with its dependencies
	return &Repository{
		TestData: NewPgTestDataRepository(connection),
	}
}