// Notion Source: https://sinalalenakhsh.notion.site/12-First-Project-d828d312c2e44551bffa7c4c1589979c

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// تعریف آدرس سایت کاریابی آلمان
	url := "https://www.arbeitsagentur.de/en/find-job/jobsearch?execution=e2s1&_eventId=reset_searchForm&toggleOffertentyp=stellenangebote&page=0&size=10&facets=%7B%22distance%22%3A%5B%2210%22%5D%7D&anwendungsbereich=en&ort=Dortmund&plz=&radius=10&beruf=Direktvermittlung"

	// دریافت صفحه HTML سایت کاریابی
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// خواندن محتوای صفحه HTML
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// چاپ محتوای صفحه HTML برای بررسی
	fmt.Println(string(body))
}
