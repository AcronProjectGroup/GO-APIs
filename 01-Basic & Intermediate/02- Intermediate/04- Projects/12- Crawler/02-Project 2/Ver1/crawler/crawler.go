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
func (c *Crawler) Crawl(url string) {
	// Make HTTP request
	// Use goquery to parse HTML and extract data and links
	// Send data to indexer
}
