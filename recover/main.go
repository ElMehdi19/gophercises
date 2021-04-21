package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func main() {
	port := 5000
	mux := http.NewServeMux()

	mux.HandleFunc("/", hello)
	mux.HandleFunc("/panic", panicDemo)
	mux.HandleFunc("/panic-after", panicAfterDemo)

	log.Printf("Running on http://127.0.0.1:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), recoverMw(mux, true)))
}

func funcThatPanics() {
	panic("Oh geez Rick!!")
}

func recoverMw(h http.Handler, dev bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				if dev {
					w.WriteHeader(http.StatusInternalServerError)
					stack := debug.Stack()
					fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, string(stack))
					return
				}
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
				return
			}
		}()
		h.ServeHTTP(w, r)
	}
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
}
