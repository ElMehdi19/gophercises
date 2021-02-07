package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/ElMehdi19/gophercises/task/db"
	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add new task.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		if len(task) == 0 {
			fmt.Println("Please type a valid task.")
			return
		}
		key, err := db.CreateTask(task)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Printf("Added task `%s` with key %d", task, key)
	},
}

func init() {
	RootCommand.AddCommand(addCommand)
}
