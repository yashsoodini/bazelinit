package golang

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/yashsoodini/bazelinit/lib/template"
)

type templateData struct {
	ModulePath string
}

//go:embed templates
var templateFS embed.FS

// Setup initializes a new Go repository with Bazel.
func Setup(modulePath string, workingDir string) error {
	fmt.Println("Initializing Go repository with Bazel...")

	tmplData := templateData{
		ModulePath: modulePath,
	}
	outputs := template.GenerateOutputs(templateFS, "templates", tmplData)

	for path, content := range outputs {
		fullPath := filepath.Join(workingDir, path)
		dir := filepath.Dir(fullPath)

		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("creating directory %s: %w", dir, err)
		}
		if err := os.WriteFile(fullPath, content, 0644); err != nil {
			return fmt.Errorf("writing file %s: %w", fullPath, err)
		}

		fmt.Printf("Created %s\n", fullPath)
	}

	fmt.Println("Go repository with Bazel initialized successfully!")

	return nil
}
