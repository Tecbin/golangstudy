package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("#feed-main-list > li:nth-child(1) > div > div.z-feed-content > h5 > a.feed-nowrap").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Html())

	})
	// document.querySelector("#feed-main-list > li:nth-child(1) > div > div.z-feed-content > h5 > a:nth-child(2) > div")
	dom.Find("#feed-main-list > li:nth-child(1) > div > div.z-feed-content > h5 > a:nth-child(2) > div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Html())

	})
}
