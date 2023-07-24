// Custom data types are defined using the Go structs feature

package main

import "fmt"

func main() {
	// Define Product Structure(struct)
	type Product struct {
		name, category string
		price          float64
	}
	

	// Stored in variable
	kayak := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}


	// Use it with dot kayak.category
	fmt.Println(kayak.name, kayak.category, kayak.price)
	
	
	// Change value for the variable 
	kayak.price = 300
	fmt.Println("Changed price:", kayak.price)
}
