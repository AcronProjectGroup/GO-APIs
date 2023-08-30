// The opposite process is also true so
// that a method that receives a value can be invoked using a pointer

package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

func (product Product) printDetails() {
	fmt.Println("Name:", product.name, "Category:", product.category,
		"Price", product.calcTax(0.2, 100))
}

func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}

func main() {
	kayak := &Product{"Kayak", "Watersports", 275}
	kayak.printDetails()
}

/*
This feature means that you can write methods based on how you want them to behave, using pointers
to avoid value copying or to allow the receiver to be modified by a method.

One effect of this feature is that value and pointer types are considered 
the same when it comes to method overloading, 
meaning that a method named printDetails whose receiver type 
is Product will conflict with a printDetails method whose receiver type is *Product.

*/
