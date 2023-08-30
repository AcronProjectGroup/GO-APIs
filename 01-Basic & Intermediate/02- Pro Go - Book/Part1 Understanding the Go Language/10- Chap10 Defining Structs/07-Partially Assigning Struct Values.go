// Values do not have to be provided for all fields when creating a struct value

package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
	}
	
	kayak := Product{
		name:     "Kayak",
		category: "Watersports",
	}
	
	fmt.Println(kayak.name, kayak.category, kayak.price)
	
	kayak.price = 300
	
	fmt.Println("Changed price:", kayak.price)
}

/*

No initial value is provided for the price field for the struct assigned to the kayak variable. 
When no field is provided, the zero value for the field’s type is used.

As the output shows, omitting an initial value 
doesn’t prevent a value from being assigned to a field subsequently.
The zero types are assigned to all fields if you define a struct-typed variable 
but don’t assign a value to it

*/


// Output:
// Kayak Watersports 0
// Changed price: 300