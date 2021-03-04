package hn

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUp() (string, func()) {
	mux := http.NewServeMux()
	items := `[
		1, 2, 3, 4
	]`
	mux.HandleFunc("/topstories.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, items)
	})

	item := `{
		"by": "Jax Teller",
		"descendants": 66,
		"id": 1,
		"kids": [26346957, 26346970, 26347248],
		"score": 19,
		"time": 1614876580,
		"title": "SAMCRO FOR LIFE",
		"type": "story",
		"url": "https://mehdi.codes"
	}`
	mux.HandleFunc("/item/1.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, item)
	})

	server := httptest.NewServer(mux)
	return server.URL, func() {
		server.Close()
	}
}

func TestGetItems(t *testing.T) {
	server, tearDown := setUp()
	defer tearDown()

	client := Client{apiBase: server}
	ids, err := client.GetItems()
	if err != nil {
		t.Errorf("client.GetItems() error: %s", err.Error())
	}

	if len(ids) != 4 {
		t.Errorf("want %d; got %d", 4, len(ids))
	}
}

func TestGetItem(t *testing.T) {
	server, tearDown := setUp()
	defer tearDown()

	client := Client{apiBase: server}
	item, err := client.GetItem(1)

	if err != nil {
		t.Errorf("client.GetItem() error: %s", err.Error())
	}

	want := "Jax Teller"
	if item.By != want {
		t.Errorf("want %s; got %s", want, item.By)
	}

	want = "SAMCRO FOR LIFE"
	if item.Title != want {
		t.Errorf("want %s; got %s", want, item.Title)
	}
}
