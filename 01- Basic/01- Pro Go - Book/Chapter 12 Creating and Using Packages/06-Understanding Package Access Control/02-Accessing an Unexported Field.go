package main

import (
	"fmt"
	"UsingCustomPackage05/packages/store"
)

func main() {
	product := store.Product{
		Name:     "Kayak",
		Category: "Watersports",
		// price:    279,
	}

	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	// fmt.Println("Price:", product.price)
}

/*

The first change attempts to set a value for 
the price field when using the literal syntax to create a Product value.
The second change attempts to read the value of the price field.


Output:

	.\main.go:13:9: cannot refer to unexported field 'price' in struct literal of type
	store.Product
	.\main.go:18:34: product.price undefined (cannot refer to unexported field or method price)



To resolve these errors, 
I can either export the price field or export methods or functions that provide
access to the field value.


defines a constructor function for creating Product values and methods
for getting and setting the price field.

*/
