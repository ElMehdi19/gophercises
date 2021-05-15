package mws

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/ElMehdi19/gophercises/recover_chroma/utils"
)

func RecoverMw(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("panic:", err)

				if utils.IsDevMode() {
					stack := debug.Stack()
					stackWithLinks := utils.CreateLinks(string(stack))
					fmt.Fprintf(w, "<h1><b>panic: </b>%s</h1><pre>%s</pre>", err, stackWithLinks)
					return
				}

				http.Error(w, "something went wrong", http.StatusInternalServerError)
			}
		}()
		handler.ServeHTTP(w, r)
	}
}
