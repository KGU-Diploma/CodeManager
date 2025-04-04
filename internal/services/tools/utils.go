package tools

import (
	"os"
	"fmt"
)

func WriteSourceToFile(filename, source string) error {
	err := os.WriteFile(filename, []byte(source), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filename, err)
	}
	return nil
}