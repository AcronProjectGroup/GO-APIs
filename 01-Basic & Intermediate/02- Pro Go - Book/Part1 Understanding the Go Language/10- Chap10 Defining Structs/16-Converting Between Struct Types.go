// A struct type can be converted into any other struct type that has the same fields

// meaning all the fields have
// the same name and type and are defined in the same order

package main

import "fmt"

func main() {
	
	type Product struct {
		name, category string
		price          float64
	}
	
	type Item struct {
		name     string
		category string
		price    float64
	}
	
	prod := Product{name: "Kayak", category: "Watersports", price: 275.00}
	item := Item{name: "Kayak", category: "Watersports", price: 275.00}
	
	fmt.Println("prod == item:", prod == Product(item))

}
