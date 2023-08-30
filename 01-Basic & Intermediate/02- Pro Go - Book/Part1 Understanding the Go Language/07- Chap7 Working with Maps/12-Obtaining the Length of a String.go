package main
import (
	"fmt"
	"strconv"
)
func main() {
	var price string = "€48.95"
	
	var currency string = string(price[0])
	
	var amountString string = price[1:]
	
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	
	fmt.Println("Length:", len(price))
	
	fmt.Println("Currency:", currency)
	
	
	if (parseErr == nil) {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}

}

/* Output:

Length: 8
Currency: â
Parse Error: strconv.ParseFloat: parsing "\x82\xac48.95": invalid syntax

*/