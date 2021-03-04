package hn

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiBase = "https://hacker-news.firebaseio.com/v0"

// Client quiet hn client
type Client struct {
	apiBase string
}

func (c *Client) defaultify() {
	if c.apiBase == "" {
		c.apiBase = apiBase
	}
}

// GetItems will return a slice of ids of top HN items
func (c *Client) GetItems() ([]int, error) {
	c.defaultify()
	resp, err := http.Get(fmt.Sprintf("%s/topstories.json", c.apiBase))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var ids []int
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&ids); err != nil {
		return nil, err
	}
	return ids, nil
}
