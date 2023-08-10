package main

import (
	"fmt"
)

func main() {
	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChannel)
	for {
		details := <-dispatchChannel
		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
			"x", details.Product.Name)
	}

}

// The Go runtime will terminate the program if all the
// goroutines are blocked, which you can see by compiling and executing the project

/* Output:
Order count: 4
Dispatch to Charlie : 9 x Lifejacket
Dispatch to Charlie : 6 x Kayak
Dispatch to Alice : 2 x Corner Flags
Dispatch to Charlie : 3 x Kayak
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /home/ubuntu/Desktop/Repo/Test/concurrency/main.go:11 +0x87
exit status 2

*/

