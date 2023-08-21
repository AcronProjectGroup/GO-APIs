package main

import (
	"fmt"
)

func main() {
	dispatchChannel := make(chan DispatchNotification, 100)
	
	go DispatchOrders(dispatchChannel)
	
	for details := range dispatchChannel {
		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
			"x", details.Product.Name)
	}
	
	fmt.Println("Channel has been closed")
	
	// go DispatchOrders(dispatchChannel)
	// for {
	// 	if details, open := <-dispatchChannel; open {
	// 		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
	// 			"x", details.Product.Name)
	// 	} else {
	// 		fmt.Println("Channel has been closed")
	// 		break
	// 	}
	// }
	
}

/* Output:

Order count: 2
Dispatch to Bob : 8 x Corner Flags
Dispatch to Charlie : 5 x Corner Flags
Channel has been closed

*/


