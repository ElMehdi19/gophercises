package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/ElMehdi19/gophercises/quiet_hn/hn"
)

func main() {
	numStories := flag.Int("num_stories", 30, "how many stories to display")
	port := flag.Int("port", 5000, "port to start the web server on")
	flag.Parse()

	tpl := template.Must(template.ParseFiles("./index.gohtml"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler(*numStories, tpl))

	log.Printf("Running on http://127.0.0.1:%d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))

}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()

		stories, err := getStoriesFromCache(numStories)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
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

var (
	cache    []item
	cacheExp time.Time
)

func getStoriesFromCache(numStories int) ([]item, error) {
	if cacheExp.Sub(time.Now()) > 0 {
		return cache, nil
	}

	items, err := getStoriesAsync(numStories)
	if err != nil {
		return nil, err
	}
	cache = items
	cacheExp = time.Now().Add(15 * time.Second)

	return cache, nil
}

func getStoriesAsync(numStories int) ([]item, error) {
	var client hn.Client
	ids, err := client.GetItems()
	if err != nil {
		return nil, err
	}

	type result struct {
		idx  int
		item item
		err  error
	}

	resultChan := make(chan result)
	for i, id := range ids {
		go func(i, id int) {
			story, err := client.GetItem(id)
			if err != nil {
				resultChan <- result{err: err, idx: i}
			} else {
				resultChan <- result{item: parseItem(story), idx: i}
			}
		}(i, id)
		if i >= numStories {
			break
		}
	}

	var results []result
	for i := 0; i < numStories; i++ {
		results = append(results, <-resultChan)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].idx < results[j].idx
	})

	var stories []item
	for _, res := range results {
		if res.err != nil {
			continue
		}
		if isStoryLink(res.item) {
			stories = append(stories, res.item)
		}
	}

	return stories, nil
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
		newItem.Host = strings.TrimPrefix(uri.Host, "www.")
	}
	return newItem
}

func isStoryLink(hnItem item) bool {
	return hnItem.Type == "story" && hnItem.URL != ""
}
