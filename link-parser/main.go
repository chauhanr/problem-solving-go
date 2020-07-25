package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var htmlPath string

func init() {
	flag.StringVar(&htmlPath, "p", "", "To specify the path to html")
}

func main() {
	flag.Parse()
	if htmlPath == "" {
		fmt.Errorf("The path to html file cannot be empty")
		return
	}
	f, err := os.Open(htmlPath)
	if err != nil {
		fmt.Errorf("Unable to open file %s, error: %s", htmlPath, err)
		return
	}

	r := bufio.NewReader(f)
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Errorf("Unable to parse %s, error: %s", htmlPath, err)
		return
	}
	links := ExtractLinks(doc)
	fmt.Printf("Links: %v\n", links)
}

type Link struct {
	Href string `json:"Href"`
	Text string `json:"Text"`
}

func ExtractLinks(doc *html.Node) []Link {
	var dfs func(*html.Node)
	var links []Link
	dfs = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			attrs := n.Attr
			link := Link{}
			for _, attr := range attrs {
				if attr.Key == "href" {
					link.Href = attr.Val
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				s := extractText(c)
				if s != "" {
					link.Text = link.Text + s
				}
			}
			link.Text = strings.TrimSpace(link.Text)

			if link.Href != "" || link.Text != "" {
				links = append(links, link)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			dfs(c)
		}
	}
	dfs(doc)
	return links
}

func extractText(n *html.Node) string {
	if n != nil && n.Type == html.TextNode {
		return n.Data
	}
	c := n.FirstChild
	if c != nil && c.Type == html.TextNode {
		return c.Data
	}
	return ""
}
