package main

import (
	"fmt"
	"log"

	"github.com/ElMehdi19/gophercises/quiet_hn/hn"
)

func main() {
	client := hn.Client{}
	ids, err := client.GetItems()
	if err != nil {
		log.Fatalf(err.Error())
	}
	for i, id := range ids {
		if i >= 30 {
			return
		}
		fmt.Printf("%d ", id)
	}
}
