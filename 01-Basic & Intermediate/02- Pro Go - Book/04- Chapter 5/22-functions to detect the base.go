package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := "0b1100100"
	int1, int1err := strconv.ParseInt(val1, 0, 8)
	if int1err == nil {
		smallInt := int8(int1)
		fmt.Println("Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}

/*

Prefix		Description
------		------------------------------------------------------
  0b		 This prefix denotes a binary value, such as 0b1100100.
  0o		 This prefix denotes an octal value, such as 0o144.
  0x		 This prefix denotes a hex value, such as 0x64.

*/

