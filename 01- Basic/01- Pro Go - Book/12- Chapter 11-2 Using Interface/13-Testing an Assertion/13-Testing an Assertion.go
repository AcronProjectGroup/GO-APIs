package main

import "fmt"

type Service struct {
	myStr string
	myInt float64
	myFloat float64
	mySlice []string
}

type Product struct {
	myStr string
	mySTR string
	myFloat float64
}

// type Expense interface {
// 	getName() string
// 	getCost(annual bool) float64
// }

func main() {
	expenses := []Service{
		{"Boat Cover", 12, 89.50, []string{}},
		{"Paddle Protect", 12, 8, []string{}},
		// &Product{"Kayak", "Watersports", 275},
	}

	fmt.Println(expenses)

	// for _, expense := range expenses {
	// 	if s, ok := expense.(Service); ok {
	// 		fmt.Println("Service:", s.description, "Price:",
	// 			s.monthlyFee*float64(s.durationMonths))
	// 	} else {
	// 		fmt.Println("Expense:", expense.getName(),
	// 			"Cost:", expense.getCost(true))
	// 	}
	// }
}
