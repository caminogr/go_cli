package scrape

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

func GetURLs() []string {
	// Set base url
	baseURL := ("http://hoge.com")

	doc, err := goquery.NewDocument(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	urls := []string{}
	path := ""
	// Set element need to scrape
	doc.Find("ol li a").Each(func(i int, s *goquery.Selection) {
		path, _ = s.Attr("href")

		urls = append(urls, baseURL + path)
	})

	return urls
}
