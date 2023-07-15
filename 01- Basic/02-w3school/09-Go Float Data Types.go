/*

Type		Size		Range
-------		--------	---------------------
float32 	32 bits 	-3.4e+38 to 3.4e+38.
float64 	64 bits 	-1.7e+308 to +1.7e+308.

*/

package main

import (
	"fmt"
)

func main() {
	var x float32 = 123.78
	var y float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)
}




