package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main()  {
	startTime := time.Now()

	file := "website.txt"
	reader, _ := os.Open(file)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		getWebSite(scanner.Text())
	}

	fmt.Println("-----------------------------------------------------")
	fmt.Println("Totall time:", time.Since(startTime))
}


func getWebSite(website string) {
	if res, err := http.Get(website); err != nil {
		fmt.Println("[***]",website, "	is down.")
	} else {
		fmt.Printf("[%d] %s	is up.\n", res.StatusCode, website)
	}
}

/* Output:

[403] https://google.com/       		is up.
[***] https://youtube.com/      		is down.
[***] https://facebook.com/     		is down.
[***] https://twitter.com/      		is down.
[200] https://www.wikipedia.org/        is up.
[***] https://instagram.com/    		is down.
[200] https://us.yahoo.com/     		is up.
[***] https://whatsapp.com/     		is down.
[200] https://www.amazon.com/   		is up.
[200] https://zoom.us/  				is up.
[200] https://outlook.live.com/owa/     is up.
[200] https://greensock.com/gsap/       is up.
[***] https://netflix.com/      		is down.
-----------------------------------------------------
Totall time: 1m7.365008608s

*/