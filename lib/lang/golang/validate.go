package golang

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/mod/module"
)

func ValidateCommand(cmd *cobra.Command, args []string) error {
	modulePath := cmd.Flag("module_path").Value.String()
	if modulePath == "" {
		return fmt.Errorf("required flag \"module_path\" not set")
	}
	return validateModulePath(modulePath)
}

func validateModulePath(modulePath string) error {
	if err := module.CheckPath(modulePath); err != nil {
		return fmt.Errorf("%s is not a valid Go module path: %v", modulePath, err)
	}
	return nil
}
