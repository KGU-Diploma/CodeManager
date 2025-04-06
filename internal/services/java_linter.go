package services

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"regexp"
	"CodeManager/internal/services/tools"
)

type JavaLinter struct{}

func NewJavaLinter() *JavaLinter {
	return &JavaLinter{}
}

func (l *JavaLinter) Lint(source string) ([]string, error) {
	tempFile := "/tmp/Main.java"
	err := tools.WriteSourceToFile(tempFile, source)
	if err != nil {
		return nil, fmt.Errorf("failed to write source to file: %w", err)
	}

	cmd := exec.Command(
		"docker", "run", "--rm",
		"-v", "/tmp:/app",
		"java-linter",
		"javac", "/app/Main.java",
	)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	cmd.Run()
	output := out.String()
	issues := l.ExtractLinterResult(output)

	return issues, nil
}


func (l *JavaLinter) ExtractLinterResult(output string) []string {
	var issues []string
	lines := strings.Split(output, "\n")
	r := regexp.MustCompile(`^.+\.java:\d+:\s+error:`)

	if len(output) == 0 {
		return issues
	}

	for _, line := range lines {
		if r.MatchString(line) {
			issues = append(issues, line)
		}
	}

	return issues
}
