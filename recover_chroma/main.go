package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	mws "github.com/ElMehdi19/gophercises/recover_chroma/middlewares"
	"github.com/ElMehdi19/gophercises/recover_chroma/routes"
	"github.com/ElMehdi19/gophercises/recover_chroma/utils"
)

var port int

func main() {
	flag.IntVar(&port, "port", 5000, "port to run the web app on")
	flag.Parse()

	handler := newHandler()

	log.Printf("Running on http://127.0.0.1:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}

func newHandler() http.HandlerFunc {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.Home)
	mux.HandleFunc("/panic", routes.PanicDemo)

	if utils.IsDevMode() {
		mux.HandleFunc("/debug/", routes.SourceCodeHandler)
	}

	return mws.LoggingMw(mws.RecoverMw(mux))
}
