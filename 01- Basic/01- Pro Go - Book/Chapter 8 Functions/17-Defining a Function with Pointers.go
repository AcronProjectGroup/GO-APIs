package main

import "fmt"

func swapValues(first, second *int) {
	fmt.Println("Before swap:", *first, *second)

	temp := *first

	*first = *second

	*second = temp

	fmt.Println("After swap:", *first, *second)
}



func main() {
	val1, val2 := 10, 20
	fmt.Println("Before calling function", val1, val2)
	swapValues(&val1, &val2)
	fmt.Println("After calling function", val1, val2)
}

/* Output:

Before calling function 10 20
Before swap: 10 20
After swap: 20 10
After calling function 20 10

*/