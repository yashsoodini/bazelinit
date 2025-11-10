/*
Copyright © 2025 Yash Soodini yash@soodini.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yashsoodini/bazelinit/lib/lang/cpp"
	"github.com/yashsoodini/bazelinit/lib/lang/golang"
)

// rootCmd represents the root bazelinit command.
var rootCmd = &cobra.Command{
	Use:   "bazelinit",
	Short: "bazelinit is a CLI application for initializing code repos with bazel.",
	Long: `
bazelinit is a CLI application for initializing code repos with bazel. It initializes
a git repository in the current directory, adds bazel and gazelle configuration for the
specified language, and adds all bazel files to .gitignore.`,
}

// golangCmd represents the golang subcommand.
var golangCmd = &cobra.Command{
	Use:   "go",
	Short: "Initializes a Go repository with Bazel.",
	Long: `
golang is a subcommand for bazelinit that initializes a Go repository with Bazel.
It sets up the necessary Bazel and Gazelle configurations for Go, and adds all
Bazel-related files to .gitignore.`,
	PreRunE: golang.ValidateCommand,
	RunE: func(cmd *cobra.Command, args []string) error {
		return golang.Setup()
	},
}

// cppCmd represents the cpp subcommand.
var cppCmd = &cobra.Command{
	Use:   "c++",
	Short: "Initializes a C++ repository with Bazel.",
	Long: `c++ is a subcommand for bazelinit that initializes a C++ repository with Bazel.
It sets up the necessary Bazel configurations for C++, and adds all
Bazel-related files to .gitignore.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cpp.Setup()
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
	golangCmd.Flags().String("module_path", "", "The module path of the Go repository (e.g. github.com/foo/bar)")
	rootCmd.AddCommand(golangCmd)
	rootCmd.AddCommand(cppCmd)
}
