// Formatting Floating-Point Values

// Converting a Floating-Point Number

package main

import (
	"fmt"
	"strconv"
)

func main() {
	val := 49.95

	Fstring := strconv.FormatFloat(val, 'f', 2, 64)
	Estring := strconv.FormatFloat(val, 'e', -1, 64)

	fmt.Println("Format F: " + Fstring)
	fmt.Println("Format E: " + Estring)

}

// Output:
// Format F: 49.95
// Format E: 4.995e+01

// The first argument to the FormatFloat function is the value to process. 
// The second argument is a byte value, which specifies the format of the string. 
// The byte is usually expressed as a rune literal value
// The third argument to the FormatFloat function specifies the number of digits that will follow the decimal point.




// Commonly Used Format Options for Floating-Point String Formatting

/*

Function	Description
--------	--------------------------------
f			The floating-point value will be expressed in the form ±ddd.ddd without an exponent, such as 49.95.

e, E		The floating-point value will be expressed in the form ±ddd.ddde±dd, such as 4.995e+01 or
			4.995E+01. The case of the letter denoting the exponent is determined by the case of the rune
			used as the formatting argument.

g, G		The floating-point value will be expressed using format e/E for large exponents or format f for
			smaller values.


*/







