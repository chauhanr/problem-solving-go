package main

import (
	"bufio"
	"os"
	"testing"

	"golang.org/x/net/html"
)

var testcases = []struct {
	Path  string
	Links []Link
}{
	{
		Path: "ex1.html",
		Links: []Link{
			{
				Href: "https://www.twitter.com/joncalhoun",
				Text: "Check me out on twitter",
			},
			{
				Href: "https://github.com/gophercises",
				Text: "Gophercises is on Github!",
			},
		},
	},
	{
		Path: "ex3.html",
		Links: []Link{
			{
				Href: "/dog-cat",
				Text: "dog cat",
			},
		},
	},
	{
		Path: "ex2.html",
		Links: []Link{
			{
				Href: "#",
				Text: "Login",
			},
			{
				Href: "/lost",
				Text: "Lost? Need help?",
			},
			{
				Href: "https://twitter.com/marcusolsson",
				Text: "@marcusolsson",
			},
		},
	},
}

func TestExtractLinks(t *testing.T) {

	for _, tc := range testcases {
		f, err := os.Open(tc.Path)
		if err != nil {
			t.Errorf("Error finding file %s, error: %s\n", tc.Path, err)
		}
		r := bufio.NewReader(f)
		doc, err := html.Parse(r)
		if err != nil {
			t.Errorf("Error parsing file %s, error: %s\n", tc.Path, err)
		}

		links := ExtractLinks(doc)
		if len(links) != len(tc.Links) {
			t.Errorf("Expected %d links but got %d as result of Extract Link method\n", len(tc.Links), len(links))
		} else {
			for index, aln := range links {
				eln := tc.Links[index]
				if eln.Href != aln.Href {
					t.Errorf("Expected Href: %s but got %s\n", eln.Href, aln.Href)
				}
				if eln.Text != aln.Text {
					t.Errorf("Expected Text: %s but got %s\n", eln.Text, aln.Text)
				}
			}
		}

	}
}
