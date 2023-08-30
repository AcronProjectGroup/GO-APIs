/*

Syntax
for statement1; statement2; statement3 {
   // code to be executed for each iteration
}

*/

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

// Result:
// 0
// 1
// 2
// 3
// 4
