package main

import "fmt"

func printSuppliers(product string, suppliers []string) {

	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}

}


func main() {

	printSuppliers("Kayak", []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"})

	printSuppliers("Lifejacket", []string{"Sail Safe Co"})

}

/* Output:

Product: Kayak Supplier: Acme Kayaks
Product: Kayak Supplier: Bob's Boats
Product: Kayak Supplier: Crazy Canoes
Product: Lifejacket Supplier: Sail Safe Co

*/

/*

The second parameter defined by the printSuppliers function accepts 
a variable number of suppliers
using a string slice. 
This works, but it can be awkward 
because it requires slices to be constructed even
when only a single string is required

*/