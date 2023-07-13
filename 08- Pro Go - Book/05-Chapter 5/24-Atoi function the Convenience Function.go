package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := "100"

	int1, int1err := strconv.Atoi(val1)

	if int1err == nil {
		var intResult int = int1
		fmt.Println("Parsed value:", intResult)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}

}


// The Atoi function accepts only the value to be parsed and doesn’t support parsing nondecimal values. 
// The type of the result is int, instead of the int64 produced by the ParseInt function.



