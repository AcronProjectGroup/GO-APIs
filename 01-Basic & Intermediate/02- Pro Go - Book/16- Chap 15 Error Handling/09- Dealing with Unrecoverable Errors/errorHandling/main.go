/*

Instead of printing out a message when a category cannot be found, the main function panics, which is
done using the built-in panic function


The panic function is invoked with an argument, which can be any value that will help explain the
panic.

*/


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
			panic(message.CategoryError)
			// fmt.Println(message.CategoryError)
		}
	}



}

/* Output:

1
2
Watersports Total: $328.95
Chess Total: $1291.00
panic: Cannot find category: Running

goroutine 1 [running]:
main.main()
        /home/ubuntu/Desktop/Repo/Test/errorHandling/main.go:30 +0x352
exit status 2


 */
