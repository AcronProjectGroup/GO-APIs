// Methods can define parameters and results, just like regular functions

package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}


func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}


func (product *Product) printDetails() {
	fmt.Println(
		"Name:", product.name, 
		"Category:", product.category,
		"Price:", product.calcTax(0.2, 100),
	)
}


func main() {
	products := []*Product{
		newProduct("Kayak", "Watersports", 275),
		newProduct("Lifejacket", "Watersports", 48.95),
		newProduct("Soccer Ball", "Soccer", 19.50),
	}

	for _, p := range products {
		p.printDetails()
	}

}

/* Output:

Name: Kayak Category: Watersports Price: 330
Name: Lifejacket Category: Watersports Price: 48.95
Name: Soccer Ball Category: Soccer Price: 19.5

*/
