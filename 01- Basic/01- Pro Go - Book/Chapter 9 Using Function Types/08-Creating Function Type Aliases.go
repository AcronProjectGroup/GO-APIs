// Using a Type Alias

// The alias is created with the type keyword, followed by a name for the alias and then by the type

// they can simplify code and make the use of a
// specific function signature easier to identify.



/*
	The alias name can be used instead of the function type, like this:
	...
	func selectCalculator(price float64) calcFunc {
	...
*/


package main

import "fmt"

type calcFunc func(float64) float64


func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}

func calcWithoutTax(price float64) float64 {
	return price
}

func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func selectCalculator(price float64) calcFunc {
	if price > 100 {
		return calcWithTax
	}
	return calcWithoutTax
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
