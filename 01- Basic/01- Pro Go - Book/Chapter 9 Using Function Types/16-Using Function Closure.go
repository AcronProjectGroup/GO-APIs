

// The closure feature is the link between the factory function and the calculator function. 
// The calculator function relies on two variables to produce a result.

// The closure feature allows a function to access variables—and parameters—in the surrounding code.

// In this case, the calculator function relies on the parameters of the factory function. 
// When the calculator function is invoked, the parameter values are used to produce a result,


package main

import "fmt"

type calcFunc func(float64) float64

func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}


func priceCalcFactory(threshold, rate float64) calcFunc {
	return func(price float64) float64 {
		if price > threshold {
			return price + (price * rate)
		}
		return price
	}
}


func main() {
	watersportsProducts := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.50,
		"Stadium":     79500,
	}
	waterCalc := priceCalcFactory(100, 0.2)
	soccerCalc := priceCalcFactory(50, 0.1)

	for product, price := range watersportsProducts {
		printPrice(product, price, waterCalc)
	}

	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc)
	}

}


/* Output:

Product: Kayak Price: 330
Product: Lifejacket Price: 48.95
Product: Soccer Ball Price: 19.5
Product: Stadium Price: 87450

*/