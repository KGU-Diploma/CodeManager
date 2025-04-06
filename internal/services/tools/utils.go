package tools

import (
	"os"
	"fmt"
	"strings"
)

func WriteSourceToFile(filename, source string) error {
	err := os.WriteFile(filename, []byte(source), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filename, err)
	}
	return nil
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