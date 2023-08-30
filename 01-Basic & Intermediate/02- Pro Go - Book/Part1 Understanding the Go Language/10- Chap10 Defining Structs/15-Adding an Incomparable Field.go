// Structs cannot be compared if the struct type defines fields with incomparable types,
// such as slices

package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
		otherNames     []string
	}
	
	p1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p2 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p3 := Product{name: "Kayak", category: "Boats", price: 275.00}
	
	// the Go comparison operator cannot be applied to slices, which means that
	// Product values cannot be compared.
	fmt.Println("p1 == p2:", p1 == p2)
	fmt.Println("p1 == p3:", p1 == p3)
}

/* Output: 

./15-Adding an Incomparable Field.go:21:27: invalid operation: p1 == p2 (struct containing []string cannot be compared)
./15-Adding an Incomparable Field.go:22:27: invalid operation: p1 == p3 (struct containing []string cannot be compared)

*/