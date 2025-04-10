package repositories

import (
	"SolutionService/internal/repositories/models"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type PgTestDataRepository struct {
	db *sqlx.DB
}

func NewPgTestDataRepository(connection *sqlx.DB) *PgTestDataRepository {
	return &PgTestDataRepository{connection}
}

func (r *PgTestDataRepository) GetTestDataByTaskId(taskId string) ([]models.TestData, error) {
	var testData []models.TestData

	query := `SELECT * FROM t_programming WHERE id_task = $1`

	err := r.db.Select(&testData, query, taskId)
	if err != nil {
		slog.Error("Error Getting Test Data By Task Identifier", "error", err)
		return nil, err
	}

	return testData, nil
}
