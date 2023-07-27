package main

import "fmt"

type Expense interface {
	getName2() string
	getCost2(annual bool) float64
}

func main() {
	expenses := []Expense{
		Service3{"Boat Cover", 12, 89.50, []string{}},
		Service3{"Paddle Protect", 12, 8, []string{}},
	}

	for _, expense := range expenses {

		// fmt.Println(expense)

		s := expense.(Service3)
		fmt.Println(
			"Service3:", s.description,
			"Price:", s.monthlyFee*float64(s.durationMonths),
		)

		}


}
