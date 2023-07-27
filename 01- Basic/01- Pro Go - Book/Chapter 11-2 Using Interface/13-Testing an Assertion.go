package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}
	for _, expense := range expenses {
		if s, ok := expense.(Service); ok {
			fmt.Println("Service:", s.description, "Price:",
				s.monthlyFee*float64(s.durationMonths))
		} else {
			fmt.Println("Expense:", expense.getName(),
				"Cost:", expense.getCost(true))
		}
	}
}
