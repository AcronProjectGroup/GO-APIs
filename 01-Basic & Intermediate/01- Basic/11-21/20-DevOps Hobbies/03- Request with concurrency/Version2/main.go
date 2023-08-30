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

	file := "website.txt"
	reader, _ := os.Open(file)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		go getWebSite(scanner.Text(), &WG)
		WG.Add(1)
	}

	WG.Wait()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("Totall time:", time.Since(startTime))
}


func getWebSite(website string, WG *sync.WaitGroup) {
	if res, err := http.Get(website); err != nil {
		fmt.Println("[***]",website, "	is down.")
	} else {
		fmt.Printf("[%d] %s	is up.\n", res.StatusCode, website)
	}

	WG.Done()
}

