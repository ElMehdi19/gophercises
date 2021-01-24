package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ElMehdi19/gophercises/link"
)

func main() {
	fileName := flag.String("html", "example1.html", "`HTML file` to parse links from")
	flag.Parse()

	file, err := os.Open(*fileName)
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
	for _, link := range links {
		fmt.Println(link)
	}
}
