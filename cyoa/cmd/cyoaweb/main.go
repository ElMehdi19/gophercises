package main

import (
    "flag"
    "fmt"
    "os"

    "github.com/ElMehdi19/gophercises/cyoa"
)

func exit(err error){
    fmt.Errorf("%s", err)
    os.Exit(1)
}

func main(){
    fileName := flag.String("story", "gopher.json", "JSON file containing the story")
    flag.Parse()

    file, err := os.Open(*fileName)
    if err != nil {
        exit(err)
    }
    defer file.Close()

    story, err := cyoa.JsonStoryParser(file)
    if err != nil {
        exit(err)
    }

    fmt.Printf("%+v\n", story)
}
