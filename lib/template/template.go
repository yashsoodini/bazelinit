package template

import (
	"bytes"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"
)

// GenerateOutputs parses the provided templates, executes them on the provided data,
// and returns a map of file paths to their generated content.
//
// templatesFS: The embedded filesystem containing the templates.
// templatesRoot: The root directory within the embedded filesystem where templates are located.
// templateData: The data structure to be passed to the templates for execution.
func GenerateOutputs(templatesFS fs.FS, templatesRoot string, templateData any) map[string][]byte {
	outputs := make(map[string][]byte)

	err := fs.WalkDir(templatesFS, templatesRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		relativePath, err := filepath.Rel(templatesRoot, path)
		if err != nil {
			return err
		}

		tmplContent, err := fs.ReadFile(templatesFS, path)
		if err != nil {
			return err
		}

		var outputPath string
		var outputContent []byte

		if filepath.Ext(path) == ".tmpl" {
			outputPath = strings.TrimSuffix(relativePath, ".tmpl")

			tmpl, err := template.New(path).Parse(string(tmplContent))
			if err != nil {
				return fmt.Errorf("failed to parse template %s: %w", path, err)
			}

			var buf bytes.Buffer
			if err := tmpl.Execute(&buf, templateData); err != nil {
				return fmt.Errorf("executing template %s: %w", path, err)
			}

			outputContent = buf.Bytes()
		} else {
			outputPath = relativePath
			outputContent = tmplContent
		}

		outputs[outputPath] = outputContent
		return nil
	})
	if err != nil {
		panic(err)
	}

	return outputs
}
