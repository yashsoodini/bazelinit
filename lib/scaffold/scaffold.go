package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
)

// WriteFiles writes the given map of file paths and their content to the specified directory.
func WriteFiles(dir string, filePathToContents map[string][]byte) error {
	for path, content := range filePathToContents {
		fullPath := filepath.Join(dir, path)
		dir := filepath.Dir(fullPath)

		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("creating directory %s: %w", dir, err)
		}
		if err := os.WriteFile(fullPath, content, 0644); err != nil {
			return fmt.Errorf("writing file %s: %w", fullPath, err)
		}

		fmt.Printf("Created %s\n", fullPath)
	}
	return nil
}
