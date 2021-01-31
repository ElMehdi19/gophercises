package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/ElMehdi19/gophercises/link"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "`Page URL` to build a sitemap for")
	flag.Parse()

	hrefs := get(*urlFlag)
	for _, href := range hrefs {
		fmt.Println(href)
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
