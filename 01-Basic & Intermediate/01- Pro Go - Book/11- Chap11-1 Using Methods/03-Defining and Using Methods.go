package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}


/*

The printDetails function receives a pointer to a Product, which it uses to write out the value of the
name, category, and price fields. The key point for this section is the way that the printDetails function is
invoked:
...
printDetails(p)
...


The function name is followed by arguments enclosed in parentheses.

*/
func printDetails(product *Product) {
	fmt.Println("Name:", product.name, "Category:", product.category,
		"Price", product.price)
}


func main() {
	products := []*Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}

	for _, p := range products {
		printDetails(p)
	}

}
