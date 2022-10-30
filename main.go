package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var raw = `
<!DOCTYPE html>
<html>
  <body>
    <h1>My First Heading</h1>
      <p>My first paragraph.</p>
      <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
      <img src="foo.jpg" width="104" height="142">
      <img src="bar.png" width="104" height="142">
      <img src="baz.webm" width="104" height="142">
      <div>
        <img src="f00.jpeg" width="104" height="142">
        <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Obcaecati, dicta?</p>
      </div>
  </body>
</html>`

func main() {
	doc, err := html.Parse(bytes.NewReader([]byte(raw)))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse: %s", err)
		os.Exit(1)
	}

	words, images := countWordsAndImages(doc)

	fmt.Printf("%d words, and %d images", words, images)
}

func countWordsAndImages(doc *html.Node) (int, int) {
	var words, images int

	visit(doc, &words, &images)

	return words, images
}

func visit(n *html.Node, words, pics *int) {

	if n.Type == html.TextNode {
		*words += len(strings.Fields(n.Data))
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		*pics++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, words, pics)
	}
}
