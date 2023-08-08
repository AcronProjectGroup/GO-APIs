package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main()  {
	startTime := time.Now()

	var WG sync.WaitGroup

	// JOBS
	jobs := make(chan string, 5)

	// WORKERS
	const workers = 3
	for worker := 1; worker <= workers; worker++ {
		go getWebSite(jobs, &WG, worker)
	}
	
	// Job producer
	file := "website.txt"
	reader, _ := os.Open(file)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		jobs <- scanner.Text()
		WG.Add(1)
	}

	WG.Wait()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("Totall time:", time.Since(startTime))
}


func getWebSite(websites chan string, WG *sync.WaitGroup, worker int) {
	for website := range websites {
		if res, err := http.Get(website); err != nil {
			fmt.Printf("%s is down, processed by Worker %d.\n", website, worker)
		} else {
			fmt.Printf("[%d] %s	is up. processed by Worker %d.\n", res.StatusCode, website, worker)
		}
		WG.Done()
	}
}

/* Output:

https://youtube.com/ is down, processed by Worker 1.
https://facebook.com/ is down, processed by Worker 2.
https://twitter.com/ is down, processed by Worker 1.
https://instagram.com/ is down, processed by Worker 1.
[403] https://google.com/       is up. processed by Worker 3.
https://whatsapp.com/ is down, processed by Worker 3.
[200] https://www.wikipedia.org/        is up. processed by Worker 2.
[200] https://us.yahoo.com/     is up. processed by Worker 1.
[200] https://www.amazon.com/   is up. processed by Worker 3.
[200] https://outlook.live.com/owa/     is up. processed by Worker 1.
https://netflix.com/ is down, processed by Worker 1.
[200] https://zoom.us/  is up. processed by Worker 2.
[200] https://greensock.com/gsap/       is up. processed by Worker 3.
-----------------------------------------------------
Totall time: 3.702024911s

*/