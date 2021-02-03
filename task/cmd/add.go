package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add new task.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Println(task)
	},
}

func init() {
	RootCommand.AddCommand(addCommand)
}
