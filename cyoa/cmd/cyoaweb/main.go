package main

import (
    "encoding/json"
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

    decoder := json.NewDecoder(file)
    story := cyoa.Story{}
    if err := decoder.Decode(&story); err != nil {
        exit(err)
    }

    fmt.Printf("%+v\n", story)
}
