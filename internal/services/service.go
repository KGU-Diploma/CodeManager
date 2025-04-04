package services

import (
	"CodeManager/internal/dto"
)

type (
	Piston interface {
		ExecuteCode(req dto.ExecuteRequest) (*dto.ExecuteResponse, error)
		GetRuntimes() ([]dto.RuntimeResponse, error)
	}

	Linter interface {
		Lint(source string) ([]string, error)
		ExtractLinterResult(out string) []string
	}

	Service struct {
		Piston
	}
)


func NewService() *Service {
	return &Service{Piston: NewPistonService()}
}