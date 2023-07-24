// Assigning a struct to a new variable or using a struct as a function parameter
// creates a new value that copies
// the field values

package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
	}
	
	p1 := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	
	p2 := p1
	
	// in this line don't change into p2 because p1 is another place in memory address
	p1.name = "Original Kayak"
	
	fmt.Println("P1:", p1.name)
	fmt.Println("P2:", p2.name)
}

/* Output:

P1: Original Kayak
P2: Kayak


*/
