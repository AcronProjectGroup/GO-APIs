package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}
type Person struct {
	name, city string
}

func processItems(items ...interface{}) {
	for _, item := range items {
		switch value := item.(type) {
		case ProductZ1:
			fmt.Println("ProductZ1:", value.name, "Price:", value.price)
		case *ProductZ1:
			fmt.Println("ProductZ1 Pointer:", value.name, "Price:", value.price)
		case ServiceZ1:
			fmt.Println("ServiceZ1:", value.description, "Price:",
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
func main() {
	var expense Expense = &ProductZ1{"Kayak", "Watersports", 275}
	data := []interface{}{
		expense,
		ProductZ1{"Lifejacket", "Watersports", 48.95},
		ServiceZ1{"Boat Cover", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}
	processItems(data...)
}
