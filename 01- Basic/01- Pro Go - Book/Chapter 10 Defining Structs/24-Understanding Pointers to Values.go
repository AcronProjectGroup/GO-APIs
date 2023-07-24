/*

Earlier examples have used pointers in two steps. The first step is to create a value and assign it to a variable,
like this:
...
kayak := Product {
	name: "Kayak",
	category: "Watersports",
	price: 275,
}
...
The second step is to use the address operator to create a pointer, like this:
...
calcTax(&kayak)

*/

// There is no need to assign a struct value to a variable before creating a pointer,
// and the address operator
// can be used directly with the literal struct syntax

package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

func calcTax(product *Product) {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
}



func main() {

	// can be used directly with the literal struct syntax
	kayak := &Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}



	calcTax(kayak)
	fmt.Println("Name:", kayak.name, "Category:",
		kayak.category, "Price", kayak.price)

		
}
