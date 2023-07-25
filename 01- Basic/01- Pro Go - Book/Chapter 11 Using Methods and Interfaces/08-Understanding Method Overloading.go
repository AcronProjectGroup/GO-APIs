/*

Go does not support method overloading, where multiple methods can be defined with the same name
but different parameters. Instead, each combination of method name and receiver type must be unique,
regardless of the other parameters that are defined.

*/

// Methods with the Same Name

package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

type Supplier struct {
	name, city string
}

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (product *Product) printDetails() {
	fmt.Println("Name:", product.name, "Category:", product.category,
		"Price", product.calcTax(0.2, 100))
}

func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}

func (supplier *Supplier) printDetails() {
	fmt.Println("Supplier:", supplier.name, "City:", supplier.city)
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

	suppliers := []*Supplier{
		{"Acme Co", "New York City"},
		{"BoatCo", "Chicago"},
	}

	for _, s := range suppliers {
		s.printDetails()
	}

}

/* Output:

Name: Kayak Category: Watersports Price 330
Name: Lifejacket Category: Watersports Price 48.95
Name: Soccer Ball Category: Soccer Price 19.5
Supplier: Acme Co City: New York City
Supplier: BoatCo City: Chicago

*/