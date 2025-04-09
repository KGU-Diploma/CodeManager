package models

import "github.com/google/uuid"

type TestAnswer struct {
	Id uuid.UUID `db:"id"`
	TaskId uuid.UUID `db:"id_task"`
	Answer string `db:"answer"`
	IsCorrect bool `db:"is_correct"`
}