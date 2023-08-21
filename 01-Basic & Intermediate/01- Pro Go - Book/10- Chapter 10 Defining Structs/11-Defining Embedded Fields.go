// If a field is defined without a name, it is known as an embedded field,
// and it is accessed using the name of its type

// Defining Embedded Fields
package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
	}
	
	type StockLevel struct {
		Product
		count int
		float64
	}
	
	stockItem := StockLevel{
		Product: Product{"Kayak", "Watersports", 275.00},
		count:   100,
		float64: 123,
	}
	fmt.Println("Name:", stockItem.Product.name)
	fmt.Println("Count:", stockItem.count)
	fmt.Println("float64:", stockItem.float64)
}


/*

The StockLevel struct type has two fields. 
The first field is embedded and is defined just using a type,
which is the Product struct type


*/
