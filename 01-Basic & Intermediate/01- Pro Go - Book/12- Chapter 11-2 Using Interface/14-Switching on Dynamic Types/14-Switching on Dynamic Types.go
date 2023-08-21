/*

Go switch statements can be used to access dynamic types

*/

package main

import "fmt"

type Expense interface {
	getName2() string
	getCost2(annual bool) float64
}


func main() {
	for _, expense := range expenses {
		switch value := expense.(type) {
		case Service4:
			fmt.Println("Service4:", value.description, "Price:",
				value.monthlyFee*float64(value.durationMonths))
		case *Product4:
			fmt.Println("Product4:", value.name, "Price:", value.price)
		default:
			fmt.Println("Expense:", expense.getName2(),
				"Cost:", expense.getCost2(true))
		}
	}
}

var expenses = []Expense{
	Service4{"Boat Cover", 12, 89.50, []string{}},
	Service4{"Paddle Protect", 12, 8, []string{}},
	&Product4{"Kayak", "Watersports", 275},
}
