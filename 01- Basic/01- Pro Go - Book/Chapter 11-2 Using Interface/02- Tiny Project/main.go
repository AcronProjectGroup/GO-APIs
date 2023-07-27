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

type Account struct {
	accountNumber int
	expenses      []Expense
}

func main() {


	/*
	This example creates a Product struct value, assigns it to an Expense variable, alters the value of the
	struct value’s price field, and writes out the field value directly and via an interface method. Compile and
	execute the code
	The Product value was copied when it was assigned to the Expense variable, which means that the
	change to the price field does not affect the result from the getCost method.
	*/
	product := Product{"Kayak", "Watersports", 275}
	var expense Expense = product
	product.price = 100
	fmt.Println("Product field value:", product.price)
	fmt.Println("Expense method result:", expense.getCost(false))
	// ----------------------------------------------------------------------------------------
	
	
	// A pointer to the struct value can be used when making the assignment to the interface variable
	// var expense2 Expense = &product
	// fmt.Println("Expense method result:", expense2.getCost(false))
	
	
	
	// ----------------------------------------------------------------------------------------
	
	account := Account{
		accountNumber: 12345,
		expenses: []Expense{
			Product{"Kayak", "Watersports", 275},
			Service{"Boat Cover", 12, 89.50},
		},
	}

	// ----------------------------------------------------------------------------------------

	// Using Interface
	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.50},
	}

	// ----------------------------------------------------------------------------------------

	/* Explaination:
	The Account struct has an expenses field whose type is a slice of Expense values, which can be used just
	like any other field. Compile and execute the project, which produces the following output:
	Expense: Kayak Cost: 275
	Expense: Boat Cover Cost: 1074
	Total: 1349
	*/
	for _, expense := range account.expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
	fmt.Println("Total:", calcTotal(account.expenses))

	// ----------------------------------------------------------------------------------------
	for _, expense := range expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
	/* Output:
	Expense: Kayak Cost: 275
	Expense: Boat Cover Cost: 1074
	*/

	// ----------------------------------------------------------------------------------------
	fmt.Println("Total:", calcTotal(expenses))
	/* Ouput :
	Expense: Kayak Cost: 275
	Expense: Boat Cover Cost: 1074
	Total: 1349
	*/

}
