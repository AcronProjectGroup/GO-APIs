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


