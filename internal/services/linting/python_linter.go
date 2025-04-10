package linting

import (
	"SolutionService/internal/services/container"
	"SolutionService/internal/services/tools"
	"fmt"
	"log/slog"
	"path/filepath"
	"regexp"
	"strings"
)

type PythonLinter struct {
	runner container.ContainerRunner
}

func NewPythonLinter(runner container.ContainerRunner) *PythonLinter {
	return &PythonLinter{
		runner: runner,
	}
}

func (l *PythonLinter) Lint(source string) ([]string, error) {
	codePath, cleanup, err := tools.WriteToTempFile(source, "main.py")
	if err != nil {
		slog.Error("Error writing to temporary file", "error", err)
		return nil, err
	}
	defer cleanup()

	output, err := l.runner.RunContainer(filepath.Dir(codePath), "python-linter")
	if err != nil {
		slog.Error("Error starting container", "error", err)
		return nil, fmt.Errorf("docker failed: %w", err)
	}
	slog.Info("Linting output", "output", output)
	return l.ExtractLinterResult(output), nil
}

func (l *PythonLinter) ExtractLinterResult(output string) []string {
	var issues []string

	if len(output) == 0 {
		return issues
	}

	lines := strings.Split(output, "\n")
	r := regexp.MustCompile(`^.+:\d+:\d+:\s+[A-Z]\d{3}`)

	for _, line := range lines {
		if r.MatchString(line) {
			issues = append(issues, line)
		}
	}

	return issues
}
