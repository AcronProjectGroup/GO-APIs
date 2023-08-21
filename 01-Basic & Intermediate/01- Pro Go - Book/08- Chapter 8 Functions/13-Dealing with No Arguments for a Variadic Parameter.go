package main

import "fmt"

func printSuppliers(product string, suppliers ...string) {

	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}

}


func main() {

	printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")

	printSuppliers("Lifejacket", "Sail Safe Co")

	printSuppliers("Soccer Ball")

}

/* Output:

Product: Kayak Supplier: Acme Kayaks
Product: Kayak Supplier: Bob's Boats
Product: Kayak Supplier: Crazy Canoes
Product: Lifejacket Supplier: Sail Safe Co


*/


/* Explaination:

Go uses nil as the parameter value, which can cause problems with code
that assumes there will be at least one value in the slice.

There is no output for the Soccer Ball product 
because nil slices have zero length, so the for loop is
never executed.


*/