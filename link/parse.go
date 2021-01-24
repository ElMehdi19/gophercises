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
	nodes := linkNodes(doc)
	for i, node := range nodes {
		// checking out Node attributes
		fmt.Printf("Attrs for link #%d: %s\n", i, node.Attr)
	}
	return []Link{}, nil
}

func linkNodes(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	nodes := []*html.Node{}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		nodes = append(nodes, linkNodes(child)...)
	}
	return nodes
}
