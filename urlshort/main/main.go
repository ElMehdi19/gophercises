package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ElMehdi19/gophercises/urlshort"
)

func exit(err error) {
	log.Fatalf("Error: %s", err)
	os.Exit(1)
}

func readJSON() []byte {
	file, err := os.Open("paths.json")
	if err != nil {
		exit(err)
	}
	defer file.Close()
	lines, err := ioutil.ReadAll(file)
	if err != nil {
		exit(err)
	}
	return lines
}

func main() {
	mux := defaultMux()
	paths := map[string]string{
		"/fb":  "https://fb.me/Mehdiii",
		"/tw":  "https://twitter.com/DMehdi19",
		"/git": "https://github.com/ElMehdi19",
	}
	mapHandler := urlshort.MapHandler(paths, mux)

	jsonPaths := readJSON()
	jsonHandler, err := urlshort.JSONHandler(jsonPaths, mapHandler)
	if err != nil {
		panic(err)
	}

	port := ":5000"
	log.Printf("Running on http://127.0.0.1:%s", port)
	log.Fatal(http.ListenAndServe(port, jsonHandler))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})
	return mux
}
