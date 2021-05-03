package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Home page</h1>")
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}
