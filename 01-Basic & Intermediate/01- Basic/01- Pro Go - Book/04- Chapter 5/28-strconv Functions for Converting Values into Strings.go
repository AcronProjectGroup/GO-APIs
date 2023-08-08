// strconv Functions for Converting Values into Strings

/*

Function				Description
---------				---------------------------------
FormatBool(val)			This function returns the string true or false based on the value of the
						specified bool.

FormatInt(val, base)	This function returns a string representation of the specified int64 value,
						expressed in the specified base.
FormatUint(val, base)	This function returns a string representation of the specified uint64 value,
						expressed in the specified base.

FormatFloat(val, format, precision, size)
						This function returns a string representation of the specified float64 value,
						expressed using the specified format, precision, and size.

Itoa(val)				This function returns a string representation of the specified int value,
						expressed using base 10.

*/

// Formatting Boolean Values

package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := true
	val2 := false
	str1 := strconv.FormatBool(val1)
	str2 := strconv.FormatBool(val2)
	fmt.Println("Formatted value 1: " + str1)
	fmt.Println("Formatted value 2: " + str2)
}



