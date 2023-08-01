// Struct values are comparable if all their fields can be compared.

package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
	}
	
	p1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p2 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p3 := Product{name: "Kayak", category: "Boats", price: 275.00}
	
	fmt.Println("p1 == p2:", p1 == p2)
	fmt.Println("p1 == p3:", p1 == p3)
}


// The struct values p1 and p2 are equal because all their fields are equal.

// The struct values p1 and p3 are not equal 
// because the values assigned to their category fields are different.


/* Output:

p1 == p2: true
p1 == p3: false

*/