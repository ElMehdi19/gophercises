package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/alecthomas/chroma/quick"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Home page</h1>")
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func sourceCodeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	file, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	quick.Highlight(w, buf.String(), "go", "html", "monokai")
}
