package models

type TestData struct {
	Id     string `db:"id"`
	TaskId string `db:"task_id"`
	Input  string `db:"input"`
	Output string `db:"output"`
}