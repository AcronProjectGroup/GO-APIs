// Syntax
// func FunctionName(param1 type, param2 type) type {
//   // code to be executed
//   return output
// }

package main

import (
	"fmt"
)

func myFunction(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(myFunction(1, 2))
}

// Result:
// 3
