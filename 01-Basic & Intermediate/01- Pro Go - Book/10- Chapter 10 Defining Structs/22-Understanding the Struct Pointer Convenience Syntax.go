// pointers are required to ensure that structs are not needlessly
// duplicated and that changes made by functions affect the values received as parameters




// This code works, but it is hard to read, 
// especially when there are multiple references in the same block
// of code, such as the body of the calcTax method.
package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

func calcTax(product *Product) {
	if (*product).price > 100 {
		(*product).price += (*product).price * 0.2
	}
}


func main() {
	kayak := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}

	calcTax(&kayak)

	fmt.Println("Name:", kayak.name, "Category:",
		kayak.category, "Price", kayak.price)

}








