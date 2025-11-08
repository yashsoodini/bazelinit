/*
Copyright © 2025 Yash Soodini yash@soodini.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bazelinit",
	Short: "bazelinit is a CLI application for initializing code repos with bazel.",
	Long: `bazelinit is a CLI application for initializing code repos with bazel. It initializes
	a git repository in the current directory, adds bazel and gazelle configuration for the specified
	language, and adds all bazel files to .gitignore.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bazelinit is a work in-progress...")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
