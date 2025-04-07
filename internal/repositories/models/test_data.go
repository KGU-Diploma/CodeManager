package models

import "github.com/google/uuid"

type TestData struct {
	Id     uuid.UUID `db:"id"`
	TaskId uuid.UUID `db:"id_task"`
	Input  string `db:"c_input"`
	Output string `db:"c_output"`
}