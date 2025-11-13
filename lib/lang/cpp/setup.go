package cpp

import (
	"embed"
	"fmt"

	"github.com/yashsoodini/bazelinit/lib/scaffold"
	"github.com/yashsoodini/bazelinit/lib/template"
)

//go:embed templates
var templateFS embed.FS

type templateData struct{}

// Setup initializes a new C++ repository with Bazel.
func Setup(workingDir string) error {
	fmt.Println("Initializing C++ repository with Bazel...")

	outputs := template.GenerateOutputs(templateFS, "templates", templateData{})

	if err := scaffold.WriteFiles(workingDir, outputs); err != nil {
		return fmt.Errorf("failed to write files: %w", err)
	}

	fmt.Println("C++ repository with Bazel initialized successfully!")

	return nil
}
