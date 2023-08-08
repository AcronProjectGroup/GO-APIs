// The zero value for a struct type is a struct value whose fields are assigned their zero type.
// The zero value for a pointer to a struct is nil

// Examining Zero Types

package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

func main() {
	var prod Product
	
	var prodPtr *Product
	
	fmt.Println("Value:", prod.name, prod.category, prod.price)
	
	fmt.Println("Pointer:", prodPtr)

}

/* Output:

Value:   0
Pointer: <nil>

*/
