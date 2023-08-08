package main

import (
		"fmt"
		"time"
	)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	for category, group := range data {
		go group.TotalPrice(category)
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}


func (group ProductGroup) TotalPrice(category string) (total float64) {
	for _, p := range group {
		fmt.Println(category, "product:", p.Name)
		total += p.Price
		// The new statement adds 100 milliseconds to each iteration of the for loop in the TotalPrice method.
		time.Sleep(time.Second)
	}
	fmt.Println(category, "subtotal:", ToCurrency(total))
	return
}
