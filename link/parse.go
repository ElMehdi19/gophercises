package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link to be parsed from html tree
type Link struct {
	href      string
	innerText string
}

// Parse to be used to parse html string
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%+v\n", doc)
	dfs(doc, "")
	return []Link{}, nil
}

func dfs(node *html.Node, padding string) {
	if node.Type == html.ElementNode {
		fmt.Printf("%s<%s>\n", padding, node.Data)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		dfs(child, padding+"  ")
	}
}
