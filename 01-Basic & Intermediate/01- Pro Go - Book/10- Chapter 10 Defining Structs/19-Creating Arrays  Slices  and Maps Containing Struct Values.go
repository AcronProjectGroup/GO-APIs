// Creating Arrays, Slices, and Maps Containing Struct Values

package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
		//otherNames []string
	}

	type StockLevel struct {
		Product
		Alternate Product
		count     int
	}

	array := [1]StockLevel{
		{
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	fmt.Println("Array:", array[0].Product.name)

	slice := []StockLevel{
		{
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
		{
			Product:   Product{"Benze", "Mercidess", 123230},
			Alternate: Product{"Unumaq", "U700", 43123},
			count: 2,
		},
	}
	fmt.Println("Slice:", slice[0].Product.name)

	for index, Value := range slice {
		fmt.Println("Slice-Index:", index, "All-Values:", Value)
	}

	kvp := map[string]StockLevel{
		"kayak": {
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}

	fmt.Println("Map:", kvp["kayak"].Product.name)
}
