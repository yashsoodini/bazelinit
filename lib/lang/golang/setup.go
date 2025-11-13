package golang

import (
	"embed"
	"fmt"

	"github.com/yashsoodini/bazelinit/lib/scaffold"
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

	if err := scaffold.WriteFiles(workingDir, outputs); err != nil {
		return fmt.Errorf("failed to write files: %w", err)
	}

	fmt.Println("Go repository with Bazel initialized successfully!")

	return nil
}
