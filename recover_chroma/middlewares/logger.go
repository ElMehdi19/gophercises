package mws

import (
	"log"
	"net/http"
)

func LoggingMw(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		method := r.Method
		remote := r.RemoteAddr

		log.Printf("%s %s %s\n", method, path, remote)
		handler.ServeHTTP(w, r)
	}
}
