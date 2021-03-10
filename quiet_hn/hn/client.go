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

// Item hacker news item
type Item struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"Time"`
	Title       string `json:"title"`
	Type        string `json:"type"`

	// Only one can exist
	// Url for story
	// Text for comment, reply...
	URL  string `json:"url"`
	Text string `json:"text"`
}

// GetItem returns an hacker new Item
func (c *Client) GetItem(id int) (Item, error) {
	c.defaultify()
	var item Item
	resp, err := http.Get(fmt.Sprintf("%s/item/%d.json", c.apiBase, id))
	if err != nil {
		return item, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&item); err != nil {
		return item, err
	}

	return item, nil
}
