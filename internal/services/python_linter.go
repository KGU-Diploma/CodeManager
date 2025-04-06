
package services

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"regexp"
	"CodeManager/internal/services/tools"
)

type PythonLinter struct{}

func NewPythonLinter() *PythonLinter {
	return &PythonLinter{}
}

func (l *PythonLinter) Lint(source string) ([]string, error) {
	tempFile := "/tmp/temp_code.py"
	err := tools.WriteSourceToFile(tempFile, source)
	if err != nil {
		return nil, fmt.Errorf("failed to write source to file: %w", err)
	}

	cmd := exec.Command(
		"docker", "run", "--rm",
		"-v", "/tmp:/app",
		"python-linter",
		"/app/temp_code.py",
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err = cmd.Run()
	output := out.String()
	if err != nil {
		// Если Ruff что-то нашёл, мы возвращаем как результат, а не ошибку
		if strings.Contains(output, "Found") {
			result := l.ExtractLinterResult(output)
			return result, nil
		}
		// Иначе — это реальная проблема с запуском линтера
		return nil, fmt.Errorf("error running ruff in Docker: %w\nOutput:\n%s", err, output)
	}
	
	result := l.ExtractLinterResult(output)

	return result, nil
}

func (l *PythonLinter) ExtractLinterResult(output string) []string {
	var issues []string
	
	if len(output) == 0 {
		return issues
	}

	lines := strings.Split(output, "\n")

	// Пример строки: temp_code.py:1:8: F401 [*] `os` imported but unused
	r := regexp.MustCompile(`^.+:\d+:\d+:\s+[A-Z]\d{3}`)

	for _, line := range lines {
		if r.MatchString(line) {
			issues = append(issues, line)
		}
	}

	return issues
}