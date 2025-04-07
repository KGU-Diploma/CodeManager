package linting

import (
	"CodeManager/internal/services"
	"CodeManager/internal/services/container"
	"fmt"
)

type LinterFactory struct{
	runner container.ContainerRunner
}

func NewLinterFactory(runner container.ContainerRunner) *LinterFactory {
	return &LinterFactory{
		runner: runner,
	}
}

func (f *LinterFactory) NewLinter(language string) (services.Linter, error) {
	switch language {
	case "python3":
		return NewPythonLinter(f.runner), nil
	case "java":
		return NewJavaLinter(f.runner), nil
	default:
		return nil, fmt.Errorf("unsupported language: %s", language)
	}
}