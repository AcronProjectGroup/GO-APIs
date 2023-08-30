// The Go comparison operators cannot be used to compare functions
// but they can be used to determine
// whether a function has been assigned to a variable

// Checking for Assignment

package main

import "fmt"

func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}


func calcWithoutTax(price float64) float64 {
	return price
}


func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	for product, price := range products {
		var calcFunc func(float64) float64

		// The zero value for function types is nil
		fmt.Println("Function assigned:", calcFunc == nil)

		if price > 100 {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}

		fmt.Println("Function assigned:", calcFunc == nil)

		totalPrice := calcFunc(price)
		fmt.Println("Product:", product, "Price:", totalPrice)
	}

}
/* Output:

Function assigned: true
Function assigned: false
Product: Kayak Price: 330
Function assigned: true
Function assigned: false
Product: Lifejacket Price: 48.95

*/