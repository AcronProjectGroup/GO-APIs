// Each clause in an if statement has its own scope, 
// which means that variables can be accessed only within
// the clause in which they are defined.


package main

import "fmt"

func main() {
	kayakPrice := 275.00
	if kayakPrice > 500 {
		scopedVar := 500
		fmt.Println("Price is greater than", scopedVar)
	} else if kayakPrice < 100 {
		scopedVar := "Price is less than 100"
		fmt.Println(scopedVar)
	} else {
		scopedVar := false
		fmt.Println("Matched: ", scopedVar)
	}
}

