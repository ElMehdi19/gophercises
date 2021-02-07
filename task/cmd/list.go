package cmd

import (
	"fmt"
	"os"

	"github.com/ElMehdi19/gophercises/task/db"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List uncompleted tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Println("Error occured:", err.Error())
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks, just take a nap.")
			return
		}

		fmt.Println("You have to complete the following tasks:")
		for i, task := range tasks {
			fmt.Printf("Task %d: %s\n", i, task.Value)
		}
	},
}

func init() {
	RootCommand.AddCommand(listCommand)
}
