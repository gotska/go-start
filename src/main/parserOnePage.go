package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"

	"log"

)

func getElementById(id string, n *html.Node) (element *html.Node, ok bool) {
	for _, a := range n.Attr {
		if a.Key == "class" && a.Val == id {
			return n, true
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if element, ok = getElementById(id, c); ok {
			return
		}
	}
	return
}

func main() {
	resp, err := http.Get("http://korrespondent.net/ukraine/3733657-kuda-zakhodyt-konflykt-lutsenko-y-sytnyka")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	root, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	element, ok := getElementById("post-item__title", root)
	if !ok {
		log.Fatal("element not found")
	}

	fmt.Println(element.FirstChild.Data)

	log.Fatal("element missing value")
}

