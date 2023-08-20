// In this example, the select statement is used to receive values from two channels
// If none of the case statements can be performed, the default clause is executed.

package main

import (
	"fmt"
	"time"
)

func enumerateProducts(channel chan<- *Product) {
	for _, p := range ProductList[:3] {
		channel <- p
		time.Sleep(time.Millisecond * 800)
	}
	close(channel)
}
func main() {
	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
	productChannel := make(chan *Product)
	go enumerateProducts(productChannel)
	openChannels := 2
	for {
		select {
		case details, ok := <-dispatchChannel:
			if ok {
				fmt.Println("Dispatch to", details.Customer, ":",
					details.Quantity, "x", details.Product.Name)
			} else {
				fmt.Println("Dispatch channel has been closed")
				dispatchChannel = nil
				openChannels--
			}
		case product, ok := <-productChannel:
			if ok {
				fmt.Println("Product:", product.Name)
			} else {
				fmt.Println("Product channel has been closed")
				productChannel = nil
				openChannels--
			}
		default:
			if openChannels == 0 {
				goto alldone
			}
			fmt.Println("-- No message ready to be received")
			time.Sleep(time.Millisecond * 500)
		}
	}
alldone:
	fmt.Println("All values received")
}


/*

A nil channel is never ready and will not be chosen, allowing the select statement to move onto other
case statements, whose channels may still be open.



*/