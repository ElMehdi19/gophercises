package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/ElMehdi19/gophercises/link"
)

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls []loc `xml:"url"`
}

const header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "`Page URL` to build a sitemap for")
	maxDepth := flag.Int("depth", 3, "BFS max depth")
	flag.Parse()

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", " ")
	var toXML urlset
	hrefs := bfs(*urlFlag, *maxDepth)
	for _, href := range hrefs {
		toXML.Urls = append(toXML.Urls, loc{href})
	}

	fmt.Print(header)
	if err := enc.Encode(toXML); err != nil {
		log.Fatal(err)
	}
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	links, _ := link.Parse(resp.Body)
	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()
	hrefs := getHrefs(links, base)
	return filterHrefs(hrefs, base)
}

func getHrefs(links []link.Link, base string) []string {
	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}
	return hrefs
}

func filterHrefs(hrefs []string, base string) []string {
	var filtred []string
	for _, href := range hrefs {
		if strings.HasPrefix(href, base) {
			filtred = append(filtred, href)
		}
	}
	return filtred
}

type empty struct{}

func bfs(urlStr string, maxDepth int) []string {
	visited := map[string]empty{}
	var current map[string]empty
	var next = map[string]empty{
		urlStr: empty{},
	}

	for i := 0; i <= maxDepth; i++ {
		current, next = next, map[string]empty{}
		for url := range current {
			if _, ok := visited[url]; ok {
				continue
			}
			visited[url] = empty{}
			for _, href := range get(url) {
				next[href] = empty{}
			}
		}
	}

	var urlList []string
	for href := range visited {
		urlList = append(urlList, href)
	}
	return urlList
}
