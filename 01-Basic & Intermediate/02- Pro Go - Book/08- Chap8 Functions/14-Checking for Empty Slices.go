package main

import "fmt"

func printSuppliers(product string, suppliers ...string) {

	if len(suppliers) == 0 {
		fmt.Println("Product:", product, "Supplier: (none)")
	} else {
		for _, supplier := range suppliers {
			fmt.Println("Product:", product, "Supplier:", supplier)
		}
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
Product: Soccer Ball Supplier: (none)

*/