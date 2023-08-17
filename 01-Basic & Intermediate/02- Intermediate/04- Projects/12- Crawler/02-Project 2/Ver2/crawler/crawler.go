package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type Crawler struct {
	// other fields
}

func NewCrawler( /* parameters */ ) *Crawler {
	return &Crawler{
		// initialization
	}
}

func (c *Crawler) Crawl(url string) (*goquery.Document, error) {
	// Fetch the URL
	res, err := c.fetch(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Parse the page with goquery
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	// Extract the page text
	pageText := doc.Find("body").Text()

	// Find all links in the page
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			// Handle the link (e.g., send to a channel or add to a list)
			// Remember to make sure the link is absolute, not relative!
		}
	})
	

	// Return the parsed page
	return doc, nil
}
