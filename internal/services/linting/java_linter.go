package linting

import (
	"SolutionService/internal/services/container"
	"SolutionService/internal/services/tools"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

type JavaLinter struct{
	runner container.ContainerRunner
}

func NewJavaLinter(runner container.ContainerRunner) *JavaLinter {
	return &JavaLinter{
        runner: runner,
    }
}

func (l *JavaLinter) Lint(source string) ([]string, error) {
	codePath, cleanup, err := tools.WriteToTempFile(source, "Main.java")
	if err != nil {
		return nil, err
	}
	defer cleanup()

	output, err := l.runner.RunContainer(filepath.Dir(codePath), "java-linter")
	if err != nil {
		return nil, fmt.Errorf("docker failed: %w", err)
	}

	return l.ExtractLinterResult(output), nil
}

func (l *JavaLinter) ExtractLinterResult(output string) []string {
    var issues []string
    seen := make(map[string]bool)
    
    pmdRe := regexp.MustCompile(`/code/Main\.java:\d+:\t(.+?):\t(.+)`)
    checkstyleRe := regexp.MustCompile(`\[(ERROR|WARN)\] (.+\.java:\d+:.+)`)

    for _, line := range strings.Split(output, "\n") {
        var issue string

		if matches := pmdRe.FindStringSubmatch(line); len(matches) > 2 {
            rule := matches[1]
            message := matches[2]
            
            if rule == "NoPackage" {
                continue
            }
            
            issue = fmt.Sprintf("%s: %s", rule, message)
            
        } else if matches := checkstyleRe.FindStringSubmatch(line); len(matches) > 2 {
            issue = matches[2]
        }
        
        if issue != "" && !seen[issue] {
            seen[issue] = true
            issues = append(issues, issue)
        }
    }
    
    return issues
}