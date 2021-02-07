package cmd

import (
	"fmt"
	"log"

	"github.com/ElMehdi19/gophercises/task/db"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List uncompleted tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListTasks()
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, task := range tasks {
			fmt.Printf("Task %d: %s\n", task.Key, task.Value)
		}
	},
}

func init() {
	RootCommand.AddCommand(listCommand)
}
