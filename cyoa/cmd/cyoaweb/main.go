package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ElMehdi19/gophercises/cyoa"
)

func exit(err error) {
	fmt.Println(fmt.Errorf("%s", err))
	os.Exit(1)
}

func main() {
	fileName := flag.String("story", "gopher.json", "JSON file containing the story")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		exit(err)
	}
	defer file.Close()

	story, err := cyoa.JSONStoryParser(file)
	if err != nil {
		exit(err)
	}

	handler := cyoa.NewHandler(story)
	port := 5000

	log.Printf("Running on http://127.0.0.1:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
