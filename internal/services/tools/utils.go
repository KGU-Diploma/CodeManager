package tools

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

// WriteToTempFile создает временный файл с указанным содержимым
// Возвращает путь к файлу и функцию очистки (defer)
func WriteToTempFile(content, pattern string) (string, func(), error) {
    dir, err := os.MkdirTemp("", "code-*")
    if err != nil {
		slog.Error("Failed to create temp dir", "error", err)
        return "", nil, fmt.Errorf("failed to create temp dir: %w", err)
    }

    cleanup := func() {
        os.RemoveAll(dir)
    }

    filename := filepath.Join(dir, pattern)
    if err := writeSourceToFile(filename, content); err != nil {
        cleanup()
		slog.Error("Failed to write source to file", "filename", filename, "error", err)
        return "", nil, fmt.Errorf("failed to write temp file: %w", err)
    }

    return filename, cleanup, nil
}


func CompareExpectedAndActual(expected, actual string) bool {
	// Normalize line endings and trim whitespace
	cleanExpected := normalizeOutput(expected)
	cleanActual := normalizeOutput(actual)

	return cleanExpected == cleanActual
}

// normalizeOutput cleans up the string by trimming and unifying line endings.
func normalizeOutput(s string) string {
	// Replace Windows-style line endings with Unix-style
	s = strings.ReplaceAll(s, "\r\n", "\n")

	// Trim spaces from each line and remove trailing newline
	lines := strings.Split(strings.TrimSpace(s), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	// Join cleaned lines back
	return strings.Join(lines, "\n")
}

func writeSourceToFile(filename, source string) error {
	err := os.WriteFile(filename, []byte(source), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filename, err)
	}
	return nil
}