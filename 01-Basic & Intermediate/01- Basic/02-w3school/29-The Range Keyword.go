/*

Syntax
for index, value := array|slice|map {
   // code to be executed for each iteration
}


*/

package main

import (
	"fmt"
)

func main() {
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
}

// Result:
// 0      apple
// 1      orange
// 2      banana
