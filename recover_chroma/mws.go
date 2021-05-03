package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func loggingMw(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		method := r.Method
		remote := r.RemoteAddr

		log.Printf("%s %s %s\n", method, path, remote)
		handler.ServeHTTP(w, r)
	}
}

func recoverMw(handler http.Handler, devMode bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("panic:", err)

				if devMode {
					stack := debug.Stack()
					fmt.Fprintf(w, "<h1><b>panic: </b>%s</h1><pre>%s</pre>", err, string(stack))
					return
				}

				http.Error(w, "something went wrong", http.StatusInternalServerError)
			}
		}()
		handler.ServeHTTP(w, r)
	}
}
