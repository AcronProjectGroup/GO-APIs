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
		&Product2 { "Kayak", "Watersports", 275 },
	}

	for _, expense := range expenses {
		if s, ok := expense.(Service3); ok {
			fmt.Println(
				"Service3:", s.description, "Price:",
				s.monthlyFee*float64(s.durationMonths),
			)
		} else {
			fmt.Println(
				"Expense:", expense.getName2(),
				"Cost:", expense.getCost2(true),
			)
		}
	}


}



/*

The bool value can be used with an if statement to execute statements for a specific dynamic type.
Compile and execute the project; 

you will see the following output:

Service: Boat Cover Price: 1074
Service: Paddle Protect Price: 96
Expense: Kayak Cost: 275


*/