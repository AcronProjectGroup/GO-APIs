// A constructor function is responsible
// for creating struct values using values received through parameters

// Defining a Constructor Function
package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}


func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}



func main() {
	
	products := [2]*Product{
		newProduct("Kayak", "Watersports", 275),
		newProduct("Hat", "Skiing", 42.50),
	}


	for _, p := range products {
		fmt.Println("Name:", p.name, "Category:", p.category, "Price", p.price)
	}

	
}
