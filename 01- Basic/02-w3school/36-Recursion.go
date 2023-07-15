// Recursion is a common mathematical and programming concept.
// This has the benefit of meaning that you can loop through data to reach a result.

package main

import (
	"fmt"
)

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	fmt.Println(factorial_recursion(4))
}

// Result:
// 24
