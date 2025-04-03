package services

import (
	"CodeManager/internal/dto"
	"fmt"
)

type (
	Piston interface {
		ExecuteCode(req dto.ExecuteRequest) (*dto.ExecuteResponse, error)
		GetRuntimes() ([]dto.RuntimeResponse, error)
	}

	Linter interface {
		Lint(source string) ([]string, error)
	}

	Service struct {
		Piston
		Linter
	}
)

type LinterFactory struct{}

func (f *LinterFactory) NewLinter(language string) (Linter, error) {
	switch language {
	case "python":
		return NewPythonLinter(), nil
	default:
		return nil, fmt.Errorf("unsupported language: %s", language)
	}
}


func NewService(language string) (*Service, error) {
	factory := &LinterFactory{}
	linter, err := factory.NewLinter(language)
	if err != nil {
		return nil, err
	}
	return &Service{
		Piston: NewPistonService(),
		Linter: linter,
	}, nil
}