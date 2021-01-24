package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link to be parsed from html tree
type Link struct {
	Href      string
	InnerText string
}

// Parse to be used to parse html string
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	links := []Link{}
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
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

func innerText(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}
	if node.Type != html.ElementNode {
		return ""
	}
	var text strings.Builder
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		text.WriteString(innerText(child))
	}
	return text.String()
}

func buildLink(node *html.Node) Link {
	var link Link
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
			break
		}
	}
	link.InnerText = strings.Join(strings.Fields(innerText(node)), " ")
	return link
}
