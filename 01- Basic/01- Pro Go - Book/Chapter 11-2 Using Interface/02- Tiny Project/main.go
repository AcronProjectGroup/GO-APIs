package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCost(true)
	}
	return
}


func main() {
	// Using Interface
	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.50},
	}
	for _, expense := range expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
	
	fmt.Println("Total:", calcTotal(expenses))

}

/* Output:

Expense: Kayak Cost: 275
Expense: Boat Cover Cost: 1074

*/
