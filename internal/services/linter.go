package services

import "fmt"

type LinterFactory struct{}

func NewLinterFactory() *LinterFactory {
	return &LinterFactory{}
}

func (f *LinterFactory) NewLinter(language string) (Linter, error) {
	switch language {
	case "python3":
		return NewPythonLinter(), nil
	default:
		return nil, fmt.Errorf("unsupported language: %s", language)
	}
}