/*

but for this chapter, there is a single code file containing a single package, which
means that methods can be defined only for types defined in the main.go file.
But this doesn’t limit methods to just structs, because the type keyword can be used to create aliases to
any type, and methods can be defined for the alias.

*/

// Defining a Method for a Type Alias
package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

type ProductList []Product


func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}

func main() {
	products := ProductList{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}


	for category, total := range products.calcCategoryTotals() {
		fmt.Println("Category: ", category, "Total:", total)
	}

}
