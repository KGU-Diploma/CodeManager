
package services

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type PythonLinter struct{}

func NewPythonLinter() *PythonLinter {
	return &PythonLinter{}
}

func (l *PythonLinter) Lint(source string) ([]string, error) {
	var issues []string

	tempFile := "/tmp/temp_code.py"
	err := WriteSourceToFile(tempFile, source)
	if err != nil {
		return nil, fmt.Errorf("failed to write source to file: %w", err)
	}

	cmd := exec.Command("docker", "run", "--rm", "-v", fmt.Sprintf("%s:/app", "/tmp"), "python-linter", tempFile)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error running ruff in Docker: %w", err)
	}

	output := out.String()
	if len(output) == 0 {
		issues = append(issues, "No issues found!")
	} else {
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			if line != "" {
				issues = append(issues, line)
			}
		}
	}

	return issues, nil
}
