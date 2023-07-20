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

}


/* Output:

Product: Kayak Supplier: Acme Kayaks
Product: Kayak Supplier: Bob's Boats
Product: Kayak Supplier: Crazy Canoes
Product: Lifejacket Supplier: Sail Safe Co


*/