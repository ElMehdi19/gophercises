package cmd

import (
	"fmt"
	"strconv"

	"github.com/ElMehdi19/gophercises/task/db"
	"github.com/spf13/cobra"
)

var doCommand = &cobra.Command{
	Use:   "do",
	Short: "Mark task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var keys []int
		for _, arg := range args {
			task, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("couldn't parse argument", arg)
				continue
			}
			keys = append(keys, task)
			for _, key := range keys {
				err := db.DeleteTask(key)
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Printf("Task %d done congrats.\n", key)
			}
		}
	},
}

func init() {
	RootCommand.AddCommand(doCommand)
}
