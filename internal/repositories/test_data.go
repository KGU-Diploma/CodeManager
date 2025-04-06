package repositories

import (
	"CodeManager/internal/repositories/models"

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

	query := `SELECT * FROM t_programming WHERE task_id = $1`

	err := r.db.Get(&testData, query, taskId)
	if err != nil {
		return nil, err
	}

	return testData, nil
}