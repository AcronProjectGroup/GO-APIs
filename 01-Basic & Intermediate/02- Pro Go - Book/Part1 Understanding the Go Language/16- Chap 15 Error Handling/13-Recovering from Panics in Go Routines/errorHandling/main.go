/*
The main function uses a goroutine to invoke the processCategories function, which panics if the
TotalPriceAsync function sends an error.
*/

package main

import (
	"fmt"
)

type CategoryCountMessage struct {
	Category string
	Count    int
}

func processCategories(categories []string, outChan chan<- CategoryCountMessage) {
	defer func() {
		if arg := recover(); arg != nil {
			fmt.Println(arg)

			// close(outChan)

		}
	}()
	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			outChan <- CategoryCountMessage{
				Category: message.Category,
				Count:    int(message.Total),
			}
		
			// close(outChan)

		} else {
			panic(message.CategoryError)
			
		}
	}
	close(outChan)
}

func main() {

	categories := []string{"Watersports", "Chess", "Running"}

	channel := make(chan CategoryCountMessage)
	go processCategories(categories, channel)
	for message := range channel {
		fmt.Println(message.Category, "Total:", message.Count)
	}

}

/* Output:

Watersports Total: 328
Chess Total: 1291
Cannot find category: Running
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [chan receive]:
main.main()
        C:/errorHandling/main.go:39 +0x1c5
exit status 2


 */
