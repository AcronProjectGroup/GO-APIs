// In Go, you can name the return values of a function.

package main

import (
	"fmt"
)

func myFunction(x int, y int) (result int) {
	result = x + y
	return
}

func main() {
	fmt.Println(myFunction(1, 2))
}

// Result:
// 3
