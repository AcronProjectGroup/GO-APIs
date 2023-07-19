package main
import (
	"fmt"
	"strconv"
)
func main() {
	var price []rune = []rune("€48.95")
	
	var currency string = string(price[0])
	
	var amountString string = string(price[1:])
	
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

Length: 6
Currency: €
Amount: 48.95

the rune type is an alias for int32, 
which means that printing out a rune
value will display the numeric value 
used to represent the character.



*/