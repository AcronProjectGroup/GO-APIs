package main

import (
	"fmt"
	"strconv"
)

func main() {
	var price string = "$48.95"

	var currency byte = price[0]
	fmt.Println("byte currency:", currency)

	var amountString string = price[1:]
	fmt.Println("amount String:", amountString)



	amount, parseErr := strconv.ParseFloat(amountString, 64)

	fmt.Println("Currency:", currency)

	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}

}

/* Output:

byte currency: 36
amount String: 48.95
Currency: 36
Amount: 48.95

*/