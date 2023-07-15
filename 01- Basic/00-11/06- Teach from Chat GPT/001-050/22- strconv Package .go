// strconv Package

// Source: FreeCodCamp YouTube Channel
// Notion: https://sinalalenakhsh.notion.site/11-strconv-Package-1f900fdae9434070939c662bd7cc970e

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number1 int = 40
	fmt.Printf("%v, %T\n", number1, number1)
	var stringNumber string
	stringNumber = strconv.Itoa(number1)
	fmt.Printf("%v, %T\n", stringNumber, stringNumber)

	//------------------------
	//------------------------
	//------------------------

	var number2 string = "40"
	fmt.Printf("%v, %T\n", number2, number2)

	var stringNumber2 int
	stringNumber2, _ = strconv.Atoi(number2) // Integer to ASCII
	stringNumber2 = stringNumber2 + 1
	fmt.Printf("%v, %T\n", stringNumber2, stringNumber2)

}
