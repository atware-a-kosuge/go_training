package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n != nil {
		if n.Type == html.ElementNode {
			if n.Data == "a" || n.Data == "link" {
				for _, a := range n.Attr {
					if a.Key == "href" {
						links = append(links, "<"+n.Data+">"+a.Val)
					}
				}
			}

			if n.Data == "img" || n.Data == "script" {
				for _, a := range n.Attr {
					if a.Key == "src" {
						links = append(links, "<"+n.Data+">"+a.Val)
					}
				}
			}
		}
		links = visit(links, n.FirstChild)
		links = visit(links, n.NextSibling)
	}
	return links
}
