package main

import (
	"fmt"
	"strconv"
)

func main() {
	val := 275

	base10String := strconv.Itoa(val)

	base2String := strconv.FormatInt(int64(val), 2)

	fmt.Println("Base 10: " + base10String)
	fmt.Println("Base 2: " + base2String)

}


// The Itoa function accepts an int value, which is explicitly converted to an int64 
// and passed to the ParseInt function





