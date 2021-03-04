package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"
	"time"

	"github.com/ElMehdi19/gophercises/quiet_hn/hn"
)

func main() {
	numStories := 30
	port := 5000

	tpl := template.Must(template.ParseFiles("./index.gohtml"))

	http.HandleFunc("/", handler(numStories, tpl))

	log.Printf("Running on http://127.0.0.1:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		var client hn.Client
		ids, err := client.GetItems()
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		var stories []item
		for _, id := range ids {
			hnItem, err := client.GetItem(id)

			if err != nil {
				continue
			}
			parsedItem := parseItem(hnItem)

			if isStoryLink(parsedItem) {
				stories = append(stories, parsedItem)
			}

			if len(stories) >= numStories {
				break
			}
		}

		data := templateData{
			Stories: stories,
			Time:    time.Now().Sub(started),
		}

		if err := tpl.Execute(w, data); err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	})
}

type item struct {
	hn.Item
	Host string
}

type templateData struct {
	Stories []item
	Time    time.Duration
}

func parseItem(hnItem hn.Item) item {
	newItem := item{Item: hnItem}
	uri, err := url.Parse(newItem.URL)
	if err == nil {
		newItem.Host = strings.TrimPrefix(uri.Host, "www")
	}
	return newItem
}

func isStoryLink(hnItem item) bool {
	return hnItem.Type == "story" && hnItem.URL != ""
}
