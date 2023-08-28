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

	/*
		The recover built-in function allows a program to manage behavior of a
		panicking goroutine. Executing a call to recover inside a deferred
		function (but not any function called by it) stops the panicking sequence
		by restoring normal execution and retrieves the error value passed to the
		call of panic. If recover is called outside the deferred function it will
		not stop a panicking sequence. In this case, or when the goroutine is not
		panicking, or if the argument supplied to panic was nil, recover returns
		nil. Thus the return value from recover reports whether the goroutine is
		panicking.
		func recover() any
	*/
	defer func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println("Error:", err.Error())
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
