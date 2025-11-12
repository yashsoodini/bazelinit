package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yashsoodini/bazelinit/lib/lang/cpp"
	"github.com/yashsoodini/bazelinit/lib/lang/golang"
)

var workingDirectory string

// rootCmd represents the root bazelinit command.
var rootCmd = &cobra.Command{
	Use:   "bazelinit",
	Short: "bazelinit is a CLI application for initializing code repos with bazel.",
	Long: `
bazelinit is a CLI application for initializing code repos with bazel.
It provides subcommands to set up Bazel for different programming languages.`,
}

// golangCmd represents the golang subcommand.
var golangCmd = &cobra.Command{
	Use:   "go",
	Short: "Initializes a Go repository with Bazel.",
	Long: `
The "go" subcommand initializes a Go repository with Bazel.
It creates a MODULE.bazel file with the necessary Go dependencies,
a basic BUILD file, and a go.mod file.`,
	PreRunE: golang.ValidateCommand,
	RunE: func(cmd *cobra.Command, args []string) error {
		modulePath := cmd.Flag("module_path").Value.String()
		return golang.Setup(modulePath, workingDirectory)
	},
}

// cppCmd represents the cpp subcommand.
var cppCmd = &cobra.Command{
	Use:   "c++",
	Short: "Initializes a C++ repository with Bazel.",
	Long: `The "c++" subcommand is intended to initialize a C++ repository with Bazel.
This feature is currently a work-in-progress.`,
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
	var err error
	workingDirectory, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	golangCmd.Flags().String("module_path", "", "The module path of the Go repository (e.g. github.com/foo/bar)")
	rootCmd.AddCommand(golangCmd)
	rootCmd.AddCommand(cppCmd)
}
