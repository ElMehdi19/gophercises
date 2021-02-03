package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCommand = &cobra.Command{
	Use:   "do",
	Short: "Mark task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []int
		for _, arg := range args {
			task, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("couldn't parse argument", arg)
				continue
			}
			tasks = append(tasks, task)
		}
		fmt.Println(tasks)
	},
}

func init() {
	RootCommand.AddCommand(doCommand)
}
