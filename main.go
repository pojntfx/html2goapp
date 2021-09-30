package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func Traverse(doc *html.Node) {
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.TextNode && strings.TrimSpace(node.Data) != "" {
			fmt.Println("Text:", node.Data)
		}

		if node.Type == html.ElementNode {
			fmt.Println("Element:", node.DataAtom.String(), node.Attr)
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}

	crawler(doc)
}

func main() {
	root, err := html.Parse(strings.NewReader(input))
	if err != nil {
		panic(err)
	}

	Traverse(root)
}

const input = `<div class="asdf">
	<h1>Hey!</h2>
	<div>
		<p>asdfasdfasdf</p>
		<p class="yay!">Hmmm ...</p>
	</div>
</div>`
