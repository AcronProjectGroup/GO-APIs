package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := "0"
	bool1, b1err := strconv.ParseBool(val1)
	if b1err == nil {
		fmt.Println("Parsed value:", bool1)
	} else {
		fmt.Println("Cannot parse", val1)
	}

	val2 := "01"
	bool1, b1err = strconv.ParseBool(val2)
	if b1err == nil {
		fmt.Println("Parsed value:", bool1)
	} else {
		fmt.Println("Cannot parse", val2)
	}

}



