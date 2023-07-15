// This is still possible when using an initialization statement,
// but you have to ensure that the parentheses are applied
// only to the expression,

package main

import (
	"fmt"
	"strconv"
)

func main() {
	priceString := "275"

	
	//-----------------------------------------------(here      )
	if kayakPrice, err := strconv.Atoi(priceString); (err == nil) {
		fmt.Println("Price:", kayakPrice)
	} else {
		fmt.Println("Error:", err)
	}
}
