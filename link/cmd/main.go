package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ElMehdi19/gophercises/link"
)

func main() {
	file, err := os.Open("example1.html")
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	defer file.Close()

	htmlBytes, _ := ioutil.ReadAll(file)
	htmlStr := string(htmlBytes)
	reader := strings.NewReader(htmlStr)
	links, err := link.Parse(reader)
	if err != nil {
		log.Fatal("Couldn't parse you html")
	}
	_ = links
}
