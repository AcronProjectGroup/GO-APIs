/*
Type assertions can be applied only to interfaces, and they are used to tell the compiler that
an interface value has a specific dynamic type. Type conversions can be applied only to specific types,
not interfaces, and only if the structure of those types is compatible, such as converting between struct
types that have the same fields.
*/
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
		s := expense.(Service)
		fmt.Println("Service:", s.description, "Price:",
			s.monthlyFee*float64(s.durationMonths))
	}
}

/*  Output:
panic: interface conversion: main.Expense is *main.Product, not main.Service

The Go runtime has tried to perform the assertion and failed. 
To avoid this issue, there is a special form
of type assertion that indicates whether an assertion could be performed

*/