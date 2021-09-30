package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func Traverse(doc *html.Node) string {
	var crawler func(*html.Node) string
	crawler = func(node *html.Node) string {
		rv := ""
		tag := ""

		if node.Type == html.TextNode && strings.TrimSpace(node.Data) != "" {
			rv += node.Data
		}

		if node.Type == html.ElementNode {
			tag = node.DataAtom.String()
			rv += "<" + tag + ">"
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			rv += crawler(child)
		}

		if tag == "" {
			return rv
		}

		return rv + "</" + tag + ">"
	}

	return crawler(doc)
}

func main() {
	root, err := html.Parse(strings.NewReader(input))
	if err != nil {
		panic(err)
	}

	fmt.Println(Traverse(root))
}

const input = `<div class="asdf">
	<h1>Hey!</h2>
	<div>
		<p>asdfasdfasdf</p>
		<p class="yay!">Hmmm ...</p>
	</div>
</div>`
