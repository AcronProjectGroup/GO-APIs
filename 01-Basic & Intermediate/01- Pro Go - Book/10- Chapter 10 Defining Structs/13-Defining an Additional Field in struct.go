/*


As noted earlier, field names must be unique with the struct type,
which means that you can define only one embedded field for a specific type.
If you need to define two fields of the same type, then you will need
to assign a name to one of them

*/

package main

import "fmt"

func main() {
	
	type Product struct {
		name, category string
		price          float64
	}

	// The StockLevel type has two fields whose type is Product, 
	// but only one can be an embedded field.
	type StockLevel struct {
		Product
		Alternate Product
		count     int
	}
	
	stockItem := StockLevel{
		Product:   Product{"Kayak", "Watersports", 275.00},
		Alternate: Product{"Lifejacket", "Watersports", 48.95},
		count:     100,
	}
	
	fmt.Println("Name:", stockItem.Product.name)
	fmt.Println("Alt Name:", stockItem.Alternate.name)
}
