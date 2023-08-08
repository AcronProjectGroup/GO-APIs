
// Defining a Channel

// To address this issue, Go provides channels, which are conduits through which data can be sent and
// received. I am going to introduce a channel into the example in steps

package main

import (
		"fmt"
		"time"
	)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	// Channels are strongly typed, which means that they will carry values of a specified type or interface.
	var channel chan float64 = make(chan float64)
	for category, group := range data {
		go group.TotalPrice(category)
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}


func (group ProductGroup) TotalPrice(category string) (total float64) {
	for _, p := range group {
		fmt.Println(category, "product:", p.Name)
		total += p.Price
		time.Sleep(time.Second)
	}
	fmt.Println(category, "subtotal:", ToCurrency(total))
	return
}
