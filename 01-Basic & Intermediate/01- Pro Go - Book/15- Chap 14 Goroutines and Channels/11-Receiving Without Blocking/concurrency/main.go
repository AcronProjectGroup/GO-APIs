// A select statement has a similar structure to a switch statement
// except that the case statements are channel operations.



package main

import (
	"fmt"
	"time"
)

// func receiveDispatches(channel <-chan DispatchNotification) {
// 	for details := range channel {
// 		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
// 			"x", details.Product.Name)
// 	}
// 	fmt.Println("Channel has been closed")
// }

func main() {
	dispatchChannel := make(chan DispatchNotification, 100)
	// var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	// var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel
	go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
	// receiveDispatches((<-chan DispatchNotification)(dispatchChannel))

	for {
		select {
		case details, ok := <-dispatchChannel:
			if ok {
				fmt.Println("Dispatch to", details.Customer, ":",
					details.Quantity, "x", details.Product.Name)
			} else {
				fmt.Println("Channel has been closed")
				goto alldone
			}
		case details, ok := <-dispatchChannel:
			if !ok {
				fmt.Println(details, "*************************")
				goto alldone
			}
		default:
			fmt.Println("-- No message ready to be received")
			time.Sleep(time.Millisecond * 500)
		}
	}
alldone:
	fmt.Println("All values received")

}

/* Output:

-- No message ready to be received
Order count: 7
Dispatch to Charlie : 7 x Stadium
-- No message ready to be received
Dispatch to Alice : 4 x Kayak
-- No message ready to be received
-- No message ready to be received
-- No message ready to be received
Dispatch to Bob : 5 x Corner Flags
-- No message ready to be received
-- No message ready to be received
-- No message ready to be received
-- No message ready to be received
-- No message ready to be received
-- No message ready to be received
{ <nil> 0} *************************
All values received


 */
