// Struct values can be defined without using names,
// as long as the types of the values correspond to the order
// in which fields are defined by the struct type

// Omitting Field Names
package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
	}
	
	var kayak = Product{"Kayak", "Watersports", 275.00}
	
	fmt.Println("Name:", kayak.name)
	fmt.Println("Category:", kayak.category)
	fmt.Println("Price:", kayak.price)
}

/*

The literal syntax used to define the struct value contains just values, 
which are assigned to the struct fields in the order in which they are specified.

Output:
Name: Kayak
Category: Watersports
Price: 275


*/
