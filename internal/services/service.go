package services

import (
	"CodeManager/internal/dto"
)

type (
	Piston interface {
		ExecuteCode(req dto.ExecuteRequest) (*dto.PistonExecuteResponse, error)
		GetRuntimes() ([]dto.RuntimeResponse, error)
	}

	Linter interface {
		Lint(source string) ([]string, error)
		ExtractLinterResult(out string) []string
	}

	Service struct {
		Piston *PistonService
	}
)


func NewService() *Service {
	return &Service{Piston: NewPistonService()}
}