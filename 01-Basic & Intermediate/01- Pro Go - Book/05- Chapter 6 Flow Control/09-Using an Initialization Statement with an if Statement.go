
// The most common use for this feature is to initialize a variable that 
// is subsequently used in the expression

package main

import (
	"fmt"
	"strconv"
)

func main() {
	priceString := "275"

	// this is scope of main()
	kayakPrice := false
	fmt.Println(kayakPrice)

	// this is scope of if statement
		if kayakPrice, err := strconv.Atoi(priceString); err == nil {
			fmt.Println("Price:", kayakPrice)
		} else {
			fmt.Println("Error:", err)
		}
}


