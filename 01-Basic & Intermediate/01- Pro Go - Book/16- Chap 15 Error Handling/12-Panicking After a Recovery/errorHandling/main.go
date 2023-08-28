/*
Selectively Panicking After a Recovery in the main.go File in the errorHandling Folder
*/

package main

import (
	"fmt"
)

func main() {


	defer func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println("Error:", err.Error())
				panic(err)
			} else if str, ok := arg.(string); ok {
				fmt.Println("Message:", str)
			} else {
				fmt.Println("Panic recovered")
			}
		}
	}()

	
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
		}
	}

}

/* Output:

1
2
Watersports Total: $328.95
Chess Total: $1291.00
Error: Cannot find category: Running
panic: Cannot find category: Running [recovered]
        panic: Cannot find category: Running

goroutine 1 [running]:
main.main.func1()
        /home/ubuntu/Desktop/Repo/Test/errorHandling/main.go:18 +0x1b3
panic({0x488820?, 0xc000014080?})
        /usr/local/go/src/runtime/panic.go:914 +0x21f
main.main()
        /home/ubuntu/Desktop/Repo/Test/errorHandling/main.go:40 +0x385
exit status 2

*/
