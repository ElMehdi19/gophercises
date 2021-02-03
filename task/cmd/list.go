package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List uncompleted tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List command called.")
	},
}

func init() {
	RootCommand.AddCommand(listCommand)
}
