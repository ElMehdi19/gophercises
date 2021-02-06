package main

import (
	"log"
	"os"

	"github.com/ElMehdi19/gophercises/task/cmd"
	"github.com/ElMehdi19/gophercises/task/db"
)

func main() {
	must(db.Init("my.db"))
	must(cmd.RootCommand.Execute())
}

func must(err error) {
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
