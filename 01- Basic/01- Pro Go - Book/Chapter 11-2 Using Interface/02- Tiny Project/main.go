package main

import "fmt"

type Expense interface {
	getName1() string
	getCost1(annual bool) float64
}

func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCost1(true)
	}
	return
}

type Account struct {
	accountNumber int
	expenses      []Expense
}

func main() {


	/*
	This example creates a Product1 struct value, assigns it to an Expense variable, alters the value of the
	struct value’s price field, and writes out the field value directly and via an interface method. Compile and
	execute the code
	The Product1 value was copied when it was assigned to the Expense variable, which means that the
	change to the price field does not affect the result from the getCost1 method.
	*/
	product := Product1{"Kayak", "Watersports", 275}
	var expense Expense = product
	product.price = 100
	fmt.Println("Product1 field value:", product.price)
	fmt.Println("Expense method result:", expense.getCost1(false))
	// ----------------------------------------------------------------------------------------
	
	
	// A pointer to the struct value can be used when making the assignment to the interface variable
	// var expense2 Expense = &product
	// fmt.Println("Expense method result:", expense2.getCost1(false))
	
	
	
	// ----------------------------------------------------------------------------------------
	
	account := Account{
		accountNumber: 12345,
		expenses: []Expense{
			Product1{"Kayak", "Watersports", 275},
			Service1{"Boat Cover", 12, 89.50},
		},
	}

	// ----------------------------------------------------------------------------------------

	// Using Interface
	expenses := []Expense{
		Product1{"Kayak", "Watersports", 275},
		Service1{"Boat Cover", 12, 89.50},
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
		fmt.Println("Expense:", expense.getName1(), "Cost:", expense.getCost1(true))
	}
	fmt.Println("Total:", calcTotal(account.expenses))

	// ----------------------------------------------------------------------------------------
	for _, expense := range expenses {
		fmt.Println("Expense:", expense.getName1(), "Cost:", expense.getCost1(true))
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
