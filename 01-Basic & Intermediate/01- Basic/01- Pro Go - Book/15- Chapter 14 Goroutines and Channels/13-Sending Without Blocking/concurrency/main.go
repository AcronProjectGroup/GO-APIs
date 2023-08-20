// 
// 

package main

import (
	"fmt"
	"time"
)

func enumerateProducts(channel chan<- *Product) {
	for _, p := range ProductList {
		select {
		case channel <- p:
			fmt.Println("Sent product:", p.Name)
		default:
			fmt.Println("Discarding product:", p.Name)
			time.Sleep(time.Second)
		}
	}
	close(channel)
}


func main() {
	productChannel := make(chan *Product, 5)
	go enumerateProducts(productChannel)
	time.Sleep(time.Second)
	for p := range productChannel {
		fmt.Println("Received product:", p.Name)
	}
}

/*

Output:

Sent product: Kayak
Sent product: Lifejacket
Sent product: Soccer Ball
Sent product: Corner Flags
Sent product: Stadium
Discarding product: Thinking Cap
Discarding product: Unsteady Chair
Received product: Kayak
Received product: Lifejacket
Received product: Soccer Ball
Received product: Corner Flags
Received product: Stadium
Sent product: Bling-Bling King
Received product: Bling-Bling King

*/
