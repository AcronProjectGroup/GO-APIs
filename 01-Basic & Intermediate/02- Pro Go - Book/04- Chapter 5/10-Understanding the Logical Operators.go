/*

Operator Description
||	This operator returns true if either operand is true. If the first operand is true, then the second
		operand will not be evaluated.

&&	This operator returns true if both operands are true. If the first operand is false, then the
		second operand will not be evaluated.

!	This operator is used with a single operand. It returns true if the operand is false and false if
		the operand is true.


*/

package main

import (
	"fmt"
	// "math"
)

func main() {
	maxMph := 50
	
	passengerCapacity := 4
	
	airbags := true
	
	familyCar := passengerCapacity > 2 && airbags
	
	sportsCar := maxMph > 100 || passengerCapacity == 2
	
	canCategorize := !familyCar && !sportsCar
	
	fmt.Println(familyCar) // true
	
	fmt.Println(sportsCar) // false
	
	fmt.Println(canCategorize) // false
}














