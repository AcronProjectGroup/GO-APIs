package main

import "fmt"

type Expense interface {
	getName2() string
	getCost2(annual bool) float64
}

func main() {
	var e1 Expense = &Product2{name: "Kayak"}
	var e2 Expense = &Product2{name: "Kayak"}

	var e3 Expense = Service2{description: "Boat Cover"}
	var e4 Expense = Service2{description: "Boat Cover"}
	
	fmt.Println("e1 == e2", e1 == e2)
	fmt.Println("e3 == e4", e3 == e4)

}


/* Output:

	e1 == e2 false
	e3 == e4 true

The first two Expense values are not equal. 
That’s because the dynamic type for these values is a pointer
type, and pointers are equal only if they point to the same memory location.


*/