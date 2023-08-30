/*

Go allows the user of the empty interface
which means an interface that defines no methods to represent
any type, which can be a useful way
to group disparate types that share no common features

*/

package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

type Person struct {
	name, city string
}

func main() {

	var expense1 Expense = &Product{"Kayak", "Watersports", 275}
	
	data := []interface{}{
		expense1,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}
	for _, item := range data {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price)
		case Service:
			fmt.Println("Service:", value.description, "Price:",
				value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default:", value)
		}
	}
}

/* Output:

Product Pointer: Kayak Price: 275
Product: Lifejacket Price: 48.95
Service: Boat Cover Price: 1074
Person: Alice City: London
Person Pointer: Bob City: New York
Built-in type: This is a string
Built-in type: 100
Built-in type: true

*/