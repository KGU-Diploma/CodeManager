package dto

import "github.com/google/uuid"

type CreateTestAnswerRequest struct {
	UserId  uuid.UUID `json:"user_id" binding:"required"`
	AnswerId uuid.UUID `json:"answer_id" binding:"required"`
}

type CreateTestAnswerResponse struct {
	Success bool `json:"success"`
}