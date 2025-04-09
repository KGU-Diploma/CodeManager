package repositories

import (
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PgSolutionRepository struct {
	db *sqlx.DB
}

func NewPgSolutionRepository(connection *sqlx.DB) *PgSolutionRepository {
	return &PgSolutionRepository{connection}
}

func (r *PgSolutionRepository) CreateSolution(taskId, userId, answerId uuid.UUID, isCorrect bool, subbmittedAt time.Time, answer *string) error {
	query := `insert into t_solution (id_task, id_user, id_answer, is_correct, submitted_at, answer) values ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, taskId, userId, answerId, isCorrect, subbmittedAt, answer)
	if err != nil {
		slog.Error("Error writing solution to database", "error", err)
		return err
	}

	return nil
}
