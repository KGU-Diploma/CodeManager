package repositories

import (
	"SolutionService/internal/repositories/models"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PgTestAnswerRepository struct {
	db *sqlx.DB
}

func NewPgTestAnswerRepository(connection *sqlx.DB) *PgTestAnswerRepository {
	return &PgTestAnswerRepository{connection}
}

func (r *PgTestAnswerRepository) GetAllByTaskId(taskId uuid.UUID) ([]models.TestAnswer, error) {
	var testAnswers []models.TestAnswer

	query := `SELECT * FROM t_test_answer WHERE id_task = $1`

	err := r.db.Select(&testAnswers, query, taskId)
	if err != nil {
		slog.Error("Error getting all tasks by Id", "error", err)
		return nil, err
	}

	return testAnswers, nil
}
