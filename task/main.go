package main

import (
	"github.com/ElMehdi19/gophercises/task/cmd"
	"github.com/ElMehdi19/gophercises/task/db"
)

func main() {
	cmd.RootCommand.Execute()
	db.Db.Close()
}
