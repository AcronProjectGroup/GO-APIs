/*
A method whose receiver is a pointer type can also be invoked through a regular value of the underlying
type, meaning that a method whose type is *Product, for example, can be used with a Product value
*/
package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

//	type Supplier struct {
//			 name, city string
//	}
//
//	func newProduct(name, category string, price float64) *Product {
//			 return &Product{ name, category, price }
//	}

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

// func (supplier *Supplier) printDetails() {
//		 fmt.Println("Supplier:", supplier.name, "City:", supplier.city)
// }

func main() {
	
	// The kayak variable is assigned a Product value but is used with the printDetails method, whose
	// receiver is *Product.
	// Go takes care of the mismatch and invokes the method seamlessly.
	kayak := Product{"Kayak", "Watersports", 275}
	kayak.printDetails()


	
}
