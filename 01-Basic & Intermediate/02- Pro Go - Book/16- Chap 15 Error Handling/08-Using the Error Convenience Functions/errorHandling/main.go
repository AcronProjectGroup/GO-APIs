


package main

import (
	"fmt"
)

func main() {
	categories := []string{"Watersports", "Chess", "Running"}

	channel := make(chan ChannelMessage, 10)
	
	
	fmt.Println("1")
	go Products.TotalPriceAsync(categories, channel)
	fmt.Println("2")
	
	
	for message := range channel {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		} else {
			fmt.Println(message.CategoryError)
		}
	}



}

/* Output:

1
2
Watersports Total: $328.95
Chess Total: $1291.00
Running (no such category)


 */
