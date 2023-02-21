// Source : A Tour of Go 

package main

import (
	"fmt"
	"strings"
)

func main() {
	vari := strings.Fields("  foo bar  baz   ")
	typeVari := len(vari)
	fmt.Printf("Type length fields are: %T\n", typeVari)
}



// package main

// func WordCount(s string) map[string]int {
// 	return map[string]int{"x": 1}
// }

// func main() {
// 	wc.Test(WordCount)
// }

