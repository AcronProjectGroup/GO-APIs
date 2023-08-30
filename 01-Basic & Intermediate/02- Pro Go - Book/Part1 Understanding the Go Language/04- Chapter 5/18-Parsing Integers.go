package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := "100000231"
	
	int1, int1err := strconv.ParseInt(val1, 0, 10)
	
	if int1err == nil {
		fmt.Println("Parsed value:", int1)
	} else {
		fmt.Println("Cannot parse", val1)
	}
}



