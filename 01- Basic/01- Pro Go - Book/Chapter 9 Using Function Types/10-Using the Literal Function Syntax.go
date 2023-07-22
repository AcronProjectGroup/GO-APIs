package main

import "fmt"

type calcFunc func(float64) float64

//	func calcWithTax(price float64) float64 {
//		 return price + (price * 0.2)
//	}
//
//	func calcWithoutTax(price float64) float64 {
//		 return price
//	}



func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}



func selectCalculator(price float64) calcFunc {
	if price > 100 {
		var withTax calcFunc = func(price float64) float64 {
			return price + (price * 0.2)
		}
		return withTax
	}
	withoutTax := func(price float64) float64 {
		return price
	}
	return withoutTax
}





func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}
}

/*  
	Go does not support arrow functions, where functions are expressed more concisely using the =>
	operator, without the func keyword and a code block surrounded by braces. In Go, functions must always be
	defined with the keyword and a body.
*/


/*
	Literal functions can also be used with the short variable declaration syntax:
	withoutTax := func (price float64) float64 {
	    return price
	}
*/


