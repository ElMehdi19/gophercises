package main

import (
	"net/http"
	"net/url"
	"strings"
	"text/template"

	"github.com/ElMehdi19/gophercises/quiet_hn/hn"
)

func main() {

}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		if err := tpl.Execute(w, stories); err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	})
}

type item struct {
	hn.Item
	Host string
}

func parseItem(hnItem hn.Item) item {
	newItem := item{Item: hnItem}
	uri, err := url.Parse(newItem.Host)
	if err == nil {
		newItem.Host = strings.TrimPrefix(uri.Host, "www")
	}
	return newItem
}

func isStoryLink(hnItem item) bool {
	return hnItem.Type == "story" && hnItem.URL != ""
}
