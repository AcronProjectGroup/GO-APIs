// The comparison operators compare two values,
// returning the bool value true if they are the same and false otherwise.

/*

==	This operator returns true if the operands are equal.
!=	This operator returns true if the operands are not equal.
<	This operator returns true if the first operand is less than the second operand.
>	This operator returns true if the first operand is greater than the second operand.
<=	This operator returns true if the first operand is less than or equal to the second operand.
>=	This operator returns true if the first operand is greater than or equal to the second operand.

*/

package main

import (
	"fmt"
	// "math"
)

func main() {
	first := 100
	const second = 200.00
	equal := first == second
	fmt.Println(equal)
	
	notEqual := first != second
	fmt.Println(notEqual)
	
	lessThan := first < second
	fmt.Println(lessThan)
	
	lessThanOrEqual := first <= second
	fmt.Println(lessThanOrEqual)
	
	greaterThan := first > second
	fmt.Println(greaterThan)
	
	greaterThanOrEqual := first >= second
	fmt.Println(greaterThanOrEqual)
	
}
