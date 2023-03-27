// strconv Package 

// Source: FreeCodCamp YouTube Channel

package main

import (
	"fmt"
)

func main() {
	var number1 int = 40
	fmt.Printf("%v, %T\n", number1, number1)



	var stringNumber string
	stringNumber = string(number1)
	fmt.Printf("%v, %T\n", stringNumber, stringNumber)

	

}