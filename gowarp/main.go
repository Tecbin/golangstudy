package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("")
	/*	url :=

		res, err := http.NewRequest("GET", url, nil)
		res.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3776.0 Safari/537.36")
		client := &http.Client{Timeout: time.Second * 5}
		resp, err := client.Do(req)
	*/

	if err != nil {
		log.Fatal(err)
	}

	//    defer resp.Body.Close()
	//  data, err := ioutil.ReadAll(resp.Body

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	/*	doc.Find("ul#feed-main-list").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			band := s.Find("li>div>div>h5>a").Text()
			title := s.Find("div#z-highlight").Text()
			fmt.Printf("Review %d: %s - %s\n", i, band, title)
		})
	*/

	b := doc.Find("#feed-main-list > li:nth-child(1) > div > div.z-feed-content > h5 > a.feed-nowrap").Text()
	// document.querySelector("#feed-main-list > li:nth-child(1) > div > div.z-feed-content > h5 > a:nth-child(2) > div")
	p := doc.Find("#feed-main-list > li:nth-child(1) > div > div.z-feed-content > h5 > a:nth-child(2) > div").Text()
	fmt.Printf("Review %s:%s\n", b, p)

}
func main() {

	for i := 0; i < 1000; i = i + 1 {
		ExampleScrape()
		fmt.Printf("dd:%d \n", i)
		r := rand.New(rand.NewSource(time.Now().Unix()))
		time1 := 180 + r.Intn(100)
		time.Sleep(time.Duration(time1) * time.Second)

	}
}
