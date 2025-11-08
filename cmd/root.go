/*
Copyright © 2025 Yash Soodini yash@soodini.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yashsoodini/bazelinit/lang/cpp"
	"github.com/yashsoodini/bazelinit/lang/golang"
)

// rootCmd represents the root bazelinit command.
var rootCmd = &cobra.Command{
	Use:   "bazelinit",
	Short: "bazelinit is a CLI application for initializing code repos with bazel.",
	Long: `
bazelinit is a CLI application for initializing code repos with bazel. It initializes
a git repository in the current directory, adds bazel and gazelle configuration for the
specified language, and adds all bazel files to .gitignore.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		lang := cmd.Flag("lang").Value.String()
		switch lang {
		case "go", "c++":
		default:
			return fmt.Errorf("invalid --lang: %s (want go|c++)", lang)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		lang := cmd.Flag("lang").Value.String()
		switch lang {
		case "go":
			golang.Setup()
		case "c++":
			cpp.Setup()
		default:
			// This should never happen due to PreRunE validation.
			return fmt.Errorf("invalid --lang: %s (want go|c++)", lang)
		}
		return nil
	},
}

// Execute executes the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("lang", "l", "go", "The programming language used in the repository.")
}
