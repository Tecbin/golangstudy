package main

import (
	"fmt"

	htmlquery "github.com/antchfx/xquery/htmlquery"
)

func main() {
	doc, err := htmlquery.LoadURL("http://search.smzdm.com/?c=home&s=")
	if err != nil {
		panic(err)
	}
	// Find all news item.
	list, err := htmlquery.querySelector("#feed-main-list > li:nth-child(1) > div > div.z-feed-content > h5 > a:nth-child(2) > div")
	if err != nil {
		panic(err)
	}
	fmt.Printf(list)

}
