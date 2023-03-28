package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	jobTitle := getInput()
	searchJob(jobTitle)
}

func getInput() string {
	var jobTitle string
	fmt.Print("Enter a job title to search: ")
	fmt.Scanln(&jobTitle)
	return jobTitle
}

func searchJob(jobTitle string) {
	baseURL := "https://de.indeed.com"
	searchPath := "/jobs"
	searchURL := fmt.Sprintf("%s%s?q=%s", baseURL, searchPath, url.QueryEscape(jobTitle))

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Error: status code %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	results := doc.Find(".jobsearch-SerpJobCard")

	if results.Length() == 0 {
		fmt.Println("No results found.")
		return
	}

	fmt.Printf("%d results found:\n\n", results.Length())

	results.Each(func(i int, s *goquery.Selection) {
		title := s.Find(".jobtitle").Text()
		company := s.Find(".company").Text()
		location := s.Find(".location").Text()
		link, _ := s.Find(".jobtitle").Attr("href")
		link = fmt.Sprintf("%s%s", baseURL, link)

		fmt.Printf("%d) %s\n   Company: %s\n   Location: %s\n   Link: %s\n\n", i+1, title, company, location, link)
	})
}
