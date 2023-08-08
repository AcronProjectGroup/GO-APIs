/*

Channels can be safely shared between multiple goroutines, and the effect of the changes made in this
section is that the Go routines created to invoke the TotalPrice method all send their results through the
channel created by the CalcStoreTotal function, where they are received and processed.

*/


// Sending a Result Using a Channel


package main

import (
		"fmt"
		"time"
	)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64)
	for category, group := range data {
		go group.TotalPrice(category, channel)
	}
	for index := 0; index < len(data); index++ {
		// The arrow is placed before the channel to receive a value from it
		storeTotal += <- channel
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}


func (group ProductGroup) TotalPrice(category string, resultChannel chan float64) {
	var total float64
	for _, p := range group {
		fmt.Println(category, "product:", p.Name)
		total += p.Price
		time.Sleep(time.Second)
	}
	fmt.Println(category, "subtotal:", ToCurrency(total))
	resultChannel <- total
}
