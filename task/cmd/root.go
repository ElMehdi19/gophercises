package cmd

import "github.com/spf13/cobra"

// RootCommand task CLI root command
var RootCommand = &cobra.Command{
	Use:   "task",
	Short: "Task manager CLI tool",
}
